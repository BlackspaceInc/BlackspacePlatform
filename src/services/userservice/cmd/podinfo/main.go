package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/keratin/authn-go/authn"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/api"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/authentication"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/database"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/grpc"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/signals"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/version"
	"github.com/BlackspaceInc/common/circuitbreaker"
	"github.com/BlackspaceInc/common/counters/infrastructure"
	"github.com/BlackspaceInc/common/messaging/rabbitmq"
	"github.com/BlackspaceInc/common/tracing"
)

func main() {
	// flags definition for the service
	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)
	fs.Int("PORT", 9100, "HTTP port")
	fs.Int("port-metrics", 0, "metrics port")
	fs.Int("grpc-port", 0, "gRPC port")
	fs.String("grpc-service-name", "podinfo", "gPRC service name")
	fs.String("level", "info", "log level debug, info, warn, error, flat or panic")
	fs.StringSlice("backend-url", []string{}, "backend service URL")
	fs.Duration("http-client-timeout", 2*time.Minute, "client timeout duration")
	fs.Duration("http-server-timeout", 30*time.Second, "server read and write timeout duration")
	fs.Duration("http-server-shutdown-timeout", 5*time.Second, "server graceful shutdown timeout duration")
	fs.String("data-path", "/data", "data local path")
	fs.String("config-path", "", "config dir path")
	fs.String("config", "config.yaml", "config file name")
	fs.String("ui-path", "./ui", "UI local path")
	fs.String("ui-logo", "", "UI logo")
	fs.String("ui-color", "cyan", "UI color")
	fs.String("ui-message", fmt.Sprintf("greetings from podinfo v%v", version.VERSION), "UI message")
	fs.Bool("random-delay", false, "between 0 and 5 seconds random delay")
	fs.Bool("random-error", false, "1/3 chances of a random response error")
	fs.Int("stress-cpu", 0, "Number of CPU cores with 100 load")
	fs.Int("stress-memory", 0, "MB of data to load into memory")
	fs.String("DATABASE_CONNECTION_ADDRESS", "postgresql://doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db.ondigitalocean.com:25060/user-service-db?sslmode=require", "database connection address")
	fs.String("JWT_SIGNER", "blackspace", "JWT signin name")
	fs.String("AUTH_SERVICE_NAME", "authentication_service", "authentication service name")
	fs.Bool("ENABLE_AUTH_SERVICE_PRIVATE_INTEGRATION", false, "enables communication with authentication service")
	fs.String("AMQP_CONSUMER_QUEUES", "", "set of queue names that this service consumes messages from")
	fs.Int("NUM_CONSUMING_QUEUES", 0, "number of consuming queues")
	fs.String("AMQP_PRODUCER_QUEUES", "email-service:direct,discovery:direct", "set of queue names that this service pushes messages to")
	fs.Int("NUM_PRODUCING_QUEUES", 2, "number of producing queues")
	fs.Bool("IS_PRODUCTION", false, "development or production")
	fs.String("AMQP_SERVER_URL", "amqp://guest:guest@localhost:5672/", "url of the rabbitmq server")
	fs.String("ZIPKIN_SERVER_URL", "http://zipkin:9411", "url of the zipkin server")
	fs.String("SERVICE_NAME", "USER_MANAGEMENT_SERVICE", "service name")
	fs.StringSlice("INTERACTING_SERVICES", []string{"email-service", "authentication-service"}, "the list of services this service will be interacting with")
	fs.String("AUTHN_USERNAME", "blackspaceinc", "username of authentication client")
	fs.String("AUTHN_PASSWORD", "blackspaceinc", "password of authentication client")
	fs.String("AUTHN_ISSUER", "http://localhost", "authentication service issuer")
	fs.String("AUTHN_AUDIENCE", "localhost", "authentication service audience")
	fs.String("AUTHN_PRIVATE_BASE_URL", "http://localhost", "authentication service private url")
	fs.String("AUTHN_PORT", "8404", "authentication service port")
	fs.Bool("ENABLE_CPU_STRESS_TEST", false, "cpu stress testing flag")

	versionFlag := fs.BoolP("version", "v", false, "get version number")

	var (
		mc            *rabbitmq.RabbitMQClient               = nil
		infraCounters *infrastructure.InfrastructureCounters = nil
	)

	// parse flags
	err := fs.Parse(os.Args[1:])
	switch {
	case err == pflag.ErrHelp:
		os.Exit(0)
	case err != nil:
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		fs.PrintDefaults()
		os.Exit(2)
	case *versionFlag:
		fmt.Println(version.VERSION)
		os.Exit(0)
	}

	// bind flags and environment variables
	viper.BindPFlags(fs)
	viper.RegisterAlias("backendUrl", "backend-url")
	hostname, _ := os.Hostname()
	viper.SetDefault("jwt-secret", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	viper.SetDefault("ui-logo", "https://d33wubrfki0l68.cloudfront.net/33a12d8be0bc50be4738443101616e968c7afb8f/cba76/images/scalable.png")
	viper.Set("hostname", hostname)
	viper.Set("version", version.VERSION)
	viper.Set("revision", version.REVISION)
	viper.SetEnvPrefix("PODINFO")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	// load config from file
	if _, err := os.Stat(filepath.Join(viper.GetString("config-path"), viper.GetString("config"))); err == nil {
		viper.SetConfigName(strings.Split(viper.GetString("config"), ".")[0])
		viper.AddConfigPath(viper.GetString("config-path"))
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %v\n", err)
		}
	}

	// configure logging
	logger, _ := initZap(viper.GetString("level"))
	defer logger.Sync()
	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	// initialize authentication client in order to establish communication with the
	// authentication service. This serves as a singular source of truth for authentication needs
	authUsername := viper.GetString("AUTHN_USERNAME")
	authPassword := viper.GetString("AUTHN_PASSWORD")
	issuer := viper.GetString("AUTHN_ISSUER")
	audience := viper.GetString("AUTHN_AUDIENCE")
	privateURL := viper.GetString("AUTHN_PRIVATE_BASE_URL") + ":" + viper.GetString("AUTHN_PORT")
	authn, err := initAuthnClient(authUsername, authPassword, audience, issuer, privateURL)
	// crash the process if we cannot connect to the authentication service
	if err != nil {
		logger.Fatal("failed to initialized authentication service client", zap.Error(err))
	}

	// perform a test request to the authentication service
	_, err = authn.ServerStats()
	if err != nil {
		logger.Fatal("failed to initialized authentication service client", zap.Error(err))
	}

	logger.Info("successfully initialized authentication service client")

	dbConnString := viper.GetString("DATABASE_CONNECTION_ADDRESS")
	// connect to backend database
	db, err := database.New(dbConnString, logger)
	if err != nil {
		logger.Fatal("failed to initialized database connection", zap.Error(err))
	}
	defer db.Engine.Close()

	logger.Info("successfully initialized database connection")

	{
		cpuStressTestEnabled := viper.GetBool("ENABLE_CPU_STRESS_TEST")
		if cpuStressTestEnabled {
			logger.Info("commencing cpu stress test")
			// start stress tests if any
			beginStressTest(viper.GetInt("stress-cpu"), viper.GetInt("stress-memory"), logger)
			logger.Info("completed cpu stress test")
		}
	}

	// validate port
	if _, err := strconv.Atoi(viper.GetString("PORT")); err != nil {
		port, _ := fs.GetInt("PORT")
		viper.Set("PORT", strconv.Itoa(port))
	}

	// load gRPC server config
	var grpcCfg grpc.Config
	if err := viper.Unmarshal(&grpcCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	// start gRPC server
	if grpcCfg.Port > 0 {
		grpcSrv, _ := grpc.NewServer(&grpcCfg, logger)
		go grpcSrv.ListenAndServe()
	}

	// load HTTP server config
	var srvCfg api.Config
	if err := viper.Unmarshal(&srvCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	consumingQueues := viper.GetString("AMQP_CONSUMER_QUEUES")
	numConsumingQueues := viper.GetInt("NUM_CONSUMING_QUEUES")
	producingQueues := viper.GetString("AMQP_PRODUCER_QUEUES")
	numProducingQueues := viper.GetInt("NUM_PRODUCING_QUEUES")
	interactingServices := viper.GetStringSlice("INTERACTING_SERVICES")

	// initialize counters
	infraCounters = infrastructure.New(srvCfg.ServiceName, db.Engine)
	logger.Info("initialized infrastructure level counters")

	// initialize connection to the queues
	// process consuming and producing queues
	queues := parseAndCreateQueueReference(consumingQueues, producingQueues, numProducingQueues, numConsumingQueues, logger)
	// initiate broker connections
	// TODO add retry logic
	mc = initializeMessaging(&srvCfg, queues, logger)
	logger.Info("initialized amqp broker connection")

	// initialize hystrix circuit breaker. this allows us to implement service level
	// api retry logic as well as provides us with a circuit breaker
	initializeTracing(&srvCfg)
	logger.Info("initialized distributed tracing client")

	client := &http.Client{}
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport

	cb := circuitbreaker.NewCircuitBreaker(logger, 3, mc, client)
	cb.ConfigureHystrix(interactingServices)
	logger.Info("initialized and configured circuit breaker")

	// log version and port
	logger.Info("Starting podinfo",
		zap.String("version", viper.GetString("version")),
		zap.String("revision", viper.GetString("revision")),
		zap.String("port", srvCfg.Port),
	)

	// obtain authentication port and service name
	authSrvPort := viper.GetString("AUTHN_PORT")
	duration := viper.GetDuration("http-client-timeout")
	authnHandler := initAuthnCustomHandler(authSrvPort, duration, authUsername, authPassword, cb)
	logger.Info("initialized custom authentication service wrapper client")

	// start HTTP server
	srv, _ := api.NewServer(&srvCfg, logger, db, mc, infraCounters, authnHandler, authn, cb)
	stopCh := signals.SetupSignalHandler()

	configureServiceKeys(srv, logger)
	srv.ListenAndServe(stopCh)
}

// initAuthnCustomHandler initializes connection to custom authentication service wrapper
func initAuthnCustomHandler(authSrvPort string, timeout time.Duration, username, password string, cb *circuitbreaker.CircuitBreaker) *authentication.Authentication {
	authConn := "http://localhost"
	enableAuth := viper.GetBool("ENABLE_AUTH_SERVICE_PRIVATE_INTEGRATION")

	// create a connection wrapper to the authentication service
	auth := authentication.NewAuthenticationService(authConn, authSrvPort, enableAuth, timeout, username, password, cb)
	return auth
}

// configureServiceKeys calls the authentication service and obtains configuration parameters such as jwt signing key
func configureServiceKeys(srv *api.Server, logger *zap.Logger) {
	// ping the authentication service for the public jwt key
	aggErr, jwtConfig := srv.AuthnClient.Handler.GetJwtPublicKey()
	if aggErr != nil {
		logger.Panic("failed to obtain jwt public key uri", zap.Error(aggErr.Error))
	}

	// from the config object should call a get request to the jwks_uri
	// which is the url at which the public key resides
	aggErr, jwtKeys := srv.AuthnClient.Handler.GetJwks(jwtConfig.JwtPublicKeyURI)
	if aggErr != nil {
		logger.Panic("failed to obtain jwt public key", zap.Error(aggErr.Error))
	}

	// extract the public key and store as within the server
	srv.JwtConfig = jwtConfig
	srv.Keys = jwtKeys
	logger.Info("Successfully obtained jwt config parameters from authentication service")
}

// initAuthnClient initializes an instance of the authn client primarily useful in
// communicating with the authentication service securely
func initAuthnClient(username, password, audience, issuer, url string) (*authn.Client, error) {
	// Authentication.
	return authn.NewClient(authn.Config{
		// The AUTHN_URL of your Keratin AuthN server. This will be used to verify tokens created by
		// AuthN, and will also be used for API calls unless PrivateBaseURL is also set.
		Issuer: issuer,

		// The domain of your application (no protocol). This domain should be listed in the APP_DOMAINS
		// of your Keratin AuthN server.
		Audience: audience,

		// Credentials for AuthN's private endpoints. These will be used to execute admin actions using
		// the Client provided by this library.
		//
		// TIP: make them extra secure in production!
		Username: username,
		Password: password,

		// RECOMMENDED: Send private API calls to AuthN using private network routing. This can be
		// necessary if your environment has a firewall to limit public endpoints.
		PrivateBaseURL: url,
	})
}

// Initializes zap logger utility
func initZap(logLevel string) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	switch logLevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	}

	zapEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:       level,
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zapEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapConfig.Build()
}

var stressMemoryPayload []byte

// beginStressTest performs cpu stress testing
func beginStressTest(cpus int, mem int, logger *zap.Logger) {
	done := make(chan int)
	if cpus > 0 {
		logger.Info("starting CPU stress", zap.Int("cores", cpus))
		for i := 0; i < cpus; i++ {
			go func() {
				for {
					select {
					case <-done:
						return
					default:

					}
				}
			}()
		}
	}

	if mem > 0 {
		path := "/tmp/podinfo.data"
		f, err := os.Create(path)

		if err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}

		if err := f.Truncate(1000000 * int64(mem)); err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}

		stressMemoryPayload, err = ioutil.ReadFile(path)
		f.Close()
		os.Remove(path)
		if err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}
		logger.Info("starting CPU stress", zap.Int("memory", len(stressMemoryPayload)))
	}
}

// initialize zipkin connection
func initializeTracing(cfg *api.Config) {
	tracing.InitTracing(cfg.ZipkinServerUrl, cfg.ServiceName)
}

// initializeMessaging initializes connections to an AMQP broker of choice
func initializeMessaging(cfg *api.Config, queues rabbitmq.Queues, logger *zap.Logger) *rabbitmq.RabbitMQClient {
	if cfg.AmqpServerUrl == "" {
		panic("No 'amqp_server_url' set in configuration, cannot start")
	}
	logger.Info("amqp broker details", zap.String("url", cfg.AmqpServerUrl))

	// create new rabbitMQ connection client
	amqpClient, err := rabbitmq.New(cfg.AmqpServerUrl, queues, logger)
	if err != nil {
		logger.Error("error occurred while initiating connection to rabbitmq broker.", zap.Error(err))
		panic(err.Error())
	}

	return amqpClient
}

// parseAndCreateQueueReference creates various rabbitMQ queue object reference
// for consuming and producing queues based on bounded queuename and exchange mappings
func parseAndCreateQueueReference(consumingQueues, producingQueues string, numProducerQueues, numConsumerQueues int, logger *zap.Logger) rabbitmq.Queues {
	// parse both string based on comma seperator
	consumerQueueSet := strings.SplitN(consumingQueues, ",", numConsumerQueues)
	producerQueueSet := strings.SplitN(producingQueues, ",", numProducerQueues)

	logger.Info("consuming queues", zap.Any("queues", consumerQueueSet))
	logger.Info("producing queues", zap.Any("queues", producerQueueSet))

	queueToExchangeMapping := make(map[string]rabbitmq.Exchange)

	// populate the queue to exchange mapping
	populateQueueToExchangeMapping(consumerQueueSet, queueToExchangeMapping, logger)
	populateQueueToExchangeMapping(producerQueueSet, queueToExchangeMapping, logger)

	// initialize a queue from the queue to exchange mapping
	return rabbitmq.InitiateQueues(queueToExchangeMapping)
}

// populateQueueToExchangeMapping creates a queue to exchange mapping stored within a hashmap for both consuming and producing queues
func populateQueueToExchangeMapping(queueSet []string, queueToExchangeMapping map[string]rabbitmq.Exchange, logger *zap.Logger) {
	// from parsed set of queues extract the queue name and exchange type and create a queue instance
	for _, queueName := range queueSet {
		if queueName != "" {
			// parse the queueName
			queueDetails := strings.SplitN(queueName, ":", 2)
			logger.Info("details", zap.Any("queue details", queueDetails))
			// extract actual name of queue and exchange type form queue name reference obtained through env. variables
			name := queueDetails[0]
			exchangeType := queueDetails[1]

			// TODO: Look into the impact of each respective params and enable where fitting
			// bind the queue to the actual exchange of interest
			exchange := rabbitmq.Exchange{
				ExchangeName: name,
				ExchangeType: exchangeType,
				Durable:      true,
				AutoDelete:   false,
				Internal:     false,
				NoWait:       false,
				Args:         nil,
			}

			// store the queue to exchange mapping in a hashmap of interest
			if _, ok := queueToExchangeMapping[queueName]; !ok {
				queueToExchangeMapping[name] = exchange
			}
		}
	}
}

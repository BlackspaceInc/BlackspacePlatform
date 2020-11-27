package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/keratin/authn-go/authn"
	"github.com/sony/gobreaker"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/api"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/authentication"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/grpc"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/logging"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/signals"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/version"
)

func main() {
	// flags definition
	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)
	fs.Int("port", 9898, "HTTP port")
	fs.Int("secure-port", 0, "HTTPS port")
	fs.Int("port-metrics", 0, "metrics port")
	fs.Int("grpc-port", 0, "gRPC port")
	fs.String("grpc-service-name", "authentication_handler_service", "gPRC service name")
	fs.String("level", "info", "log level debug, info, warn, error, flat or panic")
	fs.StringSlice("backend-url", []string{}, "backend service URL")
	fs.Duration("HTTP_CLIENT_TIMEOUT", 2*time.Minute, "client timeout duration")
	fs.Duration("http-server-timeout", 30*time.Second, "server read and write timeout duration")
	fs.Duration("http-server-shutdown-timeout", 5*time.Second, "server graceful shutdown timeout duration")
	fs.String("data-path", "/data", "data local path")
	fs.String("config-path", "", "config dir path")
	fs.String("cert-path", "/data/cert", "certificate path for HTTPS port")
	fs.String("config", "config.yaml", "config file name")
	fs.String("ui-path", "./ui", "UI local path")
	fs.String("ui-logo", "", "UI logo")
	fs.String("ui-color", "#34577c", "UI color")
	fs.String("ui-message", fmt.Sprintf("greetings from authentication_handler_service v%v", version.VERSION), "UI message")
	fs.Bool("h2c", false, "allow upgrading to H2C")
	fs.Bool("random-delay", false, "between 0 and 5 seconds random delay by default")
	fs.String("random-delay-unit", "s", "either s(seconds) or ms(milliseconds")
	fs.Int("random-delay-min", 0, "min for random delay: 0 by default")
	fs.Int("random-delay-max", 5, "max for random delay: 5 by default")
	fs.Bool("random-error", false, "1/3 chances of a random response error")
	fs.Bool("unhealthy", false, "when set, healthy state is never reached")
	fs.Bool("unready", false, "when set, ready state is never reached")
	fs.Int("stress-cpu", 0, "number of CPU cores with 100 load")
	fs.Int("stress-memory", 0, "MB of data to load into memory")
	fs.String("cache-server", "", "Redis address in the format <host>:<port>")
	// authentication service specific flags
	fs.String("AUTHN_USERNAME", "blackspaceinc", "username of authentication client")
	fs.String("AUTHN_PASSWORD", "blackspaceinc", "password of authentication client")
	fs.String("AUTHN_ISSUER", "http://localhost", "authentication service issuer")
	fs.String("AUTHN_DOMAINS", "localhost", "authentication service domains")
	fs.String("AUTHN_PRIVATE_BASE_URL", "http://authentication_service",
		"authentication service private url. should be local host if these are not running on docker containers. " +
		"However if running in docker container with a configured docker network, the url should be equal to the service name")
	fs.String("AUTHN_INTERNAL_PORT", "3000", "authentication service port")
	fs.Bool("ENABLE_AUTH_SERVICE_PRIVATE_INTEGRATION", true, "enables communication with authentication service")
	// logging specific configurations
	fs.String("ENABLE_LOG_TO_STDERR", "false", `feature flag used to toggle on or off wether or not to log to stderr or to log to a
																specific location. the location must be a log file`)
	fs.String("ENABLE_LOG_TO_STDERR_AND_FILES", "false", `feature flag used to toggle on or off wether or not to log to stderr and to log to a
																specific location. the location must be a log file`)
	fs.String("LOG_DIR", "./logs", `the directory at which the application should write
																					log entries.`)
	fs.String("LOG_FILE", "authentication_handler_service.log", `the file at which the application should write
																					log entries.`)
	fs.Int("LOG_LEVEL_VERBOSITY", 3, `number for log level verbosity`)

	versionFlag := fs.BoolP("version", "v", false, "get version number")

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
	viper.SetDefault("ui-logo", "https://raw.githubusercontent.com/github.com/blackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/gh-pages/cuddle_clap.gif")
	viper.Set("hostname", hostname)
	viper.Set("version", version.VERSION)
	viper.Set("revision", version.REVISION)
	viper.SetEnvPrefix("authentication_handler_service")
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

	logging.ConfigureStructuredLogging()
	defer klog.Flush()

	authnServiceClient := NewAuthServiceClient(err)
	klog.Info("successfully initialized authentication service client")

	// start stress tests if any
	beginStressTest(viper.GetInt("stress-cpu"), viper.GetInt("stress-memory"))

	// validate port
	if _, err := strconv.Atoi(viper.GetString("port")); err != nil {
		port, _ := fs.GetInt("port")
		viper.Set("port", strconv.Itoa(port))
	}

	// validate secure port
	if _, err := strconv.Atoi(viper.GetString("secure-port")); err != nil {
		securePort, _ := fs.GetInt("secure-port")
		viper.Set("secure-port", strconv.Itoa(securePort))
	}

	// validate random delay options
	if viper.GetInt("random-delay-max") < viper.GetInt("random-delay-min") {
		klog.Fatal("`--random-delay-max` should be greater than `--random-delay-min`")
	}

	switch delayUnit := viper.GetString("random-delay-unit"); delayUnit {
	case
		"s",
		"ms":
		break
	default:
		klog.Fatal("`random-delay-unit` accepted values are: s|ms")
	}

	// load gRPC server config
	var grpcCfg grpc.Config
	if err := viper.Unmarshal(&grpcCfg); err != nil {
		klog.Fatal("config unmarshal failed", "error", err.Error())
	}

	// start gRPC server
	if grpcCfg.Port > 0 {
		grpcSrv, _ := grpc.NewServer(&grpcCfg)
		go grpcSrv.ListenAndServe()
	}

	// load HTTP server config
	var srvCfg api.Config
	if err := viper.Unmarshal(&srvCfg); err != nil {
		klog.Fatal("config unmarshal failed", "error", err.Error())
	}

	// log version and revisions
	klog.Info("Starting authentication_handler_service",
		"version", viper.GetString("version"),
		"revision", viper.GetString("revision"),
		"port", srvCfg.Port)

	// start HTTP server
	srv, _ := api.NewServer(&srvCfg, authnServiceClient)
	stopCh := signals.SetupSignalHandler()
	srv.ListenAndServe(stopCh)
}

func NewAuthServiceClient(err error) *api.AuthServiceClientWrapper {
	// initialize authentication client in order to establish communication with the
	// authentication service. This serves as a singular source of truth for authentication needs
	authUsername := viper.GetString("AUTHN_USERNAME")
	authPassword := viper.GetString("AUTHN_PASSWORD")
	issuer := viper.GetString("AUTHN_ISSUER")
	domains := viper.GetString("AUTHN_DOMAINS")
	authnUrl := viper.GetString("AUTHN_PRIVATE_BASE_URL")
	privateURL := authnUrl + ":" + viper.GetString("AUTHN_INTERNAL_PORT")
	authSrvPort := viper.GetString("AUTHN_PORT")
	duration := viper.GetDuration("HTTP_CLIENT_TIMEOUT")

	authnClient, err := initAuthnClient(authUsername, authPassword, domains, issuer, privateURL)
	// crash the process if we cannot connect to the authentication service
	if err != nil {
		klog.Fatal("failed to initialized authentication service client", "error", err.Error())
	}

	// perform a test request to the authentication service
	_, err = authnClient.ServerStats()
	if err != nil {
		klog.Fatal("failed to connect to authentication service", "error", err.Error())
	}

	klog.Info("successfullly established connection to authentication service")

	authnHandler := initAuthnHandler(authnUrl, authSrvPort, duration, authUsername, authPassword, nil)
	klog.Info("initialized custom authentication service wrapper client")

	// attempt to connect to the authentication service if not then crash process
	return &api.AuthServiceClientWrapper{
		Client:  authnClient,
		Handler: authnHandler,
	}
}

// initAuthnHandler initializes connection to custom authentication service wrapper
func initAuthnHandler(authnUrl, authSrvPort string,
							timeout time.Duration,
							username, password string,
							cb *gobreaker.CircuitBreaker) *authentication.Authentication {
	enableAuth := viper.GetBool("ENABLE_AUTH_SERVICE_PRIVATE_INTEGRATION")

	// create a connection wrapper to the authentication service
	auth := authentication.NewAuthenticationService(authnUrl, authSrvPort, enableAuth, timeout, username, password, cb)
	return auth
}

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
		Development: false,
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

func beginStressTest(cpus int, mem int) {
	done := make(chan int)
	if cpus > 0 {
		klog.Info("starting CPU stress", "cores", cpus)
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
		path := "/tmp/authentication_handler_service.data"
		f, err := os.Create(path)

		if err != nil {
			klog.Error("memory stress failed", "error", err.Error())
		}

		if err := f.Truncate(1000000 * int64(mem)); err != nil {
			klog.Error("memory stress failed", "error", err.Error())
		}

		stressMemoryPayload, err = ioutil.ReadFile(path)
		f.Close()
		os.Remove(path)
		if err != nil {
			klog.Error("memory stress failed", "error", err.Error())
		}
		klog.Info("starting CPU stress", "memory", len(stressMemoryPayload))
	}
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

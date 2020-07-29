// Package BlackSpace Backend API.
//
// This serves as the user's microservice api definition for the BlackSpace Platform
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Yoan Yomba<yoanyombapro@gmail.com.com> http://BlackSpace.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/BlackspaceInc/common/circuitbreaker"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/keratin/authn-go/authn"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"

	"github.com/BlackspaceInc/common/counters/infrastructure"
	"github.com/BlackspaceInc/common/messaging/rabbitmq"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/authentication"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/database"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/fscache"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
	mw "github.com/BlackspaceInc/Backend/user-management-service/pkg/middleware"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

var (
	healthy int32
	ready   int32
	watcher *fscache.Watcher
)

type FluxConfig struct {
	GitUrl    string `mapstructure:"git-url"`
	GitBranch string `mapstructure:"git-branch"`
}

type AuthServiceClientWrapper struct {
	Client  *authn.Client
	Handler *authentication.Authentication
}

type Config struct {
	HttpClientTimeout         time.Duration `mapstructure:"http-client-timeout"`
	HttpServerTimeout         time.Duration `mapstructure:"http-server-timeout"`
	HttpServerShutdownTimeout time.Duration `mapstructure:"http-server-shutdown-timeout"`
	BackendURL                []string      `mapstructure:"backend-url"`
	UILogo                    string        `mapstructure:"ui-logo"`
	UIMessage                 string        `mapstructure:"ui-message"`
	UIColor                   string        `mapstructure:"ui-color"`
	UIPath                    string        `mapstructure:"ui-path"`
	DataPath                  string        `mapstructure:"data-path"`
	ConfigPath                string        `mapstructure:"config-path"`
	Port                      string        `mapstructure:"port"`
	PortMetrics               int           `mapstructure:"port-metrics"`
	Hostname                  string        `mapstructure:"hostname"`
	RandomDelay               bool          `mapstructure:"random-delay"`
	RandomError               bool          `mapstructure:"random-error"`
	JWTSecret                 string        `mapstructure:"jwt-secret"`
	JWTSigningAuthority       string        `mapstructure:"JWT_SIGNER"`
	IsProduction              bool          `mapstructure:"IS_PRODUCTION"`
	AmqpServerUrl             string        `mapstructure:"AMQP_SERVER_URL"`
	ZipkinServerUrl           string        `mapstructure:"ZIPKIN_SERVER_URL"`
	ServiceName               string        `mapstructure:"SERVICE_NAME"`
	Issuer                    string        `mapstructure:"AUTHN_ISSUER"`
	AuthenticationServiceName string        `mapstructure:"AUTH_SERVICE_NAME"`
}

type Server struct {
	router         *mux.Router
	logger         *zap.Logger
	config         *Config
	db             *database.Db
	JwtConfig      *authentication.JwtConfiguration
	Keys           *authentication.JsonKeys
	AmqpClient     *rabbitmq.RabbitMQClient
	Telemetry      *infrastructure.InfrastructureCounters
	AuthnClient    *AuthServiceClientWrapper
	CircuitBreaker *circuitbreaker.CircuitBreaker
}

func NewServer(
	config *Config,
	logger *zap.Logger,
	db *database.Db,
	amqpClient *rabbitmq.RabbitMQClient,
	telemetry *infrastructure.InfrastructureCounters,
	authnHandler *authentication.Authentication,
	authnClient *authn.Client,
	cb *circuitbreaker.CircuitBreaker) (*Server, error) {

	srv := &Server{
		router:     mux.NewRouter(),
		logger:     logger,
		config:     config,
		db:         db,
		AmqpClient: amqpClient,
		Telemetry:  telemetry,
		AuthnClient: &AuthServiceClientWrapper{
			Client:  authnClient,
			Handler: authnHandler,
		},
		CircuitBreaker: cb,
	}

	return srv, nil
}

func (s *Server) registerHandlers() {
	authMw := mw.NewJwtMiddleware(s.config.JWTSigningAuthority, s.config.JWTSecret)
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	s.router.HandleFunc("/", s.indexHandler).HeadersRegexp("User-Agent", "^Mozilla.*").Methods("GET")
	s.router.HandleFunc("/", s.infoHandler).Methods("GET")
	s.router.HandleFunc("/version", s.versionHandler).Methods("GET")
	s.router.HandleFunc("/echo", s.echoHandler).Methods("POST")
	s.router.HandleFunc("/headers", s.echoHeadersHandler).Methods("GET", "POST")
	s.router.HandleFunc("/healthz", s.healthzHandler).Methods("GET")
	s.router.HandleFunc("/readyz", s.readyzHandler).Methods("GET")
	s.router.HandleFunc("/readyz/enable", s.enableReadyHandler).Methods("POST")
	s.router.HandleFunc("/readyz/disable", s.disableReadyHandler).Methods("POST")
	s.router.HandleFunc("/store", s.storeWriteHandler).Methods("POST")
	s.router.HandleFunc("/store/{hash}", s.storeReadHandler).Methods("GET").Name("store")
	s.router.HandleFunc("/configs", s.configReadHandler).Methods("GET")
	s.router.HandleFunc("/api/info", s.infoHandler).Methods("GET")
	s.router.HandleFunc("/api/echo", s.echoHandler).Methods("POST")
	s.router.HandleFunc("/ws/echo", s.echoWsHandler)
	s.router.HandleFunc("/v1/user/signup", s.SignUpHandler).Methods("POST")   // sign up a user
	s.router.HandleFunc("/v1/user/login", s.loginUserHandler).Methods("POST") // logs in a user into the system

	s.router.HandleFunc("/v1/user/logout", authMw.Handler(s.logoutHandler)).Methods("DELETE")                 // log user out of system
	s.router.HandleFunc("/v1/user/{id:[0-9]+}", authMw.Handler(s.updatedUserAccountHandler)).Methods("PUT")   // updates a user account
	s.router.HandleFunc("/v1/user/{id:[0-9]+}", authMw.Handler(s.getUserAccountHandler)).Methods("GET")       // get user account by id
	s.router.HandleFunc("/v1/user/{id:[0-9]+}", authMw.Handler(s.deleteUserAccountHandler)).Methods("DELETE") // deletes a user account by id

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// serve swagger files in a seamless manner
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml", Title: "BlackSpace Backend API"}
	sh := middleware.Redoc(ops, nil)
	s.router.Handle("/docs", sh)
	s.router.Handle("/swagger.yaml", http.FileServer(http.Dir("./pkg/api/docs")))
}

func (s *Server) registerMiddlewares() {
	prom := NewPrometheusMiddleware()
	s.router.Use(prom.Handler)
	httpLogger := NewLoggingMiddleware(s.logger)
	s.router.Use(httpLogger.Handler)
	s.router.Use(versionMiddleware)
	s.router.Use(mw.AuthnMiddleware(s.AuthnClient.Client))
}

func (s *Server) ListenAndServe(stopCh <-chan struct{}) {
	go s.startMetricsServer()

	s.registerHandlers()
	s.registerMiddlewares()

	srv := &http.Server{
		Addr:         ":" + s.config.Port,
		WriteTimeout: s.config.HttpServerTimeout,
		ReadTimeout:  s.config.HttpServerTimeout,
		IdleTimeout:  2 * s.config.HttpServerTimeout,
		Handler:      s.router,
	}

	// load configs in memory and start watching for changes in the config dir
	if stat, err := os.Stat(s.config.ConfigPath); err == nil && stat.IsDir() {
		var err error
		watcher, err = fscache.NewWatch(s.config.ConfigPath)
		if err != nil {
			s.logger.Error("config watch error", zap.Error(err), zap.String("path", s.config.ConfigPath))
		} else {
			watcher.Watch()
		}
	}

	// run server in background
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatal("HTTP server crashed", zap.Error(err))
		}
	}()

	// signal Kubernetes the server is ready to receive traffic
	atomic.StoreInt32(&healthy, 1)
	atomic.StoreInt32(&ready, 1)

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(context.Background(), s.config.HttpServerShutdownTimeout)
	defer cancel()

	// all calls to /healthz and /readyz will fail from now on
	atomic.StoreInt32(&healthy, 0)
	atomic.StoreInt32(&ready, 0)

	s.logger.Info("Shutting down HTTP server", zap.Duration("timeout", s.config.HttpServerShutdownTimeout))

	// wait for Kubernetes readiness probe to remove this instance from the load balancer
	// the readiness check interval must be lower than the timeout
	if viper.GetString("level") != "debug" {
		time.Sleep(3 * time.Second)
	}

	// attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Warn("HTTP server graceful shutdown failed", zap.Error(err))
	} else {
		s.logger.Info("HTTP server stopped")
	}
}

func (s *Server) startMetricsServer() {
	if s.config.PortMetrics > 0 {
		mux := http.DefaultServeMux
		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%v", s.config.PortMetrics),
			Handler: mux,
		}

		srv.ListenAndServe()
	}
}

func (s *Server) printRoutes() {
	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
}

func (s *Server) ExtractJwtFromHeader(r *http.Request) (*helper.TokenValidationResponse, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	// extract the id from the token
	claims := helper.JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(splitToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		if claims.StandardClaims.Issuer != s.config.JWTSigningAuthority {
			return nil, errors.New("invalid token")
		} else {
			return &helper.TokenValidationResponse{
				User:      claims.User,
				Id:        claims.Id,
				ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
			}, nil
		}
	}

	return nil, errors.New("invalid authorization token")
}

func (s *Server) GenerateAndSignJwtToken(userID uint32, user *models.UserORM) (*AccessTokens, error) {
	id := int(userID)
	idStr := strconv.Itoa(id)
	// sign jwt token
	claims := &helper.JwtCustomClaims{
		idStr,
		*user,
		jwt.StandardClaims{
			Issuer:    s.config.Issuer,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.config.JWTSecret))

	// create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix
	rt, err := refreshToken.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &AccessTokens{
		AccessToken:  t,
		RefreshToken: rt,
	}, err
}

type AccessTokens struct {
	AccessToken  string
	RefreshToken string
}

type ArrayResponse []string
type MapResponse map[string]string

// Package BlackSpace Authentication Handler Service API.
//
// This serves as the authentication handler microservice api definition for the BlackSpace Platform
//
// Terms Of Service:
//
// there are no TOS at this moment
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
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"strings"
	"sync/atomic"
	"time"

	core_metrics "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-metrics"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/keratin/authn-go/authn"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	core_logging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-logging/json"

	_ "github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/api/docs"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/authentication"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/fscache"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/metrics"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/middleware"
)

// @title authentication_handler_service API
// @version 2.0
// @description Go microservice template for Kubernetes.

// @contact.name Source Code
// @contact.url https://github.com/blackspaceInc/BlackspacePlatform/src/services/authentication_handler_service

// @license.name MIT License
// @license.url https://github.com/blackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/blob/master/LICENSE

// @host localhost:9898
// @BasePath /
// @schemes http https

var (
	healthy int32
	ready   int32
	watcher *fscache.Watcher
)

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
	CertPath                  string        `mapstructure:"cert-path"`
	Port                      string        `mapstructure:"port"`
	SecurePort                string        `mapstructure:"secure-port"`
	PortMetrics               int           `mapstructure:"port-metrics"`
	Hostname                  string        `mapstructure:"hostname"`
	H2C                       bool          `mapstructure:"h2c"`
	RandomDelay               bool          `mapstructure:"random-delay"`
	RandomDelayUnit           string        `mapstructure:"random-delay-unit"`
	RandomDelayMin            int           `mapstructure:"random-delay-min"`
	RandomDelayMax            int           `mapstructure:"random-delay-max"`
	RandomError               bool          `mapstructure:"random-error"`
	Unhealthy                 bool          `mapstructure:"unhealthy"`
	Unready                   bool          `mapstructure:"unready"`
	JWTSecret                 string        `mapstructure:"jwt-secret"`
	CacheServer               string        `mapstructure:"cache-server"`
}

type Server struct {
	router      *mux.Router
	config      *Config
	pool        *redis.Pool
	handler     http.Handler
	authnClient *AuthServiceClientWrapper
	logger      core_logging.ILog
	metrics     *metrics.CoreMetrics
	metricsEngine *core_metrics.CoreMetricsEngine
}

func NewServer(config *Config, client *AuthServiceClientWrapper, logging core_logging.ILog, serviceMetrics *metrics.CoreMetrics,
	metricsEngineConf *core_metrics.CoreMetricsEngine) (*Server,
	error) {
	srv := &Server{
		router:      mux.NewRouter(),
		config:      config,
		authnClient: client,
		logger:      logging,
		metrics:     serviceMetrics,
		metricsEngine: metricsEngineConf,
	}

	return srv, nil
}

func (s *Server) registerHandlers() {
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.Handle("/v1/metrics", promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, core_metrics.HandlerWithReset(s.metricsEngine.Registry,
		core_metrics.HandlerOpts{})))
	s.router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	s.router.HandleFunc("/", s.indexHandler).HeadersRegexp("User-Agent", "^Mozilla.*").Methods("GET")
	s.router.HandleFunc("/", s.infoHandler).Methods("GET")
	s.router.HandleFunc("/version", s.versionHandler).Methods("GET")
	s.router.HandleFunc("/healthz", s.healthzHandler).Methods("GET")
	s.router.HandleFunc("/readyz", s.readyzHandler).Methods("GET")
	s.router.HandleFunc("/readyz/enable", s.enableReadyHandler).Methods("POST")
	s.router.HandleFunc("/readyz/disable", s.disableReadyHandler).Methods("POST")
	s.router.HandleFunc("/api/info", s.infoHandler).Methods("GET")

	s.router.HandleFunc("/v1/account/create", s.createAccountHandler).Methods("POST")
	s.router.HandleFunc("/v1/account/update/{id:[0-9]+}", s.updateAccountHandler).Methods("POST", "PUT")
	s.router.HandleFunc("/v1/account/delete/{id:[0-9]+}", s.deleteAccountHandler).Methods("DELETE")
	s.router.HandleFunc("/v1/account/lock/{id:[0-9]+}", s.lockAccountHandler).Methods("POST")
	s.router.HandleFunc("/v1/account/unlock/{id:[0-9]+}", s.unlockAccountHandler).Methods("POST")
	s.router.HandleFunc("/v1/account/{id:[0-9]+}", s.getAccountHandler).Methods("GET")
	s.router.HandleFunc("/v1/account/login", s.loginAccountHandler).Methods("POST")
	s.router.HandleFunc("/v1/account/logout", s.logoutHandler).Methods("POST")

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	s.router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		doc, err := swag.ReadDoc()
		if err != nil {
			s.logger.Error(err, "swagger error", "path", "/swagger.json")
		}
		w.Write([]byte(doc))
	})
}

func (s *Server) registerMiddlewares() {
	prom := NewPrometheusMiddleware(s.metrics)
	s.router.Use(prom.Handler)
	httpLogger := NewLoggingMiddleware()
	s.router.Use(httpLogger.Handler)
	s.router.Use(versionMiddleware)
	s.router.Use(middleware.CorsMiddleware)
	authMw := middleware.NewAuthnMw(s.authnClient.Client, s.logger)
	s.router.Use(authMw.AuthenticationMiddleware)

	if s.config.RandomDelay {
		randomDelayer := NewRandomDelayMiddleware(s.config.RandomDelayMin, s.config.RandomDelayMax, s.config.RandomDelayUnit)
		s.router.Use(randomDelayer.Handler)
	}
	if s.config.RandomError {
		s.router.Use(randomErrorMiddleware)
	}
}

func (s *Server) ListenAndServe(stopCh <-chan struct{}) {
	go s.startMetricsServer()

	s.registerHandlers()
	s.registerMiddlewares()

	if s.config.H2C {
		s.handler = h2c.NewHandler(s.router, &http2.Server{})
	} else {
		s.handler = s.router
	}

	s.printRoutes()

	// load configs in memory and start watching for changes in the config dir
	if stat, err := os.Stat(s.config.ConfigPath); err == nil && stat.IsDir() {
		var err error
		watcher, err = fscache.NewWatch(s.config.ConfigPath)
		if err != nil {
			s.logger.Error(err, "config watch error", "path", s.config.ConfigPath)
		} else {
			watcher.Watch()
		}
	}

	// start redis connection pool
	ticker := time.NewTicker(30 * time.Second)
	s.startCachePool(ticker, stopCh)

	// create the http server
	srv := s.startServer()

	// create the secure server
	secureSrv := s.startSecureServer()

	// signal Kubernetes the server is ready to receive traffic
	if !s.config.Unhealthy {
		atomic.StoreInt32(&healthy, 1)
	}
	if !s.config.Unready {
		atomic.StoreInt32(&ready, 1)
	}

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(context.Background(), s.config.HttpServerShutdownTimeout)
	defer cancel()

	// all calls to /healthz and /readyz will fail from now on
	atomic.StoreInt32(&healthy, 0)
	atomic.StoreInt32(&ready, 0)

	// close cache pool
	if s.pool != nil {
		_ = s.pool.Close()
	}

	s.logger.Info("Shutting down HTTP/HTTPS server", "timeout", s.config.HttpServerShutdownTimeout)

	// wait for Kubernetes readiness probe to remove this instance from the load balancer
	// the readiness check interval must be lower than the timeout
	if viper.GetString("level") != "debug" {
		time.Sleep(3 * time.Second)
	}

	// determine if the http server was started
	if srv != nil {
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.ErrorM(err, "HTTP server graceful shutdown failed")
		}
	}

	// determine if the secure server was started
	if secureSrv != nil {
		if err := secureSrv.Shutdown(ctx); err != nil {
			s.logger.ErrorM(err, "HTTPS server graceful shutdown failed")
		}
	}
}

func (s *Server) startServer() *http.Server {

	// determine if the port is specified
	if s.config.Port == "0" {

		// move on immediately
		return nil
	}

	srv := &http.Server{
		Addr:         ":" + s.config.Port,
		WriteTimeout: s.config.HttpServerTimeout,
		ReadTimeout:  s.config.HttpServerTimeout,
		IdleTimeout:  2 * s.config.HttpServerTimeout,
		Handler:      s.handler,
	}

	// start the server in the background
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.FatalM(err, "HTTP server crashed")
		}
	}()

	// return the server and routine
	return srv
}

func (s *Server) startSecureServer() *http.Server {

	// determine if the port is specified
	if s.config.SecurePort == "0" {

		// move on immediately
		return nil
	}

	srv := &http.Server{
		Addr:         ":" + s.config.SecurePort,
		WriteTimeout: s.config.HttpServerTimeout,
		ReadTimeout:  s.config.HttpServerTimeout,
		IdleTimeout:  2 * s.config.HttpServerTimeout,
		Handler:      s.handler,
	}

	cert := path.Join(s.config.CertPath, "tls.crt")
	key := path.Join(s.config.CertPath, "tls.key")

	// start the server in the background
	go func() {
		if err := srv.ListenAndServeTLS(cert, key); err != http.ErrServerClosed {
			s.logger.FatalM(err, "HTTPS server crashed")
		}
	}()

	// return the server
	return srv
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

type ArrayResponse []string
type MapResponse map[string]string

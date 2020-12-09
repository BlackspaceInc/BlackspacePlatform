package core_tracing

import (
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
)

// Init creates a new instance of Jaeger tracer.
func Init(serviceName string, metricsFactory metrics.Factory) (opentracing.Tracer, io.Closer){
	l, _ := zap.NewProduction()
	jaegerLogger := jaegerLoggerAdapter{l}

	cfg, err := config.FromEnv()
	if err != nil {
		l.Fatal(err.Error(), zap.String("error", "cannot parse Jaeger env vars"))
	}
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	// TODO(ys) a quick hack to ensure random generators get different seeds, which are based on current time.
	time.Sleep(100 * time.Millisecond)
	metricsFactory = metricsFactory.Namespace(metrics.NSOptions{Name: serviceName, Tags: nil})
	tracer, closer , err := cfg.NewTracer(
		config.Logger(jaegerLogger),
		config.Metrics(metricsFactory),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		l.Fatal(err.Error(), zap.String("error", "cannot initialize Jaeger Tracer"))
	}
	return tracer, closer
}

type jaegerLoggerAdapter struct {
	logger *zap.Logger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(msg,zap.Any(msg, fmt.Sprintf(msg, args...)))
}

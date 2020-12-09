package tracing

import (
	"io"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerprom "github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

type ServiceTracer struct {
	Tracer opentracing.Tracer
	Closer io.Closer
	Reporter jaeger.Reporter
}

func NewTracer(serviceName string, hostname string, production bool) *ServiceTracer {
	l, _ := zap.NewProduction()
	jaegerLogger := jaegerLoggerAdapter{l}

	factory := jaegerprom.New()
	metrics := jaeger.NewMetrics(factory, map[string]string{"lib": "jaeger", "serviceName": serviceName})

	transport, err := jaeger.NewUDPTransport(hostname, 0)
	if err != nil {
		l.Fatal(err.Error(), zap.String("error", "cannot create transport object"))
	}

	reporter := jaeger.NewCompositeReporter(
		jaeger.NewLoggingReporter(jaegerLogger),
		jaeger.NewRemoteReporter(transport,
			jaeger.ReporterOptions.Metrics(metrics),
			jaeger.ReporterOptions.Logger(jaegerLogger),
			jaeger.ReporterOptions.BufferFlushInterval(1 * time.Second),
		),
	)

	var recordAllSpans = false
	if production {
		recordAllSpans = true
	}

	sampler := jaeger.NewConstSampler(recordAllSpans)

	tracer, closer := jaeger.NewTracer(serviceName,
		sampler,
		reporter,
		jaeger.TracerOptions.Metrics(metrics),
	)

	return &ServiceTracer{
		Tracer: tracer,
		Closer: closer,
		Reporter: reporter,
	}
}

func (s *ServiceTracer) OpenTracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wireCtx, _ := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))

		serverSpan := opentracing.StartSpan(r.URL.Path,
			ext.RPCServerOption(wireCtx))
		defer serverSpan.Finish()

		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
		next.ServeHTTP(w, r)
	})
}


type jaegerLoggerAdapter struct {
	logger *zap.Logger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(msg)
}


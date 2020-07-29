package metrics

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	onceEnable         sync.Once
	prometheusRegistry = prometheus.NewRegistry()
	defaultRegistry    Registry
)

// Options control config
type Options struct {
	FlushInterval          time.Duration
	EnableGoRuntimeMetrics bool
	DbConn                 *gorm.DB
	ServiceName            string
}

// MetricsExporter is a prom exporter for go
type MetricsExporter struct {
	FlushInterval time.Duration
	lc            sync.RWMutex
	lg            sync.RWMutex
	ls            sync.RWMutex
	counters      map[string]*prometheus.CounterVec
	gauges        map[string]*prometheus.GaugeVec
	summaries     map[string]*prometheus.SummaryVec
	histograms    map[string]*prometheus.HistogramVec
	gaugeFuncs    map[string]*prometheus.GaugeFunc
	ServiceName   string
}

//NewMetricsExporter create a prometheus exporter
func NewMetricsExporter(options Options) Registry {
	if options.EnableGoRuntimeMetrics {
		onceEnable.Do(func() {
			EnableRunTimeMetrics()
			fmt.Print("go runtime metrics is exported")
		})

	}
	metricsExporter := &MetricsExporter{
		FlushInterval: options.FlushInterval,
		lc:            sync.RWMutex{},
		lg:            sync.RWMutex{},
		ls:            sync.RWMutex{},
		summaries:     make(map[string]*prometheus.SummaryVec),
		counters:      make(map[string]*prometheus.CounterVec),
		gauges:        make(map[string]*prometheus.GaugeVec),
		histograms:    make(map[string]*prometheus.HistogramVec),
		gaugeFuncs:    make(map[string]*prometheus.GaugeFunc),
		ServiceName:   options.ServiceName,
	}

	RegisterInfraMetrics(options, metricsExporter)
	return metricsExporter
}

func RegisterInfraMetrics(options Options, metricsExporter *MetricsExporter) {
	RegisterCommonDatabaseCounters(options.ServiceName, options.DbConn, metricsExporter)
	RegisterCommonRuntimeCounters(options.ServiceName, metricsExporter)
	RegisterCommonRequestCounters(options.ServiceName, metricsExporter)
}

// EnableRunTimeMetrics enable runtime metrics
func EnableRunTimeMetrics() {
	GetSystemPrometheusRegistry().MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	GetSystemPrometheusRegistry().MustRegister(prometheus.NewGoCollector())
}

// GetSystemPrometheusRegistry return prometheus registry which go chassis use
func GetSystemPrometheusRegistry() *prometheus.Registry {
	return prometheusRegistry
}

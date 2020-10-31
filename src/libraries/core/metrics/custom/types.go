package custom

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// NewRegistry create a registry
type NewRegistry func(opts Options) Registry

// Registry holds all of metrics collectors
// name is a unique ID for different type of metrics
type Registry interface {
	// CreateGauge creates a gauge metric type
	CreateGauge(opts GaugeOpts) error
	// CreateCounter creates a counter metric type
	CreateCounter(opts CounterOpts) error
	// CreateSummary creates a summary metric type
	CreateSummary(opts SummaryOpts) error
	// CreateHistogram creates a histogram metric type
	CreateHistogram(opts HistogramOpts) error
	// CreateGaugeFunc create a gauge function metric type
	CreateGaugeFunc(opts GaugeOpts, f func() float64) error

	// GaugeSet sets a gauge counter type to a value specified by val
	GaugeSet(name string, val float64, labels map[string]string) error
	// CounterAdd sets a gauge counter to a value specified by val
	CounterAdd(name string, val float64, labels map[string]string) error
	// SummaryObserve specifies a counter of summary type
	SummaryObserve(name string, val float64, Labels map[string]string) error
	// HistogramObserve set/increments a histogram counter type
	HistogramObserve(name string, val float64, labels map[string]string) error
	// ObtainGaugeFunc attempts to obtain a gauge function by name
	ObtainGaugeFunc(name string) (*prometheus.GaugeFunc, error)
	// Monitor observes the request size and latency associated with a specific route and records it in a summary Blackspace metric
	Monitor(next http.Handler) http.Handler
}

// CounterOpts is options to create a counter options
type CounterOpts struct {
	Name      string
	Namespace string
	Subsystem string
	Help      string
	Labels    []string
}

// GaugeOpts is options to create a gauge collector
type GaugeOpts struct {
	Name      string
	Namespace string
	Subsystem string
	Help      string
	Labels    []string
}

// SummaryOpts is options to create summary collector
type SummaryOpts struct {
	Name       string
	Namespace  string
	Subsystem  string
	Help       string
	Labels     []string
	Objectives map[float64]float64
}

// HistogramOpts is options to create histogram collector
type HistogramOpts struct {
	Name      string
	Namespace string
	Subsystem string
	Help      string
	Labels    []string
	Buckets   []float64
}

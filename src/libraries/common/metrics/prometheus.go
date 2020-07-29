package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// CreateGauge creates a metric of type gauge
func (c *MetricsExporter) CreateGauge(opts GaugeOpts) error {
	name := c.ServiceName + opts.Name
	c.lg.RLock()
	_, ok := c.gauges[name]
	c.lg.RUnlock()
	if ok {
		return fmt.Errorf("metric [%s] is duplicated", opts.Name)
	}
	c.lg.Lock()
	defer c.lg.Unlock()
	gVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: opts.Help,
	}, opts.Labels)
	c.gauges[name] = gVec
	GetSystemPrometheusRegistry().MustRegister(gVec)
	return nil
}

// CreateCounter creates a metric of type counter
func (c *MetricsExporter) CreateCounter(opts CounterOpts) error{
	name := c.ServiceName + opts.Name
	c.lc.RLock()
	_, ok := c.counters[name]
	c.lc.RUnlock()
	if ok {
		return fmt.Errorf("metric [%s] is duplicated", opts.Name)
	}
	c.lc.Lock()
	defer c.lc.Unlock()
	v := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: opts.Help,
	}, opts.Labels)
	c.counters[name] = v
	GetSystemPrometheusRegistry().MustRegister(v)
	return nil
}

// CreateSummary creates a metric of type summary
func (c *MetricsExporter) CreateSummary(opts SummaryOpts) error {
	name := c.ServiceName + opts.Name
	c.ls.RLock()
	_, ok := c.summaries[name]
	c.ls.RUnlock()
	if ok {
		return fmt.Errorf("metric [%s] is duplicated", opts.Name)
	}
	c.ls.Lock()
	defer c.ls.Unlock()
	v := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       name,
		Help:       opts.Help,
		Objectives: opts.Objectives,
	}, opts.Labels)
	c.summaries[name] = v
	GetSystemPrometheusRegistry().MustRegister(v)
	return nil
}

// CreateHistogram creates a metric of type histogram
func (c *MetricsExporter) CreateHistogram(opts HistogramOpts) error {
	name := c.ServiceName + opts.Name
	c.ls.RLock()
	_, ok := c.histograms[name]
	c.ls.RUnlock()
	if ok {
		return fmt.Errorf("metric [%s] is duplicated", opts.Name)
	}
	c.ls.Lock()
	defer c.ls.Unlock()
	v := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    opts.Help,
		Buckets: opts.Buckets,
	}, opts.Labels)
	c.histograms[name] = v
	GetSystemPrometheusRegistry().MustRegister(v)
	return nil
}

func (c *MetricsExporter) CreateGaugeFunc(opts GaugeOpts, f func() float64) error {
	name := c.ServiceName + opts.Name
	c.lg.RLock()
	_, ok := c.gaugeFuncs[name]
	c.lg.RUnlock()
	if ok {
		return fmt.Errorf("metric [%s] is duplicated", opts.Name)
	}
	c.lg.Lock()
	defer c.lg.Unlock()
	gFunc := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: name,
		Help: opts.Help,
	}, f)
	c.gaugeFuncs[name] = &gFunc
	GetSystemPrometheusRegistry().MustRegister(gFunc)
	return nil
}

// GaugeSet adds a new counter (gauge) and sets initial value
func (c *MetricsExporter) GaugeSet(name string, val float64, labels map[string]string) error {
	name = c.ServiceName + name
	c.lg.RLock()
	gVec, ok := c.gauges[name]
	c.lg.RUnlock()
	if !ok {
		return fmt.Errorf("metrics do not exists, create it first")
	}
	gVec.With(labels).Set(val)
	return nil
}

// CounterAdd adds a new counter and sets its initial start value
func (c *MetricsExporter) CounterAdd(name string, val float64, labels map[string]string) error {
	name = c.ServiceName + name
	c.lc.RLock()
	v, ok := c.counters[name]
	c.lc.RUnlock()
	if !ok {
		return fmt.Errorf("metrics do not exists, create it first")
	}
	v.With(labels).Add(val)
	return nil
}

// SummaryObserve creates a summary counter and initializes it to a start value
func (c *MetricsExporter) SummaryObserve(name string, val float64, labels map[string]string) error {
	name = c.ServiceName + name
	// define critical section in which the counter will be attempted to be added to the metrics exports
	// counter type mapping
	c.ls.RLock()
	v, ok := c.summaries[name]
	c.ls.RUnlock()

	// if counter type is not found withing the list create it
	if !ok {
		return fmt.Errorf("metrics do not exists, create it first")
	}
	v.With(labels).Observe(val)
	return nil
}

// HistogramObserve creates a histogram counter and initializes it to some start value
func (c *MetricsExporter) HistogramObserve(name string, val float64, labels map[string]string) error {
	name = c.ServiceName + name
	c.ls.RLock()
	v, ok := c.histograms[name]
	c.ls.RUnlock()
	if !ok {
		return fmt.Errorf("metrics do not exists, create it first")
	}
	v.With(labels).Observe(val)
	return nil
}

func (c *MetricsExporter) ObtainGaugeFunc(name string) (*prometheus.GaugeFunc, error){
	name = c.ServiceName + name
	c.ls.RLock()
	v, ok := c.gaugeFuncs[name]
	c.ls.RUnlock()
	if !ok {
		return nil, fmt.Errorf("metrics do not exists, create it first")
	}
	return v, nil
}

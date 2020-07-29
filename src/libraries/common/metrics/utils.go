package metrics

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func registerCounter(serviceName string, db *gorm.DB, metricsExporter *MetricsExporter, value *CounterMetadata) {
	// obtain the counter type
	switch value.CounterType {
	case Enum.GaugeFunc:
		registerGaugeFuncMetrics(serviceName, db, metricsExporter, value)
	case Enum.SummaryVec:
		registerSummaryMetrics(serviceName, value, metricsExporter)
	case Enum.CounterVec:
		registerCounterMetrics(serviceName, value, metricsExporter)
	case Enum.GaugeVec:
		registerGaugeMetrics(serviceName, value, metricsExporter)
	case Enum.HistogramVec:
		registerHistogramMetrics(serviceName, value, metricsExporter)
	default:
		registerCounterMetrics(serviceName, value, metricsExporter)
	}
}

func registerHistogramMetrics(serviceName string, value *CounterMetadata, metricsExporter *MetricsExporter) {
	opts := HistogramOpts{
		Name:     value.CounterName,
		Namespace: value.CounterNameSpace(serviceName),
		Subsystem: value.CounterSubSystem,
		Help:      value.CounterHelpString,
		Labels:    value.CounterLabels,
		Buckets:   value.CounterBuckets,
	}
	if err := metricsExporter.CreateHistogram(opts); err != nil {
		fmt.Sprintf("metric [%s] wast not able to be created due to the following reason : [%s]", opts.Name, err.Error())
	}
}

func registerGaugeMetrics(serviceName string, value *CounterMetadata, metricsExporter *MetricsExporter) {
	opts := GaugeOpts{
		Name:      value.CounterName,
		Namespace: value.CounterNameSpace(serviceName),
		Subsystem: value.CounterSubSystem,
		Help:      value.CounterHelpString,
		Labels:    value.CounterLabels,
	}
	if err := metricsExporter.CreateGauge(opts); err != nil {
		fmt.Sprintf("metric [%s] wast not able to be created due to the following reason : %s", opts.Name, err.Error())
	}
}

func registerCounterMetrics(serviceName string, value *CounterMetadata, metricsExporter *MetricsExporter) {
	opts := CounterOpts{
		Name:      value.CounterName,
		Namespace: value.CounterNameSpace(serviceName),
		Subsystem: value.CounterSubSystem,
		Help:      value.CounterHelpString,
		Labels:    value.CounterLabels,
	}
	if err := metricsExporter.CreateCounter(opts); err != nil {
		fmt.Sprintf("metric [%s] wast not able to be created due to the following reason : %s", opts.Name, err.Error())
	}
}

func registerSummaryMetrics(serviceName string, value *CounterMetadata, metricsExporter *MetricsExporter) {
	opts := SummaryOpts{
		Name:       value.CounterName,
		Namespace:  value.CounterNameSpace(serviceName),
		Subsystem:  value.CounterSubSystem,
		Help:       value.CounterHelpString,
		Labels:     value.CounterLabels,
		Objectives: nil,
	}
	if err := metricsExporter.CreateSummary(opts); err != nil {
		fmt.Sprintf("metric [%s] wast not able to be created due to the following reason : %s", opts.Name, err.Error())
	}
}

func registerGaugeFuncMetrics(serviceName string, db *gorm.DB, metricsExporter *MetricsExporter, value *CounterMetadata) {
	opts := GaugeOpts{
		Name:      value.CounterName,
		Namespace: value.CounterNameSpace(serviceName),
		Subsystem: value.CounterSubSystem,
		Help:      value.CounterHelpString,
		Labels:    value.CounterLabels,
	}
	function := value.CounterFunction(db)
	if err := metricsExporter.CreateGaugeFunc(opts, function); err != nil {
		fmt.Sprintf("metric [%s] wast not able to be created due to the following reason : %s", opts.Name, err.Error())
	}
}

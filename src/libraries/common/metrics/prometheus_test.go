package metrics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BlackspaceInc/common/metrics"
)

func TestCounterMetrics(t *testing.T){
	t.Run("TestName:AddCounterVec", TestAddDatabaseOperationLatency)
	t.Run("TestName:SetGaugeCounter", TestGaugeSet)
	t.Run("TestName:ObserveSummary", TestSummaryObserve)
	t.Run("TestName:CreateHistogram", TestCreateHistogram)
}

func TestGaugeSet(t *testing.T) {
	// generate a random string
	randMetricName := "_" + metrics.GenerateRandomString(10)

	err := metricExporter.GaugeSet(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.Error(t, err)

	err = metricExporter.CreateGauge(metrics.GaugeOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.NoError(t, err)
	err = metricExporter.CreateGauge(metrics.GaugeOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.Error(t, err)

	err = metricExporter.GaugeSet(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.NoError(t, err)
}

func TestSummaryObserve(t *testing.T) {
	// generate a random string
	randMetricName := "_" + metrics.GenerateRandomString(10)

	err := metricExporter.SummaryObserve(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.Error(t, err)

	err = metricExporter.CreateSummary(metrics.SummaryOpts{
		Name:  randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.NoError(t, err)
	err = metricExporter.CreateSummary(metrics.SummaryOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.Error(t, err)
	err = metricExporter.SummaryObserve(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.NoError(t, err)
}

func TestCreateHistogram(t *testing.T) {
	// generate a random string
	randMetricName := "_" + metrics.GenerateRandomString(10)

	err := metricExporter.HistogramObserve(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.Error(t, err)

	err = metricExporter.CreateHistogram(metrics.HistogramOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.NoError(t, err)
	err = metricExporter.CreateHistogram(metrics.HistogramOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.Error(t, err)

	err = metricExporter.HistogramObserve(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.NoError(t, err)
}

func TestAddDatabaseOperationLatency(t *testing.T){
	// generate a random string
	randMetricName := "_" + metrics.GenerateRandomString(10)
	// error out since counter does not exist
	err := metricExporter.CounterAdd(randMetricName, 1, map[string]string{
		"service": "s",
	})
	assert.Error(t, err)

	err = metricExporter.CreateCounter(metrics.CounterOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.NoError(t, err)

	// should error out since we will have duplicate counters
	err = metricExporter.CreateCounter(metrics.CounterOpts{
		Name:   randMetricName,
		Help:   "1",
		Labels: []string{"service"},
	})
	assert.Error(t, err)

	// should not throw error since we are associating a counter to a label that does exist
	err = metricExporter.CounterAdd(randMetricName, 1, map[string]string{
		"service": "user_management",
	})
	assert.NoError(t, err)
}

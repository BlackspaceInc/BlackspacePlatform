package metrics_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"

	"github.com/BlackspaceInc/common/metrics"
)

var (
	serviceName = "test_service"
	db = metrics.SetupTests()
	option = metrics.Options{
		FlushInterval:          0,
		EnableGoRuntimeMetrics: false,
		DbConn:                 db,
		ServiceName:           serviceName,
	}
	metricExporter = metrics.NewMetricsExporter(option)
)

func TestDatabaseCountersPresent(t *testing.T) {
	t.Run("TestName:OpenConnectionCounterExists", TestOpenConnectionsCounterExists)
	t.Run("TestName:IdleConnectionsCounterExists", TestIdleConnectionsCounterExists)
	t.Run("TestName:ConnectionsInUseCounterExists", TestConnectionsInUseCounterExists)
	t.Run("TestName:DatabaseConnectionWaitDurationCounterExists", TestDatabaseConnectionWaitDurationCounterExists)
	t.Run("TestName:DatabaseOperationLatency", TestDatabaseOperationLatency)
}

func TestOpenConnectionsCounterExists(t *testing.T) {
	var counterName = metrics.OpenConnections
	validateGaugeFuncExists(t, counterName)
}

func TestIdleConnectionsCounterExists(t *testing.T) {
	var counterName = metrics.IdleConnections
	validateGaugeFuncExists(t, counterName)
}

func TestConnectionsInUseCounterExists(t *testing.T) {
	var counterName = metrics.ConnectionsInUse
	validateGaugeFuncExists(t, counterName)
}

func TestDatabaseConnectionWaitDurationCounterExists(t *testing.T) {
	var counterName = metrics.DatabaseConnectionWaitDuration
	validateGaugeFuncExists(t, counterName)
}

func TestDatabaseOperationLatency(t *testing.T){
	var counterName = metrics.DatabaseOperationLatency
	label := map[string]string{
		"operation" : "test_db",
	}
	err := metricExporter.SummaryObserve(counterName, 20, label)
	assert.Empty(t,err)
}

func validateGaugeFuncExists(t *testing.T, counterName string) {
	gaugeFuncMetric, err := metricExporter.ObtainGaugeFunc(counterName)
	assert.Empty(t, err)
	assert.NotEmpty(t, gaugeFuncMetric)
}

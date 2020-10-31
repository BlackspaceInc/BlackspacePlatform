package custom

import (
	"github.com/jinzhu/gorm"

	"github.com/BlackspaceInc/common/metrics"
)

// RegisterCommonDatabaseCounters generate common database specific counters of interest for a given service
func RegisterCommonDatabaseCounters(serviceName string, db *gorm.DB, metricsExporter *MetricsExporter){
	for _, value := range databaseCounters {
		metrics.registerCounter(serviceName, db, metricsExporter, value)
	}
}

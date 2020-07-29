package metrics

import (
	"github.com/jinzhu/gorm"
)

// RegisterCommonDatabaseCounters generate common database specific counters of interest for a given service
func RegisterCommonDatabaseCounters(serviceName string, db *gorm.DB, metricsExporter *MetricsExporter){
	for _, value := range databaseCounters {
		registerCounter(serviceName, db, metricsExporter, value)
	}
}

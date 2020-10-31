package custom

import (
	"github.com/jinzhu/gorm"
)

const (
	OpenConnections                = "_open_connections"
	IdleConnections                = "_idle_connections_open"
	ConnectionsInUse               = "_connections_in_use"
	DatabaseConnectionWaitDuration = "_database_connection_wait_duration"
	DatabaseOperationLatency       = "_database_operation_latency"
)

var databaseCounters = map[string] *CounterMetadata{
	OpenConnections: &CounterMetadata{
		CounterType:       Enum.GaugeFunc,
		CounterName:       OpenConnections,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  SUBSYSTEM,
		CounterHelpString: "number of open database connections",
		CounterLabels:     nil,
		CounterFunction:   func (db *gorm.DB) func()float64{
			return func() float64 {
				return float64(db.DB().Stats().MaxOpenConnections)
			}
		},
		CounterBuckets: nil,
	},
	IdleConnections: &CounterMetadata{
		CounterType:       Enum.GaugeFunc,
		CounterName:       IdleConnections,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  SUBSYSTEM,
		CounterHelpString: "number of idle database connections opened",
		CounterLabels:     nil,
		CounterFunction:   func (db *gorm.DB) func()float64{
			return func() float64 {
				return float64(db.DB().Stats().Idle)
			}
		},
		CounterBuckets: nil,
	},
	ConnectionsInUse: &CounterMetadata{
		CounterType:       Enum.GaugeFunc,
		CounterName:       ConnectionsInUse,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  SUBSYSTEM,
		CounterHelpString: "number of database connections in use",
		CounterLabels:     nil,
		CounterFunction:   func (db *gorm.DB) func()float64{
			return func() float64 {
				return float64(db.DB().Stats().InUse)
			}
		},
		CounterBuckets: nil,
	},
	DatabaseConnectionWaitDuration: &CounterMetadata{
		CounterType:       Enum.GaugeFunc,
		CounterName:       DatabaseConnectionWaitDuration,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  SUBSYSTEM,
		CounterHelpString: "time blocked waiting for a new connection to the database",
		CounterLabels:     nil,
		CounterFunction:   func (db *gorm.DB) func()float64{
			return func() float64 {
				return float64(db.DB().Stats().InUse)
			}
		},
		CounterBuckets: nil,
	},
	DatabaseOperationLatency: &CounterMetadata{
		CounterType:       Enum.SummaryVec,
		CounterName:       DatabaseOperationLatency,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  SUBSYSTEM,
		CounterHelpString: "time blocked waiting for a new connection to the database",
		CounterLabels:     []string{"operation"},
		CounterFunction:   nil,
	},
}

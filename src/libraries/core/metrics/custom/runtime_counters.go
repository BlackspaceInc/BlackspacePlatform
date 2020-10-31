package custom

import (
	"runtime"

	"github.com/jinzhu/gorm"
)

var RUNTIME_SUBSYSTEM = "Runtime"

const (
	NumberOfGoRoutines = "_number_of_go_routines"
)

var runtimeCounters = map[string] *CounterMetadata{
	NumberOfGoRoutines: &CounterMetadata{
		CounterType:       Enum.GaugeFunc,
		CounterName:       NumberOfGoRoutines,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  RUNTIME_SUBSYSTEM,
		CounterHelpString: "number of goroutines that currently exist",
		CounterLabels:     nil,
		CounterFunction:   func (db *gorm.DB) func()float64{
			return func() float64 { return float64(runtime.NumGoroutine())}
		},
		CounterBuckets: nil,
	},
}

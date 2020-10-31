package custom

import (
	"github.com/jinzhu/gorm"
)

const SUBSYSTEM string = "database"

type CounterAlias string

type CounterType struct {
	GaugeFunc    CounterAlias
	CounterVec   CounterAlias
	GaugeVec     CounterAlias
	SummaryVec   CounterAlias
	HistogramVec CounterAlias
}

var Enum = &CounterType{
	GaugeFunc:    "GAUGE_FUNC",
	CounterVec:   "COUNTER_VEC",
	GaugeVec:     "GAUGE_VEC",
	SummaryVec:   "SUMMARY_VEC",
	HistogramVec: "HISTOGRAM_VEC",
}

type CounterMetadata struct {
	CounterType       CounterAlias
	CounterName       string
	CounterNameSpace  func(string) string
	CounterSubSystem  string
	CounterHelpString string
	CounterLabels     []string
	CounterFunction   func(db *gorm.DB) func()float64
	CounterBuckets    []float64
}

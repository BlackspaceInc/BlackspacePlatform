package custom

var REQUEST_SUBSYSTEM = "Request"

const (
	RequestInfo = "_request_info"
)

var requestsCounters = map[string] *CounterMetadata{
	RequestInfo: &CounterMetadata{
		CounterType:       Enum.SummaryVec,
		CounterName:       RequestInfo,
		CounterNameSpace:  func(serviceName string) string {return serviceName},
		CounterSubSystem:  REQUEST_SUBSYSTEM,
		CounterHelpString: "number of goroutines that currently exist",
		CounterLabels:     []string{"size", "duration"},
		CounterFunction:   nil,
		CounterBuckets:    nil,
	},
}

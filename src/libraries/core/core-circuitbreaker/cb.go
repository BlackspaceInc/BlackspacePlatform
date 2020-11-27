package core_circuitbreaker


import (
	"github.com/sony/gobreaker"
)


var cb *gobreaker.CircuitBreaker

func init(){
	var st gobreaker.Settings
	st.MaxRequests = 10
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 5 && failureRatio >= 0.6
	}
	st.Interval = 1
	cb = gobreaker.NewCircuitBreaker(st)
}

package custom

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BlackspaceInc/common/metrics"
)

func (c *MetricsExporter) Monitor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			routeName := r.RequestURI
			counterName := c.ServiceName + "_" + r.RequestURI
			_ = c.CreateSummary(metrics.SummaryOpts{
				Name:      counterName,
				Namespace: c.ServiceName,
				Subsystem: "Requests",
				Help:      routeName + "" + "duration and request size",
				Labels: []string{
					"size",
					"duration",
				},
				Objectives: nil,
			})

			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			// Store duration of request
			_ = c.SummaryObserve(counterName, duration.Seconds(), map[string]string{"duration": "request_duration"})

			// Store size of response, if possible.
			size, err := strconv.Atoi(w.Header().Get("Content-Length"))
			if err == nil {
				_ = c.SummaryObserve(counterName, float64(size), map[string]string{"size": "request_size"})
			}
		})
}

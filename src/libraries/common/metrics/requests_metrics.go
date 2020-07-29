package metrics

// RegisterCommonRuntimeCounters generate common runtime specific counters of interest for a given service
func RegisterCommonRequestCounters(serviceName string, metricsExporter *MetricsExporter){
	for _, value := range requestsCounters {
		registerCounter(serviceName, nil, metricsExporter, value)
	}
}

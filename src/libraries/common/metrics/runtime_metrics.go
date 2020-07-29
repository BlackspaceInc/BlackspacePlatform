package metrics

// RegisterCommonRuntimeCounters generate common runtime specific counters of interest for a given service
func RegisterCommonRuntimeCounters(serviceName string, metricsExporter *MetricsExporter){
	for _, value := range runtimeCounters {
		registerCounter(serviceName, nil, metricsExporter, value)
	}
}

## Blackspace Core Metrics Library
---
This document outlines how to effectively make use of this library.

```go
// define a core engine registry object to which the version info would be tied to
CoreEngine := NewCoreMetricsEngineInstance("blackspace_platform",nil)

// Define a
LoginRequestCounter = NewGaugeVec(&GaugeOpts{
        Namespace: "blackspace_platform"
		Subsystem:  "authentication_service",
		Name:      "login_request_counter",
		Help:      "Number of log in requests",
	}, []string{"Request"})

CoreEngine.RegisterMetric(LoginRequestCounter)

LoginRequestCounter.WithLabelValues("Request").Observe(1.0)
```

## Blackspace Core Metrics Library
---
This document outlines how to effectively make use of this library.

```go
// define a version info object which outlines
// the library version useful especially in cases when metrics are deprecated
// across versions
var currentVersion = apimachineryversion.Info{
		Major:      "1",
		Minor:      "17",
		GitVersion: "v1.17.0-alpha-1.12345",
	}

// define a registry object to which the version info would be tied to
registry := newPlatformRegistry(currentVersion)

// Define a
LoginRequestCounter = NewGaugeVec(&GaugeOpts{
		Subsystem: "Authentication Service",
		Name:      "login_request_counter",
		Help:      "Number of log in requests",
	}, []string{"Request"})

registry.MustRegister(LoginRequestCounter)

LoginRequestCounter.WithLabelValues("Request").Observe(1.0)
```

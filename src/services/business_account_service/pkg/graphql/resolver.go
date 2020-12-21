package graphql

import (
	core_logging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-logging/json"
	core_metrics "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-metrics"
	core_tracing "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-tracing"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db      *database.Db
	Logger  core_logging.ILog
	Tracer  *core_tracing.TracingEngine
	Metrics *core_metrics.CoreMetricsEngine
}

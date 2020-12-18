package graphql

import (
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db *database.Db
}

package database

import (
	"context"
	"os"
	"time"

	core_logging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-logging/json"
	core_metrics "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-metrics"
	core_tracing "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-tracing"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/utils"
)

// IDatabase provides an interface which any database tied to this service should implement
type IDatabase interface {
	// To be implemented
}

// Db witholds connection to a postgres database as well as a logging handler
type Db struct {
	Engine                                   *gorm.DB
	Logger                                   core_logging.ILog
	TracingEngine                            *core_tracing.TracingEngine
	MetricsEngine                            *core_metrics.CoreMetricsEngine
	AuthenticationHandlerServiceBaseEndpoint string
}

// Tx is a type serving as a function decorator for common database transactions
type Tx func(ctx context.Context, tx *gorm.DB) error

// CmplxTx is a type serving as a function decorator for complex database transactions
type CmplxTx func(ctx context.Context, tx *gorm.DB) (interface{}, error)

// type of database
var postgres = "postgres"

// database operation types
const (
	DB_CONNECTION_ATTEMPT = "DB_CONNECTION_ATTEMPT"
)

var maxConnectionRetryAttempts = 5

// New creates a database connection and returns the connection object
func New(ctx context.Context, connectionString string, tracingEngine *core_tracing.TracingEngine, metricsEngine *core_metrics.CoreMetricsEngine,
	logger core_logging.ILog, svcEndpoint string) (*Db,
	error) {

	if connectionString == utils.EMPTY || tracingEngine == nil || metricsEngine == nil || logger == nil {
		// crash the process
		os.Exit(1)
	}

	// generate a span for the database connection
	ctx, span := utils.StartRootOperationSpan(ctx, DB_CONNECTION_ATTEMPT, tracingEngine, logger)
	defer span.Finish()

	logger.Info("Attempting database connection operation")
	conn, err := ConnectToDb(connectionString, logger)
	if err != nil {
		logger.FatalM(err, errors.ErrFailedToConnectToDatabase.Error())
	}
	logger.Info("Successfully connected to the database")

	// configure db
	logger.Info("Attempting database connection configuration")
	conn = configureDbConnection(conn)
	logger.Info("Successfully configured database connection object")

	logger.Info("Attempting database schema migration")
	err = MigrateSchemas(conn, logger)
	if err != nil {
		logger.FatalM(err, errors.ErrFailedToPerformDatabaseMigrations.Error())
	}
	logger.Info("Successfully migrated database")

	var endpoint = svcEndpoint
	if endpoint == "" {
		endpoint = "http://authentication-handler-service:9898/v1/account"
	}
	return &Db{
		Engine:                                   conn,
		Logger:                                   logger,
		TracingEngine:                            tracingEngine,
		MetricsEngine:                            metricsEngine,
		AuthenticationHandlerServiceBaseEndpoint: svcEndpoint,
	}, nil
}

func ConnectToDb(connectionString string, logger core_logging.ILog) (*gorm.DB, error) {
	retries := 0
	for retries < maxConnectionRetryAttempts {
		// perform connection request
		conn, err := gorm.Open(postgres, connectionString)
		if err != nil {
			if retries == maxConnectionRetryAttempts {
				logger.Error(err, errors.ErrFailedToConnectToDatabase.Error())
				return nil, err
			}
			retries += 1
		} else {
			return conn, nil
		}

		time.Sleep(1 * time.Second)
	}
	return nil, errors.ErrExceededMaxRetryAttempts
}

// configureDbConnection configures the database connection object
func configureDbConnection(conn *gorm.DB) *gorm.DB {
	conn.SingularTable(true)
	conn.LogMode(false)
	conn = conn.Set("gorm:auto_preload", true)
	return conn
}

// MigrateSchemas creates or updates a given set of models based on a schema
// if it does not exist or migrates the model schemas to the latest version
func MigrateSchemas(db *gorm.DB, logger core_logging.ILog, models ...interface{}) error {
	migration := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20200416",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(models...).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(models...).Error
			},
		},
	})

	err := migration.Migrate()
	if err != nil {
		// TODO: emit metric
		logger.ErrorM(err, errors.ErrFailedToPerformDatabaseMigrations.Error())
		return err
	}

	return nil
}

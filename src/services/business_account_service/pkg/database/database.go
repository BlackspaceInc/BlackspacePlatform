package database

import (
	"context"
	"os"
	"time"

	core_logging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-logging/json"
	core_metrics "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-metrics"
	core_tracing "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-tracing"
	"github.com/giantswarm/retry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	svcErrors "github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/proto"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/saga"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/utils"
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
	Saga                                     *saga.SagaCoordinator
	MaxConnectionAttempts                    int
	MaxRetriesPerOperation                   int
	RetryTimeOut                             time.Duration
	OperationSleepInterval                   time.Duration
}

// Tx is a type serving as a function decorator for common database transactions
type Tx func(ctx context.Context, tx *gorm.DB) error

// CmplxTx is a type serving as a function decorator for complex database transactions
type CmplxTx func(ctx context.Context, tx *gorm.DB) (interface{}, error)

// database operation types
const (
	DB_CONNECTION_ATTEMPT = "DB_CONNECTION_ATTEMPT"
)

// New creates a database connection and returns the connection object
func New(ctx context.Context, connectionString string, tracingEngine *core_tracing.TracingEngine, metricsEngine *core_metrics.CoreMetricsEngine,
	logger core_logging.ILog, svcEndpoint string) (*Db,
	error) {
	// generate a span for the database connection
	ctx, span := utils.StartRootOperationSpan(ctx, DB_CONNECTION_ATTEMPT, tracingEngine, logger)
	defer span.Finish()

	if connectionString == utils.EMPTY || tracingEngine == nil || metricsEngine == nil || logger == nil {
		// crash the process
		os.Exit(1)
	}

	logger.Info("Attempting database connection operation")
	conn, err := ConnectToDb(connectionString, logger)
	if err != nil {
		logger.FatalM(err, svcErrors.ErrFailedToConnectToDatabase.Error())
	}
	logger.Info("Successfully connected to the database")

	// configure db
	logger.Info("Attempting database connection configuration")
	conn = configureDbConnection(conn)
	logger.Info("Successfully configured database connection object")

	logger.Info("Attempting database schema migration")
	err = MigrateSchemas(conn, logger, &proto.BusinessAccountORM{}, &proto.MediaORM{}, &proto.TopicsORM{})
	if err != nil {
		logger.FatalM(err, svcErrors.ErrFailedToPerformDatabaseMigrations.Error())
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
		Saga:                                     saga.NewSagaCoordinator(logger),
		MaxConnectionAttempts:                    5,
		MaxRetriesPerOperation:                   3,
		RetryTimeOut:                             1 * time.Millisecond,
		OperationSleepInterval:                   1 * time.Second,
	}, nil
}

// ConnectToDb attempts to connect to the database using retries
func ConnectToDb(connectionString string, logger core_logging.ILog) (*gorm.DB, error) {
	var connection = make(chan *gorm.DB, 1)

	err := retry.Do(
		func(conn chan<- *gorm.DB) func() error {
			return func() error {
				newConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
				conn <- newConn
				return err
			}
		}(connection),
		retry.MaxTries(5),
		retry.Timeout(time.Second*10),
		retry.Sleep(1*time.Second),
	)

	if err != nil {
		logger.Error(err, svcErrors.ErrExceededMaxRetryAttempts.Error())
		return nil, svcErrors.ErrExceededMaxRetryAttempts
	}

	return <-connection, nil
}

// configureDbConnection configures the database connection object
func configureDbConnection(conn *gorm.DB) *gorm.DB {
	conn.FullSaveAssociations = true
	conn.SkipDefaultTransaction = false
	conn = conn.Set("gorm:auto_preload", true)
	return conn
}

// MigrateSchemas creates or updates a given set of model based on a schema
// if it does not exist or migrates the model schemas to the latest version
func MigrateSchemas(db *gorm.DB, logger core_logging.ILog, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		// TODO: emit metric
		logger.ErrorM(err, svcErrors.ErrFailedToPerformDatabaseMigrations.Error())
		return err
	}

	return nil
}

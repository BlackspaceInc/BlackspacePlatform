package errors

import (
	"errors"
)

var (
	ErrFailedToConnectToDatabase = errors.New("failed to connect to database")
	ErrFailedToPerformDatabaseMigrations = errors.New("failed to perform database migrations")
	ErrInvalidInputArguments = errors.New("invalid input arguments")
	ErrInvalidEnvironmentVariableConfigurations = errors.New("invalid environment variable configurations")
	ErrFailedToStartGRPCServer = errors.New("failed to start grpc server")
	ErrHttpServerFailedGracefuleShutdown = errors.New("http server failed to perform graceful shutdown")
	ErrHttpsServerFailedGracefuleShutdown = errors.New("https server failed to perform graceful shutdown")
	ErrHttpServerCrashed = errors.New("Http Server crashed")
	ErrHttpsServerCrashed = errors.New("Https Server crashed")
	ErrSwaggerGenError = errors.New("swagger generation error")
	ErrFailedToWatchConfigDirectory = errors.New("failed to watch config directory")
	ErrExceededMaxRetryAttempts = errors.New("exceeded max retry attemps")
)

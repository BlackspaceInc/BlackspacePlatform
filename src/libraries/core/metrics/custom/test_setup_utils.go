package custom

import (
	"math/rand"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var connSettings = "postgresql://doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db.ondigitalocean.com:25060/test-db?sslmode=require"
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// init initializes a connection to the database initially and performs package level
// cleanup handler initialization
func SetupTests() *gorm.DB {
	testDbInstance := Setup()
	if testDbInstance == nil {
		os.Exit(1)
	}

	return testDbInstance
}

// Setup sets up database connection prior to testing
func Setup() *gorm.DB{
	// database connection string
	// initialize connection to the database
	db := Initialize(connSettings)
	return db
}

// Initialize creates a singular connection to the backend database instance
func Initialize(connSettings string) *gorm.DB {
	var err error
	// configure logging
	logger := zap.L()
	defer logger.Sync()
	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	// connect to database
	dbInstance, err := newDbConnection(connSettings, logger)
	if err != nil {
		logger.Info("Error connecting to database", zap.Error(err))
		os.Exit(1)
	}

	return dbInstance
}

// New creates a database connection and returns the connection object
func newDbConnection(connString string, logger *zap.Logger) (*gorm.DB, error) {
	conn, err := gorm.Open("postgres", connString)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Successfully connected to the database")

	conn.SingularTable(true)
	conn.LogMode(false)
	conn = conn.Set("gorm:auto_preload", true)

	logger.Info("Migrating database schema")

	logger.Info("Successfully migrated database")

	return conn, nil
}

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

package database

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/database"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

// TODO - change this later should be using local database from container
var connSettings = "postgresql://doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db." +
	"ondigitalocean.com:25060/test-db?sslmode=require"

// init initializes a connection to the database initially and performs package level
// cleanup handler initialization
func SetupTests() *database.Db {
	testDbInstance := Setup()
	if testDbInstance == nil {
		os.Exit(1)
	}

	return testDbInstance
}

// Setup sets up database connection prior to testing
func Setup() *database.Db {
	// database connection string
	// initialize connection to the database
	db := Initialize(connSettings)
	// spin up/migrate tables for testing
	_ = database.MigrateSchemas(db.Engine, db.Logger)
	return db
}

// Initialize creates a singular connection to the backend database instance
func Initialize(connSettings string) *database.Db {
	var err error
	// configure logging
	logger := zap.L()
	defer logger.Sync()
	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	// connect to database
	dbInstance, err := database.New(connSettings, logger)
	if err != nil {
		logger.Info("Error connecting to database", zap.Error(err))
		os.Exit(1)
	}

	return dbInstance
}

// DeleteCreatedEntities sets up GORM `onCreate` hook and return a function that can be deferred to
// remove all the entities created after the hook was set up
// You can use it as
//
// func TestSomething(t *testing.T){
//     db, _ := gorm.Open(...)
//
//     cleaner := DeleteCreatedEntities(db)
//     defer cleaner()
//
// }
func DeleteCreatedEntities(db *gorm.DB) func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}

// GenerateTestUser creates a test user
func NewUser(firstname, lastname, email, username, password string, randomize bool) (*models.User, error) {
	str, err := NewUUID()
	if err != nil {
		return nil, err
	}

	randStr := ""
	if randomize {
		randStr = str
	}

	return &models.User{
		Id:                   0,
		CreatedAt:            nil,
		DeletedAt:            nil,
		UpdatedAt:            nil,
		FirstName:            firstname + randStr,
		LastName:             lastname + randStr,
		Gender:               nil,
		Email:                email + randStr + "@email.com",
		Password:             password + randStr,
		BirthDate:            "",
		IsActive:             false,
		IsOnline:             false,
		Username:             username + randStr,
		Authnid:              0,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random)
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

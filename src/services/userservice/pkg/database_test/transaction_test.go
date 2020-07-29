package database_test

import (
	"errors"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

// TestTransaction Tests the result of a transaction
func TestTransaction(t *testing.T) {
	// success scenarios
	t.Run("TestName:Passed_TransactionPassed", TransactionPassed)
	// failure scenarios
	t.Run("TestName:Failed_Transaction", TransactionFailed)
}

func TestComplexTransaction(t *testing.T) {
	// success scenarios
	t.Run("TestName:Passed_ComplextTransaction", ComplextTransactionPassed)
	// failure scenarios
	t.Run("TestName:Failed_ComplextTransaction", ComplextTransactionFailed)
}

func TransactionFailed(t *testing.T) {
	f := func(tx *gorm.DB) error {
		return errors.New("failed transactions")
	}

	err := db.PerformTransaction(f)
	assert.NotEmpty(t, err)
}

func ComplextTransactionFailed(t *testing.T) {
	f := func(tx *gorm.DB) (interface{}, error) {
		return nil, errors.New("failed transactions")
	}

	output, err := db.PerformComplexTransaction(f)
	assert.NotEmpty(t, err)
	assert.Empty(t, output)
}

func TransactionPassed(t *testing.T) {
	f := func(tx *gorm.DB) error {
		return nil
	}

	err := db.PerformTransaction(f)
	assert.Empty(t, err)
}

func ComplextTransactionPassed(t *testing.T) {
	f := func(tx *gorm.DB) (interface{}, error) {
		return true, nil
	}

	output, err := db.PerformComplexTransaction(f)
	assert.Empty(t, err)
	assert.NotEmpty(t, output)
}

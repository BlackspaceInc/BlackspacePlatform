package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBusinessAccountActiveStatus(t *testing.T){
	t.Run("TestName:SetAccountInactive", SetAccountInactive)
	t.Run("TestName:SetAccountActive", SetAccountActive)
}

// SetAccountActive test that an account can be set as active
func SetAccountActive(t *testing.T){
	ctx := context.TODO()
	var authnId uint32 = uint32(GenerateRandomId(20, 100))
	account := GenerateRandomizedAccount()
	// create account first
	result, err := db.CreateBusinessAccount(ctx, account, authnId)
	ExpectNoErrorOccured(t, err, result)

	// update account
	err = db.SetBusinessAccountStatusAndSave(ctx, result,true)
	ExpectNoErrorOccured(t, err, result)

	obtainedAccount, err := db.GetBusinessAccount(ctx, result.Id)
	ExpectNoErrorOccured(t, err, obtainedAccount)
	assert.True(t, obtainedAccount.IsActive)
}

// SetAccountInactive test that an account can be set as active
func SetAccountInactive(t *testing.T){
	ctx := context.TODO()
	var authnId uint32 = uint32(GenerateRandomId(20, 100))
	account := GenerateRandomizedAccount()
	// create account first
	result, err := db.CreateBusinessAccount(ctx, account, authnId)
	ExpectNoErrorOccured(t, err, result)

	// update account
	err = db.SetBusinessAccountStatusAndSave(ctx, result,false)
	ExpectNoErrorOccured(t, err, result)

	obtainedAccount, err := db.GetBusinessAccount(ctx, result.Id)
	ExpectNoErrorOccured(t, err, obtainedAccount)
	assert.False(t, obtainedAccount.IsActive)
}

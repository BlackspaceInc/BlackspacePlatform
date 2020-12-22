package database_test

import (
	"context"
	"testing"
)

func TestGetBusinessAccount(t *testing.T){
	t.Run("TestName:GetBusinessAccount", GetBusinessAccount)
	t.Run("TestName:GetBusinessAccountByEmail", GetBusinessAccountByEmail)
	t.Run("TestName:GetBusinessAccountById", GetBusinessAccountById)
	t.Run("TestName:GetBusinessAccountDoesntExist", GetBusinessAccountDoesntExist)
}

// GetBusinessAccount test that an account can be set as obtained correctly
func GetBusinessAccount(t *testing.T){
	ctx := context.TODO()
	var authnId uint32 = uint32(GenerateRandomId(20, 100))
	account := GenerateRandomizedAccount()
	// create account first
	result, err := db.CreateBusinessAccount(ctx, account, authnId)
	ExpectNoErrorOccured(t, err, result)

	// update account
	obtainedAccount, err := db.GetBusinessAccount(ctx, result.Id)
	ExpectValidAccountObtained(t, err, obtainedAccount, result)
}

// GetBusinessAccountByEmail test that an account can be obtained by email correctly
func GetBusinessAccountByEmail(t *testing.T){
	ctx := context.TODO()
	var authnId uint32 = uint32(GenerateRandomId(20, 100))
	account := GenerateRandomizedAccount()
	// create account first
	result, err := db.CreateBusinessAccount(ctx, account, authnId)
	ExpectNoErrorOccured(t, err, result)

	obtainedAccount := db.GetBusinessByEmail(ctx, result.Email)
	ExpectValidAccountObtained(t, nil, obtainedAccount, result)
}

// GetBusinessAccountById ensures that we can obtain an account by id
func GetBusinessAccountById(t *testing.T){
	ctx := context.TODO()
	var authnId uint32 = uint32(GenerateRandomId(20, 100))
	account := GenerateRandomizedAccount()
	// create account first
	result, err := db.CreateBusinessAccount(ctx, account, authnId)
	ExpectNoErrorOccured(t, err, result)

	obtainedAccount := db.GetBusinessById(ctx, result.Id)
	ExpectValidAccountObtained(t, nil, obtainedAccount, result)
}

// GetBusinessAccountDoesntExist ensures we obtain the proper error when attempting to get an account that does not exist
func GetBusinessAccountDoesntExist(t *testing.T){
	ctx := context.TODO()

	// generate random id
	id := GenerateRandomId(500, 1000)
	obtainedAccount, err := db.GetBusinessAccount(ctx, uint32(id))
	ExpectAccountDoesNotExistError(t, err, obtainedAccount)
}

package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/database"
	svcErrors "github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/proto"
)

func TestBusinessCreateAccount(t *testing.T) {
	t.Run("TestName:CreateBusinessAccount", CreateAccountNonExistentAccount)
	t.Run("TestName:CreateDuplicateBusinessAccounts", CreateDuplicateBusinessAccounts)
	t.Run("TestName:CreateBusinessAccountWithFaultyAuthnId", CreateBusinessAccountWithFaultyAuthnId)
	t.Run("TestName:CreateBusinessAccountWithNoCompanyName", CreateBusinessAccountWithNoCompanyName)
	t.Run("TestName:CreateBusinessAccountWithNoEmail", CreateBusinessAccountWithNoEmail)
	t.Run("TestName:CreateBusinessAccountWithNoPassword", CreateAccountNonExistentAccount)
}

// CreateAccountNonExistentAccount test an account is properly created in the backend database if a record
// doesnt already exist
func CreateAccountNonExistentAccount(t *testing.T) {
	account := GenerateRandomizedAccount()

	createdAccount, err := db.CreateBusinessAccount(context.Background(), account, 1)
	assert.Empty(t, err)
	assert.NotEmpty(t, createdAccount, "user record cannot be empty")
	// assert the user record returned is actually active
	assert.Equal(t, createdAccount.IsActive, true, "user record should be activated")
}

// CreateDuplicateBusinessAccounts tests that duplicate accounts can have the expected error return type
func CreateDuplicateBusinessAccounts(t *testing.T){
	account := GenerateRandomizedAccount()

	var authnId uint32 = 500
	// create account twice
	createdAccount, err := db.CreateBusinessAccount(context.Background(), account, authnId)
	assert.Empty(t, err)

	createdAccount, err = db.CreateBusinessAccount(context.Background(), createdAccount, authnId)
	// an error should occur
	ExpectAccountAlreadyExistError(t, err, createdAccount)
}

// CreateBusinessAccountWithFaultyAuthnId tests the proper errors are returned for accounts with faulty authnn Ids
func CreateBusinessAccountWithFaultyAuthnId(t *testing.T){
	account := GenerateRandomizedAccount()

	var authnId uint32 = 0
	// create account with faulty Id and ensure the proper expected error is returned
	createdAccount, err := db.CreateBusinessAccount(context.Background(), account, authnId)
	ExpectInvalidArgumentsError(t, err, createdAccount)
}

// CreateBusinessAccountWithNoCompanyName tests the proper errors are returned for accounts with no company name
func CreateBusinessAccountWithNoCompanyName(t *testing.T){
	account := GenerateRandomizedAccount()
	// remove company name
	account.CompanyName = ""

	var authnId uint32 = 501
	createdAccount, err := db.CreateBusinessAccount(context.Background(), account, authnId)
	ExpectInvalidArgumentsError(t, err, createdAccount)
}

// CreateBusinessAccountWithNoEmail tests the proper errors are returned for accounts with no email
func CreateBusinessAccountWithNoEmail(t *testing.T){
	account := GenerateRandomizedAccount()
	// remove company name
	account.Email = ""

	var authnId uint32 = 502
	createdAccount, err := db.CreateBusinessAccount(context.Background(), account, authnId)
	ExpectInvalidArgumentsError(t, err, createdAccount)
}

// CreateBusinessAccountWithNoPassword tests the proper errors are returned for accounts with no password
func CreateBusinessAccountWithNoPassword(t *testing.T){
	account := GenerateRandomizedAccount()
	// remove company name
	account.Password = ""

	var authnId uint32 = 503
	createdAccount, err := db.CreateBusinessAccount(context.Background(), account, authnId)
	ExpectInvalidArgumentsError(t, err, createdAccount)
}

// ExpectInvalidArgumentsError ensure the invalid error is present
func ExpectInvalidArgumentsError(t *testing.T, err error, account *proto.BusinessAccount) {
	assert.NotEmpty(t, err)
	assert.EqualError(t, err, svcErrors.ErrInvalidInputArguments.Error())
	assert.Nil(t, account)
}

// ExpectAccountAlreadyExistError ensures the account already exist error is present
func ExpectAccountAlreadyExistError(t *testing.T, err error, createdAccount *proto.BusinessAccount) {
	assert.NotEmpty(t, err)
	assert.EqualError(t, err, svcErrors.ErrAccountAlreadyExist.Error())
	assert.Nil(t, createdAccount)
}

// ExpectAccountDoesNotExistError ensures the account does not exist error is present
func ExpectAccountDoesNotExistError(t *testing.T, err error, createdAccount *proto.BusinessAccount) {
	assert.NotEmpty(t, err)
	assert.EqualError(t, err, svcErrors.ErrAccountDoesNotExist.Error())
	assert.Nil(t, createdAccount)
}

// ExpectCannotUpdatePasswordError ensure the invalid error is present
func ExpectCannotUpdatePasswordError(t *testing.T, err error, account *proto.BusinessAccount) {
	assert.NotEmpty(t, err)
	assert.EqualError(t, err, svcErrors.ErrCannotUpdatePassword.Error())
	assert.Nil(t, account)
}

func GenerateRandomizedAccount() *proto.BusinessAccount {
	randStr := database.GenerateRandomString(15)
	account := testBusinessAccount
	account.Email = account.Email + randStr
	account.CompanyName = account.CompanyName + randStr
	return account
}

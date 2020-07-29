package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const passwordtoHash string = "teststring"
const emptyPasswordToHash string = ""

type MockInterface struct{}
type FaultyMockInterface struct{}

func (t MockInterface) GetPasswordToHash() string {
	return passwordtoHash
}

func (t FaultyMockInterface) GetPasswordToHash() string {
	return emptyPasswordToHash
}

// TestValidateAndHashPassword test whether a password is properly validated and hashed
func TestValidateAndHashPassword(t *testing.T) {
	t.Run("TestName:Passed_TestValidateAndHashPassword", ValidateAndHashPasswordValidPassword)
	t.Run("TestName:Failed_TestInValidateAndHashPassword", ValidateAndHashPasswordInValidPassword)
}

// TestComparePassword tests whether a password can be adequately compared
func TestComparePassword(t *testing.T) {
	t.Run("TestName:ComparePasswords", ComparePasswords)
}

// ValidateAndHashPasswordValidPassword Tests a valid password
func ValidateAndHashPasswordValidPassword(t *testing.T) {
	// arrange
	var (
		mock = MockInterface{}
	)

	// act
	hashPassword, err := db.ValidateAndHashPassword(mock)

	// assert
	assert.Empty(t, err)
	assert.NotEmpty(t, hashPassword)
}

// ValidateAndHashPasswordInValidPassword Tests wether a given invalid hashed password
func ValidateAndHashPasswordInValidPassword(t *testing.T) {
	// arrange
	var (
		mock = FaultyMockInterface{}
	)

	// act
	hashPassword, err := db.ValidateAndHashPassword(mock)

	// assert
	assert.NotEmpty(t, err)
	assert.Empty(t, hashPassword)
}

// ComparePasswords Tests if a hashed password is equal to a plain counterpart
func ComparePasswords(t *testing.T) {
	// arrange
	var (
		mock = MockInterface{}
	)

	// act
	hashPassword, err := db.ValidateAndHashPassword(mock)
	// assert
	assert.Empty(t, err)
	assert.NotEmpty(t, hashPassword)

	valid := db.ComparePasswords(hashPassword, []byte(mock.GetPasswordToHash()))
	assert.True(t, valid)

	valid = db.ComparePasswords(hashPassword, []byte("random string for password"))
	assert.False(t, valid)
}

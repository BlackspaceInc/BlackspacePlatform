package database_test

import (
	"context"
	"math/rand"
	// "os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/database"
	dbTestUtil "github.com/BlackspaceInc/Backend/user-management-service/pkg/database_test"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

var (
	db        *database.Db
	firstname = "test"
	lastname  = "test"
	password  = "test"
	email     = "test"
	username  = "test"
)

func TestMain(m *testing.M) {
	db = setupTestEnvironment()
	defer db.Engine.Close()

	// defer deleting all created entries
	cleanupHandler := dbTestUtil.DeleteCreatedEntities(db.Engine)
	defer cleanupHandler()

	_ = m.Run()
	return
	// os.Exit(code)
}

// setupTestEnvironment sets up test environment
func setupTestEnvironment() *database.Db {
	// prepare database connection and cleanup handler
	return dbTestUtil.SetupTests()
}

// TestCreateUser Tests The Create User Db Operation
func TestCreateUser(t *testing.T) {
	t.Run("TestName:Passed_UserDoesNotInitiallyExist", CreateUserNonExistentUser)
	t.Run("TestName:Passed_UserExistsAndInactive", CreateUserInactiveExistingUser)
	// failure scenarios
	t.Run("TestName:Failed_UserAlreadyExists", CreateUserUserAlreadyExists)
	t.Run("TestName:Failed_InvalidUserParameters", CreateUserInvalidParams)
}

// TestUpdateUser Test Update User Db Operation
func TestUpdateUser(t *testing.T) {
	// success scenario
	// Update Existing User
	t.Run("TestName:Passed_UpdateExistingUser", UpdateUserExistingRecord)
	// failure scenario
	// Update Non Existing User
	t.Run("TestName:Failed_UpdateNonExistingUser", UpdateUserNonExistingRecord)
}

// TestGetUser Test Get User Db Operation
func TestGetUser(t *testing.T) {
	// success scenario
	// Get An Existing User
	t.Run("TestName:Passed_GetExistingUser", GetUserExistingRecord)
	// failure scenario
	// Get a non existing user (random id)
	t.Run("TestName:Failed_GetNonExistingUser", GetUserNonExistingRecord)
}

// TestDeleteUser Test Delete User Db Operation
func TestDeleteUser(t *testing.T) {
	// success scenario
	// Delete a user record that exists
	t.Run("TestName:Passed_DeleteExistingUser", DeleteUserExistingRecord)
	// failure scenario
	// Delete a user record that does not exist
	t.Run("TestName:Failed_DeleteNonExistingUser", DeleteUserNonExistingRecord)
}

// TestGetUserIfExists tests if a user exists
func TestGetUserIfExists(t *testing.T) {
	// success scenario
	t.Run("TestName:Passed_GetUserIfExistWithValidParams", GetUserIfExistsValidParams)

	// failure scenario
	t.Run("TestName:Failed_GetUserIfExistWithInValidParams", GetUserIfExistsInvalidParams)
	t.Run("TestName:Failed_GetUserIfExistsNonExistentUser", GetUserIfExistsNonExistentUser)
}

// GetUserIfExistsValidParams Attempts to get a user if it exists under the condition that its
// parameters are valid
func GetUserIfExistsValidParams(t *testing.T) {
	// arrange
	ctx := context.TODO()
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	// act
	userOrm := UserToOrm(ctx, t, user)
	err = db.Engine.Save(&userOrm).Error
	assert.Empty(t, err)

	var tests = []struct {
		TestName string
		Id       uint32
		Email    string
		Username string
	}{
		{"ValidIdEmailUsernamer", userOrm.Id, userOrm.Email, userOrm.Username},
		{"ValidEmailUsername", 0, userOrm.Email, userOrm.Username},
		{"ValidUsername", 0, "", userOrm.Username},
		{"ValidEmail", 0, userOrm.Email, ""},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			// assert
			// get user by id
			exists, userRecord, err := db.GetUserIfExists(ctx, tt.Id, tt.Username, tt.Email)
			assert.Empty(t, err)
			assert.True(t, exists)
			assert.NotEmpty(t, userRecord)
		})
	}
}

// GetUserIfExistsInvalidParams attempts to get a user if it exists based on invalid input
// parameters
func GetUserIfExistsInvalidParams(t *testing.T) {
	var randomStr = "randomstring for testing purposes"
	var randomId = GenerateRandomId()
	// arrange
	ctx := context.TODO()
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	// act
	userOrm := UserToOrm(ctx, t, user)
	err = db.Engine.Save(&userOrm).Error
	assert.Empty(t, err)

	var tests = []struct {
		TestName string
		Id       uint32
		Email    string
		Username string
	}{
		{"ValidIdEmailUsernamer", randomId, randomStr, randomStr},
		{"ValidEmailUsername", 0, randomStr, randomStr},
		{"ValidUsername", 0, "", randomStr},
		{"ValidEmail", 0, randomStr, ""},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			// assert
			// get user by id
			exists, userRecord, err := db.GetUserIfExists(ctx, tt.Id, tt.Username, tt.Email)
			assert.NotEmpty(t, err)
			assert.False(t, exists)
			assert.Empty(t, userRecord)
		})
	}
}

// GetUserIfExistsNonExistentUser Attempts to get a user that does not exist
func GetUserIfExistsNonExistentUser(t *testing.T) {
	var randomId = GenerateRandomId()
	ctx := context.TODO()
	// arrange
	exists, userRecord, err := db.GetUserIfExists(ctx, randomId, "", "")
	assert.Nil(t, userRecord)
	assert.NotEmpty(t, err)
	assert.False(t, exists)
}

// DeleteUserNonExistingRecord attempts to delete a user record that does not exist
func DeleteUserNonExistingRecord(t *testing.T) {
	// arrange
	// generate a random id
	ctx := context.TODO()
	id := GenerateRandomId()

	// act
	err := db.DeleteUser(ctx, id)

	// assert
	assert.NotEmpty(t, err)
}

// DeleteUserExistingRecord deletes an existing user record
func DeleteUserExistingRecord(t *testing.T) {
	// arrange
	ctx := context.TODO()
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	// act
	// save the record in the db
	userRecord := UserToOrm(ctx, t, user)
	err = db.Engine.Save(&userRecord).Error
	assert.Empty(t, err)

	// attempt deletion
	err = db.DeleteUser(ctx, userRecord.Id)
	assert.Empty(t, err)

	// attempt to obtain the user which should fail
	err = db.Engine.Find(&models.UserORM{Id: userRecord.Id}).Error
	assert.NotEmpty(t, err)
}

// GetUserNonExistingRecord attempts to get a non existing user record from the backend
func GetUserNonExistingRecord(t *testing.T) {
	// arrange
	// generate a random id
	ctx := context.TODO()

	// act
	userRecord, err := db.GetUser(ctx, 0)

	// assert
	assert.NotEmpty(t, err)
	assert.Nil(t, userRecord)
}

// GenerateRandomId generates a random id value
func GenerateRandomId() uint32 {
	rand.Seed(time.Now().UnixNano())
	id := rand.Uint32() + 1000
	return id
}

// GetUserExistingRecord attempts to get existing record from db
func GetUserExistingRecord(t *testing.T) {
	// arrange
	// create a user object
	ctx := context.TODO()
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	// convert user to orm type
	userOrm := UserToOrm(ctx, t, user)

	// save the user record in the db
	err = db.Engine.Save(&userOrm).Error
	assert.Empty(t, err)

	// attempt to obtain the user now by id
	userRecord, err := db.GetUser(ctx, userOrm.Id)
	assert.Empty(t, err)
	assert.NotEmpty(t, userRecord)

	// assert the user names and emails are the same since they can only be unique in the system
	assert.Equal(t, userOrm.Email, userRecord.Email, "email fields must be the same")
	assert.Equal(t, userOrm.Username, userRecord.Username, "username fields must be the same")
}

// Converts a user object to orm type
func UserToOrm(ctx context.Context, t *testing.T, user *models.User) models.UserORM {
	userOrm, err := user.ToORM(ctx)
	assert.Empty(t, err)
	assert.NotEmpty(t, userOrm)
	return userOrm
}

// UpdateUserNonExistingRecord attempts to update a non existing user record
func UpdateUserNonExistingRecord(t *testing.T) {
	// arrange
	// create test user object
	ctx := context.TODO()
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	updatedUser, err := db.UpdateUser(ctx, user)
	assert.NotEmpty(t, err)
	assert.Empty(t, updatedUser)
}

// UpdateUserExistingRecord attempts to update an existing user record
func UpdateUserExistingRecord(t *testing.T) {
	// arrange
	var testUsername = "test"
	// create the user in the backend store
	ctx := context.TODO()
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	userOrm, err := user.ToORM(ctx)
	assert.NotEmpty(t, userOrm, "user record cannot be empty")
	assert.Empty(t, err)

	// act
	// save the user instance in the database and attempt to update
	err = db.Engine.Save(&userOrm).Error
	assert.Empty(t, err)

	// convert userOrm to user type to perform update function call
	userOrm.IsActive = false
	userOrm.Username = testUsername
	transformedUser, err := userOrm.ToPB(ctx)

	updatedUser, err := db.UpdateUser(ctx, &transformedUser)

	// assert
	assert.Empty(t, err)
	assert.NotEmpty(t, updatedUser, "updated user record should not be null")
	assert.Equal(t, updatedUser.Username, testUsername, "usernames must match")
	assert.Equal(t, updatedUser.IsActive, false, "user active status must be false")
}

// CreateUserUserAlreadyExists tests the create user database call under the assumption
// that the user record already exists. Should fail
func CreateUserUserAlreadyExists(t *testing.T) {
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	ctx := context.TODO()

	createdUser, err := db.CreateUser(ctx, user)
	assert.Empty(t, err)
	assert.NotEmpty(t, createdUser, "user record should not be empty")

	// reattempt creation of already present user record which is inactive
	createdUser, err = db.CreateUser(ctx, user)
	assert.NotEmpty(t, err)
	assert.Empty(t, createdUser)
}

// CreateUserInvalidParams attempts to create a user with invalid parameters
func CreateUserInvalidParams(t *testing.T) {
	empty := ""
	user, err := dbTestUtil.NewUser(empty, empty, empty, empty, empty, false)
	assert.Empty(t, err)

	ctx := context.TODO()

	createdUser, err := db.CreateUser(ctx, user)
	assert.NotEmpty(t, err)
	assert.Empty(t, createdUser, "user record should be empty")
}

// TestCreateUserNonExistentUser test a user is properly created in the backend database if he/she
// doesnt already exist
func CreateUserNonExistentUser(t *testing.T) {
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	ctx := context.TODO()

	createdUser, err := db.CreateUser(ctx, user)
	assert.Empty(t, err)

	assert.NotEmpty(t, createdUser, "user record cannot be empty")
	// assert the user record returned is actually active
	assert.Equal(t, createdUser.IsActive, true, "user record should be activated")
	assert.Equal(t, createdUser.Username, user.Username, "user fields do not match")
}

// TestCreateUserNonExistentUser test a user is properly created in the backend database if he/she
// doesnt already exist
func CreateUserInactiveExistingUser(t *testing.T) {
	user, err := dbTestUtil.NewUser(firstname, lastname, email, username, password, true)
	assert.Empty(t, err)

	ctx := context.TODO()

	createdUser, err := db.CreateUser(ctx, user)
	assert.Empty(t, err)

	// make user inactive
	createdUser.IsActive = false
	err = db.Engine.Save(&createdUser).Error
	assert.Empty(t, err)

	// reattempt creation of already present user record which is inactive
	updatedUser, err := db.CreateUser(ctx, user)
	assert.NotEmpty(t, updatedUser, "user record cannot be empty")
	assert.Equal(t, updatedUser.IsActive, true,
		"ser already present should be returned and be made active after create user db call")
}

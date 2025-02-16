package repository

import (
	"context"
	"testing"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/server/security"
	"github.com/stretchr/testify/assert"
)

var defaultTestingUser = model.User{
	Name:     "Test User",
	Email:    "test@example.com",
	Password: "password",
}

func TestCreateUser(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	t.Run("it can create a new user", func(t *testing.T) {
		err := CreateUser(ctx, defaultTestingUser)
		assert.NoError(t, err)

		createdUser, err := FindUserByEmail(ctx, defaultTestingUser.Email)
		assert.NoError(t, err)
		assert.Equal(t, defaultTestingUser.Name, createdUser.Name)
		assert.NotEmpty(t, createdUser.Password)
	})

	t.Run("it encrypts the user password", func(t *testing.T) {
		err := CreateUser(ctx, defaultTestingUser)
		assert.NoError(t, err)

		createdUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)
		assert.NotEqual(t, defaultTestingUser.Password, createdUser.Password)
		assert.True(t, security.VerifyPassword(createdUser.Password, defaultTestingUser.Password))
	})
}

func TestFindUserByID(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	CreateUser(ctx, defaultTestingUser)
	createdUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)

	t.Run("it can find a user by their ID", func(t *testing.T) {
		foundUser, err := FindUserByID(ctx, createdUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, createdUser.Email, foundUser.Email)
	})

	t.Run("it returns an error when the user does not exist", func(t *testing.T) {
		_, err := FindUserByID(ctx, 999)
		assert.Error(t, err)
	})
}

func TestFindUserByEmail(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	CreateUser(ctx, defaultTestingUser)

	t.Run("it can find a user by their email", func(t *testing.T) {
		foundUser, err := FindUserByEmail(ctx, defaultTestingUser.Email)
		assert.NoError(t, err)
		assert.Equal(t, defaultTestingUser.Name, foundUser.Name)
	})

	t.Run("it returns an error when the user does not exist", func(t *testing.T) {
		_, err := FindUserByEmail(ctx, "invalid-user@example.com")
		assert.Error(t, err)
	})
}

func TestUserExistsByEmail(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	CreateUser(ctx, defaultTestingUser)

	t.Run("it returns true when the user exists", func(t *testing.T) {
		exists := UserExistsByEmail(ctx, defaultTestingUser.Email)
		assert.True(t, exists)
	})

	t.Run("it returns false when the user does not exist", func(t *testing.T) {
		exists := UserExistsByEmail(ctx, "invalid-user@example.com")
		assert.False(t, exists)
	})
}

func TestUpdateUser(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	CreateUser(ctx, defaultTestingUser)

	t.Run("it can update a user", func(t *testing.T) {
		beforeUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)

		err := UpdateUser(ctx, beforeUser, model.User{
			Name: "Updated User",
		})
		assert.NoError(t, err)

		afterUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)
		assert.Equal(t, "Updated User", afterUser.Name)
	})

	t.Run("it only updates the provided fields", func(t *testing.T) {
		beforeUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)

		err := UpdateUser(ctx, beforeUser, model.User{
			Email: "new-email@example.com",
		})
		assert.NoError(t, err)

		afterUser, _ := FindUserByEmail(ctx, "new-email@example.com")
		assert.Equal(t, beforeUser.Name, afterUser.Name)
		assert.Equal(t, beforeUser.Password, afterUser.Password)
	})

	t.Run("it re-hashes the password when it is updated", func(t *testing.T) {
		beforeUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)

		err := UpdateUser(ctx, beforeUser, model.User{
			Password: "new-password",
		})
		assert.NoError(t, err)

		afterUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)
		assert.True(t, security.VerifyPassword(afterUser.Password, "new-password"))
		assert.False(t, security.VerifyPassword(afterUser.Password, "password"))
	})
}

func TestDeleteUserAndRelatedRecordsById(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	CreateUser(ctx, defaultTestingUser)

	t.Run("it can delete a user and their related records", func(t *testing.T) {
		createdUser, _ := FindUserByEmail(ctx, defaultTestingUser.Email)
		err := DeleteUserAndRelatedRecordsById(ctx, createdUser.ID)
		assert.NoError(t, err)

		_, err = FindUserByID(ctx, createdUser.ID)
		assert.Error(t, err)
	})
}

package repository

import (
	"context"
	"testing"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/stretchr/testify/assert"
)

func setupPriorityTesting() (context.Context, model.User, func()) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	ctx := context.Background()

	user := model.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	database.
		GetConnectionWithContext(ctx).
		Create(&user)

	createdUser, _ := FindUserByEmail(ctx, user.Email)

	return ctx, createdUser, func() {
		database.Disconnect()
	}
}

func TestCreateDefaultPrioritiesForUserId(t *testing.T) {
	ctx, user, close := setupPriorityTesting()
	defer close()

	t.Run("it creates the default priorities for a user", func(t *testing.T) {
		err := CreateDefaultPrioritiesForUserId(ctx, user.ID)
		assert.NoError(t, err)

		priorities, err := GetPrioritiesForUserId(ctx, user.ID)
		assert.NoError(t, err)
		assert.Len(t, priorities, 4)
	})
}

func TestGetPrioritiesForUser(t *testing.T) {
	ctx, user, close := setupPriorityTesting()
	defer close()

	CreateDefaultPrioritiesForUserId(ctx, user.ID)

	t.Run("it can get priorities for a user", func(t *testing.T) {
		priorities, err := GetPrioritiesForUser(ctx, &user)
		assert.NoError(t, err)
		assert.Len(t, priorities, 4)
	})
}

func TestGetPrioritiesForUserId(t *testing.T) {
	ctx, user, close := setupPriorityTesting()
	defer close()

	CreateDefaultPrioritiesForUserId(ctx, user.ID)

	t.Run("it can get priorities for a user by ID", func(t *testing.T) {
		priorities, err := GetPrioritiesForUserId(ctx, user.ID)
		assert.NoError(t, err)
		assert.Len(t, priorities, 4)
	})
}

func TestDeletePrioritiesForUserId(t *testing.T) {
	ctx, user, close := setupPriorityTesting()
	defer close()

	CreateDefaultPrioritiesForUserId(ctx, user.ID)

	t.Run("it can delete priorities for a user by ID", func(t *testing.T) {
		err := DeletePrioritiesForUserId(ctx, user.ID)
		assert.NoError(t, err)

		priorities, err := GetPrioritiesForUserId(ctx, user.ID)
		assert.NoError(t, err)
		assert.Len(t, priorities, 0)
	})
}

package repository

import (
	"context"
	"testing"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/stretchr/testify/assert"
)

func setupCategoryTesting() (context.Context, model.User, func()) {
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

func TestCreateDefaultCategoriesForUserId(t *testing.T) {
	ctx, user, close := setupCategoryTesting()
	defer close()

	t.Run("it creates the default categories for a user", func(t *testing.T) {
		err := CreateDefaultCategoriesForUserId(ctx, user.ID)
		assert.NoError(t, err)

		categories, err := GetCategoriesForUserId(ctx, user.ID)
		assert.NoError(t, err)
		assert.Len(t, categories, 4)
	})
}

func TestGetCategoriesForUser(t *testing.T) {
	ctx, user, close := setupCategoryTesting()
	defer close()

	CreateDefaultCategoriesForUserId(ctx, user.ID)

	t.Run("it can get categories for a user", func(t *testing.T) {
		categories, err := GetCategoriesForUser(ctx, &user)
		assert.NoError(t, err)
		assert.Len(t, categories, 4)
	})
}

func TestGetCategoriesForUserId(t *testing.T) {
	ctx, user, close := setupCategoryTesting()
	defer close()

	CreateDefaultCategoriesForUserId(ctx, user.ID)

	t.Run("it can get categories for a user by ID", func(t *testing.T) {
		categories, err := GetCategoriesForUserId(ctx, user.ID)
		assert.NoError(t, err)
		assert.Len(t, categories, 4)
	})
}

func TestDeleteCategoriesForUserId(t *testing.T) {
	ctx, user, close := setupCategoryTesting()
	defer close()

	CreateDefaultCategoriesForUserId(ctx, user.ID)

	t.Run("it can delete categories for a user by ID", func(t *testing.T) {
		err := DeleteCategoriesForUserId(ctx, user.ID)
		assert.NoError(t, err)

		categories, err := GetCategoriesForUserId(ctx, user.ID)
		assert.NoError(t, err)
		assert.Len(t, categories, 0)
	})
}

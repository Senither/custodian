package repository

import (
	"context"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
)

func CreateDefaultCategoriesForUserId(ctx context.Context, userId uint) error {
	categories := []model.Category{
		{Name: "House Stuff", UserId: userId},
		{Name: "Work", UserId: userId},
		{Name: "Learning", UserId: userId},
		{Name: "Meeting", UserId: userId},
	}

	result := database.
		GetConnectionWithContext(ctx).
		Create(&categories)

	return result.Error
}

func GetCategoriesForUser(ctx context.Context, user *model.User) ([]model.Category, error) {
	return GetCategoriesForUserId(ctx, user.ID)
}

func GetCategoriesForUserId(ctx context.Context, userId uint) ([]model.Category, error) {
	var categories []model.Category

	result := database.
		GetConnectionWithContext(ctx).
		Where("user_id = ?", userId).
		Find(&categories)

	return categories, result.Error
}

func DeleteCategoriesForUserId(ctx context.Context, id uint) error {
	result := database.
		GetConnectionWithContext(ctx).
		Unscoped().
		Delete(&model.Category{}, "user_id = ?", id)

	return result.Error
}

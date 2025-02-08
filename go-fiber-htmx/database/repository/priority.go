package repository

import (
	"context"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
)

func CreateDefaultPrioritiesForUserId(ctx context.Context, userId uint) error {
	priorities := []model.Priority{
		{Name: "Low", UserId: userId},
		{Name: "Medium", UserId: userId},
		{Name: "High", UserId: userId},
		{Name: "Highest", UserId: userId},
	}

	result := database.
		GetConnectionWithContext(ctx).
		Create(&priorities)

	return result.Error
}

func GetPrioritiesForUser(ctx context.Context, user *model.User) ([]model.Priority, error) {
	return GetPrioritiesForUserId(ctx, user.ID)
}

func GetPrioritiesForUserId(ctx context.Context, userId uint) ([]model.Priority, error) {
	var priorities []model.Priority

	result := database.
		GetConnectionWithContext(ctx).
		Where("user_id = ?", userId).
		Find(&priorities)

	return priorities, result.Error
}

func DeletePrioritiesForUserId(ctx context.Context, id uint) error {
	result := database.
		GetConnectionWithContext(ctx).
		Unscoped().
		Delete(&model.Priority{}, "user_id = ?", id)

	return result.Error
}

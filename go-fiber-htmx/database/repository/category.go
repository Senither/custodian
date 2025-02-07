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

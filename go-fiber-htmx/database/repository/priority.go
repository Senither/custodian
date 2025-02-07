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

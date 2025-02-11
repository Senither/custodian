package repository

import (
	"context"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
)

func CreateTaskForUser(ctx context.Context, user *model.User, task model.Task) error {
	task.UserId = user.ID

	return database.
		GetConnectionWithContext(ctx).
		Create(&task).
		Error
}

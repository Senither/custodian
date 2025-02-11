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

func GetTasksForUserWithRelations(ctx context.Context, user *model.User) ([]model.Task, error) {
	var tasks []model.Task

	err := database.
		GetConnectionWithContext(ctx).
		Where("user_id = ?", user.ID).
		Order("status ASC").
		Order("created_at DESC").
		Preload("Category").
		Preload("Priority").
		Find(&tasks).
		Error

	return tasks, err
}

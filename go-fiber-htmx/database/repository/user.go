package repository

import (
	"context"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
)

func CreateUser(ctx context.Context, user model.User) error {
	return database.
		GetConnectionWithContext(ctx).
		Create(&user).
		Error
}

func FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	result := database.
		GetConnectionWithContext(ctx).
		Model(model.User{Email: email}).
		First(&user)

	return user, result.Error
}

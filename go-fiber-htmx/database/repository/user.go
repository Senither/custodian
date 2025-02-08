package repository

import (
	"context"

	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/server/security"
)

func CreateUser(ctx context.Context, user model.User) error {
	hash, err := security.EncryptPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	return CreateUserWithoutPasswordEncryption(ctx, user)
}

func CreateUserWithoutPasswordEncryption(ctx context.Context, user model.User) error {
	result := database.
		GetConnectionWithContext(ctx).
		Create(&user)

	if result.Error != nil {
		return result.Error
	}

	go CreateDefaultPrioritiesForUserId(ctx, user.ID)
	go CreateDefaultCategoriesForUserId(ctx, user.ID)

	return nil
}

func FindUserByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User

	result := database.
		GetConnectionWithContext(ctx).
		Model(model.User{}).
		Where("id = ?", id).
		First(&user)

	return user, result.Error
}

func FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	result := database.
		GetConnectionWithContext(ctx).
		Model(model.User{}).
		Where("email = ?", email).
		First(&user)

	return user, result.Error
}

func UserExistsByEmail(ctx context.Context, email string) bool {
	var exists bool

	database.
		GetConnectionWithContext(ctx).
		Model(model.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists)

	return exists
}

func DeleteUserById(ctx context.Context, id uint) error {
	result := database.
		GetConnectionWithContext(ctx).
		Unscoped().
		Delete(&model.User{}, id)

	return result.Error
}

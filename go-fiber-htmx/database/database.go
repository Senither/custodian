package database

import (
	"context"

	"github.com/glebarez/sqlite"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/database/model"
	"gorm.io/gorm"
)

var connection *gorm.DB

func InitiateDatabaseConnection() error {
	con, err := gorm.Open(sqlite.Open(config.Get().Database.Url), &gorm.Config{})
	if err != nil {
		return err
	}

	con.AutoMigrate(
		&model.User{},
	)

	connection = con

	return nil
}

func GetConnectionWithContext(ctx context.Context) *gorm.DB {
	return connection.WithContext(ctx)
}

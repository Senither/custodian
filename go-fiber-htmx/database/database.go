package database

import (
	"context"

	"github.com/glebarez/sqlite"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/database/model"
	"gorm.io/gorm"
)

var connection *gorm.DB

// A DSN string for an in-memory SQLite database, this is used
// during testing to create a new database for each test.
const MemorySQLiteDSN = "file::memory:?cache=shared"

func InitiateDatabaseConnection(dsn string) error {
	con, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	con.AutoMigrate(
		&model.Session{},
		&model.User{},
		&model.Priority{},
		&model.Category{},
		&model.Task{},
	)

	connection = con

	return nil
}

func GetConnectionWithContext(ctx context.Context) *gorm.DB {
	db := connection.WithContext(ctx)

	if config.Get().Application.Debug {
		return db.Debug()
	}

	return db
}

func Disconnect() error {
	sqlDB, err := connection.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

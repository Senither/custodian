package model

import (
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupSessionTesting() (*gorm.DB, *DatabaseSessionStorage, func()) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Session{})

	store := NewDatabaseSessionStorage(db, session.Config{
		Expiration: time.Hour,
	})

	return db, store, func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}

func TestDatabaseSessionStorage_Get(t *testing.T) {
	_, store, close := setupSessionTesting()
	defer close()

	t.Run("it returns an empty session if not found", func(t *testing.T) {
		data, err := store.Get("non-existent-key")
		assert.NoError(t, err)
		assert.Empty(t, data)
	})

	t.Run("it returns the session data if found", func(t *testing.T) {
		store.Set("test-key", []byte("test-data"), time.Hour)
		data, err := store.Get("test-key")
		assert.NoError(t, err)
		assert.Equal(t, []byte("test-data"), data)
	})
}

func TestDatabaseSessionStorage_Set(t *testing.T) {
	_, store, close := setupSessionTesting()
	defer close()

	t.Run("it sets the session data", func(t *testing.T) {
		err := store.Set("test-key", []byte("test-data"), time.Hour)
		assert.NoError(t, err)

		data, _ := store.Get("test-key")
		assert.Equal(t, []byte("test-data"), data)
	})
}

func TestDatabaseSessionStorage_Delete(t *testing.T) {
	_, store, close := setupSessionTesting()
	defer close()

	t.Run("it deletes the session data", func(t *testing.T) {
		store.Set("test-key", []byte("test-data"), time.Hour)
		err := store.Delete("test-key")
		assert.NoError(t, err)

		data, _ := store.Get("test-key")
		assert.Empty(t, data)
	})
}

func TestDatabaseSessionStorage_Reset(t *testing.T) {
	_, store, close := setupSessionTesting()
	defer close()

	t.Run("it resets all session data", func(t *testing.T) {
		store.Set("test-key-1", []byte("test-data-1"), time.Hour)
		store.Set("test-key-2", []byte("test-data-2"), time.Hour)

		err := store.Reset()
		assert.NoError(t, err)

		data1, _ := store.Get("test-key-1")
		data2, _ := store.Get("test-key-2")
		assert.Empty(t, data1)
		assert.Empty(t, data2)
	})
}

func TestDatabaseSessionStorage_Close(t *testing.T) {
	_, store, close := setupSessionTesting()
	defer close()

	t.Run("it always returns nil", func(t *testing.T) {
		err := store.Close()
		assert.NoError(t, err)
	})
}

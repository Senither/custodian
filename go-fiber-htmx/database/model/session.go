package model

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Session struct {
	ID        string    `gorm:"primaryKey"`
	Data      []byte    `gorm:"type:text"`
	ExpiresAt time.Time `gorm:"index"`
}

type DatabaseSessionStorage struct {
	DB     *gorm.DB
	config session.Config
}

func NewDatabaseSessionStorage(db *gorm.DB, config session.Config) *DatabaseSessionStorage {
	return &DatabaseSessionStorage{
		DB: db,
		config: session.Config{
			Expiration: config.Expiration,
		},
	}
}

func (s *DatabaseSessionStorage) Get(key string) ([]byte, error) {
	var session Session

	result := s.DB.Where(Session{ID: key}).Attrs(Session{
		Data:      []byte{},
		ExpiresAt: time.Now().Add(s.config.Expiration),
	}).FirstOrInit(&session)

	return session.Data, result.Error
}

func (s *DatabaseSessionStorage) Set(key string, value []byte, exp time.Duration) error {
	session := Session{
		ID:        key,
		Data:      value,
		ExpiresAt: time.Now().Add(exp),
	}

	return s.DB.Save(&session).Error
}

func (s *DatabaseSessionStorage) Delete(key string) error {
	return s.DB.Delete(&Session{ID: key}).Error
}

func (s *DatabaseSessionStorage) Reset() error {
	return s.DB.Exec("DELETE FROM sessions").Error
}

func (s *DatabaseSessionStorage) Close() error {
	return nil
}

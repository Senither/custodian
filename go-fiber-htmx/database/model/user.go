package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string     `json:"name"`
	Email           string     `gorm:"index,unique;type:text collate nocase"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"-"` // Ignored when marshaling to JSON
}

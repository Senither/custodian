package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string     `json:"name"`
	Email           string     `json:"email" gorm:"index,unique;type:text collate nocase"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"-"` // Ignored when marshaling to JSON
	Tasks           []Task     `json:"tasks"`
	Priorities      []Priority `json:"priorities"`
	Categories      []Category `json:"categories"`
}

type Task struct {
	gorm.Model
	UserId     uint   `json:"user_id"`
	PriorityId uint   `json:"priority_id"`
	CategoryId uint   `json:"category_id"`
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	User       User
	Priority   Priority
	Category   Category
}

type Priority struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}

type Category struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}

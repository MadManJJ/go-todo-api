package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" validate:"required"`
	Description string `json:"description"`
	IsDone      bool   `gorm:"default:false" json:"is_done"`
	UserID      uint   `json:"user_id" validate:"required"` // Foreign key to the User
}

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Firstname string `gorm:"size:100" json:"firstname" validate:"required"`
	Lastname  string `gorm:"size:100" json:"lastname" validate:"required"`
	Password  string `gorm:"not null" json:"-" validate:"required,min=6"` // Hidden in JSON
	Todos     []Todo `json:"todos,omitempty"` // One-to-many relationship
}

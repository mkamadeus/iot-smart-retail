package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       *uuid.UUID `gorm:"type:uuid" json:"id"`
	CardID   string     `gorm:"unique"`
	Name     string     `json:"name" gorm:"not null"`
	Birthday *time.Time `json:"birthday" gorm:"not null"`
}

type CreateUserRequest struct {
	Name     string
	Birthday time.Time
}

func (r *CreateUserRequest) Build() *User {
	return &User{
		Name:     r.Name,
		Birthday: &r.Birthday,
	}
}

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	CardID    string    `gorm:"unique"`
	Name      string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Gender    string    `gorm:"not null"`
}

type CreateUserRequest struct {
	Name      string    `json:"name"`
	CardID    string    `json:"card_id"`
	Birthdate time.Time `json:"birthdate"`
	Gender    string    `json:"gender"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CardID    string    `json:"card_id"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	Gender    string    `json:"gender"`
}

func (r *CreateUserRequest) Build() *User {
	return &User{
		ID:        uuid.New(),
		CardID:    r.CardID,
		Name:      r.Name,
		Birthdate: r.Birthdate,
		Gender:    r.Gender,
	}
}

func (u *User) BuildResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		CardID:    u.CardID,
		Name:      u.Name,
		Birthdate: u.Birthdate,
		Gender:    u.Gender,
	}
}

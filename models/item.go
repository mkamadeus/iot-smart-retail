package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid" json:"id"`
	Name string
}

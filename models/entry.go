package models

import (
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	UserID    string `json:"userId"`
	User      User
	Timestamp *time.Time `json:"timestamp"`
	Type      string     `json:"type"`
}

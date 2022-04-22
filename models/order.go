package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TransactionID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Transaction   Transaction
	ItemID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Item          Item
	Count         uint
}

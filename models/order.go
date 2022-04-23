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

type ItemRequestWithCount struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	ItemID        uuid.UUID `json:"item_id"`
	Count         uint      `json:"count"`
}

func (r *ItemRequestWithCount) Build() *Order {
	return &Order{
		TransactionID: r.TransactionID,
		ItemID:        r.ItemID,
		Count:         r.Count,
	}
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid" json:"id"`
	UserID uuid.UUID
	User   User
	ItemID uuid.UUID
	Item   Item
}

type CreateTransactionRequest struct {
	UserID uuid.UUID
	ItemID uuid.UUID
}

// TODO: implement cart
type TransactionResponse struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Cart   string
}

func (r *CreateTransactionRequest) Build() *Transaction {
	return &Transaction{}
}

// TODO: implement cart
func (t *Transaction) BuildResponse() *TransactionResponse {
	return &TransactionResponse{
		ID:     t.ID,
		UserID: t.UserID,
		Cart:   "TBA",
	}
}

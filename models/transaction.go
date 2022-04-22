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
}

type CreateTransactionRequest struct {
	UserID uuid.UUID
	ItemID uuid.UUID
}

type TransactionResponse struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Order  Order
}

func (r *CreateTransactionRequest) Build() *Transaction {
	return &Transaction{}
}

func (t *Transaction) BuildResponse() *TransactionResponse {
	return &TransactionResponse{
		ID:     t.ID,
		UserID: t.UserID,
	}
}

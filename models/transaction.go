package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	UserID uuid.UUID `gorm:"type:uuid"`
	User   User
}

type TransactionDAO struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Items  []ItemWithCountDAO
}

type CreateTransactionRequest struct {
	UserID  uuid.UUID
	ItemIDs []uuid.UUID
}

type TransactionResponse struct {
	ID     uuid.UUID                `json:"id"`
	UserID uuid.UUID                `json:"user_id"`
	Items  []*ItemResponseWithCount `json:"items"`
}

func (r *CreateTransactionRequest) Build() *Transaction {
	return &Transaction{
		ID:     uuid.New(),
		UserID: r.UserID,
	}
}

func (t *TransactionDAO) BuildResponse() *TransactionResponse {

	itemResponses := make([]*ItemResponseWithCount, len(t.Items))
	for i, item := range t.Items {
		itemResponses[i] = item.BuildResponse()
	}

	return &TransactionResponse{
		ID:     t.ID,
		UserID: t.UserID,
		Items:  itemResponses,
	}
}

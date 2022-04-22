package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid"`
	Name string
}

type CreateItemRequest struct {
	Name string `json:"name"`
}

func (r *CreateItemRequest) Build() *Item {
	return &Item{
		ID:   uuid.New(),
		Name: r.Name,
	}
}

type ItemResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (i *Item) BuildResponse() *ItemResponse {
	return &ItemResponse{
		ID:   i.ID,
		Name: i.Name,
	}
}

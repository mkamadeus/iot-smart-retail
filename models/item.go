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

type ItemDAO struct {
	ID   uuid.UUID
	Name string
}

type ItemWithCountDAO struct {
	ID    uuid.UUID
	Name  string
	Count uint
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

type ItemResponseWithCount struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Count uint      `json:"count"`
}

func (i *ItemDAO) BuildResponse() *ItemResponse {
	return &ItemResponse{
		ID:   i.ID,
		Name: i.Name,
	}
}

func (i *ItemWithCountDAO) BuildResponse() *ItemResponseWithCount {
	return &ItemResponseWithCount{
		ID:    i.ID,
		Name:  i.Name,
		Count: i.Count,
	}
}

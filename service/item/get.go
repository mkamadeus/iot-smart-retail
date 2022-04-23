package item

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) Get(id uuid.UUID) (*models.ItemDAO, error) {
	var item *models.Item
	result := s.Database.First(item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ItemDAO{
		ID:   item.ID,
		Name: item.Name,
	}, nil
}

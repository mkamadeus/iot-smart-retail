package item

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) Get(id uuid.UUID) (*models.Item, error) {
	var item *models.Item
	result := s.Database.First(item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

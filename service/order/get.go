package order

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) Get(id uuid.UUID) (*models.Order, error) {
	var item *models.Order
	result := s.Database.First(item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

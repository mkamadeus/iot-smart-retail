package item

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) Create(item *models.Item) (*models.Item, error) {
	result := s.Database.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

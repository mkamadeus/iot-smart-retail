package item

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) GetAll() ([]*models.Item, error) {
	items := make([]*models.Item, 0)
	result := s.Database.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

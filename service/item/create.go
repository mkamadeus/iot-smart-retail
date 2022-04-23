package item

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) Create(item *models.Item) (*models.ItemDAO, error) {
	result := s.Database.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ItemDAO{
		ID:   item.ID,
		Name: item.Name,
	}, nil
}

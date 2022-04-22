package order

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) Create(user *models.Order) (*models.Order, error) {
	result := s.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

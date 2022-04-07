package transaction

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) Create(user *models.User) (*models.User, error) {
	result := s.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

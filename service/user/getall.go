package user

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) GetAll() ([]*models.User, error) {
	users := make([]*models.User, 0)
	result := s.Database.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

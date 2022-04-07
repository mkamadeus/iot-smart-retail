package user

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) Get(id *uuid.UUID) (*models.User, error) {
	var user *models.User
	result := s.Database.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

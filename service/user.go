package service

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
	"gorm.io/gorm"
)

type UserService struct {
	Database *gorm.DB
}

func NewUserService(Database *gorm.DB) *UserService {
	return &UserService{
		Database: Database,
	}
}

func (s *UserService) GetAll() ([]*models.User, error) {
	users := make([]*models.User, 0)
	result := s.Database.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserService) Get(id *uuid.UUID) (*models.User, error) {
	var user *models.User
	result := s.Database.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *UserService) Create(user *models.User) (*models.User, error) {
	result := s.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

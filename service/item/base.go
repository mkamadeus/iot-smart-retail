package item

import "gorm.io/gorm"

type Service struct {
	Database *gorm.DB
}

func New(database *gorm.DB) *Service {
	return &Service{
		Database: database,
	}
}

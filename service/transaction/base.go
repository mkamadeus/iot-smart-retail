package transaction

import "gorm.io/gorm"

type Service struct {
	Database *gorm.DB
}

func New(Database *gorm.DB) *Service {
	return &Service{
		Database: Database,
	}
}

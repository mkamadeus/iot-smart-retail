package transaction

import (
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) GetAll() ([]*models.Transaction, error) {
	txns := make([]*models.Transaction, 0)
	result := s.Database.Find(&txns)
	if result.Error != nil {
		return nil, result.Error
	}
	return txns, nil
}

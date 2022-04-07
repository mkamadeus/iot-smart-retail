package transaction

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) Get(id *uuid.UUID) (*models.Transaction, error) {
	var txn *models.Transaction
	result := s.Database.First(txn, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return txn, nil
}

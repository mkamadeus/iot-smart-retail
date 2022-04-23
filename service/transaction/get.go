package transaction

import (
	_ "embed"

	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

//go:embed get.sql
var getQuery string

func (s *Service) Get(id *uuid.UUID) (*models.TransactionDAO, error) {
	var txn *models.TransactionDAO
	// result := s.Database.First(txn, id)

	// result := s.Database.Raw(getQuery, )
	result := s.Database.Raw(getQuery).Scan(&txn)
	if result.Error != nil {
		return nil, result.Error
	}

	return txn, nil
}

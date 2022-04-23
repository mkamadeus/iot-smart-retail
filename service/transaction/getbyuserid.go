package transaction

import (
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (s *Service) GetByUserID(userId *uuid.UUID) ([]*models.TransactionDAO, error) {
	// var txns []*models.BaseTransaction
	// result := s.Database.Where(map[string]interface{}{"user_id": userId}).Find(&txns)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return txns, nil
	return make([]*models.TransactionDAO, 0), nil
}

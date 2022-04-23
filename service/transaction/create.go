package transaction

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) Create(txn *models.Transaction) (*models.TransactionDAO, error) {
	result := s.Database.Create(txn)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.TransactionDAO{
		ID:     txn.ID,
		UserID: txn.UserID,
	}, nil
}

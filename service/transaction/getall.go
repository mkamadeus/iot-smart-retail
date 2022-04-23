package transaction

import (
	_ "embed"
	"fmt"

	"github.com/mkamadeus/iot-smart-retail/models"
)

//go:embed getall.sql
var getallQuery string

func (s *Service) GetAll() ([]*models.TransactionDAO, error) {
	txns := make([]*models.TransactionDAO, 0)
	// result := s.Database.Find(&txns)

	val := map[string]interface{}{}
	result := s.Database.Raw(getallQuery).Scan(&val)
	fmt.Println(result.RowsAffected)
	fmt.Println(val)

	if result.Error != nil {
		return nil, result.Error
	}
	return txns, nil
}

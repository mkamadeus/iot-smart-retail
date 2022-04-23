package transaction

import (
	_ "embed"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

//go:embed getbyuserid.sql
var getbyuseridQuery string

func (s *Service) GetByUserID(userID *uuid.UUID) ([]*models.TransactionDAO, error) {
	txns := make([]*models.TransactionDAO, 0)

	val := []map[string]interface{}{}
	result := s.Database.Raw(getbyuseridQuery, userID).Scan(&val)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, v := range val {
		id, err := uuid.Parse(v["id"].(string))
		if err != nil {
			return nil, err
		}
		userID, err := uuid.Parse(v["user_id"].(string))
		if err != nil {
			return nil, err
		}
		items := make([]models.ItemWithCountDAO, 0)
		err = json.Unmarshal([]byte(v["items"].(string)), &items)
		if err != nil {
			return nil, err
		}

		txns = append(txns, &models.TransactionDAO{
			ID:     id,
			UserID: userID,
			Items:  items,
		})
	}

	return txns, nil
}

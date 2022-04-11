package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type CheckInRequest struct {
	CardID uuid.UUID `json:"card_id"`
}

func (req *CheckInRequest) UnmarshalJSON(b []byte) error {
	var v []interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	cardID, err := uuid.Parse(v[0].(string))
	if err != nil {
		return err
	}
	req.CardID = cardID
	return nil
}

type CheckOutRequest struct {
	CardID uuid.UUID `json:"card_id"`
}

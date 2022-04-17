package models

type CheckInRequest struct {
	CardID string `json:"card_id"`
}

type CheckOutRequest struct {
	CardID string `json:"card_id"`
}

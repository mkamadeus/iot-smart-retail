package transaction

import "github.com/mkamadeus/iot-smart-retail/service/transaction"

type Handler struct {
	Service *transaction.Service
}

func New(service *transaction.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

package transaction

import (
	"github.com/mkamadeus/iot-smart-retail/service/item"
	"github.com/mkamadeus/iot-smart-retail/service/order"
	"github.com/mkamadeus/iot-smart-retail/service/transaction"
)

type Handler struct {
	TransactionService *transaction.Service
	ItemService        *item.Service
	OrderService       *order.Service
}

func New(txn *transaction.Service, item *item.Service, order *order.Service) *Handler {
	return &Handler{
		TransactionService: txn,
		ItemService:        item,
		OrderService:       order,
	}
}

package api

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/handler/api/item"
	"github.com/mkamadeus/iot-smart-retail/handler/api/transaction"
	"github.com/mkamadeus/iot-smart-retail/handler/api/user"
	"github.com/mkamadeus/iot-smart-retail/handler/pubsub/entry"
)

type Handler struct {
	User        *user.Handler
	Transaction *transaction.Handler
	Item        *item.Handler
	Entry       *entry.Handler
}

var HandlerSet = wire.NewSet(New, user.New, item.New, transaction.New)

func New(uh *user.Handler, th *transaction.Handler, ih *item.Handler) *Handler {
	return &Handler{
		User:        uh,
		Transaction: th,
		Item:        ih,
	}
}

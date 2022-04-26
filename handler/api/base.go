package api

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/handler/api/item"
	"github.com/mkamadeus/iot-smart-retail/handler/api/sse"
	"github.com/mkamadeus/iot-smart-retail/handler/api/transaction"
	"github.com/mkamadeus/iot-smart-retail/handler/api/user"
)

type Handler struct {
	User        *user.Handler
	Transaction *transaction.Handler
	Item        *item.Handler
	SSE         *sse.Handler
}

var HandlerSet = wire.NewSet(New, user.New, item.New, transaction.New, sse.New)

func New(uh *user.Handler, th *transaction.Handler, ih *item.Handler, sseh *sse.Handler) *Handler {
	return &Handler{
		User:        uh,
		Transaction: th,
		Item:        ih,
		SSE:         sseh,
	}
}

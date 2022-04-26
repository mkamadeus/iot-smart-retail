package pubsub

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/handler/pubsub/entry"
	"github.com/mkamadeus/iot-smart-retail/handler/pubsub/storefront"
)

type Handler struct {
	Entry      *entry.Handler
	Storefront *storefront.Handler
}

var HandlerSet = wire.NewSet(New, entry.New, storefront.New)

func New(eh *entry.Handler, sh *storefront.Handler) *Handler {
	return &Handler{
		Entry:      eh,
		Storefront: sh,
	}
}

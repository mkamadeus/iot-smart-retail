package pubsub

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/handler/pubsub/entry"
)

type Handler struct {
	Entry *entry.Handler
}

var HandlerSet = wire.NewSet(New, entry.New)

func New(eh *entry.Handler) *Handler {
	return &Handler{
		Entry: eh,
	}
}

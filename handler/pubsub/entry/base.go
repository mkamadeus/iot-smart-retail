package entry

import "github.com/mkamadeus/iot-smart-retail/service/entry"

type Handler struct {
	Service *entry.Service
}

func New(service *entry.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

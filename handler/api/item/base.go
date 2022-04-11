package item

import "github.com/mkamadeus/iot-smart-retail/service/item"

type Handler struct {
	Service *item.Service
}

func New(service *item.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

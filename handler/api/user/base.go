package user

import (
	"github.com/mkamadeus/iot-smart-retail/service/user"
)

type Handler struct {
	Service *user.Service
}

func New(service *user.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

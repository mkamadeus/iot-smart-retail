package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/service/user"
)

type handler interface {
	Get(ctx *fiber.Ctx) error
}

type Handler struct {
	Service *user.Service
}

func New(service *user.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

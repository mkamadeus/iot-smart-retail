package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) Create(ctx *fiber.Ctx) error {
	var request models.CreateUserRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	user := request.Build()
	created, err := h.Service.Create(user)
	if err != nil {
		return err
	}
	return ctx.JSON(created.BuildResponse())
}

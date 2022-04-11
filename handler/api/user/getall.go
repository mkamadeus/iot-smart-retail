package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) GetAll(ctx *fiber.Ctx) error {
	users, err := h.Service.GetAll()
	if err != nil {
		return err
	}
	response := make([]*models.UserResponse, len(users))
	for i, user := range users {
		response[i] = user.BuildResponse()
	}

	return ctx.JSON(response)
}

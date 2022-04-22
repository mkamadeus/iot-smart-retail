package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) Create(ctx *fiber.Ctx) error {
	var request models.CreateItemRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	item := request.Build()
	created, err := h.Service.Create(item)
	if err != nil {
		return err
	}
	return ctx.JSON(created.BuildResponse())
}

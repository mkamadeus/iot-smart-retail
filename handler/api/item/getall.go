package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) GetAll(ctx *fiber.Ctx) error {
	items, err := h.Service.GetAll()
	if err != nil {
		return err
	}
	response := make([]*models.ItemResponse, len(items))
	for i, item := range items {
		response[i] = item.BuildResponse()
	}

	return ctx.JSON(response)
}

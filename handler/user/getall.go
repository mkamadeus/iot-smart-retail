package user

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetAll(ctx *fiber.Ctx) error {
	users, err := h.Service.GetAll()
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

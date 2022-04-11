package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	converted, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	txn, err := h.Service.Get(&converted)
	if err != nil {
		return err
	}
	return ctx.JSON(txn.BuildResponse())
}

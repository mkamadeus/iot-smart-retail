package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	converted, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	txn, err := h.TransactionService.Get(&converted)
	if err != nil {
		return err
	}

	response := &models.TransactionResponse{
		ID:     txn.ID,
		UserID: txn.UserID,
	}

	// get items

	return ctx.JSON(response)
}

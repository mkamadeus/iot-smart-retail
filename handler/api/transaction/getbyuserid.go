package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) GetByUserID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	converted, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	txns, err := h.TransactionService.GetByUserID(&converted)
	if err != nil {
		return err
	}

	response := make([]*models.TransactionResponse, len(txns))
	for i, txn := range txns {
		response[i] = txn.BuildResponse()
	}

	// get items

	return ctx.JSON(response)
}

package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) GetAll(ctx *fiber.Ctx) error {
	txns, err := h.TransactionService.GetAll()
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

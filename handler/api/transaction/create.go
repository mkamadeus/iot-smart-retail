package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) Create(ctx *fiber.Ctx) error {
	var request models.CreateTransactionRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	// build txns
	txn := request.Build()
	created, err := h.TransactionService.Create(txn)

	//
	if err != nil {
		return err
	}
	return ctx.JSON(created.BuildResponse())
}

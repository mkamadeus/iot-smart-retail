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

	// check item validity
	items := make([]*models.ItemResponse, len(request.ItemIDs))
	for _, itemID := range request.ItemIDs {
		item, err := h.ItemService.Get(itemID)
		if err != nil {
			return err
		}
		items = append(items, item.BuildResponse())
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

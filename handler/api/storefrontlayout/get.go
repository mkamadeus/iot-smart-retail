package storefrontlayout

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetLayout(ctx *fiber.Ctx) error {
	result, err := h.Service.GetStoreLayout()
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}

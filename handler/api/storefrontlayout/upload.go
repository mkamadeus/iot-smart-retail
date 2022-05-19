package storefrontlayout

import "github.com/gofiber/fiber/v2"

func (h *Handler) UploadLayout(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("upload-layout")
	if err != nil {
		return err
	}

	if err := h.Service.UploadStoreLayout(fileHeader); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"message": "Upload is successful!"})
}

package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberApp() *fiber.App {
	// init app
	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("smart retail")
	})

	return app
}

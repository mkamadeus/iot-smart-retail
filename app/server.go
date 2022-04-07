package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mkamadeus/iot-smart-retail/handler"
)

func NewFiberServer(handlers *handler.Handler) *fiber.App {
	// init app
	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("smart retail")
	})

	// users
	app.Post("/users", handlers.User.Create)
	app.Get("/users", handlers.User.GetAll)
	app.Get("/users/:id", handlers.User.Get)

	return app
}

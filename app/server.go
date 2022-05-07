package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mkamadeus/iot-smart-retail/handler/api"
)

func NewFiberServer(handlers *api.Handler) *fiber.App {
	// init app
	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Hello world")
		return ctx.JSON("smart retail")
	})

	// users
	app.Post("/users", handlers.User.Create)
	app.Get("/users", handlers.User.GetAll)
	app.Get("/users/:id", handlers.User.Get)
	app.Get("/users/:id/transactions", handlers.Transaction.GetByUserID)

	// items
	app.Post("/items", handlers.Item.Create)
	app.Get("/items", handlers.Item.GetAll)

	// txns
	app.Get("/transactions/", handlers.Transaction.GetAll)

	app.Get("/sse", handlers.SSE.SendToClients)

	app.Post("/layout", handlers.Layout.UploadLayout)
	app.Get("/layout", handlers.Layout.GetLayout)

	return app
}

package app

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/config"
	"gorm.io/gorm"
)

type App struct {
	Server         *fiber.App
	DatabaseClient *gorm.DB
	MQTTClient     *mqtt.Client
	CacheClient    *redis.Client
	Config         *config.Config
	SSERoutine     *SSERoutine
}

func NewApp(server *fiber.App, db *gorm.DB, mc *mqtt.Client, cache *redis.Client, conf *config.Config, sseRoutine *SSERoutine) *App {
	return &App{
		Server:         server,
		DatabaseClient: db,
		MQTTClient:     mc,
		CacheClient:    cache,
		Config:         conf,
		SSERoutine:     sseRoutine,
	}
}

func (app *App) Listen() {
	done := make(chan interface{})
	defer close(done)
	app.SSERoutine.Run(done)
	err := app.Server.Listen(fmt.Sprintf(":%d", app.Config.Server.Port))
	if err != nil {
		fmt.Println("Can't run the server")
	}
}

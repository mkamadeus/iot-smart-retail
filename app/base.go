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
}

func NewApp(server *fiber.App, db *gorm.DB, mc *mqtt.Client, cache *redis.Client, conf *config.Config) *App {
	return &App{
		Server:         server,
		DatabaseClient: db,
		MQTTClient:     mc,
		CacheClient:    cache,
		Config:         conf,
	}
}

func (app *App) Listen() {
	app.Server.Listen(fmt.Sprintf(":%d", app.Config.Server.Port))
}

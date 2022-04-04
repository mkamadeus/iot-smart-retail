//go:build wireinject
// +build wireinject

package app

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type App struct {
	Server         *fiber.App
	DatabaseClient *gorm.DB
	MQTTClient     *mqtt.Client
	CacheClient    *redis.Client
}

func NewApp(server *fiber.App, db *gorm.DB, mc *mqtt.Client, cache *redis.Client) *App {
	return &App{
		Server:         server,
		DatabaseClient: db,
		MQTTClient:     mc,
		CacheClient:    cache,
	}
}

func InitializeApp() (*App, error) {
	wire.Build(NewApp, NewFiberApp, NewCache, NewDB, NewMQTTClient, NewConfig)
	return &App{}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Injectors from base.go:

func InitializeApp() (*App, error) {
	app := NewFiberApp()
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := NewDB(config)
	if err != nil {
		return nil, err
	}
	client, err := NewMQTTClient(config)
	if err != nil {
		return nil, err
	}
	redisClient := NewCache(config)
	appApp := NewApp(app, db, client, redisClient)
	return appApp, nil
}

// base.go:

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

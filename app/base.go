package app

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	Server         *fiber.App
	DatabaseClient *gorm.DB
	MQTTClient     *mqtt.Client
	CacheClient    *redis.Client
}

func InitializeApp() {

}

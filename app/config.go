package app

import (
	"os"
	"strconv"
)

type Config struct {
	Port     uint
	MQTT     mqttConfig
	Database databaseConfig
	Cache    cacheConfig
}

type mqttConfig struct {
	ServerURL     string
	ClientID      string
	TapTopic      string
	RegisterTopic string
	DisplayTopic  string
}

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type cacheConfig struct {
	Host     string
	Port     string
	Password string
}

func GetConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	appConfig := &Config{
		Port: uint(port),
		MQTT: mqttConfig{
			ServerURL:     os.Getenv("MQTT_SERVER_URL"),
			TapTopic:      os.Getenv("MQTT_TAP_TOPIC"),
			RegisterTopic: os.Getenv("MQTT_REGISTER_TOPIC"),
			DisplayTopic:  os.Getenv("MQTT_DISPLAY_TOPIC"),
		},
		Database: databaseConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DATABASE"),
		},
		Cache: cacheConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	return appConfig, nil
}

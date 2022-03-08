package app

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Port uint
	MQTT mqttConfig
}

type mqttConfig struct {
	ServerURL     string
	TapTopic      string
	RegisterTopic string
	DisplayTopic  string
}

var AppConfig *Config
var lock = &sync.Mutex{}

func GetConfig() (*Config, error) {
	if AppConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			return nil, err
		}
		AppConfig = &Config{
			Port: uint(port),
			MQTT: mqttConfig{
				ServerURL:     os.Getenv("MQTT_SERVER_URL"),
				TapTopic:      os.Getenv("MQTT_TAP_TOPIC"),
				RegisterTopic: os.Getenv("MQTT_REGISTER_TOPIC"),
				DisplayTopic:  os.Getenv("MQTT_DISPLAY_TOPIC"),
			},
		}
	}

	return AppConfig, nil
}

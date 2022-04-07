package config

import "os"

type MQTTConfig struct {
	ServerURL     string
	ClientID      string
	TapTopic      string
	RegisterTopic string
	DisplayTopic  string
}

func NewMQTTClientConfig() *MQTTConfig {
	return &MQTTConfig{
		ServerURL:     os.Getenv("MQTT_SERVER_URL"),
		TapTopic:      os.Getenv("MQTT_TAP_TOPIC"),
		RegisterTopic: os.Getenv("MQTT_REGISTER_TOPIC"),
		DisplayTopic:  os.Getenv("MQTT_DISPLAY_TOPIC"),
	}
}

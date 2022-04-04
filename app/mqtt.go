package app

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTClient(config *Config) (*mqtt.Client, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.MQTT.ServerURL)
	opts.SetClientID(config.MQTT.ClientID)
	opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)
	return &client, nil

}

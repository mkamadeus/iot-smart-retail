package external

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mkamadeus/iot-smart-retail/config"
)

func NewMQTTClient(c *config.Config) (*mqtt.Client, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.MQTT.ServerURL)
	opts.SetClientID(c.MQTT.ClientID)
	opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)
	return &client, nil

}

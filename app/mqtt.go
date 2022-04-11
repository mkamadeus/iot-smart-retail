package app

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mkamadeus/iot-smart-retail/config"
	"github.com/mkamadeus/iot-smart-retail/handler/pubsub"
)

func NewMQTTClient(c *config.Config, handlers *pubsub.Handler) (*mqtt.Client, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.MQTT.ServerURL)
	opts.SetClientID(c.MQTT.ClientID)
	opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	client.Subscribe(c.MQTT.TapTopic, 1, handlers.Entry.CheckIn)

	return &client, nil

}

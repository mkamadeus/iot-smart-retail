package storefront

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (h *Handler) ReceiveMessage(_ mqtt.Client, m mqtt.Message) {
	msg := string(m.Payload()[:])
	h.Service.Notifier <- msg
}

package storefront

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mkamadeus/iot-smart-retail/utils"
)

func (h *Handler) ReceiveMessage(_ mqtt.Client, m mqtt.Message) {
	msg := utils.B2S(m.Payload())
	h.Service.Notifier <- msg
}

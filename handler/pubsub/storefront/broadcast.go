package storefront

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) ReceiveMessage(_ mqtt.Client, m mqtt.Message) {
	var msg models.MQTTMessage
	err := json.Unmarshal(m.Payload(), &msg)
	if err != nil {
		return
	}
	fmt.Println(msg)
	h.Service.Notifier <- msg.Message
}

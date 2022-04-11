package entry

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) CheckIn(c mqtt.Client, m mqtt.Message) {
	var req models.CheckInRequest
	err := json.Unmarshal(m.Payload(), &req)

	fmt.Println(req)
	if err != nil {
		return
	}

	err = h.Service.CheckIn(&req.CardID)
	if err != nil {
		return
	}
}

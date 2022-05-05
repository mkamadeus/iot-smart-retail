package entry

import (
	"encoding/json"
	"fmt"
	"github.com/mkamadeus/iot-smart-retail/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mkamadeus/iot-smart-retail/models"
)

func (h *Handler) CheckOut(c mqtt.Client, m mqtt.Message) {
	var req models.CheckOutRequest
	fmt.Println(utils.B2S(m.Payload()))
	err := json.Unmarshal(m.Payload(), &req)

	fmt.Println(req)
	if err != nil {
		return
	}

	err = h.Service.CheckOut(req.CardID)
	if err != nil {
		return
	}
}

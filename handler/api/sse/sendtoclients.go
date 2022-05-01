package sse

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mkamadeus/iot-smart-retail/utils"
	"github.com/valyala/fasthttp"
)

func (h *Handler) SendToClients(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", "chunked")

	sseChannel := make(chan string)
	h.Service.Clients = append(h.Service.Clients, sseChannel)

	ctx.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			data := <-sseChannel
			parsedMessage := utils.SplitByCharacter(data, '-')
			fmt.Fprintf(w, "event: %s\ndata: %s\n\n", parsedMessage[0], parsedMessage[1])

			w.Flush()
		}
	}))
	return nil
}

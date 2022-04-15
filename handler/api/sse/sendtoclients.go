package sse

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func (h *Handler) SendToClients(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", "chunked")

	sseChannel := make(chan string)
	h.Service.Clients = append(h.Service.Clients, sseChannel)

	defer func() {
		close(sseChannel)
		sseChannel = nil
	}()

	ctx.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			data := <-sseChannel
			fmt.Printf("data: %v \n\n", data)
			fmt.Fprintf(w, "data: %v \n\n", data)

			w.Flush()
		}
	}))

	return nil
}

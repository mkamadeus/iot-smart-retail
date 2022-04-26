package sse

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"strings"
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
			parsedMessage := splitByDash(data)
			fmt.Fprintf(w, "event: %s\ndata: %s\n\n", parsedMessage[0], parsedMessage[1])

			w.Flush()
		}
	}))
	return nil
}

func splitByDash(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		if r == '-' {
			return true
		}
		return false
	})
}

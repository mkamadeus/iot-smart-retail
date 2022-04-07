package handler

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/handler/entry"
	"github.com/mkamadeus/iot-smart-retail/handler/item"
	"github.com/mkamadeus/iot-smart-retail/handler/transaction"
	"github.com/mkamadeus/iot-smart-retail/handler/user"
)

type Handler struct {
	User *user.Handler
	// Transaction transaction.Handler
	// Item        item.Handler
	// Entry       entry.Handler
}

var HandlerSet = wire.NewSet(New, user.New, item.New, transaction.New, entry.New)

// func New(userService user.Handler, transactionService transaction.Handler, itemService item.Handler, entryService entry.Handler) *Handler {
func New(userService *user.Handler) *Handler {
	return &Handler{
		User: userService,
		// Transaction: transactionService,
		// Item:        itemService,
		// Entry:       entryService,
	}
}

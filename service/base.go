package service

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/service/entry"
	"github.com/mkamadeus/iot-smart-retail/service/item"
	"github.com/mkamadeus/iot-smart-retail/service/transaction"
	"github.com/mkamadeus/iot-smart-retail/service/user"
)

var ServiceSet = wire.NewSet(user.New, item.New, transaction.New, entry.New)

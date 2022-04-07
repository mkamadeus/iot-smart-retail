//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/config"
	"github.com/mkamadeus/iot-smart-retail/external"
	"github.com/mkamadeus/iot-smart-retail/handler"
	"github.com/mkamadeus/iot-smart-retail/service"
)

func InitializeApp() (*App, error) {
	wire.Build(NewApp, NewFiberServer, external.ExternalSet, config.ConfigSet, handler.HandlerSet, service.ServiceSet)
	return &App{}, nil
}

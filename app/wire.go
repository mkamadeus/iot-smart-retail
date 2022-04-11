//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/mkamadeus/iot-smart-retail/config"
	"github.com/mkamadeus/iot-smart-retail/handler/api"
	"github.com/mkamadeus/iot-smart-retail/handler/pubsub"
	"github.com/mkamadeus/iot-smart-retail/repository"
	"github.com/mkamadeus/iot-smart-retail/service"
)

func InitializeApp() (*App, error) {
	wire.Build(
		NewApp,
		NewFiberServer,
		NewMQTTClient,
		repository.RepositorySet,
		config.ConfigSet,
		api.HandlerSet,
		pubsub.HandlerSet,
		service.ServiceSet,
	)
	return &App{}, nil
}

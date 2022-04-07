package config

import (
	"github.com/google/wire"
)

type Config struct {
	Server   *ServerConfig
	MQTT     *MQTTConfig
	Database *DatabaseConfig
	Cache    *CacheConfig
}

var ConfigSet = wire.NewSet(New, NewCacheConfig, NewDatabaseConfig, NewMQTTClientConfig, NewServerConfig)

func New(server *ServerConfig, cache *CacheConfig, db *DatabaseConfig, mqtt *MQTTConfig) (*Config, error) {
	appConfig := &Config{
		Server:   server,
		Cache:    cache,
		Database: db,
		MQTT:     mqtt,
	}

	return appConfig, nil
}

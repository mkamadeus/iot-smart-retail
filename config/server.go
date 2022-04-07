package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Port uint
}

func NewServerConfig() (*ServerConfig, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	return &ServerConfig{
		Port: uint(port),
	}, nil
}

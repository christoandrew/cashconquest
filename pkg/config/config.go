package config

import (
	"os"

	database "github.com/christo-andrew/maybe-go/pkg/database"
)

type ServerConfig struct {
	Host           string
	Port           string
	EnvPath        string
	DatabaseConfig database.DatabaseConfig
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Host:           os.Getenv("HOST"),
		Port:           os.Getenv("PORT"),
		DatabaseConfig: database.DefaultDatabaseConfig(),
		EnvPath:        ".env",
	}
}

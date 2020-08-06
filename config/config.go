package config

import "time"

type ServerConfig struct {
	Address        string
	HandlerTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
}

type DatabaseConfig struct {
	URI    string
	Driver string
}

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func New() *AppConfig {
	return &AppConfig{
		Server: ServerConfig{
			Address:        getEnv("ADDRESS", "localhost:4000"),
			HandlerTimeout: getEnvAsDuration("HANDLER_TIMEOUT", 10),
			ReadTimeout:    getEnvAsDuration("READ_TIMEOUT", 5),
			WriteTimeout:   getEnvAsDuration("WRITE_TIMEOUT", 10),
			IdleTimeout:    getEnvAsDuration("IDLE_TIMEOUT", 20),
		},
		Database: DatabaseConfig{
			URI:    getEnv("DB_URI", ""),
			Driver: getEnv("DB_DRIVER", "postgres"),
		},
	}
}

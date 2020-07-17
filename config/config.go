package config

type ServerConfig struct {
	Address string
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
			Address: getEnv("ADDRESS", "localhost:4000"),
		},
		Database: DatabaseConfig{
			URI:    getEnv("DB_URI", ""),
			Driver: getEnv("DB_DRIVER", "postgres"),
		},
	}
}

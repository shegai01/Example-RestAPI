package api

import "github.com/shegai01/Example-RestAPI/storage"

// general instance for API server of REST application
type Config struct {
	// port
	BindAddr string `toml:"bind_addr"`
	//logger level
	LoggerLevel string `toml:"logger_level"`
	//storage config
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}

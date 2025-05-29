package api

import "github.com/shegai01/Example-RestAPI/storage"

type Config struct {
	BindAddr string          `toml:"bind_addr"`
	LogLevel string          `toml:"log_level"`
	Storage  *storage.Config `toml:"storage"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Storage:  storage.NewConfig(),
	}

}

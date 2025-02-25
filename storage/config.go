package storage

// add string for configure database
type Config struct {
	DataBaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}

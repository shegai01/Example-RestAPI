package api

import (
	"github.com/sirupsen/logrus"
)

// base API server instance description
type API struct {
	// unexported feild !!
	config *Config
	logger *logrus.Logger
}

// API constructor build API instance
func NewAPI(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}

}

// start http sercer/confirue, router, database
func (api *API) Start() error {
	return nil
}

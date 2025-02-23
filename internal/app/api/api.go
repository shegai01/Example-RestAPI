package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// base API server instance description
type API struct {
	// unexported feild !!
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// API constructor build API instance
func NewAPI(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}

// start http sercer/confirue, router, database
func (api *API) Start() error {
	api.configureLoggerField()
	//configure router
	api.configureRouterFeild()
	// start http-server
	api.logger.Info("starting api server at port: ", api.config.BindAddr)
	return http.ListenAndServe(api.config.BindAddr, api.router)
}

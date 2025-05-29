package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shegai01/Example-RestAPI/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}
func (api *API) Start() error {
	if err := api.configureLogger(); err != nil {
		return err
	}
	api.logger.Info("api startin at port: ", api.config.BindAddr)
	api.configRouter()
	if err := api.configStorage(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.BindAddr, api.router)
}

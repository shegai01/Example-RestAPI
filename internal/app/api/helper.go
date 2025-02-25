package api

import (
	"net/http"

	"github.com/shegai01/Example-RestAPI/storage"
	"github.com/sirupsen/logrus"
)

// try configure api
func (api *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(log_level)
	return nil
}

// try configure router
func (api *API) configureRouterFeild() {
	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world this is api"))
	})
}

// try configure storage
func (api *API) configureStorageFeild() error {
	storage := storage.New(api.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	api.storage = storage
	return nil
}

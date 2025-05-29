package api

import (
	"net/http"

	"github.com/shegai01/Example-RestAPI/storage"
	"github.com/sirupsen/logrus"
)

func (a *API) configureLogger() error {
	log, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log)
	return nil

}

func (a *API) configRouter() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hey say api"))
	})

}
func (a *API) configStorage() error {
	storage := storage.NewStorage(a.config.Storage)
	if err := storage.Open(); err != nil { //connection db else error
		return err
	}
	a.storage = storage
	return nil
}

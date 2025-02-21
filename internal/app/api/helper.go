package api

import "github.com/sirupsen/logrus"

// try configure api
func (api *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(log_level)
	return nil
}

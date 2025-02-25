package storage

import "database/sql"

// instance storage
type Storage struct {
	config *Config
	db     *sql.DB
}

// constuction storage
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DataBaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	return nil
}

// close connection method
func (storage *Storage) Close() {
	storage.db.Close()

}

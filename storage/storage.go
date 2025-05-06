package storage

type Storage struct {
	Config *StorageConfig
}

func NewStorage(config *StorageConfig) *ConfigStorage {
	return &ConfigStorage{
		DataBaseURI: ,
	}
}

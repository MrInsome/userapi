package storage

import (
	"encoding/json"
	"net/http"
	"os"
	"refactoring/internal/config"
	"refactoring/internal/data"
)

type FileStorage struct {
	*config.Configs
}

func NewFileStorage(c *config.Configs) *FileStorage {
	return &FileStorage{c}
}

func (fs *FileStorage) ReadStore() (data.UserStore, error) {
	var store data.UserStore
	StoreFS := http.Dir(fs.Directory)
	_, err := StoreFS.Open(fs.Name)
	if err == nil {
		file, err := os.ReadFile(fs.Name)
		if err != nil {
			return store, err
		}

		err = json.Unmarshal(file, &store)
		if err != nil {
			return store, err
		}

		return store, nil
	} else {
		return store, err
	}
}

func (fs *FileStorage) WriteStore(store data.UserStore) error {
	file, err := json.Marshal(store)
	if err != nil {
		return err
	}
	StoreFS := http.Dir(fs.Directory)
	_, err = StoreFS.Open(fs.Name)
	if err == nil {
		err = os.WriteFile(fs.Name, file, 0644)
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}

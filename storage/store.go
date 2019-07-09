package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const (
	// configName is the name of the config.
	configName = "config.json"

	// Permissions are the directory permissions for storage.
	Permissions = 0755
)

// Store represents the data store while hiding implementation behind the interface.
type Store interface {
	Load() error
	Read() (*Config, error)
	Write(config *Config) error
}

type store struct {
	path       string
	configName string
	writeMutex sync.Mutex
}

// Read will return the current configuration.
func (s *store) Read() (*Config, error) {
	body, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", s.path, configName))
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := json.Unmarshal(body, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (s *store) Write(config *Config) error {
	// Lock access to the file.
	s.writeMutex.Lock()

	// Unlock it when we're done writing.
	defer s.writeMutex.Unlock()

	// Marshal the data into json.
	body, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Write to the file, replacing the existing config with the new updated one.
	return ioutil.WriteFile(
		fmt.Sprintf("%s/%s", s.path, configName),
		body,
		Permissions,
	)
}

// Load will create the directory and config file if it doesn't
// exist, and will load a default config, if the config exists
// it will be set on the store.
func (s *store) Load() error {
	// if the config doesn't exist, create it.
	configExists, err := s.configExists()
	if err != nil {
		return err
	}

	if !configExists {
		c := &Config{
			Games: make([]Game, 0),
		}

		// Write a new config with default settings.
		return s.Write(c)
	}

	return nil
}

func (s *store) configExists() (bool, error) {
	_, err := os.Stat(fmt.Sprintf("%s/%s", s.path, configName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err

	}

	return true, nil
}

// NewStore returns a new store with all dependencies set up.
func NewStore(path string) Store {
	return &store{
		path: path,
	}
}

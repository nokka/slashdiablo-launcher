package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/therecipe/qt/core"
)

const (
	configName  = "config.json"
	permissions = 0755
)

// Store represents the data store available while hiding implementation
// details behind the interface.
type Store interface {
	Load() error
	Read() (*Config, error)
}

type store struct {
	path       string
	configName string
}

// Read will return the current configuration.
func (s *store) Read() (*Config, error) {
	body, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", s.path, s.configName))
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
	// Marshal the data into json.
	body, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Write to the file, replacing the existing config with the new updated one.
	return ioutil.WriteFile(
		fmt.Sprintf("%s/%s", s.path, s.configName),
		body,
		permissions,
	)
}

// Load will create the directory and config file if it doesn't
// exist, and will load a default config, if the config exists
// it will be set on the store.
func (s *store) Load() error {
	// Get the target specific data storage location.
	err := s.getLocation()
	if err != nil {
		return err
	}

	// Check if the data directory exists.
	pathExists, err := s.pathExists()
	if err != nil {
		return err
	}

	// If the path doesn't exist, create it.
	if !pathExists {
		err := os.Mkdir(s.path, permissions)
		if err != nil {
			return err
		}
	}

	// if the config doesn't exist, create it.
	configExists, err := s.configExists()
	if err != nil {
		return err
	}

	if !configExists {
		c := &Config{
			D2Instances: 1,
		}
		// Write a new config with default settings.
		return s.Write(c)
	}

	return nil
}

func (s *store) pathExists() (bool, error) {
	// Returns an error if the path does not exist.
	f, err := os.Stat(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err

	}

	return f.IsDir(), nil
}

func (s *store) configExists() (bool, error) {
	// Returns an error if the config does not exist.
	_, err := os.Stat(s.configName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err

	}

	return true, nil
}

func (s *store) getLocation() error {
	locations := core.QStandardPaths_StandardLocations(core.QStandardPaths__AppLocalDataLocation)
	if len(locations) == 0 {
		return errors.New("missing app data location")
	}

	// Grab the first available location.
	s.path = locations[0]

	return nil
}

// NewStore returns a new store with all dependencies set up.
func NewStore() Store {
	return &store{
		configName: configName,
	}
}

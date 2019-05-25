package config

import (
	"fmt"

	"github.com/nokka/slash-launcher/log"
	"github.com/nokka/slash-launcher/storage"
)

// Service is responsible for all things related to configuration.
type Service interface {
	Read() (*storage.Config, error)
	Update(fields map[string]interface{}) error
}

type service struct {
	store  storage.Store
	logger log.Logger
}

// Update will update the configuration with the given fields.
func (s *service) Read() (*storage.Config, error) {
	conf, err := s.store.Read()
	if err != nil {
		s.logger.Log("failed to read config", err)
		return nil, err
	}

	return conf, err
}

// Update will update the configuration with the given fields.
func (s *service) Update(fields map[string]interface{}) error {
	fmt.Println("UPDATING FIELDS")
	fmt.Println(fields)
	return nil
}

// NewService returns a service with all the dependencies.
func NewService(
	store storage.Store,
	logger log.Logger,
) Service {
	return &service{
		store:  store,
		logger: logger,
	}
}

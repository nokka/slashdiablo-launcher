package config

import (
	"github.com/nokka/slash-launcher/log"
	"github.com/nokka/slash-launcher/storage"
)

// Service is responsible for all things related to configuration.
type Service interface {
	Read() (*storage.Config, error)
	Update(request UpdateConfigRequest) error
}

type service struct {
	store  storage.Store
	logger log.Logger
}

// Read will read the configuration and return it.
func (s *service) Read() (*storage.Config, error) {
	conf, err := s.store.Read()
	if err != nil {
		s.logger.Log("failed to read config", err)
		return nil, err
	}

	return conf, err
}

// UpdateConfigRequest is the data available to update the config with.
type UpdateConfigRequest struct {
	D2Location  *string
	D2Instances *int
	D2Maphack   *bool
	HDLocation  *string
	HDInstances *int
	HDMaphack   *bool
}

// Update will update the configuration with the given fields.
func (s *service) Update(request UpdateConfigRequest) error {
	conf, err := s.store.Read()
	if err != nil {
		s.logger.Log("failed to read config", err)
		return err
	}

	if request.D2Location != nil {
		conf.D2Location = *request.D2Location
	}

	if request.D2Instances != nil {
		conf.D2Instances = *request.D2Instances
	}

	if request.D2Maphack != nil {
		conf.D2Maphack = *request.D2Maphack
	}

	if request.HDLocation != nil {
		conf.HDLocation = *request.HDLocation
	}

	if request.HDInstances != nil {
		conf.HDInstances = *request.HDInstances
	}

	if request.HDMaphack != nil {
		conf.HDMaphack = *request.HDMaphack
	}

	err = s.store.Write(conf)
	if err != nil {
		s.logger.Log("failed to write config", err)
		return err
	}

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

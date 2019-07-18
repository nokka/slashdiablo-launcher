package config

import (
	"github.com/nokka/slash-launcher/log"
	"github.com/nokka/slash-launcher/storage"
)

// Service is responsible for all things related to configuration.
type Service interface {
	Read() (*storage.Config, error)
	UpdateGame(request UpdateGameRequest) error
}

type service struct {
	store     storage.Store
	gameModel *GameModel
	logger    log.Logger
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

// UpdateGameRequest ...
type UpdateGameRequest struct {
	ID        int    `json:"id"`
	Location  string `json:"location"`
	Instances int    `json:"instances"`
	Maphack   bool   `json:"maphack"`
	HD        bool   `json:"hd"`
}

// UpdateGame will update the configuration with the given game.
func (s *service) UpdateGame(request UpdateGameRequest) error {
	conf, err := s.store.Read()
	if err != nil {
		s.logger.Log("failed to read config", err)
		return err
	}

	// Look for game to update, and mutate if found.
	for i := 0; i < len(conf.Games); i++ {
		if request.ID == conf.Games[i].ID {
			conf.Games[i].Location = request.Location
			conf.Games[i].Instances = request.Instances
			conf.Games[i].Maphack = request.Maphack
			conf.Games[i].HD = request.HD
		}
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
	gameModel *GameModel,
	logger log.Logger,
) Service {
	return &service{
		store:     store,
		gameModel: gameModel,
		logger:    logger,
	}
}

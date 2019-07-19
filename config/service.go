package config

import (
	"fmt"

	"github.com/nokka/slashdiablo-launcher/log"
	"github.com/nokka/slashdiablo-launcher/storage"
)

// Service is responsible for all things related to configuration.
type Service interface {
	Read() (*storage.Config, error)
	AddGame()
	UpsertGame(request UpdateGameRequest) error
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

// AddGame adds a new game to the game model.
func (s *service) AddGame() {
	g := NewGame(nil)

	// Generate ID next in the sequence.
	// TODO: Generate a stronger id.
	g.ID = len(s.gameModel.Games()) + 1

	fmt.Println("GENERATED ID", g.ID)
	g.Instances = 1

	s.gameModel.AddGame(g)
}

// UpdateGameRequest ...
type UpdateGameRequest struct {
	ID        int    `json:"id"`
	Location  string `json:"location"`
	Instances int    `json:"instances"`
	Maphack   bool   `json:"maphack"`
	HD        bool   `json:"hd"`
}

// UpsertGame will upsert the game to the config.
func (s *service) UpsertGame(request UpdateGameRequest) error {
	conf, err := s.store.Read()
	if err != nil {
		s.logger.Log("failed to read config", err)
		return err
	}

	// If the item to be updated isn't found, create a new one.
	var found bool

	// Look for game to update, and mutate if found.
	for i := 0; i < len(conf.Games); i++ {
		if conf.Games[i].ID == request.ID {
			found = true
			conf.Games[i].Location = request.Location
			conf.Games[i].Instances = request.Instances
			conf.Games[i].Maphack = request.Maphack
			conf.Games[i].HD = request.HD
		}
	}

	// Game wasn't found, append a new one.
	if !found {
		fmt.Println("WASNT FOUND, APPEND")
		fmt.Println(request)
		g := storage.Game{
			ID:        request.ID,
			Location:  request.Location,
			Instances: request.Instances,
			Maphack:   request.Maphack,
			HD:        request.HD,
		}

		// Add game to the config.
		conf.Games = append(conf.Games, g)
	}

	err = s.store.Write(conf)
	if err != nil {
		s.logger.Log("failed to write config", err)
		return err
	}

	// Updates game model with the new information.
	var updatedIndex int
	games := s.gameModel.Games()
	for i := 0; i < len(games); i++ {
		if games[i].ID == request.ID {
			updatedIndex = i
			games[i].Location = request.Location
			games[i].Instances = request.Instances
			games[i].Maphack = request.Maphack
			games[i].HD = request.HD
		}
	}

	// Notify the UI of the change.
	s.gameModel.updateGame(updatedIndex)

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

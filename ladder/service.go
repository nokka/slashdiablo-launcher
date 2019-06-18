package ladder

import (
	"github.com/nokka/slash-launcher/log"
)

// Service is responsible for all things related to the Slashdiablo ladder.
type Service interface {
	SetLadderCharacters(mode string) error
}

type service struct {
	sdClient    Client
	ladderModel *TopLadderModel
	logger      log.Logger
}

// GetLadder will fetch the ladder from the Slashdiablo API.
func (s *service) SetLadderCharacters(mode string) error {
	characters, err := s.sdClient.GetLadder(mode)
	if err != nil {
		return err
	}

	// Set the top 10 ladder positions.
	topChars := characters[:10]

	for _, char := range topChars {
		c := NewCharacter(nil)
		c.Rank = char.Rank
		c.Name = char.Name
		c.Class = char.Class
		c.Level = char.Level
		s.ladderModel.AddCharacter(c)
	}

	return nil
}

// NewService returns a service with all the dependencies.
func NewService(
	sdClient Client,
	ladderModel *TopLadderModel,
	logger log.Logger,
) Service {
	return &service{
		sdClient:    sdClient,
		ladderModel: ladderModel,
		logger:      logger,
	}
}

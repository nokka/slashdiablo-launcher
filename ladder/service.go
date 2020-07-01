package ladder

import (
	"errors"

	ladderClient "github.com/nokka/slashdiablo-launcher/clients/ladder"
)

// Service is responsible for all things related to the Slashdiablo ladder.
type Service interface {
	SetLadderCharacters(mode string) error
}

type service struct {
	client      ladderClient.Client
	ladderModel *TopLadderModel
}

// SetLadderCharacters will fetch the ladder from the Slashdiablo API.
func (s *service) SetLadderCharacters(mode string) error {
	characters, err := s.client.GetLadder(mode)
	if err != nil {
		return err
	}

	if len(characters) >= 10 {
		// Just grab the top 10 characters.
		topChars := characters[:10]

		for _, char := range topChars {
			s.ladderModel.AddCharacter(newCharacter(char))
		}
	} else {
		return errors.New("missing ladder characters")
	}

	return nil
}

// newCharacter will create a new QObject character that we can pass to the model.
func newCharacter(char ladderClient.Character) *Character {
	c := NewCharacter(nil)
	c.Rank = char.Rank
	c.Name = char.Name
	c.Class = getShortClassName(char.Class)
	c.Level = char.Level
	c.Title = char.Title
	c.Status = char.Status
	return c
}

func getShortClassName(name string) string {
	switch name {
	case "Assassin":
		// Assassin gets shortened to "Ass", we don't want that.
		return "Asn"
	default:
		return name[:3]
	}
}

// NewService returns a service with all the dependencies.
func NewService(
	client ladderClient.Client,
	ladderModel *TopLadderModel,
) Service {
	return &service{
		client:      client,
		ladderModel: ladderModel,
	}
}

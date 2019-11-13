package news

import (
	"github.com/nokka/slashdiablo-launcher/clients/slashdiablo"
)

// Service is responsible for all things related to the news.
type Service interface {
	SetNewsItems() error
}

type service struct {
	client    slashdiablo.Client
	newsModel *Model
}

// SetNewsItems will fetch the news from the Slashdiablo server.
func (s *service) SetNewsItems() error {
	i := NewItem(nil)
	i.Title = "derp"
	i.Text = "Definitely news..."
	i.Date = "8 NOV"
	i.Year = "2019"

	s.newsModel.AddItem(i)

	return nil
}

// NewService returns a service with all the dependencies.
func NewService(
	client slashdiablo.Client,
	newsModel *Model,
) Service {
	return &service{
		client:    client,
		newsModel: newsModel,
	}
}

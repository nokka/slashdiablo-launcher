package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

// JSONItem represents a news item from JSON, before it's turned into a model item.
type JSONItem struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Date  string `json:"date"`
	Year  string `json:"year"`
}

// SetNewsItems will fetch the news from the Slashdiablo server.
func (s *service) SetNewsItems() error {
	contents, err := s.client.GetNews()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(contents)
	if err != nil {
		return err
	}

	var newsItems []JSONItem
	if err := json.Unmarshal(bytes, &newsItems); err != nil {
		return err
	}

	if len(newsItems) > 3 {
		newsItems = newsItems[:3]
	}

	for _, i := range newsItems {
		// Shorten text if it's too long.
		if len(i.Text) > 275 {
			i.Text = fmt.Sprintf("%s...", i.Text[:272])
		}
		s.newsModel.AddItem(newItem(i))
	}

	return nil
}

// newItem will create a new QObject item that we can pass to the model.
func newItem(item JSONItem) *Item {
	i := NewItem(nil)
	i.Title = item.Title
	i.Text = item.Text
	i.Date = item.Date
	i.Year = item.Year
	return i
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

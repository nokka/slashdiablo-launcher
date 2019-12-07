package bridge

import (
	"github.com/nokka/slashdiablo-launcher/log"
	"github.com/nokka/slashdiablo-launcher/news"
	"github.com/therecipe/qt/core"
)

// NewsBridge is the connection between QML and the news model.
type NewsBridge struct {
	core.QObject

	// Dependencies.
	newsService news.Service
	logger      log.Logger

	// Properties.
	_ bool `property:"loading"`
	_ bool `property:"error"`

	// Models.
	NewsModel *core.QAbstractListModel `property:"items"`

	// Slots.
	_ func() `slot:"getNews"`
}

// Connect will connect the QML signals to functions in Go.
func (b *NewsBridge) Connect() {
	b.ConnectGetNews(b.getNews)
}

func (b *NewsBridge) getNews() {
	go func() {
		// Tell the GUI that we're fetching data.
		b.SetLoading(true)

		// Set the news items on the model.
		err := b.newsService.SetNewsItems()

		// Stop loading when we're done fetching news data.
		b.SetLoading(false)

		if err != nil {
			b.logger.Error(err)
			b.SetError(true)
			return
		}
	}()

	return
}

// NewNews sets up a news bridge with all dependencies.
func NewNews(ns news.Service, nm *news.Model, logger log.Logger) *NewsBridge {
	l := NewNewsBridge(nil)

	// Setup dependencies.
	l.newsService = ns
	l.logger = logger

	// Setup model.
	l.SetItems(nm)

	// Set initial state.
	l.SetLoading(false)
	l.SetError(false)

	return l
}

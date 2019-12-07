package news

import "github.com/therecipe/qt/core"

// Item represents a news item in the model.
type Item struct {
	core.QObject

	Title string
	Text  string
	Date  string
	Year  string
	Link  string
}

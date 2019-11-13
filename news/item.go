package news

import "github.com/therecipe/qt/core"

// Item ...
type Item struct {
	core.QObject

	Title string
	Text  string
	Date  string
	Year  string
}

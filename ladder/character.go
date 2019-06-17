package ladder

import "github.com/therecipe/qt/core"

// Character ...
type Character struct {
	core.QObject

	Name  string
	Class string
	Level int
	Rank  int
}

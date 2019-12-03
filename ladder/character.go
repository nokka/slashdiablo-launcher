package ladder

import "github.com/therecipe/qt/core"

// Character represents a Diablo character in the model.
type Character struct {
	core.QObject

	Name   string
	Class  string
	Level  int
	Rank   int
	Title  string
	Status string
}

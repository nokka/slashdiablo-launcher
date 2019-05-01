package d2

import "fmt"

// Launcher is responible for launching Diablo II.
type Launcher struct {
	Path string
}

// Exec will exec the game.
func (l *Launcher) Exec() {
	fmt.Println("LAUNCH on path", l.Path)
}

// NewLauncher returns a launcher with all the dependencies setup.
func NewLauncher(path string) *Launcher {
	return &Launcher{
		Path: path,
	}
}

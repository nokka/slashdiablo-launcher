package bridge

import (
	"fmt"

	"github.com/nokka/slash-launcher/d2"
	"github.com/therecipe/qt/core"
)

// DiabloBridge is the connection between QML and Go, it facilitates
// a way to setup signals that can be interpreted in Go code.
type DiabloBridge struct {
	core.QObject

	// Services.
	D2service *d2.Service

	// Properties.
	_ float32 `property:"patchProgress"`

	// Functions.
	_ func() `signal:"launchGame"`
	_ func() `slot:"patchGame"`
}

// Connect will connect the QML signals to functions in Go.
func (q *DiabloBridge) Connect() {
	q.ConnectLaunchGame(q.launchGame)
	q.ConnectPatchGame(q.patchGame)
}

func (q *DiabloBridge) patchGame() {
	fmt.Println("PATCHING GAME")
	// Run this on a seperate thread so we don't block the UI.
	go func() {
		done := make(chan bool, 1)

		// Let the patcher run, it returns a channel
		// where we get the progress from, and another channel withe errors.
		progress, errors := q.D2service.Patch(done)

		for {
			select {
			case percentage := <-progress:
				fmt.Println("Patching progress", percentage)
				q.SetPatchProgress(percentage)
			case err := <-errors:
				fmt.Println("Received error", err)
				// @TODO: Add QML signal.
				return
			case <-done:
				return
			}
		}
	}()
}

func (q *DiabloBridge) launchGame() {
	err := q.D2service.Exec()
	if err != nil {
		fmt.Println(err)
		// @TODO: Add QML signal.
	}
}

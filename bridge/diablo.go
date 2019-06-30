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
	_ bool    `property:"patching"`
	_ bool    `property:"errored"`
	_ bool    `property:"playable"`
	_ float32 `property:"patchProgress"`
	_ string  `property:"status"`

	// Functions.
	_ func() `slot:"launchGame"`
	_ func() `slot:"checkForUpdates"`
}

// Connect will connect the QML signals to functions in Go.
func (q *DiabloBridge) Connect() {
	q.ConnectLaunchGame(q.launchGame)
	q.ConnectCheckForUpdates(q.checkForUpdates)
}

func (q *DiabloBridge) launchGame() {
	err := q.D2service.Exec()
	if err != nil {
		fmt.Println(err)
		// @TODO: Add QML signal.
	}
}

func (q *DiabloBridge) checkForUpdates() {
	// Tell the GUI we've started patching.
	q.SetPatching(true)
	q.SetPlayable(false)

	// Run this on a seperate thread so we don't block the UI.
	go func() {
		done := make(chan bool, 1)

		// Let the patcher run, it returns a channel
		// where we get the progress from, and another channel with errors.
		progress, state := q.D2service.Patch(done)

		for {
			select {
			case percentage := <-progress:
				fmt.Println("Patching progress", percentage)
				q.SetPatchProgress(percentage)
			case current := <-state:
				fmt.Println("Received state", current)
				if current.Error != nil {
					q.SetErrored(true)
					q.SetPatching(false)
				}

				if current.Message != "" {
					q.SetStatus(current.Message)
				}
			case <-done:
				fmt.Println("Got done")
				q.SetPatching(false)
				q.SetPlayable(true)
				return
			}
		}
	}()
}

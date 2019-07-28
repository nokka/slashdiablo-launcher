package bridge

import (
	"github.com/nokka/slashdiablo-launcher/d2"
	"github.com/nokka/slashdiablo-launcher/log"
	"github.com/therecipe/qt/core"
)

// DiabloBridge is the connection between QML and Go, it facilitates
// a way to setup signals that can be interpreted in Go code.
type DiabloBridge struct {
	core.QObject

	// Dependencies.
	d2service *d2.Service
	logger    log.Logger

	// Properties.
	_ bool    `property:"patching"`
	_ bool    `property:"errored"`
	_ bool    `property:"playable"`
	_ bool    `property:"validVersion"`
	_ float32 `property:"patchProgress"`
	_ string  `property:"status"`

	// Slots.
	_ func() `slot:"launchGame"`
	_ func() `slot:"validateVersion"`
	_ func() `slot:"applyPatches"`
	_ func() `slot:"runDEPFix"`
}

// Connect will connect the QML signals to functions in Go.
func (q *DiabloBridge) Connect() {
	q.ConnectLaunchGame(q.launchGame)
	q.ConnectApplyPatches(q.applyPatches)
	q.ConnectValidateVersion(q.validateVersion)
	q.ConnectRunDEPFix(q.runDEPFix)
}

func (q *DiabloBridge) launchGame() {
	err := q.d2service.Exec()
	if err != nil {
		q.logger.Error(err)
		// @TODO: Add QML signal.
	}
}

func (q *DiabloBridge) applyPatches() {
	// Tell the GUI we've started patching.
	q.SetPatching(true)
	q.SetPlayable(false)

	// Run this on a seperate thread so we don't block the UI.
	go func() {
		done := make(chan bool, 1)

		// Let the patcher run, it returns a channel
		// where we get the progress from, and another channel with errors.
		progress, state := q.d2service.Patch(done)

		for {
			select {
			case percentage := <-progress:
				q.SetPatchProgress(percentage)
			case current := <-state:
				if current.Error != nil {
					// Log the error to persistant logging store.
					q.logger.Error(current.Error)

					// Update bridge state.
					q.SetErrored(true)
					q.SetPatching(false)
				}

				if current.Message != "" {
					q.SetStatus(current.Message)
				}
			case <-done:
				q.SetPatching(false)
				q.validateVersion()
				return
			}
		}
	}()
}

func (q *DiabloBridge) validateVersion() {
	valid, err := q.d2service.ValidateGameVersions()
	if err != nil {
		q.logger.Error(err)
		q.SetErrored(true)
		return
	}

	if valid {
		q.SetPlayable(true)
	}

	q.SetValidVersion(valid)
}

func (q *DiabloBridge) runDEPFix() {
	err := q.d2service.RunDEPFix()
	if err != nil {
		q.logger.Error(err)
		// @TODO: Add QML signal.
	}
}

// NewDiablo ...
func NewDiablo(d2s *d2.Service, logger log.Logger) *DiabloBridge {
	b := NewDiabloBridge(nil)

	// Set dependencies.
	b.d2service = d2s
	b.logger = logger

	// Set initial state.
	b.SetPatching(false)
	b.SetErrored(false)
	b.SetPlayable(false)

	return b
}

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
	d2service d2.Service
	logger    log.Logger

	// Properties.
	_ bool    `property:"patching"`
	_ bool    `property:"errored"`
	_ bool    `property:"validVersion"`
	_ bool    `property:"validatingVersion"`
	_ bool    `property:"launching"`
	_ float32 `property:"patchProgress"`
	_ string  `property:"status"`
	_ string  `property:"gateway"`
	_ int     `property:"launchDelay"`

	// Slots.
	_ func()                 `slot:"launchGame"`
	_ func()                 `slot:"validateVersion"`
	_ func()                 `slot:"applyPatches"`
	_ func(path string) bool `slot:"applyDEP"`
	_ func(gateway string)   `slot:"updateGateway"`
	_ func(delay int)        `slot:"updateLaunchDelay"`
}

// Connect will connect the QML signals to functions in Go.
func (b *DiabloBridge) Connect() {
	b.ConnectLaunchGame(b.launchGame)
	b.ConnectApplyPatches(b.applyPatches)
	b.ConnectValidateVersion(b.validateVersion)
	b.ConnectApplyDEP(b.applyDEP)
	b.ConnectUpdateGateway(b.updateGateway)
	b.ConnectUpdateLaunchDelay(b.updateLaunchDelay)
}

func (b *DiabloBridge) launchGame() {
	// Tell the GUI that we're launching games.
	b.SetLaunching(true)

	// Do the work on another thread not to lock the GUI.
	go func() {
		err := b.d2service.Exec()
		if err != nil {
			b.logger.Error(err)
		}

		// Done launching.
		b.SetLaunching(false)
	}()

}

func (b *DiabloBridge) applyPatches() {
	// Tell the GUI we've started patching.
	b.SetPatching(true)
	b.SetValidVersion(false)

	// Run this on a separate thread so we don't block the UI.
	go func() {
		done := make(chan bool, 1)

		// Let the patcher run, it returns a channel
		// where we get the progress from, and another channel with errors.
		progress, state := b.d2service.Patch(done)

		for {
			select {
			case percentage := <-progress:
				b.SetPatchProgress(percentage)
			case current := <-state:
				if current.Error != nil {
					// Log the error to persistent logging store.
					b.logger.Error(current.Error)

					// Update bridge state.
					b.SetErrored(true)
					b.SetPatching(false)
				}

				if current.Message != "" {
					b.SetStatus(current.Message)
				}
			case <-done:
				b.SetPatching(false)
				b.validateVersion()
				return
			}
		}
	}()
}

func (b *DiabloBridge) validateVersion() {
	// Update GUI and reset errors.
	b.SetValidatingVersion(true)
	b.SetErrored(false)

	// Do the work on another thread not to lock the GUI.
	go func() {
		valid, err := b.d2service.ValidateGameVersions()
		if err != nil {
			b.logger.Error(err)
			b.SetErrored(true)
		}

		b.SetValidVersion(valid)
		b.SetValidatingVersion(false)
	}()
}

func (b *DiabloBridge) applyDEP(path string) bool {
	err := b.d2service.ApplyDEP(path)
	if err != nil {
		b.logger.Error(err)
		return false
	}

	return true
}

func (b *DiabloBridge) updateGateway(gateway string) {
	err := b.d2service.SetGateway(gateway)
	if err != nil {
		b.logger.Error(err)
	}

	// Gateway was successfully saved, set gateway on the bridge.
	b.SetGateway(gateway)
}

func (b *DiabloBridge) updateLaunchDelay(delay int) {
	err := b.d2service.SetLaunchDelay(delay)
	if err != nil {
		b.logger.Error(err)
	}

	// Launch delay was successfully saved, set launch delay on the bridge.
	b.SetLaunchDelay(delay)
}

// NewDiablo returns a new Diablo bridge with all dependencies set up.
func NewDiablo(d2s d2.Service, gateway string, launchDelay int, logger log.Logger) *DiabloBridge {
	b := NewDiabloBridge(nil)

	// Set dependencies.
	b.d2service = d2s
	b.logger = logger

	// Set initial state.
	b.SetPatching(false)
	b.SetLaunching(false)
	b.SetErrored(false)
	b.SetValidVersion(false)
	b.SetValidatingVersion(false)
	b.SetGateway(gateway)
	b.SetLaunchDelay(launchDelay)

	return b
}

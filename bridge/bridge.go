package bridge

import (
	"fmt"
	"os"

	"github.com/nokka/slash-launcher/d2"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

// QmlBridge is the connection between QML and Go, it facilitates
// a way to setup signals that can be interpreted in Go code.
type QmlBridge struct {
	core.QObject

	// Reference to main view.
	View *quick.QQuickView

	// Game launcher
	D2service *d2.Service

	// Patching progress.
	_ float32 `property:"patchProgress"`

	_ func() `signal:"closeLauncher"`
	_ func() `signal:"minimizeLauncher"`
	_ func() `signal:"launchGame"`

	_ func() `slot:"patchGame"`
	_ func() `slot:"checkGameLocation"`
}

// Connect will connect the QML signals to functions in Go.
func (q *QmlBridge) Connect() {
	q.ConnectCloseLauncher(func() {
		os.Exit(0)
	})

	q.ConnectMinimizeLauncher(func() {
		//q.View.SetWindowState(core.Qt__WindowMinimized)
	})

	q.ConnectLaunchGame(func() {
		q.D2service.Exec()
	})

	q.ConnectPatchGame(func() {
		/*go func() {
			// Let the patcher run, it returns a channel
			// where we get the progress from.
			progress := q.D2service.Patch()

			// Range over the percentages complete.
			for percentage := range progress {
				q.SetPatchProgress(percentage)
			}
		}()*/
	})

	q.ConnectCheckGameLocation(func() {
		fmt.Println("Checking game location")
	})
}

package bridge

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

// QmlBridge is the connection between QML and Go, it facilitates
// a way to setup signals that can be interpreted in Go code.
type QmlBridge struct {
	core.QObject

	// Reference to main view.
	View *quick.QQuickView

	_ func() `slot:"closeLauncher"`
	_ func() `slot:"minimizeLauncher"`
}

// Connect will connect the QML signals to functions in Go.
func (q *QmlBridge) Connect() {
	q.ConnectCloseLauncher(func() {
		fmt.Println("Closing Launcher")
		os.Exit(0)
	})

	q.ConnectMinimizeLauncher(func() {
		fmt.Println("Minimizing Launcher")
		q.View.SetWindowState(core.Qt__WindowMinimized)
	})
}

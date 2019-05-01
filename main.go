package main

import (
	"os"

	"github.com/nokka/slash-launcher/bridge"
	"github.com/nokka/slash-launcher/d2"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// Enable high dpi scaling, useful for devices with high pixel density displays.
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Create the new QApplication.
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create the view and configure it.
	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetFlags(core.Qt__FramelessWindowHint)

	// Set our main.qml to the view.
	view.SetSource(core.NewQUrl3("qml/main.qml", 0))

	// Create a new QML bridge that will bridge the client to Go.
	var qmlBridge = bridge.NewQmlBridge(nil)

	// Setup the bridge dependencies.
	qmlBridge.D2Launcher = d2.NewLauncher("/")
	qmlBridge.View = view

	// Connect the QML signals on the bridge to Go.
	qmlBridge.Connect()

	// Set the bridge on the view.
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	// Make the view visible.
	view.Show()

	// Finally, execute the application.
	app.Exec()
}

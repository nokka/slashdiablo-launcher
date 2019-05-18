package main

import (
	"fmt"
	"os"

	"github.com/nokka/slash-launcher/bridge"
	"github.com/nokka/slash-launcher/d2"
	"github.com/nokka/slash-launcher/github"
	"github.com/nokka/slash-launcher/storage"
	"github.com/nokka/slash-launcher/window"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// Environment variables set when building.
	var (
		githubOwner      = envString("GITHUB_OWNER", "")
		githubRepository = envString("GITHUB_REPO", "")
	)

	// Enable high dpi scaling, useful for devices with high pixel density displays.
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Create the new QApplication.
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create the view and configure it.
	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetFlags(core.Qt__FramelessWindowHint |
		core.Qt__WindowMinimizeButtonHint |
		core.Qt__Window,
	)

	// Create a new QML bridge that will bridge the GUI to Go.
	var qmlBridge = bridge.NewQmlBridge(nil)

	// Setup local storage.
	store := storage.NewStore()
	if err := store.Load(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Setup services.
	gs := github.NewService(githubOwner, githubRepository)
	d2s := d2.NewService("/tmp", gs, store)

	// Setup the bridge dependencies.
	qmlBridge.D2service = d2s
	qmlBridge.View = view

	// Connect the QML signals on the bridge to Go.
	qmlBridge.Connect()

	// Set the bridge on the view.
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	// Set our main.qml to the view.
	view.SetSource(core.NewQUrl3("qml/main.qml", 0))
	//view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))

	// Allows for windows to minimize on Darwin.
	window.AllowMinimize(view.WinId())

	// Center the view.
	view.SetPosition2(
		(widgets.QApplication_Desktop().Screen(0).Width()-view.Width())/2,
		(widgets.QApplication_Desktop().Screen(0).Height()-view.Height())/2,
	)

	// Make the view visible.
	view.Show()

	// Finally, execute the application.
	app.Exec()
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

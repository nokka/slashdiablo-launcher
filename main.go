package main

import (
	"errors"
	"os"

	"github.com/nokka/slash-launcher/bridge"
	"github.com/nokka/slash-launcher/config"
	"github.com/nokka/slash-launcher/d2"
	"github.com/nokka/slash-launcher/github"
	"github.com/nokka/slash-launcher/log"
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

	// Set app context.
	core.QCoreApplication_SetApplicationName("slashdiablo.launcher")
	core.QCoreApplication_SetOrganizationName("slashdiablo.com")

	// Enable high dpi scaling, useful for devices with high pixel density displays.
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Create base application.
	app := widgets.NewQApplication(0, nil)

	// Create a new frameless window, this is the root window.
	fw := window.NewFramelessWindow(1.0, 1024, 600)

	// Create a new QML widget, this is what we will draw the application on.
	qmlWidget := newQmlWidget()

	// Setup layout that will be added to the root window.
	layout := newLayout()
	layout.AddWidget(qmlWidget, 0, 0)

	// Add the layout to the window, this is the only item on the base window.
	fw.SetupContent(layout)

	configPath, err := getConfigPath()
	if err != nil {
		os.Exit(0)
	}

	// Data directory is a requirement for the app.
	os.MkdirAll(configPath, storage.Permissions)

	// Setup file logger.
	logger := log.NewLogger(configPath)

	// Setup local storage.
	store := storage.NewStore(configPath)
	if err := store.Load(); err != nil {
		logger.Log("unable to load config", err)
		os.Exit(0)
	}

	conf, err := store.Read()
	if err != nil {
		logger.Log("unable to read config", err)
		os.Exit(0)
	}

	// Setup services.
	gs := github.NewService(githubOwner, githubRepository)
	cs := config.NewService(store, logger)
	d2s := d2.NewService(gs, cs, logger)

	// Models.
	lm := NewLadderModel(nil)

	// Setup QML bridge with all dependencies.
	qmlBridge := bridge.NewQmlBridge(nil)
	qmlBridge.D2service = d2s

	// Setup config bridge with all dependencies.
	configBridge := bridge.NewConfigBridge(nil)
	configBridge.Configuration = cs

	// Set ladder model on the bridge.
	qmlBridge.SetLadderCharacters(lm)

	// Set values from disk on the config bridge.
	configBridge.SetD2Location(conf.D2Location)
	configBridge.SetD2Instances(conf.D2Instances)
	configBridge.SetHDLocation(conf.HDLocation)
	configBridge.SetHDInstances(conf.HDInstances)

	// Add bridge to QML.
	qmlWidget.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	qmlBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("settings", configBridge)
	configBridge.Connect()

	// Make sure the window is allowed to minimize.
	window.AllowMinimize(fw.WinId())

	// Set the source for our drawable widget to our qml entrypoint.
	//qmlWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))
	qmlWidget.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))

	fw.Show()
	fw.Widget.SetFocus2()

	app.Exec()
}

// newQmlWidget returns a configured QML widget.
func newQmlWidget() *quick.QQuickWidget {
	var qwidget = quick.NewQQuickWidget(nil)
	qwidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	return qwidget
}

// newLayout returns a configured layout.
func newLayout() *widgets.QVBoxLayout {
	layout := widgets.NewQVBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)
	return layout
}

// getConfigPath returns the target specific application data directory.
func getConfigPath() (string, error) {
	locations := core.QStandardPaths_StandardLocations(
		core.QStandardPaths__AppLocalDataLocation,
	)
	if len(locations) == 0 {
		return "", errors.New("failed to locate application data directory")
	}

	// Grab the first available location.
	return locations[0], nil
}

// envString extracts a string from os environment.
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

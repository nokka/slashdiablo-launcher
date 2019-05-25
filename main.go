package main

import (
	"errors"
	"os"

	"github.com/nokka/slash-launcher/bridge"
	"github.com/nokka/slash-launcher/d2"
	"github.com/nokka/slash-launcher/github"
	"github.com/nokka/slash-launcher/log"
	"github.com/nokka/slash-launcher/storage"
	"github.com/nokka/slash-launcher/window"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
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

	ga := gui.NewQGuiApplication(len(os.Args), os.Args)
	ga.SetWindowIcon(gui.NewQIcon5(":/qml/assets/tmp_icon.png"))

	configPath, err := getConfigPath()
	if err != nil {
		os.Exit(0)
	}

	// Setup logger.
	logger := log.NewLogger(configPath)

	// The data directory is a requirement for the app.
	os.MkdirAll(configPath, storage.Permissions)

	// Setup local storage.
	store := storage.NewStore(configPath)
	if err := store.Load(); err != nil {
		logger.Log("unable to load config", err)
		os.Exit(0)
	}

	config, err := store.Read()
	if err != nil {
		logger.Log("unable to read config", err)
		os.Exit(0)
	}

	// Setup services.
	gs := github.NewService(githubOwner, githubRepository)
	d2s := d2.NewService("/tmp", gs, store, logger)

	// Create a new QML bridge that will bridge the GUI to Go.
	var qmlBridge = bridge.NewQmlBridge(nil)

	// Initiate the config bridge.
	configBridge := bridge.NewConfigBridge(nil)
	configBridge.SetD2Location(config.D2Location)
	configBridge.SetD2Instances(config.D2Instances)
	configBridge.SetHDLocation(config.HDLocation)
	configBridge.SetHDInstances(config.HDInstances)
	configBridge.Connect()

	// Setup the bridge dependencies.
	qmlBridge.D2service = d2s

	// Connect the QML signals on the bridge to Go.
	qmlBridge.Connect()

	// Setup QML engine.
	qmlEngine := qml.NewQQmlApplicationEngine(nil)

	qmlEngine.ConnectObjectCreated(func(object *core.QObject, url *core.QUrl) {
		if object.ObjectName() == "mainWindow" {
			window.AllowMinimize(gui.NewQWindowFromPointer(object.Pointer()).WinId())
		}
	})

	// Connect the qml bridge to QML.
	qmlEngine.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	qmlEngine.RootContext().SetContextProperty("settings", configBridge)

	// Set our main.qml to the view.
	//qmlEngine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	qmlEngine.Load(core.NewQUrl3("qml/main.qml", 0))

	// Finally, exec the application.
	gui.QGuiApplication_Exec()
}

// Returns the target specific application data directory.
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

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

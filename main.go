package main

import (
	"bufio"
	"errors"
	"fmt"
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

// Launcher ...
type Launcher struct {
	fw *window.QFramelessWindow

	app *widgets.QApplication
	win *widgets.QMainWindow
}

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
	//core.QCoreApplication_SetAttribute(core.Qt__AA_UseSoftwareOpenGL, true)

	a := &Launcher{}
	a.app = widgets.NewQApplication(0, nil)

	a.fw = window.CreateQFramelessWindow(1.0)

	qmlWidget := newQmlWidget()

	layout := widgets.NewQVBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)
	layout.AddWidget(qmlWidget, 0, 0)

	a.fw.SetupContent(layout)
	a.fw.SetupWidgetColor(0, 0, 0)

	// TODO: Refactor
	lm := NewLadderModel(nil)

	lm.AddCharacter(&Character{
		Name:  "Meanski",
		Class: "pal",
		Level: "99",
	})

	lm.AddCharacter(&Character{
		Name:  "Nolan",
		Class: "sor",
		Level: "95",
	})

	lm.AddCharacter(&Character{
		Name:  "Nolan",
		Class: "sor",
		Level: "95",
	})

	configPath, err := getConfigPath()
	if err != nil {
		os.Exit(0)
	}

	// Data directory is a requirement for the app.
	os.MkdirAll(configPath, storage.Permissions)

	// Setup logger.
	logger := log.NewLogger(configPath)

	// STD LOGGING
	r, w, err := os.Pipe()
	if err != nil {
		os.Exit(0)
	}

	os.Stderr = w

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			logger.Log("stdout", line)
		}
	}()

	fmt.Printf("CAPTURING STDOUT \n")

	// STD LOGGING END

	// Setup local storage.
	store := storage.NewStore(configPath)
	if err := store.Load(); err != nil {
		logger.Log("unable to load config", err)
		os.Exit(0)
	}

	fmt.Printf("STORE SETUP \n")

	conf, err := store.Read()
	if err != nil {
		logger.Log("unable to read config", err)
		os.Exit(0)
	}

	fmt.Printf("STORE READ \n")

	// Setup services.
	gs := github.NewService(githubOwner, githubRepository)
	cs := config.NewService(store, logger)
	d2s := d2.NewService(gs, cs, logger)

	fmt.Printf("SERVICES SETUP \n")

	// Create a new QML bridge that will bridge the GUI to Go.
	qmlBridge := bridge.NewQmlBridge(nil)
	qmlBridge.D2service = d2s

	// Initiate the config bridge.
	configBridge := bridge.NewConfigBridge(nil)
	configBridge.Configuration = cs

	// Setup bridges.
	qmlWidget.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	qmlWidget.RootContext().SetContextProperty("settings", configBridge)

	fmt.Printf("QML CONNECTED \n")
	// Connect the QML signals on the bridge to Go.
	qmlBridge.Connect()
	configBridge.Connect()

	// Start using the connections.
	qmlBridge.SetLadderCharacters(lm)
	configBridge.SetD2Location(conf.D2Location)
	configBridge.SetD2Instances(conf.D2Instances)
	configBridge.SetHDLocation(conf.HDLocation)
	configBridge.SetHDInstances(conf.HDInstances)

	window.AllowMinimize(a.fw.WinId())

	fmt.Printf("MINIMIZE SETUP \n")

	qmlWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))

	fmt.Printf("QML SOURCE SET TO MAIN.qml \n")

	a.fw.Show()
	a.fw.Widget.SetFocus2()
	a.app.Exec()
}

func newQmlWidget() *quick.QQuickWidget {
	var quickWidget = quick.NewQQuickWidget(nil)
	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	//quickWidget.SetSource(core.NewQUrl3("qrc:/qml/main. qml", 0))
	//quickWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))

	return quickWidget
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

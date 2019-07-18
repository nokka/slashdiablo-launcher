package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/nokka/goqmlframeless"
	"github.com/nokka/slash-launcher/bridge"
	"github.com/nokka/slash-launcher/config"
	"github.com/nokka/slash-launcher/d2"
	"github.com/nokka/slash-launcher/github"
	"github.com/nokka/slash-launcher/ladder"
	"github.com/nokka/slash-launcher/log"
	"github.com/nokka/slash-launcher/storage"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// Environment variables set when building.
	var (
		githubOwner      = envString("GITHUB_OWNER", "")
		githubRepository = envString("GITHUB_REPO", "")
		githubToken      = envString("GITHUB_TOKEN", "")
		ladderAddress    = envString("LADDER_ADDRESS", "")
		debugMode        = envBool("DEBUG_MODE", false)
	)

	// Set app context.
	core.QCoreApplication_SetApplicationName("slashdiablo.launcher")
	core.QCoreApplication_SetOrganizationName("slashdiablo.com")
	core.QCoreApplication_SetOrganizationDomain("slashdiablo.com")

	// Enable high dpi scaling, useful for devices with high pixel density displays.
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Create base application.
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create new frameless window.
	fw := goqmlframeless.NewWindow(goqmlframeless.Options{
		Width:        1024,
		Height:       600,
		Alpha:        1.0,
		Color:        goqmlframeless.RGB{R: 11, G: 3, B: 13},
		BorderRadius: 5,
		ShadowSize:   0,
	})

	// QML Widget that will be used to draw on.
	qmlWidget := quick.NewQQuickWidget(nil)
	qmlWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	// Add QML widget to layout.
	fw.Layout.AddWidget(qmlWidget, 0, 0)

	configPath, err := getConfigPath()
	if err != nil {
		os.Exit(0)
	}

	// Data directory is a requirement for the app.
	os.MkdirAll(configPath, storage.Permissions)

	// Setup file logger.
	logger := log.NewLogger(configPath)

	// Enable debugger if it was enabled through the env variable.
	if debugMode {
		enableDebugger(logger)
	}

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

	// Models.
	lm := ladder.NewTopLadderModel(nil)
	gm := config.NewGameModel(nil)

	// Setup clients.
	gc := github.NewClient(githubOwner, githubRepository, githubToken)
	lc := ladder.NewClient(ladderAddress)

	// Setup services.
	cs := config.NewService(store, gm, logger)
	d2s := d2.NewService(gc, cs, logger)
	ls := ladder.NewService(lc, lm, logger)

	// Setup QML bridges with all dependencies.
	diabloBridge := bridge.NewDiabloBridge(nil)
	diabloBridge.D2service = d2s
	diabloBridge.SetPatching(false)
	diabloBridge.SetErrored(false)
	diabloBridge.SetPlayable(false)

	configBridge := newConfigBridge(cs, conf, gm)

	ladderBridge := bridge.NewLadderBridge(nil)
	ladderBridge.LadderService = ls
	ladderBridge.SetCharacters(lm)
	ladderBridge.SetLoading(false)
	ladderBridge.SetError(false)

	// Add bridges to QML.
	qmlWidget.RootContext().SetContextProperty("diablo", diabloBridge)
	diabloBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("settings", configBridge)
	configBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("ladder", ladderBridge)
	ladderBridge.Connect()

	// Make sure the window is allowed to minimize.
	goqmlframeless.AllowMinimize(fw.WinId())

	// Set the source for our drawable widget to our qml entrypoint.
	qmlWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))
	//qmlWidget.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))

	fw.Show()
	app.Exec()
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

func newConfigBridge(cs config.Service, conf *storage.Config, gm *config.GameModel) *bridge.ConfigBridge {
	configBridge := bridge.NewConfigBridge(nil)

	for _, game := range conf.Games {
		g := config.NewGame(nil)
		g.Location = game.Location
		g.Instances = game.Instances
		g.Maphack = game.Maphack
		g.HD = game.HD

		fmt.Println("ADDING GAME")
		fmt.Println(g)
		gm.AddGame(g)
	}

	configBridge.Configuration = cs
	configBridge.SetD2Location(conf.D2Location)
	configBridge.SetD2Instances(conf.D2Instances)
	configBridge.SetD2Maphack(conf.D2Maphack)
	configBridge.SetHDLocation(conf.HDLocation)
	configBridge.SetHDInstances(conf.HDInstances)
	configBridge.SetHDMaphack(conf.HDMaphack)
	configBridge.SetNrOfGames(len(conf.Games))
	configBridge.SetGames(gm)

	return configBridge
}

// enableDebugger will capture stdout and stderr output.
func enableDebugger(logger log.Logger) {
	r, w, err := os.Pipe()
	if err != nil {
		os.Exit(0)
	}

	os.Stdout = w
	os.Stderr = w

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			logger.Debug(line + "\n")
		}
	}()
}

// envString extracts a string from os environment.
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

// envBool extracts a bool from os environment.
func envBool(env string, fallback bool) bool {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}

	v, err := strconv.ParseBool(e)
	if err != nil {
		return fallback
	}

	return v
}

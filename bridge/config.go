package bridge

import (
	"encoding/json"

	"github.com/nokka/slashdiablo-launcher/config"
	"github.com/nokka/slashdiablo-launcher/log"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

// ConfigBridge is the connection between QML and the Go config.
type ConfigBridge struct {
	core.QObject
	configPath string

	// Dependencies.
	config config.Service
	logger log.Logger

	// Models.
	GameModel *core.QAbstractListModel `property:"games"`

	// Properties.
	_ string   `property:"buildVersion"`
	_ []string `property:"availableHDMods"`
	_ []string `property:"availableMaphackMods"`
	_ bool     `property:"prerequisitesLoaded"`
	_ bool     `property:"prerequisitesError"`

	// Slots.
	_ func()                 `slot:"addGame"`
	_ func(body string) bool `slot:"upsertGame"`
	_ func(id string)        `slot:"deleteGame"`
	_ func() bool            `slot:"persistGameModel"`
	_ func()                 `slot:"getPrerequisites"`
	_ func()                 `slot:"openConfigPath"`
}

// Connect will connect the QML signals to functions in Go.
func (c *ConfigBridge) Connect() {
	c.ConnectUpsertGame(c.upsertGame)
	c.ConnectAddGame(c.addGame)
	c.ConnectDeleteGame(c.deleteGame)
	c.ConnectPersistGameModel(c.persistGameModel)
	c.ConnectGetPrerequisites(c.getPrerequisites)
	c.ConnectOpenConfigPath(c.openConfigPath)
}

// addGame will add a game to the game model.
func (c *ConfigBridge) addGame() {
	c.config.AddGame()
}

// upsertGame will update the game model.
func (c *ConfigBridge) upsertGame(body string) bool {
	var request config.UpdateGameRequest
	if err := json.Unmarshal([]byte(body), &request); err != nil {
		c.logger.Error(err)
		return false
	}

	err := c.config.UpsertGame(request)
	if err != nil {
		c.logger.Error(err)
		return false
	}

	return true
}

// deleteGame will delete the given id from the game model.
func (c *ConfigBridge) deleteGame(id string) {
	err := c.config.DeleteGame(id)
	if err != nil {
		c.logger.Error(err)
	}
}

// persistGameModel will persist the current game model to the config.
func (c *ConfigBridge) persistGameModel() bool {
	if err := c.config.PersistGameModel(); err != nil {
		c.logger.Error(err)
		return false
	}
	return true
}

// getPrerequisites will fetch all required config data.
func (c *ConfigBridge) getPrerequisites() {
	go func() {
		// Tell the UI that we're fetching prerequisites.
		c.SetPrerequisitesLoaded(false)
		c.SetPrerequisitesError(false)

		mods, err := c.config.GetAvailableMods()
		if err != nil {
			c.SetPrerequisitesError(true)
			c.logger.Error(err)
			return
		}

		// Default option for no mod at all.
		defaultMods := []string{config.ModVersionNone}
		c.SetAvailableHDMods(append(defaultMods, mods.HD...))
		c.SetAvailableMaphackMods(append(defaultMods, mods.Maphack...))

		c.SetPrerequisitesLoaded(true)
	}()
}

// openConfigPath will open the config path in the file explorer.
func (c *ConfigBridge) openConfigPath() {
	gui.QDesktopServices_OpenUrl(core.QUrl_FromLocalFile(c.configPath))
}

// NewConfig returns a new config bridge with all dependencies set up.
func NewConfig(cs config.Service, gm *config.GameModel, configPath string, logger log.Logger) *ConfigBridge {
	b := NewConfigBridge(nil)

	b.configPath = configPath

	// Setup dependencies.
	b.config = cs
	b.logger = logger

	// Setup model.
	b.SetGames(gm)

	// Set initial state.
	b.SetPrerequisitesLoaded(false)
	b.SetPrerequisitesError(false)

	return b
}

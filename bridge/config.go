package bridge

import (
	"encoding/json"
	"fmt"

	"github.com/nokka/slashdiablo-launcher/config"
	"github.com/nokka/slashdiablo-launcher/log"
	"github.com/therecipe/qt/core"
)

// ConfigBridge is the connection between QML and the Go config.
type ConfigBridge struct {
	core.QObject

	// Dependencies.
	config config.Service
	logger log.Logger

	// Models.
	GameModel *core.QAbstractListModel `property:"games"`

	// Properties.
	_ string   `property:"buildVersion"`
	_ []string `property:"availableHDMods"`
	_ []string `property:"availableMaphackMods"`

	// Slots.
	_ func()                 `slot:"addGame"`
	_ func(body string) bool `slot:"upsertGame"`
	_ func(id string)        `slot:"deleteGame"`
	_ func() bool            `slot:"persistGameModel"`
	_ func()                 `slot:"getAvailableMods"`
}

// Connect will connect the QML signals to functions in Go.
func (c *ConfigBridge) Connect() {
	c.ConnectUpsertGame(c.upsertGame)
	c.ConnectAddGame(c.addGame)
	c.ConnectDeleteGame(c.deleteGame)
	c.ConnectPersistGameModel(c.persistGameModel)
	c.ConnectGetAvailableMods(c.getAvailableMods)
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

// persistGameModel will persist the current game model to the config.
func (c *ConfigBridge) getAvailableMods() {
	// Do the work on another thread not to lock the GUI.
	//go func() {
	mods, err := c.config.GetAvailableMods()
	if err != nil {
		c.logger.Error(err)
	}

	fmt.Println("AVAILABLE MODS IN BRIDGE")
	fmt.Println(mods)

	// Default option for no mod at all.
	allMods := []string{config.HDVersionNone}
	c.SetAvailableHDMods(append(allMods, mods.HD...))
	//}()
}

// NewConfig returns a new config bridge with all dependencies set up.
func NewConfig(cs config.Service, gm *config.GameModel, logger log.Logger) *ConfigBridge {
	configBridge := NewConfigBridge(nil)

	// Setup dependencies.
	configBridge.config = cs
	configBridge.logger = logger

	// Setup model.
	configBridge.SetGames(gm)

	return configBridge
}

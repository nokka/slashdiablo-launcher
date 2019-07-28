package bridge

import (
	"encoding/json"

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

	// Slots.
	_ func(body string) bool `slot:"upsertGame"`
	_ func()                 `slot:"addGame"`
	_ func(id int)           `slot:"deleteGame"`
}

// Connect will connect the QML signals to functions in Go.
func (c *ConfigBridge) Connect() {
	c.ConnectUpsertGame(c.upsertGame)
	c.ConnectAddGame(c.addGame)
	c.ConnectDeleteGame(c.deleteGame)
}

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

func (c *ConfigBridge) addGame() {
	c.config.AddGame()
}

func (c *ConfigBridge) deleteGame(id int) {
	err := c.config.DeleteGame(id)
	if err != nil {
		c.logger.Error(err)
	}
}

// NewConfig ...
func NewConfig(cs config.Service, gm *config.GameModel, logger log.Logger) *ConfigBridge {
	configBridge := NewConfigBridge(nil)

	// Setup dependencies.
	configBridge.config = cs
	configBridge.logger = logger

	// Setup model.
	configBridge.SetGames(gm)

	return configBridge
}

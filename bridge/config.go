package bridge

import (
	"encoding/json"
	"fmt"

	"github.com/nokka/slashdiablo-launcher/config"
	"github.com/therecipe/qt/core"
)

// ConfigBridge is the connection between QML and the Go config.
type ConfigBridge struct {
	core.QObject

	// Services.
	Configuration config.Service

	// Models.
	GameModel *core.QAbstractListModel `property:"games"`

	// Functions.
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
		fmt.Println(err)
		// TODO: Add logger for the error.
		return false
	}

	err := c.Configuration.UpsertGame(request)
	if err != nil {
		// TODO: Add logger for the error.
		return false
	}

	return true
}

func (c *ConfigBridge) addGame() {
	c.Configuration.AddGame()
}

func (c *ConfigBridge) deleteGame(id int) {
	c.Configuration.DeleteGame(id)
}

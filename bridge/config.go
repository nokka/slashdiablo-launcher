package bridge

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

// ConfigBridge is the connection between QML and the Go config.
type ConfigBridge struct {
	core.QObject

	_ string `property:"D2Location"`
	_ int    `property:"D2Instances"`
	_ string `property:"HDLocation"`
	_ int    `property:"HDInstances"`

	_ func(D2Location string, HDLocation string) bool `slot:"setGamePaths"`
}

// Connect will connect the QML signals to functions in Go.
func (c *ConfigBridge) Connect() {
	c.ConnectSetGamePaths(c.setGamePaths)
}

func (c *ConfigBridge) setGamePaths(D2Location string, HDLocation string) bool {
	fmt.Println("SET GAME PATHS RUNNING")
	c.SetD2Location(D2Location)
	return true
}

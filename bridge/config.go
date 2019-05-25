package bridge

import (
	"github.com/nokka/slash-launcher/config"
	"github.com/therecipe/qt/core"
)

// ConfigBridge is the connection between QML and the Go config.
type ConfigBridge struct {
	core.QObject
	Configuration config.Service

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
	if err := c.Configuration.Update(config.UpdateConfigRequest{
		D2Location: &D2Location,
		HDLocation: &HDLocation,
	}); err != nil {
		return false
	}

	// Update was successful, update QML.
	c.SetD2Location(D2Location)
	c.SetHDLocation(HDLocation)

	return true
}

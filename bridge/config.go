package bridge

import (
	"github.com/nokka/slash-launcher/config"
	"github.com/therecipe/qt/core"
)

// ConfigBridge is the connection between QML and the Go config.
type ConfigBridge struct {
	core.QObject

	// Services.
	Configuration config.Service

	// Properties.
	_ string `property:"D2Location"`
	_ int    `property:"D2Instances"`
	_ bool   `property:"D2Maphack"`
	_ string `property:"HDLocation"`
	_ int    `property:"HDInstances"`
	_ bool   `property:"HDMaphack"`

	// Functions.
	_ func(D2Location string, D2Instances int, D2Maphack bool, HDLocation string, HDInstances int, HDMaphack bool) bool `slot:"update"`
}

// Connect will connect the QML signals to functions in Go.
func (c *ConfigBridge) Connect() {
	c.ConnectUpdate(c.update)
}

func (c *ConfigBridge) update(
	D2Location string,
	D2Instances int,
	D2Maphack bool,
	HDLocation string,
	HDInstances int,
	HDMaphack bool,
) bool {
	// Save updates to the persistant storage.
	if err := c.Configuration.Update(config.UpdateConfigRequest{
		D2Location:  &D2Location,
		D2Instances: &D2Instances,
		D2Maphack:   &D2Maphack,
		HDLocation:  &HDLocation,
		HDInstances: &HDInstances,
		HDMaphack:   &HDMaphack,
	}); err != nil {
		return false
	}

	// Update was successful, update QML.
	c.SetD2Location(D2Location)
	c.SetD2Instances(D2Instances)
	c.SetD2Maphack(D2Maphack)
	c.SetHDLocation(HDLocation)
	c.SetHDInstances(HDInstances)
	c.SetHDMaphack(HDMaphack)

	return true
}

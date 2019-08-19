package config

import (
	"github.com/therecipe/qt/core"
)

// Game represents a diablo installation in the configuration.
type Game struct {
	core.QObject

	ID            int    `json:"id"`
	Location      string `json:"location"`
	Instances     int    `json:"instances"`
	Maphack       bool   `json:"maphack"`
	OverrideBHCfg bool   `json:"override_bh_cfg"`
	HD            bool   `json:"hd"`
}

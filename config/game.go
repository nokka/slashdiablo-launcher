package config

import (
	"github.com/therecipe/qt/core"
)

const (
	// HDVersionNone is used to determine that no hd version has been chosen for a game.
	HDVersionNone = "none"
)

// Game represents a diablo installation in the configuration.
type Game struct {
	core.QObject

	ID            string   `json:"id"`
	Location      string   `json:"location"`
	Instances     int      `json:"instances"`
	Maphack       bool     `json:"maphack"`
	OverrideBHCfg bool     `json:"override_bh_cfg"`
	HD            bool     `json:"hd"`
	Flags         []string `json:"flags"`
	HDVersion     string   `json:"hd_version"`
}

// GameMods represents the mods available for a Diablo II game.
type GameMods struct {
	HD []string `json:"hd"`
}

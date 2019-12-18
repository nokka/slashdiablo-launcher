package config

import (
	"github.com/therecipe/qt/core"
)

const (
	// ModVersionNone is used to determine that no mod version has been chosen for a game.
	ModVersionNone = "none"
)

// Game represents a diablo installation in the configuration.
type Game struct {
	core.QObject

	ID             string   `json:"id"`
	Location       string   `json:"location"`
	Instances      int      `json:"instances"`
	OverrideBHCfg  bool     `json:"override_bh_cfg"`
	Flags          []string `json:"flags"`
	HDVersion      string   `json:"hd_version"`
	MaphackVersion string   `json:"maphack_version"`
}

// GameMods represents the mods available for a Diablo II game.
type GameMods struct {
	HD      []string `json:"hd"`
	Maphack []string `json:"maphack"`
}

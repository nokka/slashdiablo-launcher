package storage

// DefaultLaunchDelay is used if a launch delay hasn't been set by a user.
const DefaultLaunchDelay = 1000

// Config is the configuration required to run the app.
type Config struct {
	Games       []Game `json:"games"`
	LaunchDelay int    `json:"launch_delay"`
}

// Game represents a game setup by the user.
type Game struct {
	ID             string   `json:"id"`
	Location       string   `json:"location"`
	Instances      int      `json:"instances"`
	OverrideBHCfg  bool     `json:"override_bh_cfg"`
	Flags          []string `json:"flags"`
	HDVersion      string   `json:"hd_version"`
	MaphackVersion string   `json:"maphack_version"`
}

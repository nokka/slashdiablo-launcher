package storage

// Config is the configuration required to run the app.
type Config struct {
	Games   []Game `json:"games"`
	Gateway string `json:"gateway"`
}

// Game represents a game setup by the user.
type Game struct {
	ID            string   `json:"id"`
	Location      string   `json:"location"`
	Instances     int      `json:"instances"`
	Maphack       bool     `json:"maphack"`
	OverrideBHCfg bool     `json:"override_bh_cfg"`
	HD            bool     `json:"hd"`
	Flags         []string `json:"flags"`
	HDVersion     string   `json:"hd_version"`
}

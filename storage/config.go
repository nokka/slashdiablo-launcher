package storage

// Config is the configuration required to run the app.
type Config struct {
	Games []Game `json:"games"`
}

// Game represents a game setup by the user.
type Game struct {
	ID        int    `json:"id"`
	Location  string `json:"location"`
	Instances int    `json:"instances"`
	Maphack   bool   `json:"maphack"`
	HD        bool   `json:"hd"`
}

package storage

// Config is the configuration required to run the app.
type Config struct {
	D2Location  string `json:"d2_location"`
	D2Instances int    `json:"d2_instances"`
	D2Maphack   bool   `json:"d2_maphack"`
	HDLocation  string `json:"hd_location"`
	HDInstances int    `json:"hd_instances"`
	HDMaphack   bool   `json:"hd_maphack"`
	Games       []Game `json:"games"`
}

// Game represents a game setup by the user.
type Game struct {
	ID        int    `json:"id"`
	Location  string `json:"location"`
	Instances int    `json:"instances"`
	Maphack   bool   `json:"maphack"`
	HD        bool   `json:"hd"`
}

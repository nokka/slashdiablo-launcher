package storage

// Config is the configuration required to run the app.
type Config struct {
	D2Location  string `json:"d2_location"`
	D2Instances int    `json:"d2_instances"`
	HDLocation  string `json:"hd_location"`
	HDInstances int    `json:"hd_instances"`
}

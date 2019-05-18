package storage

// Config is the configuration required to run the app.
type Config struct {
	Location    *string `json:"location"`
	D2Instances int     `json:"d2_instances"`
	HDInstances int     `json:"hd_instances"`
}

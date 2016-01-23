package config

import "os"

// Configuration contains the configuration
type Configuration struct {
	Endpoint string `json:"cluster_endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Load builds a config obj
func Load(cf string) (*Configuration, error) {
	return &Configuration{
		Endpoint: os.Getenv("AION_ENDPOINT"),
		Username: os.Getenv("AION_USERNAME"),
		Password: os.Getenv("AION_PASSWORD"),
	}, nil
}

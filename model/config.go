package model

import "os"

type Config struct {
	Port string
}

func (config *Config) Load() error {
	port := os.Getenv("PORT")
	if port == "" {
		config.Port = "8080"
	} else {
		config.Port = port
	}
	return nil
}

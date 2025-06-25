package config

import "os"

// Config holds service configuration values.
type Config struct {
	Port string
}

// LoadConfig reads environment variables and returns a Config.
func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else if port[0] != ':' {
		port = ":" + port
	}
	return Config{Port: port}
}

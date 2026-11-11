package config

import (
	"os"
	"time"
)

// Config holds configuration values read from environment variables.
type Config struct {
	Port              string
	DBDriver          string
	PostgresDSN       string
	ConsensusInterval time.Duration
	LogLevel          string
}

// Load environment variables with sensible defaults.
func Load() *Config {
	port := getEnv("PORT", "8080")
	dbDriver := getEnv("DB_DRIVER", "sqlite")
	pg := getEnv("POSTGRES_DSN", "")
	ci := getEnv("CONSENSUS_INTERVAL", "10")
	logLevel := getEnv("LOG_LEVEL", "info")

	interval, err := time.ParseDuration(ci + "s")
	if err != nil {
		interval = 10 * time.Second
	}

	return &Config{
		Port:              port,
		DBDriver:          dbDriver,
		PostgresDSN:       pg,
		ConsensusInterval: interval,
		LogLevel:          logLevel,
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

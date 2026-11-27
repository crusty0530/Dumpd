package config

import (
	"fmt"
	"strings"
)

// Constant Values
const (
	DefaultPort    = "8080"
	envDatabaseURL = "DATABASE_URL"
	envPort        = "PORT"
)

// Define environment variables
type Config struct {
	Port        string
	DatabaseURL string
}

// InvalidConfigError lists every missing environment variable.
type InvalidConfigError struct {
	MissingKeys []string
}

// Error makes *InvalidConfigError satisfy the error interface.
func (e *InvalidConfigError) Error() string {
	return fmt.Sprintf("missing required environment variables: %s", strings.Join(e.MissingKeys, ", "))
}

// Load reads the environment variables
func Load(lookup func(string) string) (Config, error) {
	config := Config{Port: lookup(envPort), DatabaseURL: lookup(envDatabaseURL)}
	var missingKeys []string

	if config.Port == "" {
		config.Port = DefaultPort
	}

	if config.DatabaseURL == "" {
		missingKeys = append(missingKeys, envDatabaseURL)
	}

	if len(missingKeys) != 0 {
		return Config{}, &InvalidConfigError{MissingKeys: missingKeys}
	}

	return config, nil
}

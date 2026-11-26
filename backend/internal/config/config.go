package config

import (
	"fmt"
	"strings"
)

const DefaultPort = "8080"

// Define environment variables
type Config struct {
	Port string
}

// InvalidConfigError lists every missing environment variables.
type InvalidConfigError struct {
	MissingKeys []string
}

// Error makes *InvalidConfigError satisfy the error interface.
func (e *InvalidConfigError) Error() string {
	return fmt.Sprintf("missing required environment variables: %s", strings.Join(e.MissingKeys, ", "))
}

// Load the environment variables
func Load(lookup func(string) string) (Config, error) {
	config := Config{Port: lookup("PORT")}

	if config.Port == "" {
		config.Port = DefaultPort
	}

	return config, nil
}

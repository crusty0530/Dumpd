package config

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		env         map[string]string
		wantPort    string
		wantMissing []string
	}{
		{name: "PORT not set", env: map[string]string{envDatabaseURL: "postgres://test"}, wantPort: DefaultPort},
		{name: "PORT set to 3000", env: map[string]string{envPort: "3000", envDatabaseURL: "postgres://test"}, wantPort: "3000"},
		{name: "DATABASE_URL missing", env: map[string]string{envPort: DefaultPort}, wantMissing: []string{envDatabaseURL}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lookup := func(key string) string { return tc.env[key] }
			cfg, err := Load(lookup)

			if tc.wantMissing == nil {
				if err != nil {
					t.Fatalf("Load returned an unexpected error: %v", err)
				}
				if cfg.Port != tc.wantPort {
					t.Errorf("port: got %q, want %q", cfg.Port, tc.wantPort)
				}
				return
			}

			var configErr *InvalidConfigError
			if !errors.As(err, &configErr) {
				t.Fatalf("expected *InvalidConfigError, got %T (%v)", err, err)
			}

			if !slices.Equal(configErr.MissingKeys, tc.wantMissing) {
				t.Errorf("missing keys: got %v, want %v", configErr.MissingKeys, tc.wantMissing)
			}

			message := configErr.Error()
			for _, key := range tc.wantMissing {
				if !strings.Contains(message, key) {
					t.Errorf("error message %q does not mention %q", message, key)
				}
			}
		})
	}
}

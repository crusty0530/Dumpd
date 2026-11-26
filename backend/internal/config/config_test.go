package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		wantPort string
	}{
		{name: "PORT not set", env: map[string]string{}, wantPort: DefaultPort},
		{name: "PORT set to 3000", env: map[string]string{"PORT": "3000"}, wantPort: "3000"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lookup := func(key string) string { return tc.env[key] }
			cfg, err := Load(lookup)

			if err != nil {
				t.Fatalf("Load returned an unexpected error: %v", err)
			}

			if cfg.Port != tc.wantPort {
				t.Errorf("got %q, want %q", cfg.Port, tc.wantPort)
			}
		})
	}
}

package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	originalPort := os.Getenv("PORT")
	defer os.Setenv("PORT", originalPort)

	tests := []struct {
		name         string
		envPort      string
		expectedPort string
	}{
		{
			name:         "Environment variable PORT is set",
			envPort:      "8080",
			expectedPort: "8080",
		},
		{
			name:         "Environment variables are not set so it takes default values",
			envPort:      "",
			expectedPort: defaultPort,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envPort != "" {
				os.Setenv("PORT", tt.envPort)
			}
			cfg := LoadConfig()

			if cfg.Port != tt.expectedPort {
				t.Errorf("got %s, want %s", cfg.Port, tt.expectedPort)
			}
		})
	}
}

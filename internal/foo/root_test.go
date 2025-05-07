package foo

import (
	"testing"

	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "basic successful run",
			config: Config{
				Verbose: true,
				Name:    "test-service",
				Options: viper.New(),
			},
			wantErr: false,
		},
		{
			name: "run with empty name",
			config: Config{
				Verbose: false,
				Name:    "",
				Options: viper.New(),
			},
			wantErr: false,
		},
		{
			name: "run with nil options",
			config: Config{
				Verbose: true,
				Name:    "test-service",
				Options: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig(t *testing.T) {
	v := viper.New()
	config := Config{
		Verbose: true,
		Name:    "test-name",
		Options: v,
	}

	if !config.Verbose {
		t.Error("Config.Verbose = false, want true")
	}

	if got := config.Name; got != "test-name" {
		t.Errorf("Config.Name = %v, want %v", got, "test-name")
	}

	if got := config.Options; got != v {
		t.Errorf("Config.Options = %v, want %v", got, v)
	}
}

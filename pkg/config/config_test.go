package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	os.Args = []string{"test", "--data-path", "test", "--port", "8081", "--host", "localhost"}

	t.Run("Default", func(t *testing.T) {
		gotConfig, err := NewConfig()
		if err != nil {
			t.Errorf("NewConfig() error = %v", err)
		}
		if gotConfig.DataPath != "test" {
			t.Errorf("NewConfig() = %v, want %v", gotConfig.DataPath, "test")
		}
		if gotConfig.Port != 8081 {
			t.Errorf("NewConfig() = %v, want %v", gotConfig.Port, 8081)
		}
		if gotConfig.Host != "localhost" {
			t.Errorf("NewConfig() = %v, want %v", gotConfig.Host, "localhost")
		}
	})

}

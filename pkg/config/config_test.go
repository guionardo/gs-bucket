package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	os.Args = []string{"test", "--data-path", "test", "--port", "8081", "--host", "localhost"}

	t.Run("Default", func(t *testing.T) {
		gotConfig, err := NewConfig()
		if err != nil {
			t.Errorf("NewConfig() error = %v", err)
		}
		expected := &Config{
			DataPath: "test",
			Port:     8081,
			Host:     "localhost",
		}
		if !reflect.DeepEqual(gotConfig, expected) {
			t.Errorf("NewConfig() = %v, want %v", gotConfig, expected)
		}

	})

}

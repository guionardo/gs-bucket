package server

import (
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tmp := t.TempDir()
	tests := []struct {
		name    string
		wantCfg *Config
		wantErr bool
		preRun  func()
	}{
		{"No envs", nil, true, nil},
		{"Default port", &Config{
			RepositoryFolder: tmp,
			HttpPort:         8080,
		}, false, func() {
			os.Setenv(RepositoryFolder, tmp)
		}},
		{"Especific port", &Config{
			RepositoryFolder: tmp,
			HttpPort:         1234,
		}, false, func() {
			os.Setenv(HttpPort, "1234")
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.preRun != nil {
				tt.preRun()
			}
			gotCfg, err := GetConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("GetConfig() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

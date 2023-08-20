package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/guionardo/gs-bucket/domain"
)

func TestLocalRepository_save_and_load_data(t *testing.T) {
	tmp := t.TempDir()
	repo, _ := CreateLocalRepository(tmp)
	hash := ""
	data := []byte(`{"name":"Guionardo"}`)
	t.Run("Save", func(t *testing.T) {
		file1, _ := domain.CreateFileFromData("test1.json", data, 10*time.Minute, "owner")
		_ = repo.SaveFile(file1, data)
		if err := repo.saveFileMap(); err != nil {
			t.Errorf("LocalRepository.saveFileMap() error = %v", err)
		}
		hash = file1.Slug
	})
	t.Run("Load", func(t *testing.T) {
		if err := repo.loadFileMap(); err != nil {
			t.Errorf("LocalRepository.loadFileMap() error = %v", err)
		}
		file := repo.Get(hash)
		if file == nil || file.Name != "test1.json" {
			t.Error("LocalRepository.Get() failed to get file")
		}
	})
	t.Run("Load file", func(t *testing.T) {
		if readData, err := repo.Load(hash); err != nil {
			t.Errorf("LocalRepository.Load error = %v", err)
		} else if !reflect.DeepEqual(readData, data) {
			t.Errorf("Expected data %v got %v", data, readData)
		}

	})

}

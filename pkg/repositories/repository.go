package repositories

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/guionardo/gs-bucket/pkg/config"
)

type Repository struct {
	DataPath string
}

func NewRepository(c *config.Config) (repository *Repository, err error) {
	if len(c.DataPath) == 0 {
		err = errors.New("data path is empty")
		return
	}

	stat, err := os.Stat(c.DataPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(c.DataPath, os.ModePerm)
		if err != nil {
			return
		}
	}
	if stat, err = os.Stat(c.DataPath); err != nil || !stat.IsDir() {
		err = errors.New("data path is not a directory")
		return
	}

	return &Repository{DataPath: c.DataPath}, nil
}

func (repository *Repository) GetFile(fileName string) (file *File, err error) {
	localFile := path.Join(repository.DataPath, fileName)
	if stat, err := os.Stat(localFile); err != nil || stat.IsDir() {
		return nil, fmt.Errorf("file not found %s", fileName)
	}

	return ReadFile(localFile)

}

func (repository *Repository) SaveFile(file *File) error {
	localFile := path.Join(repository.DataPath, file.Code)
	return file.Save(localFile)
}

func (repository *Repository) DeleteFile(file *File) {
	localFile := path.Join(repository.DataPath, file.Code)
	file.Delete(localFile)
}

func (repository *Repository) GetAll() (files []*File, err error) {

	fi, err := ioutil.ReadDir(repository.DataPath)
	if err != nil {
		return
	}
	files = make([]*File, 0, len(fi))
	for _, f := range fi {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".meta") {
			continue
		}
		filename := strings.TrimSuffix(f.Name(), ".meta")
		if file, err := repository.GetFile(filename); err == nil {
			if file.ValidUntil.Before(time.Now()) {
				repository.DeleteFile(file)
			} else {
				file.Content = nil
				files = append(files, file)
			}
		}

	}
	return
}

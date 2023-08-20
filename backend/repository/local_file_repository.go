package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/guionardo/gs-bucket/domain"
)

type LocalRepository struct {
	rootPath  string
	fileMap   map[string]*domain.File
	indexFile string

	lock sync.RWMutex
}

func (r *LocalRepository) loadFileMap() error {
	content, err := os.ReadFile(r.indexFile)
	if err == nil {
		err = json.Unmarshal(content, &r.fileMap)
	}
	if err != nil {
		log.Printf("LocalRepository.loadFileMap error %v", err)
	}
	return err
}

func (r *LocalRepository) saveFileMap() error {
	content, err := json.Marshal(r.fileMap)
	if err == nil {
		err = os.WriteFile(r.indexFile, content, 0666)
	}
	if err != nil {
		log.Printf("LocalRepository.saveFileMap error %v", err)
	}

	return err
}

func (r *LocalRepository) Save() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.saveFileMap()
}

func (r *LocalRepository) Purge() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.loadFileMap()
	removed := 0
	for key, file := range r.fileMap {
		if file.Expired() {
			delete(r.fileMap, key)
			os.Remove(path.Join(r.rootPath, file.Name))
			removed++
		}
	}
	if removed > 0 {
		r.saveFileMap()
	}
}

func CreateLocalRepository(rootPath string) (*LocalRepository, error) {
	if _, err := os.Stat(rootPath); err != nil {
		return nil, err
	}
	log.Printf("Local repository: %s", rootPath)
	return &LocalRepository{
		rootPath:  rootPath,
		indexFile: path.Join(rootPath, ".index"),
		fileMap:   make(map[string]*domain.File),
	}, nil
}

func (r *LocalRepository) Get(hash string) *domain.File {
	return r.fileMap[hash]
}

func (r *LocalRepository) SaveFile(file *domain.File, data []byte) (err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if file.CreationDate.IsZero() {
		file.CreationDate = time.Now()
	}
	if err = os.WriteFile(path.Join(r.rootPath, file.Name), data, 0644); err != nil {
		return err
	}
	r.fileMap[file.Slug] = file
	return r.saveFileMap()
}

func (r *LocalRepository) Load(hash string) (data []byte, err error) {
	if file, ok := r.fileMap[hash]; !ok {
		err = fmt.Errorf("file not found %s", hash)
	} else if file.Expired() {
		err = fmt.Errorf("file expired %s", hash)
	} else {
		fileName := path.Join(r.rootPath, file.Name)
		data, err = os.ReadFile(fileName)
	}
	return
}

func (r *LocalRepository) Delete(hash string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if file, ok := r.fileMap[hash]; ok {
		delete(r.fileMap, hash)
		os.Remove(path.Join(r.rootPath, file.Name))
		return r.saveFileMap()
	}
	return fmt.Errorf("file not found in repository: %s", hash)
}

func (r *LocalRepository) GetFileCount() int {
	return len(r.fileMap)
}

func (r *LocalRepository) GetFileSize() int64 {
	var size int64
	for _, file := range r.fileMap {
		size += int64(file.Size)
	}
	return size
}

func (r *LocalRepository) List() ([]*domain.File, error) {
	list := make([]*domain.File, len(r.fileMap))
	index := 0
	for _, file := range r.fileMap {
		list[index] = file
		index++
	}
	return list, nil
}

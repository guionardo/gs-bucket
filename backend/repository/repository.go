package repository

import "github.com/guionardo/gs-bucket/domain"

type Repository interface {
	SaveFile(file *domain.File, data []byte) error
	Save() error
	Load(hash string) ([]byte, error)
	Get(hash string) *domain.File
	Delete(hash string) error
	List() ([]*domain.File, error)
	Purge()

	GetFileCount() int
	GetFileSize() int64
}

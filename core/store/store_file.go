package store

import (
	"app/models"
	"os"

	"gorm.io/gorm"
)

type FileStore struct {
	db *gorm.DB
}

func NewFileStore(db *gorm.DB) *FileStore {
	return &FileStore{
		db: db,
	}
}

func (s *FileStore) GetByID(id int64) (*models.File, error) {
	var ret models.File
	err := s.db.Take(&ret, "id = ?", id).Error
	return &ret, err
}

func (s *FileStore) Create(f *models.File) error {
	return s.db.Create(f).Error
}

func (s *FileStore) DeleteByID(id int64) error {
	var f models.File
	if err := s.db.Take(&f, "id = ?", id).Error; err != nil {
		return err
	}
	if err := s.db.Delete(&f).Error; err != nil {
		return err
	}
	return os.Remove(f.Path)
}

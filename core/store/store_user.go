package store

import (
	"app/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) SearchActives(username string, exclude string) ([]models.User, error) {
	var ret []models.User
	if err := s.db.Where("username LIKE ? AND is_active=? AND username != ?", "%"+username+"%", true, exclude).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *UserStore) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetActivesByUsername(usernames []string) ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users, "username IN (?) AND is_active=?", usernames, true).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}

func (s *UserStore) Create(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.db.Create(user).Error
}

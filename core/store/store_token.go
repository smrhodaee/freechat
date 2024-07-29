package store

import (
	"app/models"
	"app/tools"
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TokenStore struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewTokenStore(db *gorm.DB, rdb *redis.Client) *TokenStore {
	return &TokenStore{
		db:  db,
		rdb: rdb,
	}
}

func (s *TokenStore) GenerateRandom(username string, length int, expire time.Duration) (*models.Token, error) {
	var token models.Token
	token.Username = username
	token.Value = tools.RandomToken(length)
	err := s.rdb.Set(context.Background(), token.Key(), token.Username, expire).Err()
	return &token, err
}

func (s *TokenStore) GetUser(value string) (*models.User, error) {
	var user models.User
	var err error
	token := models.Token{
		Value: value,
	}
	token.Username, err = s.rdb.Get(context.Background(), token.Key()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	if err := s.db.Where("username=?", token.Username).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *TokenStore) Delete(token string) error {
	return s.rdb.Del(context.Background(), models.Token{
		Value: token,
	}.Key()).Err()
}

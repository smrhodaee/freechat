package store

import (
	"app/configs"
	"app/models"
	"app/tools"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type ChaptaStore struct {
	rdb *redis.Client
}

func NewChaptaStore(rdb *redis.Client) *ChaptaStore {
	return &ChaptaStore{
		rdb: rdb,
	}
}

func (s *ChaptaStore) Delete(uuid string) error {
	return s.rdb.Del(context.Background(), configs.GetChaptaKey(uuid)).Err()
}

func (s *ChaptaStore) GenerateRandom(length int, expire time.Duration) (*models.Chapta, error) {
	var chapta models.Chapta
	var err error
	chapta.UUID, err = tools.GenerateUUID()
	if err != nil {
		return nil, err
	}
	chapta.Code = tools.RandomChapta(length)
	ctx := context.Background()
	if err := s.rdb.Set(ctx, configs.GetChaptaKey(chapta.UUID), chapta.Code, expire).Err(); err != nil {
		return nil, err
	}
	return &chapta, nil
}

func (s *ChaptaStore) IsValid(chapta *models.Chapta) (bool, error) {
	code, err := s.rdb.Get(context.Background(), configs.GetChaptaKey(chapta.UUID)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return strings.Compare(code, chapta.Code) == 0, nil
}

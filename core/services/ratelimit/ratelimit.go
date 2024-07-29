package ratelimit

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	countKey    = "count:%s:%s:%d" // "count:<rule-name>:<value>:<type>"
	exceededKey = "exceeded:%s:%s" // "exceededKey:<rule-name>:<value>"
)

type Item struct {
	Count  int64         `json:"count"`
	Expire time.Duration `json:"expire"`
}

type Rule struct {
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

type Service struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Service {
	return &Service{
		rdb: rdb,
	}
}

func (s *Service) Incr(rule *Rule, value string) error {
	ctx := context.Background()
	exceededK := fmt.Sprintf(exceededKey, rule.Name, value)
	for i, item := range rule.Items {
		countK := fmt.Sprintf(countKey, rule.Name, value, i)
		value, err := s.rdb.Incr(ctx, countK).Result()
		if err != nil {
			return err
		}
		if value == 1 {
			if err := s.rdb.Set(ctx, countK, 1, item.Expire).Err(); err != nil {
				return err
			}
		}
		if value > item.Count {
			ttl, err := s.rdb.TTL(ctx, countK).Result()
			if err != nil {
				return err
			}
			if ttl >= time.Second {
				if err := s.rdb.Set(ctx, exceededK, 1, ttl).Err(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *Service) IsExceeded(rule *Rule, value string) (bool, error) {
	key := fmt.Sprintf(exceededKey, rule.Name, value)
	res, err := s.rdb.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

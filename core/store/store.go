package store

import (
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Store struct {
	User    *UserStore
	Chapta  *ChaptaStore
	Token   *TokenStore
	Room    *RoomStore
	Message *MessageStore
	File    *FileStore
}

func New(db *gorm.DB, rdb *redis.Client, cache *cache.Cache) *Store {
	return &Store{
		User:    NewUserStore(db),
		Chapta:  NewChaptaStore(rdb),
		Token:   NewTokenStore(db, rdb),
		Room:    NewRoomStore(db),
		Message: NewMessageStore(db),
		File:    NewFileStore(db),
	}
}
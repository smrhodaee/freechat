package models

import (
	"time"
)

type MessageType uint8

const (
	MessageTypeText MessageType = iota
	MessageTypeImage
)

type Message struct {
	ID        int64       `gorm:"primary_key" json:"id"`
	RoomName  string      `json:"room_name"`
	Username  string      `json:"username"`
	Text      string      `json:"Text"`
	ValueID   int64       `json:"value_id"`
	Type      MessageType `json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	Room      *Room       `gorm:"foreignKey:RoomName"`
}

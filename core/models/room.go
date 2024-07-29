package models

import "time"

type RoomType uint8

const (
	RoomTypeStorage RoomType = iota
	RoomTypeDirect
	RoomTypeGroup
	RoomTypeChannel
)

type Room struct {
	Name      string      `gorm:"primary_key" json:"name"`
	Title     string      `json:"title"`
	Type      RoomType    `json:"type"`
	IsActive  bool        `json:"is_active"`
	CreatedAt time.Time   `json:"created_at"`
	Members   RoomMembers `gorm:"foreignKey:RoomName" json:"members"`
}

func (Room) Table() string {
	return "room"
}

type MemberRule uint8

const (
	MemberRoleNormal MemberRule = iota
	MemberRoleAdmin
	MemberRoleOwner
)

type RoomMember struct {
	ID             int64      `gorm:"primary_key" json:"id"`
	RoomName       string     `json:"room_name"`
	Username       string     `json:"username"`
	Role           MemberRule `json:"role"`
	CreatedAt      time.Time  `json:"created_at"`
	LastActivityAt time.Time  `json:"last_activity_at"`
}

func (RoomMember) Table() string {
	return "room_member"
}

type RoomMembers []RoomMember

func (rms RoomMembers) GetUsernames() []string {
	var ret []string
	for _, m := range rms {
		ret = append(ret, m.Username)
	}
	return ret
}

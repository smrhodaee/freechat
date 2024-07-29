package api

import (
	"app/configs"
	"app/models"
	"app/router"
	"time"
)

type UserResponse struct {
	Username string `json:"username"`
}

func (UserResponse) Clone(data any) UserResponse {
	obj, ok := data.(models.User)
	var ret UserResponse
	if ok {
		ret.Username = obj.Username
	}
	return ret
}

//------------------------------------------------------

type MemberResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (MemberResponse) Clone(data any) MemberResponse {
	obj, ok := data.(models.RoomMember)
	var ret MemberResponse
	if ok {
		ret.Username = obj.Username
		ret.Role = configs.GetMemberRole(obj.Role)
	}
	return ret
}

type RoomResponse struct {
	Name      string           `json:"name"`
	Title     string           `json:"title"`
	Type      string           `json:"type"`
	CreatedAt time.Time        `json:"created_at"`
	Members   []MemberResponse `json:"members"`
}

func (RoomResponse) Clone(data any) RoomResponse {
	obj, ok := data.(models.Room)
	var ret RoomResponse
	if ok {
		ret.Name = obj.Name
		ret.Title = obj.Title
		switch obj.Type {
		case models.RoomTypeDirect:
			ret.Type = "direct"
		case models.RoomTypeChannel:
			ret.Type = "channel"
		case models.RoomTypeGroup:
			ret.Type = "group"
		case models.RoomTypeStorage:
			ret.Type = "storage"
		}
		ret.CreatedAt = obj.CreatedAt
		ret.Members = router.Collection(MemberResponse{}, obj.Members)
	}
	return ret
}

//------------------------------------------------------

type ChatResponse struct {
	ID        int64     `json:"id"`
	RoomName  string    `json:"room_name"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	ValueID   int64     `json:"value_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	TextType  = "TEXT"
	ImageType = "IMAGE"
)

func (ChatResponse) Clone(data any) ChatResponse {
	msg, ok := data.(models.Message)
	var ret ChatResponse
	if ok {
		ret.ID = msg.ID
		ret.RoomName = msg.RoomName
		ret.Username = msg.Username
		ret.Text = msg.Text
		if msg.Type == models.MessageTypeText {
			ret.Type = TextType
		} else if msg.Type == models.MessageTypeImage {
			ret.Type = ImageType
		}
		ret.ValueID = msg.ValueID
		ret.CreatedAt = msg.CreatedAt
	}
	return ret
}

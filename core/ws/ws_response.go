package ws

import (
	"app/configs"
	"app/models"
	"time"
)

type CreateMessageResponse struct {
	ID        int64     `json:"id"`
	RoomName  string    `json:"room_name"`
	Username  string    `json:"username"`
	ValueID   int64     `json:"value_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
}

func (CreateMessageResponse) Clone(data any) CreateMessageResponse {
	obj, ok := (data).(*models.Message)
	var ret CreateMessageResponse
	if ok {
		ret.ID = obj.ID
		ret.RoomName = obj.RoomName
		ret.Username = obj.Username
		ret.ValueID = obj.ValueID
		ret.Text = obj.Text
		ret.CreatedAt = obj.CreatedAt
		ret.Type = configs.GetMessageType(obj.Type)
	}
	return ret
}

func NewCreateMessageResponse(msg *models.Message) *WsMessage {
	return NewMessage(nil, createMessageType, CreateMessageResponse{}.Clone(msg))
}

type DeleteMessageResponse struct {
	ID int64 `json:"id"`
}

func NewDeleteMessageResponse(id int64) *WsMessage {
	return NewMessage(nil, deleteChatType, DeleteMessageResponse{ID: id})
}
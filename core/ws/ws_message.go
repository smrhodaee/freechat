package ws

type WsMessageType string

const (
	createMessageType WsMessageType = "CHAT"
	deleteChatType    WsMessageType = "DELETECHAT"
	errorMessageType  WsMessageType = "ERROR"
)

type WsMessageBody struct {
	Type  WsMessageType `json:"type"`
	Value interface{}   `json:"value"`
}

type WsMessage struct {
	Client *WsClient     `json:"client"`
	Body   WsMessageBody `json:"body"`
}

func NewMessage(client *WsClient, t WsMessageType, value interface{}) *WsMessage {
	return &WsMessage{
		Client: client,
		Body: WsMessageBody{
			Type:  t,
			Value: value,
		},
	}
}

// func (m *WsMessage) GetChatMessage(s *store.Store) (*ChatMessage, error) {
// 	var msg ChatMessage
// 	data, err := json.Marshal(m.Body.Value)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := json.Unmarshal(data, &msg); err != nil {
// 		return nil, err
// 	}
// 	if err := msg.ValidateAndNormalize(s, m.Client.Connection); err != nil {
// 		return nil, err
// 	}
// 	return &msg, nil
// }

// type DeleteChatSchema struct {
// 	ID int64 `json:"id"`
// }

// type ChatMessage struct {
// 	ID        int64        `json:"id"`
// 	RoomName  string       `json:"room_name"`
// 	Username  string       `json:"username"`
// 	ValueID   int64        `json:"value_id"`
// 	Text      string       `json:"text"`
// 	CreatedAt time.Time    `json:"created_at"`
// 	Type      string       `json:"type"`
// 	Room      *models.Room `json:"-"`
// }

// func (m *ChatMessage) ValidateAndNormalize(s *store.Store, c *websocket.Conn) error {
// 	msg, err := s.Message.GetByID(m.ID)
// 	if err != nil {
// 		return err
// 	}
// 	if msg == nil {
// 		return router.NotFoundError()
// 	}
// 	m.Text = msg.Text
// 	m.RoomName = msg.RoomName
// 	m.Username = msg.Username
// 	m.CreatedAt = msg.CreatedAt
// 	m.ValueID = msg.ValueID
// 	m.Room = &msg.Room
// 	if !permissions.CanBroadCast(router.GetReqInfoWs(c).User, msg) {
// 		return router.PermissionError()
// 	}
// 	return nil
// }

// func NewDeleteChatMessage(id int64) *WsMessage {
// 	return NewMessage(nil, deleteChatType, DeleteChatSchema{
// 		ID: id,
// 	})
// }

// func NewChatMessage(client *WsClient, msg *ChatMessage) *WsMessage {
// 	return NewMessage(client, chatMessageType, msg)
// }

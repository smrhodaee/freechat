package ws

import (
	"github.com/gofiber/contrib/websocket"
)

type WsClient struct {
	Username   string
	Connection *websocket.Conn
}

func NewClient(username string, c *websocket.Conn) *WsClient {
	return &WsClient{
		Username:   username,
		Connection: c,
	}
}

func (wc *WsClient) SendMessage(msg *WsMessage) error {
	return wc.Connection.WriteJSON(msg.Body)
}

func (wc *WsClient) ReadMessage() (*WsMessage, error) {
	var msg = new(WsMessage)
	msg.Client = wc
	err := wc.Connection.ReadJSON(&msg.Body)
	return msg, err
}

func (wc *WsClient) SendError(message string) error {
	return wc.SendMessage(NewMessage(wc, errorMessageType, message))
}

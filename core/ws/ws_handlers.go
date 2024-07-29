package ws

import (
	"app/router"

	"github.com/gofiber/contrib/websocket"
)

func (ws *WS) chatHandler(c *websocket.Conn) {
	reqInfo := router.GetReqInfoWs(c)
	client := NewClient(reqInfo.Username, c)
	defer func(client *WsClient) {
		ws.unregister <- client
	}(client)
	ws.register <- client
	for {
		msg, err := client.ReadMessage()
		if err != nil {
			return
		}
		ws.message <- msg
	}
}

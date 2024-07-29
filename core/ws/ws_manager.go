package ws

func (ws *WS) Broadcast(message *WsMessage, usernames []string) {
	for _, username := range usernames {
		ws.lock.Lock()
		for _, client := range ws.sessions[username] {
			ws.handleConnectionError(client.SendMessage(message), client)
		}
		ws.lock.Unlock()
	}
}

func (ws *WS) handleRegister(client *WsClient) {
	ws.sessions[client.Username] = append(ws.sessions[client.Username], client)
}

func (ws *WS) handleUnRegister(client *WsClient) {
	for i := len(ws.sessions[client.Username]) - 1; i >= 0; i-- {
		if ws.sessions[client.Username][i] == client {
			ws.sessions[client.Username] = append(ws.sessions[client.Username][:i], ws.sessions[client.Username][i+1:]...)
		}
	}

	if len(ws.sessions[client.Username]) == 0 {
		delete(ws.sessions, client.Username)
	}

	if client.Connection != nil {
		client.Connection.Close()
	}
}

func (ws *WS) handleConnectionError(err error, client *WsClient) {
	if err != nil {
		ws.log.Error("websocket connection:", err)
		ws.unregister <- client
	}
}

// func (ws *WS) handleMessage(message *WsMessage) error {
// 	switch message.Body.Type {
// 	case chatMessageType:
// 		msg, err := message.GetChatMessage(ws.store)
// 		if err != nil {
// 			return err
// 		}
// 		ws.Broadcast(NewChatMessage(message.Client, msg), msg.Room.Members.GetUsernames())
// 	}
// 	return nil
// }

// func (ws *WS) handleError(err error, client *WsClient) {
// 	if err != nil {
// 		fErr := router.ParseError(err)
// 		if fErr.Code == fiber.StatusInternalServerError {
// 			ws.log.Error("websocket:", err)
// 		}
// 		ws.handleConnectionError(client.SendError(fErr.Message), client)
// 	}
// }

func (ws *WS) HandleChannels() {
	for {
		select {
		case client := <-ws.register:
			ws.handleRegister(client)
		case client := <-ws.unregister:
			ws.handleUnRegister(client)
			// case message := <-ws.message:
			// 	ws.handleError(ws.handleMessage(message), message.Client)
		}
	}
}

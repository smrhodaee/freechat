import { getToken } from "./common"
import routes from "./routes"

type onChatMessageType = (msg: CreateMessageResponse) => void
type onChatDeleteType = (id: string) => void
type onErrorType = (msg: string) => void

export class WsChat {
    socket: WebSocket
    constructor(onChatMessage: onChatMessageType, onChatDelete: onChatDeleteType, onError: onErrorType) {
        this.socket = new WebSocket(routes.wsChats(getToken()))
        this.socket.addEventListener('message', e => {
            const data = JSON.parse(e.data)
            switch (data['type']) {
                case "ERROR":
                    onError(data['value'])
                    break;
                case "CHAT":
                    onChatMessage(data["value"])
                    break;
                case "DELETECHAT":
                    onChatDelete(data['value']["id"])
                    break;
            }
        })
        this.socket.addEventListener('open', () => console.log("connection is open"))
    }
}

export interface CreateMessageResponse {
    id: number
    room_name: string
    username: string
    text: string
    value_id: string
    type: string
    created_at: number
}

export default ""
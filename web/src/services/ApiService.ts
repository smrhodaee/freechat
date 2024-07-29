import { GlobalResponse, RequestInput, authHeader, extractGlobal, getRequest, jsonHeader, jsonRequest, request } from "./common"
import routes from "./routes"

class Service {
    async getChapta(): Promise<GlobalResponse<ChaptaResponse>> {
        return request<ChaptaResponse>({
            "url": routes.chapta(),
            "handle": async (resp: Response) => {
                if (resp.status == 200) {
                    let blob = await resp.blob()
                    let text = await blob.text()
                    let uuid = text.slice(text.indexOf("CHAPTA_UUID") + 12, text.length)
                    return new GlobalResponse<ChaptaResponse>("Chapta Created Successfully", 200, false, {
                        "uuid": uuid,
                        "url": window.URL.createObjectURL(blob),
                    })
                }
                return await extractGlobal(resp)
            }
        })
    }

    async registerOrLogin(data: RegisterOrLoginRequest): Promise<GlobalResponse<LoginOrRegisterResponse>> {
        return jsonRequest(routes.registerOrLogin(), data, false)
    }

    async findRoomsOfUser(): Promise<GlobalResponse<RoomResponse[]>> {
        return getRequest(routes.roomsOfUser())
    }

    async logout() {
        return getRequest(routes.logout())
    }

    async findChatsOfRoom(roomName: string): Promise<GlobalResponse<ChatResponse[]>> {
        return jsonRequest(routes.findChatsOfRoom(), { "room_name": roomName })
    }

    async findUsers(username: string): Promise<GlobalResponse<UserResponse[]>> {
        return jsonRequest(routes.findUsers(), { "username": username })
    }

    async createGroup(data: CreateGroupRequest) {
        return jsonRequest(routes.createGroup(), data)
    }

    async deleteChat(id: string) {
        return jsonRequest(routes.deleteChat(), { "id": id })
    }

    async getImageMessage(id: string): Promise<GlobalResponse<ImageMessage>> {
        return request({
            "url": routes.getImageMessage(),
            headers: { ...authHeader(), ...jsonHeader() },
            method: "POST",
            body: JSON.stringify({ "id": id }),
            "handle": async (resp: Response) => {
                if (resp.status == 200) {
                    return new GlobalResponse<ImageMessage>("Image get successfully", 200, false, {
                        "url": window.URL.createObjectURL(await resp.blob()),
                    })
                }
                return await extractGlobal(resp);
            }
        })
    }

    async submitChat(req: SubmitMessageRequest, file?: File): Promise<GlobalResponse<ChatResponse>> {
        const data = new FormData()
        data.append("room_name", req.room_name)
        data.append("text", req.text)
        if (file) data.append("file", file)
        return request(new RequestInput(routes.submitChat(), "POST", { ...authHeader() }, data))
    }
}

export const APIService = new Service()


export interface CreateGroupRequest {
    name: string
    title: string
    usernames: string[]
}

export interface SubmitMessageRequest {
    room_name: string
    text: string
}

export interface RegisterOrLoginRequest {
    username: string
    password: string
    uuid: string
    code: string
}


export interface ChatResponse {
    id: number
    room_name: string
    username: string
    text: string
    value_id: string
    type: string
    created_at: number
}

export interface RoomResponse {
    name: string
    title: string
    type: string
    created_at: number
}

export interface ChaptaResponse {
    uuid: string
    url: string
}

export interface ImageMessage {
    url: string
}

export interface LoginOrRegisterResponse {
    token: string
}

export interface UserResponse {
    username: string 
}

export default ""
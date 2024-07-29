const API_URL = `http://${window.location.hostname}:8080`
const WS_URL = `ws://${window.location.hostname}:8080`

export default {
    chapta: () => `${API_URL}/chapta.png`,
    registerOrLogin: () => `${API_URL}/auth/login`,
    createGroup: () => `${API_URL}/rooms/group`,
    roomsOfUser: () => `${API_URL}/users/rooms/`,
    logout: () => `${API_URL}/auth/logout`,
    findChatsOfRoom: () => `${API_URL}/rooms/msgs`,
    wsChats: (token: string) => `${WS_URL}/chat/?token=${token}`,
    deleteChat: ( )=> `${API_URL}/delete-message`,
    getImageMessage: () => `${API_URL}/get-image-chat`,
    submitChat: () => `${API_URL}/msgs/`,
    findUsers: () => `${API_URL}/users`
}

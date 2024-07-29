export default class Room {
    name: string
    title: string
    type: string
    unreadMessages: number

    constructor(name = "", title = "", type="", unreadMessages = 0) {
        this.name = name;
        this.title = title;
        this.type = type;
        this.unreadMessages = unreadMessages;
    }
}
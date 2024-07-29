export default class Chat {
    id: string
    username: string
    text: string
    url?: string
    value_id: string
    type: string
    isOwner: boolean

    constructor(id: string, username: string, text: string, value_id: string, type: string, isOwner: boolean) {
        this.id = id
        this.username = username
        this.text = text
        this.type = type
        this.isOwner = isOwner
        this.value_id = value_id
    }
}

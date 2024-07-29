export function serverUnavailable<T>(): GlobalResponse<T> {
    return {
        "status": 0,
        "error": true,
        "message": "Server Unavailable"
    }
}

export function getToken(): string {
    return localStorage.getItem("token") || ""
}

export function authHeader(needed = true) {
    return {
        "Authorization": needed ? getToken() : ""
    }
}

export function jsonHeader() {
    return {
        "Content-Type": "application/json",
    }
}

export async function jsonRequest<T>(url: string, data: any, auth: boolean = true): Promise<GlobalResponse<T>> {
    return request(new RequestInput(url, "POST", {
        ...authHeader(auth),
        ...jsonHeader(),
    }, JSON.stringify(data)))
}

export async function getRequest<T>(url: string): Promise<GlobalResponse<T>> {
    return request(new RequestInput(url, "GET", authHeader(true)))
}

export async function deleteRequest<T>(url: string): Promise<GlobalResponse<T>> {
    return request(new RequestInput(url, "DELETE", authHeader(true)))
}

export class RequestInput {
    url: string
    method?: string = "GET"
    headers?: HeadersInit
    body?: BodyInit
    handle: Function = extractGlobal
    constructor(url: string, method = "GET", headers?: HeadersInit, body?: BodyInit, handle?: Function) {
        this.url = url
        this.method = method
        this.headers = headers
        this.body = body
        if (handle) this.handle = handle
    }
}

export async function request<T>(req: RequestInput): Promise<GlobalResponse<T>> {
    try {
        const resp = await fetch(req.url, {
            method: req.method,
            headers: req.headers,
            body: req.body,
        })
        return await req.handle(resp)
    } catch (e) {
        return serverUnavailable()
    }
}

export async function extractGlobal<T>(response: Response): Promise<GlobalResponse<T>> {
    const json = await response.json()
    let ret: GlobalResponse<T> = {
        "status": response.status,
        "error": !json["status"],
        "message": json["message"],
    }
    if (json["data"]) ret["data"] = json["data"][0]
    return ret
}


export class GlobalResponse<T> {
    status: number
    error: boolean
    message: string
    data?: T
    constructor(message: string, status = 200, error = false, data?: T) {
        this.message = message
        this.status = status
        this.error = error
        this.data = data
    }
}
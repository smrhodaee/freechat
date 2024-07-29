export default
    class Timer {
    startTime!: number
    handler: (ms: number) => void
    interval!: number
    ms: number
    now!: number
    down!: boolean;

    constructor(ms: number, handler: (ms: number) => void) {
        this.ms = ms
        this.handler = handler
    }

    start() {
        this.now = this.ms
        this.down = false;
        this.resume()
    }

    resume() {
        if (!this.down) {
            this.startTime = (new Date()).getTime()
            this.interval = setInterval(() => this.tick(), 10)
        }
    }

    pause() {
        this.now = this.tick()
        clearInterval(this.interval)
    }

    tick() {
        let tmp = Math.max(0, this.now - (((new Date()).getTime() - this.startTime)))
        if (tmp == 0) clearInterval(this.interval)
        this.handler(tmp)
        return tmp
    }

    stop() {
        clearInterval(this.interval)
        this.handler(0)
        this.down = true
    }
}
export const state = () => ({
  socket: null,
})

export const mutations = {
  connect(state, url, handlers) {
    state.socket = new WebSocket(url)

    this.socket.onmessage = (msg) => {
      const data = JSON.parse(msg.data)
      const handle = handlers[data.type]
      if (!handle) {
        console.error('unknown message type')
        return
      }
      handle(data)
    }
  },
}

<template>
  <div>
    <h1>Chat</h1>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
  middleware: ['authenticated'],

  data: () => ({
    dialogs: [],
  }),

  computed: {
    me() {
      return this.$store.state.username
    },

    socket() {
      return this.$store.chat.state.socket
    },
  },

  created() {
    this.connect(`ws://127.0.0.1:8080/ws?me=${this.me}`, {
      1: this.getMessage,
    })
  },

  methods: {
    sendMessage() {
      this.socket.send(JSON.stringify({ type: 1, recipient: '', content: '' }))
    },

    getMessage(data) {
      const { sender, content } = data
      if (sender === this.me) {
        return
      }

      const dialog = this.dialogs.filter((d) => d.companion === sender)[0]
      if (dialog) dialog.messages.push({ sender, content })
    },

    ...mapMutations({
      connect: 'chat/connect',
    }),
  },
}
</script>

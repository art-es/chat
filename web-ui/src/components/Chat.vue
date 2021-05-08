<template>
  <div>
    <div class="username">Username: {{ me }}</div>
    <div>
      <h1>List of users</h1>
      <ul>
        <li v-for="(user, index) in users" :key="index">
          {{ user }}
        </li>
      </ul>
    </div>
    <Dialog
      v-if="currentDialog"
      :me="me"
      :companion="currentDialog.companion"
      :messages="currentDialog.messages"
    >
  </div>
</template>

<script>
import Dialog from './Chat/Dialog.vue';

export default {
  props: ['me'],

  components: {
    Dialog,
  },

  data: () => ({
    socket: null,
    dialogs: [],
    currentDialog: null,
    users: [],
  }),

  methods: {
    connectToSocket() {
      this.socket = new WebSocket(`ws://127.0.0.1:8080/ws?me=${this.me}`);

      const handlers = {
        1: this.getMessage,
      };

      this.socket.onmessage = (msg) => {
        const data = JSON.parse(msg.data);
        const handle = handlers[data.type];
        if (!handle) {
          console.error('unknown message type');
          return;
        }
        handle(data);
      };
    },

    sendMessage() {
      this.socket.send(JSON.stringify({ type: 1, recipient: '', content: '' }));
    },

    getMessage(data) {
      const { sender, content } = data;
      if (sender === this.me) {
        return;
      }

      const dialog = this.dialogs.filter((d) => d.companion === sender)[0];
      if (dialog) dialog.messages.push({ sender, content });
    },
  },

  created() {
    this.connectToSocket();
  },
};
</script>

<style scoped>
.me {
  margin-bottom: 20px;
}
</style>

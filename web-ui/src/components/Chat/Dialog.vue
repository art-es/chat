<template>
  <div>
    <h2>Dialog with {{ companion }}</h2>
    <div class="chat-board">
      <div
        v-for="(message, index) in messages"
        :key="index"
        class="chat-board_message"
        :class="messageCSSClass(message.from)"
      >
        <p>{{ message.text }}</p>
        <span :class="`details-${message.from === me ? 'right' : 'left'}`">
          {{ message.from }}
        </span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['me', 'companion', 'messages'],

  methods: {
    messageCSSClass(from) {
      if (from === 'System') {
        return 'chat-board_message--system';
      }
      if (from === this.username) {
        return 'chat-board_message--darker';
      }
      return '';
    },
  },
};
</script>

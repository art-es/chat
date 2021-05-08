<template>
  <div>
    <h2>Room #{{ roomId }}</h2>
    <div class="chat-board">
      <div
        v-for="(message, index) in messages"
        :key="index"
        class="container"
        :class="containerClass(message.from)"
      >
        <p>{{ message.text }}</p>
        <span :class="`details-${message.from === username ? 'right' : 'left'}`">
          {{ message.from }}
        </span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['username', 'roomId', 'messages'],

  methods: {
    containerClass(from) {
      if (from === 'System') {
        return 'system';
      }
      if (from === this.username) {
        return 'darker';
      }
      return '';
    },
  },
};
</script>

<style scoped>
.container {
  border: 2px solid #dedede;
  background-color: #f1f1f1;
  border-radius: 5px;
  padding: 10px;
  margin: 10px 0;
}

.darker {
  border-color: #ccc;
  background-color: #ddd;
}

.system {
  border-color: wheat;
}

.container::after {
  content: '';
  clear: both;
  display: table;
}

.details-right {
  float: right;
  color: #aaa;
}

.details-left {
  float: left;
  color: #999;
}
</style>

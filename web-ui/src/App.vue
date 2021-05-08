<template>
  <div id="app">
    <SignIn v-if="!authenticated" @setUsername="setUsername" />
    <Chat v-else :me="username" />
  </div>
</template>

<script>
import SignIn from './components/SignIn.vue';
import Chat from './components/Chat.vue';

export default {
  name: 'app',

  components: {
    SignIn,
    Chat,
  },

  data: () => ({
    username: '',
  }),

  computed: {
    authenticated() {
      return Boolean(this.username);
    },
  },

  methods: {
    setUsername(username) {
      this.username = username;
    },
  },
};
</script>

<style lang="scss">
/* Forms */
.btn {
  background-color: #f1f1f1;
  border: 1px solid #dedede;
  border-radius: 3px;
  color: #000;
  padding: 10px 24px;
  cursor: pointer;
  margin-right: 10px;
}

.input {
  height: 25px;
}

/* Chat */
.chat-board {
  &_message {
    border: 2px solid #dedede;
    background-color: #f1f1f1;
    border-radius: 5px;
    padding: 10px;
    margin: 10px 0;

    &::after {
      content: '';
      clear: both;
      display: table;
    }

    &--system {
      border-color: wheat;
    }

    &--darker {
      border-color: #ccc;
      background-color: #ddd;
    }
  }

  &_details {
    &--right {
      float: right;
      color: #aaa;
    }

    &--left {
      float: left;
      color: #999;
    }
  }
}
</style>

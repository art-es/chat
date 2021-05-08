<template>
  <div>
    <label style="display: block">
      Username:
      <input type="text" v-model="value" class="input" />
    </label>

    <span style="color: red">{{ error }}</span><br />
    <button class="btn" @click="submit">Submit</button>
  </div>
</template>

<script>
export default {
  data: () => ({
    value: '',
    error: '',
  }),

  methods: {
    submit() {
      if (this.value === '') {
        this.error = 'Please, enter not empty string';
        return;
      }
      this.$emit('setUsername', this.value);
    },
  },
};
</script>

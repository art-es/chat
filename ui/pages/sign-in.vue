<template>
  <div>
    <el-row style="margin-bottom: 5px">
      <el-col :span="12">
        <el-alert
          v-for="(error, index) in errors"
          :key="index"
          :title="error.title"
          type="error"
          show-icon
          :close="closeError(index)"
          style="margin-bottom: 5px"
        />
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-input v-model="value" placeholder="Please enter username" clearable>
          <el-button slot="append" @click="submit">Sign In</el-button>
        </el-input>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
  data: () => ({
    value: '',
    errors: [],
  }),

  methods: {
    submit() {
      try {
        this.setUsername(this.value)
        this.$router.push({ path: '/' })
      } catch (e) {
        this.errors.push({ title: e.message })
      }
    },

    closeError(index) {
      return () => {
        this.errors.splice(index, 1)
      }
    },

    ...mapMutations(['setUsername']),
  },
}
</script>

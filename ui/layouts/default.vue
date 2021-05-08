<template>
  <div>
    <el-container style="height: 500px; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu :default-active="activeMenuItem">
          <el-menu-item
            v-for="menuItem in menuItems"
            :key="menuItem.route"
            :index="menuItem.index"
            @click="$router.push({ path: menuItem.path })"
          >
            {{ menuItem.title }}
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header style="text-align: right; font-size: 12px">
          <el-dropdown v-if="authenticated">
            <i class="el-icon-setting" style="margin-right: 15px"></i>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click="exit">Exit</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
          <span v-if="authenticated">{{ this.$store.state.username }}</span>
          <NuxtLink v-else to="/sign-in">Sign In</NuxtLink>
        </el-header>

        <el-main>
          <Nuxt />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
  data() {
    const menuItems = [
      { title: 'Home', path: '/', route: 'index' },
      { title: 'Chat', path: '/chat', route: 'chat' },
    ].map((item, idx) => {
      item.index = String(idx)
      return item
    })

    return { menuItems }
  },

  computed: {
    activeMenuItem() {
      const menuItem = this.menuItems.filter(
        (item) => item.route === this.$route.name
      )
      return menuItem[0] ? menuItem[0].index : '0'
    },

    ...mapGetters(['authenticated']),
  },

  methods: {
    exit() {
      this.logout()
      this.$router.push({ path: '/' })
    },

    ...mapMutations(['logout']),
  },
}
</script>

<style>
.el-header {
  background-color: #b3c0d1;
  color: #333;
  line-height: 60px;
}

.el-aside {
  color: #333;
}
</style>

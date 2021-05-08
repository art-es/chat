export const state = () => ({
  username: null,
})

export const getters = {
  authenticated: (state) => Boolean(state.username),
}

export const mutations = {
  setUsername(state, username) {
    if (username === '') {
      throw new Error('username must not be an empty string')
    }
    state.username = username
  },

  logout(state) {
    state.username = ''
  },
}

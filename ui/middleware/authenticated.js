export default function ({ store, redirect }) {
  if (!store.getters.authenticated) {
    redirect('/sign-in')
  }
}

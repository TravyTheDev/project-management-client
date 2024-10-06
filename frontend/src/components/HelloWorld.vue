<script lang="ts" setup>
import {onMounted, ref} from 'vue'
import {GetUser, Login, Logout} from '../../wailsjs/go/main/App'

type User = {
  id: number;
  username: string;
  email: string;
}

const user = ref<User>()
const email = ref("")
const password = ref("")

const login = async () => {
  await Login(email.value, password.value)
  window.location.reload()
}

const logout = async () => {
  await Logout()
  user.value = undefined
}

const getUser = async () => {
  const res = await GetUser()
  user.value = res
}

onMounted(() => {
  getUser()
})

</script>

<template>
  <main>
    <div v-if="user?.id">
      <p>Hello {{ user.username }}</p>
      <button @click="logout">LOGOUT</button>
    </div>
    <div v-else>
      <p>email:</p>
      <input v-model="email" type="text">
      <p>password:</p>
      <input v-model="password" type="text">
      <button @click="login">LOGIN</button>
    </div>
  </main>
</template>

<style scoped>
</style>

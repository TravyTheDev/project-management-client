<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
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
    <div v-if="user">
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
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>

<script lang="ts" setup>
import { RouterView, useRouter } from 'vue-router';
import { onMounted, provide, ref } from 'vue';
import { types } from '../wailsjs/go/models';
import { getUser } from './functions';

const router = useRouter()
const loginUser = ref<types.User>()

const loadUser = async () => {
  loginUser.value = await getUser()
  if(loginUser.value?.id){
    router.push('/main')
  }else{
    router.push('/login')
  }
}

provide("loginUser", loginUser)
const colorTheme = ref("light")
//TODO SETTINGS DB TABLE
document.body.classList.add(colorTheme.value)
onMounted(() => {
  loadUser()
})

</script>

<template>
  <RouterView class="bg-backgroundColor h-[100vh] text-fontColor" />
</template>

<style>
</style>

<script lang="ts" setup>
import { RouterView, useRouter } from 'vue-router';
import { onMounted, provide, ref } from 'vue';
import { projects } from '../wailsjs/go/models';
import { GetUser } from '../wailsjs/go/main/App';
import { getUser } from './functions';

const router = useRouter()
const loginUser = ref<projects.User>()

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
//TODO SEARCH PROJECTS
//TODO NOTIFICATIONS
document.body.classList.add(colorTheme.value)
onMounted(() => {
  loadUser()
})
// const getItemKey = (item: any) => item;
</script>

<template>
  <RouterView class="bg-backgroundColor h-[100vh] text-fontColor" />
</template>

<style>
</style>

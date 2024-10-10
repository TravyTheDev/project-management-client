<script lang="ts" setup>
import { RouterView, useRoute, useRouter } from 'vue-router';
import Layout from './components/Layout.vue'
import { onMounted, provide, ref } from 'vue';
import { projects } from '../wailsjs/go/models';
import { GetUser } from '../wailsjs/go/main/App';

const route = useRoute()
const router = useRouter()
const loginUser = ref<projects.User>()

const getUser = async () => {
  try {
    const res = await GetUser()
    loginUser.value = res
    if (loginUser.value.id === 0) {
      router.push('/login')
    }else {
      router.push('/main')
    }
  } catch (error) {
    console.log(error)
  }
}

provide("loginUser", loginUser)
const colorTheme = ref("light")
//TODO SETTINGS DB TABLE
//TODO SEARCH PROJECTS
//TODO NOTIFICATIONS
document.body.classList.add(colorTheme.value)
onMounted(() => {
  getUser()
})
const getItemKey = (item: any) => item;
</script>

<template>
  <Layout :key="getItemKey(loginUser)" class="bg-backgroundColor h-[100vh] text-fontColor">
    <RouterView :key="route.path" />
  </Layout>
</template>

<style>
</style>

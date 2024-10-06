import { createWebHistory, createRouter } from 'vue-router'

import Main from './pages/Main.vue'
import Login from './pages/Login.vue'
import Timeline from './pages/Timeline.vue'



const routes = [
  { path: '/', component: Main },
  { path: '/login', component: Login },
  { path: '/timeline', component: Timeline },


]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
import { createWebHistory, createRouter } from 'vue-router'

import Main from './pages/Main.vue'
import Login from './pages/Login.vue'
import Timeline from './pages/Timeline.vue'
import Project from './pages/Project.vue'



const routes = [
  { path: '/', component: Main },
  { path: '/login', component: Login },
  { path: '/timeline', component: Timeline },
  { path: '/project/:id', component: Project },


]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
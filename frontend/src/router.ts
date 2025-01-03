import { createWebHashHistory, createRouter } from 'vue-router'

import Main from './pages/Main.vue'
import Login from './pages/Login.vue'
import Timeline from './pages/Timeline.vue'
import Urgency from './pages/Urgency.vue'
import Project from './pages/Project.vue'
import Register from './pages/Register.vue'
import Layout from './components/Layout.vue'



const routes = [
  {
    path: '/',
    component: Layout,
    children: [
      {
        path: 'main',
        component: Main,
      },
      {
        path: 'main/reload',
        component: Main,
      },
      {
        path: 'urgency',
        component: Urgency,
      },
      {
        path: 'timeline',
        component: Timeline,
      },
      {
        path: '/project/:id',
        component: Project,
      },
      {
        path: '/register',
        component: Register,
      },
    ],
  },
  { path: '/login', component: Login },

]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'

import Dashboard from '../views/Dashboard.vue'
import Deployments from "@/views/Deployments.vue";
import ReplicaSets from "@/views/ReplicaSets.vue";
import Namespaces from "@/views/Namespaces.vue";
import Pods from "@/views/Pods.vue";

const routes = [
  { path: '/', component: Dashboard },
  { path: '/pods', component: Pods },
  { path: '/deployments', component: Deployments },
  { path: '/replicasets', component: ReplicaSets },
  { path: '/namespaces', component: Namespaces },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

// Workaround for https://github.com/vitejs/vite/issues/11804
router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (!localStorage.getItem('vuetify:dynamic-reload')) {
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    } else {
      console.error('Dynamic import error, reloading page did not fix it', err)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router

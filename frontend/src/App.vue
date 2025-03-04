<template>
  <ContextSelector :show="contextSelectorShow" @close="contextSelectorShow = false"/>
  <v-app>
    <v-layout>

      <v-app-bar
          style="--draggable:drag"
          color="primary"
          density="compact"
      >
        <template v-slot:prepend>
          <v-btn @click="toggleRail" icon="mdi-kubernetes"></v-btn>
          <v-btn @click="contextSelectorShow = true" icon="mdi-feature-search-outline"></v-btn>
        </template>
        <template v-slot:title>
          {{ contextStore.current }}
          <v-btn icon color="warning">
            <v-icon icon="mdi-alert-box-outline"></v-icon>
            <v-tooltip
                activator="parent"
                location="bottom"
            >
              Ram Usage
            </v-tooltip>
          </v-btn>
        </template>
        <template v-slot:append>
          <v-btn @click="toggleRail" icon="mdi-cog"></v-btn>
        </template>
      </v-app-bar>

      <v-navigation-drawer
          :rail="isRail"
          v-model="drawler"
          permanent
      >
        <v-list density="compact" nav>
          <v-list-item
              v-for="item in sidebarItems"
              :title="item.title"
              :to="item.goTo"
              :prepend-icon="item.icon"
              :color="item.color"
          ></v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <v-container fluid>
          <RouterView />
        </v-container>
      </v-main>
    </v-layout>
  </v-app>
</template>

<script setup>
import { ref, watch } from 'vue'
import ContextSelector from "@/components/ContextSelector.vue";

// Context state
import { useContextStore } from '@/stores/context'
const contextStore = useContextStore()
const selectedContext = ref("")

const isRail = ref(true)
const group = ref(null)
const contextSelectorShow = ref(false)
const drawler = true

const sidebarItems = [
  {
    title: 'Dashboard',
    goTo: '/',
    icon: 'mdi-view-dashboard',
    color: "info"
  },
  {
    title: 'Namespaces',
    goTo: '/namespaces',
    icon: 'custom:namespace',
  },
  {
    title: 'Pods',
    goTo: '/pods',
    icon: 'custom:pod',
  },
  {
    title: 'Deployments',
    goTo: '/deployments',
    icon: 'custom:deployment',
  },
  {
    title: 'ReplicaSets',
    goTo: '/replicasets',
    icon: 'custom:replicaset',
  },

]

// Functions ==========
function toggleRail() {
  isRail.value = !isRail.value
}
// ====================

// Watchers ===========
watch(contextStore, async (ctx) => {
  selectedContext.value = ctx.current
});
// ====================
</script>

<style>
</style>
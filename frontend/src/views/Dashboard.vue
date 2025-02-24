<template>
  <v-container fluid class="fill-height d-flex flex-column">
    <!-- Fixed Sections -->
    <v-container fluid style="flex: none;">
      <v-row>
        <v-col cols="12">
          <v-alert
              density="compact"
              text="Cluster Almost got its limits by usage of the RAM"
              title="Memory Usage"
              type="warning"
              v-model="clusterAlert"
              closable
          ></v-alert>
        </v-col>
        <v-col cols="4">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-card-title>
              <v-icon icon="custom:pod"></v-icon>
              Nodes
            </v-card-title>
            <v-card-subtitle>{{ clusterStateStore.nodes.ready }}/{{ clusterStateStore.nodes.total }}</v-card-subtitle>
            <v-card-text>
              <v-progress-linear
                :location="null"
                buffer-color="deep-orange"
                buffer-opacity="1"
                :buffer-value="clusterStateStore.nodes.total"
                color="green"
                :max="clusterStateStore.nodes.total"
                min="0"
                :model-value="clusterStateStore.nodes.ready"
                rounded
              ></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-col>

        <v-col cols="4">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-card-title>
              <v-icon icon="custom:pod"></v-icon>
              Pods
            </v-card-title>
            <v-card-subtitle>{{ clusterStateStore.pods.ready }}/{{ clusterStateStore.pods.total }}</v-card-subtitle>
            <v-card-text>
              <v-progress-linear
                  :location="null"
                  buffer-color="deep-orange"
                  buffer-opacity="1"
                  color="green"
                  :max="clusterStateStore.pods.total"
                  min="0"
                  :model-value="clusterStateStore.pods.ready"
                  :buffer-value="clusterStateStore.pods.total"
                  rounded
              ></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-col>

        <v-col cols="4">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-card-title>
              <v-icon icon="custom:pod"></v-icon>
              Events
            </v-card-title>
            <v-card-subtitle>823/999</v-card-subtitle>
            <v-card-text>
              <v-progress-linear
                  :location="null"
                  buffer-color="deep-orange"
                  buffer-opacity="1"
                  color="green"
                  max="999"
                  min="0"
                  model-value="823"
                  buffer-value="999"
                  rounded
              ></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="6">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-card-title>
              <v-icon icon="custom:pod"></v-icon>
              CPU
            </v-card-title>
            <v-card-subtitle>{{ clusterStateStore.cpu.used }}/{{ clusterStateStore.cpu.total }} Cores</v-card-subtitle>
            <v-card-text>
              <v-progress-linear
                  :location="null"
                  color="green"
                  max="100"
                  min="0"
                  model-value="45"
                  rounded
              ></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-col>

        <v-col cols="6">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-card-title>
              <v-icon icon="custom:pod"></v-icon>
              Memory
            </v-card-title>
            <v-card-subtitle>{{ clusterStateStore.ram.used }}/{{ clusterStateStore.ram.total }} Gb</v-card-subtitle>
            <v-card-text>
              <v-progress-linear
                  :location="null"
                  color="deep-orange"
                  max="100"
                  min="0"
                  model-value="92"
                  rounded
              ></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>

    <!-- Expandable EVENTS Section -->
    <v-container fluid style="flex-grow: 1;">
      <v-row class="fill-height">
        <v-col cols="12" class="d-flex flex-column">
          <v-card class="v-card py-4 flex-grow-1" color="surface-variant" rounded="lg" variant="outlined">
            <v-card-text class="v-card__text">
              <v-data-table :items="items"></v-data-table>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>

<script setup>
import { watch, ref } from "vue";

const clusterAlert = ref(false)

// Cluster state
import { useClusterStateStore } from '@/stores/clusterState'
const clusterStateStore = useClusterStateStore()

const items = [
  {
    name: 'African Elephant',
    species: 'Loxodonta africana',
    diet: 'Herbivore',
    habitat: 'Savanna, Forests',
  }
]

// Watchers ===========
watch(clusterStateStore, async (ctx) => {
  console.log(ctx)
});
// ====================
</script>

<style scoped>
</style>
<template>
    <v-dialog
        v-model="isVisible"
        max-width="600"
    >
      <v-card
          width="auto"
          prepend-icon="mdi-feature-search-outline"
          title="Cluster contexts"
      >
        <v-card-text>
          <v-row dense>
            <v-col
                cols="12"
            >
              <v-autocomplete
                  :items="contexts"
                  v-model="selectedContexts"
              ></v-autocomplete>
            </v-col>
          </v-row>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
              class="ms-auto"
              text="Ok"
              @click="emit('close')"
          ></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import {defineProps, defineEmits, computed, onMounted, ref, watch} from "vue";
import { GetContexts } from '../../wailsjs/go/main/App'

// Cluster state
import { useClusterStateStore } from '@/stores/clusterState'
const clusterStateStore = useClusterStateStore()

// Context state
import { useContextStore } from '@/stores/context'
const contextStore = useContextStore()

const props = defineProps<{ show: boolean }>();
const emit = defineEmits(["close"]);
const contexts = ref(["1", "2", "3", "4", "5"]);
const selectedContexts = ref("")

const isVisible = computed({
  get: () => props.show,
  set: (value) => {
    if (!value) emit("close")
  }
});

// Functions ==========
function initCluster() {
  clusterStateStore.setNodesState({ total: 9, ready: 7 })
  clusterStateStore.setPodsState({ total: 30, ready: 28 })
  clusterStateStore.setCpuState({ total: 38, used: 15 })
  clusterStateStore.setRamState({ total: 144, used: 139 })
}
// ====================

// Watchers ===========
watch(selectedContexts, async (ctx) => {
  contextStore.set(ctx)
  initCluster()
});
// ====================

// On load ==========
onMounted(() => {
  GetContexts().then((result) => {
    let data = JSON.parse(result)
    let ctxList = []
    for (let key in data) {
      ctxList.push(key)
    }
    contexts.value = ctxList
    selectedContexts.value = ctxList[0]
    contextStore.set(selectedContexts.value)
    initCluster()
  });
})
// ====================
</script>
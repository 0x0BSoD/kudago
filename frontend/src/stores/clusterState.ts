import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

import { GetDashboardData } from '../../wailsjs/go/main/App'
import { formatBytes } from "@/helpers";

interface Nodes {
    total: number,
    ready: number
}

interface Cpu {
    total: number,
    used: number
}

interface Ram {
    total: number,
    used: number,
    total_unit: string,
    used_unit: string
}

interface Pods {
    total: number,
    ready: number
}

interface ClusterState {
    nodes: Nodes,
    cpu: Cpu,
    ram: Ram,
    pods: Pods,
}

export const useClusterStateStore = defineStore('clusterState', () => {
    // State
    const nodes = ref<Nodes>({ready: 0, total: 0})
    const cpu = ref<Cpu>({used: 0, total: 0})
    const ram = ref<Ram>({used: 0, total: 0, used_unit: "B", total_unit: "B"})
    const pods = ref<Pods>({ready: 0, total: 0})

    // Polling state
    const isPolling = ref<boolean>(false)
    const pollingInterval = ref<number>(1000) // Default 30 seconds
    const loading = ref<boolean>(false)
    const InitialLoading = ref<boolean>(false)
    const error = ref<string | null>(null)

    // Timer reference
    let pollingTimer: number | null = null

    // Getters
    const getNodes = computed((): Nodes => nodes.value)
    const getCpu = computed((): Cpu => cpu.value)
    const getRam = computed((): Ram => ram.value)
    const getPods = computed((): Pods => pods.value)

    // Actions for state updates
    function setNodesState(data: Nodes): void {
        nodes.value = data
    }

    function setCpuState(data: Cpu): void {
        cpu.value = data
    }

    function setRamState(data: Ram): void {
        ram.value = data
    }

    function setPodsState(data: Pods): void {
        pods.value = data
    }

    // API fetching action
    async function fetchClusterState(): Promise<void> {
        if (InitialLoading.value) return

        try {
            loading.value = true
            error.value = null

            const response = await GetDashboardData()

            const data = JSON.parse(response)

            // Update state with received data
            if (data) {
                setNodesState({
                    total: data.total,
                    ready: data.ready,
                })
                setCpuState({
                    total: data.cpu_total,
                    used: data.cpu_used,
                })

                const [total_num, total_unit] = formatBytes(data.mem_total)
                const [used_num, used_unit] = formatBytes(data.mem_used)
                setRamState({
                    total: total_num,
                    total_unit: total_unit,
                    used: used_num,
                    used_unit: used_unit,
                })

                setPodsState({
                    total: data.pods_total,
                    ready: data.pods_ready,
                })
            }

        } catch (err) {
            error.value = err instanceof Error ? err.message : 'Unknown error'
            console.error('Error fetching cluster state:', err)
        } finally {
            loading.value = false
        }
    }

    // Polling control actions
    function startPolling(): void {
        if (isPolling.value) return

        isPolling.value = true

        // Immediate first fetch
        InitialLoading.value = true
        fetchClusterState()

        // Set up interval for subsequent fetches
        InitialLoading.value = false
        pollingTimer = window.setInterval(fetchClusterState, pollingInterval.value)
    }

    function stopPolling(): void {
        if (!isPolling.value) return

        isPolling.value = false

        if (pollingTimer !== null) {
            window.clearInterval(pollingTimer)
            pollingTimer = null
        }
    }

    function setPollingInterval(interval: number): void {
        pollingInterval.value = interval

        // Restart polling with new interval if already active
        if (isPolling.value) {
            stopPolling()
            startPolling()
        }
    }

    // Clean up when store is disposed
    function $dispose(): void {
        stopPolling()
    }

    return {
        // State
        nodes,
        cpu,
        ram,
        pods,
        isPolling,
        pollingInterval,
        loading,
        error,

        // Getters
        getNodes,
        getCpu,
        getRam,
        getPods,

        // State update actions
        setNodesState,
        setCpuState,
        setRamState,
        setPodsState,

        // Polling actions
        fetchClusterState,
        startPolling,
        stopPolling,
        setPollingInterval
    }
})
import { defineStore } from 'pinia'

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
    used: number
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

export const useClusterStateStore = defineStore('clusterState',{
    state: (): ClusterState => ({
        nodes: { ready: 0, total: 0 },
        cpu: { used: 0, total: 0 },
        ram: { used: 0, total: 0 },
        pods: { ready: 0, total: 0 }
    }),

    getters: {
        getNodes: (state): Nodes => state.nodes,
        getCpu: (state): Cpu => state.cpu,
        getRam: (state): Ram => state.ram,
        getPods: (state): Pods => state.pods
    },

    actions: {
        setNodesState(data: Nodes): void {
            this.nodes = data;
        },
        setCpuState(data: Cpu): void {
            this.cpu = data;
        },
        setRamState(data: Ram): void {
            this.ram = data;
        },
        setPodsState(data: Pods): void {
            this.pods = data;
        },
    }
})
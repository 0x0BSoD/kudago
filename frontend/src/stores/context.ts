import { defineStore } from 'pinia'

interface ContextState {
    current: string;
}

export const useContextStore = defineStore('context',{
    state: (): ContextState => ({
        current: 'nil' as string
    }),

    getters: {
        get: (state): string => state.current
    },

    actions: {
        set(ctx: string): void {
            this.current = ctx;
        },
    }
})
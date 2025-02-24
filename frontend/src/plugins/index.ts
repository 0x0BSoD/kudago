// Plugins
import vuetify from './vuetify'
import router from '../router'
import pina from './pina'

// Types
import type { App } from 'vue'

export function registerPlugins (app: App) {
  app
      .use(vuetify)
      .use(pina)
      .use(router)
}

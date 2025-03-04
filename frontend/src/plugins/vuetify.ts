import {defineAsyncComponent, h} from 'vue'

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import {createVuetify, type IconProps, type IconSet} from 'vuetify'

// Custom Icons
const k8sIcons: Record<string, any> = {
    'pod': defineAsyncComponent(() => import('@/assets/icons/pod.svg')),
    'namespace': defineAsyncComponent(() => import('@/assets/icons/namespace.svg')),
    'deployment': defineAsyncComponent(() => import('@/assets/icons/deployment.svg')),
    'replicaset': defineAsyncComponent(() => import('@/assets/icons/replica-set.svg')),
}

const k8sIconsSet: IconSet = {
    component: (props: IconProps) => {
        if (typeof props.icon === "string") {
            const IconComponent = k8sIcons[props.icon.replace('custom:', '')]
            if (!IconComponent) {
                return h('span', {class: 'mdi mdi-help-circle'})
            }
            return h(
                'i',
                { class: 'v-icon notranslate v-theme--light v-icon--size-default' }, // Vuetify expects an `i` element wrapper
                [h(IconComponent, { class: 'v-icon__svg', width: '3em', height: '3em' })]
            )
        }
    }
}

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
    icons: {
        defaultSet: 'mdi',
        sets: {
            custom: k8sIconsSet
        }
    },
    theme: {
        defaultTheme: 'dark'
    }
})

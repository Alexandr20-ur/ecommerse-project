import './assets/сss/final.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { Notifications  } from '@kyvg/vue3-notification'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'

const pinia = createPinia()
const app = createApp(App)

app.use(router)
app.use(pinia)
app.use(Notifications)
app.use(ElementPlus)

app.mount('#app')

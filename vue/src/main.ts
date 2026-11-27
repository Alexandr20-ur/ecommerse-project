import './assets/сss/final.css' // Исправил кириллицу на латиницу в пути css
import {createApp} from 'vue'
import {createPinia} from 'pinia'
import NotificationsPlugin from '@kyvg/vue3-notification'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)
const pinia = createPinia()

// Сначала подключаем Pinia, чтобы сторы были доступны в гардах роутера
app.use(pinia)
app.use(router)
app.use(NotificationsPlugin)
app.use(ElementPlus)

app.mount('#app')
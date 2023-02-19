import { createApp } from 'vue'
import App from './App.vue'
import Notifications from 'notiwind'

const app = createApp(App)
app.use(Notifications)
app.mount('#app')

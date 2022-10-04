import {createApp} from 'vue'
import App from './App.vue'
import WaveUI from 'wave-ui'
import './style.css'

const app = createApp(App)
new WaveUI(app, {})
app.mount('#app')

import {createApp} from 'vue'
import App from './App.vue'
import './styles/index.less';

const app = createApp(App)
import router from './router/index'
app.use(router)
app.mount('#app')

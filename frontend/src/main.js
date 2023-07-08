import 'primevue/resources/themes/tailwind-light/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

import { createApp } from 'vue'
import './style.scss'
import PrimeVue from 'primevue/config';
import App from './App.vue'

const app = createApp(App)
app.use(PrimeVue)
app.mount('#app')

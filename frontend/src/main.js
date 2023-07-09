import 'primevue/resources/themes/tailwind-light/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

import { createApp } from 'vue'
import './style.scss'
import PrimeVue from 'primevue/config';
import App from './App.vue'

//global components
import SelectButton from 'primevue/selectbutton'
import ToggleButton from 'primevue/togglebutton'
import Chart from "vue3-apexcharts";


const app = createApp(App)
app.use(PrimeVue)

app.component('SelectButton', SelectButton)
app.component('ToggleButton', ToggleButton)
app.component('Chart', Chart)

app.mount('#app')

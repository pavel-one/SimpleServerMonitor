<template>
  <div>
    <logo />
    <temps  :data="temps"/>
  </div>
</template>

<script>
import Logo from './components/Logo.vue'
import Temps from "./components/TempCharts.vue";
import store from "./store/store.js";
import {helpers} from "./support/helpers.js";

export default {
  components: {
    Temps,
    Logo
  },
  data() {
    return {
      temps: null
    }
  },
  mounted() {
    store.Connection.onmessage = event => {
      const json = JSON.parse(event.data)
      const e = helpers.getEvent(json.event)

      if (e.channel === 'temp' && e.name === 'all') {
        this.temps = json.data
      }

      if (e.channel === 'temp' && e.name === 'add') {
        // console.log(this.temps)
        console.log("added", json.data)

        // this.temps.labels.push(json.data.labels[0])
      }

      // console.log("event:", e, json)
      // console.log('app:', json)
    }

    setTimeout(() => {
      this.temps = {
        labels: [ 'January', 'February', 'March'],
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: [40, 20, 12]
          }
        ]
      }
    }, 5000)
  }
}
</script>

<style scoped>

</style>

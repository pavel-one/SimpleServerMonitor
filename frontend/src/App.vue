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
    console.log(store.Connection)
    store.Connection.onmessage = event => {
      const json = JSON.parse(event.data)
      const e = helpers.getEvent(json.event)

      if (e.channel === 'temp' && e.name === 'all') {
        const dataset = helpers.formatAllTemps(json.data)
        console.log(dataset)
        this.temps = dataset
      }

      console.log("event:", e, json)
      // console.log('app:', json)
    }
  }
}
</script>

<style scoped>

</style>

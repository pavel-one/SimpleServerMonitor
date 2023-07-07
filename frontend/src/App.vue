<template>
  <div>
    <logo />
    <VueApexCharts width="500" type="line" :options="options" :series="series"></VueApexCharts>
  </div>
</template>

<script>
import Logo from './components/Logo.vue'
import store from "./store/store.js";
import {helpers} from "./support/helpers.js";

import VueApexCharts from "vue3-apexcharts";

export default {
  components: {
    Logo,
    VueApexCharts
  },
  data() {
    return {
      temps: null,
      options: {
        chart: {
          id: 'vuechart-example'
        },
        xaxis: {
          categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
        }
      },
      series: [{
        name: 'series-1',
        data: [30, 40, 45, 50, 49, 60, 70, 91]
      }]
    }
  },
  methods: {
    addTemp: async function (sensor_id, temp) {
      // for (const item of this.temps.datasets) {
      //   if (item.sensor_id === sensor_id) {
      //     item.data[item.data.length] = 10
      //   }
      // }
    }
  },
  mounted() {
    store.Connection.onmessage = async event => {
      const json = JSON.parse(event.data)
      const e = helpers.getEvent(json.event)

      if (e.channel === 'temp' && e.name === 'all') {
        this.temps = json.data
        return
      }

      if (e.channel === 'temp' && e.name === 'add') {
        // let newChart = copy(this.temps)
        //
        // newChart.labels.push(json.data.labels[0])
        //
        // for (const item of json.data.datasets) {
        //   for (let nItem of newChart.datasets) {
        //     if (nItem.sensor_id === item.sensor_id) {
        //       nItem.data.push(item.data[0])
        //     }
        //   }
        // }

      }

    }
  }
}
</script>

<style scoped>

</style>

<template>
  <div>
    <logo />
    <Line ref="chart" style="width: 1000px; height: 400px" v-if="temps" :data="temps" :options="options"/>
  </div>
</template>

<script>
import Logo from './components/Logo.vue'
import store from "./store/store.js";
import {helpers} from "./support/helpers.js";

import {CategoryScale, Chart as ChartJS, Legend, LinearScale, LineElement, PointElement, Title, Tooltip, Colors} from 'chart.js'
import {Line} from "vue-chartjs";

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
    Colors
)

export default {
  components: {
    Line,
    Logo
  },
  data() {
    return {
      temps: null,
      options: {
        responsive: true
      }
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
        let newChart = copy(this.temps)

        newChart.labels.push(json.data.labels[0])

        for (const item of json.data.datasets) {
          for (let nItem of newChart.datasets) {
            if (nItem.sensor_id === item.sensor_id) {
              nItem.data.push(item.data[0])
            }
          }
        }

      }

    }
  }
}
</script>

<style scoped>

</style>

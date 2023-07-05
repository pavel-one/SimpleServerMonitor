<template>
  <Line style="width: 1000px; height: 400px" v-if="data" :data="data" :options="options"/>
</template>

<script>
import {CategoryScale, Chart as ChartJS, Legend, LinearScale, LineElement, PointElement, Title, Tooltip, Colors} from 'chart.js'
import {Line} from 'vue-chartjs'

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
    Line
  },
  props: {
    data: {
      type: Object,
    }
  },
  data() {
    return {
      options: {
        responsive: true
      }
    }
  }
}
</script>

<template>
  <Line style="width: 1000px; height: 400px" v-if="data" :data="data" :options="options"/>
</template>

<script>
import {CategoryScale, Chart as ChartJS, Legend, LinearScale, LineElement, PointElement, Title, Tooltip} from 'chart.js'
import {Line} from 'vue-chartjs'

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
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
      //     data: {
      //       labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
      //       datasets: [
      //         {
      //           label: 'Data One',
      //           backgroundColor: '#f87979',
      //           data: [40, 39, 10, 40, 39, 80, 40]
      //         },
      //         {
      //           label: 'Data two',
      //           backgroundColor: 'blue',
      //           data: [35, 11, 10, 40, 39, 80, 40]
      //         }
      //       ]
      //     },
      options: {
        responsive: true,
        // maintainAspectRatio: false,
        // animations: {
        //   tension: {
        //     duration: 1000,
        //     easing: 'linear',
        //     from: 1,
        //     to: 0,
        //     loop: true
        //   }
        // },
        transitions: {
          show: {
            animations: {
              x: {
                from: 0
              },
              y: {
                from: 0
              }
            }
          },
          hide: {
            animations: {
              x: {
                to: 0
              },
              y: {
                to: 0
              }
            }
          }
        }
      }
    }
  }
}
</script>

<template>
  <div>
    <logo/>

    <div class="card flex justify-content-center">
      <SelectButton v-model="tempSelector.selected" optionValue="value" :options="tempSelector.options" optionLabel="name"/>
    </div>
    <VueApexCharts v-if="temps" width="500" type="line" :options="options" :series="temps"></VueApexCharts>
    <div class="card flex justify-content-center">
      <ToggleButton v-model="update" onLabel="Stop" offLabel="Updating"
                    onIcon="pi pi-times" offIcon="pi pi-check" class="w-9rem" />
    </div>
  </div>
</template>

<script>
import Logo from './components/Logo.vue'
import store from "./store/store.js";
import {helpers} from "./support/helpers.js";
import SelectButton from 'primevue/selectbutton'
import ToggleButton from 'primevue/togglebutton'

import VueApexCharts from "vue3-apexcharts";

export default {
  components: {
    Logo,
    VueApexCharts,
    SelectButton,
    ToggleButton
  },
  data() {
    return {
      temps: null,
      options: {
        chart: {
          id: 'vuechart'
        },
        xaxis: {
          categories: []
        }
      },
      update: true,
      tempSelector: {
        selected: 'second',
        options: [
          {
            name: 'Seconds',
            value: 'second'
          },
          {
            name: 'Minutes',
            value: 'minute'
          },
          {
            name: 'Hours',
            value: 'hour',
          },
          {
            name: 'Days',
            value: 'day',
          },
          {
            name: 'Months',
            value: 'month',
          },
        ]
      }
    }
  },
  mounted() {
    store.Connection.onmessage = async event => {
      const json = JSON.parse(event.data)
      const e = helpers.getEvent(json.event)

      if (e.channel === 'temp' && e.name === 'all') {
        if (!this.update) {
          return
        }

        this.options.xaxis.categories = json.data.labels
        this.temps = json.data.datasets
        return
      }

      if (e.channel === 'temp' && e.name === 'add') {
        if (!this.update) {
          return
        }

        this.options.xaxis.categories.shift()
        this.options.xaxis.categories.push(json.data.labels[0])

        for (const item of json.data.datasets) {
          for (let nItem of this.temps) {
            if (nItem.sensor_id === item.sensor_id) {
              nItem.data.shift()
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

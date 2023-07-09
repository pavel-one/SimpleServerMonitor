<template>
  <div>
    <div class="card flex justify-content-center">
      <select-button v-model="tempSelector.selected" optionValue="value" :options="tempSelector.options"
                    optionLabel="name"/>
    </div>
    <chart  v-if="temps" width="500" type="line" :options="options" :series="temps"></chart>
    <div class="card flex justify-content-center">
      <toggle-button v-model="update" onLabel="Stop" offLabel="Updating"
                    onIcon="pi pi-times" offIcon="pi pi-check" class="w-9rem"/>
    </div>
  </div>
</template>

<script>
import store from "../store/store.js";
import {helpers} from "../support/helpers.js";

export default {
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
            name: '1 minute',
            value: 'second'
          },
          {
            name: '1 hour',
            value: 'minute'
          },
          {
            name: '1 day',
            value: 'hour',
          },
          {
            name: '1 month',
            value: 'day',
          },
          {
            name: '1 year',
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

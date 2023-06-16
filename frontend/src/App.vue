<template>
  <div>
    <logo />
    <bar-chart></bar-chart>
  </div>
</template>

<script>
import Logo from './components/Logo.vue'
import BarChart from "./components/Charts.vue";

export default {
  components: {
    BarChart,
    Logo
  }
}
</script>

<style scoped>

</style>

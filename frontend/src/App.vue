<template>
  <div>
    <logo/>
    <TempCharts />
  </div>
</template>

<script>
import Logo from './components/Logo.vue'
import TempCharts from "./components/TempCharts.vue";

export default {
  components: {
    TempCharts,
    Logo,
  },
}
</script>
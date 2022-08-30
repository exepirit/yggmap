<template>
  <NavBar title="Yggdrasil network stats" />
  <SummaryStatPage />
</template>

<script>
import NavBar from './components/shared/NavBar.vue'
import SummaryStatPage from './feature-summary-stat/SummaryStatPage.vue'

export default {
  name: 'App',
  components: {
    NavBar,
    SummaryStatPage
  }
}

</script>

<style>
@import'~bootstrap/dist/css/bootstrap.css'
</style>

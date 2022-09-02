<template>
  <NavBar title="Yggdrasil network stats" />
  <NetworkMapPage />
</template>

<script>
import NavBar from './components/shared/NavBar.vue'
import NetworkMapPage from './feature-network-map/NetworkMapPage.vue';

export default {
  name: 'App',
  components: {
    NavBar,
    NetworkMapPage,
  }
}

</script>

<style>
@import'~bootstrap/dist/css/bootstrap.css'
</style>

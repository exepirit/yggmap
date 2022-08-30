<template>
  <div class="container">
    <NodesCountCard :nodesCount="nodesCount"/>
  </div>
</template>

<script>
import axios from 'axios'
import NodesCountCard from './NodesCountCard.vue'

export default {
  components: {
    NodesCountCard,
  },
  data() {
    return {
      loaded: false,
      nodesCount: 0,
    }
  },
  mounted() {
    axios
      .get("/api/v1/node/active")
      .then(response => {
        this.nodesCount = response.data.length
        this.loaded = true
      })
  }
}
</script>

<style>

</style>
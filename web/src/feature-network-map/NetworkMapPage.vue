<template>
  <div class="container">
    <ForceDirectedTree :graph="network"/>
  </div>
</template>

<script>
import axios from 'axios';
import ForceDirectedTree from './ForceDirectedTree.vue';

export default {
  components: {
    ForceDirectedTree,
  },
  data() {
    return {
      network: {
        nodes: [],
        edges: []
      }
    };
  },
  mounted() {
    axios
      .get("/api/v1/network/spanningTree")
      .then(response => {
        this.network = response.data;
      });
  }
}
</script>

<style>
</style>
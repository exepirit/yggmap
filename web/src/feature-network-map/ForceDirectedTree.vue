<template>
  <svg class="graph"></svg>
</template>

<script>
import * as d3 from 'd3';
import ForceGraph from './forceGraph';

export default {
  created() {
    this.$watch('graph', () => {
      this.initGraph();
    });
  },
  mounted() {
    this.initGraph();
  },
  props: {
    graph: {
      type: Object,
      required: true,
    },
  },
  methods: {
    initGraph() {
      ForceGraph({
        nodes: this.graph.nodes,
        links: this.graph.edges,
      }, {
        nodeId: (node) => node.publicKey,
        nodeTitle: (node) => node.publicKey,
        linkSource: ({from}) => from,
        linkTarget: ({to}) => to,
        nodeRadius: 5,
        onClick: (id) => {
          console.log(id);
          this.updateColors()
        },
      })
    },
    updateColors() {
      d3.select("g.nodes")
        .attr("fill", () => {
          return `rgb(0,255,0)`;
        })
    }
  }
}
</script>

<style>
</style>
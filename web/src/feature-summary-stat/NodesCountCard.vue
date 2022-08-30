<template>
  <h1 class="text-center">{{ nodesCount }}</h1>
  <p class="text-center">
    <b>nodes online</b>
  </p>
</template>

<script>
export default {
  props: {
    nodesCount: Number,
  }
}
</script>

<style>

</style>
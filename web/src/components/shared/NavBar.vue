<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">{{ title }}</a>
    </div>
  </nav>
</template>

<script>
export default {
  props: {
    title: String,
  }
}
</script>

<style>

</style>
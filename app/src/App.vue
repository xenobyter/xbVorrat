<template>
  <boxlist />
  <unitlist />
  <information />
</template>

<script>
import boxlist from "./components/boxlist.vue";
import unitlist from "./components/unitlist"
import information from "./components/information.vue";

export default {
  name: "App",
  components: {
    boxlist,
    unitlist,
    information,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
body {
  margin-left: 0px;
}
</style>

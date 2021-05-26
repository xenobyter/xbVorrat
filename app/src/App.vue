<template>
  <boxlist />
  <information />
</template>

<script>
import boxlist from "./components/boxlist.vue";
// import information from "./components/information.vue";

export default {
  name: "App",
  components: {
    boxlist,
    // information,
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

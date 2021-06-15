<template >
  <navigation v-on:nav="showNav" />
  <boxlist v-if="show=='boxlist'" />
  <unitlist v-if="show=='unitlist'" />
  <articlelist v-if="show=='articlelist'" />
  <information />
</template>

<script>
import navigation from "./components/navigation.vue";
import boxlist from "./components/boxlist.vue";
import unitlist from "./components/unitlist";
import articlelist from "./components/articlelist";
import information from "./components/information.vue";

export default {
  name: "App",
  components: {
    navigation,
    boxlist,
    unitlist,
    articlelist,
    information,
  },
  data() {
    return {
      show: "articlelist",
    };
  },
  methods: {
    showNav(item) {
      this.show = item;
    },
  },
};
</script>

<style>
:root {
  --text-color: #2c3e50;
  --listen-hintergrund: gainsboro;
  --rand-links: 1.5rem;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: var(--text-color);
}
body {
  margin-left: 1.5rem;
  margin-right: 1.5rem;
}
h1 {
  text-align: center;
}
</style>

<template >
  <navigation v-on:nav="showNav" />
  <boxlist v-if="show == 'boxlist'" v-on:nav="showNav" />
  <unitlist v-if="show == 'unitlist'" v-on:nav="showNav" />
  <articlelist v-if="show == 'articlelist'" v-on:nav="showNav" />
  <information v-if="show == 'information'" v-on:nav="showNav" />
  <mainview v-if="show == ''" v-on:nav="showNav" />
  <!-- TODO: Touchoptimierung -->
</template>

<script>
import navigation from "./components/navigation.vue";
import boxlist from "./components/boxlist.vue";
import unitlist from "./components/unitlist";
import articlelist from "./components/articlelist";
import information from "./components/information.vue";
import mainview from "./components/mainview.vue";
import { version } from "../package";

export default {
  name: "App",
  components: {
    navigation,
    boxlist,
    unitlist,
    articlelist,
    information,
    mainview,
  },
  data() {
    return {
      show: "",
      appVersion: version,
    };
  },
  methods: {
    showNav(item) {
      this.show = item;
    },
  },
  created() {
    document.title = "xbVorrat v" + this.appVersion;
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
  overflow: hidden; /* Hide scrollbars */
}
h1 {
  text-align: center;
}
.status {
  background-color: var(--listen-hintergrund);
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: absolute;
  width: calc(100vw - 2.5rem);
  height: 1.4rem;
  bottom: 0px;
  margin-left: 0px;
}
</style>

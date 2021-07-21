<template >
  <navigation v-on:nav="showNav" />
  <boxlist v-if="show == 'boxlist'" v-on:nav="showNav" />
  <unitlist v-if="show == 'unitlist'" v-on:nav="showNav" />
  <articlelist v-if="show == 'articlelist'" v-on:nav="showNav" />
  <information v-if="show == 'information'" v-on:nav="showNav" />
  <mainview v-if="show == ''" v-on:nav="showNav" />
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
/* nunito-200 - latin */
@font-face {
  font-family: "Nunito";
  font-style: normal;
  font-weight: 200;
  src: local(""), url("/fonts/nunito-v16-latin-200.woff2") format("woff2"),
    url("/fonts/nunito-v16-latin-200.woff") format("woff");
}
/* nunito-300 - latin */
@font-face {
  font-family: "Nunito";
  font-style: normal;
  font-weight: 300;
  src: local(""), url("/fonts/nunito-v16-latin-300.woff2") format("woff2"),
    url("/fonts/nunito-v16-latin-300.woff") format("woff");
}
/* nunito-regular - latin */
@font-face {
  font-family: "Nunito";
  font-style: normal;
  font-weight: 400;
  src: local(""), url("/fonts/nunito-v16-latin-regular.woff2") format("woff2"),
    url("/fonts/nunito-v16-latin-regular.woff") format("woff");
}
:root {
  --text-color: black;
  --listen-hintergrund: gainsboro;
  --rand-links: 1.5rem;
}
#app {
  color: var(--text-color);
  font-family: "Nunito", sans-serif;
  font-weight: 300;
}
body {
  margin-left: 1.5rem;
  margin-right: 1.5rem;
  overflow-x: hidden;
}
h1 {
  text-align: center;
}
.status {
  background-color: var(--listen-hintergrund);
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: fixed;
  width: calc(100vw - 3rem);
  height: 1.4rem;
  bottom: 0px;
  margin-left: 0px;
}
.icon {
  height: 1rem;
  width: auto;
}
button {
  margin: 0.15rem 0 0.15rem 0.15rem;
}
</style>

<template>
  <div class="info">
    <h1>Informationen</h1>
    <h2>Tastaturbedienung</h2>
    <div>
      Auf der Hauptseite kann mit "+" ein neuer Vorratseintrag angelegt werden.
      Mit "ESC" kann die Anlage abgebrochen werden.
    </div>
    <div>
      Die Unterfunktionen zur Anlage von Boxen, Einheiten und Artikeln sind über
      die Tasten "2" bis "4" erreichbar. Diese Informationsseite wird über die
      "5" geöffnet. Von jeder Unterfunktion führt die "1" zur Hauptseite zurück.
    </div>
  </div>
</template>

<script>
export default {
  name: "information",
  emits: ["nav"],
  methods: {
    keyHandler(e) {
      switch (e.key) {
        case "1":
          this.$emit("nav", "");
          break;
      }
    },
  },
  mounted() {
    window.addEventListener("keyup", this.keyHandler);
  },
  unmounted() {
    window.removeEventListener("keyup", this.keyHandler);
  },
};
</script>

<style scoped>
</style>

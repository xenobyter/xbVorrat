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
    <h2>Sortierung</h2>
    <div>
      Eine Sortierung der Vorratsliste nach Artikel, Box und Verfallsdatum kann
      durch Klick auf die entsprechende Spaltenüberschrift ausgelöst werden. Die
      Sortierrichtung wird durch einen erneuten Klick geändert.
    </div>
    <h2>Lizenzen und Erwähnungen</h2>
    <div>
      Die App verwendet als Schriftart <a href="https://fonts.google.com/specimen/Nunito">Nunito</a> und hosted diese lokal. Alle Symbole sind aus den <a href="https://fonts.google.com/icons">Material Icons</a> und stehen unter <a href="https://www.apache.org/licenses/LICENSE-2.0.html">Apache license 2.0</a>.
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
div {
  margin-bottom: 0.5rem;
}
</style>

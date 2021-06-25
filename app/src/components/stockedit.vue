<template>
  <div class="dialog">
    <div class="header">
      <h2 v-if="stock.articlestr != ''">Bearbeiten</h2>
      <h2 v-else>Erstellen</h2>
    </div>
    <div class="body">
      <div>
        <label for="selArticle">Artikel:</label>
        <div>
          {{ res.article }}
        </div>
        <select v-on:change="selectArticle($event)" id="selArticle" ref="start">
          <option>Auswahl:</option>
          <option
            v-for="article in articles"
            v-bind:key="article.id"
            :selected="article.id == stock.article"
            :value="article.id"
          >
            {{ article.name }}
          </option>
        </select>
      </div>
      <div>
        <label for="selBox">Box:</label>
        <div>
          {{ res.box }}
        </div>
        <select v-on:change="selectBox($event)" id="selBox">
          <option>Auswahl:</option>
          <option
            v-for="box in boxes"
            v-bind:key="box.id"
            :selected="box.id == stock.box"
            :value="box.id"
          >
            {{ box.name }}
          </option>
        </select>
      </div>
      <div>
        <label for="inDatum">Datum:</label>
        <div></div>
        <input
          type="date"
          v-on:change="inputDatum($event)"
          id="inDatum"
          :value="expiry"
        />
      </div>
      <div>
        <label for="inAnzahl">Anzahl:</label>
        <div></div>
        <input
          type="number"
          min="0"
          id="inAnzahl"
          :value="res.quantity"
          v-on:change="inputQuantity($event)"
        />
      </div>
      <div>
        <label for="inGroesse">Größe:</label>
        <div></div>
        <input
          type="number"
          min="0"
          step="0.5"
          id="inGroesse"
          :value="res.size"
          v-on:change="inputSize($event)"
        />
      </div>
    </div>
    <div class="footer">
      <button v-on:click="doEdit" v-if="stock.articlestr != ''">
        Bearbeiten
      </button>
      <button v-on:click="doAdd" v-else>Erstellen</button>
      <button v-on:click="closeEdit">Schliessen</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "stockedit",
  props: { stock: Object },
  data() {
    return {
      res: Object,
      articles: [
        {
          id: 1,
          name: "Testartikel 1",
          unit: 1,
        },
        {
          id: 2,
          name: "Testartikel 2",
          unit: 2,
        },
      ],
      boxes: [
        { id: 1, name: "Box1", notiz: "Fehler beim Abruf" },
        { id: 2, name: "Box2", notiz: "Fehler beim Abruf" },
      ],
      expiry: String,
    };
  },
  methods: {
    closeEdit() {
      this.$emit("closeedit");
    },
    doEdit() {
      this.$emit("doedit", this.res);
    },
    doAdd() {
      this.$emit("doadd", this.res);
    },
    onChange(event) {
      this.res.unit = parseInt(event.target.value);
    },
    articlesGET() {
      axios
        .get("http://localhost:8081/api/articles", { timeout: 900 })
        .then((response) => {
          this.articles = response.data;
          this.unitsGET();
        })
        .catch((e) => {
          console.log("GET", e);
          this.status = "GET: " + e.message;
        });
    },
    boxesGET() {
      axios
        .get("http://localhost:8081/api/boxes", { timeout: 900 })
        .then((response) => {
          this.boxes = response.data;
        })
        .catch((e) => {
          console.log("GET", e);
          this.status = "GET: " + e.message;
        });
    },
    selectArticle(event) {
      this.res.article = parseInt(event.target.value);
    },
    selectBox(event) {
      this.res.box = parseInt(event.target.value);
    },
    inputDatum(event) {
      this.res.expiry = this.datumZuPunkt(event.target.value);
      this.expiry = event.target.value;
    },
    datumZuStrich(str) {
      if (!str) return;
      const d = str.split(".");
      return [d[2], d[1], d[0]].join("-");
    },
    datumZuPunkt(str) {
      const d = str.split("-");
      return [d[2], d[1], d[0]].join(".");
    },
    inputQuantity(event) {
      this.res.quantity = parseInt(event.target.value);
    },
    inputSize(event) {
      this.res.size = parseFloat(event.target.value);
    },
  },
  mounted() {
    this.res = { ...this.stock };
    this.articlesGET();
    this.boxesGET();
    this.expiry = this.datumZuStrich(this.res.expiry);
    this.$nextTick(() => {
      this.$refs.start.focus();
    });
  },
};
</script>

<style scoped>
.dialog {
  z-index: 1;
  height: 100vh;
  width: 100vw;
  position: absolute;
  top: 0%;
  left: 0%;
  background-color: black;
  background-color: rgba(0, 0, 0, 0.4);
  -webkit-transition: 0.5s;
  overflow: auto;
  transition: all 0.3s linear;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.header,
.body,
.footer {
  background-color: white;
  box-shadow: 5px 5px 5px;
  margin: 0 0.5em 0 0.5em;
  padding: 0 0 0.5em 0.5em;
}
.body {
  display: flex;
  flex-direction: column;
}
.body div {
  display: flex;
  padding: 0 0 0.25rem 0;
}
label {
  min-width: 4rem;
}
.body div div {
  min-width: 2rem;
}
.body div input,
.body div select {
  -ms-box-sizing: border-box;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  min-width: 11rem;
}
</style>
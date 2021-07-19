<template>
  <div class="dialog">
    <div class="header">
      <h2 v-if="article.name != ''">Bearbeiten</h2>
      <h2 v-else>Erstellen</h2>
    </div>
    <div class="body">
      <div class="id">
        <label for="id">ID:</label>
        <input type="number" id="id" v-model="res.id" disabled />
      </div>
      <div class="name">
        <label for="name">Name:</label>
        <input type="text" id="name" v-model="res.name" ref="name" v-on:keydown.esc="closeEdit" />
      </div>
      <div class="unit">
        <label for="unit">Einheit:</label>
        <select v-on:change="onChange($event)">
          <option>Auswahl:</option>
          <option
            v-for="unit in units"
            v-bind:key="unit.id"
            :selected="unit.id == article.unit"
            :value="unit.id"
          >
            {{ unit.long }}
          </option>
        </select>
        <input type="text" id="notiz" v-model="res.unit" disabled />
      </div>
    </div>
    <div class="footer">
      <button v-on:click="doEdit" v-if="article.name != ''">Bearbeiten</button>
      <button v-on:click="doAdd" v-else>Erstellen</button>
      <button v-on:click="closeEdit">Schliessen</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "articleedit",
  props: { article: Object, units: Array },
  data() {
    return { res: Object };
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
      this.res.unit = parseInt(event.target.value)
    },
  },
  mounted() {
    this.res = { ...this.article };
    this.$nextTick(() => {
      this.$refs.name.focus();
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
.id,
.name,
.unit {
  display: flex;
}
label {
  min-width: 4em;
  margin-bottom: 0.1rem;
}
input {
  border: 0px;
  margin-bottom: 0.1rem;
}
#notiz {
  padding-left: 0.5rem;
}
select {
  border: 0px;
  outline: 0px;
}
</style>
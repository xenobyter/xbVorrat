<template>
  <div class="dialog">
    <div class="header">
      <h2 v-if="box.name != ''">Bearbeiten</h2>
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
      <div class="notiz">
        <label for="notiz">Notiz:</label>
        <input type="text" id="notiz" v-model="res.notiz" />
      </div>
    </div>
    <div class="footer">
      <button v-on:click="doEdit" v-if="box.name != ''">Bearbeiten</button>
      <button v-on:click="doAdd" v-else>Erstellen</button>
      <button v-on:click="closeEdit">Schliessen</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "boxedit",
  props: { box: Object },
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
      this.$emit("doadd", this.res)
    }
  },
  mounted() {
    this.res = { ...this.box };
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
.notiz {
  display: flex;
}
label {
  min-width: 4em;
}
input {
  border: 0px;
}
</style>
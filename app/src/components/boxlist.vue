<template>
  <div id="boxlist">
    <h1>Boxen</h1>
    <ul>
      <li v-for="box in boxes" :key="box.id">
        <div class="box">
          <div class="boxLeft">{{ box.id }}</div>
          <div class="boxMain">
            <div class="boxName">{{ box.name }}</div>
            <div class="boxNotiz">{{ box.notiz }}</div>
          </div>
          <div class="boxRight">
            <a v-on:click="boxEDIT(box)" class="edit">&#128393;</a>
            <a v-on:click="boxDELETE(box)" class="delete">&#128465;</a>
          </div>
        </div>
      </li>
      <li>
        <div class="box boxRight">
          <a v-on:click="boxAdd()" class="add">&#10750;</a>
        </div>
      </li>
      <boxedit
        v-if="showedit"
        v-on:closeedit="showedit = false"
        v-on:doedit="boxPATCH"
        v-on:doadd="boxPUT"
        :box="box"
      />
    </ul>
    <div class="status">{{ status }}</div>
  </div>
</template>

<script>
import axios from "axios";
import boxedit from "./boxedit.vue";

export default {
  name: "boxlist",
  components: {
    boxedit,
  },
  data() {
    return {
      api: process.env.VUE_APP_API,
      boxes: [],
      showedit: false,
      box: Object,
      status: "",
    };
  },
  methods: {
    boxesGET() {
      axios
        .get(this.api + "/boxes", { timeout: 900 })
        .then((response) => {
          this.boxes = response.data;
        })
        .catch((e) => {
          this.boxes = [
            { id: 1, name: "Fehler", notiz: "Fehler beim Abruf" },
            { id: 2, name: "Fehler", notiz: "Fehler beim Abruf" },
          ];
          console.log("GET", e);
          this.status = "GET: " + e.message;
        });
    },
    boxDELETE(box) {
      this.box = box;
      const index = this.boxes.indexOf(this.box);
      axios
        .delete(this.api + "/boxes/" + box.id, { timeout: 900 })
        .then((response) => {
          console.log("Status:", response.status);
          this.boxes.splice(index, 1);
        })
        .catch((e) => {
          console.error("DELETE", e.message);
          this.status = "DELETE: " + e.message;
        });
    },
    boxEDIT(box) {
      this.box = box;
      this.showedit = true;
    },
    boxPATCH(box) {
      const index = this.boxes.indexOf(this.box);
      this.showedit = false;
      axios
        .patch(
          this.api + "/boxes/" + box.id,
          {
            name: box.name,
            notiz: box.notiz,
          },
          { timeout: 900 }
        )
        .then((response) => {
          this.boxes[index] = { ...box };
          console.log("Status:", response.status);
        })
        .catch((e) => {
          console.log("PATCH", e);
          this.status = "PATCH: " + e.message;
        });
    },
    boxAdd() {
      this.box = { name: "", notiz: "" };
      this.showedit = true;
    },
    boxPUT(box) {
      this.showedit = false;
      axios
        .put(
          this.api + "/boxes",
          {
            name: box.name,
            notiz: box.notiz,
          },
          { timeout: 900 }
        )
        .then((response) => {
          console.log("Status:", response.status);
          //BUG: Fehler beim Push der ersten Box in eine leere Liste
          this.boxes.push(box);
          const index = this.boxes.indexOf(box);
          this.boxes[index].id = response.data.id;
        })
        .catch((e) => {
          console.log("PUT", e);
          this.status = "PUT: " + e.message;
        });
    },
  },
  mounted() {
    this.boxesGET();
  },
};
</script>

<style scoped>
ul {
  margin: 0;
  padding: 0;
  list-style-type: none;
}
li {
  margin-bottom: 0.5rem;
}
.box {
  overflow: hidden;
  display: flex;
  background-color: var(--listen-hintergrund);
}
.boxLeft,
.boxMain {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.boxLeft {
  padding: 0 0.5rem 0 0;
  min-width: 2rem;
  max-width: 4rem;
  text-align: right;
}
.boxMain {
  flex: 1;
}
.boxNotiz {
  font-size: small;
  color: rgb(160, 160, 160);
  min-height: 1.2rem;
}
.boxRight {
  justify-content: flex-end;
}
.status {
  background-color: var(--listen-hintergrund);
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: absolute;
  width: 100vw;
  height: 1.4rem;
  bottom: 0px;
  margin-left: 0px;
}
</style>
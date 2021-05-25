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
      boxes: [],
      showedit: false,
      box: Object,
    };
  },
  methods: {
    boxesGET() {
      axios
        .get("http://localhost:8081/api/boxes", { timeout: 900 })
        .then((response) => {
          this.boxes = response.data;
        })
        .catch((error) => {
          this.boxes = [
            { id: 1, name: "Fehler", notiz: "Fehler beim Abruf" },
            { id: 2, name: "Fehler", notiz: "Fehler beim Abruf" },
          ];
          console.log("GET", error);
        });
    },
    boxDELETE(box) {
      this.box = box;
      const index = this.boxes.indexOf(this.box);
      //TODO: Im BE prüfen, ob die Box leer ist
      axios
        .delete("http://localhost:8081/api/boxes/" + box.id, { timeout: 900 })
        .then((response) => {
          console.log("Status:", response.status);
          this.boxes.splice(index, 1);
        })
        .catch((e) => {
          console.error("DELETE", e.message); 
          // TODO: Fehler nicht nur loggen sondern in der App anzeigen, eine Art Statusleiste oder ähnliches
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
          "http://localhost:8081/api/boxes/" + box.id,
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
        .catch((error) => {
          console.log("PATCH", error);
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
          "http://localhost:8081/api/boxes",
          {
            name: box.name,
            notiz: box.notiz,
          },
          { timeout: 900 }
        )
        .then((response) => {
          console.log("Status:", response.status);
          this.boxes.push(box);
          const index = this.boxes.indexOf(box);
          this.boxes[index].id = response.data.id;
        })
        .catch((error) => {
          console.log("PUT", error);
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
  margin: 0.5em;
}
.box {
  overflow: hidden;
  margin: 0 1em 0 1em;
  display: flex;
  background-color: gainsboro;
  justify-content: space-between;
}
.boxLeft,
.boxMain {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.boxLeft {
  padding: 0 0.5em 0 0;
  min-width: 2em;
  max-width: 4em;
  text-align: right;
}
.boxMain {
  flex: 1;
}
.boxNotiz {
  font-size: small;
  color: rgb(160, 160, 160);
  min-height: 1.2em;
}
.boxRight {
  justify-content: flex-end;
}
</style>>
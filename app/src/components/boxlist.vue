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
            <button v-on:click="boxEDIT(box)" class="edit">
              <img src="/edit_black_24dp.svg" alt="&#10750;" class="icon" />
            </button>
            <button v-on:click="boxDELETE(box)" class="delete">
              <img src="/delete_black_24dp.svg" alt="&#128465;" class="icon" />
            </button>
          </div>
        </div>
      </li>
      <li>
        <div class="box boxRight">
          <button v-on:click="boxAdd()" class="add">
            <img src="/add_black_24dp.svg" alt="&#10750;" class="icon" />
          </button>
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
  emits: ["nav"],
  components: {
    boxedit,
  },
  data() {
    return {
      api: process.env.VUE_APP_API,
      boxes: [
        { id: 1, name: "Box1", notiz: "Notiz1" },
        { id: 2, name: "Box2", notiz: "Notiz2" },
      ],
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
          this.boxesGET();
        })
        .catch((e) => {
          console.log("PUT", e);
          this.status = "PUT: " + e.message;
        });
    },
    keyHandler(e) {
      switch (e.key) {
        case "+":
          this.boxAdd();
          break;
        case "1":
          this.$emit("nav", "");
          break;
      }
    },
  },
  mounted() {
    window.addEventListener("keyup", this.keyHandler);
    this.boxesGET();
  },
  unmounted() {
    window.removeEventListener("keyup", this.keyHandler);
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
</style>
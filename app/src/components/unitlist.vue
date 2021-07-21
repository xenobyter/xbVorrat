<template>
  <div id="unitlist">
    <h1>Einheiten</h1>
    <ul>
      <li v-for="unit in units" :key="unit.id">
        <div class="unit">
          <div class="unitLeft">{{ unit.id }}</div>
          <div class="unitMain">
            <div class="unitName">{{ unit.unit }}</div>
            <div class="unitNotiz">{{ unit.long }}</div>
          </div>
          <div class="unitRight">
            <button v-on:click="unitEDIT(unit)" class="edit">
              <img
                src="/icons/edit_black_24dp.svg"
                alt="&#128393;"
                class="icon"
              />
            </button>
            <button v-on:click="unitDELETE(unit)" class="delete">
              <img
                src="/icons/delete_black_24dp.svg"
                alt="&#128465;"
                class="icon"
              />
            </button>
          </div>
        </div>
      </li>
      <li>
        <div class="unit unitRight">
          <button v-on:click="unitAdd()" class="add">
            <img src="/icons/add_black_24dp.svg" alt="&#10750;" class="icon" />
          </button>
        </div>
      </li>
      <unitedit
        v-if="showedit"
        v-on:closeedit="showedit = false"
        v-on:doedit="unitPATCH"
        v-on:doadd="unitPUT"
        :unit="unit"
      />
    </ul>
    <div class="status">{{ status }}</div>
  </div>
</template>

<script>
import axios from "axios";
import unitedit from "./unitedit.vue";

export default {
  name: "unitlist",
  emits: ["nav"],
  components: {
    unitedit,
  },
  data() {
    return {
      api: process.env.VUE_APP_API,
      units: [
        { id: 1, unit: "kg", long: "Kilogramm" },
        { id: 2, unit: "l", long: "Liter" },
      ],
      showedit: false,
      unit: Object,
      status: "",
    };
  },
  methods: {
    unitsGET() {
      axios
        .get(this.api + "/units", { timeout: 900 })
        .then((response) => {
          this.units = response.data;
        })
        .catch((e) => {
          console.log("GET", e);
          this.status = "GET: " + e.message;
        });
    },
    unitDELETE(unit) {
      this.unit = unit;
      const index = this.units.indexOf(this.unit);
      axios
        .delete(this.api + "/units/" + unit.id, { timeout: 900 })
        .then((response) => {
          console.log("Status:", response.status);
          this.units.splice(index, 1);
        })
        .catch((e) => {
          console.error("DELETE", e.message);
          this.status = "DELETE: " + e.message;
        });
    },
    unitEDIT(unit) {
      this.unit = unit;
      this.showedit = true;
    },
    unitPATCH(unit) {
      const index = this.units.indexOf(this.unit);
      this.showedit = false;
      axios
        .patch(
          this.api + "/units/" + unit.id,
          {
            unit: unit.unit,
            long: unit.long,
          },
          { timeout: 900 }
        )
        .then((response) => {
          this.units[index] = { ...unit };
          console.log("Status:", response.status);
        })
        .catch((e) => {
          console.log("PATCH", e);
          this.status = "PATCH: " + e.message;
        });
    },
    unitAdd() {
      this.unit = { unit: "", long: "" };
      this.showedit = true;
    },
    unitPUT(unit) {
      this.showedit = false;
      console.log(unit);
      axios
        .put(
          this.api + "/units",
          {
            unit: unit.unit,
            long: unit.long,
          },
          { timeout: 900 }
        )
        .then((response) => {
          console.log("Status:", response.status);
          this.unitsGET();
        })
        .catch((e) => {
          console.log("PUT", e);
          this.status = "PUT: " + e.message;
        });
    },
    keyHandler(e) {
      if (this.showedit) return;
      switch (e.key) {
        case "+":
          this.unitAdd();
          break;
        case "1":
          this.$emit("nav", "");
          break;
      }
    },
  },
  mounted() {
    window.addEventListener("keyup", this.keyHandler);
    this.unitsGET();
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
.unit {
  overflow: hidden;
  display: flex;
  background-color: var(--listen-hintergrund);
}
.unitLeft,
.unitMain {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.unitLeft {
  padding: 0 0.5rem 0 0;
  min-width: 2rem;
  max-width: 4rem;
  text-align: right;
}
.unitMain {
  flex: 1;
}
.unitNotiz {
  font-size: small;
  color: rgb(160, 160, 160);
  min-height: 1.2rem;
}
.unitRight {
  justify-content: flex-end;
}
</style>
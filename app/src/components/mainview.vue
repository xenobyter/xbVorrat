<template>
  <div class="main">
    <h1>Vorrat</h1>
    <table>
      <tr>
        <th v-on:click="sort('articlestr')">
          Name
          <img
            v-if="sortBy == 'articlestr' && order == 'asc'"
            src="/icons/sort_black_24dp.svg"
            style="transform: scaleY(-1)"
            alt="Aufsteigend"
          />
          <img
            v-if="sortBy == 'articlestr' && order == 'desc'"
            src="/icons/sort_black_24dp.svg"
            alt="Absteigend"
          />
          <img
            v-if="sortBy != 'articlestr'"
            src="/icons/sort_black_24dp.svg"
            alt="Ausgeblendet"
            style="visibility: hidden"
          />
        </th>
        <th v-on:click="sort('boxstr')">
          Box
          <img
            v-if="sortBy == 'boxstr' && order == 'asc'"
            src="/icons/sort_black_24dp.svg"
            style="transform: scaleY(-1)"
            alt="Aufsteigend"
          />
          <img
            v-if="sortBy == 'boxstr' && order == 'desc'"
            src="/icons/sort_black_24dp.svg"
            alt="Absteigend"
          />
          <img
            v-if="sortBy != 'boxstr'"
            src="/icons/sort_black_24dp.svg"
            alt="Ausgeblendet"
            style="visibility: hidden"
          />
        </th>
        <th v-on:click="sort('expiry')">
          Datum
          <img
            v-if="sortBy == 'expiry' && order == 'asc'"
            src="/icons/sort_black_24dp.svg"
            style="transform: scaleY(-1)"
            alt="Aufsteigend"
          />
          <img
            v-if="sortBy == 'expiry' && order == 'desc'"
            src="/icons/sort_black_24dp.svg"
            alt="Absteigend"
          />
          <img
            v-if="sortBy != 'expiry'"
            src="/icons/sort_black_24dp.svg"
            alt="Ausgeblendet"
            style="visibility: hidden"
          />
        </th>
        <th colspan="4">Menge</th>
        <th></th>
      </tr>
      <tr v-for="stock in stocks" v-bind:key="stock.id">
        <td>{{ stock.articlestr }}</td>
        <td>{{ stock.boxstr }}</td>
        <td>{{ stock.expiry }}</td>
        <td class="nopad">{{ stock.quantity }}</td>
        <td class="nopad">&#215;</td>
        <td class="nopad">{{ stock.size }}</td>
        <td>{{ stock.unitstr }}</td>
        <td class="action">
          <button v-on:click="stockEDIT(stock)" class="edit">
            <img
              src="/icons/edit_black_24dp.svg"
              alt="&#128393;"
              class="icon"
            />
          </button>
          <button v-on:click="stockDELETE(stock)" class="delete">
            <img
              src="/icons/delete_black_24dp.svg"
              alt="&#128465;"
              class="icon"
            />
          </button>
        </td>
      </tr>
      <tr>
        <td colspan="8" class="action">
          <button v-on:click="stockADD()" class="add">
            <img src="/icons/add_black_24dp.svg" alt="&#10750;" class="icon" />
          </button>
        </td>
      </tr>
      <tr v-for="stock in stocks" v-bind:key="stock.id"></tr>
    </table>
    <div class="status">{{ status }}</div>
    <stockedit
      v-if="showedit"
      v-on:closeedit="showedit = false"
      v-on:doedit="stockPATCH"
      v-on:doadd="stockPUT"
      :stock="stock"
    />
  </div>
</template>

<script>
import axios from "axios";
import stockedit from "./stockedit.vue";

export default {
  name: "mainview",
  emits: ["nav"],
  components: {
    stockedit,
  },
  data() {
    return {
      api: process.env.VUE_APP_API,
      stocks: [
        {
          id: 1,
          article: 1,
          articlestr: "Mehl",
          box: 1,
          boxstr: "Box1",
          size: 0.5,
          unitstr: "kg",
          quantity: 2,
          expiry: "31.12.2021",
        },
        {
          id: 2,
          article: 1,
          articlestr: "Mehl2",
          box: 1,
          boxstr: "Box1",
          size: 0.5,
          unitstr: "kg",
          quantity: 2,
          expiry: "31.12.2021",
        },
      ],
      stock: Object,
      status: "",
      showedit: false,
      sortBy: "id",
      order: "asc",
      tblHeader: { articlestr: "Name" },
    };
  },
  methods: {
    stocksGET() {
      axios
        .get(
          this.api +
            "/stocks" +
            "?sort=" +
            this.sortBy +
            "&order=" +
            this.order,
          { timeout: 900 }
        )
        .then((response) => {
          this.stocks = response.data;
        })
        .catch((e) => {
          console.log("GET", e);
          this.status = `GET: ${e.message}`;
        });
    },
    stockEDIT(stock) {
      this.stock = stock;
      this.showedit = true;
    },
    stockPATCH(stock) {
      this.showedit = false;
      axios
        .patch(
          this.api + "/stocks/" + stock.id,
          {
            article: stock.article,
            box: stock.box,
            size: stock.size,
            quantity: stock.quantity,
            expiry: stock.expiry,
          },
          { timeout: 900 }
        )
        .then((response) => {
          console.log("Status:", response.status);
          this.stocksGET();
        })
        .catch((e) => {
          console.log("PATCH", e);
          this.status = `PATCH: ${e.message}`;
        });
    },
    stockADD() {
      this.stock = { articlestr: "" };
      this.showedit = true;
    },
    stockPUT(stock) {
      this.showedit = false;
      axios
        .put(
          this.api + "/stocks",
          {
            article: stock.article,
            box: stock.box,
            size: stock.size,
            quantity: stock.quantity,
            expiry: stock.expiry,
          },
          { timeout: 900 }
        )
        .then((response) => {
          console.log("Status:", response.status);
          this.stocksGET();
        })
        .catch((e) => {
          console.log("PUT", e);
          this.status = `PUT: ${e.message}`;
        });
    },
    stockDELETE(stock) {
      if (confirm(`${stock.articlestr} aus ${stock.boxstr} löschen?`)) {
        axios
          .delete(this.api + "/stocks/" + stock.id, { timeout: 900 })
          .then((response) => {
            console.log("Status:", response.status);
            this.stocksGET();
          })
          .catch((e) => {
            console.error("DELETE", e.message);
            this.status = "DELETE: " + e.message;
          });
      }
    },
    keyHandler(e) {
      if (this.showedit) return
      switch (e.key) {
        case "+":
          this.stockADD();
          break;
        case "2":
          this.$emit("nav", "boxlist");
          break;
        case "3":
          this.$emit("nav", "unitlist");
          break;
        case "4":
          this.$emit("nav", "articlelist");
          break;
        case "5":
          this.$emit("nav", "information");
          break;
      }
    },
    sort(arg) {
      if (this.sortBy != arg) {
        this.sortBy = arg;
      } else {
        this.order = this.order == "asc" ? "desc" : "asc";
      }
      this.stocksGET();
    },
  },
  mounted() {
    window.addEventListener("keyup", this.keyHandler);
    this.stocksGET();
  },
  unmounted() {
    window.removeEventListener("keyup", this.keyHandler);
  },
};
</script>

<style scoped>
table {
  width: calc(100vw - 3rem);
  border-collapse: collapse;
  margin-bottom: 2rem;
}
td,
th {
  border-top: 1px solid black;
  border-bottom: 1px solid black;
  padding: 0 0.1rem 0 0.1rem;
  font-size: medium;
  text-align: left;
}
td:first-child,
th:first-child {
  border-left: 1px solid black;
}
td:last-child,
th:last-child {
  border-right: 1px solid black;
}
td.action {
  text-align: right;
}
td.nopad {
  padding: 0;
  margin: 0;
  width: 0.6rem;
}
th img {
  vertical-align: middle;
}
@media screen and (max-width: 320px) {
  td,
  th {
    font-size: x-small;
  }
}
</style>
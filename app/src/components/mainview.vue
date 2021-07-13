<template>
  <div class="main">
    <h1>Vorrat</h1>
    <table>
      <tr>
        <th>Name</th>
        <th>Box</th>
        <th>Datum</th>
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
          <a v-on:click="stockEDIT(stock)" class="edit">&#128393;</a>
          <a v-on:click="stockDELETE(stock)" class="delete">&#128465;</a>
        </td>
      </tr>
      <tr>
        <td colspan="8" class="action">
          <a v-on:click="stockADD()" class="add">&#10750;</a>
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
      ],
      stock: Object,
      status: "",
      showedit: false,
    };
  },
  methods: {
    stocksGET() {
      axios
        .get(this.api + "/stocks/rich", { timeout: 900 })
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
      console.log(stock);
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
  },
  mounted() {
    this.stocksGET();
  },
};
</script>

<style scoped>
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
table {
  width: calc(100vw - 2.5rem);
  border-collapse: collapse;
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
@media screen and (max-width: 306px) {
  td,
  th {
    font-size: x-small;
  }
}
</style>
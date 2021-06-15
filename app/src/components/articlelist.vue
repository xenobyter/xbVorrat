<template>
  <div id="articlelist">
    <h1>Artikel</h1>
    <ul>
      <li v-for="article in articles" :key="article.id">
        <div class="article">
          <div class="articleLeft">{{ article.id }}</div>
          <div class="articleMain">
            <div class="articleName">{{ article.name }}</div>
            <div class="articleNotiz">{{ article.long }}</div>
          </div>
          <div class="articleRight">
            <a v-on:click="articleEDIT(article)" class="edit">&#128393;</a>
            <a v-on:click="articleDELETE(article)" class="delete">&#128465;</a>
          </div>
        </div>
      </li>
      <li>
        <div class="article articleRight">
          <a v-on:click="articleAdd()" class="add">&#10750;</a>
        </div>
      </li>
      <articleedit
        v-if="showedit"
        v-on:closeedit="showedit = false"
        v-on:doedit="articlePATCH"
        v-on:doadd="articlePUT"
        :article="article"
        :units="units"
      />
    </ul>
    <div class="status">{{ status }}</div>
  </div>
</template>

<script>
import axios from "axios";
import articleedit from "./articleedit.vue";

export default {
  name: "articlelist",
  components: {
    articleedit,
  },
  data() {
    return {
      articles: [
        {
          id: 1,
          name: "Testartikel 1",
          unit: 1,
          long: "Kilogramm",
          short: "kg",
        },
        {
          id: 2,
          name: "Testartikel 2",
          unit: 2,
          long: "Liter",
          short: "l",
        },
      ],
      units: [
        { id: 1, unit: "kg", long: "Kilogramm" },
        { id: 2, unit: "l", long: "Liter" },
      ],
      showedit: false,
      article: Object,
      status: "",
    };
  },
  methods: {
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
    unitsGET() {
      axios
        .get("http://localhost:8081/api/units", { timeout: 900 })
        .then((response) => {
          this.units = response.data;
          this.untisMerge();
        })
        .catch((e) => {
          console.log("GET", e);
          this.status = "GET: " + e.message;
        });
    },
    untisMerge() {
      this.articles.forEach((article) => {
        const unit = this.units.find((u) => {
          return u.id == article.unit;
        });
        article.long = unit.long;
        article.short = unit.unit;
      });
    },
    articleDELETE(article) {
      this.article = article;
      const index = this.articles.indexOf(this.article);
      axios
        .delete("http://localhost:8081/api/articles/" + article.id, {
          timeout: 900,
        })
        .then((response) => {
          console.log("Status:", response.status);
          this.articles.splice(index, 1);
        })
        .catch((e) => {
          console.error("DELETE", e.message);
          this.status = "DELETE: " + e.message;
        });
    },
    articleEDIT(article) {
      this.article = article;
      this.showedit = true;
    },
    articlePATCH(article) {
      const index = this.articles.indexOf(this.article);
      this.showedit = false;
      axios
        .patch(
          "http://localhost:8081/api/articles/" + article.id,
          {
            name: article.name,
            unit: article.unit,
          },
          { timeout: 900 }
        )
        .then((response) => {
          article.long = this.getUnitById(article.unit).long
          this.articles[index] = { ...article };
          console.log("Status:", response.status);
        })
        .catch((e) => {
          console.log("PATCH", e);
          this.status = "PATCH: " + e.message;
        });
    },
    articleAdd() {
      this.article = { name: "", unit: "" };
      this.showedit = true;
    },
    articlePUT(article) {
      this.showedit = false;
      axios
        .put(
          "http://localhost:8081/api/articles",
          {
            name: article.name,
            unit: article.unit,
          },
          { timeout: 900 }
        )
        .then((response) => {
          console.log("Status:", response.status);
          article.long = this.getUnitById(article.unit).long;
          this.articles.push(article);
          const index = this.articles.indexOf(article);
          this.articles[index].id = response.data.id;
        })
        .catch((e) => {
          console.log("PUT", e);
          this.status = "PUT: " + e.message;
        });
    },
    getUnitById(id) {
      const unit = this.units.find((u) => {
        return u.id == id;
      });
      return unit;
    },
  },
  mounted() {
    this.articlesGET();
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
.article {
  overflow: hidden;
  display: flex;
  background-color: var(--listen-hintergrund);
}
.articleLeft,
.articleMain {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.articleLeft {
  padding: 0 0.5rem 0 0;
  min-width: 2rem;
  max-width: 4rem;
  text-align: right;
}
.articleMain {
  flex: 1;
}
.articleNotiz {
  font-size: small;
  color: rgb(160, 160, 160);
  min-height: 1.2rem;
}
.articleRight {
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
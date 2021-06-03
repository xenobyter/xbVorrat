<template>
  <div class="row">
    <input type="checkbox" id="hamburg" v-on:click="checked = !checked" :checked="checked"/>
    <label for="hamburg" class="hamburg">
      <span class="line"></span>
      <span class="line"></span>
      <span class="line"></span>
    </label>
  </div>
  <div class="content" v-if="checked">
    <a v-on:click="show('')">Start</a>
    <a v-on:click="show('boxlist')">Boxen</a>
    <a v-on:click="show('unitlist')">Einheiten</a>
  </div>
</template>

<script>
export default {
  name: "navigation",
  emits: ["nav"],
  data() {
    return {
      checked: false,
    };
  },
  methods: {
    show(item) {
      this.checked =false;
      this.$emit("nav", item);
    },
  },
};
</script>

<style scoped >
label.hamburg {
  display: block;
  background: var(--listen-hintergrund);
  width: 55px;
  height: 50px;
  position: relative;
  margin-right: auto;
  border-radius: 4px;
}
@media only screen and (min-width: 270px) {
  label.hamburg {
    position: absolute;
    margin-top: -0.75rem;
  }
}
input#hamburg {
  display: none;
}
.line {
  position: absolute;
  left: 10px;
  height: 4px;
  width: 35px;
  background: var(--text-color);
  border-radius: 2px;
  display: block;
  transition: 0.5s;
  transform-origin: center;
}
.line:nth-child(1) {
  top: 12px;
}
.line:nth-child(2) {
  top: 24px;
}
.line:nth-child(3) {
  top: 36px;
}

#hamburg:checked + .hamburg .line:nth-child(1) {
  transform: translateY(12px) rotate(-45deg);
}
#hamburg:checked + .hamburg .line:nth-child(2) {
  opacity: 0;
}
#hamburg:checked + .hamburg .line:nth-child(3) {
  transform: translateY(-12px) rotate(45deg);
}

.content {
  position: absolute;
  background-color: var(--listen-hintergrund);
  z-index: 1;
  margin-top: 60px;
  height: auto;
  border: 1px solid;
  box-shadow: 5px 5px 5px;
  display: flex;
  flex-direction: column;
}

.content a {
  padding: 0 0.5rem 0.2rem 0.5rem;
}

.content a:nth-child(1) {
  padding-top: 0.2rem;
}
</style>
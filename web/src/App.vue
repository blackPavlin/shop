<template>
  <el-container id="app-container">
    <el-header>
      <nav-bar />
    </el-header>
    <router-view />
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import NavBar from "./components/NavBar.vue";

export default defineComponent({
  name: "AppView",
  components: {
    NavBar,
  },
});
</script>

<style lang="scss">
body,
html {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

body,
html,
#app,
#app-container {
  margin: 0;
  padding: 0;
  min-height: 100vh;
}
</style>

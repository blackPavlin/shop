<template>
  <el-menu :router="true" :default-active="$route.path">
    <el-menu-item>
      <el-icon></el-icon>
      <template #title></template>
    </el-menu-item>

    <el-menu-item>
      <el-icon></el-icon>
      <template #title></template>
    </el-menu-item>

    <el-menu-item>
      <el-icon></el-icon>
      <template #title></template>
    </el-menu-item>
  </el-menu>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "CategorySideBar",
  components: {},
  setup() {
    return {};
  },
});
</script>

<template>
  <el-container class="products">
    <el-aside width="250px">
      <category-side-bar />
    </el-aside>
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CategorySideBar from "@/components/CategorySideBar.vue";

export default defineComponent({
  name: "ProductsView",
  components: {
    CategorySideBar,
  },
  setup() {
    return {};
  },
});
</script>

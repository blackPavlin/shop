<template>
  <el-menu class="admin-sidebar" :router="true" :default-active="$route.path">
    <el-menu-item index="/admin/category" route="/admin/category">
      <el-icon></el-icon>
      <template #title>Управление категориями</template>
    </el-menu-item>

    <el-menu-item index="/admin/product" route="/admin/product">
      <el-icon></el-icon>
      <template #title>Управление товарами</template>
    </el-menu-item>

    <el-menu-item index="/admin/order" route="/admin/order">
      <el-icon></el-icon>
      <template #title>Управление заказами</template>
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

<style lang="scss" scoped>
.admin-sidebar {
  height: 100%;

  .el-menu-item {
    user-select: none;
  }
}
</style>

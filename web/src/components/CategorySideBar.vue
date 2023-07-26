<template>
  <el-menu class="category-menu" :router="true" :default-active="$route.path">
    <el-menu-item
      v-for="category in categories"
      :key="category.id"
      :index="`/products/${category.id}`"
      :route="`/products/${category.id}`"
      >{{ category.name }}</el-menu-item
    >
  </el-menu>
</template>

<script lang="ts">
import { defineComponent, onMounted, computed } from "vue";
import { useCategoryStore } from "@/store/category.store";

export default defineComponent({
  name: "CategorySideBar",
  setup() {
    const categoryStore = useCategoryStore();
    const categories = computed(() => categoryStore.getCategories);

    onMounted(async () => {
      try {
        await categoryStore.loadCategories();
      } catch (error) {
        console.log(error);
      }
    });

    return {
      categories,
    };
  },
});
</script>

<style lang="scss" scoped>
.category-menu {
  height: 100%;
}

// .category-menu__item {}
</style>

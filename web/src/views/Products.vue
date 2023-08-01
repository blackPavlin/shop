<template>
  <el-container class="products">
    <el-aside width="250px">
      <category-side-bar />
    </el-aside>
    <el-main>
      <el-space wrap>
        <product-cart
          v-for="product of products"
          :key="product.id"
          :product="product"
        ></product-cart>
      </el-space>
    </el-main>
  </el-container>
</template>

<script lang="ts">
import { defineComponent, computed, onMounted } from "vue";
import { ElNotification } from "element-plus";
import { useProductStore } from "@/store/product.store";
import CategorySideBar from "@/components/CategorySideBar.vue";
import ProductCart from "@/components/ProductCart.vue";

export default defineComponent({
  name: "ProductsView",
  components: {
    CategorySideBar,
    ProductCart,
  },
  setup() {
    const productStore = useProductStore();
    const products = computed(() => productStore.getProducts);

    const loadProducts = async (): Promise<void> => {
      return productStore.loadProducts().catch((error) => {
        ElNotification.error(error.message);
      });
    };

    onMounted(async () => loadProducts());

    return {
      products,
    };
  },
});
</script>

<template>
  <div class="product-list">
    <el-scrollbar>
      <el-space wrap>
        <product-cart v-for="i in 70" :key="i"></product-cart>
      </el-space>
    </el-scrollbar>

    <!-- <div class="pagination">
      <el-pagination background layout="prev, pager, next" :total="100" />
    </div> -->
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ProductCart from "@/components/ProductCart.vue";

export default defineComponent({
  name: "ProductList",
  components: {
    ProductCart,
  },
  setup() {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.product-list {
  height: 100%;
}

.pagination {
  display: flex;
  justify-content: center;
}
</style>

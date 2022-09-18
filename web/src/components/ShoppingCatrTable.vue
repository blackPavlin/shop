<template>
  <div class="shopping-cart-table">
    <el-table :data="products">
      <el-table-column label="Название"></el-table-column>

      <el-table-column label="Цена"></el-table-column>

      <el-table-column label="Количество"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ShoppingCartTable",
  props: {
    products: {
      type: Array,
      required: true,
    },
  },
  setup() {
    return {};
  },
});
</script>

<template>
  <div class="shopping-cart-table">
    <el-table :data="products">
      <el-table-column label="Название"></el-table-column>

      <el-table-column label="Цена"></el-table-column>

      <el-table-column label="Количество" prop="amount"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { PropType } from "vue";
import type { CartList } from "@/api/models/CartList";

export default defineComponent({
  name: "ShoppingCartTable",
  props: {
    products: {
      type: Array as PropType<CartList>,
      required: true,
    },
  },
  setup() {
    return {};
  },
});
</script>

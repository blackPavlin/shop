<template>
  <el-main class="cart">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>Корзина</span>
        </div>

        <el-button-group>
          <el-button type="success" @click="createOrder">
            Заказать
            <el-icon><Check /></el-icon>
          </el-button>
          <el-button type="danger" @click="clearCart">
            Очистить
            <el-icon class="el-icon--right"><Delete /></el-icon>
          </el-button>
        </el-button-group>
      </template>

      <shopping-catr-table :products="products" />
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { Delete, Check } from "@element-plus/icons-vue";
import { ElMessageBox, ElMessage } from "element-plus";
import ShoppingCatrTable from "@/components/ShoppingCatrTable.vue";

import { useCartStore } from "@/store/cart.store";

export default defineComponent({
  name: "CartView",
  components: {
    Delete,
    Check,
    ShoppingCatrTable,
  },
  setup() {
    const cartStore = useCartStore();

    const products = computed(() => cartStore.getCartProducts);

    const createOrder = () => {
      console.log("");
    };

    const clearCart = () => {
      ElMessageBox.confirm("Вы уверены, что хотите очистить корзину?")
        .then(async () => {
          console.log("");
        })
        .catch(() => {
          ElMessage({
            type: "info",
            message: "Удаление отменено",
          });
        });
    };

    return {
      products,
      createOrder,
      clearCart,
    };
  },
});
</script>

<style lang="scss">
.el-card__header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>

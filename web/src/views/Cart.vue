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

      <el-table :data="products">
        <el-table-column label="Название" prop="product.name"></el-table-column>
        <el-table-column label="Цена" prop="product.price"></el-table-column>
        <el-table-column label="Количество" prop="amount"></el-table-column>
        <el-table-column>
          <template #default="scope">
            <el-button
              type="primary"
              :icon="Edit"
              @click="openUpdateModal(scope.row.id)"
            />
            <el-button
              type="danger"
              :icon="Delete"
              @click="deleteCartProduct(scope.row.id)"
            />
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, computed, onMounted, ref } from "vue";
import { Edit, Delete, Check } from "@element-plus/icons-vue";
import { ElMessageBox, ElMessage, ElNotification } from "element-plus";
import { useCartStore } from "@/store/cart.store";

export default defineComponent({
  name: "CartView",
  components: {
    Delete,
    Check,
  },
  setup() {
    const cartStore = useCartStore();

    onMounted(() => {
      cartStore.loadCart().catch((error) => {
        ElNotification.error(error.message);
      });
    });

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

    const isUpdateDialogVisible = ref(false);

    const openUpdateModal = (cartId: number) => {
      console.log(cartId);
      isUpdateDialogVisible.value = true;
    };

    // const updateCartProduct = () => {};

    const deleteCartProduct = (cartId: number) => {
      ElMessageBox.confirm("Вы уверены, что хотите удалить товар?", "Warning", {
        confirmButtonText: "Удалить",
        cancelButtonText: "Отмена",
        type: "warning",
      })
        .then(() => {
          cartStore
            .deleteCart(String(cartId))
            .then(async () => {
              ElNotification.success("Товар успешно удалён");
              await cartStore.loadCart();
            })
            .catch((error) => {
              ElNotification.error(error.message);
            });
        })
        .catch(() => {
          ElNotification.info("Удаление отменено");
        });
    };

    return {
      products,
      createOrder,
      clearCart,
      isUpdateDialogVisible,
      openUpdateModal,
      deleteCartProduct,

      Edit,
      Delete,
    };
  },
});
</script>

<style lang="scss">
.el-card__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>

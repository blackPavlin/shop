<template>
  <el-menu mode="horizontal" :router="true" :default-active="$route.path">
    <el-menu-item index="/products" route="/products">
      <el-icon><goods /></el-icon>
      <template #title>Товары</template>
    </el-menu-item>

    <el-menu-item index="/registration" route="/registration" v-if="!isAuth">
      <el-icon><user /></el-icon>
      <template #title>Регистрация</template>
    </el-menu-item>

    <el-menu-item index="/login" route="/login" v-if="!isAuth">
      <el-icon><arrow-left /></el-icon>
      <template #title>Вход</template>
    </el-menu-item>

    <el-menu-item index="/cart" route="/cart" v-if="isAuth">
      <el-icon><shopping-cart /></el-icon>
      <template #title>Корзина</template>
    </el-menu-item>

    <el-menu-item index="/orders" route="/orders" v-if="isAuth">
      <el-icon><sold-out /></el-icon>
      <template #title>Заказы</template>
    </el-menu-item>

    <el-menu-item index="/logout" route="/logout" v-if="isAuth">
      <el-icon><arrow-right /></el-icon>
      <template #title>Выход</template>
    </el-menu-item>
  </el-menu>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import {
  ArrowLeft,
  ArrowRight,
  ShoppingCart,
  SoldOut,
  Goods,
  User,
} from "@element-plus/icons-vue";
import { useAuthStore } from "@/store/auth.store";

export default defineComponent({
  name: "NavBar",
  components: {
    ArrowLeft,
    ArrowRight,
    ShoppingCart,
    SoldOut,
    Goods,
    User,
  },
  setup() {
    const store = useAuthStore();
    const isAuth = computed(() => store.isAuth);

    return { isAuth };
  },
});
</script>

<style lang="scss" scoped>
.el-menu--horizontal {
  justify-content: flex-end;
}

.el-menu--horizontal {
  .el-menu-item {
    user-select: none;
  }
}
</style>

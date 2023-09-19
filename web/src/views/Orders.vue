<template>
  <el-main class="options">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>Заказы</span>
        </div>
      </template>
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "OrdersView",
});
</script>

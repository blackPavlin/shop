<template>
  <el-main class="address">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>Адреса доставки</span>
        </div>

        <el-button-group>
          <el-button type="primary">Добавить</el-button>
        </el-button-group>
      </template>

      <el-table :data="addresses"></el-table>
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, onMounted, computed } from "vue";
import { ElNotification } from "element-plus";
import { useAddressStore } from "@/store/address.store";

export default defineComponent({
  name: "AddressView",
  setup() {
    const addressStore = useAddressStore();

    onMounted(() => {
      addressStore.loadAddresses().catch((error) => {
        ElNotification.error(error.message);
      });
    });

    const addresses = computed(() => addressStore.getAddresses);

    return {
      addresses,
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

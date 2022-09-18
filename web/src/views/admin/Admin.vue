<template>
  <el-container class="admin">
    <el-aside>
      <admin-side-bar />
    </el-aside>
    <el-main></el-main>
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import AdminSideBar from "@/components/admin/AdminSideBar.vue";

export default defineComponent({
  name: "AdminView",
  components: {
    AdminSideBar,
  },
  setup() {
    return {};
  },
});
</script>

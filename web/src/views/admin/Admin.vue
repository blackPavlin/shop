<template>
  <el-container class="admin">
    <el-aside width="250px">
      <admin-side-bar />
    </el-aside>
    <router-view />
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

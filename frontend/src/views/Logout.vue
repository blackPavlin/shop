<template>
  <el-main class="logout"></el-main>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount } from "vue";
import { useAuthStore } from "@/store/auth.store";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "LogoutView",
  setup() {
    const store = useAuthStore();
    const router = useRouter();

    onBeforeMount(() => {
      store.clearToken();
      router.push("/");
    });
  },
});
</script>

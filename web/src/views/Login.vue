<template>
  <el-main class="login">
    <el-card>
      <h2>Login</h2>

      <el-form
        class="login-form"
        :model="form"
        :rules="rules"
        ref="formRef"
        @submit.prevent="login(formRef)"
      >
        <el-form-item prop="email">
          <el-input
            v-model="form.email"
            placeholder="Email"
            :prefix-icon="User"
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="form.password"
            placeholder="Password"
            :prefix-icon="Lock"
            type="password"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            class="login-button"
            type="primary"
            :loading="loading"
            native-type="submit"
            >Login</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import { FormRules, FormInstance, ElNotification } from "element-plus";
import { User, Lock } from "@element-plus/icons-vue";
import { useAuthStore } from "@/store/auth.store";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "LoginView",
  setup() {
    const form = reactive({
      email: "",
      password: "",
    });

    const rules = reactive<FormRules>({
      email: [
        {
          type: "email",
          required: true,
          message: "Email is required",
          trigger: "blur",
        },
      ],
      password: [
        { required: true, message: "Password is required", trigger: "blur" },
      ],
    });

    const formRef = ref<FormInstance>();

    const loading = ref(false);

    const store = useAuthStore();
    const router = useRouter();

    const login = async (formEl?: FormInstance) => {
      if (!formEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }

        loading.value = true;
        store
          .login({
            email: form.email,
            password: form.password,
          })
          .then(() => {
            router.push("/");
          })
          .catch((error) => {
            ElNotification.warning(error.message);
          })
          .finally(() => {
            loading.value = false;
          });
      });
    };

    return { form, rules, formRef, loading, login, User, Lock };
  },
});
</script>

<style lang="scss" scoped>
h2 {
  font-family: "Open Sans";
  letter-spacing: 1px;
  font-family: Roboto, sans-serif;
  padding-bottom: 20px;
}

.login {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-button {
  width: 100%;
  margin-top: 40px;
}

.login-form {
  width: 290px;
}

.login .el-card {
  width: 340px;
  display: flex;
  justify-content: center;
}

.login .el-input input {
  padding-left: 35px;
}
</style>

<template>
  <el-main class="registration">
    <el-card>
      <h2>Registration</h2>

      <el-form
        class="registration-form"
        :model="form"
        :rules="rules"
        ref="formRef"
        @submit.prevent="registration(formRef)"
      >
        <el-form-item prop="name">
          <el-input
            v-model="form.name"
            placeholder="Name"
            :prefix-icon="User"
          ></el-input>
        </el-form-item>

        <el-form-item prop="phone">
          <el-input
            v-model="form.phone"
            placeholder="Phone"
            :prefix-icon="Phone"
          ></el-input>
        </el-form-item>

        <el-form-item prop="address">
          <el-input
            v-model="form.address"
            placeholder="Address"
            :prefix-icon="AddLocation"
          ></el-input>
        </el-form-item>

        <el-form-item prop="email">
          <el-input
            v-model="form.email"
            placeholder="Email"
            :prefix-icon="Message"
          ></el-input>
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="form.password"
            placeholder="Password"
            :prefix-icon="Lock"
            type="password"
          ></el-input>
        </el-form-item>

        <el-form-item prop="confirmPassword">
          <el-input
            v-model="form.confirmPassword"
            placeholder="Confirm Password"
            :prefix-icon="Lock"
            type="password"
          ></el-input>
        </el-form-item>

        <el-form-item>
          <el-button
            class="registration-button"
            type="primary"
            native-type="submit"
            :loading="loading"
            >Registration</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { FormInstance, FormRules, ElNotification } from "element-plus";
import { defineComponent, reactive, ref } from "vue";
import {
  User,
  Lock,
  AddLocation,
  Phone,
  Message,
} from "@element-plus/icons-vue";
import { useAuthStore } from "@/store/auth.store";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "RegistrationView",
  setup() {
    const form = reactive({
      name: "",
      phone: "",
      address: "",
      email: "",
      password: "",
      confirmPassword: "",
    });

    const rules = reactive<FormRules>({
      name: [
        {
          type: "string",
          required: true,
          message: "Name is required",
          trigger: "blur",
        },
      ],
      phone: [
        {
          type: "string",
          required: true,
          message: "Phone is required",
          trigger: "blur",
        },
      ],
      address: [
        {
          type: "string",
          required: true,
          message: "Address is required",
          trigger: "blur",
        },
      ],
      email: [
        {
          type: "email",
          required: true,
          message: "Email is required",
          trigger: "blur",
        },
      ],
      password: [
        {
          type: "string",
          required: true,
          message: "Password is required",
          trigger: "blur",
        },
      ],
      confirmPassword: [],
    });

    const formRef = ref<FormInstance>();

    const loading = ref(false);

    const store = useAuthStore();
    const router = useRouter();

    const registration = (formEl?: FormInstance) => {
      if (!formEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }

        loading.value = true;

        store
          .registration(form)
          .then(() => {
            return store.login(form.email, form.password);
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

    return {
      form,
      rules,
      formRef,
      loading,
      registration,
      User,
      Lock,
      AddLocation,
      Phone,
      Message,
    };
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

.registration {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.registration-button {
  width: 100%;
  margin-top: 40px;
}

.registration-form {
  width: 290px;
}

.registration .el-card {
  width: 340px;
  display: flex;
  justify-content: center;
}

.registration .el-input input {
  padding-left: 35px;
}
</style>

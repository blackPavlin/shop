<template>
  <el-main class="admin-product">
    <el-card>
      <template #header>
        <span>Управление товарами</span>
      </template>

      <el-card>
        <template #header>
          <span>Создать товар</span>
        </template>

        <el-form
          :model="form"
          :rules="rules"
          ref="formRef"
          @submit.prevent="createProduct(formRef)"
          label-width="100px"
          size="large"
        >
          <el-form-item prop="categoryId" label="Категория">
            <el-select v-model="form.categoryId" placeholder="Категория">
              <el-option
                v-for="category in categories"
                :key="category.id"
                :label="category.name"
                :value="category.id"
              ></el-option>
            </el-select>
          </el-form-item>

          <el-form-item prop="name" label="Название">
            <el-input v-model="form.name" placeholder="Название"></el-input>
          </el-form-item>

          <el-form-item prop="description" label="Описание">
            <el-input
              v-model="form.description"
              placeholder="Описание"
              type="textarea"
            ></el-input>
          </el-form-item>

          <el-form-item prop="amount" label="Количество">
            <el-input-number
              v-model="form.amount"
              placeholder="Количество"
              :min="0"
            ></el-input-number>
          </el-form-item>

          <el-form-item prop="price" label="Цена">
            <el-input-number
              v-model="form.price"
              placeholder="Цена"
              :min="0"
            ></el-input-number>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" native-type="submit">Создать</el-button>
          </el-form-item>
        </el-form>
      </el-card>

      <el-card>
        <template #header>
          <span>Товары</span>
        </template>

        <el-form size="large"></el-form>

        <el-table :data="products">
          <el-table-column label="ID" prop="id"></el-table-column>

          <el-table-column label="Название" prop="name"></el-table-column>
        </el-table>
      </el-card>
    </el-card>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, computed, onMounted } from "vue";
import { FormRules, FormInstance, ElNotification } from "element-plus";
import { Plus } from "@element-plus/icons-vue";
import { useCategoryStore } from "@/store/category.store";
import { useProductStore } from "@/store/product.store";

export default defineComponent({
  name: "ProductAdmin",
  components: {},
  setup() {
    const form = reactive({
      categoryId: 0,
      name: "",
      description: "",
      amount: 0,
      price: 0,
    });

    const rules = reactive<FormRules>({
      categoryId: [
        {
          required: true,
          message: "Category is required",
          trigger: "blur",
        },
      ],
      name: [
        {
          required: true,
          message: "Name is required",
          trigger: "blur",
        },
        {
          max: 100,
          message: "Max length 100",
          trigger: "blur",
        },
      ],
      description: [],
      amount: [
        {
          required: true,
          message: "Amount is required",
          trigger: "blur",
        },
      ],
      price: [
        {
          required: true,
          message: "Price is required",
          trigger: "blur",
        },
      ],
    });

    const formRef = ref<FormInstance>();

    const categoryStore = useCategoryStore();
    const productStore = useProductStore();
    const categories = computed(() => categoryStore.getCategories);
    const products = computed(() => productStore.getProducts);

    const loadCategories = async (): Promise<void> => {
      return categoryStore.loadCategories().catch((error) => {
        ElNotification.error(error.message);
      });
    };
    const loadProducts = async (): Promise<void> => {
      return productStore.loadProducts().catch((error) => {
        ElNotification.error(error.message);
      });
    };

    const createProduct = (formEl?: FormInstance): void => {
      if (!formEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }

        productStore
          .createProduct(form)
          .then(async () => {
            formEl.resetFields();
            ElNotification.success("Product successful created");

            await loadProducts();
          })
          .catch((error) => {
            ElNotification.error(error.message);
          });
      });
    };

    onMounted(() => loadCategories());
    onMounted(() => loadProducts());

    return {
      form,
      rules,
      formRef,
      categories,
      products,
      createProduct,
      Plus,
    };
  },
});
</script>

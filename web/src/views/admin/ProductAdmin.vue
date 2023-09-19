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
          @submit.prevent="createProduct(formRef, uploadRef)"
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

          <el-form-item label="Изображения">
            <el-upload
              v-model:file-list="fileList"
              ref="uploadRef"
              list-type="picture-card"
              :auto-upload="false"
              multiple
            >
              <el-icon><Plus /></el-icon>
            </el-upload>
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

        <el-table :data="products">
          <el-table-column label="ID" prop="id"></el-table-column>
          <el-table-column label="Название" prop="name"></el-table-column>
          <el-table-column label="Количество" prop="amount"></el-table-column>
          <el-table-column label="Цена" prop="price"></el-table-column>
          <el-table-column>
            <template #default="scope">
              <el-button
                type="primary"
                :icon="Edit"
                @click="openUpdateModal(scope.row.id)"
              ></el-button>
              <el-button
                type="danger"
                :icon="Delete"
                @click="deleteProduct(scope.row.id)"
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </el-card>

    <el-dialog title="Редактирование товара" v-model="isUpdateDialogVisible">
      <el-form
        :model="updateForm"
        :rules="rules"
        ref="updateRef"
        @submit.prevent="updateProduct(updateRef)"
        label-width="100px"
        size="large"
      >
        <el-form-item prop="name" label="Название">
          <el-input v-model="updateForm.name" placeholder="Название"></el-input>
        </el-form-item>

        <el-form-item prop="description" label="Описание">
          <el-input
            v-model="updateForm.description"
            placeholder="Описание"
            type="textarea"
          ></el-input>
        </el-form-item>

        <el-form-item prop="amount" label="Количество">
          <el-input-number
            v-model="updateForm.amount"
            placeholder="Количество"
            :min="0"
          ></el-input-number>
        </el-form-item>

        <el-form-item prop="price" label="Цена">
          <el-input-number
            v-model="updateForm.price"
            placeholder="Цена"
            :min="0"
          ></el-input-number>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" native-type="submit">Сохранить</el-button>
          <el-button @click="isUpdateDialogVisible = false">Отмета</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, computed, onMounted } from "vue";
import {
  FormRules,
  FormInstance,
  UploadInstance,
  ElNotification,
  ElMessageBox,
  UploadUserFile,
  UploadRawFile,
} from "element-plus";
import { Plus, Delete, Edit } from "@element-plus/icons-vue";
import { useCategoryStore } from "@/store/category.store";
import { useProductStore } from "@/store/product.store";
import { useImageStore } from "@/store/image.storage";
import { CreateProductRequest } from "@/api/models/CreateProductRequest";
import { Product } from "@/api/models/Product";

export default defineComponent({
  name: "ProductAdmin",
  components: {
    Plus,
  },
  setup() {
    const form = reactive<CreateProductRequest>({
      categoryId: 0,
      name: "",
      description: "",
      amount: 0,
      price: 0,
    });

    const updateForm = reactive<Product>({
      id: 0,
      categoryId: 0,
      name: "",
      description: "",
      amount: 0,
      price: 0,
      images: [],
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
    const updateRef = ref<FormInstance>();
    const uploadRef = ref<UploadInstance>();

    const fileList = ref<UploadUserFile[]>();

    const categoryStore = useCategoryStore();
    const productStore = useProductStore();
    const imageStore = useImageStore();

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

    const createProduct = (
      formEl?: FormInstance,
      uploadEl?: UploadInstance,
    ): void => {
      if (!formEl || !uploadEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }

        productStore
          .createProduct(form)
          .then(async (product) => {
            console.log(fileList.value);

            if (fileList.value && fileList.value.length) {
              const files = fileList.value
                .map((file) => file.raw)
                .filter((file): file is UploadRawFile => Boolean(file));
              await imageStore.uploadImages(String(product.id), { files });
            }

            uploadEl.clearFiles();
            formEl.resetFields();
            ElNotification.success("Product successful created");

            await loadProducts();
          })
          .catch((error) => {
            ElNotification.error(error.message);
          });
      });
    };

    const isUpdateDialogVisible = ref(false);

    const openUpdateModal = (productId: number) => {
      productStore
        .getProduct(String(productId))
        .then((product) => {
          updateForm.id = product.id;
          updateForm.categoryId = product.categoryId;
          updateForm.name = product.name;
          updateForm.description = product.description;
          updateForm.amount = product.amount;
          updateForm.price = product.price;
          updateForm.images = product.images;

          isUpdateDialogVisible.value = true;
        })
        .catch((error) => {
          ElNotification.error(error.message);
        });
    };

    const updateProduct = (formEl?: FormInstance) => {
      if (!formEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }
      });

      productStore
        .updateProduct(String(updateForm.id), updateForm)
        .then(async () => {
          formEl.resetFields();
          isUpdateDialogVisible.value = false;
          ElNotification.success("Товар успешно изменён");

          await loadProducts();
        })
        .catch((error) => {
          ElNotification.error(error.message);
        });
    };

    const deleteProduct = (productId: number) => {
      ElMessageBox.confirm("Вы уверены, что хотите удалить товар?", "Warning", {
        confirmButtonText: "Удалить",
        cancelButtonText: "Отмена",
        type: "warning",
      })
        .then(() => {
          productStore
            .deleteProduct(String(productId))
            .then(async () => {
              ElNotification.success("Товар успешно удалён");
              await loadProducts();
            })
            .catch((error) => {
              ElNotification.error(error.message);
            });
        })
        .catch(() => {
          ElNotification.info("Удаление отменено");
        });
    };

    onMounted(() => loadCategories());
    onMounted(() => loadProducts());

    return {
      form,
      updateForm,
      rules,
      formRef,
      updateRef,
      uploadRef,
      fileList,
      categories,
      products,
      isUpdateDialogVisible,
      createProduct,
      openUpdateModal,
      updateProduct,
      deleteProduct,

      Delete,
      Edit,
    };
  },
});
</script>

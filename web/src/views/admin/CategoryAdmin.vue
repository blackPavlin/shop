<template>
  <el-main class="admin-category">
    <el-card>
      <template #header>
        <span>Управление категориями</span>
      </template>

      <el-card>
        <template #header>
          <span>Создать категорию</span>
        </template>

        <el-form
          :model="form"
          :rules="rules"
          ref="formRef"
          @submit.prevent="createCategory(formRef)"
          size="large"
        >
          <el-form-item prop="name">
            <el-input
              v-model="form.name"
              placeholder="Название категории"
            ></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" native-type="submit">Создать</el-button>
          </el-form-item>
        </el-form>
      </el-card>

      <el-card>
        <template #header>
          <span>Категории</span>
        </template>

        <el-form size="large">
          <el-form-item>
            <el-input placeholder="Название категории"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" native-type="submit">Найти</el-button>
          </el-form-item>
        </el-form>

        <el-table :data="categories">
          <el-table-column label="ID" prop="id"></el-table-column>

          <el-table-column label="Название" prop="name"></el-table-column>

          <el-table-column>
            <template #default="scope">
              <el-button
                type="primary"
                :icon="Edit"
                @click="openUpdateModal(scope.row.id, scope.row.name)"
              ></el-button>

              <el-button
                type="danger"
                :icon="Delete"
                @click="deleteCategory(scope.row.id)"
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </el-card>

    <el-dialog title="Редактирование категории" v-model="isUpdateDialogVisible">
      <el-form
        :model="updateForm"
        :rules="updateRules"
        ref="updateFormRef"
        @submit.prevent.stop="updateCategory(updateFormRef)"
        size="large"
      >
        <el-form-item prop="name">
          <el-input
            v-model="updateForm.name"
            placeholder="Новое название категории"
          ></el-input>
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
  ElNotification,
  ElMessageBox,
} from "element-plus";
import { Delete, Edit } from "@element-plus/icons-vue";
import { useCategoryStore } from "@/store/category.store";
import { CreateCategoryRequest } from "@/api/models/CreateCategoryRequest";
import { Category } from "@/api/models/Category";

export default defineComponent({
  name: "CategoryAdmin",
  components: {},
  setup() {
    const form = reactive<CreateCategoryRequest>({
      name: "",
    });

    const updateForm = reactive<Category>({
      id: 0,
      name: "",
    });

    const rules = reactive<FormRules>({
      name: [
        {
          required: true,
          message: "Name is required",
          trigger: "blur",
        },
        {
          min: 3,
          message: "Min length 3",
          trigger: "blur",
        },
        {
          max: 50,
          message: "Max length 50",
          trigger: "blur",
        },
      ],
    });

    const updateRules = reactive<FormRules>({
      name: [
        {
          required: true,
          message: "Name is required",
          trigger: "blur",
        },
        {
          min: 3,
          message: "Min length 3",
          trigger: "blur",
        },
        {
          max: 50,
          message: "Max length 50",
          trigger: "blur",
        },
      ],
    });

    const formRef = ref<FormInstance>();
    const updateFormRef = ref<FormInstance>();

    const store = useCategoryStore();
    const categories = computed(() => store.getCategories);

    const loadCategories = async (): Promise<void> => {
      return store.loadCategories().catch((error) => {
        ElNotification.error(error.message);
      });
    };

    const createCategory = (formEl?: FormInstance): void => {
      if (!formEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }

        store
          .createCategory(form)
          .then(async () => {
            formEl.resetFields();
            ElNotification.success("Category successful created");

            await loadCategories();
          })
          .catch((error) => {
            ElNotification.error(error.message);
          });
      });
    };

    const isUpdateDialogVisible = ref(false);

    const openUpdateModal = (categoryId: number, name: string) => {
      updateForm.id = categoryId;
      updateForm.name = name;

      isUpdateDialogVisible.value = true;
    };

    const updateCategory = (formEl?: FormInstance): void => {
      if (!formEl) {
        return;
      }

      formEl.validate((valid) => {
        if (!valid) {
          return;
        }

        store
          .updateCategory(updateForm)
          .then(async () => {
            formEl.resetFields();
            isUpdateDialogVisible.value = false;
            ElNotification.success("Категория успешно изменена");

            await loadCategories();
          })
          .catch((error) => {
            ElNotification.error(error.message);
          });
      });
    };

    const deleteCategory = (id: number): void => {
      ElMessageBox.confirm(
        "Вы уверены, что хотите удалить категорию?",
        "Warning",
        {
          confirmButtonText: "Удалить",
          cancelButtonText: "Отмена",
          type: "warning",
        },
      )
        .then(() => {
          store
            .deleteCategory(id)
            .then(async () => {
              ElNotification.success("Категория успешно удалена");
              await loadCategories();
            })
            .catch((error) => {
              ElNotification.error(error.message);
            });
        })
        .catch(() => {
          ElNotification.info("Удаление отменено");
        });
    };

    onMounted(async () => loadCategories());

    return {
      form,
      rules,
      formRef,
      updateForm,
      updateRules,
      updateFormRef,
      categories,
      isUpdateDialogVisible,
      createCategory,
      openUpdateModal,
      updateCategory,
      deleteCategory,

      Delete,
      Edit,
    };
  },
});
</script>

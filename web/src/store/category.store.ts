import { defineStore } from "pinia";
import { Category } from "@/api/models/Category";
import { CategoryList } from "@/api/models/CategoryList";
import { CreateCategoryRequest } from "@/api/models/CreateCategoryRequest";
import { CategoryService } from "@/api/services/CategoryService";

type State = {
  categories: CategoryList;
};

export const useCategoryStore = defineStore("category", {
  state: (): State => ({
    categories: [],
  }),
  actions: {
    async loadCategories(): Promise<void> {
      this.categories = await CategoryService.getCategory();
    },
    async createCategory(body: CreateCategoryRequest): Promise<void> {
      await CategoryService.postCategory(body);
    },
    async updateCategory(body: Category): Promise<void> {
      await CategoryService.patchCategory(body);
    },
    async deleteCategory(categoryId: number): Promise<void> {
      await CategoryService.deleteCategory(categoryId);
    },
  },
  getters: {
    getCategories: (state): CategoryList => state.categories,
  },
});

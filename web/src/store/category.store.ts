import { defineStore } from "pinia";
import { CategoryService, CategoryList, CreateCategoryRequest } from "@/api";

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
  },
  getters: {
    getCategories: (state): CategoryList => state.categories,
  },
});

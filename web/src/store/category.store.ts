import { defineStore } from "pinia";
import { CategoryService } from "@/api/services/CategoryService";

import { CategoryList } from "@/api/models/CategoryList";

type State = {
  categories: CategoryList;
};

export const useCategoryStore = defineStore("category", {
  state: (): State => ({
    categories: [],
  }),
  actions: {
    async loadCategories(): Promise<void> {
      const categories = await CategoryService.getCategory();
      this.categories = categories;
    },
  },
  getters: {
    getCategories: (state): CategoryList => state.categories,
  },
});

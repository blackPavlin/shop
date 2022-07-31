import { defineStore } from "pinia";
import client from "@/plugins/api.plugin";

type State = {
  categories: string[];
};

interface GetCategoriesResponse {
  categories: string[];
}

export const useCategoryStore = defineStore("category", {
  state: (): State => ({
    categories: [],
  }),
  actions: {
    async loadCategories(): Promise<void> {
      const { data } = await client.get<GetCategoriesResponse>("/category");
      this.categories = data.categories;
    },
  },
  getters: {
    getCategories: (state): string[] => state.categories,
  },
});

import { defineStore } from "pinia";

type State = {
  products: Product[];
};

interface Product {
  id: string;
}

export const useProductStore = defineStore("product", {
  state: (): State => ({
    products: [],
  }),
  actions: {},
  getters: {},
});

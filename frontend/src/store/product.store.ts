import { defineStore } from "pinia";
import client from "@/plugins/api.plugin";

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

import { defineStore } from "pinia";
import { Product } from "@/api/models/Product";
import { ProductList } from "@/api/models/ProductList";
import { ProductService } from "@/api/services/ProductService";
import { CreateProductRequest } from "@/api/models/CreateProductRequest";

type State = {
  products: ProductList;
};

export const useProductStore = defineStore("product", {
  state: (): State => ({
    products: [],
  }),
  actions: {
    async loadProducts(): Promise<void> {
      this.products = await ProductService.getProduct();
    },
    async createProduct(body: CreateProductRequest): Promise<void> {
      await ProductService.postProduct(body);
    },
  },
  getters: {
    getProducts: (state) => state.products,
  },
});

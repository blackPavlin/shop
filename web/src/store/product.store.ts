import { defineStore } from "pinia";
import { Product } from "@/api/models/Product";
import { ProductService } from "@/api/services/ProductService";
import { CreateProductRequest } from "@/api/models/CreateProductRequest";

type State = {
  products: Product[];
  quantity: number;
};

export const useProductStore = defineStore("product", {
  state: (): State => ({
    products: [],
    quantity: 0,
  }),
  actions: {
    async loadProducts(): Promise<void> {
      const { items, quantity } = await ProductService.getProduct();
      this.products = items;
      this.quantity = quantity;
    },
    async getProduct(productId: string): Promise<Product> {
      return ProductService.getProduct1(productId);
    },
    async createProduct(body: CreateProductRequest): Promise<Product> {
      return ProductService.postProduct(body);
    },
    async updateProduct(
      productId: string,
      body: CreateProductRequest,
    ): Promise<Product> {
      return ProductService.patchProduct(productId, body);
    },
    async deleteProduct(productId: string): Promise<void> {
      await ProductService.deleteProduct(productId);
    },
  },
  getters: {
    getProducts: (state) => state.products,
  },
});

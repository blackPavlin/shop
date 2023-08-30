import { defineStore } from "pinia";
import { CartService } from "@/api/services/CartService";
import type { CartList } from "@/api/models/CartList";

type State = {
  products: CartList;
};

export const useCartStore = defineStore("cart", {
  state: (): State => ({
    products: [],
  }),
  actions: {
    async loadCart(): Promise<void> {
      this.products = await CartService.getCart();
    },
  },
  getters: {
    getCartProducts(): CartList {
      return this.products;
    },
  },
});

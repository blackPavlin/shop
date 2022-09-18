import { defineStore } from "pinia";
import { CartService } from "@/api/services/CartService";

import { CartProductList } from "@/api/models/CartProductList";

type State = {
  products: CartProductList;
};

export const useCartStore = defineStore("cart", {
  state: (): State => ({
    products: [],
  }),
  actions: {
    async loadCart(): Promise<void> {
      const cart = await CartService.getCart();
      this.products = cart.products;
    },
  },
  getters: {
    getCartProducts(): CartProductList {
      return this.products;
    },
  },
});

import { defineStore } from "pinia";
import client from "@/plugins/api.plugin";

type State = {
  products: CartProduct[];
};

export interface Cart {
  products: CartProduct[];
}

export interface CartProduct {
  product_id: string;
}

interface GetCartResponse {
  cart: Cart;
}

export const useCartStore = defineStore("cart", {
  state: (): State => ({
    products: [],
  }),
  actions: {
    async loadCart(): Promise<void> {
      const { data } = await client.get<GetCartResponse>("/cart");
      this.products = data.cart.products;
    },
  },
  getters: {
    getCartProducts(): CartProduct[] {
      return this.products;
    },
  },
});

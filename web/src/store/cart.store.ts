import { defineStore } from "pinia";
import { CartService } from "@/api/services/CartService";
import type { Cart, CartList, CartProduct } from "@/api";

type State = {
  products: CartList;
  productsAmountByProductID: Map<number, number>;
};

export const useCartStore = defineStore("cart", {
  state: (): State => ({
    products: [],
    productsAmountByProductID: new Map<number, number>(),
  }),
  actions: {
    async loadCart(): Promise<void> {
      this.products = await CartService.getCart();
      for (const product of this.products) {
        this.productsAmountByProductID.set(product.product.id, product.amount);
      }
    },
    async updateCart(body: CartProduct): Promise<Cart> {
      return CartService.patchCart(body);
    },
    async deleteCart(cartId: string): Promise<void> {
      await CartService.deleteCart1(cartId);
    },
    async deleteCarts(): Promise<void> {
      await CartService.deleteCart();
    },
  },
  getters: {
    getCartProducts(): CartList {
      return this.products;
    },
    getProductAmountByProductId(): Map<number, number> {
      return this.productsAmountByProductID;
    },
  },
});

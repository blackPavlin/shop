import { defineStore } from "pinia";
import { OrderService } from "@/api/services/OrderService";

import { OrderList } from "@/api/models/OrderList";

type State = {
  orders: OrderList;
};

export const useOrderStore = defineStore("order", {
  state: (): State => ({
    orders: [],
  }),
  actions: {
    async loadOrders(): Promise<void> {
      const orders = await OrderService.getOrder();
      this.orders = orders;
    },
  },
  getters: {
    getOrders(): OrderList {
      return this.orders;
    },
  },
});

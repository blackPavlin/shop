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
      this.orders = await OrderService.getOrder();
    },
  },
  getters: {
    getOrders(): OrderList {
      return this.orders;
    },
  },
});

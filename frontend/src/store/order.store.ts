import { defineStore } from "pinia";
import client from "@/plugins/api.plugin";

type State = {
  orders: Order[];
};

interface Order {
  id: string;
}

interface GetOrdersResponse {
  orders: Order[];
}

export const useOrderStore = defineStore("order", {
  state: (): State => ({
    orders: [],
  }),
  actions: {
    async loadOrders(): Promise<void> {
      const { data } = await client.get<GetOrdersResponse>("/order");
      this.orders = data.orders;
    },
  },
  getters: {
    getOrders(): Order[] {
      return this.orders;
    },
  },
});

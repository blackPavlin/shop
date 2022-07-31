import { defineStore } from "pinia";
import client from "@/plugins/api.plugin";

type State = {
  user: User | null;
};

interface User {
  name: string;
  phone: string;
  address: string;
  email: string;
  role: "admin" | "user";
}

interface GetUserResponse {
  user: User;
}

export const useUserStore = defineStore("user", {
  state: (): State => ({
    user: null,
  }),
  actions: {
    async loadUser(): Promise<void> {
      const { data } = await client.get<GetUserResponse>("/user");
      this.user = data.user;
    },
  },
  getters: {
    getUser(): User | null {
      return this.user;
    },
  },
});

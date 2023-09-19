import { defineStore } from "pinia";
import { UserService } from "@/api/services/UserService";
import type { User } from "@/api";

type State = {
  user: User | null;
};

export const useUserStore = defineStore("user", {
  state: (): State => ({
    user: null,
  }),
  actions: {
    async loadUser(): Promise<void> {
      this.user = await UserService.getUser();
    },
  },
  getters: {
    getUser(): User | null {
      return this.user;
    },
  },
});

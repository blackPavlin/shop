import { defineStore } from "pinia";
import { UserService } from "@/api/services/UserService";

import { User } from "@/api/models/User";

type State = {
  user: User | null;
};

export const useUserStore = defineStore("user", {
  state: (): State => ({
    user: null,
  }),
  actions: {
    async loadUser(): Promise<void> {
      const user = await UserService.getUser();
      this.user = user;
    },
  },
  getters: {
    getUser(): User | null {
      return this.user;
    },
  },
});

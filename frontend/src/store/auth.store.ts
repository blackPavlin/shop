import { defineStore } from "pinia";

type State = {
  token: string | null;
};

export const useAuthStore = defineStore("auth", {
  state: (): State => ({
    token: localStorage.getItem("token"),
  }),
  actions: {
    setToken(token: string): void {
      localStorage.setItem("token", token);
    },
    clearToken(): void {
      localStorage.removeItem("token");
    },
  },
  getters: {
    isAuth: (state): boolean => Boolean(state.token),
    getToken: (state): string | null => state.token,
  },
});

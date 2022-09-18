import { defineStore } from "pinia";
import { OpenAPI } from "@/api";
import { AuthorizationService } from "@/api/services/AuthorizationService";

import { LoginRequest } from "@/api/models/LoginRequest";
import { SignupRequest } from "@/api/models/SignupRequest";

OpenAPI.TOKEN = localStorage.getItem("token") || "";

type State = {
  token: string | null;
};

export const useAuthStore = defineStore("auth", {
  state: (): State => ({
    token: localStorage.getItem("token"),
  }),
  actions: {
    async login(body: LoginRequest): Promise<void> {
      const { token } = await AuthorizationService.postAuthLogin(body);
      this.setToken(token);
    },
    async registration(body: SignupRequest): Promise<void> {
      const { token } = await AuthorizationService.postAuthSignup(body);
      this.setToken(token);
    },
    setToken(token: string): void {
      this.token = token;
      localStorage.setItem("token", token);
    },
    clearToken(): void {
      this.token = null;
      localStorage.removeItem("token");
    },
  },
  getters: {
    isAuth: (state): boolean => Boolean(state.token),
    getToken: (state): string | null => state.token,
  },
});

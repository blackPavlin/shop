import { defineStore } from "pinia";
import client from "@/plugins/api.plugin";

type State = {
  token: string | null;
};

type LoginResponse = {
  token: string;
};

type RegistarationRequest = {
  name: string;
  phone: string;
  address: string;
  email: string;
  password: string;
};

type RegistrationResponse = {
  success: boolean;
};

export const useAuthStore = defineStore("auth", {
  state: (): State => ({
    token: localStorage.getItem("token"),
  }),
  actions: {
    async login(email: string, password: string): Promise<void> {
      const { data } = await client.post<LoginResponse>("/auth/signin", {
        email,
        password,
      });

      this.setToken(data.token);
    },
    async registration(
      body: RegistarationRequest
    ): Promise<RegistrationResponse> {
      const { data } = await client.post<RegistrationResponse>(
        "/auth/signup",
        body
      );

      return data;
    },
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

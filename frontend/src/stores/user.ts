import { defineStore } from "pinia";
import { siteAPI } from "./api";
import { User } from "@/internal/types/user";
import { Failure } from "@/internal/entity/failure";
import { Result } from "neverthrow";

interface profile {
  user: User;
}

export const useUserStore = defineStore("user", {
  state: () => ({
    user: null as User | null,
    init: false,
  }),
  getters: {},
  actions: {
    async auth(login: string, password: string): Promise<Result<any, Failure>> {
      return siteAPI.post("/user/auth", {
        login: login,
        password: password,
      });
    },
    async profile() {
      const res = await siteAPI.get("/user/profile");
      res.match(
        (r: profile) => {
          this.user = r.user;
        },
        (err: Failure) => {}
      );

      this.init = true;
    },
    async logout() {
      return siteAPI.post("/user/logout", {});
    },
  },
});

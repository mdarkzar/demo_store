import { Failure } from "@/internal/entity/failure";
import { Notification } from "@/internal/types/notification";
import { defineStore } from "pinia";
import { useToast } from "vue-toastification";
import { siteAPI } from "./api";

const toast = useToast();

interface messageList {
  messageList: Array<Notification> | null;
}

export const useNotificationStore = defineStore("notification", {
  state: () => ({
    messageList: null as Array<Notification> | null,
  }),
  getters: {},
  actions: {
    async loadMessages() {
      const res = await siteAPI.get("/notification/new");
      res.match(
        (r: messageList) => {
          this.messageList = r.messageList;
        },
        (err: Failure) => {
          toast.error(err.message);
        }
      );
    },
  },
});

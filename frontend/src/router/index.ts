import LoginView from "@/views/LoginView.vue";
import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import { useUserStore } from "@/stores/user";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/login",
      name: "Вход в Demo Store",
      component: LoginView,
    },
    {
      path: "/",
      name: "home",
      component: HomeView,
      meta: { title: "Demo Store" },
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();

  if (!userStore.init) {
    await userStore.profile();
  }

  if (to.path !== "/login" && userStore.user == null)
    next({ path: "/login", query: { to: to.path } });
  else next();
});

export default router;

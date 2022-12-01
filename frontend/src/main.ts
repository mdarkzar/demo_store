import { createApp } from "vue";
import { createPinia } from "pinia";

import ElementPlus from "element-plus";
import "font-awesome/css/font-awesome.min.css";
import "mdi-icons/css/materialdesignicons.min.css";
import "element-plus/dist/index.css";
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";
import App from "./App.vue";
import router from "./router";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

import "./assets/main.scss";

const app = createApp(App);

app.use(createPinia());
app.use(ElementPlus, { zIndex: 6000, size: "large" });
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.use(Toast, {
  position: "bottom-center",
  timeout: 5000,
  closeOnClick: true,
  pauseOnFocusLoss: true,
  pauseOnHover: true,
  draggable: true,
  draggablePercent: 0.6,
  showCloseButtonOnHover: false,
  hideProgressBar: true,
  closeButton: "button",
  icon: true,
  rtl: false,
});

// stores

app.use(router);

app.mount("#app");

import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      "/api": { target: "http://127.0.0.1:9000" },
    },
    port: 3000,
  },
  plugins: [vue()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
      "vue-i18n": "vue-i18n/dist/vue-i18n.runtime.esm-bundler.js",
      web3: resolve(__dirname, "./node_modules/web3/dist/web3.min.js"),
      vue: "vue/dist/vue.esm-bundler.js",
    },
  },
});

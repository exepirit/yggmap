import { defineConfig } from "vite";
import preact from "@preact/preset-vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    preact({
      prerender: {
        enabled: false,
        renderTarget: "#app",
        additionalPrerenderRoutes: ["/404"],
      },
    }),
  ],
  server: {
    proxy: {
      "/api": "http://127.0.0.1:8080",
    },
  },
});

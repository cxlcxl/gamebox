import { defineConfig, loadEnv } from "vite"
import vue from "@vitejs/plugin-vue"
import path from "path"

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), "")
  return {
    plugins: [vue()],
    resolve: {
      alias: {
        "~bootstrap": path.resolve(__dirname, "node_modules/bootstrap"),
        "@": path.resolve(__dirname, "src"),
        "@v": path.resolve(__dirname, "src/views"),
      },
    },
    define: {
      "process.env": env,
    },
    server: {
      port: 8001,
      hot: true,
      proxy: {
        "/api/": {
          target: process.env.VUE_APP_BASE_API,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ""),
        },
      },
    },
  }
})

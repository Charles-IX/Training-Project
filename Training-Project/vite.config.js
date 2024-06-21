import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  // plugins: [vue()],
  // base: './',
  // build: {
  //   outDir: resolve(__dirname, 'dist')
  // }
  plugins: [vue()],
  base: './', // 确保相对路径
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    rollupOptions: {
      input: {
        main: path.resolve(__dirname, 'index.html')
      }
    }
  }
})

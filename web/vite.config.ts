import * as path from 'path'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import pkg from './package.json'

process.env.VITE_APP_VERSION = pkg.version
if (process.env.NODE_ENV === 'production') {
  process.env.VITE_APP_BUILD_EPOCH = new Date().getTime().toString()
  process.env.VITE_BACKEND_URL = '';
} else {
  process.env.VITE_BACKEND_URL = process.env.BACKEND_URL ?? 'http://127.0.0.1:8082';
}
process.env.VITE_LOCALSTORAGE_URL_COUNT = process.env.LOCALSTORAGE_URL_COUNT ?? '5';

export default defineConfig({
  plugins: [
    vue({
      script: {
        refSugar: true,
      },
    }),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})

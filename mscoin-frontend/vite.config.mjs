import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve, dirname } from 'path'
import { fileURLToPath } from 'url'
import { localAcceptanceMockPlugin } from './dev/localAcceptanceMocks.mjs'

// Vite 中获取 __dirname 的替代方案
const __dirname = dirname(fileURLToPath(import.meta.url))

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    localAcceptanceMockPlugin(),
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
      '@components': resolve(__dirname, 'src/components'),
      '@pages': resolve(__dirname, 'src/pages'),
      '@pages-vue3': resolve(__dirname, 'src/pages-vue3'),
      '@assets': resolve(__dirname, 'src/assets'),
      '@config': resolve(__dirname, 'src/config'),
      '@utils': resolve(__dirname, 'src/utils'),
      '@js': resolve(__dirname, 'src/assets/js'),
      'vue': 'vue/dist/vue.esm-bundler.js'
    }
  },
  server: {
    port: 3000,
    host: true,
    proxy: {
      '/uc/asset/contract-transaction/all': {
        target: 'http://127.0.0.1:8086',
        changeOrigin: true
      },
      '/uc/contract-wallet': {
        target: 'http://127.0.0.1:8086',
        changeOrigin: true
      },
      '/uc/contract': {
        target: 'http://127.0.0.1:8086',
        changeOrigin: true
      },
      '/uc': {
        target: 'http://127.0.0.1:8888',
        changeOrigin: true
      },
      '/market': {
        target: 'http://127.0.0.1:8889',
        changeOrigin: true
      },
      '/otc': {
        target: 'http://127.0.0.1:8888',
        changeOrigin: true
      },
      '/crowd': {
        target: 'http://127.0.0.1:8888',
        changeOrigin: true
      },
      '/exchange': {
        target: 'http://127.0.0.1:8890',
        changeOrigin: true
      },
      '/swap': {
        target: 'http://127.0.0.1:8086',
        changeOrigin: true
      },
      '/socket.io': {
        target: 'http://127.0.0.1:8889',
        changeOrigin: true,
        ws: true
      }
    }
  },
  optimizeDeps: {
    include: ['vue', 'vue-router', 'vuex', 'axios', 'moment', 'jquery', 'element-plus']
  },
  build: {
    outDir: 'dist-vite',
    assetsDir: 'static',
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'index.html')
      },
      output: {
        manualChunks: {
          'vendor': ['vue', 'vue-router', 'vuex', 'axios', 'element-plus', '@element-plus/icons-vue'],
          'utils': ['moment', 'jquery']
        }
      }
    }
  },
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true
      }
    }
  },
  define: {
    'process.env': {},
    global: 'globalThis'
  }
})

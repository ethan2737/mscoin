import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue2'
import { resolve, dirname } from 'path'
import { fileURLToPath } from 'url'

// Vite 中获取 __dirname 的替代方案
const __dirname = dirname(fileURLToPath(import.meta.url))

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
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
      '@js': resolve(__dirname, 'src/assets/js')
    }
  },
  server: {
    port: 3000,
    host: true,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
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
          'vendor': ['vue', 'vue-router', 'vuex', 'vue-resource', 'vue-i18n', 'iview'],
          'utils': ['axios', 'moment', 'socket.io-client', 'jquery']
        }
      }
    }
  },
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true,
        modifyVars: {
          hack: `true; @import '~iview/src/styles/index.less';`
        }
      }
    }
  }
})

// Vue 3 入口文件
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import axios from 'axios'

// Vue 3 路由和状态管理
import { createRouter, createWebHashHistory } from 'vue-router'
import { createStore } from 'vuex'

// Vue I18n for Vue 3
import { createI18n } from 'vue-i18n'

// 导入 Vue 3 路由配置
import routes from './config/routes-vue3.js'
import storeConfig from './config/store-vue3.js'
import { runtimeContract } from './config/runtime-vue3'

// 导入语言包
import zh from './assets/lang/zh.js'
import en from './assets/lang/en.js'

// 创建 i18n 实例
const i18n = createI18n({
  legacy: false,
  locale: 'zh',
  fallbackLocale: 'en',
  messages: {
    zh: zh,
    en: en
  }
})

// 导入全局样式
import 'swiper/dist/css/swiper.css'
import './assets/icons/iconfont.css'
// import './assets/icons/style.css' - 文件不存在

import moment from 'moment'
import $ from 'jquery'

// 导入主组件
import App from './App.vue'

// 创建 Vue 3 应用
const app = createApp(App)

// 使用 Element Plus
app.use(ElementPlus)

// 使用 Vue I18n
app.use(i18n)

// 注册 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 配置全局属性
app.config.globalProperties.host = runtimeContract.host
app.config.globalProperties.wshost = runtimeContract.wshost
app.config.globalProperties.api = runtimeContract.api
app.config.globalProperties.$runtime = runtimeContract
app.config.globalProperties.$http = axios
app.config.globalProperties.$moment = moment
app.config.globalProperties.$ = $

// 添加全局过滤器
app.config.globalProperties.toFloor = function(number, scale = 8) {
  if (new Number(number) == 0) {
    return 0
  }
  var __str = number + ''
  if (__str.indexOf('e') > -1 || __str.indexOf('E') > -1) {
    var __num = new Number(number).toFixed(scale + 1)
    __str = __num + ''
    return __str.substring(0, __str.length - 1)
  } else if (__str.indexOf('.') > -1) {
    if (scale == 0) {
      return __str.substring(0, __str.indexOf('.'))
    }
    return __str.substring(0, __str.indexOf('.') + scale + 1)
  } else {
    return __str
  }
}

// 时间格式化过滤器
app.config.globalProperties.$filters = {
  timeFormat: (tick) => moment(tick).format('HH:mm:ss'),
  dateFormat: (tick) => moment(tick).format('YYYY-MM-DD HH:mm:ss'),
  toFixed: (number, scale) => new Number(number).toFixed(scale),
  toPercent: (point) => {
    var str = Number(point * 100).toFixed(1)
    str += '%'
    return str
  }
}

// 创建 Vue Router 5.x 实例
const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 路由守卫
router.beforeEach(() => {
  return true
})

router.afterEach((to, from) => {
  window.scrollTo(0, 0)
})

// 创建 Vuex 4.x store
const store = createStore(storeConfig)

// 提供 store, router 和 i18n 给 Composition API 使用
app.provide('store', store)
app.provide('router', router)
app.provide('i18n', i18n.global)
app.provide('runtime', runtimeContract)

// 挂载应用
app.use(router)
app.use(store)
app.mount('#app')

console.log('Vue 3 app bootstrapped')

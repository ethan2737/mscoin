<template>
  <div :class="pageView">
    <div class="page-content">
      <div class="time_download" style="display: none;">
        <div class="leftwrapper">
          <el-icon class="clock"><Clock /></el-icon>
          <span>{{ formattedTime }}&nbsp;&nbsp;{{ utc }}</span>
        </div>
      </div>
      <div class="layout">
        <div class="layout-ceiling">
          <router-link to="/">
            <div class="layout-logo"></div>
          </router-link>
          <div class="layout-ceiling-main">
            <div class="header_nav">
              <el-menu
                :default-active="activeNav"
                mode="horizontal"
                :ellipsis="false"
                background-color="#172636"
                text-color="#828ea1"
                active-text-color="#f0a70a"
              >
                <router-link to="/">
                  <el-menu-item index="nav-index">{{ $t("header.index") }}</el-menu-item>
                </router-link>
                <router-link to="/exchange">
                  <el-menu-item index="nav-exchange">{{ $t("header.exchange") }}</el-menu-item>
                </router-link>
                <router-link to="/ctc">
                  <el-menu-item index="nav-ctc">{{ $t("header.ctc") }}</el-menu-item>
                </router-link>
                <router-link to="/otc/trade/usdt">
                  <el-menu-item index="nav-otc">{{ $t("header.otc") }}</el-menu-item>
                </router-link>
                <router-link to="/swapindex/1">
                  <el-menu-item index="nav-swapindex">{{ $t("header.swap") }}</el-menu-item>
                </router-link>
                <router-link to="/swapindex/2">
                  <el-menu-item index="nav-secswap">{{ $t("header.secswap") }}</el-menu-item>
                </router-link>
                <router-link to="/activity">
                  <el-menu-item index="nav-lab">{{ $t("header.lab") }}</el-menu-item>
                </router-link>
                <router-link to="/mining">
                  <el-menu-item index="nav-mining">{{ $t("header.mining") }}</el-menu-item>
                </router-link>
                <router-link to="/crowdfunding">
                  <el-menu-item index="nav-crowdfunding">{{ $t("header.crowdfunding") }}</el-menu-item>
                </router-link>
              </el-menu>
            </div>
            <div class="header_nav_mobile_triggle" @click="toggleMemu()">
              <el-icon style="font-size: 26px; color: #cccccc;"><Menu /></el-icon>
            </div>
            <div class="header_nav" style="float:right;margin-left: 10px;">
              <el-dropdown @command="changelanguage" trigger="click">
                <div class="lang-title" style="display: flex; align-items: center; cursor: pointer; height: 50px; line-height: 50px;">
                  <img v-if="lang === '简体中文'" class="lang-img" src="./assets/images/lang-zh.png" alt="中文">
                  <img v-else-if="lang === 'English'" class="lang-img" src="./assets/images/lang-en.png" alt="English">
                </div>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="en">
                      <img src="./assets/images/lang-en.png" style="width: 20px; margin-right: 5px; vertical-align: middle;">ENGLISH
                    </el-dropdown-item>
                    <el-dropdown-item command="zh">
                      <img src="./assets/images/lang-zh.png" style="width: 20px; margin-right: 5px; vertical-align: middle;">简体中文
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>

            <div class="rr login-container">
              <div class="login_register isLogin" v-if="isLogin">
                <div class="mymsg">
                  <router-link to="/uc/safe">{{ $t("header.usercenter") }}</router-link>
                </div>
                <el-dropdown trigger="click">
                  <div class="user-dropdown-trigger">
                    <el-icon class="user-icon"><User /></el-icon>
                    <span class="username-text">{{ member.username }}</span>
                    <el-icon class="arrow-icon"><ArrowDown /></el-icon>
                  </div>
                  <template #dropdown>
                    <el-dropdown-menu class="user-dropdown-menu">
                      <el-dropdown-item class="dropdown-item">
                        <router-link to="/uc/money" class="dropdown-item-link">
                          <el-icon><Coin /></el-icon>
                          <span>{{ $t("header.assetmanage") }}</span>
                        </router-link>
                      </el-dropdown-item>
                      <el-dropdown-item class="dropdown-item">
                        <router-link to="/uc/entrust/current" class="dropdown-item-link">
                          <el-icon><Switch /></el-icon>
                          <span>{{ $t("header.trademanage") }}</span>
                        </router-link>
                      </el-dropdown-item>
                      <el-dropdown-item class="dropdown-item">
                        <router-link to="/uc/contract/entrust/current" class="dropdown-item-link">
                          <el-icon><Switch /></el-icon>
                          <span>{{ $t("header.contractmanage") }}</span>
                        </router-link>
                      </el-dropdown-item>
                      <el-dropdown-item class="dropdown-item">
                        <router-link to="/uc/innovation/orders" class="dropdown-item-link">
                          <el-icon><Switch /></el-icon>
                          <span>{{ $t("header.innovationmanage") }}</span>
                        </router-link>
                      </el-dropdown-item>
                      <el-dropdown-item divided class="dropdown-divider">
                        <div @click="logout" class="dropdown-item-link">
                          <el-icon><SwitchButton /></el-icon>
                          <span>{{ $t("common.logout") }}</span>
                        </div>
                      </el-dropdown-item>
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </div>
              <div class="login_register" v-else>
                <el-menu
                  mode="horizontal"
                  :ellipsis="false"
                  background-color="#172636"
                  text-color="#828ea1"
                  active-text-color="#f0a70a"
                >
                  <el-sub-menu index="2">
                    <template #title>
                      <span>{{ $t("common.loginregister") }}</span>
                    </template>
                    <router-link to="/login" id="login">
                      <el-menu-item index="1-1">{{ $t("common.login") }}</el-menu-item>
                    </router-link>
                    <router-link to="/register" id="register">
                      <el-menu-item index="1-2" style="color: #f0a70a;">{{ $t("common.register") }}</el-menu-item>
                    </router-link>
                  </el-sub-menu>
                </el-menu>
              </div>
            </div>
          </div>
        </div>
      </div>
      <router-view v-if="isRouterAlive"></router-view>
    </div>
    <el-drawer v-model="navDrawerModal" size="40%" direction="rtl" class="header_nav_mobile">
      <el-menu
        :default-active="activeNav"
        mode="vertical"
        background-color="#2b323a"
        text-color="#bdc2ca"
        active-text-color="#f0a70a"
      >
        <router-link to="/">
          <el-menu-item index="nav-index">{{ $t("header.index") }}</el-menu-item>
        </router-link>
        <router-link to="/ctc">
          <el-menu-item index="nav-ctc">{{ $t("header.ctc") }}</el-menu-item>
        </router-link>
        <router-link to="/otc/trade/usdt">
          <el-menu-item index="nav-otc">{{ $t("header.otc") }}</el-menu-item>
        </router-link>
        <router-link to="/activity">
          <el-menu-item index="nav-lab">{{ $t("header.lab") }}</el-menu-item>
        </router-link>
        <router-link to="/mining">
          <el-menu-item index="nav-mining">{{ $t("header.mining") }}</el-menu-item>
        </router-link>
        <router-link to="/crowdfunding">
          <el-menu-item index="nav-crowdfunding">{{ $t("header.crowdfunding") }}</el-menu-item>
        </router-link>
        <router-link to="/partner">
          <el-menu-item index="nav-invite">{{ $t("header.invite") }}</el-menu-item>
        </router-link>
        <router-link to="/notice/item/0">
          <el-menu-item index="nav-service">{{ $t("header.service") }}</el-menu-item>
        </router-link>

        <el-sub-menu index="nav-login" v-if="!isLogin">
          <template #title>
            <span>{{ $t("common.loginregister") }}</span>
          </template>
          <router-link to="/login">
            <el-menu-item index="1-1" style="padding-left: 20px !important;">{{ $t("common.login") }}</el-menu-item>
          </router-link>
          <router-link to="/register">
            <el-menu-item index="1-2" style="padding-left: 20px !important;">{{ $t("common.register") }}</el-menu-item>
          </router-link>
        </el-sub-menu>

        <el-sub-menu index="nav_personal" v-if="isLogin">
          <template #title>
            <span>{{ $t("header.usercenter") }}</span>
          </template>
          <router-link to="/uc/safe">
            <el-menu-item index="nav_safe" style="padding-left: 20px !important;">{{ $t("uc.member.securitysetting") }}</el-menu-item>
          </router-link>
          <router-link to="/uc/money">
            <el-menu-item index="nav_assets" style="padding-left: 20px !important;">{{ $t("header.assetmanage") }}</el-menu-item>
          </router-link>
          <router-link to="/uc/innovation/minings">
            <el-menu-item index="nav_innnovationmanage" style="padding-left: 20px !important;">{{ $t("header.innovationmanage") }}</el-menu-item>
          </router-link>
        </el-sub-menu>

        <div v-if="isLogin" @click="logout" style="margin-top: 5px; padding: 8px 24px 8px 5px; color: #bdc2ca; cursor: pointer;">
          {{ $t("common.logout") }}
        </div>

        <div style="height: 1px; width: 100%; background: rgb(59, 69, 85); margin-top: 10px; margin-bottom: 10px;"></div>

        <el-sub-menu index="lang">
          <template #title>
            <span>{{ languageValue }}</span>
          </template>
          <el-menu-item index="zh" style="padding-left: 20px !important;">
            <img src="./assets/images/lang-zh.png" style="width: 20px; margin-right: 5px; vertical-align: middle;">简体中文
          </el-menu-item>
          <el-menu-item index="en" style="padding-left: 20px !important;">
            <img src="./assets/images/lang-en.png" style="width: 20px; margin-right: 5px; vertical-align: middle;">ENGLISH
          </el-menu-item>
        </el-sub-menu>

        <router-link to="/app">
          <el-menu-item index="nav-appdownload">{{ $t("header.appdownlaod") }}</el-menu-item>
        </router-link>
      </el-menu>
    </el-drawer>
    <div class="shoujiShow">
      <div class="sjShow_content">
        <div>
          <router-link :to="{ path: '/help/list', query: { cate: '0', cateTitle: '新手指南' } }">{{ $t("footer.xszn") }}</router-link>
        </div>
        <div>
          <router-link :to="{ path: '/help/list', query: { cate: '1', cateTitle: '常见问题' } }">{{ $t("footer.cjwt") }}</router-link>
        </div>
        <div>
          <router-link :to="{ path: '/help/list', query: { cate: '2', cateTitle: '交易指南' } }">{{ $t("footer.jyzn") }}</router-link>
        </div>
        <div>
          <router-link :to="{ path: '/help/list', query: { cate: '3', cateTitle: '币种资料' } }">{{ $t("footer.bzzl") }}</router-link>
        </div>
      </div>
    </div>
    <div class="footer">
      <div class="footer_content">
        <div class="footer_left">
          <img src="./assets/images/logo-bottom.png" style="margin: 0;" alt="">
          <p style="letter-spacing: 2px;">{{ $t("footer.gsmc") }}</p>
          <p>Copyright © 2023 - MSZLU.COM All rights reserved.&nbsp;&nbsp;</p>
        </div>
        <div class="footer_right">
          <ul>
            <li class="footer_title">
              <span>{{ $t("footer.gsjj") }}</span>
            </li>
            <li>
              <router-link target="_blank" to="/about">{{ $t("footer.gywm") }}</router-link>
            </li>
            <li>
              <router-link target="_blank" :to="'/helpdetail?cate=6&id=39&cateTitle=其他'">{{ $t("footer.jrwm") }}</router-link>
            </li>
            <li>
              <router-link target="_blank" to="/notice/item/0">{{ $t("footer.notice") }}</router-link>
            </li>
            <li class="wechatclick">
              <el-popover placement="right" width="80" trigger="hover">
                <template #reference>
                  <a href="javascript:;" class="wechat">{{ $t("footer.apidoc") }}</a>
                </template>
                <p style="text-align: center;">{{ $t("footer.zwkf") }}</p>
              </el-popover>
            </li>
          </ul>
          <ul>
            <li class="footer_title">
              <span>{{ $t("footer.bzzx") }}</span>
            </li>
            <li>
              <router-link :to="{ path: '/help/list', query: { cate: '0', cateTitle: '新手指南' } }">{{ $t("footer.xszn") }}</router-link>
            </li>
            <li>
              <router-link :to="{ path: '/help/list', query: { cate: '1', cateTitle: '常见问题' } }">{{ $t("footer.cjwt") }}</router-link>
            </li>
            <li>
              <router-link :to="{ path: '/help/list', query: { cate: '2', cateTitle: '交易指南' } }">{{ $t("footer.jyzn") }}</router-link>
            </li>
            <li>
              <router-link :to="{ path: '/help/list', query: { cate: '3', cateTitle: '币种资料' } }">{{ $t("footer.bzzl") }}</router-link>
            </li>
          </ul>
          <ul>
            <li class="footer_title">
              <span>{{ $t("footer.tkxy") }}</span>
            </li>
            <li>
              <router-link target="_blank" :to="'/helpdetail?cate=5&id=2&cateTitle=条款协议'">{{ $t("footer.mztk") }}</router-link>
            </li>
            <li>
              <router-link target="_blank" :to="'/helpdetail?cate=5&id=3&cateTitle=条款协议'">{{ $t("footer.ystk") }}</router-link>
            </li>
            <li>
              <router-link target="_blank" :to="'/helpdetail?cate=5&id=5&cateTitle=条款协议'">{{ $t("footer.fwtk") }}</router-link>
            </li>
            <li>
              <router-link target="_blank" :to="'/helpdetail?cate=5&id=38&cateTitle=条款协议'">{{ $t("footer.fltk") }}</router-link>
            </li>
          </ul>
          <ul>
            <li class="footer_title">
              <span>{{ $t("footer.lxwm") }}</span>
            </li>
            <li class="wechatclick">
              <el-popover placement="right" width="200" trigger="hover">
                <template #reference>
                  <a href="javascript:;" class="wechat">{{ $t("footer.kfyx") }}</a>
                </template>
                <p style="text-align: center;">service@MSZLU.COM</p>
              </el-popover>
            </li>
            <li class="wechatclick">
              <el-popover placement="right" width="200" trigger="hover">
                <template #reference>
                  <a href="javascript:;" class="wechat">{{ $t("footer.swhz") }}</a>
                </template>
                <p style="text-align: center;">support@MSZLU.COM</p>
              </el-popover>
            </li>
            <li class="wechatclick">
              <el-popover placement="right" width="200" trigger="hover">
                <template #reference>
                  <a href="javascript:;" class="wechat">{{ $t("footer.sbsq") }}</a>
                </template>
                <p style="text-align: center;">list@MSZLU.COM</p>
              </el-popover>
            </li>
            <li class="wechatclick">
              <el-popover placement="right" width="200" trigger="hover">
                <template #reference>
                  <a href="javascript:;" class="wechat">{{ $t("footer.tsjb") }}</a>
                </template>
                <p style="text-align: center;">ceo@MSZLU.COM</p>
              </el-popover>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <el-backtop :bottom="140" />
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 根组件 App.vue
 * 包含全局布局：头部导航、移动端抽屉、底部信息
 */
import { ref, reactive, computed, onMounted, onBeforeUnmount, provide, nextTick, inject } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage, ElMenu, ElSubMenu, ElMenuItem, ElDropdown, ElDropdownMenu, ElDropdownItem, ElDrawer, ElPopover, ElBacktop, ElIcon } from 'element-plus'
import { Clock, Menu, ArrowDown, User, Coin, Switch, SwitchButton } from '@element-plus/icons-vue'
import axios from 'axios'
import { runtimeContract } from './config/runtime-vue3'
import { establishAuthenticatedSession, clearAuthenticatedSession } from './utils/auth-session'

// 使用 vue-i18n
const { t: $t, locale } = useI18n()

// 注入 store, router 和 i18n
const store = inject('store')
const router = inject('router')

const host = ''
const api = runtimeContract.api

// 状态
const isRouterAlive = ref(true)
const pageView = ref('page-view')
const utc = ref(null)
const formattedTime = ref(null)
const navDrawerModal = ref(false)
const timeTimer = ref(null)

// 计算属性
const activeNav = computed(() => store?.state.activeNav || 'nav-index')
const isLogin = computed(() => store?.getters.isLogin || false)
const member = computed(() => store?.getters.member || { username: '' })
const lang = computed(() => store?.state.lang || '简体中文')
const languageValue = computed(() => store?.getters.lang || '简体中文')

// 提供 reload 方法给子组件
const reload = () => {
  isRouterAlive.value = false
  nextTick(() => {
    isRouterAlive.value = true
  })
}

provide('reload', reload)

// 工具方法
const strpo = (str) => {
  if (!str) return ''
  if (str.length > 4) {
    return str.slice(0, 4) + '···'
  }
  return str
}

const toggleMemu = () => {
  navDrawerModal.value = !navDrawerModal.value
}

const isMobile = () => {
  return /phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone/i.test(navigator.userAgent)
}

const getQueryVariable = (key) => {
  const after = window.location.search || window.location.hash.split('?')[1]
  if (after) {
    const reg = new RegExp('(^|&)' + key + '=([^&]*)(&|$)')
    const r = after.match(reg)
    if (r != null) {
      return decodeURIComponent(r[2])
    }
  }
  return null
}

// 语言切换
const changelanguage = (name) => {
  if (name === 'en') {
    store?.commit('setlang', 'English')
    locale.value = 'en'
    reload()
  }
  if (name === 'zh') {
    store?.commit('setlang', '简体中文')
    locale.value = 'zh'
    reload()
  }
}

// 初始化
const initialize = () => {
  store?.commit('navigate', 'nav-index')
  store?.commit('recoveryMember')
  store?.commit('initLang')
  store?.commit('initLoginTimes')

  // 同步 i18n locale
  const currentLang = store?.state.lang || '简体中文'
  locale.value = currentLang === 'English' ? 'en' : 'zh'

  checkLogin()
}

const redirectAfterLogout = () => {
  const currentPath = router?.currentRoute?.value?.path || '/'

  if (currentPath.startsWith('/uc')) {
    router.replace({
      path: '/login',
      query: { returnUrl: currentPath }
    })
    return
  }

  reload()
}

// 退出登录
const logout = async () => {
  const token = localStorage.getItem('TOKEN')
  clearAuthenticatedSession({ storage: localStorage, store })
  navDrawerModal.value = false
  redirectAfterLogout()

  if (token) {
    try {
      await axios.post(`${host}/uc/loginout`, {}, {
        headers: { 'x-auth-token': token }
      })
    } catch (error) {
      console.warn('logout request failed', error)
    }
  }

  ElMessage.success($t('common.logout'))
}

// 检查登录状态
const checkLogin = () => {
  const token = localStorage.getItem('TOKEN')
  axios.post(`${host}/uc/check/login`, {}, {
    headers: { 'x-auth-token': token }
  }).then(response => {
    const result = response.data
    if (result.code === 0 && result.data === false) {
      clearAuthenticatedSession({ storage: localStorage, store })
    }
  }).catch(() => {})
}

// 处理外部 token
const handleExternalToken = () => {
  const token = getQueryVariable('token')
  if (token && token !== 'null' && token !== '') {
    establishAuthenticatedSession({
      loginData: { token },
      axiosInstance: axios,
      memberInfoUrl: `${host}${api.uc.memberInfo}`,
      storage: localStorage,
      tokenKey: runtimeContract.tokenKey,
      memberKey: runtimeContract.memberKey,
      store
    }).catch(() => {})
  }
}

// 路由监听
const handleRouteChange = (to) => {
  pageView.value = 'page-view'

  if (to.path === '/reg') {
    pageView.value = 'page-view2'
    if (!isMobile()) {
      const code = getQueryVariable('code')
      if (code !== undefined && code !== '' && code !== null) {
        router.replace('/register?code=' + code)
      } else {
        router.replace('/register')
      }
    }
  }

  if (to.path === '/exchange' && isMobile()) {
    router.replace('/reg')
  }

  if (to.path === '/app') {
    pageView.value = 'page-view2'
  }

  if (to.path.length > 11 && to.path.substr(0, 9) === '/envelope') {
    pageView.value = 'page-view3'
  }
}

// 页面标题更新
const updatePageTitle = () => {
  switch (activeNav.value) {
    case 'nav-exchange':
      window.document.title = (lang.value === '简体中文' ? '交易中心' : 'Exchange') + ' - MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
      break
    case 'nav-service':
      window.document.title = (lang.value === '简体中文' ? '公告' : 'Announcement') + ' - MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
      break
    case 'nav-about':
      window.document.title = (lang.value === '简体中文' ? '关于' : 'About') + ' - MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
      break
    case 'nav-lab':
      window.document.title = (lang.value === '简体中文' ? '公益创新室' : 'Lab') + ' - MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
      break
    case 'nav-mining':
      window.document.title = (lang.value === '简体中文' ? '矿机' : 'Mining') + ' - MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
      break
    case 'nav-invite':
      window.document.title = (lang.value === '简体中文' ? '全球传递爱' : 'Promotion') + ' - MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
      break
    default:
      window.document.title = 'MSCOIN | MSCOIN 官网 - 全球比特币交易平台 | 全球数字货币交易平台'
      break
  }
}

// 生命周期
onMounted(() => {
  // 初始化时间
  const d = new Date()
  const gmtHours = d.getTimezoneOffset() / 60
  utc.value = 'GMT ' + (gmtHours > 0 ? '-' : '+') + ' ' + String(gmtHours)[1]

  // 启动时间定时器
  timeTimer.value = setInterval(() => {
    formattedTime.value = new Date().getTime()
  }, 1000)

  // 隐藏初始加载层
  const initLoading = document.getElementById('initLoading')
  if (initLoading) {
    document.body.removeChild(initLoading)
  }

  // 初始化
  initialize()
  handleExternalToken()

  // 监听路由变化
  router?.afterEach((to) => {
    handleRouteChange(to)
    updatePageTitle()
  })

  // 初始更新标题
  updatePageTitle()
})

onBeforeUnmount(() => {
  if (timeTimer.value) {
    clearInterval(timeTimer.value)
  }
})
</script>

<style scoped lang="scss">
// 样式保持与原 App.vue 一致，仅替换 iView 类名为 Element Plus
@media screen and (max-width: 768px) {
  .header_nav_mobile_triggle {
    display: inline-block !important;
  }
  .footer_content {
    padding: 70px 2% 85px 5%;
  }
  .page-view, .page-view2 {
    .page-content {
      .layout {
        height: 45px;
        .layout-ceiling {
          padding: 5px 10px !important;
          .layout-ceiling-main {
            height: 35px !important;
            line-height: 35px !important;
          }
          .layout-logo {
            width: 200px !important;
            height: 35px !important;
          }
        }
      }
    }
  }
}

.header_nav_mobile_triggle {
  display: none;
  float: right;
  padding: 0 5px 0 20px;
}

.page-view2 .nav-pdf {
  color: #333;
  font-size: 14px;
}

.nav-pdf {
  font-size: 14px;
  color: #fff;
}

.page-view {
  height: 100%;

  .page-content {
    .time_download {
      padding: 0 80px;
      height: 35px;
      background-color: #000;
      line-height: 35px;
      overflow: hidden;

      .leftwrapper {
        float: left;

        .clock {
          display: inline-block;
          vertical-align: middle;
          color: #fff;
        }

        span {
          color: #fff;
          line-height: 35px;
          font-size: 12px;
        }
      }
    }

    .layout {
      width: 100%;
      position: absolute;
      z-index: 10;

      .layout-ceiling {
        padding: 5px 20px;

        .layout-logo {
          width: 300px;
          height: 48px;
          background: url(./assets/images/logo.png) no-repeat;
          background-size: 100% 100%;
          float: left;
          position: absolute;
          z-index: 10;
        }

        .layout-ceiling-main {
          height: 50px;
          line-height: 50px;
          margin-left: 218px;

          .header_nav {
            // 移除 router-link 默认下划线
            & :deep(a) {
              text-decoration: none !important;
            }

            & :deep(.el-menu) {
              background: transparent;
              border-bottom: none;

              &.el-menu--vertical {
                background: transparent;
              }

              .el-sub-menu__title {
                &:hover {
                  background: transparent !important;
                }
              }

              .el-menu-item {
                font-size: 16px;
                font-weight: 500;
                color: #e0e0e0 !important;
                text-decoration: none !important;
                border-bottom: none !important;

                &:hover {
                  background: transparent !important;
                  color: #fff !important;
                  text-decoration: none !important;
                  border-bottom: none !important;
                }

                &.is-active {
                  color: #f0a70a !important;
                }
              }

              // 激活项的容器边框下划线（非文字下划线）
              .el-menu-item.is-active::after {
                content: '';
                position: absolute;
                bottom: 0;
                left: 0;
                right: 0;
                height: 3px;
                background: #ffa800;
              }
            }
          }

          .rr {
            float: right;
            z-index: 10;

            .mymsg {
              float: left;
              padding-right: 20px;

              a {
                color: #e0e0e0 !important;
                font-size: 16px !important;
                font-weight: 500;
                text-decoration: none !important;
                display: inline-block;
                padding-right: 20px;
                border-right: 1px solid #828ea1;
                line-height: 50px;
                height: 50px;
                vertical-align: top;

                &:hover {
                  color: #fff !important;
                  text-decoration: none !important;
                }
              }
            }

            .login_register {
              float: left;
              padding-right: 20px;
              border-right: 1px solid #273c55;
              line-height: 50px;

              & :deep(.el-menu) {
                background: transparent;

                #login,
                #register {
                  display: inline-block;
                  min-width: 60px;
                  height: 100%;
                  text-align: center;
                  line-height: 20px;
                  margin-left: 0;
                  box-sizing: border-box;

                  .el-menu-item {
                    height: 100%;
                    color: #828ea1;
                  }

                  &:hover {
                    .el-menu-item {
                      color: #fff;
                    }
                  }
                }

                #login {
                  border-right: 1px solid #273c55;
                }

                #register {
                  .el-menu-item {
                    color: #f0a70a !important;
                  }

                  &:hover {
                    .el-menu-item {
                      color: #fff;
                    }
                  }
                }
              }
            }

            .isLogin {
              :deep(.el-dropdown) {
                display: block;
                float: left;

                .el-dropdown-link {
                  color: #828ea1;

                  &:hover {
                    color: #fff;
                  }
                }
              }

              // 用户下拉触发器样式 - 修复对齐问题
              .user-dropdown-trigger {
                display: flex;
                align-items: center;
                cursor: pointer;
                color: #e0e0e0;
                font-size: 16px;
                font-weight: 500;
                height: 50px;
                line-height: 50px;

                .user-icon,
                .arrow-icon {
                  vertical-align: middle;
                  display: inline-flex;
                  align-items: center;
                }

                .username-text {
                  margin: 0 6px;
                  white-space: nowrap;
                  text-decoration: none;
                }
              }

              // 用户下拉菜单样式
              .user-dropdown-menu {
                background-color: #27313e;
                border: 1px solid #394559;
                border-radius: 4px;
                padding: 8px 0;
                min-width: 180px;
                box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);

                .dropdown-item {
                  padding: 0;

                  .dropdown-item-link {
                    display: flex;
                    align-items: center;
                    padding: 10px 16px;
                    color: #ccc;
                    text-decoration: none;
                    transition: all 0.2s ease;

                    &:hover {
                      background-color: #2f3e51;
                      color: #f0ac19;

                      .el-icon {
                        color: #f0ac19;
                      }
                    }

                    .el-icon {
                      margin-right: 8px;
                      font-size: 16px;
                      color: #828ea1;
                    }

                    span {
                      white-space: nowrap;
                    }
                  }
                }

                .dropdown-divider {
                  border-top: 1px solid #394559;
                  margin-top: 4px;
                  padding-top: 4px;

                  .dropdown-item-link {
                    color: #ccc;

                    &:hover {
                      color: #f0ac19;
                    }

                    .el-icon {
                      color: #828ea1;

                      &:hover {
                        color: #f0ac19;
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}

.page-view2 {
  :deep(.el-select__selection) {
    background-color: #0c1621;

    &:hover {
      border-color: transparent;
    }

    &:focus {
      border-color: transparent;
    }
  }

  .page-content {
    .layout {
      background: #172636;
      width: 100%;
      z-index: 10;
      position: absolute;
      top: 0;

      .layout-ceiling {
        .layout-ceiling-main {
          .header_nav {
            display: none;
          }
        }

        .rr {
          display: none;
        }
      }
    }
  }

  .footer {
    .footer_content {
      .footer_right {
        display: none;
      }
    }
  }
}

.page-view3 {
  background: linear-gradient(150deg, #c3333d, #bc000d, #ff1d2c);
  min-height: 100%;
  background-color: #fff;

  .page-content {
    padding-bottom: 20px !important;

    .layout {
      display: none;
    }

    .time_download {
      display: none;
    }
  }

  .footer {
    display: none;
  }
}

.appdownload {
  :deep(.el-popover) {
    background-color: #27313e;
    color: #fff;
    padding-top: 10px;
  }
}

.wechatclick {
  :deep(.el-popover) {
    background-color: #27313e;
    color: #fff;
  }
}
</style>

<style lang="scss">
// 全局样式
.container_test {
  padding-top: 60px;
}

.ivu-table-filter-list .ivu-table-filter-select-item {
  color: #ccc;
  &:hover {
    background-color: #27313e;
    color: #f0ac19;
  }
}

.ivu-table-filter-list .ivu-table-filter-select-item-selected {
  color: #f0ac19;
  &:hover {
    color: #f0ac19;
  }
}

.ivu-table-filter i.on {
  color: #fff;
}

.ivu-message {
  color: #333;
}

.ivu-poptip-inner {
  background-color: #27313e;
  color: #fff;

  .ivu-poptip-body-content-inner {
    color: #fff;
  }
}

.ivu-poptip-popper {
  .ivu-poptip-arrow:after {
    left: 0 !important;
    border-right-color: #27313e !important;
  }
}

.exchange .ivu-checkbox-checked .ivu-checkbox-inner {
  background-color: #f0a70a;
  border-color: #f0a70a;
}

.ivu-modal-confirm-head {
  text-align: center;
  margin-bottom: 15px;
}

.ivu-modal-header p,
.ivu-modal-header-inner {
  color: #fff;
}

.ivu-modal-body {
  border-radius: 5px;

  .ivu-modal-confirm {
    .ivu-modal-confirm-body {
      font-size: 14px;
    }
  }
}

.ivu-modal-confirm-footer .ivu-btn-primary {
  background-color: #f0a70a;
  border-color: #f0a70a;
}

.ivu-modal-confirm-footer .ivu-btn-text {
  &:hover {
    color: #f0a70a;
  }
}

.ivu-table-wrapper {
  background-color: #192330;

  .ivu-table {
    box-shadow: 0px 0px 4px #27313e;
    background-color: #192330;
    color: #ccc;

    &:before {
      background: transparent;
    }

    &:after {
      background: #192330;
    }

    .ivu-table-header {
      th {
        background-color: #27313e;
        border: none;
        color: #ccc;
      }
    }

    .ivu-table-row:hover {
      background: #1e2834;
    }

    .ivu-table-row td {
      background-color: transparent;
      border: none;
      border-bottom: 1px solid #27313e;
      color: #fff;
    }
  }
}

.ivu-table td {
  background-color: #192330;
  border-bottom: 1px solid #27313e;
}

.ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
  background: none;

  &:after {
    background: none;
  }
}

.ivu-select-dropdown .ivu-select-item {
  color: #ccc;
  padding: 6px 16px;
}

.page-view {
  height: 100%;

  .page-content {
    .layout {
      .layout-ceiling {
        .layout-ceiling {
          background: #172636;
          box-shadow: 0 0 5px 5px rgba(0, 0, 0, 0.1);

          .layout-ceiling-main {
            .header_nav {
              .ivu-menu-vertical.ivu-menu-light {
                .ivu-menu-submenu-title {
                  i.ivu-icon.ivu-icon-ios-arrow-down.ivu-menu-submenu-title-icon {
                    &:before {
                      content: '';
                    }
                  }
                }
              }
            }

            .rr {
              .login_register .ivu-menu-submenu-title .ivu-icon {
                &:before {
                  content: '';
                }
              }
            }
          }
        }
      }
    }
  }
}

.page-view2 {
  height: 100%;

  .page-content {
    .layout {
      .layout-ceiling {
        .layout-ceiling-main {
          .header_nav {
            .ivu-menu-vertical.ivu-menu-light {
              .ivu-menu-submenu-title {
                i.ivu-icon.ivu-icon-ios-arrow-down.ivu-menu-submenu-title-icon {
                  &:before {
                    content: '';
                  }
                }
              }
            }
          }

          .rr {
            .login_register .ivu-menu-submenu-title .ivu-icon {
              &:before {
                content: '';
              }
            }
          }
        }
      }
    }
  }
}

html,
body {
  height: 100%;
  font-size: 14px;
  background: #0b1520;
  color: #fff;
}

::-webkit-scrollbar {
  width: 3px;
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #39557a;
  border-radius: 25px;
}

.ivu-carousel-dots li button {
  width: 30px;
  height: 10px;
  border-radius: 14px;
}

.ivu-menu-dark,
.ivu-menu-dark.ivu-menu-vertical .ivu-menu-opened {
  background: #18202a;
}

#checkbox {
  width: 10px;
}

.login_title {
  color: #000;
  text-align: center;
  height: 80px;
  font-size: 25px;
}

.login_right .ivu-select-dropdown {
  background: #fff;
}

.ivu-form-inline .ivu-form-item {
  display: block;
  margin-right: 0;
}

.layout {
  position: absolute;
}

.layout-copy {
  text-align: center;
  padding: 10px 0 20px;
  color: #9ea7b4;
}

.layout-ceiling-main {
  height: 50px;
  line-height: 50px;
  margin-left: 128px;
}

.layout-ceiling-main .rr {
  float: right;
}

.layout-ceiling-main .ivu-menu-vertical .ivu-menu-item,
.ivu-menu-vertical .ivu-menu-submenu-title {
  padding: 0;
}

.layout-ceiling-main .ivu-menu-item {
  font-size: 14px;
}

.layout-ceiling-main a {
  color: #fff;
  display: inline-block;
  line-height: 40px;
  height: 40px;
  text-align: center;
  margin-left: 20px;
}

.shoujiShow {
  display: none;
}

@media screen and (max-width: 768px) {
  .header_nav {
    display: none;
  }

  .login-container {
    display: none;
  }

  .footer_right {
    display: none;
  }

  .rightwrapper {
    display: none;
  }

  .shoujiShow {
    display: block !important;
  }

  .sjShow_content {
    display: flex;
    position: relative;
    top: -150px;
  }

  .sjShow_content div {
    width: 25%;
    text-align: center;
  }

  .sjShow_content div a {
    color: #ccc;
  }
}

.header_nav {
  float: left;
}

.ivu-dropdown-rel a {
  width: 100%;
}

.ivu-dropdown-menu {
  width: 120px;
}

.layout-ceiling-main .ivu-select-dropdown {
  background: #27313e;
  margin-left: 25px;

  .ivu-dropdown-item {
    padding: 10px 16px;
    color: #ccc;

    &:hover {
      color: #f0ac19;
    }
  }
}

.ivu-select-dropdown a {
  width: 100%;
  text-align: left;
  margin: 0;
  height: 20px;
  line-height: 20px;
}

.ivu-dropdown-item:hover {
  background-color: #27313e;
  color: #f0ac19;
}

.ivu-dropdown-item img {
  width: 14px;
  vertical-align: middle;
}

.ivu-radio-inner:after {
  background: #18202a;
}

.user_center {
  height: 900px;
}

.ivu-menu-item {
  text-align: center;
}

.ivu-menu-vertical .ivu-menu-submenu .ivu-menu-item {
  padding-left: 0 !important;
  padding-right: 0;
  color: rgba(130, 142, 161, 1);
  font-size: 14px;
  border-right: 0 !important;
}

.ivu-menu-dark.ivu-menu-vertical .ivu-menu-submenu .ivu-menu-item-active,
.ivu-menu-dark.ivu-menu-vertical .ivu-menu-submenu .ivu-menu-item-active:hover {
  background: #1855fd !important;
}

.rr .ivu-menu-vertical.ivu-menu-light:after {
  width: 0;
}

.layout_menu_right {
  margin-left: 3%;
  background: #18202a;
  color: #fff;
  padding-bottom: 130px;
}

.menu_right_title {
  font-size: 16px;
  line-height: 45px;
  margin: 0 10px;
  padding-left: 20px;
  border-bottom: 1px solid #263142;
}

.category .ivu-radio-group.ivu-radio-group-button {
  width: 100%;
}

.category .ivu-radio-group label {
  font-size: 14px;
}

.category .ivu-radio-group-button .ivu-radio-wrapper {
  background: #28313e;
  color: #979797;
  border: 0;
  padding: 0 25px;
}

.category .ivu-radio-group-button .ivu-radio-wrapper-checked {
  color: #fff;
  background: #2f3e52;
  box-shadow: none;
}

.category .ivu-radio-wrapper span {
  padding-left: 0;
}

.purse_address_left {
  float: left;
  width: 86%;
}

.purse_address p {
  font-size: 10px;
  line-height: 25px;
  color: #979797;
}

.purse_address_left_icon {
  line-height: 40px;
  height: 40px;
  width: 100%;
}

.purse_address_left_icon img {
  vertical-align: middle;
  margin-right: 10px;
}

.purse_address span {
  font-size: 14px;
  float: left;
  color: #fff;
  padding: 0 20px;
  background: #28313e;
  width: 21%;
}

.purse_address_left_icon label {
  float: left;
  width: 72%;
  height: 40px;
  border: 2px solid #28313e;
  padding-left: 20px;
}

#qrcode canvas {
  width: 120px;
}

#qrcode img {
  width: 100%;
}

.chart_container #chart_updated_time {
  float: left;
}

.page-content {
  min-height: 100%;
  padding-bottom: 200px;
}

.footer {
  width: 100%;
  overflow: hidden;
  margin-top: -200px;
}

.footer_content {
  height: 300px;
  padding: 80px 10%;
  color: #53575c;
  color: rgba(255, 255, 255, 0.8);
  background: #192330;
}

.footer_left {
  float: left;
  font-size: 14px;
}

.footer_left img {
  margin: 15px 0;
  width: 300px;
}

.footer_left p {
  margin: 10px 0;
  color: #828ea1;
}

.footer_right {
  float: right;
  text-align: left;
}

.footer_right ul {
  float: left;
  margin: 0 30px;
}

.footer_right ul li {
  list-style-type: none;
}

.footer_right ul li a {
  color: #828ea1;
  line-height: 30px;
  display: block;
  font-size: 12px;
}

.footer_right ul li a:hover {
  color: #fff;
}

.footer_title {
  font-size: 13px;
  height: 40px;
}

.ivu-select-selected-value {
  color: #bbbec4;
}

.ivu-col {
  text-align: center;
}

.page-view {
  .page-content {
    .layout {
      .layout-ceiling {
        .rr {
          .login_register {
            .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
              color: #fff;
            }
          }

          .isLogin {
            .ivu-dropdown {
              display: inline-block;

              .ivu-select-dropdown {
                padding: 0;
                margin: 0;

                .ivu-dropdown-menu {
                  .ivu-dropdown-item {
                    border-radius: 5px;
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}

.changelanguage {
  .ivu-dropdown {
    .ivu-select-dropdown {
      z-index: 901;
    }
  }
}

.ivu-page-next,
.ivu-page-prev {
  background-color: #192330;
}

.ivu-page-item {
  background-color: #192330;
  border-color: #27313e;
}

.ivu-page-item-jump-next,
.ivu-page-item-jump-prev,
.ivu-page-next,
.ivu-page-prev {
  border-color: #27313e;
}

.ivu-page-item-active {
  font-weight: bold;
}

.ivu-page-next:hover,
.ivu-page-prev:hover {
  border-color: #f0ac19;
}

.ivu-page-next:hover a,
.ivu-page-prev:hover a {
  color: #f0ac19;
}

.ivu-page-item-jump-next a,
.ivu-page-item-jump-next a {
  color: #666;
}

.ivu-page-item-jump-prev a:hover,
.ivu-page-item-jump-next a:hover {
  color: #f0ac19;
}

.ivu-page-item:hover {
  border-color: #f0ac19;
}

.ivu-page-item:hover a {
  color: #f0ac19;
}

.ivu-page-item.ivu-page-item-active a {
  color: #f0ac19;
}

.ivu-page-disabled {
  a {
    cursor: not-allowed;

    .ivu-icon {
      cursor: not-allowed;
    }
  }
}

.ivu-input,
.ivu-input-number-input,
.ivu-input-number {
  background-color: #192330;
  color: #fff;
  border-color: #27313e;

  &:hover {
    border-color: #27313e;
  }

  &:focus {
    border-color: #27313e;
    box-shadow: none;
  }
}

.ivu-input[disabled]:hover,
fieldset[disabled] .ivu-input:hover {
  border-color: #27313e;
}

.ivu-input[disabled],
fieldset[disabled] .ivu-input {
  background-color: #27313e;
}

.ivu-input-number-focused {
  box-shadow: none;
}

.ivu-input-number:focus {
  box-shadow: none;
}

.ivu-form .ivu-form-item-label {
  color: #ccc;
}

.ivu-input-number-handler-wrap {
  background: #27313e;
  border-left: 1px solid #192330;
}

.ivu-input-number-handler {
  border-top: 1px solid #192330;
}

.ivu-input-number-handler:hover .ivu-input-number-handler-up-inner,
.ivu-input-number-handler:hover .ivu-input-number-handler-down-inner {
  color: #ccc;
}

.ivu-input-group-append,
.ivu-input-group-prepend {
  color: #ccc;
}

.ivu-select-selection {
  background-color: #192330;
  color: #fff;
  border: 1px solid #27313e;
}

.ivu-select-selection:hover {
  border-color: #27313e;
}

.ivu-select-visible .ivu-select-selection {
  border-color: #27313e;
  box-shadow: none;
}

.ivu-select-selected-value {
  color: #fff;
}

.ivu-select-selection-focused {
  border-color: #27313e;
}

.ivu-select-dropdown {
  background-color: #192330;
}

.ivu-select-disabled .ivu-select-selection {
  background-color: #27313e;
}

.ivu-select-disabled .ivu-select-selection:hover {
  border-color: #27313e;
}

.ivu-select-item-selected {
  background-color: #192330;
  color: #ccc;
}

.ivu-select-item-focus {
  background-color: #192330;
}

.ivu-select-item:hover {
  background-color: #27313e;
  color: #f0ac19;
}

.ivu-select-multiple .ivu-select-item-selected {
  background-color: #192330;
  color: #f0ac19;
}

.ivu-select-multiple .ivu-select-item-focus,
.ivu-select-multiple .ivu-select-item-selected:hover {
  background-color: #192330;
}

.ivu-select-multiple .ivu-select-item-selected:after {
  color: #f0ac19;
}

.ivu-select-item-selected,
.ivu-select-item-selected:hover {
  background-color: #192330;
  color: #f0ac19;
}

.ivu-checkbox-inner {
  background-color: #192330;
}

.ivu-switch {
  border: 1px solid #27313e;
  background-color: #192330;
}

.ivu-switch:after {
  background-color: #ccc;
}

.ivu-tag {
  border: 1px solid #27313e;
  border-radius: 3px;
  background: #192330;
}

.ivu-tag-text {
  color: #ccc;
}

.ivu-table-wrapper {
  border: none;
}

.ivu-table-wrapper > .ivu-spin-fix {
  background-color: rgba(0, 0, 0, 0.2);
  border: none;
  border-color: #fff;
}

.ivu-spin-fix {
  background-color: rgba(0, 0, 0, 0.2);
  border: none;
  border-color: #fff;
}

.ivu-spin-dot {
  background: #f0ac19;
}

.ivu-tabs-bar {
  border-color: #f5f5f5;
}

.ivu-picker-panel-icon-btn {
  &:hover {
    color: #f0ac19;
  }
}

.ivu-date-picker-focused input {
  border-color: #1f2936;
  box-shadow: none;
}

.ivu-date-picker-cells-focused em {
  box-shadow: none;
  color: #f0ac19;

  &:after {
  }
}

.ivu-date-picker-cells-cell {
  color: #fff;
}

.ivu-date-picker-cells-cell-selected em,
.ivu-date-picker-cells-cell-selected:hover em {
  background: #27313e;
  color: #f0ac19;
}

.ivu-date-picker-cells-cell-today em:after {
  background: #27313e;
}

.ivu-date-picker-cells-cell-range:before {
  background: rgba(240, 167, 10, 0.2);
}

.ivu-date-picker-cells-cell:hover em {
  background: #27313e;
  color: #f0ac19;
}

.ivu-btn {
  border: none;
}

.ivu-btn-primary:hover {
  background: #f0ac19;
  border-color: #f0ac19;
}

.ivu-btn.ivu-btn-default {
  background-color: #27313e;
  color: #fff;

  &:hover {
    color: #f0a70a;
  }

  &:active {
    color: #f0a70a;
  }
}

.ivu-btn-text {
  color: #ccc;
  border: 1px solid #27313e;
}

.ivu-btn-primary {
  background-color: #f0ac19;
  border-color: #f0ac19;
}

.ivu-btn-text:hover {
  background-color: transparent;
  color: #f0ac19;
}

.ivu-input-group-append,
.ivu-input-group-prepend {
  background-color: #27313e;
  border: 1px solid #27313e;
}

.ivu-form-item-error .ivu-input-group-append,
.ivu-form-item-error .ivu-input-group-prepend {
  background-color: #27313e;
  border: 1px solid #27313e;
}

.ivu-form-item-error .ivu-input,
.ivu-form-item-error .ivu-input:focus,
.ivu-form-item-error .ivu-input:hover {
  border: 1px solid #27313e;
  box-shadow: none;
}

.ivu-radio-checked .ivu-radio-inner {
  border-color: #f0ac19;
}

.ivu-radio-checked:hover {
  .ivu-radio-inner {
    border-color: #f0ac19;
  }
}

.ivu-radio-inner:after {
  background: #f0ac19;
}

.ivu-switch-checked {
  border-color: #f0ac19;
  background-color: #f0ac19;
}

.ivu-switch:focus {
  box-shadow: none;
}

.ivu-radio-focus {
  box-shadow: none;
}

.ivu-modal-content {
  background-color: #192330;
}

.ivu-modal-header {
  border-bottom: 1px solid #27313e;
}

.ivu-modal-confirm-head-icon-confirm {
  color: #fff;
}

.ivu-modal-header p {
  color: #fff;
}

.ivu-modal-footer {
  border-top: 1px solid #27313e;
}

.ivu-table-sort i.on {
  color: #f0ac19;
}

.ivu-table-sort i:hover {
  color: #f0ac19;
}

.ivu-modal-confirm-head {
  font-size: 24px;
}

.ivu-modal-confirm-body {
  color: #fff;
  padding-left: 0;
}

.ivu-modal-confirm-head-title {
  color: #fff;
  margin-left: 5px;
}

.ivu-modal-confirm-footer {
  padding-top: 10px;
  border-top: 1px solid #27313e;
}

.ivu-upload-list-file:hover {
  background-color: #27313e;
}

.ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-active,
.ivu-menu-light.ivu-menu-horizontal .ivu-menu-item:hover,
.ivu-menu-light.ivu-menu-horizontal .ivu-menu-submenu-active,
.ivu-menu-light.ivu-menu-horizontal .ivu-menu-submenu:hover {
  border-bottom: 0 !important;
  color: #828ea1 !important;
}

.ivu-menu-horizontal .ivu-menu-submenu .ivu-select-dropdown .ivu-menu-item:hover {
  background: #2f3e51 !important;
}

.ivu-menu-horizontal.ivu-menu-light {
  background: transparent !important;
}

.ivu-menu-horizontal.ivu-menu-light:after {
  height: 0 !important;
}

.ivu-select-dropdown {
  border-radius: 0 !important;
}

.lang-img {
  height: 20px;
  margin-bottom: -5px;
  margin-right: 5px;
}

.lang-item {
  text-align: left;

  img {
    height: 20px;
    margin-bottom: -5px;
    margin-right: 5px;
  }

  &:hover {
    background: #2f3e51;
  }
}

.ivu-message-notice-content {
  background: #324368;
  color: #a3bbcc;
}

.social-list {
  ul {
    list-style: none;
    padding-top: 5px;

    li {
      transition: all 0.5s;
      width: 25px;
      height: 25px;
      line-height: 25px;
      border-radius: 2px;
      background: rgb(57, 69, 89);
      text-align: center;
      float: left;
      margin-right: 8px;
      color: #a3b6c6;

      &:hover {
        color: #fff;
        cursor: pointer;
      }
    }
  }
}

.ivu-tooltip-inner {
  background: #394559;
}

.ivu-tooltip-arrow {
  border-bottom-color: #394559;
}

.ivu-notice-notice {
  background: #21364d;
}

.ivu-notice-title {
  color: #fff;
}

.ivu-notice-desc {
  color: #fff;
}

.swiper-pagination-fraction,
.swiper-pagination-custom,
.swiper-container-horizontal > .swiper-pagination-bullets {
  bottom: -5px;
}

.swiper-pagination-bullet {
  background: #fff;
  border-radius: 2px;
  height: 3px;
  width: 15px;
  opacity: 0.3;
  transition: all 0.5s;
}

.swiper-pagination-bullet-active {
  background: #f0a70a !important;
  width: 30px;
  opacity: 1;
}

.login_right .ivu-select-dropdown {
  background: #212b36;
}

.login_right .ivu-select-dropdown .ivu-select-item {
  text-align: left;
}

.ivu-form-item-error .ivu-input-group-append,
.ivu-form-item-error .ivu-input-group-prepend,
.ivu-input-group-prepend {
  background-color: #17212e;
  border-bottom: 1px solid #27313e;
  border-top: none;
  border-left: none;
  border-right: none;
}

.ivu-input-group-append {
  background-color: #17212e;
  border-bottom: 1px solid #27313e;
  border-left: none;
}

.ivu-select-single .ivu-select-selection {
  background-color: #17212e;
}

.login_form {
  input::-webkit-input-placeholder {
    color: #8a8a8acf !important;
    font-size: 0.95rem !important;
    letter-spacing: 1px !important;
  }

  input:-moz-placeholder {
    color: #8a8a8a !important;
    font-size: 13px !important;
    letter-spacing: 1px !important;
  }

  input::-moz-placeholder {
    color: #8a8a8a !important;
    font-size: 13px !important;
    letter-spacing: 1px !important;
  }

  input::-ms-input-placeholder {
    color: #8a8a8a !important;
    font-size: 13px !important;
    letter-spacing: 1px !important;
  }

  .ivu-input-group-prepend {
    font-size: 0.95rem;
    letter-spacing: 1px;
  }
}

.login_form .login_right form.ivu-form.ivu-form-label-right.ivu-form-inline .password .ivu-form-item-content .ivu-input-wrapper.ivu-input-type .ivu-input {
  letter-spacing: 8px;
}

.ivu-menu-light {
  background: transparent !important;
}

.ivu-spin-fullscreen-wrapper {
  background: #46597a70 !important;
}

.ivu-spin {
  color: #f0a70a !important;
}

.ivu-poptip-popper[x-placement^='bottom'] .ivu-poptip-arrow {
  border-bottom-color: #27313e;
}

.ivu-poptip-popper[x-placement^='bottom'] .ivu-poptip-arrow:after {
  border-bottom-color: #27313e;
}

.ivu-poptip-title-inner {
  color: #ccc;
  font-size: 14px;
}

.ivu-poptip-title:after {
  background-color: #394253;
}

.tag-hot {
  display: inline-block;
  padding: 0 4px;
  background: #ff0000;
  color: #fff;
  line-height: 16px;
  font-size: 10px;
  margin-left: 5px;
  margin-top: -5px;
  border-radius: 2px;
  position: absolute;
  top: 16px;
  font-weight: 600;
}

.page {
  text-align: right;
  margin-top: 10px;

  .ivu-page {
    .ivu-page-prev,
    .ivu-page-next {
      background: transparent !important;
      color: #000;
      border: none;
    }

    .ivu-page-item {
      background-color: transparent !important;
      color: #000;
      border: none;
    }
  }
}

.ivu-progress-bg {
  border-radius: 0 !important;
  background-color: #ff8100;
  max-width: 100%;
}

.ivu-progress-success .ivu-progress-bg {
  background-color: #ff8100 !important;
}

.header_nav_mobile .ivu-menu-vertical .ivu-menu-item,
.header_nav_mobile .ivu-menu-vertical .ivu-menu-submenu-title {
  padding: 8px 24px 8px 5px;
  color: #828ea1;
}

.header_nav_mobile .ivu-drawer-wrap .ivu-drawer-no-header .ivu-drawer-content .ivu-drawer-body {
  background: #2b323a;
  padding-top: 60px;
}

.header_nav_mobile .ivu-menu-vertical.ivu-menu-light:after {
  background: transparent !important;
}

.header_nav_mobile .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
  color: #f0a70a;
}
</style>

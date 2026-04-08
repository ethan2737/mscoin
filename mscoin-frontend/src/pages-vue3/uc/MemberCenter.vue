<style lang="scss" scoped>
.mymsg_wrapper {
  padding: 70px 8% 0;
  margin-bottom: 10px;
  .wrapper {
    width: 100%;
    background: #192330;
    overflow: hidden;
    padding: 30px 15px 30px 0;

    .pc_menu {
      padding: 0 10px;
      text-align: left;
      float: left;
      background: #192330;
      min-height: 662px;
      position: static;
      border-right: 1px solid #27313e;
      color: #fff;

      .title {
        text-align: center;
        font-size: 20px;
        line-height: 40px;
        color: #fff;
        margin-bottom: 20px;
      }
    }

    .nav-rights {
      width: 100%;
      float: left;
      min-height: 600px;
    }
  }
}
</style>
<style lang="scss">
.mymsg_wrapper {
  .wrapper {
    .pc_menu {
      // 菜单样式
      .el-menu {
        background: #192330;
        border-right: none;

        .el-menu-item {
          background: #27313e;
          color: #fff;

          &:hover {
            background-color: #3c4553;
          }

          &.is-active {
            color: #f0ac19;

            a {
              color: #f0ac19;
            }
          }

          a {
            color: #fff;
            display: block;
            height: 100%;
            padding: 14px 24px;
            font-size: 14px;
          }
        }

        .el-sub-menu__title {
          height: 50px;
          padding: 15px 10px;

          &:hover {
            background-color: #3c4553;

            .content, .el-icon {
              color: #fff;
            }
          }

          .isclick {
            width: 6px;
            height: 6px;
            background-color: #f0ac19;
            display: inline-block;
            border-radius: 50%;
            margin-top: 16px;
            line-height: 50px;
            display: none;
          }

          .content {
            padding-left: 5px;
            padding-top: 5px;
            color: #fff;
          }
        }
      }

      .el-sub-menu.is-opened {
        .el-sub-menu__title {
          background-color: #3c4553;

          .isclick {
            background-color: #fff;
          }

          .content,
          .el-icon {
            color: #fff;
          }
        }
      }
    }
  }
}

.header_nav_mobile_triggle {
  display: none;
  margin-bottom: 20px;
}

@media screen and (max-width: 768px) {
  .pc_menu {
    display: none !important;
  }

  .header_nav_mobile_triggle {
    display: block !important;
    text-align: left;
    padding-left: 20px;
  }

  .mymsg_wrapper {
    padding: 40px 0 !important;

    .wrapper {
      padding: 0px 0px 30px 0px !important;

      .nav-rights {
        width: 100% !important;
      }
    }

    .account-box .account-in .account-item .account-item-in {
      padding: 15px 5px 15px 5px !important;
    }

    .account-box .account-in .account-item .account-item-in .bankInfo {
      width: 50% !important;
    }

    .account-box .account-in .account-item .account-item-in .card-number {
      width: 100px !important;
    }

    .nav-right {
      padding-left: 0px !important;
    }
  }
}

.uc_header_nav {
  .el-menu--vertical {
    text-align: left;

    .el-menu-item {
      text-align: left;
      color: #DDD;

      a {
        color: #828ea1;
      }

      &.is-active {
        a {
          color: #f0a70a !important;
        }
      }
    }
  }
}
</style>

<template>
  <div class="mymsg_wrapper">
    <div class="wrapper">
      <el-row :gutter="20" style="min-height: 600px; padding-top: 10px;">
        <el-col :xs="24" :sm="24" :md="4" :lg="4">
          <div class="header_nav_mobile_triggle" @click="toggleMemu">
            <el-icon style="font-size: 26px; color: #cccccc; vertical-align: middle;">
              <Menu />
            </el-icon>
            个人中心
          </div>

          <el-menu
            class="pc_menu"
            :default-active="activename"
            :default-openeds="opennames"
            :accordion="true"
            @open="handleOpen"
            @close="handleClose"
          >
            <div class="title">{{ userCenterTitle }}</div>

            <el-sub-menu index="1">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ accountTitle }}</span>
              </template>
              <el-menu-item index="1-1">
                <router-link to="/uc/safe">{{ securityTitle }}</router-link>
              </el-menu-item>
              <el-menu-item index="1-2">
                <router-link to="/uc/account">{{ accountSetting }}</router-link>
              </el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="2">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ assetsTitle }}</span>
              </template>
              <el-menu-item index="2-1">
                <router-link to="/uc/money">{{ personalAssets }}</router-link>
              </el-menu-item>
              <el-menu-item index="2-2">
                <router-link to="/uc/contract-money">{{ contractAssets }}</router-link>
              </el-menu-item>
              <el-menu-item index="2-3">
                <router-link to="/uc/record">{{ billDetail }}</router-link>
              </el-menu-item>
              <el-menu-item index="2-6">
                <router-link to="/uc/contract-record">{{ contractBillDetail }}</router-link>
              </el-menu-item>
              <el-menu-item index="2-4">
                <router-link to="/uc/recharge">{{ chargeTitle }}</router-link>
              </el-menu-item>
              <el-menu-item index="2-5">
                <router-link to="/uc/withdraw">{{ pickupTitle }}</router-link>
              </el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="3">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ exchangeTitle }}</span>
              </template>
              <el-menu-item index="3-1">
                <router-link to="/uc/entrust/current">{{ curDelegate }}</router-link>
              </el-menu-item>
              <el-menu-item index="3-2">
                <router-link to="/uc/entrust/history">{{ hisDelegate }}</router-link>
              </el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="7">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ swapTitle }}</span>
              </template>
              <el-menu-item index="7-1">
                <router-link to="/uc/contract/entrust/current">{{ curDelegate }}</router-link>
              </el-menu-item>
              <el-menu-item index="7-2">
                <router-link to="/uc/contract/entrust/history">{{ hisDelegate }}</router-link>
              </el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="4" v-if="false">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ otcTitle }}</span>
              </template>
              <el-menu-item index="4-1">
                <router-link to="/uc/ad">{{ myadTitle }}</router-link>
              </el-menu-item>
              <el-menu-item index="4-2">
                <router-link to="/uc/order">{{ myorderTitle }}</router-link>
              </el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="5">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ activityTitle }}</span>
              </template>
              <el-menu-item index="5-1">
                <router-link to="/uc/innovation/myorders">{{ subActivity1 }}</router-link>
              </el-menu-item>
              <el-menu-item index="5-2">
                <router-link to="/uc/innovation/myminings">{{ subActivity2 }}</router-link>
              </el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="8">
              <template #title>
                <span class="isclick"></span>
                <span class="content">{{ crowdfundingTitle }}</span>
              </template>
              <el-menu-item index="8-1">
                <router-link to="/uc/crowdfunding/list">{{ crowdfundingList }}</router-link>
              </el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-col>

        <el-col :xs="24" :sm="24" :md="20" :lg="20">
          <router-view></router-view>
        </el-col>
      </el-row>
    </div>

    <!-- 移动端抽屉 -->
    <el-drawer
      v-model="drawerVisible"
      :closable="true"
      direction="ltr"
      size="40%"
      class="header_nav_mobile uc_header_nav"
    >
      <el-menu
        :default-active="activename"
        :default-openeds="opennames"
        :accordion="true"
        @open="handleOpen"
        @close="handleClose"
      >
        <el-sub-menu index="1">
          <template #title>
            <span class="isclick"></span>
            <span class="content">{{ accountTitle }}</span>
          </template>
          <el-menu-item index="1-1">
            <router-link to="/uc/safe">{{ securityTitle }}</router-link>
          </el-menu-item>
          <el-menu-item index="1-2">
            <router-link to="/uc/account">{{ accountSetting }}</router-link>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="2">
          <template #title>
            <span class="isclick"></span>
            <span class="content">{{ assetsTitle }}</span>
          </template>
          <el-menu-item index="2-1">
            <router-link to="/uc/money">{{ personalAssets }}</router-link>
          </el-menu-item>
          <el-menu-item index="2-2">
            <router-link to="/uc/contract-money">{{ contractAssets }}</router-link>
          </el-menu-item>
          <el-menu-item index="2-3">
            <router-link to="/uc/record">{{ billDetail }}</router-link>
          </el-menu-item>
          <el-menu-item index="2-6">
            <router-link to="/uc/contract-record">{{ contractBillDetail }}</router-link>
          </el-menu-item>
          <el-menu-item index="2-4">
            <router-link to="/uc/recharge">{{ chargeTitle }}</router-link>
          </el-menu-item>
          <el-menu-item index="2-5">
            <router-link to="/uc/withdraw">{{ pickupTitle }}</router-link>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="3">
          <template #title>
            <span class="isclick"></span>
            <span class="content">{{ exchangeTitle }}</span>
          </template>
          <el-menu-item index="3-1">
            <router-link to="/uc/entrust/current">{{ curDelegate }}</router-link>
          </el-menu-item>
          <el-menu-item index="3-2">
            <router-link to="/uc/entrust/history">{{ hisDelegate }}</router-link>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="7">
          <template #title>
            <span class="isclick"></span>
            <span class="content">{{ swapTitle }}</span>
          </template>
          <el-menu-item index="7-1">
            <router-link to="/uc/contract/entrust/current">{{ curDelegate }}</router-link>
          </el-menu-item>
          <el-menu-item index="7-2">
            <router-link to="/uc/contract/entrust/history">{{ hisDelegate }}</router-link>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="5">
          <template #title>
            <span class="isclick"></span>
            <span class="content">{{ activityTitle }}</span>
          </template>
          <el-menu-item index="5-1">
            <router-link to="/uc/innovation/myorders">{{ subActivity1 }}</router-link>
          </el-menu-item>
          <el-menu-item index="5-2">
            <router-link to="/uc/innovation/myminings">{{ subActivity2 }}</router-link>
          </el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="8">
          <template #title>
            <span class="isclick"></span>
            <span class="content">{{ crowdfundingTitle }}</span>
          </template>
          <el-menu-item index="8-1">
            <router-link to="/uc/crowdfunding/list">{{ crowdfundingList }}</router-link>
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-drawer>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 用户中心布局组件
 * 迁移点：
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. 使用 Element Plus 组件替代 iView 组件
 * 4. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 *
 * 组件映射：
 * - <Menu> → <el-menu>
 * - <Submenu> → <el-sub-menu>
 * - <MenuItem> → <el-menu-item>
 * - <Row> → <el-row>
 * - <Col> → <el-col>
 * - <Drawer> → <el-drawer>
 * - <Icon> → <el-icon>
 */
import { ref, reactive, inject, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Menu } from '@element-plus/icons-vue'

// Vuex 3.x 和 Vue Router 3.x 兼容方案
const store = inject('store')
const router = inject('router')

// 文本常量
const userCenterTitle = '用户中心'
const accountTitle = '账户'
const securityTitle = '安全设置'
const accountSetting = '账户设置'
const assetsTitle = '资产'
const personalAssets = '个人资产'
const contractAssets = '合约资产'
const billDetail = '账单明细'
const contractBillDetail = '合约账单明细'
const chargeTitle = '充值'
const pickupTitle = '提现'
const exchangeTitle = '交易所'
const curDelegate = '当前委托'
const hisDelegate = '历史委托'
const swapTitle = '永续合约'
const otcTitle = 'OTC'
const myadTitle = '我的广告'
const myorderTitle = '我的订单'
const activityTitle = '活动中心'
const subActivity1 = '我的订单'
const subActivity2 = '我的挖矿'
const crowdfundingTitle = '众筹'
const crowdfundingList = '众筹列表'

// 响应式数据
const activename = ref('1-1')
const opennames = ref(['1'])
const drawerVisible = ref(false)

// 路由映射
const routeArr = reactive({
  '/uc/safe': '1-1',
  '/uc/account': '1-2',
  '/uc/money': '2-1',
  '/uc/contract-money': '2-2',
  '/uc/record': '2-3',
  '/uc/contract-record': '2-6',
  '/uc/recharge': '2-4',
  '/uc/withdraw': '2-5',
  '/uc/trade': '2-6',
  '/uc/paydividends': '2-7',
  '/uc/blc': '2-8',
  '/uc/bjc': '2-9',
  '/uc/entrust/current': '3-1',
  '/uc/entrust/history': '3-2',
  '/uc/contract/entrust/current': '7-1',
  '/uc/contract/entrust/history': '7-2',
  '/uc/crowdfunding/list': '8-1',
  '/uc/ad': '4-1',
  '/uc/order': '4-2',
  '/uc/withdraw/address': '2-4',
  '/uc/ad/create': '4-3',
  '/uc/ad/update': '4-4',
  '/uc/innovation/myorders': '5-1',
  '/uc/innovation/myminings': '5-2'
})

// 菜单展开/收起处理
const handleOpen = (key, keyPath) => {
  // Element Plus el-menu accordion 模式下自动处理
}

const handleClose = (key, keyPath) => {
  // Element Plus el-menu accordion 模式下自动处理
}

// 切换抽屉
const toggleMemu = () => {
  drawerVisible.value = !drawerVisible.value
}

// 高亮菜单
const heightLightMenu = (path) => {
  const acName = routeArr[path] || '1-1'
  const opName = acName.split('-')[0]
  opennames.value = [opName]
  activename.value = acName
}

// 链接跳转
const link = (code) => {
  const routeMap = {
    '1-1': '/uc/safe',
    '2-1': '/uc/money',
    '3-1': '/uc/entrust/current',
    '7-1': '/uc/contract/entrust/current',
    '4-1': '/uc/ad',
    '5-1': '/uc/innovation/myorders',
    '5-2': '/uc/innovation/myminings',
    '8-1': '/uc/crowdfunding/list'
  }
  router.push(routeMap[code] || '/uc/safe')
}

// 生命周期
onMounted(() => {
  store.commit('navigate', 'nav-other')
  store.state.HeaderActiveName = '0'
  store.state.HeaderActiveName = '1-6'

  // 检查登录状态
  if (!localStorage.TOKEN || !localStorage.MEMBER) {
    ElMessage.success('请先登录')
    router.push('/login')
    return
  }

  // 初始化菜单高亮
  const path = router.currentRoute.value.path
  heightLightMenu(path)
})

// 监听路由变化
watch(() => router.currentRoute.value.path, (newPath) => {
  heightLightMenu(newPath)
})
</script>

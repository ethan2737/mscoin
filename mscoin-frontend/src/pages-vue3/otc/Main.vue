<template>
  <div class="content-wraps">
    <div class="containers">
      <div class="fiat">
        <div class="to_business">
          <h3>法币交易</h3>
          <span>便捷、安全、快速买卖数字货币</span>
          <a href="javascript:void(0)" @click="goBusiness">成为商家</a>
        </div>
      </div>
      <div class="content">
        <el-menu
          mode="horizontal"
          :default-active="activeMenuName"
          @select="handleMenuSelect"
          class="tradelist"
          style="background-color: #192330; border-radius: 4px;"
        >
          <el-menu-item
            v-for="(coin, index) in coins"
            :key="index"
            :index="'coin-' + index"
            style="border: none;"
          >
            {{ coin.unit }}
          </el-menu-item>
        </el-menu>
        <router-view />
      </div>
      <div class="advantage">
        <div style="display: flex; justify-content: center; padding: 30px;">
          <div style="width: 25%; text-align: center;">
            <div style="width: 50px; height: 50px; margin: 20px auto;">
              <img src="../../assets/images/price.png" alt="" style="width: 80%;" />
            </div>
            <div style="line-height: 30px; font-size: 16px; color: #fff;">市场一口价</div>
            <div style="padding: 20px 40px; line-height: 20px; font-size: 12px; color: #999;">根据市场价格实时波动</div>
          </div>
          <div style="width: 25%; text-align: center;">
            <div style="width: 50px; height: 50px; margin: 20px auto;">
              <img src="../../assets/images/poundage.png" alt="" style="width: 80%;" />
            </div>
            <div style="line-height: 30px; font-size: 16px; color: #fff;">完全免手续费</div>
            <div style="padding: 20px 40px; line-height: 20px; font-size: 12px; color: #999;">用户所见即所得，买卖价格外，无需任何平台手续费</div>
          </div>
          <div style="width: 25%; text-align: center;">
            <div style="width: 50px; height: 50px; margin: 20px auto;">
              <img src="../../assets/images/instant.png" alt="" style="width: 80%;" />
            </div>
            <div style="line-height: 30px; font-size: 16px; color: #fff;">即时成交</div>
            <div style="padding: 20px 40px; line-height: 20px; font-size: 12px; color: #999;">引入平台服务商家，智能匹配，成交订单，无须等待撮合</div>
          </div>
          <div style="width: 25%; text-align: center;">
            <div style="width: 50px; height: 50px; margin: 20px auto;">
              <img src="../../assets/images/platedanbao.png" alt="" style="width: 80%;" />
            </div>
            <div style="line-height: 30px; font-size: 16px; color: #fff;">平台担保</div>
            <div style="padding: 20px 40px; line-height: 20px; font-size: 12px; color: #999;">平台认证商家，安全有保障，24 小时客服为交易保驾护航</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 主页
 */
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()

const host = ''
const api = {
  otc: {
    coin: '/otc/coin'
  }
}

const coins = ref([])
const activeMenuName = ref('coin-1')

const isLogin = computed(() => {
  return localStorage.getItem('TOKEN') !== null
})

const goBusiness = () => {
  if (isLogin.value) {
    router.push('/uc/ident/business')
  } else {
    ElMessage.warning('请先登录')
  }
}

const handleMenuSelect = (menuName) => {
  if (menuName.startsWith('coin')) {
    const coin = coins.value[menuName.split('-')[1]]
    router.push('/otc/trade/' + coin.unit)
  } else {
    router.push('/otc/' + menuName)
  }
}

const activeMenu = () => {
  let coin = route.params.unit || 'USDT'
  coin = coin.toUpperCase()
  let index = 0
  coins.value.forEach((item, i) => {
    if (item.unit === coin) {
      index = i
    }
  })
  activeMenuName.value = 'coin-' + index
}

const init = () => {
  axios.post(`${host}${api.otc.coin}`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      coins.value = resp.data
      activeMenu()
    }
  }).catch(() => {
    console.error('获取币种列表失败')
  })
}

watch(() => route.params, () => {
  activeMenu()
}, { deep: true })

onMounted(() => {
  init()
})
</script>

<style scoped lang="scss">
.content-wraps {
  padding: 0 12%;
  padding-top: 60px;

  .containers {
    width: 100%;
    margin: 20px 0;

    .fiat {
      border-radius: 5px;
      height: 250px;
      background: url("../../assets/images/otc_bg.jpg") no-repeat center center;
      background-size: 100%;
      display: flex;
      justify-content: center;
      align-items: center;

      .to_business {
        color: #fff;
        text-align: center;

        h3 {
          font-size: 46px;
          letter-spacing: 20px;
        }

        span {
          font-size: 20px;
          letter-spacing: 10px;
          display: block;
        }

        a {
          width: 220px;
          height: 45px;
          display: inline-block;
          background: #d0b387;
          border-radius: 5px;
          font-size: 20px;
          line-height: 45px;
          color: #000;
          margin-top: 20px;
          text-decoration: none;

          &:hover {
            background: #c9a976;
          }
        }
      }
    }

    .content {
      width: 100%;
      margin: 20px auto;
      background-color: #192330;
      border-radius: 4px;
    }

    .advantage {
      background-color: #192330;
      border-radius: 4px;

      ul {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 30px;

        li {
          width: 25%;
          list-style-type: none;
          min-height: 190px;

          div {
            text-align: center;
          }

          div.image {
            width: 50px;
            height: 50px;
            margin: 20px auto;

            img {
              width: 80%;
              vertical-align: middle;
            }
          }

          div.title {
            line-height: 30px;
            font-size: 16px;
            color: #fff;
          }

          div.content1 {
            padding: 20px 40px;
            line-height: 20px;
            font-size: 12px;
            color: #999;
          }
        }
      }
    }
  }
}
</style>

<style lang="scss">
.content-wraps {
  .containers {
    .content {
      ul.tradelist {
        &.el-menu {
          &.el-menu--horizontal {
            border-bottom: none;

            .el-menu-item {
              border: none;

              &:hover {
                color: #f0ac19;
              }

              &.is-active {
                color: #f0ac19;
              }
            }
          }
        }
      }
    }
  }
}
</style>

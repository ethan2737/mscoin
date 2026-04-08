<template>
  <div class="content-wrap" id="List">
    <div class="container">
      <div class="row">
        <div class="leftmenu left-box">
          <div class="user-info">
            <div class="avatar-box">
              <div class="user-face user-avatar-public">
                <img v-if="user.avatar != null && user.avatar != ''" :src="user.avatar" width="60px" height="60px" style="border-radius: 50%">
                <span v-else class="user-avatar-in">{{ usernameS }}</span>
              </div>
              <div class="user-name">
              </div>
            </div>
            <span class="ml10">
              {{ user.username }}
            </span>
          </div>
          <div class="deal-market-info">
            <p v-if="user.emailVerified === 1">
              <i class="iconfont icon-youxiang111"></i>
              <span class="unmarket">{{ $t('otc.checkuser.emaildone') }}</span>
            </p>
            <p v-else>
              <i class="iconfont icon-youxiang"></i>
              <span class="unmarket">{{ $t('otc.checkuser.emailundo') }}</span>
            </p>
            <p v-if="user.phoneVerified === 1">
              <i class="iconfont icon-dianhua111"></i>
              <span class="">{{ $t('otc.checkuser.teldone') }}</span>
            </p>
            <p v-else>
              <i class="iconfont icon-dianhua"></i>
              <span class="">{{ $t('otc.checkuser.telundo') }}</span>
            </p>
            <p v-if="user.realVerified === 1">
              <i class="iconfont icon-renzheng111"></i>
              <span class="">{{ $t('otc.checkuser.idcarddone') }}</span>
            </p>
            <p v-else>
              <i class="iconfont icon-renzheng"></i>
              <span class="unmarket">{{ $t('otc.checkuser.idcardundo') }}</span>
            </p>
          </div>
        </div>
        <div class="right-safe">
          <div class="trade-right-box">
            <div class="trade-price">
              <div class="tit row">
                <div class="col">{{ $t('otc.checkuser.language') }}: {{ $t('otc.checkuser.languagetext') }}</div>
                <div class="col">{{ $t('otc.checkuser.registtime') }}: {{ user.createTime }}</div>
                <div class="col">{{ $t('otc.checkuser.exchangetimes') }}: {{ user.transactions }}</div>
              </div>
            </div>
          </div>
          <div class="trade-operation">
            <div class="trade-group merchant-top">
              <i class="merchant-icon tips"></i>
              <span class="tips-word">{{ user.username ? strpro(user.username) : '' }}{{ $t('otc.checkuser.exchangeinfo') }}</span>
            </div>
          </div>
          <div class="demo-tabs-style1 tabbox">
            <el-tabs v-model="activeTab">
              <el-tab-pane :label="$t('otc.checkuser.tablabel1')" name="name1">
                <div class="order-table">
                  <el-table :columns="tableColumnsSell" :data="tableOrderSell" v-loading="loading" border>
                    <el-table-column :label="$t('otc.checkuser.col_symbol')" prop="unit" />
                    <el-table-column :label="$t('otc.checkuser.col_paymode')" prop="payMode" />
                    <el-table-column :label="$t('otc.checkuser.col_num')" prop="remainAmount" />
                    <el-table-column :label="$t('otc.checkuser.col_price') + '/BTC'" prop="price" width="170">
                      <template #default="{ row }">
                        <div>
                          <p class="price">{{ row.price }} CNY</p>
                          <p class="price2">{{ row.minLimit }} - {{ row.maxLimit }} CNY</p>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column :label="$t('otc.checkuser.col_created')" prop="createTime" width="160" />
                    <el-table-column :label="$t('otc.checkuser.col_operate')" width="120">
                      <template #default="{ row }">
                        <el-button
                          type="success"
                          long
                          @click="handleBuy(row)"
                          style="margin-right: 8px; width: 80%;"
                        >
                          {{ $t('otc.checkuser.buyin') }}
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </el-tab-pane>
              <el-tab-pane :label="$t('otc.checkuser.tablabel2')" name="name2">
                <div class="order-table">
                  <el-table :columns="tableColumnsBuy" :data="tableOrderBuy" v-loading="loading" border>
                    <el-table-column :label="$t('otc.checkuser.col_symbol')" prop="unit" />
                    <el-table-column :label="$t('otc.checkuser.col_paymode')" prop="payMode" />
                    <el-table-column :label="$t('otc.checkuser.col_num')" prop="remainAmount" />
                    <el-table-column :label="$t('otc.checkuser.col_price') + '/BTC'" prop="price" width="170">
                      <template #default="{ row }">
                        <div>
                          <p class="price">{{ row.price }} CNY</p>
                          <p class="price2">{{ row.minLimit }} - {{ row.maxLimit }} CNY</p>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column :label="$t('otc.checkuser.col_created')" prop="createTime" width="160" />
                    <el-table-column :label="$t('otc.checkuser.col_operate')" width="120">
                      <template #default="{ row }">
                        <el-button
                          type="danger"
                          long
                          @click="handleSell(row)"
                          style="margin-right: 8px; width: 80%;"
                        >
                          {{ $t('otc.checkuser.sellout') }}
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 用户审核页面
 */
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElTabs, ElTabPane, ElButton } from 'element-plus'
import axios from 'axios'
import { useStore, useRouter } from 'vue-router/composables'

const store = useStore()
const router = useRouter()

const host = 'http://localhost'

const loading = ref(true)
const hasRealName = ref(false)
const usernameS = ref('')
const activeTab = ref('name1')

const user = ref({
  username: '',
  email: true,
  mobileNo: false,
  idCard: true,
  avatar: '',
  emailVerified: 0,
  phoneVerified: 0,
  realVerified: 0,
  createTime: '',
  transactions: 0
})

const tableOrderSell = ref([])
const tableOrderBuy = ref([])

const isLogin = computed(() => store.getters.isLogin)
const member = computed(() => store.getters.member)

const tableColumnsSell = computed(() => [
  { label: '币种', prop: 'unit' },
  { label: '支付方式', prop: 'payMode' },
  { label: '数量', prop: 'remainAmount' },
  { label: '价格/BTC', prop: 'price', width: 170 },
  { label: '创建时间', prop: 'createTime', width: 160 },
  { label: '操作', width: 120 }
])

const tableColumnsBuy = computed(() => [
  { label: '币种', prop: 'unit' },
  { label: '支付方式', prop: 'payMode' },
  { label: '数量', prop: 'remainAmount' },
  { label: '价格/BTC', prop: 'price', width: 170 },
  { label: '创建时间', prop: 'createTime', width: 160 },
  { label: '操作', width: 120 }
])

const strpro = (str) => {
  let newStr = str
  str = str.slice(1)
  const re = /[\D\d]*/g
  const str2 = str.replace(re, function (s) {
    let result = ''
    for (let i = 0; i < s.length; i++) {
      result += '*'
    }
    return result
  })
  return newStr.slice(0, 1) + str2
}

const handleBuy = (row) => {
  if (!isLogin.value) {
    router.push('/login')
  } else if (!member.value.realName) {
    ElMessage.error('请先完成实名认证')
    setTimeout(() => {
      router.push('/uc/safe')
    }, 2000)
  } else {
    router.push('/otc/tradeInfo?tradeId=' + row.advertiseId)
  }
}

const handleSell = (row) => {
  if (!isLogin.value) {
    router.push('/login')
  } else if (!member.value.realName) {
    ElMessage.error('请先完成实名认证')
    setTimeout(() => {
      router.push('/uc/safe')
    }, 2000)
  } else {
    router.push('/otc/tradeInfo?tradeId=' + row.advertiseId)
  }
}

const getAdv = () => {
  axios.post(`${host}/otc/advertise/member`, {
    name: router.currentRoute.value.query.id
  }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      loading.value = false
      tableOrderBuy.value = resp.data.buy || []
      tableOrderSell.value = resp.data.sell || []
      user.value = resp.data
      usernameS.value = (user.value.username + '').replace(/^\s+|\s+$/g, '').slice(0, 1)
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

onMounted(() => {
  getAdv()
})
</script>

<style scoped lang="scss">
.container {
  padding-top: 30px;
  margin: 0 auto;
  width: 1200px;
  background: #192330;
  margin-bottom: 20px;
}

.content-wrap {
  min-height: 600px;
  padding-top: 80px;
}

.row {
  display: flex;
  flex-wrap: wrap;

  .leftmenu {
    flex: 0 0 16.666667%;
    max-width: 16.666667%;
  }

  .right-safe {
    flex: 0 0 83.333333%;
    max-width: 83.333333%;
  }
}

.leftmenu {
  margin-bottom: 60px;
  background-color: #192330;
  position: relative;
  min-height: 1px;
  padding: 50px 15px 50px 10px;
}

.left-box .user-info {
  display: flex;
  align-items: center;
  padding-bottom: 15px;
  border-bottom: 1px dashed #ebeff5;
}

.avatar-box {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.user-avatar-public {
  height: 65px;
  width: 65px;
  box-shadow: 0 1px 5px 0 rgba(71, 78, 114, 0.24);
  position: relative;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;

  > .user-avatar-in {
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    background: #f0a70a;
    height: 60px;
    width: 60px;
    color: #fff;
  }
}

.left-box span.ml10 {
  color: #fff;
  margin-left: 5px;
}

.left-box .deal-market-info {
  padding: 20px 0 20px 20px;
  border-bottom: 1px dashed #ebeff5;

  p {
    display: flex;
    align-items: center;
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    .iconfont {
      margin-right: 20px;
      font-size: 20px;

      &:before {
        background-size: 100% 100%;
        width: 20px;
        height: 20px;
        display: inline-block;
        content: "";
      }
    }
  }
}

.icon-youxiang:before {
  background-image: url(../../assets/img/t1-1.png);
}

.icon-youxiang111:before {
  background-image: url(../../assets/img/t1-2.png);
}

.icon-dianhua:before {
  background-image: url(../../assets/img/t2-1.png);
}

.icon-dianhua111:before {
  background-image: url(../../assets/img/t2-2.png);
}

.icon-renzheng:before {
  background-image: url(../../assets/img/t3-1.png);
}

.icon-renzheng111:before {
  background-image: url(../../assets/img/t3-2.png);
}

.tabbox {
  margin-left: 20px;
  padding: 20px 15px;
}

.merchant-top {
  height: 50px;
  display: flex;
  align-items: center;
  padding: 0 15px;
  color: #fff;
  margin-left: 20px;
}

.merchant-icon {
  width: 4px;
  height: 22px;
  margin-right: 10px;
  background: #f0a70a;
  display: inline-block;
  margin-left: 4px;
}

.tips-word {
  flex-grow: 2;
  text-align: left;
  font-size: 14px;
}

.tit {
  display: flex;

  .col {
    color: #a2a2a2;
    flex: 1;
  }
}

.trade-right-box {
  margin-left: 33px;
  margin-right: 15px;
  text-align: left;

  .trade-price {
    padding: 15px 0;
    border: 1px solid #27313e;
    margin-bottom: 20px;
  }
}

.trade-operation {
  padding: 20px;
  border: 1px solid #27313e;
  margin-bottom: 20px;
}

.order-table {
  :deep(.el-table) {
    background: transparent;

    th.el-table__cell {
      background: #27313e;
      color: #fff;
    }

    td.el-table__cell {
      background: transparent;
      color: #fff;
    }
  }
}

.price {
  font-size: 16px;
  font-weight: bolder;
  color: #444f71;
}

.price2 {
  font-size: 12px;
  color: #8994a3;
  margin-top: 0;
}
</style>

<style lang="scss">
.right-safe {
  .el-tabs {
    .el-tabs__header {
      .el-tabs__nav-wrap::after {
        display: none;
      }

      .el-tabs__item {
        &:hover {
          color: #f0a70a;
        }

        &.is-active {
          color: #f0a70a;
        }
      }

      .el-tabs__nav {
        .el-tabs__active-bar {
          background: #f0a70a;
        }
      }
    }

    .el-tabs__content {
      padding: 0;
    }
  }
}
</style>

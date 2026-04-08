<template>
  <div class="nav-right tradeCenter">
    <section class="list-content">
      <el-tabs v-model="activeTab" @tab-change="handleTabChange">
        <el-tab-pane :label="$t('otc.buyin')" name="buy">
          <div class="table-responsive list-table">
            <el-table :data="advertiment.ask.rows" v-loading="loading" border>
              <el-table-column :label="$t('otc.merchant')" min-width="180">
                <template #default="{ row }">
                  <div style="display: flex; align-items: center;">
                    <div class="user-face user-avatar-public">
                      <img v-if="row.avatar" :src="row.avatar" style="width: 45px; height: 45px; border-radius: 50%;" />
                      <span v-else class="user-avatar-in">{{ getAvatarInitial(row.memberName) }}</span>
                    </div>
                    <div style="margin-left: 10px;">
                      <el-button type="link" @click="goToCheckUser(row.memberName)" style="padding: 0;">
                        {{ strpro(row.memberName) }}
                      </el-button>
                      <div v-if="row.level === 2" style="display: inline-block;">
                        <img src="../../assets/images/business_v.png" style="height: 20px;" />
                      </div>
                    </div>
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.volume')" prop="transactions" width="100" align="center" />
              <el-table-column :label="$t('otc.paymethod')" prop="payMode" align="center" />
              <el-table-column :label="$t('otc.amount')" prop="remainAmount" align="center" />
              <el-table-column :label="$t('otc.limit')" align="center">
                <template #default="{ row }">
                  {{ row.minLimit }} - {{ row.maxLimit }} CNY
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.price')" align="center">
                <template #default="{ row }">
                  {{ row.price }} CNY
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.operate')" width="100" align="center">
                <template #default="{ row }">
                  <el-button
                    :type="row.advertiseType === 0 ? 'danger' : 'success'"
                    link
                    @click="handleTrade(row)"
                  >
                    {{ row.advertiseType === 0 ? $t('otc.sell') : $t('otc.buy') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <div class="page_change" style="margin: 10px; overflow: hidden;">
              <div style="float: right;">
                <el-pagination
                  v-if="advertiment.ask.totalElement > 0"
                  layout="prev, pager, next"
                  :page-size="advertiment.ask.pageNumber"
                  :total="advertiment.ask.totalElement"
                  :current-page="advertiment.ask.currentPage"
                  @current-change="changePage"
                />
              </div>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane :label="$t('otc.sellout')" name="sell">
          <div class="table-responsive list-table">
            <el-table :data="advertiment.bid.rows" v-loading="loading" border>
              <el-table-column :label="$t('otc.merchant')" min-width="180">
                <template #default="{ row }">
                  <div style="display: flex; align-items: center;">
                    <div class="user-face user-avatar-public">
                      <img v-if="row.avatar" :src="row.avatar" style="width: 45px; height: 45px; border-radius: 50%;" />
                      <span v-else class="user-avatar-in">{{ getAvatarInitial(row.memberName) }}</span>
                    </div>
                    <div style="margin-left: 10px;">
                      <el-button type="link" @click="goToCheckUser(row.memberName)" style="padding: 0;">
                        {{ strpro(row.memberName) }}
                      </el-button>
                      <div v-if="row.level === 2" style="display: inline-block;">
                        <img src="../../assets/images/business_v.png" style="height: 20px;" />
                      </div>
                    </div>
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.volume')" prop="transactions" width="100" align="center" />
              <el-table-column :label="$t('otc.paymethod')" prop="payMode" align="center" />
              <el-table-column :label="$t('otc.amount')" prop="remainAmount" align="center" />
              <el-table-column :label="$t('otc.limit')" align="center">
                <template #default="{ row }">
                  {{ row.minLimit }} - {{ row.maxLimit }} CNY
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.price')" align="center">
                <template #default="{ row }">
                  {{ row.price }} CNY
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.operate')" width="100" align="center">
                <template #default="{ row }">
                  <el-button
                    :type="row.advertiseType === 0 ? 'danger' : 'success'"
                    link
                    @click="handleTrade(row)"
                  >
                    {{ row.advertiseType === 0 ? $t('otc.sell') : $t('otc.buy') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <div class="page_change" style="margin: 10px; overflow: hidden;">
              <div style="float: right;">
                <el-pagination
                  v-if="advertiment.bid.totalElement > 0"
                  layout="prev, pager, next"
                  :page-size="advertiment.bid.pageNumber"
                  :total="advertiment.bid.totalElement"
                  :current-page="advertiment.bid.currentPage"
                  @current-change="changePage"
                />
              </div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </section>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 交易列表页面
 */
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElTabs, ElTabPane, ElButton, ElIcon } from 'element-plus'
import axios from 'axios'
import { useStore, useRouter } from 'vue-router/composables'

const store = useStore()
const router = useRouter()

const host = 'http://localhost'
const api = {
  otc: {
    advertise: '/otc/advertise/list'
  }
}

const showBorder = ref(false)
const showStripe = ref(false)
const showHeader = ref(true)
const fixedHeader = ref(false)
const loading = ref(true)
const dataCount = ref(10)
const tableSize = ref('large')
const activeTab = ref('buy')

const advertiment = reactive({
  ask: {
    rows: [],
    currentPage: 1,
    totalPage: 1,
    pageNumber: 10,
    totalElement: 0
  },
  bid: {
    rows: [],
    currentPage: 1,
    totalPage: 1,
    pageNumber: 10,
    totalElement: 0
  }
})

const isLogin = computed(() => store.getters.isLogin)
const member = computed(() => store.getters.member)
const coin = computed(() => router.currentRoute.params.pathMatch)
const lang = computed(() => store.state.lang)

const getAvatarInitial = (name) => {
  return (name + '').replace(/^\s+|\s+$/g, '').slice(0, 1)
}

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

const goToCheckUser = (memberName) => {
  if (isLogin.value) {
    router.push('/checkuser?id=' + memberName)
  } else {
    router.push('/login')
  }
}

const handleTrade = (row) => {
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

const handleTabChange = (tab) => {
  if (tab === 'sell') {
    loadAd(1, 0, advertiment.bid)
  } else {
    loadAd(1, 1, advertiment.ask)
  }
}

const changePage = (page) => {
  if (activeTab.value === 'sell') {
    loadAd(page, 0, advertiment.bid)
  } else {
    loadAd(page, 1, advertiment.ask)
  }
}

const loadAd = (pageNo, advertiseType, table) => {
  table.rows = []
  table.totalElement = 0
  table.currentPage = pageNo

  const params = {
    pageNo,
    pageSize: table.pageNumber,
    advertiseType,
    unit: coin.value
  }

  loading.value = true
  axios.post(`${host}${api.otc.advertise}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0 && resp.data.context) {
      table.rows = resp.data.context
      table.totalElement = resp.data.totalElement
    } else if (resp.code !== 0) {
      ElMessage.error(resp.message)
    }
    loading.value = false
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const reloadAd = () => {
  loadAd(1, 0, advertiment.bid)
  loadAd(1, 1, advertiment.ask)
}

// Watch for coin and lang changes
watch([coin, lang], () => {
  reloadAd()
})

onMounted(() => {
  reloadAd()
})
</script>

<style scoped lang="scss">
.nav-right {
  color: #26264c;
  padding-right: 0;

  .list-content {
    color: #fff;

    :deep(.el-tabs) {
      .el-tabs__header {
        border-bottom: none;

        .el-tabs__nav-wrap::after {
          display: none;
        }

        .el-tabs__item {
          &:hover {
            color: #f0ac19;
          }

          &.is-active {
            color: #f0ac19;
          }
        }

        .el-tabs__nav {
          .el-tabs__active-bar {
            background: #f0ac19;
          }
        }
      }

      .el-tabs__content {
        .el-tabpane {
          .el-table {
            background: transparent;

            th.el-table__cell {
              background: #27313e;
              color: #fff;
            }

            td.el-table__cell {
              background: transparent;
              color: #fff;
            }

            .el-button--link {
              color: #f0a70a;
            }
          }

          .page_change {
            margin: 10px;
            overflow: hidden;
          }
        }
      }
    }
  }
}

.user-face.user-avatar-public {
  background: #fff;
  border-radius: 50%;
  height: 45px;
  width: 45px;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 1px 5px 0 rgba(71, 78, 114, 0.24);
  position: relative;
  margin-right: 10px;

  .user-avatar-in {
    background: #f0a70a;
    border-radius: 50%;
    height: 35px;
    width: 35px;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
  }
}
</style>

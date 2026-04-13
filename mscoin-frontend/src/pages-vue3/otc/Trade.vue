<template>
  <div class="nav-right tradeCenter">
    <section class="list-content">
      <el-tabs v-model="tabPage" @tab-change="handleTabChange">
        <el-tab-pane :label="$t('otc.buyin')" name="buy">
          <div class="table-responsive list-table otc-offer-board">
            <el-table :data="advertiment.ask.rows" v-loading="loading" class="offer-table">
              <el-table-column :label="$t('otc.merchant')" min-width="180">
                <template #default="{ row }">
                  <div class="merchant-cell">
                    <div class="user-face user-avatar-public">
                      <img v-if="row.avatar" :src="row.avatar" class="merchant-avatar-image" />
                      <span v-else class="user-avatar-in">{{ getAvatarInitial(row.memberName) }}</span>
                    </div>
                    <div class="merchant-meta">
                      <el-button type="link" class="merchant-link" @click="goToCheckUser(row.memberName)">
                        {{ row.memberName }}
                      </el-button>
                      <img
                        v-if="row.level === 2"
                        src="../../assets/images/business_v.png"
                        class="merchant-badge"
                      />
                    </div>
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.volume')" prop="transactions" width="100" align="center" />
              <el-table-column :label="$t('otc.paymethod')" prop="payMode" align="center" />
              <el-table-column :label="$t('otc.amount')" prop="remainAmount" align="center" />
              <el-table-column label="限额" align="center">
                <template #default="{ row }">
                  {{ row.minLimit }} - {{ row.maxLimit }} CNY
                </template>
              </el-table-column>
              <el-table-column label="价格" align="center">
                <template #default="{ row }">
                  {{ row.price }} CNY
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.operate')" width="100" align="center">
                <template #default="{ row }">
                  <el-button
                    :type="row.advertiseType === 0 ? 'danger' : 'success'"
                    class="trade-action"
                    @click="handleTrade(row)"
                  >
                    {{ row.advertiseType === 0 ? $t('otc.sell') : $t('otc.buy') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <div class="page_change">
              <div class="page_change_inner">
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
          <div class="table-responsive list-table otc-offer-board">
            <el-table :data="advertiment.bid.rows" v-loading="loading" class="offer-table">
              <el-table-column :label="$t('otc.merchant')" min-width="180">
                <template #default="{ row }">
                  <div class="merchant-cell">
                    <div class="user-face user-avatar-public">
                      <img v-if="row.avatar" :src="row.avatar" class="merchant-avatar-image" />
                      <span v-else class="user-avatar-in">{{ getAvatarInitial(row.memberName) }}</span>
                    </div>
                    <div class="merchant-meta">
                      <el-button type="link" class="merchant-link" @click="goToCheckUser(row.memberName)">
                        {{ row.memberName }}
                      </el-button>
                      <img
                        v-if="row.level === 2"
                        src="../../assets/images/business_v.png"
                        class="merchant-badge"
                      />
                    </div>
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.volume')" prop="transactions" width="100" align="center" />
              <el-table-column :label="$t('otc.paymethod')" prop="payMode" align="center" />
              <el-table-column :label="$t('otc.amount')" prop="remainAmount" align="center" />
              <el-table-column label="限额" align="center">
                <template #default="{ row }">
                  {{ row.minLimit }} - {{ row.maxLimit }} CNY
                </template>
              </el-table-column>
              <el-table-column label="价格" align="center">
                <template #default="{ row }">
                  {{ row.price }} CNY
                </template>
              </el-table-column>
              <el-table-column :label="$t('otc.operate')" width="100" align="center">
                <template #default="{ row }">
                  <el-button
                    :type="row.advertiseType === 0 ? 'danger' : 'success'"
                    class="trade-action"
                    @click="handleTrade(row)"
                  >
                    {{ row.advertiseType === 0 ? $t('otc.sell') : $t('otc.buy') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <div class="page_change">
              <div class="page_change_inner">
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
 * Vue 3 迁移 - OTC 交易页面
 */
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElTabs, ElTabPane, ElButton, ElPagination } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'
import { buildOtcCheckUserPath, buildOtcTradeInfoPath } from './route-helpers'

const store = useStore()
const route = useRoute()
const router = useRouter()

const host = ''
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
const tabPage = ref('buy')

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
  },
  columns: []
})

const isLogin = computed(() => store.getters.isLogin)
const member = computed(() => store.getters.member)
const coin = computed(() => route.params.unit)
const lang = computed(() => store.state.lang)

const getAvatarInitial = (name) => {
  return (name + '').replace(/^\s+|\s+$/g, '').slice(0, 1)
}

const goToCheckUser = (memberName) => {
  if (isLogin.value) {
    router.push(buildOtcCheckUserPath(memberName))
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
    router.push(buildOtcTradeInfoPath(row.advertiseId))
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
  if (tabPage.value === 'sell') {
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
    color: #dfe7f5;

    :deep(.el-tabs) {
      .el-tabs__header {
        margin-bottom: 18px;
        border-bottom: none;

        .el-tabs__nav-wrap::after {
          display: none;
        }

        .el-tabs__nav-wrap {
          padding: 0 18px;
        }

        .el-tabs__item {
          min-height: 42px;
          border-radius: 999px;
          background: transparent;
          color: #8d9bb0;
          font-size: 15px;
          font-weight: 600;
          transition: all 0.2s ease;

          &:hover {
            color: #f0ac19;
            background: transparent;
          }

          &.is-active {
            color: #f0ac19;
            background: transparent;
          }
        }

        .el-tabs__nav {
          gap: 12px;

          .el-tabs__active-bar {
            background: #f0ac19;
            height: 3px;
            border-radius: 999px;
          }
        }
      }

      .el-tabs__content {
        overflow: visible;
      }
    }
  }
}

.otc-offer-board {
  padding: 16px 18px 20px;
  border: 1px solid rgba(96, 118, 145, 0.22);
  border-radius: 18px;
  background:
    linear-gradient(180deg, rgba(34, 45, 60, 0.96) 0%, rgba(21, 29, 40, 0.98) 100%);
  box-shadow:
    inset 0 1px 0 rgba(255, 255, 255, 0.03),
    0 18px 40px rgba(5, 10, 18, 0.22);

  :deep(.el-table) {
    --el-table-tr-bg-color: transparent;
    --el-table-bg-color: transparent;
    --el-bg-color: transparent;
    --el-fill-color-lighter: rgba(255, 255, 255, 0.02);
    --el-table-border-color: rgba(92, 112, 138, 0.18);
    --el-table-row-hover-bg-color: rgba(240, 172, 25, 0.08);
    color: #f4f7fb;
    background: transparent;

    &::before {
      display: none;
    }
  }

  :deep(.el-table__inner-wrapper),
  :deep(.el-table__header-wrapper),
  :deep(.el-table__body-wrapper),
  :deep(.el-table__empty-block) {
    background: transparent;
  }

  :deep(th.el-table__cell) {
    background: rgba(255, 255, 255, 0.04);
    color: #8d9bb0;
    border-bottom-color: rgba(92, 112, 138, 0.24);
    font-size: 12px;
    font-weight: 600;
    letter-spacing: 0.02em;
    text-transform: uppercase;
  }

  :deep(td.el-table__cell) {
    background: transparent;
    color: #f4f7fb;
    border-bottom-color: rgba(92, 112, 138, 0.14);
  }

  :deep(.el-table__row:last-child td.el-table__cell) {
    border-bottom: none;
  }
}

.merchant-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.merchant-avatar-image {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
}

.merchant-meta {
  display: flex;
  align-items: center;
  gap: 6px;
}

.merchant-link {
  padding: 0;
  border: none;
  background: transparent;
  box-shadow: none;
  color: #f7fbff;
  font-weight: 600;
  font-size: 14px;
  line-height: 1.2;
}

.merchant-link:hover {
  color: #f0ac19;
}

.merchant-badge {
  height: 14px;
}

.trade-action {
  min-width: 72px;
  border-radius: 6px;
  font-weight: 600;
}

.page_change {
  margin-top: 18px;
}

.page_change_inner {
  display: flex;
  justify-content: flex-end;
}

.user-face.user-avatar-public {
  background: linear-gradient(180deg, #f8fbff 0%, #dbe5f4 100%);
  border-radius: 50%;
  height: 32px;
  width: 32px;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 1px 5px 0 rgba(71, 78, 114, 0.24);
  position: relative;
  margin-right: 10px;

  .user-avatar-in {
    background: #f0a70a;
    border-radius: 50%;
    height: 24px;
    width: 24px;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    font-size: 12px;
  }
}
</style>

<style lang="scss">
.nav-right {
  .list-content {
    .el-table {
      th.el-table__cell,
      td.el-table__cell {
        text-align: center;
      }
    }

    .page_change {
      .el-pagination {
        --el-pagination-bg-color: transparent;
        --el-pagination-button-color: #dfe7f5;
        --el-pagination-hover-color: #f0ac19;
      }

      .btn-prev,
      .btn-next,
      .number {
        background: rgba(255, 255, 255, 0.04);
        border-radius: 10px;
      }

      .number.is-active {
        background: rgba(240, 172, 25, 0.16);
        color: #f0ac19;
      }
    }
  }
}
</style>

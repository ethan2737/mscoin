<template>
  <div class="nav-rights">
    <div class="nav-right">
      <div class="bill_box_order">
        <div class="order_box">
          <div class="tabs-header">
            <el-input
              v-model="ordKeyword"
              :placeholder="$t('uc.otcorder.searchtip')"
              class="search-input"
              @keyup.enter="handleSearch"
            >
              <template #suffix>
                <el-icon @click="handleSearch" style="cursor: pointer;"><Search /></el-icon>
              </template>
            </el-input>
          </div>
          <el-tabs v-model="activeTab" @tab-click="handleTabClick">
            <el-tab-pane :label="$t('uc.otcorder.unpaid')" name="1" />
            <el-tab-pane :label="$t('uc.otcorder.paided')" name="2" />
            <el-tab-pane :label="$t('uc.otcorder.finished')" name="3" />
            <el-tab-pane :label="$t('uc.otcorder.canceled')" name="0" />
            <el-tab-pane :label="$t('uc.otcorder.appealing')" name="4" />
          </el-tabs>
          <div class="order-table">
                <el-table :data="tableOrder" v-loading="loading" border style="width: 100%">
                  <el-table-column :label="$t('uc.otcorder.orderno')" min-width="60" align="center">
                    <template #default="{ row }">
                      <el-link type="primary" @click="goChat(row.orderSn)">{{ row.orderSn }}</el-link>
                    </template>
                  </el-table-column>
                  <el-table-column prop="createTime" :label="$t('uc.otcorder.created')" min-width="90" align="center" />
                  <el-table-column prop="unit" :label="$t('uc.otcorder.symbol')" align="center" />
                  <el-table-column :label="$t('uc.otcorder.type')" align="center">
                    <template #default="{ row }">
                      {{ row.type === 0 ? $t('uc.otcorder.type_buy') : $t('uc.otcorder.type_sell') }}
                    </template>
                  </el-table-column>
                  <el-table-column :label="$t('uc.otcorder.tradename')" align="center">
                    <template #default="{ row }">
                      <el-link type="primary" @click="goCheckUser(row.name)">{{ row.name }}</el-link>
                    </template>
                  </el-table-column>
                  <el-table-column prop="amount" :label="$t('uc.otcorder.amount')" align="center" />
                  <el-table-column prop="money" :label="$t('uc.otcorder.money')" align="center" />
                  <el-table-column prop="commission" :label="$t('uc.otcorder.fee')" align="center" />
                </el-table>
                <div class="pagination-container">
                  <el-pagination
                    v-if="totalPage > 0"
                    layout="total, prev, pager, next"
                    :total="totalNum"
                    :current-page="currentPage"
                    :page-size="pageSize"
                    @current-change="changePage"
                  />
                </div>
              </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 我的订单页面
 */
import { ref, reactive, inject, onMounted } from 'vue'
import { ElMessage, ElLink } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import axios from 'axios'
import { buildOtcChatPath, buildOtcCheckUserPath } from '../otc/route-helpers'

const router = inject('router')

const host = ''

const activeTab = ref('1')
const ordKeyword = ref('')
const choseBtn = ref(0)
const whichItem = ref(1)
const tableOrder = ref([])
const loading = ref(true)
const totalPage = ref(0)
const pageSize = ref(10)
const totalNum = ref(0)
const currentPage = ref(1)

const changePage = (pageNo) => {
  if (pageNo > 0) pageNo = pageNo - 1
  getOrder(whichItem.value, pageNo)
}

const getOrder = (status, pageNo) => {
  loading.value = true
  tableOrder.value = []
  const params = {
    status: status,
    pageNo: pageNo,
    pageSize: pageSize.value
  }
  currentPage.value = pageNo + 1

  axios.post(`${host}/otc/order/self`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0 && resp.data.content) {
      tableOrder.value = resp.data.content
      totalPage.value = resp.data.totalPages
      totalNum.value = resp.data.totalElements
    } else {
      ElMessage.error(resp.message || $t('common.logintip'))
    }
    loading.value = false
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
    loading.value = false
  })
}

const handleSearch = () => {
  tableOrder.value = []
  const params = {
    status: whichItem.value,
    pageNo: 0,
    pageSize: pageSize.value
  }
  if (ordKeyword.value !== '') {
    params.orderSn = ordKeyword.value
  }
  currentPage.value = 1

  axios.post(`${host}/otc/order/self`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0 && resp.data.content) {
      tableOrder.value = resp.data.content
      totalPage.value = resp.data.totalPages
      totalNum.value = resp.data.totalElements
    } else {
      ElMessage.error(resp.message)
    }
    loading.value = false
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
    loading.value = false
  })
}

const handleTabClick = (tab) => {
  const name = tab.props.name
  if (name === '1') {
    whichItem.value = 1
  } else if (name === '2') {
    whichItem.value = 2
  } else if (name === '3') {
    whichItem.value = 3
  } else if (name === '0') {
    whichItem.value = 0
  } else if (name === '4') {
    whichItem.value = 4
  }
  changePage(1)
}

const goChat = (tradeId) => {
  router.push(buildOtcChatPath(tradeId))
}

const goCheckUser = (id) => {
  router.push(buildOtcCheckUserPath(id))
}

onMounted(() => {
  changePage(1)
})
</script>

<style scoped lang="scss">
.nav-rights {
  .nav-right {
    height: auto;
    overflow: hidden;
    padding: 0 15px;

    .bill_box_order {
      width: 100%;
      padding-left: 20px;
      height: auto;

      .order_box {
        text-align: left;

        :deep(.el-tabs__header) {
          border-color: #273742 !important;

          .el-tabs__item {
            color: #fff;

            &:hover {
              color: #f0a70a;
            }

            &.is-active {
              color: #f0a70a;
            }
          }

          .el-tabs__nav-wrap::after {
            background-color: #273742;
          }

          .el-tabs__active-bar {
            background-color: #f0a70a;
          }
        }

        .search-input {
          width: 250px;
          margin-right: 20px;
        }

        .order-table {
          margin-top: 20px;

          .pagination-container {
            margin: 10px;
            overflow: hidden;
            float: right;
          }
        }
      }
    }
  }
}
</style>

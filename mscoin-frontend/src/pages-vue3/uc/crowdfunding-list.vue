<template>
  <div class="entrustcurrent">
    <div class="table">
      <el-tabs v-model="activeTab" @tab-click="handleTabClick" style="background: transparent;">
        <el-tab-pane :label="$t('uc.myCf.medicalCrowdfunding')" name="1" />
        <el-tab-pane :label="$t('uc.myCf.otherCrowdfunding')" name="2" />
        <el-tab-pane :label="$t('uc.myCf.offlineCharity')" name="3" />
        <el-tab-pane :label="$t('uc.myCf.myDonation')" name="4" />
      </el-tabs>

      <!-- Tab 1: 医疗众筹 -->
      <el-table v-show="activeTab === '1'" :data="ylList" :no-data-text="$t('common.nodata')" v-loading="loading" border>
        <el-table-column :label="$t('uc.myCf.title')" prop="fundingTitle" align="center" />
        <el-table-column :label="$t('uc.myCf.time')" align="center" width="160">
          <template #default="{ row }">
            {{ dateFormat(row.addTime) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.targetAmount')" align="center">
          <template #default="{ row }">
            {{ row.targetAmount }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.amountRaised')" align="center">
          <template #default="{ row }">
            {{ row.amountReceived }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.withdrawable')" align="center">
          <template #default="{ row }">
            {{ row.drawMoney }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.status')" align="center">
          <template #default="{ row }">
            <span :style="{ color: getStatusColor(row.status) }">
              {{ getStatusText(row.status) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.operating')" align="center" width="110">
          <template #default="{ row }">
            <el-button
              size="small"
              :disabled="row.drawMoney === 0"
              @click="handleWithdraw(row)"
              style="border: 1px solid #f0ac19; color: #f1ac19; border-radius: 10px;"
            >
              {{ $t('uc.myCf.withdraw') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- Tab 2: 其他众筹 -->
      <el-table v-show="activeTab === '2'" :data="qtList" :no-data-text="$t('common.nodata')" v-loading="loading" border>
        <el-table-column :label="$t('uc.myCf.title')" prop="fundingTitle" align="center" />
        <el-table-column :label="$t('uc.myCf.time')" align="center" width="160">
          <template #default="{ row }">
            {{ dateFormat(row.addTime) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.type')" align="center">
          <template #default="{ row }">
            {{ row.type === 1 ? '心愿' : '创业' }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.targetAmount')" align="center">
          <template #default="{ row }">
            {{ row.targetAmount }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.amountRaised')" align="center">
          <template #default="{ row }">
            {{ row.amountReceived }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.withdrawable')" align="center">
          <template #default="{ row }">
            {{ row.drawMoney }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.status')" align="center">
          <template #default="{ row }">
            <span :style="{ color: getStatusColor(row.status) }">
              {{ getStatusText(row.status) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.operating')" align="center" width="110">
          <template #default="{ row }">
            <el-button
              size="small"
              :disabled="row.drawMoney === 0"
              @click="handleWithdraw(row)"
              style="border: 1px solid #f0ac19; color: #f1ac19; border-radius: 10px;"
            >
              {{ $t('uc.myCf.withdraw') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- Tab 3: 线下公益 -->
      <el-table v-show="activeTab === '3'" :data="gyList" :no-data-text="$t('common.nodata')" v-loading="loading" border>
        <el-table-column :label="$t('uc.myCf.title')" prop="fundingTitle" align="center" />
        <el-table-column :label="$t('uc.myCf.time')" align="center" width="160">
          <template #default="{ row }">
            {{ dateFormat(row.addTime) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.type')" align="center">
          <template #default="{ row }">
            <span :style="{ color: row.type === 1 ? '#F65F47' : '#0071FE' }">
              {{ row.type === 1 ? $t('uc.myCf.wish') : $t('uc.myCf.pioneer') }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.applyAmount')" align="center">
          <template #default="{ row }">
            {{ row.applyAmount }} USDT
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.status')" align="center">
          <template #default="{ row }">
            <span :style="{ color: getStatusColor(row.status) }">
              {{ getStatusText(row.status) }}
            </span>
          </template>
        </el-table-column>
      </el-table>

      <!-- Tab 4: 我的捐款 -->
      <el-table v-show="activeTab === '4'" :data="jkList" :no-data-text="$t('common.nodata')" v-loading="loading" border>
        <el-table-column :label="$t('uc.myCf.title')" prop="fundingTitle" align="center" />
        <el-table-column :label="$t('uc.myCf.time')" align="center" width="160">
          <template #default="{ row }">
            {{ dateFormat(row.fundingTime) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.type')" align="center">
          <template #default="{ row }">
            <span :style="{ color: getFundingTypeColor(row.fundingType) }">
              {{ getFundingTypeText(row.fundingType) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.myCf.donateMoney')" align="center">
          <template #default="{ row }">
            {{ row.fundingMoney }} USDT
          </template>
        </el-table-column>
      </el-table>

      <div class="page">
        <el-pagination
          v-show="activeTab === '1'"
          layout="prev, pager, next"
          :total="total"
          :page-size="pageSize"
          :current-page="pageNo"
          @current-change="handlePageChange"
        />
        <el-pagination
          v-show="activeTab === '2'"
          layout="prev, pager, next"
          :total="total2"
          :page-size="pageSize2"
          :current-page="pageNo2"
          @current-change="handlePageChange2"
        />
        <el-pagination
          v-show="activeTab === '3'"
          layout="prev, pager, next"
          :total="total3"
          :page-size="pageSize3"
          :current-page="pageNo3"
          @current-change="handlePageChange3"
        />
        <el-pagination
          v-show="activeTab === '4'"
          layout="prev, pager, next"
          :total="total4"
          :page-size="pageSize4"
          :current-page="pageNo4"
          @current-change="handlePageChange4"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 众筹列表页面
 */
import { ref, onMounted } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElTabs, ElTabPane, ElPagination, ElButton, ElIcon } from 'element-plus'
import axios from 'axios'
import dayjs from 'dayjs'

const host = 'http://localhost'
const api = {
  crowdfunding: {
    myCrowdfundingYl: '/uc/crowdfunding/myYl',
    myCrowdfundingQt: '/uc/crowdfunding/myQt',
    myCrowdfundingGy: '/uc/crowdfunding/myGy',
    myCrowdfundingJk: '/uc/crowdfunding/myJk',
    getCash: '/uc/crowdfunding/getCash'
  }
}

const loading = ref(false)
const activeTab = ref('1')

// Page 1 - 医疗众筹
const pageSize = ref(10)
const pageNo = ref(1)
const total = ref(0)
const ylList = ref([])

// Page 2 - 其他众筹
const pageSize2 = ref(10)
const pageNo2 = ref(1)
const total2 = ref(0)
const qtList = ref([])

// Page 3 - 线下公益
const pageSize3 = ref(10)
const pageNo3 = ref(1)
const total3 = ref(0)
const gyList = ref([])

// Page 4 - 我的捐款
const pageSize4 = ref(10)
const pageNo4 = ref(1)
const total4 = ref(0)
const jkList = ref([])

const dateFormat = (tick) => {
  return dayjs(tick).format('YYYY-MM-DD')
}

const getStatusText = (status) => {
  const statusMap = {
    0: '审核中',
    1: '已通过',
    2: '已拒绝',
    3: '已结束'
  }
  return statusMap[status] || '未知'
}

const getStatusColor = (status) => {
  const colorMap = {
    0: '#F3BB01',
    1: '#05DF3F',
    2: '#E80000',
    3: '#A7A5A1'
  }
  return colorMap[status] || '#A7A5A1'
}

const getFundingTypeText = (type) => {
  const typeMap = {
    1: '心愿',
    2: '先锋',
    3: '医疗'
  }
  return typeMap[type] || '未知'
}

const getFundingTypeColor = (type) => {
  const colorMap = {
    1: '#F65F47',
    2: '#0071FE',
    3: '#0AC159'
  }
  return colorMap[type] || '#A7A5A1'
}

const handleTabClick = (tab) => {
  activeTab.value = tab.props.name
  pageNo.value = 1
  pageNo2.value = 1
  pageNo3.value = 1
  pageNo4.value = 1
  refresh()
  refresh2()
  refresh3()
  refresh4()
}

const handlePageChange = (page) => {
  pageNo.value = page
  refresh()
}

const handlePageChange2 = (page) => {
  pageNo2.value = page
  refresh2()
}

const handlePageChange3 = (page) => {
  pageNo3.value = page
  refresh3()
}

const handlePageChange4 = (page) => {
  pageNo4.value = page
  refresh4()
}

const handleWithdraw = (row) => {
  const param = {
    fundingId: row.id,
    type: row.type
  }

  axios.post(`${host}${api.crowdfunding.getCash}`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success(resp.message)
      refresh()
      refresh2()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const refresh = () => {
  loading.value = true
  const params = {
    pageNo: pageNo.value - 1,
    pageSize: pageSize.value
  }

  axios.post(`${host}${api.crowdfunding.myCrowdfundingYl}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    loading.value = false
    const resp = response.data
    if (resp.content && resp.content.length > 0) {
      total.value = resp.totalElements
      ylList.value = resp.content
    }
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const refresh2 = () => {
  loading.value = true
  const params = {
    pageNo: pageNo2.value - 1,
    pageSize: pageSize2.value
  }

  axios.post(`${host}${api.crowdfunding.myCrowdfundingQt}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    loading.value = false
    const resp = response.data
    if (resp.content && resp.content.length > 0) {
      total2.value = resp.totalElements
      qtList.value = resp.content
    }
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const refresh3 = () => {
  loading.value = true
  const params = {
    pageNo: pageNo3.value - 1,
    pageSize: pageSize3.value
  }

  axios.post(`${host}${api.crowdfunding.myCrowdfundingGy}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    loading.value = false
    const resp = response.data
    if (resp.content && resp.content.length > 0) {
      total3.value = resp.totalElements
      gyList.value = resp.content
    }
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const refresh4 = () => {
  loading.value = true
  const params = {
    pageNo: pageNo4.value - 1,
    pageSize: pageSize4.value
  }

  axios.post(`${host}${api.crowdfunding.myCrowdfundingJk}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    loading.value = false
    const resp = response.data
    if (resp.content && resp.content.length > 0) {
      total4.value = resp.totalElements
      jkList.value = resp.content
    }
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

onMounted(() => {
  refresh()
  refresh2()
  refresh3()
  refresh4()
})
</script>

<style scoped lang="scss">
.entrustcurrent {
  float: left;
  width: 100%;
  padding-left: 30px;
  background: transparent;
}

.page {
  text-align: right;
  margin-top: 20px;
}

.table {
  border-radius: 4px;
  background: transparent;

  :deep(.el-table) {
    background: transparent;
    border-radius: 4px;

    th.el-table__cell {
      background: #27313e;
      color: #fff;
    }

    td.el-table__cell {
      background: transparent;
      color: #fff;
    }
  }

  :deep(.el-tabs) {
    background: transparent;

    .el-tabs__header {
      background: #192330;
      padding: 0 16px;
      border-radius: 5px;

      .el-tabs__item {
        color: #fff;
        border: 1px solid #fff;
        padding: 3px 5px;
        margin-right: 10px;
        border-radius: 5px;

        &.is-active {
          color: #000;
          background: #f0ac19;
          border-color: #f0ac19;
        }

        &:hover {
          color: #f0a70a;
        }
      }
    }
  }
}

.order-type-btn-wrap {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 30px;
  padding: 0 16px;
  min-width: 1320px;
  background: #192330;

  button {
    width: 80px;
    border-radius: 5px;
    height: auto;
    background: #192330;
    border: 1px solid #fff;
    padding: 3px 5px;
    text-align: center;
    font-size: 12px;
    margin-right: 10px;
    color: #fff;
  }

  .btn-selected {
    color: #000 !important;
    background: #f0ac19;
    border: 1px solid #f0ac19 !important;
  }
}

.buy {
  color: #00b275 !important;
}

.sell {
  color: #f15057 !important;
}
</style>

<style lang="scss">
.entrustcurrent {
  .el-table th.el-table__cell,
  .entrustcurrent .el-table td.el-table__cell {
    text-align: center;
  }

  .el-tabs__nav-wrap::after {
    display: none;
  }

  .el-tabs__content {
    padding: 0;
  }
}
</style>

<template>
  <div class="entrusthistory">
    <div class="form-inline">
      <div class="form-item">
        <label class="form-label">{{ $t('uc.finance.trade.start_end') }}</label>
        <el-date-picker
          v-model="formItem.date"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          style="width: 180px;"
        />
      </div>
      <div class="form-item">
        <label class="form-label">{{ $t('swap.name') }}</label>
        <el-select v-model="formItem.contractCoinId" placeholder="请选择" style="width: 100px;">
          <el-option v-for="(item, index) in symbol" :key="index" :label="item.name" :value="item.id" />
        </el-select>
      </div>
      <div class="form-item">
        <label class="form-label">{{ $t('uc.finance.trade.type') }}</label>
        <el-select v-model="formItem.way" placeholder="请选择" style="width: 100px;">
          <el-option value="openLong" :label="$t('swap.openLong')" />
          <el-option value="openShort" :label="$t('swap.openShort')" />
          <el-option value="closeLong" :label="$t('swap.closeLong')" />
          <el-option value="closeShort" :label="$t('swap.closeShort')" />
        </el-select>
      </div>
      <div class="form-item">
        <el-button type="warning" @click="handleSubmit">{{ $t('uc.finance.trade.search') }}</el-button>
        <el-button style="margin-left: 8px;" @click="handleClear" class="clear_btn">{{ $t('uc.finance.trade.clear') }}</el-button>
      </div>
    </div>

    <div class="table">
      <el-tabs v-model="formItem.type" @tab-click="setSelectedType">
        <el-tab-pane :label="$t('swap.limited_price')" :name="2" />
        <el-tab-pane :label="$t('swap.trigger_price')" :name="3" />
      </el-tabs>

      <el-table :data="orders" v-loading="loading" border>
        <el-table-column label="#" width="200" align="center">
          <template #default="{ row }">
            <span :class="row.direction.toLowerCase()">
              {{ row.contractCoinName }}{{ row.contractCoinType === 'SECOND' ? '.' + row.holdTime + 'SEC' : '' }}.{{ row.leverage }}X.{{ $t('swap.' + getWay(row.direction, row.entrustType)) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.time')" width="180" align="center">
          <template #default="{ row }">
            {{ dateFormat(row.createTime) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.priceType')" align="center">
          <template #default="{ row }">
            {{ row.type === 'LIMIT_PRICE' ? '限价委托' : '计划委托' }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.price')" align="center">
          <template #default="{ row }">
            {{ toFloor(row.entrustPrice, baseCoinScale) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.num')" align="center">
          <template #default="{ row }">
            {{ row.share }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.strikePrice')" align="center">
          <template #default="{ row }">
            {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '--' }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.expand.fee')" align="center">
          <template #default="{ row }">
            <span class="sell">
              {{ row.fee === 0 ? '--' : toFloor(row.fee, baseCoinScale) + ' KICK' }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.profit')" align="center">
          <template #default="{ row }">
            <span :class="row.profit >= 0 ? 'buy' : 'sell'">
              {{ formatProfit(row) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.status')" align="center">
          <template #default="{ row }">
            <span :style="{ color: getStatusColor(row.status) }">
              {{ getStatusText(row.status) }}
            </span>
          </template>
        </el-table-column>
      </el-table>

      <div class="page">
        <el-pagination
          v-show="orders.length > 0"
          layout="prev, pager, next"
          :total="total"
          :page-size="pageSize"
          :current-page="pageNo"
          @current-change="loadDataPage"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 合约历史委托
 */
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElTabs, ElTabPane, ElButton, ElSelect, ElOption, ElDatePicker, ElPagination } from 'element-plus'
import axios from 'axios'
import dayjs from 'dayjs'
import { useStore } from 'vuex'

const store = useStore()

const host = ''
const api = {
  swap: {
    history: '/uc/contract/history',
    thumb: '/uc/contract/coin/thumb'
  }
}

const baseCoinScale = 4

const loading = ref(false)
const pageSize = ref(10)
const pageNo = ref(1)
const total = ref(10)
const symbol = ref([])
const orders = ref([])

const formItem = reactive({
  contractCoinId: '',
  type: 2,
  entrustType: '',
  direction: '',
  date: '',
  way: ''
})

const lang = computed(() => store.state.lang)
const memberId = computed(() => store.getters.member?.id || store.state.member?.id || 0)

const getWay = (direction, entrustType) => {
  if (direction === '1') {
    return entrustType === 'OPEN' ? 'openLong' : 'closeLong'
  } else {
    return entrustType === 'OPEN' ? 'openShort' : 'closeShort'
  }
}

const swapSplit = (val) => {
  const map = {
    openLong: { direction: '1', entrustType: 'OPEN' },
    openShort: { direction: '2', entrustType: 'OPEN' },
    closeLong: { direction: '1', entrustType: 'CLOSE' },
    closeShort: { direction: '2', entrustType: 'CLOSE' }
  }
  return map[val] || { direction: '', entrustType: '' }
}

const toFloor = (number, scale = 8) => {
  if (Number(number) === 0) return 0
  const str = number + ''
  if (str.indexOf('e') !== -1 || str.indexOf('E') !== -1) {
    return Number(number).toFixed(scale)
  }
  const dotIndex = str.indexOf('.')
  if (dotIndex === -1) return number
  const end = Math.min(dotIndex + scale + 1, str.length)
  return str.substring(0, end)
}

const dateFormat = (tick) => {
  return dayjs(tick).format('YYYY-MM-DD HH:mm:ss')
}

const formatProfit = (row) => {
  if (row.entrustType === 'OPEN') {
    return '--'
  }
  const profit = row.profit
  if (profit === undefined || profit === null) return '--'
  return (profit > 0 ? '+' : '') + toFloor(profit, baseCoinScale)
}

const getStatusText = (status) => {
  const statusMap = {
    ENTRUST_SUCCESS: '成功',
    ENTRUST_FAILURE: '失败',
    ENTRUST_CANCEL: '已撤销'
  }
  return statusMap[status] || '--'
}

const getStatusColor = (status) => {
  const colorMap = {
    ENTRUST_SUCCESS: '#f0a70a',
    ENTRUST_FAILURE: '#f15057',
    ENTRUST_CANCEL: '#7c7f82'
  }
  return colorMap[status] || '#A7A5A1'
}

const loadDataPage = (page) => {
  pageNo.value = page
  refresh()
}

const handleSubmit = () => {
  pageNo.value = 1
  refresh()
}

const handleClear = () => {
  formItem.contractCoinId = ''
  formItem.direction = ''
  formItem.date = ''
  formItem.entrustType = ''
  // Keep type
}

const refresh = () => {
  loading.value = true
  const { contractCoinId, type, direction, date: rangeDate, entrustType } = formItem
  const startTime = rangeDate && rangeDate[0] ? new Date(rangeDate[0]).getTime() : ''
  const endTime = rangeDate && rangeDate[1] ? new Date(rangeDate[1]).getTime() : ''

  const params = {}
  params.memberId = memberId.value
  if (contractCoinId) params.contractCoinId = contractCoinId
  if (direction) params.direction = direction
  if (type) params.type = type
  if (startTime) params.startTime = startTime
  if (endTime) params.endTime = endTime
  if (entrustType) params.entrustType = entrustType
  params.pageNo = pageNo.value - 1
  params.pageSize = pageSize.value

  orders.value = []

  axios.post(`${host}${api.swap.history}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.content && resp.content.length > 0) {
      total.value = resp.totalElements
      orders.value = resp.content
    }
    loading.value = false
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const getSymbol = () => {
  axios.post(`${host}${api.swap.thumb}`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp && resp.length > 0) {
      symbol.value = resp
    }
  }).catch(() => {
    console.error('获取币种列表失败')
  })
}

const setSelectedType = (tab) => {
  handleClear()
  formItem.type = tab.props.name
  pageNo.value = 1
  refresh()
}

watch(() => formItem.way, (val) => {
  if (val) {
    const re = swapSplit(val)
    formItem.direction = re.direction
    formItem.entrustType = re.entrustType
  }
})

watch(lang, () => {
  // 国际化更新
})

onMounted(() => {
  refresh()
  getSymbol()
})
</script>

<style scoped lang="scss">
.entrusthistory {
  float: left;
  width: 100%;
  padding-left: 30px;
}

.page {
  text-align: right;
  margin-top: 20px;
}

.form-inline {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 15px;
  margin-bottom: 20px;

  .form-item {
    display: flex;
    align-items: center;
    gap: 8px;

    .form-label {
      font-size: 14px;
      color: #fff;
      white-space: nowrap;
    }
  }
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

        &.is-active {
          color: #f0a70a;
        }

        &:hover {
          color: #f0a70a;
        }
      }
    }

    .el-tabs__nav-wrap::after {
      display: none;
    }
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
.entrusthistory {
  .el-table th.el-table__cell,
  .entrusthistory .el-table td.el-table__cell {
    text-align: center;
  }

  .el-tabs__content {
    padding: 0;
  }
}
</style>

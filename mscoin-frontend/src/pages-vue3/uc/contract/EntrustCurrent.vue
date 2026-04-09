<template>
  <div class="entrustcurrent">
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

      <el-table v-show="formItem.type === 2" :data="orders" v-loading="loading" border>
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
            {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '-' }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.frozenAmount')" align="center">
          <template #default="{ row }">
            {{ toFloor(row.principalAmount, baseCoinScale) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.action')" width="110" align="center">
          <template #default="{ row, $index }">
            <el-button
              size="small"
              @click="handleCancel($index, row)"
              style="border: 1px solid #f0ac19; color: #f1ac19; border-radius: 10px;"
            >
              {{ $t('swap.undo') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-table v-show="formItem.type === 3" :data="orders" v-loading="loading" border>
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
        <el-table-column :label="$t('swap.triggerPrice')" align="center">
          <template #default="{ row }">
            {{ row.triggerPrice > 0 ? toFloor(row.triggerPrice, baseCoinScale) : '--' }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.strikePrice')" align="center">
          <template #default="{ row }">
            {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '-' }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.frozenAmount')" align="center">
          <template #default="{ row }">
            {{ toFloor(row.principalAmount, baseCoinScale) }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('swap.action')" width="110" align="center">
          <template #default="{ row, $index }">
            <el-button
              size="small"
              @click="handleCancel($index, row)"
              style="border: 1px solid #f0ac19; color: #f1ac19; border-radius: 10px;"
            >
              {{ $t('swap.undo') }}
            </el-button>
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
 * Vue 3 迁移 - 合约当前委托
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
    current: '/uc/contract/current',
    thumb: '/uc/contract/coin/thumb',
    entrustCancel: '/uc/contract/entrust/cancel'
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

const getWay = (direction, entrustType) => {
  if (direction === '1') {
    return entrustType === 'OPEN' ? 'openLong' : 'closeLong'
  } else {
    return entrustType === 'OPEN' ? 'openShort' : 'closeShort'
  }
}

const swapMerge = (direction, entrustType) => {
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
  if (contractCoinId) params.contractCoinId = contractCoinId
  if (direction) params.direction = direction
  if (type) params.type = type
  if (startTime) params.startTime = startTime
  if (endTime) params.endTime = endTime
  if (entrustType) params.entrustType = entrustType
  params.pageNo = pageNo.value - 1
  params.pageSize = pageSize.value

  orders.value = []

  axios.post(`${host}${api.swap.current}`, params, {
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

const handleCancel = (index, order) => {
  if (!confirm('确定要撤销该委托吗？')) return

  axios.post(`${host}${api.swap.entrustCancel}/${order.id}`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success('撤销成功')
      refresh()
    } else {
      ElMessage.error(resp.message || '撤销失败')
    }
  }).catch(() => {
    ElMessage.error('请求失败')
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
.entrustcurrent {
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
.entrustcurrent {
  .el-table th.el-table__cell,
  .entrustcurrent .el-table td.el-table__cell {
    text-align: center;
  }

  .el-tabs__content {
    padding: 0;
  }
}
</style>

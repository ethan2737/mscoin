<template>
  <div class="entrustcurrent">
    <el-form :model="formItem" label-width="75px" inline>
      <el-form-item :label="$t('uc.finance.trade.start_end')">
        <el-date-picker
          v-model="formItem.date"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          style="width: 180px;"
        />
      </el-form-item>
      <el-form-item :label="$t('uc.finance.trade.symbol')">
        <el-select v-model="formItem.symbol" :placeholder="$t('common.pleaseselect')" style="width: 100px;">
          <el-option v-for="item in symbol" :key="item.symbol" :value="item.symbol" />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('uc.finance.trade.type')">
        <el-select v-model="formItem.type" :placeholder="$t('common.pleaseselect')" style="width: 70px;">
          <el-option value="LIMIT_PRICE" :label="$t('uc.finance.trade.limit')" />
          <el-option value="MARKET_PRICE" :label="$t('uc.finance.trade.market')" />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('uc.finance.trade.direction')">
        <el-select v-model="formItem.direction" :placeholder="$t('common.pleaseselect')" style="width: 70px;">
          <el-option value="BUY" :label="$t('uc.finance.trade.buy')" />
          <el-option value="SELL" :label="$t('uc.finance.trade.sell')" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="handleSubmit">{{ $t('uc.finance.trade.search') }}</el-button>
        <el-button style="margin-left: 8px" @click="handleClear" class="clear_btn">{{ $t('uc.finance.trade.clear') }}</el-button>
      </el-form-item>
    </el-form>
    <div class="table">
      <el-table
        :data="orders"
        v-loading="loading"
        element-loading-text="加载中..."
        border
        style="width: 100%"
      >
        <el-table-column type="expand">
          <template #default="props">
            <expand :skin="props.row.skin" :rows="props.row.detail" />
          </template>
        </el-table-column>
        <el-table-column prop="time" :label="$t('exchange.time')" min-width="55">
          <template #default="{ row }">
            <span>{{ dateFormat(row.time) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="symbol" :label="$t('uc.finance.trade.symbol')" width="130" />
        <el-table-column :label="$t('uc.finance.trade.type')" width="70">
          <template #default="{ row }">
            <span>{{ row.type === 'LIMIT_PRICE' ? $t('exchange.limited_price') : $t('exchange.market_price') }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('exchange.direction')" width="90">
          <template #default="{ row }">
            <span :class="row.direction.toLowerCase()">
              {{ row.direction === 'BUY' ? $t('exchange.buyin') : $t('exchange.sellout') }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('exchange.price')">
          <template #default="{ row }">
            <span :title="row.price">{{ toFloor(row.price) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('exchange.num')">
          <template #default="{ row }">
            <span :title="row.amount">{{ toFloor(row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('exchange.traded')">
          <template #default="{ row }">
            <span :title="row.tradedAmount">{{ toFloor(row.tradedAmount) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.finance.trade.turnover')">
          <template #default="{ row }">
            <span :title="row.turnover">{{ toFloor(row.turnover) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('exchange.action')" width="110">
          <template #default="{ row }">
            <el-button
              size="small"
              type="warning"
              style="border-radius: 10px"
              @click="handleCancel(row.orderId)"
            >
              {{ $t('exchange.undo') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="page">
        <el-pagination
          layout="total, prev, pager, next"
          :total="total"
          :page-size="pageSize"
          :current-page="pageNo"
          @current-change="handlePageChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, inject, onMounted, reactive, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import moment from 'moment'
import axios from 'axios'

import expand from './expand.vue'
import { getAuthHeaders, useRuntimeContract } from '../../config/runtime-vue3'

const store = inject('store')
const { api } = useRuntimeContract()
const host = ''

const loading = ref(false)
const pageSize = ref(10)
const pageNo = ref(1)
const total = ref(0)
const symbol = ref([])
const orders = ref([])

const formItem = reactive({
  symbol: '',
  type: '',
  direction: '',
  date: ''
})

const lang = computed(() => store?.state?.lang)

const dateFormat = (tick) => moment(tick).format('YYYY-MM-DD HH:mm:ss')

const toFloor = (number, scale = 8) => {
  if (Number(number) === 0) {
    return 0
  }
  const raw = `${number}`
  if (raw.includes('e') || raw.includes('E')) {
    const fixed = Number(number).toFixed(scale + 1)
    return fixed.substring(0, fixed.length - 1)
  }
  if (!raw.includes('.')) {
    return raw
  }
  if (scale === 0) {
    return raw.substring(0, raw.indexOf('.'))
  }
  return raw.substring(0, raw.indexOf('.') + scale + 1)
}

const loadOrders = () => {
  loading.value = true
  const { symbol: symbolVal, type, direction, date: rangeDate } = formItem
  const params = {
    pageNo: pageNo.value,
    pageSize: pageSize.value
  }
  if (symbolVal) params.symbol = symbolVal
  if (direction) params.direction = direction
  if (type) params.type = type
  if (rangeDate?.[0]) params.startTime = new Date(rangeDate[0]).getTime()
  if (rangeDate?.[1]) params.endTime = new Date(rangeDate[1]).getTime()

  axios.post(`${host}${api.exchange.currentOrder}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      ...getAuthHeaders()
    }
  })
    .then((response) => {
      const resp = response.data
      if (resp.code !== 0) {
        ElMessage.error(resp.message || '加载当前委托失败')
        return
      }
      const page = resp.data || {}
      total.value = page.totalElements || 0
      orders.value = (page.content || []).map((row) => ({
        ...row,
        price: row.type === 'MARKET_PRICE' ? '市价' : row.price
      }))
    })
    .catch(() => {
      ElMessage.error('加载当前委托失败')
    })
    .finally(() => {
      loading.value = false
    })
}

const loadSymbol = () => {
  axios.post(`${host}${api.market.thumb}`, {}, {
    withCredentials: true
  })
    .then((response) => {
      const resp = response.data
      if (resp.code === 0 && Array.isArray(resp.data)) {
        symbol.value = resp.data
        if (!formItem.symbol && resp.data.length > 0) {
          const preferred = resp.data.find((item) => item.symbol === 'BTC/USDT')
          formItem.symbol = (preferred || resp.data[0]).symbol
          loadOrders()
        }
      }
    })
}

const handlePageChange = (page) => {
  pageNo.value = page
  loadOrders()
}

const handleSubmit = () => {
  pageNo.value = 1
  loadOrders()
}

const handleClear = () => {
  formItem.symbol = ''
  formItem.type = ''
  formItem.direction = ''
  formItem.date = ''
  pageNo.value = 1
  loadOrders()
}

const handleCancel = (orderId) => {
  ElMessageBox.confirm('确定要撤销此订单吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(() => {
      axios.post(`${host}${api.exchange.orderCancel}/${orderId}`, {}, {
        withCredentials: true,
        headers: getAuthHeaders()
      })
        .then((response) => {
          const resp = response.data
          if (resp.code === 0) {
            ElMessage.success('撤单成功')
            loadOrders()
            return
          }
          ElMessage.error(resp.message || '撤单失败')
        })
        .catch(() => {
          ElMessage.error('撤单失败')
        })
    })
    .catch(() => {})
}

onMounted(() => {
  loadOrders()
  loadSymbol()
})

watch(lang, () => {
  loadOrders()
})
</script>

<style lang="scss" scoped>
.entrustcurrent {
  float: left;
  width: 100%;
  padding-left: 30px;
}

.page {
  text-align: right;
  margin-top: 20px;
}

.table {
  border-radius: 4px;
}

.table :deep(.el-table) {
  border-radius: 4px;
}

.table :deep(.el-table__expanded-cell) {
  padding: 10px 50px;
}

.form :deep(.el-form-item) {
  display: inline-block;
  margin-right: 16px;
}

.buy {
  color: #00b050;
}

.sell {
  color: #f0a70a;
}
</style>

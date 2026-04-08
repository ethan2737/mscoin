<template>
  <div class="entrusthistory">
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
          <el-option value="0" :label="$t('uc.finance.trade.buy')" />
          <el-option value="1" :label="$t('uc.finance.trade.sell')" />
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
        <el-table-column :label="$t('exchange.done')">
          <template #default="{ row }">
            <span :title="row.tradedAmount">{{ toFloor(row.tradedAmount) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('uc.finance.trade.turnover')">
          <template #default="{ row }">
            <span :title="row.turnover">{{ toFloor(row.turnover) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('exchange.status')" width="80">
          <template #default="{ row }">
            <span v-if="row.status === 'COMPLETED'" style="color: #f0a70a">{{ $t('exchange.finished') }}</span>
            <span v-else-if="row.status === 'CANCELED'" style="color: #f0a70a">{{ $t('exchange.canceled') }}</span>
            <span v-else>--</span>
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
/**
 * Vue 3 迁移 - 历史委托组件
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. iView Table 替换为 Element Plus Table
 * 4. h() render 函数改为 slot 语法
 * 5. 使用 inject 获取 store 和 router
 */
import { ref, reactive, computed, inject, onMounted, watch } from 'vue'
import expand from './expand.vue'
import moment from 'moment'
import axios from 'axios'

// 注入 store 和 router (Vue 3 兼容模式)
const store = inject('store')
const router = inject('router')

// 状态
const loading = ref(false)
const pageSize = ref(10)
const pageNo = ref(1)
const total = ref(10)
const symbol = ref([])
const orders = ref([])

// 表单数据
const formItem = reactive({
  symbol: '',
  type: '',
  direction: '',
  date: ''
})

// 监听 lang 变化
const lang = computed(() => store?.state?.lang)

// 方法
const dateFormat = (tick) => {
  return moment(tick).format('YYYY-MM-DD HH:mm:ss')
}

const toFloor = (number, scale = 8) => {
  if (new Number(number) == 0) {
    return 0
  }
  const __str = number + ''
  if (__str.indexOf('e') > -1 || __str.indexOf('E') > -1) {
    const __num = new Number(number).toFixed(scale + 1)
    const __str2 = __num + ''
    return __str2.substring(0, __str2.length - 1)
  } else if (__str.indexOf('.') > -1) {
    if (scale == 0) {
      return __str.substring(0, __str.indexOf('.'))
    }
    return __str.substring(0, __str.indexOf('.') + scale + 1)
  } else {
    return __str
  }
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
}

const loadOrders = () => {
  loading.value = true
  const { symbol: symbolVal, type, direction, date: rangeDate } = formItem
  const startTime = rangeDate && rangeDate[0] ? new Date(rangeDate[0]).getTime() : ''
  const endTime = rangeDate && rangeDate[1] ? new Date(rangeDate[1]).getTime() : ''

  let params = {}
  if (symbolVal) params.symbol = symbolVal
  if (direction) params.direction = direction
  if (type) params.type = type
  if (startTime) params.startTime = startTime
  if (endTime) params.endTime = endTime
  params.pageNo = pageNo.value
  params.pageSize = pageSize.value

  orders.value = []

  const host = 'http://localhost'
  axios.post(`${host}/exchange/order/personal/history`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    let rows = []
    if (resp.content && resp.content.length > 0) {
      total.value = resp.totalElements
      for (let i = 0; i < resp.content.length; i++) {
        let row = resp.content[i]
        row.price = row.type === 'MARKET_PRICE' ? '市价' : row.price
        rows.push(row)
      }
      orders.value = rows
    }
    loading.value = false
  })
  .catch(() => {
    loading.value = false
  })
}

const loadSymbol = () => {
  const host = 'http://localhost'
  axios.post(`${host}/api/market/thumb`, {}, {
    withCredentials: true
  })
  .then(response => {
    const resp = response.data
    if (resp && resp.length > 0) {
      symbol.value = resp
    }
  })
}

// 生命周期
onMounted(() => {
  loadOrders()
  loadSymbol()
})

// 监听语言变化
watch(lang, () => {
  // 语言变化时重新加载数据
  loadOrders()
})
</script>

<style lang="scss" scoped>
.entrusthistory {
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

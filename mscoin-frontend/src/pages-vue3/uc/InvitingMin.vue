<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear record">
        <div class="col-xs-12 rightarea-con">
          <div class="trade_accumulative">
            <div class="trade_accumulative_return">
              {{ $t('uc.finance.trade.accumulative_return') }}&nbsp;&nbsp;{{ Number(accumulative_return).toFixed(8) }}
            </div>
            <div class="trade_accumulat_return">
              {{ $t('uc.finance.trade.accumulat_return') }}&nbsp;&nbsp;{{ Number(accumulat_return).toFixed(8) }}
            </div>
          </div>
          <div class="form-group">
            <span>{{ $t('uc.finance.trade.start_end') }}：</span>
            <el-date-picker
              v-model="rangeDate"
              type="daterange"
              range-separator="-"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              style="width: 200px; margin-right: 35px;"
              value-format="timestamp"
              @change="changedate"
            />
            <el-button type="warning" @click="queryOrder" style="padding: 6px 30px; margin-left: 10px;">
              {{ $t('uc.finance.trade.search') }}
            </el-button>
          </div>
          <div class="order-table">
            <el-table :data="tableRecord" :no-data-text="$t('common.nodata')" v-loading="loading" border>
              <el-table-column prop="transaction_time" :label="$t('uc.finance.trade.transactionTime')" align="center">
                <template #default="{ row }">
                  {{ row._source?.transaction_time || '' }}
                </template>
              </el-table-column>
              <el-table-column prop="exchange_order_id" :label="$t('uc.finance.trade.exchangeOrderId')" align="center">
                <template #default="{ row }">
                  {{ row._source?.exchange_order_id || '' }}
                </template>
              </el-table-column>
              <el-table-column prop="symbol" :label="$t('uc.finance.trade.symbol')" align="center">
                <template #default="{ row }">
                  {{ row._source?.symbol || '' }}
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.trade.direction')" align="center">
                <template #default="{ row }">
                  {{ formatDirection(row._source) }}
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.trade.poundageAmount')" align="center">
                <template #default="{ row }">
                  <span :title="row._source?.poundage_amount">
                    {{ toFloor(row._source?.poundage_amount || 0) }} {{ row._source?.coin_id || '' }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.trade.mineAmount')" align="center">
                <template #default="{ row }">
                  <span :title="row._source?.mine_amount">
                    {{ toFloor(row._source?.mine_amount || 0) }}
                  </span>
                </template>
              </el-table-column>
            </el-table>
            <div style="margin: 10px; overflow: hidden">
              <div style="float: right;">
                <el-pagination
                  layout="total, prev, pager, next"
                  :total="total"
                  :page-size="pageSize"
                  :current-page="page"
                  @current-change="changePage"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 邀请挖矿记录页面
 */
import { ref, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const store = useStore()
const router = useRouter()

const host = ''
const api = {
  uc: {
    mylist: '/uc/asset/transaction/list'
  }
}

const loading = ref(false)
const rangeDate = ref('')
const startTime = ref('')
const endTime = ref('')
const accumulative_return = ref(0)
const accumulat_return = ref(0)
const pageSize = ref(10)
const page = ref(1)
const total = ref(0)
const tableRecord = ref([])

const toFloor = (number, scale = 8) => {
  if (Number(number) === 0) return 0
  const str = number + ''
  if (str.indexOf('e') > -1 || str.indexOf('E') > -1) {
    const num = Number(number).toFixed(scale + 1)
    return (num + '').substring(0, (num + '').length - 1)
  } else if (str.indexOf('.') > -1) {
    if (scale === 0) return str.substring(0, str.indexOf('.'))
    return str.substring(0, str.indexOf('.') + scale + 1)
  }
  return str
}

const formatDirection = (source) => {
  if (!source) return ''
  const type = source.type
  const direction = source.direction
  const lang = store.state.lang
  let str1 = ''
  let str2 = ''

  if (type === 1) {
    str1 = lang === 'English' ? 'market price' : '限价'
  } else if (type === 0) {
    str1 = lang === 'English' ? 'Present price' : '市价'
  }

  if (direction === 1) {
    str2 = lang === 'English' ? ' sell' : '卖出'
  } else if (direction === 0) {
    str2 = lang === 'English' ? ' buy' : '买入'
  }

  return str1 + str2
}

const changedate = () => {
  if (rangeDate.value && rangeDate.value[0]) {
    startTime.value = dateform(rangeDate.value[0])
    endTime.value = dateform(rangeDate.value[1])
  }
}

const dateform = (time) => {
  const date = new Date(time)
  const y = date.getFullYear()
  let m = date.getMonth() + 1
  m = m < 10 ? '0' + m : m
  let d = date.getDate()
  d = d < 10 ? '0' + d : d
  let h = date.getHours()
  h = h < 10 ? '0' + h : h
  let minute = date.getMinutes()
  let second = date.getSeconds()
  minute = minute < 10 ? '0' + minute : minute
  second = second < 10 ? '0' + second : second
  return y + '-' + m + '-' + d + ' ' + h + ':' + minute + ':' + second
}

const init = () => {
  let memberId = 0
  if (!store.getters.isLogin) {
    router.push('/login')
    return
  }
  memberId = store.getters.member.id

  let start = ''
  let end = ''
  if (startTime.value) start = startTime.value
  if (endTime.value) end = endTime.value

  const params = {
    memberId,
    page: page.value,
    limit: 10,
    startTime: start,
    endTime: end
  }

  loading.value = true
  axios.post(`${host}${api.uc.mylist}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      const data = resp.data
      accumulative_return.value = data.backAmount
      accumulat_return.value = data.preAmount
      if (data.data) {
        const trueData = data.data.hits
        total.value = trueData.total
        tableRecord.value = trueData.hits
      }
    } else {
      ElMessage.error(resp.message)
    }
    loading.value = false
  }).catch(() => {
    ElMessage.error('请求失败')
    loading.value = false
  })
}

const changePage = (pageIndex) => {
  page.value = pageIndex
  init()
}

const queryOrder = () => {
  if (!rangeDate.value || rangeDate.value.length === 0) {
    ElMessage.error('请选择搜索日期范围')
    return
  }

  try {
    page.value = 1
    init()
  } catch (ex) {
    ElMessage.error('请选择搜索日期范围')
  }
}

onMounted(() => {
  init()
})
</script>

<style scoped lang="scss">
.datedaitl {
  text-align: left;
  margin-bottom: 10px;
}

.bill_box {
  width: 100%;
  height: auto;
  overflow: hidden;
}

.form-group {
  margin-bottom: 20px;
  text-align: left;
}

.rightarea {
  background: #fff;
  padding-left: 15px !important;
  padding-right: 15px !important;
  margin-bottom: 60px !important;

  .rightarea-top {
    line-height: 75px;
    border-bottom: #f1f1f1 solid 1px;
  }

  .rightarea-con {
    padding-top: 30px;
    padding-bottom: 20px;
  }
}

.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 15px;
}

.order-table {
  margin-top: 20px;

  :deep(.el-table__cell) {
    padding-right: 0 !important;
  }
}

.trade_accumulative {
  height: auto;
  overflow: hidden;
  font-size: 18px;
  font-weight: 600;
  text-align: left;
  border-bottom: 1px solid #f5f5f5;
  padding-bottom: 20px;
  margin-bottom: 20px;

  .trade_accumulative_return {
    width: 50%;
    float: left;
  }

  .trade_accumulat_return {
    width: 50%;
    float: left;
  }
}
</style>

<style lang="scss">
th .el-table__cell {
  overflow: hidden;
  white-space: nowrap;
}
</style>

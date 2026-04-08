<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear record">
        <div class="col-xs-12 rightarea-con">
          <div class="trade_accumulative">
            <div class="trade_accumulative_return">
              {{ $t('uc.finance.inviting.accumulative_return') }}&nbsp;&nbsp;{{ accumulative_return }}
            </div>
            <div class="trade_accumulat_return">
              {{ $t('uc.finance.inviting.accumulat_return') }}&nbsp;&nbsp;{{ accumulat_return }}
            </div>
          </div>
          <div class="form-group">
            <span>{{ $t('uc.finance.inviting.start_end') }}：</span>
            <el-date-picker
              v-model="rangeDate"
              type="daterange"
              range-separator="-"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              style="width: 200px; margin-right: 35px;"
            />
            <el-button type="warning" @click="queryOrder" style="padding: 6px 50px; margin-left: 10px;">
              {{ $t('uc.finance.record.search') }}
            </el-button>
          </div>
          <div class="datedaitl">
            <span style="color: #eb6f6c">{{ $t('uc.finance.inviting.start_end') }}：</span>&nbsp;&nbsp;
            <span>{{ $t('uc.finance.inviting.chargetime') }}</span>
          </div>
          <div class="order-table">
            <el-table :data="tableRecord" :no-data-text="$t('common.nodata')" v-loading="loading" border>
              <el-table-column prop="transactionTime" :label="$t('uc.finance.trade.transactionTime')" align="center" />
              <el-table-column prop="inviterName" :label="$t('uc.finance.inviting.refereename')" align="center" />
              <el-table-column prop="inviterMobile" :label="$t('uc.finance.inviting.referetel')" align="center" />
              <el-table-column prop="mineAmount" :label="$t('uc.finance.inviting.invitingawards')" align="center" />
              <el-table-column :label="$t('uc.finance.trade.account_date')" align="center" width="160">
                <template #default="{ row }">
                  {{ formatAccountDate(row.transactionTime) }}
                </template>
              </el-table-column>
            </el-table>
            <div style="margin: 10px; overflow: hidden">
              <div style="float: right;">
                <el-pagination
                  layout="total, prev, pager, next"
                  :total="total"
                  :page-size="pageSize"
                  :current-page="page + 1"
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
 * Vue 3 迁移 - 挖矿交易记录页面
 */
import { ref, reactive, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
import { useStore, useRouter } from 'vue-router/composables'

const store = useStore()
const router = useRouter()

const host = 'http://localhost'
const api = {
  uc: {
    mylist: '/uc/asset/transaction/list'
  }
}

const loading = ref(false)
const rangeDate = ref('')
const startTime = ref('')
const endTime = ref('')
const accumulative_return = ref('1000')
const accumulat_return = ref('1000')
const pageSize = ref(10)
const page = ref(0)
const total = ref(0)
const tableRecord = ref([])

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

const formatAccountDate = (transactionTime) => {
  if (!transactionTime) return ''
  const reg = /-/g
  const a = transactionTime
  const b = a.split(' ')[0].replace(reg, '/')
  const c = a.split(' ')[1]
  const date = new Date(b).getTime() + 86400000
  const formatted = dateform(date)
  return formatted.split(' ')[0] + ' ' + c
}

const init = (pageNo) => {
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
    page: pageNo,
    limit: 10,
    inviterState: 1,
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
      total.value = data.exchangeOrders.totalElements
      tableRecord.value = data.exchangeOrders.content
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
  init(pageIndex)
}

const queryOrder = () => {
  if (!rangeDate.value || rangeDate.value.length === 0) {
    ElMessage.error('请选择搜索日期范围')
    return
  }

  try {
    let rangedate = ''
    rangedate += rangeDate.value[0].getFullYear() + '-' + (rangeDate.value[0].getMonth() + 1) + '-' + rangeDate.value[0].getDate()
    rangedate += '~'
    rangedate += rangeDate.value[1].getFullYear() + '-' + (rangeDate.value[1].getMonth() + 1) + '-' + rangeDate.value[1].getDate()

    const date = rangedate.split('~')
    startTime.value = new Date(date[0]).getTime()
    endTime.value = new Date(date[1]).getTime()
    init(1)
  } catch (ex) {
    ElMessage.error('请选择搜索日期范围')
  }
}

onMounted(() => {
  init(1)
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
  border-bottom: 1px solid #ccc;
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

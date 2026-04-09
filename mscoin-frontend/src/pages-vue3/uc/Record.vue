<template>
  <div class="nav-rights">
    <div class="nav-right">
      <div class="bill_flow_box">
        <div class="rightarea-con">
          <div class="form-group">
            <span>{{ $t('uc.finance.record.start_end') }}：</span>
            <el-date-picker
              v-model="rangeDate"
              type="daterange"
              range-separator="-"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              style="width: 200px;margin-right:30px;"
              @change="changedate"
            />
            <span>{{ $t('uc.finance.record.symbol') }}：</span>
            <el-select v-model="coinType" style="width:100px;margin-right:30px;" @change="getAddrList" clearable :placeholder="$t('common.pleaseselect')">
              <el-option v-for="item in coinList" :key="item.unit" :value="item.unit">{{ item.unit }}</el-option>
            </el-select>
            <span>{{ $t('uc.finance.record.operatetype') }}：</span>
            <el-select v-model="recordValue" clearable style="width:200px" @change="getType" :placeholder="$t('common.pleaseselect')">
              <el-option v-for="item in recordType" :key="item.value" :value="item.value">{{ item.label }}</el-option>
            </el-select>
            <el-button type="warning" @click="queryOrder" style="padding: 6px 30px;margin-left:10px;">{{ $t('uc.finance.record.search') }}</el-button>
          </div>
          <div class="order-table">
            <el-table :data="tableRecord" v-loading="loading" border style="width: 100%">
              <el-table-column prop="createTime" :label="$t('uc.finance.record.chargetime')" align="center" width="160" />
              <el-table-column :label="$t('uc.finance.record.type')" align="center">
                <template #default="{ row }">
                  {{ getRecordTypeName(row.type) }}
                </template>
              </el-table-column>
              <el-table-column prop="symbol" :label="$t('uc.finance.record.symbol')" align="center" />
              <el-table-column :label="$t('uc.finance.record.num')" align="center">
                <template #default="{ row }">
                  <span :title="row.amount">{{ toFloor(row.amount || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.record.shouldfee')" align="center">
                <template #default="{ row }">
                  <span :title="row.fee">{{ toFloor(row.fee || '0') }}{{ row.type === 3 ? ' KICK' : '' }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.record.discountfee')" align="center">
                <template #default="{ row }">
                  <span :title="row.discountFee">{{ toFloor(row.discountFee || '0') }}{{ row.type === 3 ? ' KICK' : '' }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.record.realfee')" align="center">
                <template #default="{ row }">
                  <span :title="row.realFee">{{ toFloor(row.realFee || '0') }}{{ row.type === 3 ? ' KICK' : '' }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.record.status')" align="center">
                <template #default>
                  {{ $t('uc.finance.record.finish') }}
                </template>
              </el-table-column>
            </el-table>
            <div style="margin: 10px;overflow: hidden">
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
 * Vue 3 迁移 - 账单明细页面
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. iView 组件替换为 Element Plus
 * 4. vue-resource 替换为 axios
 * 5. h() render 函数改为 slot 语法
 */
import { ref, reactive, computed, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

const store = inject('store')
const router = inject('router')

const host = ''

const loading = ref(false)
const rangeDate = ref('')
const startTime = ref('')
const endTime = ref('')
const recordValue = ref('')
const coinType = ref('')
const coinList = ref([])
const pageSize = ref(10)
const page = ref(1)
const total = ref(0)
const tableRecord = ref([])

const recordType = computed(() => [
  { value: 0, label: $t('uc.finance.record.charge') },
  { value: 1, label: $t('uc.finance.record.pickup') },
  { value: 2, label: $t('uc.finance.record.transaccount') },
  { value: 3, label: $t('uc.finance.record.exchange') },
  { value: 4, label: $t('uc.finance.record.otcbuy') },
  { value: 5, label: $t('uc.finance.record.otcsell') },
  { value: 6, label: $t('uc.finance.record.activityaward') },
  { value: 7, label: $t('uc.finance.record.promotionaward') },
  { value: 8, label: $t('uc.finance.record.dividend') },
  { value: 9, label: $t('uc.finance.record.vote') },
  { value: 10, label: $t('uc.finance.record.handrecharge') },
  { value: 11, label: $t('uc.finance.record.match') },
  { value: 12, label: $t('uc.finance.record.activitybuy') },
  { value: 13, label: $t('uc.finance.record.ctcbuy') },
  { value: 14, label: $t('uc.finance.record.ctcsell') },
  { value: 15, label: $t('uc.finance.record.redout') },
  { value: 16, label: $t('uc.finance.record.redin') },
  { value: 17, label: $t('uc.finance.record.miningfee') },
  { value: 18, label: $t('uc.finance.record.miningout') },
  { value: 19, label: $t('uc.finance.record.miningin') },
  { value: 20, label: $t('uc.finance.record.donatein') },
  { value: 21, label: $t('uc.finance.record.donateout') }
])

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

const getRecordTypeName = (type) => {
  const typeMap = {
    0: $t('uc.finance.record.charge'),
    1: $t('uc.finance.record.pickup'),
    2: $t('uc.finance.record.transaccount'),
    3: $t('uc.finance.record.exchange'),
    4: $t('uc.finance.record.otcbuy'),
    5: $t('uc.finance.record.otcsell'),
    6: $t('uc.finance.record.activityaward'),
    7: $t('uc.finance.record.promotionaward'),
    8: $t('uc.finance.record.dividend'),
    9: $t('uc.finance.record.vote'),
    10: $t('uc.finance.record.handrecharge'),
    11: $t('uc.finance.record.match'),
    12: $t('uc.finance.record.activitybuy'),
    13: $t('uc.finance.record.ctcbuy'),
    14: $t('uc.finance.record.ctcsell'),
    15: $t('uc.finance.record.redout'),
    16: $t('uc.finance.record.redin'),
    17: $t('uc.finance.record.miningfee'),
    18: $t('uc.finance.record.miningout'),
    19: $t('uc.finance.record.miningin'),
    20: $t('uc.finance.record.donatein'),
    21: $t('uc.finance.record.donateout')
  }
  return typeMap[type] || ''
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

const changePage = (pageindex) => {
  page.value = pageindex
  getList(pageindex)
}

const queryOrder = () => {
  if (!rangeDate.value || rangeDate.value.length === 0) {
    ElMessage.error('请选择搜索日期范围')
    return
  }
  page.value = 1
  getList(1)
}

const getType = () => {
  // 类型变化处理
}

const getAddrList = () => {
  axios.post(`${host}/uc/withdraw/support/coin/info`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0 && resp.data.length > 0) {
      coinList.value = resp.data
    } else {
      ElMessage.error(resp?.message || '获取币种列表失败')
    }
  })
  .catch(() => {
    ElMessage.error('获取币种列表失败')
  })
}

const getList = (pageNo) => {
  const memberId = store?.state?.member?.id || 0
  const type = recordValue.value !== '' ? recordValue.value : ''

  const params = {
    pageNo: pageNo || page.value,
    pageSize: pageSize.value,
    startTime: startTime.value,
    endTime: endTime.value,
    memberId,
    symbol: coinType.value,
    type
  }

  loading.value = true
  axios.post(`${host}/uc/asset/transaction/all`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      loading.value = false
      if (resp.data) {
        total.value = resp.data.totalElements
        tableRecord.value = resp.data.content
      }
    } else {
      ElMessage.error(resp.message)
      loading.value = false
    }
  })
  .catch(() => {
    loading.value = false
    ElMessage.error('获取账单记录失败')
  })
}

const clear = () => {
  startTime.value = ''
  endTime.value = ''
}

onMounted(() => {
  getList(1)
  getAddrList()
})
</script>

<style scoped lang="scss">
.nav-rights {
  .nav-right {
    height: auto;
    overflow: hidden;
    padding: 0 15px;

    .bill_flow_box {
      .rightarea-con {
        .form-group {
          margin-bottom: 20px;
          text-align: left;
        }
      }
    }
  }
}
</style>

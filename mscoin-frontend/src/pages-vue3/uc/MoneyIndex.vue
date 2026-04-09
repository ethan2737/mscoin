<template>
  <div class="nav-rights">
    <div class="nav-right">
      <div class="bill_box">
        <div class="rightarea-con">
          <div class="money_table">
            <div class="table-header">
              <div class="total-assets">
                <span>{{ $t('uc.finance.money.totalassets') }}：</span>
                <span style="font-size: 18px; color: #D8E1EB;">${{ totalUSDT }}</span>
                <span style="font-size: 10px; color: #828ea1; margin-left: 5px;">≈ ¥{{ totalCny }}</span>
              </div>
              <el-input
                v-model="searchKey"
                :placeholder="$t('common.searchplaceholder')"
                class="search-input"
                @input="seachInputChange"
              >
                <template #prefix>
                  <el-icon><Search /></el-icon>
                </template>
              </el-input>
            </div>
            <el-table :data="tableMoneyShow" v-loading="loading" border style="width: 100%">
              <el-table-column prop="coinType" :label="$t('uc.finance.money.cointype')" align="center" width="100" />
              <el-table-column :label="$t('uc.finance.money.balance')" align="center">
                <template #default="{ row }">
                  <span :title="row.balance">{{ toFloor(row.balance || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.money.frozen')" align="center">
                <template #default="{ row }">
                  <span :title="row.frozenBalance">{{ toFloor(row.frozenBalance || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.money.needreleased')" align="center">
                <template #default="{ row }">
                  <span :title="row.toReleased">{{ toFloor(row.toReleased || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.finance.money.operate')" align="center">
                <template #default="{ row }">
                  <el-button
                    v-if="canCharge(row)"
                    type="info"
                    size="small"
                    style="margin-right: 8px;"
                    @click="goRecharge(row.coin.unit)"
                  >
                    {{ $t('uc.finance.money.charge') }}
                  </el-button>
                  <el-button
                    v-else
                    size="small"
                    disabled
                    style="margin-right: 8px;"
                  >
                    {{ $t('uc.finance.money.charge') }}
                  </el-button>
                  <el-button
                    v-if="row.coin.canWithdraw == 1"
                    type="danger"
                    size="small"
                    style="margin-right: 8px;"
                    @click="goWithdraw(row.coin.unit)"
                  >
                    {{ $t('uc.finance.money.pickup') }}
                  </el-button>
                  <el-button
                    v-else
                    size="small"
                    disabled
                    style="margin-right: 8px;"
                  >
                    {{ $t('uc.finance.money.pickup') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>
      </div>
    </div>

    <!-- 配对对话框 -->
    <el-dialog v-model="modal" :title="$t('uc.finance.money.match')" width="450px">
      <p style="font-weight: bold; padding: 10px 0;">
        {{ $t('uc.finance.money.matchtip1') }}：{{ GCCMatchAmount }}
      </p>
      <el-form :inline="true">
        <el-form-item :label="$t('uc.finance.money.matchtip2')">
          <el-input-number v-model="matchAmount" :placeholder="$t('uc.finance.money.matchtip2')" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button type="warning" @click="matchGCC">{{ $t('uc.finance.money.match') }}</el-button>
      </template>
    </el-dialog>

    <!-- 消息对话框 -->
    <el-dialog v-model="modal_msg" :title="$t('uc.finance.money.match')" width="450px">
      <p>{{ match_msg }}</p>
      <template #footer>
        <el-button type="primary" @click="modal_msg = false">{{ $t('common.confirm') }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 法币资产页面
 * 迁移点:
 * 1. 使用 <script setup> 语法和 Composition API
 * 2. Element Plus 组件替代 iView 组件
 * 3. vue-resource 替换为 axios
 * 4. h() render 函数改为 template slot 语法
 */
import { ref, computed, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import axios from 'axios'

const store = inject('store')
const router = inject('router')

const host = ''

const GCCMatchAmount = ref(0)
const matchAmount = ref(0)
const modal = ref(false)
const modal_msg = ref(false)
const match_msg = ref('')
const loading = ref(true)
const canMatch = ref(true)
const searchKey = ref('')
const tableMoney = ref([])
const tableMoneyShow = ref([])

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

const seachInputChange = () => {
  tableMoneyShow.value = tableMoney.value.filter(item => item.coinType.indexOf(searchKey.value) === 0)
}

const getMoney = () => {
  axios.post(`${host}/uc/asset/wallet`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      tableMoney.value = resp.data
      for (let i = 0; i < tableMoney.value.length; i++) {
        tableMoney.value[i].coinType = tableMoney.value[i].coin.unit
      }
      loading.value = false
      tableMoneyShow.value = tableMoney.value
    } else {
      ElMessage.error($t('common.logintip'))
    }
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
    loading.value = false
  })
}

const getGCCMatchAmount = () => {
  axios.post(`${host}/uc/asset/wallet/match-check`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      canMatch.value = true
      GCCMatchAmount.value = resp.data
    } else {
      canMatch.value = false
      match_msg.value = resp.message
    }
    showMatchDailog()
  })
  .catch(() => {
    canMatch.value = false
    match_msg.value = '获取配对信息失败'
    showMatchDailog()
  })
}

const showMatchDailog = () => {
  if (canMatch.value) {
    modal.value = true
  } else {
    modal_msg.value = true
  }
}

const matchGCC = () => {
  if (matchAmount.value <= 0) {
    ElMessage.warning($t('uc.finance.money.matcherr1'))
  } else if (matchAmount.value > GCCMatchAmount.value) {
    ElMessage.warning($t('uc.finance.money.matcherr2'))
  } else {
    axios.post(`${host}/uc/asset/wallet/match`, { amount: matchAmount.value }, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json',
        'x-auth-token': localStorage.getItem('TOKEN')
      }
    })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success($t('uc.finance.money.matchsuccess'))
        GCCMatchAmount.value = GCCMatchAmount.value - matchAmount.value
        modal.value = false
      } else {
        ElMessage.error(resp.message)
      }
    })
    .catch(() => {
      ElMessage.error('配对失败')
    })
  }
}

const canCharge = (row) => {
  return row.coin.canRecharge === 1 &&
         ((row.address != null && row.address !== '') || row.coin.accountType === 1)
}

const goRecharge = (unit) => {
  router.push(`/uc/recharge?name=${unit}`)
}

const goWithdraw = (unit) => {
  router.push(`/uc/withdraw?name=${unit}`)
}

const totalUSDT = computed(() => {
  let usdtTotal = 0
  for (let i = 0; i < tableMoney.value.length; i++) {
    usdtTotal += (tableMoney.value[i].balance + tableMoney.value[i].frozenBalance) * tableMoney.value[i].coin.usdRate
  }
  return usdtTotal.toFixed(2)
})

const totalCny = computed(() => {
  let cnyTotal = 0
  for (let i = 0; i < tableMoney.value.length; i++) {
    cnyTotal += (tableMoney.value[i].balance + tableMoney.value[i].frozenBalance) * tableMoney.value[i].coin.cnyRate
  }
  return cnyTotal.toFixed(2)
})

onMounted(() => {
  getMoney()
})
</script>

<style scoped lang="scss">
.nav-rights {
  .nav-right {
    height: auto;
    overflow: hidden;
    padding: 0 15px;

    .bill_box {
      .rightarea-con {
        .money_table {
          .table-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 20px;

            .total-assets {
              font-size: 12px;
              color: #828ea1;
            }

            .search-input {
              width: 200px;
            }
          }
        }
      }
    }
  }
}
</style>

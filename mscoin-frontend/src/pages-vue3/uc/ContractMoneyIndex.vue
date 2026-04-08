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
                    type="primary"
                    size="small"
                    @click="openTransfer(row)"
                  >
                    {{ $t('uc.finance.money.transfer') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>
      </div>
    </div>

    <!-- 划转对话框 -->
    <el-dialog v-model="transferShow" :title="$t('uc.finance.money.transfer')" width="450px">
      <div class="transfer-info" style="margin-bottom: 20px;">
        <div style="display: flex; align-items: center; justify-content: space-between;">
          <div>
            <p style="font-size: 12px; color: #828ea1;">
              {{ direction ? $t('uc.finance.money.from') : $t('uc.finance.money.to') }}
            </p>
            <p style="font-size: 16px; color: #fff;">{{ $t('uc.finance.money.exchangeAccount') }}</p>
          </div>
          <el-icon
            :size="24"
            @click="changeDirection"
            style="cursor: pointer; transform: rotate(2160deg);"
          >
            <Refresh />
          </el-icon>
          <div>
            <p style="font-size: 12px; color: #828ea1;">
              {{ direction ? $t('uc.finance.money.to') : $t('uc.finance.money.from') }}
            </p>
            <p style="font-size: 16px; color: #fff;">{{ $t('uc.finance.money.swapAccount') }}</p>
          </div>
        </div>
      </div>

      <el-form :inline="true">
        <el-form-item :label="$t('uc.finance.money.num')">
          <el-input-number
            v-model="transferAmount"
            :placeholder="$t('uc.finance.money.transfertip')"
            :min="0"
            :max="direction ? transfer.mainBalance : transfer.balance"
            style="width: 200px;"
          />
        </el-form-item>
      </el-form>

      <div style="margin-bottom: 15px; color: #ccc;">
        ({{ $t('uc.finance.money.available') }} {{ transfer.coin?.name }}
        {{ direction ? transfer.mainBalance : transfer.balance }})
      </div>

      <template #footer>
        <el-button @click="transferShow = false">{{ $t('common.close') }}</el-button>
        <el-button type="primary" @click="doTransfer">{{ $t('common.ok') }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 合约资产页面
 */
import { ref, reactive, computed, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Search, Refresh } from '@element-plus/icons-vue'
import axios from 'axios'

const host = 'http://localhost'

const direction = ref(true)
const loading = ref(true)
const searchKey = ref('')
const tableMoney = ref([])
const tableMoneyShow = ref([])
const transferShow = ref(false)
const transferAmount = ref(null)
const transfer = reactive({
  coin: null,
  balance: 0,
  mainBalance: 0
})

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
  tableMoneyShow.value = tableMoney.value.filter(item =>
    item.coinType.indexOf(searchKey.value) === 0
  )
}

const changeDirection = () => {
  direction.value = !direction.value
}

const openTransfer = (row) => {
  transfer.coin = row.coin
  transfer.balance = row.balance
  transfer.mainBalance = row.mainBalance
  transferShow.value = true
}

const doTransfer = () => {
  const limit = direction.value ? transfer.mainBalance : transfer.balance

  if (transferAmount.value > limit) {
    ElMessage.error($t('uc.finance.money.transferout'))
    return
  }

  axios.post(`${host}/uc/contract-wallet/transfer`, {
    unit: transfer.coin.unit,
    amount: transferAmount.value,
    direction: direction.value ? 1 : 2
  }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success($t('uc.finance.money.transfersuccess'))
      transferShow.value = false
      transfer.coin = null
      transferAmount.value = null
      getMoney()
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
  })
}

const getMoney = () => {
  axios.post(`${host}/uc/contract-wallet`, {}, {
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

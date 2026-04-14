<template>
  <div class="nav-rights">
    <div class="nav-right">
      <div class="bill_box">
        <div class="rightarea-con">
          <div class="money_table">
            <div class="table-header">
              <div class="total-assets">
                <span>{{ t('uc.finance.money.totalassets') }}：</span>
                <span style="font-size: 18px; color: #D8E1EB;">${{ totalUSDT }}</span>
                <span style="font-size: 10px; color: #828ea1; margin-left: 5px;">≈ ¥{{ totalCny }}</span>
              </div>
              <el-input
                v-model="searchKey"
                :placeholder="t('common.searchplaceholder')"
                class="search-input"
                @input="seachInputChange"
              >
                <template #prefix>
                  <el-icon><Search /></el-icon>
                </template>
              </el-input>
            </div>
            <el-table :data="tableMoneyShow" v-loading="loading" border style="width: 100%">
              <el-table-column prop="coinType" :label="t('uc.finance.money.cointype')" align="center" width="100" />
              <el-table-column :label="t('uc.finance.money.balance')" align="center">
                <template #default="{ row }">
                  <span :title="row.balance">{{ toFloor(row.balance || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="t('uc.finance.money.frozen')" align="center">
                <template #default="{ row }">
                  <span :title="row.frozenBalance">{{ toFloor(row.frozenBalance || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="t('uc.finance.money.needreleased')" align="center">
                <template #default="{ row }">
                  <span :title="row.toReleased">{{ toFloor(row.toReleased || '0') }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="t('uc.finance.money.operate')" align="center">
                <template #default="{ row }">
                  <el-button
                    v-if="canCharge(row)"
                    type="info"
                    size="small"
                    style="margin-right: 8px;"
                    @click="goRecharge(row.coin.unit)"
                  >
                    {{ t('uc.finance.money.charge') }}
                  </el-button>
                  <el-button
                    v-else
                    size="small"
                    disabled
                    style="margin-right: 8px;"
                  >
                    {{ t('uc.finance.money.charge') }}
                  </el-button>
                  <el-button
                    v-if="row.coin.canWithdraw == 1"
                    type="danger"
                    size="small"
                    style="margin-right: 8px;"
                    @click="goWithdraw(row.coin.unit)"
                  >
                    {{ t('uc.finance.money.pickup') }}
                  </el-button>
                  <el-button
                    v-else
                    size="small"
                    disabled
                    style="margin-right: 8px;"
                  >
                    {{ t('uc.finance.money.pickup') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>
      </div>
    </div>
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
import { useI18n } from 'vue-i18n'

const store = inject('store')
const router = inject('router')
const { t } = useI18n()

const host = ''

const loading = ref(true)
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
  const token = localStorage.getItem('TOKEN')
  if (!token) {
    ElMessage.error(t('common.logintip'))
    loading.value = false
    return
  }

  axios.post(`${host}/uc/asset/wallet`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': token
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 4000) {
      ElMessage.error(t('common.logintip'))
      loading.value = false
      return
    }
    if (resp.code === 0) {
      tableMoney.value = resp.data
      for (let i = 0; i < tableMoney.value.length; i++) {
        tableMoney.value[i].coinType = tableMoney.value[i].coin.unit
      }
      loading.value = false
      tableMoneyShow.value = tableMoney.value
    } else {
      ElMessage.error(resp.message || '获取资产信息失败')
      loading.value = false
    }
  })
  .catch((error) => {
    console.error('getMoney error:', error)
    ElMessage.error('网络请求失败，请稍后重试')
    loading.value = false
  })
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
    background: #192330;
    min-height: 600px;

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

        // 覆盖 Element Plus 表格默认白色背景
        :deep(.el-table) {
          background-color: transparent !important;

          .el-table__body tr {
            background-color: #192330 !important;

            td {
              background-color: #192330 !important;
              color: #fff !important;
              border-color: #27313e !important;
            }
          }

          .el-table__header tr {
            background-color: #27313e !important;

            th {
              background-color: #27313e !important;
              color: #fff !important;
              border-color: #27313e !important;
            }
          }
        }
      }
    }
  }
}
</style>

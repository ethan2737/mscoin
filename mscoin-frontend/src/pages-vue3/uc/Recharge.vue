<template>
  <div class="nav-rights">
    <div class="nav-right">
      <div class="bill_box">
        <section>
          <div class="table-inner action-box open">
            <div class="action-inner">
              <div class="inner-left">
                <p class="describe">{{ $t('uc.finance.recharge.symbol') }}</p>
                <el-select v-model="coinType" style="width:100px;margin-top: 23px;" @change="changeCoin">
                  <el-option v-for="item in coinList" :key="item.coin.unit" :value="item.coin.unit">{{ item.coin.unit }}</el-option>
                </el-select>
              </div>
              <div class="inner-box deposit-address">
                <p class="describe">{{ $t('uc.finance.recharge.address') }}</p>
                <div class="title">
                  <el-input v-model="qrcode.value" readonly style="width: 400px;color:#8c979f;"></el-input>
                  <a v-show="isShowGetAddress" class="link-copy" @click="resetAddress">{{ $t('uc.finance.recharge.getaddress') }}</a>
                  <a @click="copyAddress" class="link-copy">{{ $t('uc.finance.recharge.copy') }}</a>
                  <a class="link-qrcode" @click="showEwm">
                    {{ $t('uc.finance.recharge.qrcode') }}
                    <el-dialog v-model="isShowEwm" width="280px" :show-close="true">
                      <template #header>
                        <p style="text-align: center;">{{ $t('uc.finance.recharge.qrcodeaddress') }}</p>
                      </template>
                      <div class="show-qrcode" style="text-align: center;padding: 15px 10px;border-radius:10px;background:#FFF;">
                        <qrcode-vue :value="qrcode.value" :size="qrcode.size" level="H" />
                      </div>
                    </el-dialog>
                  </a>
                </div>
                <p v-if="accountType != 0" style="margin-top: 10px;font-size:12px;color:#828ea1;">
                  Memo：<span style="font-size: 20px;color: #F90;font-weight:bold;">{{ memoCode }}</span>
                </p>
                <p v-if="accountType != 0" style="margin-top: 10px;font-size:12px;color:#828ea1;">
                  {{ $t('uc.finance.recharge.memotips') }}
                  <a style="color: #f0a70a;" @click="copyMemo" class="link-copy">{{ $t('uc.finance.recharge.copy') }} Memo</a>
                </p>
              </div>
            </div>
            <div class="action-content">
              <div class="action-body">
                <p class="acb-p1">{{ $t('common.tip') }}</p>
                <p class="acb-p2">
                  • {{ $t('uc.finance.recharge.msg3') }}{{ minRechargeAmount }}{{ coinType }}{{ $t('uc.finance.recharge.msg3_1') }}<br>
                  • {{ $t('uc.finance.recharge.msg1') }}<br>
                  • {{ $t('uc.finance.recharge.msg2') }}<br>
                  • {{ $t('uc.finance.recharge.msg4') }}<br>
                  • {{ $t('uc.finance.recharge.msg5') }}
                </p>
              </div>
            </div>
            <div class="action-content">
              <div class="action-body">
                <p class="acb-p1">{{ $t('uc.finance.recharge.record') }}</p>
                <div class="order-table">
                  <el-table :data="tableRecharge" v-loading="loading" border style="width: 100%">
                    <el-table-column prop="createTime" :label="$t('uc.finance.recharge.time')" align="center" width="200" />
                    <el-table-column prop="symbol" :label="$t('uc.finance.recharge.symbol')" align="center" width="120" />
                    <el-table-column prop="address" :label="$t('uc.finance.recharge.address')" align="center" />
                    <el-table-column prop="amount" :label="$t('uc.finance.recharge.amount')" align="center" width="100" />
                  </el-table>
                  <div style="margin: 10px;overflow: hidden">
                    <div style="float: right;">
                      <el-pagination
                        layout="total, prev, pager, next"
                        :total="dataCount"
                        :page-size="10"
                        :current-page="currentPage"
                        @current-change="changePage"
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 充值页面
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. iView 组件替换为 Element Plus
 * 4. vue-qriously 替换为 qrcode-vue
 * 5. vue-resource 替换为 axios
 * 6. v-clipboard 替换为 navigator.clipboard API
 */
import { ref, reactive, computed, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import QrcodeVue from 'qrcode-vue'
import axios from 'axios'

const store = inject('store')
const router = inject('router')

const host = ''

const accountType = ref(0)
const memoCode = ref('')
const minRechargeAmount = ref('0.001')
const isShowGetAddress = ref(false)
const isShowEwm = ref(false)
const dataCount = ref(0)
const currentPage = ref(1)
const loading = ref(true)
const coinType = ref('')
const coinList = ref([])
const tableRecharge = ref([])
const allTableRecharge = ref([])

const qrcode = reactive({
  value: '',
  size: 220,
  coinName: '',
  unit: ''
})

const copyAddress = async () => {
  try {
    await navigator.clipboard.writeText(qrcode.value)
    ElMessage.success($t('uc.finance.recharge.copysuccess') + qrcode.value)
  } catch (err) {
    ElMessage.error($t('uc.finance.recharge.copysuccess'))
  }
}

const copyMemo = async () => {
  try {
    await navigator.clipboard.writeText(memoCode.value)
    ElMessage.success('Memo ' + $t('uc.finance.recharge.copysuccess'))
  } catch (err) {
    ElMessage.error($t('uc.finance.recharge.copysuccess'))
  }
}

const showEwm = () => {
  isShowEwm.value = !isShowEwm.value
}

const changeCoin = (value) => {
  for (let i = 0; i < coinList.value.length; i++) {
    if (coinList.value[i].coin.unit === value) {
      qrcode.value = coinList.value[i].address
      qrcode.coinName = coinList.value[i].coin.name.toLowerCase()
      qrcode.unit = coinList.value[i].coin.unit
      minRechargeAmount.value = coinList.value[i].coin.minRechargeAmount

      if (coinList.value[i].coin.accountType === 1) {
        qrcode.value = coinList.value[i].coin.depositAddress
        memoCode.value = coinList.value[i].memo
        accountType.value = 1
      } else {
        accountType.value = 0
      }

      if (!qrcode.value) {
        isShowGetAddress.value = true
      } else {
        isShowGetAddress.value = false
      }
    }
  }
  getCurrentCoinRecharge()
}

const resetAddress = () => {
  if (!qrcode.value) {
    const params = { unit: qrcode.unit }
    axios.post(`${host}/uc/asset/wallet/reset-address`, params, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json',
        'x-auth-token': localStorage.getItem('TOKEN')
      }
    })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        setTimeout(() => {
          location.reload()
        }, 3000)
      } else {
        ElMessage.error(resp.message)
      }
    })
    .catch(() => {
      ElMessage.error('获取地址失败')
    })
  }
}

const getMoney = () => {
  axios.post(`${host}/uc/approve/wallet/coin/list`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      for (let i = 0; i < resp.data.length; i++) {
        const coin = resp.data[i]
        if (coin.coin.canRecharge === 1) {
          coinList.value.push(coin)
        }
      }
      if (coinList.value.length > 0 && !coinType.value) {
        coinType.value = coinList.value[0].coin.unit
      }
      changeCoin(coinType.value)
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error('获取币种列表失败')
  })
}

const getCurrentCoinRecharge = () => {
  if (coinType.value !== '') {
    const temp = []
    for (let i = 0; i < allTableRecharge.value.length; i++) {
      if (allTableRecharge.value[i].symbol === coinType.value) {
        temp.push(allTableRecharge.value[i])
      }
    }
    tableRecharge.value = temp
  } else {
    tableRecharge.value = allTableRecharge.value
  }
}

const changePage = (index) => {
  currentPage.value = index
  getList(index)
}

const getList = (pageN) => {
  const memberId = store?.state?.member?.id || 0
  const pageNo = pageN
  const pageSize = 10
  const type = 0
  const params = { memberId, pageNo, pageSize, type }

  axios.post(`${host}/uc/asset/transaction/all`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0 && resp.data) {
      allTableRecharge.value = resp.data.content || []
      dataCount.value = resp.data.totalElements || 0
      getCurrentCoinRecharge()
    }
    loading.value = false
  })
  .catch(() => {
    loading.value = false
    ElMessage.error('获取充值记录失败')
  })
}

onMounted(() => {
  coinType.value = router.currentRoute?.query?.name || ''
  getMoney()
  getList(1)
})
</script>

<style scoped lang="scss">
.table-inner {
  position: relative;
  text-align: left;
  border-radius: 3px;
}

.acb-p1 {
  font-size: 16px;
  font-weight: 400;
  line-height: 50px;
}

.acb-p2 {
  font-size: 13px;
  line-height: 24px;
  color: #8c979f;
}

.action-inner {
  width: 100%;
  display: table;
}

.action-inner .inner-box {
  display: table-cell;
  width: 100%;
}

.action-box .title .copy {
  user-select: text;
}

.action-box .title a.link-copy {
  font-size: 14px;
  margin-left: 20px;
  color: #f0a70a;
  cursor: pointer;
}

.action-box .title a.link-qrcode {
  margin-left: 20px;
  font-size: 14px;
  position: relative;
  color: #f0a70a;
  cursor: pointer;
}

.action-content {
  width: 100%;
  margin-top: 30px;
  display: table;
  color: #ccc;
}

.action-box .title {
  margin-top: 20px;
  font-size: 20px;
  user-select: none;
}

.action-box .title .show-qrcode {
  position: absolute;
  top: -100px;
  left: 40px;
  padding: 10px;
  background: #FFF;
}

.action-inner .inner-box.deposit-address {
  width: 80%;
}

p.describe {
  font-size: 16px;
  font-weight: 600;
}

.bill_box {
  width: 100%;
  height: auto;
}

@media screen and (max-width: 1200px) {
  .bill_box {
    width: 90%;
    height: auto;
    margin-left: 5%;
  }
}

.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 15px;
}
</style>

<template>
  <div class="nav-rights withdraw">
    <div class="nav-right">
      <div class="rightarea">
        <section class="trade-groups merchant-tops">
          <router-link to="/uc/withdraw/address">{{ $t('uc.finance.withdraw.addressmanager') }}</router-link>
        </section>
        <section>
          <div class="table-inner action-box">
            <div class="action-inner">
              <div class="inner-left">
                <p class="describe">{{ $t('uc.finance.withdraw.symbol') }}</p>
                <el-select v-model="coinType" style="width:100px;margin-top: 14px;" @change="getAddrList">
                  <el-option v-for="item in coinList" :key="item.unit" :value="item.unit">{{ item.unit }}</el-option>
                </el-select>
              </div>
              <div class="inner-box">
                <div class="form-group form-address">
                  <label class="describe">{{ $t('uc.finance.withdraw.address') }}</label>
                  <div class="control-input-group">
                    <el-select v-model="withdrawAddress" filterable clearable @change="onAddressChange" :placeholder="$t('common.pleaseselect')">
                      <el-option v-for="item in currentCoin.addresses" :key="item.address" :value="item.address">{{ item.remark + '(' + item.address + ')' }}</el-option>
                    </el-select>
                  </div>
                </div>
              </div>
            </div>
            <div class="form-group-container">
              <div class="form-group form-amount">
                <label class="label-amount">
                  {{ $t('uc.finance.withdraw.num') }}
                  <p class="label-fr">
                    <span>【{{ $t('uc.finance.withdraw.avabalance') }}】：<span class="label-pointer">{{ toFloor(currentCoin.balance) }}</span></span>
                    <span v-if="currentCoin.enableAutoWithdraw === 0">
                      【{{ $t('common.tip') }}】：{{ $t('uc.finance.withdraw.msg1') }} {{ currentCoin.threshold }} {{ $t('uc.finance.withdraw.msg2') }}
                    </span>
                  </p>
                </label>
                <div class="input-group">
                  <el-popover placement="top" :width="300" trigger="focus">
                    <template #reference>
                      <el-input-number
                        v-model="withdrawAmount"
                        :placeholder="$t('uc.finance.withdraw.numtip1')"
                        size="large"
                        :min="currentCoin.minAmount"
                        :max="currentCoin.maxAmount"
                        @change="computerAmount"
                        style="width: 100%;"
                      />
                    </template>
                    <div>
                      <p>{{ $t('uc.finance.withdraw.tip1') }}{{ currentCoin.withdrawScale }}{{ $t('uc.finance.withdraw.tip11') }}</p>
                      <p>{{ $t('uc.finance.withdraw.tip2') }}{{ currentCoin.minAmount }}, {{ currentCoin.maxAmount }}</p>
                    </div>
                  </el-popover>
                  <span class="input-group-addon addon-tag uppercase firstt">{{ currentCoin.unit }}</span>
                </div>
              </div>
            </div>
            <div class="form-group-container form-group-container2">
              <div class="form-group form-fee">
                <label class="label-amount">{{ $t('uc.finance.withdraw.fee') }}</label>
                <div class="input-group" style="margin-top:14px;position:relative;">
                  <el-input-number
                    readonly
                    v-model="withdrawFee"
                    :min="currentCoin.minTxFee"
                    :max="currentCoin.maxTxFee"
                    size="large"
                    style="width: 100%;"
                  />
                  <span class="input-group-addon addon-tag uppercase">{{ currentCoin.unit }}</span>
                </div>
              </div>
              <div class="form-group">
                <label>{{ $t('uc.finance.withdraw.arriamount') }}</label>
                <div class="input-group" style="margin-top:14px;position:relative;">
                  <el-input-number
                    readonly
                    v-model="withdrawOutAmount"
                    :placeholder="$t('uc.finance.withdraw.arriamount')"
                    size="large"
                    style="width: 100%;"
                  />
                  <span class="input-group-addon addon-tag uppercase">{{ currentCoin.unit }}</span>
                </div>
              </div>
            </div>
            <div class="action-foot">
              <el-button type="warning" size="large" style="height:40px;width: 100%;" @click="apply">{{ $t('uc.finance.withdraw.pickup') }}</el-button>
            </div>
            <div class="action-content pt10">
              <div class="action-body">
                <p class="acb-p1">{{ $t('common.tip') }}</p>
                <p class="acb-p2">
                  • {{ $t('uc.finance.withdraw.msg3') }}：{{ currentCoin.minAmount }} {{ coinType }}。<br>
                  • {{ $t('uc.finance.withdraw.msg5') }}<br>
                  • {{ $t('uc.finance.withdraw.msg6') }}
                </p>
              </div>
            </div>
            <div class="action-content">
              <div class="action-body">
                <p class="acb-p1">{{ $t('uc.finance.withdraw.record') }}</p>
                <div class="order-table">
                  <p class="acb-p2" style="margin-bottom:10px;">• {{ $t('uc.finance.withdraw.click') }}<i class="ivu-icon ivu-icon-funnel"></i>{{ $t('uc.finance.withdraw.filtrate') }}</p>
                  <el-table :data="tableWithdraw" v-loading="loading" border style="width: 100%">
                    <el-table-column prop="createTime" :label="$t('uc.finance.withdraw.time')" width="180" />
                    <el-table-column prop="coin.unit" :label="$t('uc.finance.withdraw.symbol')" />
                    <el-table-column prop="address" :label="$t('uc.finance.withdraw.address')" />
                    <el-table-column prop="totalAmount" :label="$t('uc.finance.withdraw.num')" />
                    <el-table-column prop="fee" :label="$t('uc.finance.withdraw.fee')" />
                    <el-table-column prop="transactionNumber" :label="$t('uc.finance.withdraw.txid')" />
                    <el-table-column :label="$t('uc.finance.withdraw.status')">
                      <template #default="{ row }">
                        <span>
                          {{ row.status === 0 ? $t('uc.finance.withdraw.status_1') :
                             row.status === 1 ? $t('uc.finance.withdraw.status_2') :
                             row.status === 2 ? $t('uc.finance.withdraw.status_3') :
                             row.status === 3 ? $t('uc.finance.withdraw.status_4') : '' }}
                        </span>
                      </template>
                    </el-table-column>
                  </el-table>
                  <div style="float: right; margin: 10px;">
                    <el-pagination
                      layout="total, prev, pager, next"
                      :total="transaction.total"
                      :page-size="transaction.pageSize"
                      :current-page="transaction.page + 1"
                      @current-change="changePage"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>

    <el-dialog v-model="modal" width="450px" :show-close="true">
      <template #header>
        <p>提示</p>
      </template>
      <el-form :model="formInline" label-width="100px">
        <el-form-item :label="$t('uc.regist.smscode')">
          <div style="position: relative;">
            <el-input v-model="formInline.code" />
            <el-button
              type="warning"
              size="small"
              style="position: absolute; right: 0; top: 0;"
              @click="sendCode"
              :disabled="codeIsSending"
            >
              {{ sendcodeValue }}
            </el-button>
          </div>
        </el-form-item>
        <el-form-item :label="$t('otc.chat.msg7')">
          <el-input v-model="formInline.fundpwd" type="password" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="cancel">{{ $t('uc.finance.withdraw.cancel') }}</el-button>
        <el-button type="warning" @click="ok">{{ $t('uc.finance.withdraw.confirm') }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 提现页面
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. iView 组件替换为 Element Plus
 * 4. vue-resource 替换为 axios
 * 5. Modal 替换为 el-dialog
 * 6. Poptip 替换为 el-popover
 */
import { ref, reactive, watch, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

const store = inject('store')
const router = inject('router')

const host = 'http://localhost'

const codeIsSending = ref(false)
const sendcodeValue = ref('发送验证码')
const countdown = ref(60)
const modal = ref(false)
const withdrawAddress = ref('')
const withdrawAmount = ref(0)
const withdrawFee = ref(0)
const withdrawOutAmount = ref(0)
const coinType = ref('')
const coinList = ref([])
const tableWithdraw = ref([])
const allTableWithdraw = ref([])
const loading = ref(true)

const formInline = reactive({
  code: '',
  fundpwd: ''
})

const currentCoin = reactive({
  unit: '',
  addresses: [],
  balance: 0,
  enableAutoWithdraw: 0,
  threshold: 0,
  minAmount: 0,
  maxAmount: 0,
  minTxFee: 0,
  maxTxFee: 0,
  withdrawScale: 8
})

const transaction = reactive({
  page: 0,
  pageSize: 10,
  total: 0
})

watch(currentCoin, () => {
  withdrawFee.value = currentCoin.minTxFee + (currentCoin.maxTxFee - currentCoin.minTxFee) / 2
}, { deep: true })

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

const round = (v, e) => {
  let t = 1
  for (; e > 0; t *= 10, e--);
  for (; e < 0; t /= 10, e++);
  return Math.round(v * t) / t
}

const accSub = (arg1, arg2) => {
  let r1, r2, m, n
  try { r1 = arg1.toString().split('.')[1].length } catch (e) { r1 = 0 }
  try { r2 = arg2.toString().split('.')[1].length } catch (e) { r2 = 0 }
  m = Math.pow(10, Math.max(r1, r2))
  n = r1 >= r2 ? r1 : r2
  return ((arg1 * m - arg2 * m) / m).toFixed(n)
}

const computerAmount = () => {
  withdrawOutAmount.value = round(accSub(withdrawAmount.value, withdrawFee.value), currentCoin.withdrawScale)
}

const onAddressChange = (data) => {
  // 地址变化处理
}

const clearValues = () => {
  withdrawAddress.value = ''
  withdrawAmount.value = 0
  withdrawOutAmount.value = 0
}

const getCurrentCoinRecharge = () => {
  if (coinType.value !== '') {
    const temp = []
    for (let i = 0; i < allTableWithdraw.value.length; i++) {
      if (allTableWithdraw.value[i].coin?.unit === coinType.value) {
        temp.push(allTableWithdraw.value[i])
      }
    }
    tableWithdraw.value = temp
  } else {
    tableWithdraw.value = allTableWithdraw.value
  }
}

const changePage = (index) => {
  transaction.page = index - 1
  getList()
}

const cancel = () => {
  modal.value = false
  formInline.code = ''
  formInline.fundpwd = ''
}

const sendCode = () => {
  axios.post(`${host}/uc/mobile/withdraw/code`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      settime()
      ElMessage.success(resp.message)
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error('发送验证码失败')
  })
}

const settime = () => {
  sendcodeValue.value = countdown.value
  codeIsSending.value = true
  const timercode = setInterval(() => {
    countdown.value--
    sendcodeValue.value = countdown.value
    if (countdown.value <= 0) {
      clearInterval(timercode)
      sendcodeValue.value = '发送验证码'
      countdown.value = 60
      codeIsSending.value = false
    }
  }, 1000)
}

const getAddrList = () => {
  clearValues()
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
      if (coinType.value) {
        for (let i = 0; i < resp.data.length; i++) {
          if (coinType.value === resp.data[i].unit) {
            Object.assign(currentCoin, resp.data[i])
            break
          }
        }
      } else {
        Object.assign(currentCoin, resp.data[0])
        coinType.value = currentCoin.unit
      }
    } else {
      ElMessage.error(resp?.message || '获取币种列表失败')
    }
  })
  .catch(() => {
    ElMessage.error('获取币种列表失败')
  })
}

const getList = () => {
  loading.value = true
  const params = {
    page: transaction.page,
    pageSize: transaction.pageSize
  }

  axios.post(`${host}/uc/withdraw/record`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      tableWithdraw.value = resp.data.content
      transaction.total = resp.data.totalElements
      transaction.page = resp.data.number
    } else {
      ElMessage.error(resp.message)
    }
    loading.value = false
  })
  .catch(() => {
    loading.value = false
    ElMessage.error('获取提现记录失败')
  })
}

const valid = () => {
  if (coinType.value === '') {
    ElMessage.error($t('uc.finance.withdraw.symboltip'))
    return false
  }
  if (!withdrawAddress.value) {
    ElMessage.error($t('uc.finance.withdraw.addresstip'))
    return false
  }
  if (!withdrawAmount.value || withdrawAmount.value < currentCoin.minAmount) {
    ElMessage.error($t('uc.finance.withdraw.numtip2') + currentCoin.minAmount)
    return false
  }
  if (withdrawAmount.value < withdrawFee.value) {
    ElMessage.error($t('uc.finance.withdraw.numtip3'))
    return false
  }
  if (!withdrawFee.value || withdrawFee.value > currentCoin.maxTxFee || withdrawFee.value < currentCoin.minTxFee) {
    ElMessage.error($t('uc.finance.withdraw.feetip1') + currentCoin.minTxFee + ' , ' + $t('uc.finance.withdraw.feetip2') + currentCoin.maxTxFee)
    return false
  }
  return true
}

const apply = () => {
  if (valid()) {
    modal.value = true
    countdown.value = 60
    sendcodeValue.value = '发送验证码'
    codeIsSending.value = false
  }
}

const ok = () => {
  if (!formInline.code) {
    modal.value = true
    ElMessage.error('请填写短信验证码')
    return
  }
  if (!formInline.fundpwd) {
    modal.value = true
    ElMessage.error($t('otc.chat.msg7tip'))
    return
  }

  const params = {
    remark: '',
    unit: currentCoin.unit,
    address: withdrawAddress.value,
    amount: withdrawAmount.value,
    fee: withdrawFee.value,
    jyPassword: formInline.fundpwd,
    code: formInline.code
  }

  axios.post(`${host}/uc/withdraw/apply/code`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      modal.value = false
      formInline.code = ''
      formInline.fundpwd = ''
      transaction.page = 0
      getList()
      clearValues()
      ElMessage.success(resp.message)
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error('提现申请失败')
  })
}

onMounted(() => {
  coinType.value = router.currentRoute?.query?.name || ''
  getAddrList()
  getList()
})
</script>

<style scoped lang="scss">
.nav-rights {
  .nav-right {
    height: auto;
    overflow: hidden;
    padding: 0 15px;

    .rightarea {
      padding-left: 15px;

      .trade-groups.merchant-tops {
        font-size: 14px;
        height: 50px;
        padding: 0 15px;
        color: #fff;
        overflow: hidden;
        display: block;
        margin-right: 0;

        a {
          display: inline-block;
          color: #f0a70a;
          width: 160px;
          height: 40px;
          border: 1px solid #f0a70a;
          line-height: 40px;
          text-align: center;
          float: right;

          &:hover {
            background: #f0a70a;
            color: #000;
          }
        }
      }

      .action-box {
        padding: 10px 20px 20px;

        .form-group-container {
          .form-group.form-amount {
            .input-group {
              .el-input-number {
                width: 100%;
              }
            }
          }
        }
      }
    }
  }
}

#sendCode {
  position: absolute;
  border: none;
  background: none;
  top: 10px;
  outline: none;
  right: 0;
  width: 30%;
  color: #f0ac19;
  cursor: pointer;
  height: 20px;
  line-height: 20px;
  border-left: 1px solid #dddee1;
}

.withdraw-form-inline {
  padding: 20px 40px 0 40px;
  .el-input {
    height: 40px;
    line-height: 40px;
  }
}

.nav-rights {
  .nav-right {
    .rightarea {
      .action-box {
        .action-inner {
          .inner-left,
          .inner-box {
            .el-select-dropdown .el-select-item {
              padding: 6px 16px;
            }
          }
        }

        .form-group-container {
          .form-group {
            .input-group {
              .el-input-number {
                width: 100%;
              }
            }
          }
        }
      }
    }
  }
}

p.describe {
  font-size: 16px;
  font-weight: 600;
}

.acb-p1 {
  font-size: 18px;
  font-weight: 600;
  line-height: 50px;
}

.acb-p2 {
  font-size: 13px;
  line-height: 24px;
  color: #8c979f;
}

.action-content.pt10 {
  padding-top: 10px;
}

.action-content {
  width: 100%;
  margin-top: 20px;
  display: table;
}

.action-content .action-body {
  display: table-cell;
  vertical-align: top;
  line-height: 20px;
  font-size: 12px;
  color: #ccc;
}

.action-foot {
  text-align: center;
  padding: 40px 170px 0;
}

.action-inner {
  width: 100%;
  display: table;
}

.action-inner .inner-box {
  display: table-cell;
  width: 80%;
}

.form-group {
  position: relative;
  margin-bottom: 20px;
  font-size: 16px;
}

.control-input-group {
  position: relative;
}

.form-group-container {
  display: table;
  width: 100%;
}

.form-group-container .form-amount {
  width: 100%;
}

.form-group-container .form-group {
  display: table-cell;
}

.form-group-container .form-group span.addon-tag:last-child {
  padding: 0;
  border: none;
  background: none;
  cursor: default;
  position: absolute;
  right: 26px;
  top: 6px;
}

.form-group-container .form-group span.addon-tag:last-child.firstt {
  top: 8px;
}

.form-group-container2 {
  padding-top: 20px;
}

.form-group-container .form-fee {
  width: 50%;
  padding: 0 20px 0 0;
}

.label-amount .label-fr {
  float: right;
  color: #aaa;
  font-size: 14px;
}

.label-amount .label-fr span {
  margin-left: 2px;
}

@media screen and (max-width: 1200px) {
  .action-foot {
    padding: 0;
  }

  .form-group-container .form-group {
    width: 100%;
    padding: 0;
    display: block;
  }
}
</style>

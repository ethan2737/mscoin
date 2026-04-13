<template>
  <div class="content-wrap">
    <div class="trade-layout" v-loading="loading">
      <aside class="merchant-panel">
        <div class="merchant-head">
          <div class="avatar-shell">
            <img
              v-if="advertiseDetail.avatar"
              :src="advertiseDetail.avatar"
              alt="merchant-avatar"
              class="avatar-image"
            />
            <span v-else class="avatar-initial">{{ usernameInitial }}</span>
          </div>
          <div class="merchant-name">{{ advertiseDetail.username }}</div>
        </div>

        <div class="merchant-status">
          <div class="status-item">
            <span>{{ advertiseDetail.emailVerified === 1 ? $t('otc.tradeinfo.emaildone') : $t('otc.tradeinfo.emailundo') }}</span>
          </div>
          <div class="status-item">
            <span>{{ advertiseDetail.phoneVerified === 1 ? $t('otc.tradeinfo.teldone') : $t('otc.tradeinfo.telundo') }}</span>
          </div>
          <div class="status-item">
            <span>{{ advertiseDetail.idCardVerified === 1 ? $t('otc.tradeinfo.idcarddone') : $t('otc.tradeinfo.idcardundo') }}</span>
          </div>
        </div>

        <div class="merchant-summary">
          <div>{{ $t('otc.tradeinfo.exchangetimes') }}: <strong>{{ advertiseDetail.transactions }}</strong></div>
          <div>{{ $t('otc.tradeinfo.location') }}: <strong>{{ $t('otc.tradeinfo.location_text') }}</strong></div>
        </div>
      </aside>

      <section class="trade-main">
        <div class="trade-card">
          <div class="card-title">{{ currentActionLabel }}</div>
          <div class="detail-grid">
            <div class="detail-item">
              <span class="detail-label">{{ $t('otc.tradeinfo.price') }}</span>
              <span class="detail-value">{{ advertiseDetail.price }} CNY / {{ advertiseDetail.unit }}</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">{{ $t('otc.tradeinfo.num') }}</span>
              <span class="detail-value">{{ advertiseDetail.number }} {{ advertiseDetail.unit }}</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">{{ $t('otc.tradeinfo.paymethod') }}</span>
              <span class="detail-value">{{ advertiseDetail.payMode }}</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">{{ $t('otc.tradeinfo.exchangelimitamount') }}</span>
              <span class="detail-value">{{ advertiseDetail.minLimit }} - {{ advertiseDetail.maxLimit }} CNY</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">{{ $t('otc.tradeinfo.exchangeperiod') }}</span>
              <span class="detail-value">{{ advertiseDetail.timeLimit }} {{ $t('otc.tradeinfo.minute') }}</span>
            </div>
          </div>
        </div>

        <div class="trade-card">
          <div class="card-title">{{ $t('otc.tradeinfo.exchangetitle') }}</div>
          <div class="trade-form">
            <el-form label-position="top">
              <div class="trade-form-row">
                <el-form-item :label="$t('otc.tradeinfo.amounttip')" class="trade-form-item">
                  <el-input v-model="buyPrice" @input="handleMoneyChange">
                    <template #prepend>CNY</template>
                  </el-input>
                  <div class="field-tip">{{ moneyFieldTip }}</div>
                </el-form-item>

                <div class="trade-switch">⇄</div>

                <el-form-item :label="$t('otc.tradeinfo.numtip')" class="trade-form-item">
                  <el-input v-model="buyAmount" @input="handleAmountChange">
                    <template #prepend>{{ advertiseDetail.unit }}</template>
                  </el-input>
                  <div class="field-tip">{{ amountWarning }}</div>
                </el-form-item>
              </div>

              <div class="confirm-panel">
                <div class="confirm-summary">
                  <span>{{ currentActionLabel }}:</span>
                  <strong>{{ buyPrice || 0 }} CNY / {{ buyAmount || 0 }} {{ advertiseDetail.unit }}</strong>
                </div>
                <el-button
                  type="warning"
                  class="submit-button"
                  :disabled="submitDisabled"
                  @click="submitOrder"
                >
                  {{ currentButtonLabel }}
                </el-button>
              </div>
            </el-form>
          </div>
        </div>

        <div class="trade-card">
          <div class="card-title">{{ $t('otc.tradeinfo.remarktitle') }}</div>
          <div class="remark-content">{{ advertiseDetail.remark || '-' }}</div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { ElForm, ElFormItem, ElInput, ElMessage, ElButton } from 'element-plus'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  buildOtcOrderPayload,
  calculateCryptoAmount,
  calculateFiatAmount,
  resolveTradeAction,
  roundTo,
  toOtcNumber
} from './trade-info'
import { buildOtcChatPath } from './route-helpers'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()

const host = ''

const loading = ref(true)
const submitting = ref(false)
const buyPrice = ref('')
const buyAmount = ref('')
const moneyWarning = ref('')
const amountWarning = ref('')
const advertiseDetail = ref({
  advertiseType: 1,
  avatar: '',
  username: '',
  emailVerified: 0,
  phoneVerified: 0,
  idCardVerified: 0,
  transactions: 0,
  price: 0,
  number: 0,
  unit: 'USDT',
  payMode: '',
  minLimit: 0,
  maxLimit: 0,
  timeLimit: 15,
  remark: '',
  coinId: 0,
  otcCoinId: 0
})

const tradeAction = computed(() => resolveTradeAction(advertiseDetail.value.advertiseType))
const currentActionLabel = computed(() => t(`otc.tradeinfo.${tradeAction.value.actionKey}`))
const currentButtonLabel = computed(() => t(`otc.tradeinfo.${tradeAction.value.buttonKey}`))
const usernameInitial = computed(() => String(advertiseDetail.value.username || '').trim().slice(0, 1))
const submitDisabled = computed(() => submitting.value || !buyPrice.value || !buyAmount.value)
const moneyRangeTip = computed(() => `${advertiseDetail.value.minLimit} - ${advertiseDetail.value.maxLimit} CNY`)
const moneyFieldTip = computed(() => moneyWarning.value || moneyRangeTip.value)

function validateRange(money) {
  if (money < advertiseDetail.value.minLimit || money > advertiseDetail.value.maxLimit) {
    moneyWarning.value = `${t('otc.tradeinfo.warning2')}${advertiseDetail.value.minLimit}~${advertiseDetail.value.maxLimit}`
    return false
  }

  moneyWarning.value = ''
  return true
}

function validateAmount(amount) {
  if (amount <= 0 || amount > toOtcNumber(advertiseDetail.value.number)) {
    amountWarning.value = `${t('otc.tradeinfo.warning4')}${minAmount.value}~${advertiseDetail.value.number}`
    return false
  }

  amountWarning.value = ''
  return true
}

const minAmount = computed(() => {
  const amount = calculateCryptoAmount(advertiseDetail.value.minLimit, advertiseDetail.value.price)
  return roundTo(amount, 8)
})

function handleMoneyChange(value) {
  const money = roundTo(value, 2)
  buyPrice.value = value
  buyAmount.value = money > 0 ? String(calculateCryptoAmount(money, advertiseDetail.value.price)) : ''
  validateRange(money)
  validateAmount(toOtcNumber(buyAmount.value))
}

function handleAmountChange(value) {
  const amount = roundTo(value, 8)
  buyAmount.value = value
  buyPrice.value = amount > 0 ? String(calculateFiatAmount(amount, advertiseDetail.value.price)) : ''
  validateAmount(amount)
  validateRange(toOtcNumber(buyPrice.value))
}

async function getAdvertiseDetail() {
  loading.value = true

  try {
    const response = await axios.post(`${host}/otc/order/pre`, {
      id: route.query.tradeId
    }, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json',
        'x-auth-token': localStorage.getItem('TOKEN')
      }
    })

    const resp = response.data
    if (resp.code !== 0) {
      ElMessage.error(resp.message)
      return
    }

    advertiseDetail.value = {
      ...advertiseDetail.value,
      ...resp.data
    }
  } catch {
    ElMessage.error('获取 OTC 广告详情失败')
  } finally {
    loading.value = false
  }
}

async function submitOrder() {
  const money = roundTo(buyPrice.value, 2)
  const amount = roundTo(buyAmount.value, 8)
  if (!validateRange(money) || !validateAmount(amount)) {
    ElMessage.error(t('otc.tradeinfo.warning5'))
    return
  }

  submitting.value = true
  const payload = buildOtcOrderPayload({
    advertiseId: route.query.tradeId,
    coinId: advertiseDetail.value.otcCoinId || advertiseDetail.value.coinId,
    price: advertiseDetail.value.price,
    money,
    amount
  })

  try {
    const response = await axios.post(`${host}${tradeAction.value.orderPath}`, payload, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json',
        'x-auth-token': localStorage.getItem('TOKEN')
      }
    })

    const resp = response.data
    if (resp.code !== 0) {
      ElMessage.error(resp.message)
      return
    }

    ElMessage.success(resp.message)
    router.push(buildOtcChatPath(resp.data, { source: 'trade' }))
  } catch {
    ElMessage.error('OTC 下单失败')
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  getAdvertiseDetail()
})
</script>

<style scoped lang="scss">
.content-wrap {
  min-height: 600px;
  padding: 80px 0 32px;
}

.trade-layout {
  width: min(1200px, calc(100vw - 64px));
  margin: 0 auto;
  display: grid;
  grid-template-columns: 280px 1fr;
  gap: 20px;
}

.merchant-panel,
.trade-card {
  background: #192330;
  border: 1px solid #27313e;
  border-radius: 12px;
  color: #fff;
}

.merchant-panel {
  padding: 24px 20px;
  height: fit-content;
}

.merchant-head {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 20px;
  border-bottom: 1px solid #27313e;
}

.avatar-shell {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.avatar-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-initial {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #f0a70a;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
}

.merchant-name {
  font-size: 16px;
  font-weight: 600;
}

.merchant-status,
.merchant-summary {
  display: grid;
  gap: 12px;
  padding-top: 20px;
}

.status-item,
.merchant-summary div {
  color: #c7d0dd;
  font-size: 14px;
}

.trade-main {
  display: grid;
  gap: 20px;
}

.trade-card {
  padding: 24px;
}

.card-title {
  margin-bottom: 20px;
  font-size: 18px;
  font-weight: 600;
}

.detail-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
}

.detail-item {
  padding: 16px;
  background: #111b28;
  border-radius: 10px;
}

.detail-label {
  display: block;
  margin-bottom: 8px;
  color: #96a2b4;
  font-size: 13px;
}

.detail-value {
  color: #fff;
  font-size: 15px;
  font-weight: 500;
}

.trade-form-row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 32px minmax(0, 1fr);
  gap: 16px;
  align-items: start;
}

.trade-form-item {
  margin-bottom: 0;
}

.trade-form-item :deep(.el-input) {
  --el-input-border-color: #27313e;
  --el-input-hover-border-color: #27313e;
  --el-input-focus-border-color: #27313e;
  --el-fill-color-blank: #111b28;
}

.trade-switch {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 84px;
  color: #f0a70a;
  font-size: 28px;
}

.field-tip {
  min-height: 20px;
  padding-top: 8px;
  color: #96a2b4;
  font-size: 12px;
}

.confirm-panel {
  margin-top: 16px;
  display: flex;
  gap: 16px;
  align-items: center;
}

.confirm-summary {
  flex: 1;
  padding: 18px 20px;
  border-radius: 10px;
  background: #111b28;
  color: #c7d0dd;
}

.confirm-summary strong {
  color: #f0a70a;
  margin-left: 8px;
}

.submit-button {
  min-width: 180px;
  height: 52px;
}

.remark-content {
  color: #c7d0dd;
  line-height: 1.7;
}

:deep(.el-input__wrapper) {
  background: #111b28;
  box-shadow: inset 0 0 0 1px #27313e;
}

:deep(.el-input__wrapper:hover),
:deep(.el-input__wrapper.is-focus) {
  background: #111b28;
  box-shadow: inset 0 0 0 1px #27313e !important;
}

:deep(.el-input-group__prepend) {
  background: #111b28;
  border-color: #27313e;
  border-right-color: #27313e;
  box-shadow:
    inset 0 1px 0 0 #27313e,
    inset 0 -1px 0 0 #27313e,
    inset 1px 0 0 0 #27313e;
  color: #96a2b4;
}

@media (max-width: 960px) {
  .trade-layout {
    width: calc(100vw - 24px);
    grid-template-columns: 1fr;
  }

  .detail-grid,
  .trade-form-row,
  .confirm-panel {
    grid-template-columns: 1fr;
  }

  .trade-switch {
    height: auto;
  }
}
</style>

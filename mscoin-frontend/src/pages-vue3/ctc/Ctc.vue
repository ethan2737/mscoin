<template>
  <div class="ctc-page">
    <img class="ctc-page__banner shoujiIf" src="../../assets/images/ctc-bg.jpg" alt="">
    <div class="ctc-page__container">
      <h1 class="shoujiIf">{{ $t('ctc.title') }}</h1>
      <p class="ctc-page__desc shoujiIf">{{ $t('ctc.desc') }}</p>

      <div class="ctc-page__content">
        <el-tabs v-model="activeTab" class="ctc-page__tabs">
          <el-tab-pane name="usdt">
            <template #label>
              <span>USDT{{ $t('ctc.trade') }}</span>
            </template>

            <div class="ctc-page__hero">
              <div class="ctc-page__grid">
                <section class="trade-card trade-card--buy">
                  <div class="trade-card__header">
                    <div>
                      <p class="trade-card__eyebrow">{{ $t('ctc.buyin') }} USDT</p>
                      <h3 class="trade-card__title">{{ quote.buy }} CNY</h3>
                    </div>
                    <span class="trade-card__badge">C2C</span>
                  </div>

                  <el-form class="trade-card__form" label-position="top">
                    <el-form-item :label="$t('ctc.buyprice')">
                      <el-input :model-value="quote.buy" disabled />
                    </el-form-item>
                    <el-form-item :label="$t('ctc.buynum')">
                      <el-input
                        v-model="tradeForm.buyAmount"
                        class="trade-card__number"
                        placeholder="请输入 50 - 50000 USDT"
                        inputmode="decimal"
                      />
                    </el-form-item>
                    <el-form-item :label="$t('ctc.payType')">
                      <el-select v-model="tradeForm.payType" class="trade-card__select">
                        <el-option
                          v-for="item in payTypeOptions"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                        />
                      </el-select>
                    </el-form-item>
                    <router-link class="trade-card__link" to="/uc/account">
                      {{ $t('ctc.payset') }}
                    </router-link>

                    <div class="trade-card__summary">
                      <span>{{ $t('ctc.payamount') }}</span>
                      <strong>{{ totalBuyMoney }} CNY</strong>
                    </div>
                    <p class="trade-card__tip">{{ $t('ctc.moneyTips') }}</p>

                    <el-button class="trade-card__action trade-card__action--buy" @click="openVerifyDialog('buy')">
                      {{ $t('ctc.buyin') }} USDT
                    </el-button>
                  </el-form>
                </section>

                <section class="trade-card trade-card--sell">
                  <div class="trade-card__header">
                    <div>
                      <p class="trade-card__eyebrow">{{ $t('ctc.sell') }} USDT</p>
                      <h3 class="trade-card__title">{{ quote.sell }} CNY</h3>
                    </div>
                    <span class="trade-card__badge">C2C</span>
                  </div>

                  <el-form class="trade-card__form" label-position="top">
                    <el-form-item :label="$t('ctc.sellprice')">
                      <el-input :model-value="quote.sell" disabled />
                    </el-form-item>
                    <el-form-item :label="$t('ctc.sellnum')">
                      <el-input
                        v-model="tradeForm.sellAmount"
                        class="trade-card__number"
                        placeholder="请输入 2 - 50000 USDT"
                        inputmode="decimal"
                      />
                    </el-form-item>
                    <p class="trade-card__balance">可卖数量：{{ walletBalance }} USDT</p>
                    <el-form-item :label="$t('ctc.receiveType')">
                      <el-select v-model="tradeForm.receiveType" class="trade-card__select">
                        <el-option
                          v-for="item in receiveTypeOptions"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                        />
                      </el-select>
                    </el-form-item>
                    <p class="trade-card__tip trade-card__tip--tight">{{ $t('ctc.useselfaccount') }}</p>

                    <div class="trade-card__summary">
                      <span>{{ $t('ctc.getamount') }}</span>
                      <strong>{{ totalSellMoney }} CNY</strong>
                    </div>
                    <p class="trade-card__tip">{{ $t('ctc.moneyTips') }}</p>

                    <el-button class="trade-card__action trade-card__action--sell" @click="openVerifyDialog('sell')">
                      {{ $t('ctc.sell') }} USDT
                    </el-button>
                  </el-form>
                </section>
              </div>

              <aside class="notice-card shoujiIf">
                <div class="notice-card__body">
                  <p class="notice-card__title">{{ $t('ctc.notice') }}</p>
                  <p>{{ $t('ctc.notice1') }}</p>
                  <p>{{ $t('ctc.notice2') }}</p>
                  <p>{{ $t('ctc.notice3') }}</p>
                  <p>{{ $t('ctc.notice4') }}</p>
                  <p>{{ $t('ctc.notice5') }}</p>
                  <router-link class="notice-card__more" target="_blank" to="/helpdetail?cate=2&id=40&cateTitle=交易指南">
                    {{ $t('ctc.moredetail') }}
                  </router-link>
                </div>
                <div class="notice-card__footer">
                  <router-link class="notice-card__button" to="/uc/safe">{{ $t('ctc.verifyset') }}</router-link>
                  <router-link class="notice-card__button" to="/uc/account">{{ $t('ctc.payset') }}</router-link>
                </div>
              </aside>
            </div>

            <section class="orders-card shoujiIf">
              <div class="orders-card__header">
                <div>
                  <p class="orders-card__eyebrow">C2C</p>
                  <h3 class="orders-card__title">快捷买卖订单</h3>
                </div>
              </div>

              <el-table :data="orders" v-loading="loading" empty-text="暂无记录" class="orders-card__table">
                <el-table-column prop="createTime" label="下单时间" min-width="170">
                  <template #default="{ row }">
                    {{ formatDateTime(row.createTime) }}
                  </template>
                </el-table-column>
                <el-table-column prop="orderSn" label="订单号" min-width="180" />
                <el-table-column label="方向" min-width="90">
                  <template #default="{ row }">
                    <span :class="directionClass(row.direction)">
                      {{ directionLabel(row.direction) }}
                    </span>
                  </template>
                </el-table-column>
                <el-table-column prop="amount" label="数量(USDT)" min-width="120" />
                <el-table-column prop="price" label="单价(CNY)" min-width="120" />
                <el-table-column label="总额(CNY)" min-width="120">
                  <template #default="{ row }">
                    {{ formatMoney(row.money) }}
                  </template>
                </el-table-column>
                <el-table-column label="状态" min-width="180">
                  <template #default="{ row }">
                    {{ orderStatusText(row) }}
                  </template>
                </el-table-column>
                <el-table-column label="操作" min-width="120" fixed="right">
                  <template #default="{ row }">
                    <el-button type="warning" plain size="small" @click="showOrderDetail(row.id)">
                      查看详情
                    </el-button>
                  </template>
                </el-table-column>
              </el-table>

              <div class="orders-card__pagination">
                <el-pagination
                  background
                  layout="prev, pager, next"
                  :total="total"
                  :page-size="pageSize"
                  :current-page="pageNo"
                  @current-change="loadDataPage"
                />
              </div>
            </section>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>

    <el-dialog v-model="verifyDialogVisible" width="460px" destroy-on-close class="ctc-dialog ctc-dialog--verify">
      <template #title>
        <h3 class="ctc-dialog__title">{{ verifyDialogTitle }}</h3>
      </template>
      <el-form label-position="top" class="verify-dialog">
        <p class="verify-dialog__hint">{{ verifyDialogHint }}</p>
        <el-form-item label="短信验证码">
          <div class="verify-dialog__row">
            <el-input v-model.trim="verifyForm.code" placeholder="请输入短信验证码" />
            <el-button class="ctc-dialog__secondary-btn" :disabled="codeIsSending" @click="sendCode">{{ sendCodeText }}</el-button>
          </div>
        </el-form-item>
        <el-form-item label="资金密码">
          <el-input v-model.trim="verifyForm.fundpwd" type="password" show-password placeholder="请输入资金密码" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button class="ctc-dialog__ghost-btn" @click="closeVerifyDialog">取消</el-button>
        <el-button class="ctc-dialog__primary-btn" type="warning" @click="submitOrder">{{ verifyDialogConfirmText }}</el-button>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailDialogVisible"
      width="760px"
      destroy-on-close
      class="ctc-dialog ctc-dialog--detail"
      @close="closeDetailDialog"
    >
      <template #title>
        <h3 class="ctc-dialog__title">订单详情</h3>
      </template>
      <p class="detail-dialog__status">{{ detailStatusText }}</p>

      <el-row class="detail-dialog__metrics" :gutter="12">
        <el-col :span="8">
          <p class="detail-dialog__metric-value">{{ directionLabel(detailOrder.direction) || '--' }}</p>
          <p class="detail-dialog__metric-label">订单方向</p>
        </el-col>
        <el-col :span="8">
          <p class="detail-dialog__metric-value">{{ formatMoney(detailOrder.amount) }} <span class="detail-dialog__unit">USDT</span></p>
          <p class="detail-dialog__metric-label">交易数量</p>
        </el-col>
        <el-col :span="8">
          <p class="detail-dialog__metric-value">{{ formatMoney(detailOrder.money) }} <span class="detail-dialog__unit">CNY</span></p>
          <p class="detail-dialog__metric-label">交易总额</p>
        </el-col>
      </el-row>

      <div class="detail-dialog__notice" v-if="detailOrder.direction === 0">
        <span class="detail-dialog__notice-tag">线下付款</span>
        请使用下方收款信息完成人民币付款，金额为
        <strong>{{ formatMoney(detailOrder.money) }} CNY</strong>
        <span v-if="countdownText" class="detail-dialog__countdown">{{ countdownText }}</span>
      </div>
      <div class="detail-dialog__notice" v-else>
        <span class="detail-dialog__notice-tag detail-dialog__notice-tag--sell">平台打款</span>
        平台将按您绑定的收款方式支付人民币，预计到账 <strong>{{ formatMoney(detailOrder.money) }} CNY</strong>
      </div>

      <div class="detail-dialog__panel">
        <p class="detail-dialog__panel-title">收付款信息</p>
        <div class="detail-dialog__row">
          <span>真实姓名</span>
          <strong>{{ detailOrder.realName || '--' }}</strong>
        </div>
        <div class="detail-dialog__row">
          <span>收付款方式</span>
          <strong>{{ payModeLabel(detailOrder.payMode) }}</strong>
        </div>

        <template v-if="detailOrder.payMode === 'bank'">
          <div class="detail-dialog__row">
            <span>开户银行</span>
            <strong>{{ detailOrder.bankInfo?.bank || '--' }}</strong>
          </div>
          <div class="detail-dialog__row">
            <span>开户支行</span>
            <strong>{{ detailOrder.bankInfo?.branch || '--' }}</strong>
          </div>
          <div class="detail-dialog__row">
            <span>银行卡号</span>
            <strong>{{ detailOrder.bankInfo?.cardNo || '--' }}</strong>
          </div>
        </template>

        <template v-if="detailOrder.payMode === 'alipay'">
          <div class="detail-dialog__row">
            <span>支付宝账号</span>
            <strong>{{ detailOrder.alipay?.aliNo || '--' }}</strong>
          </div>
          <img v-if="detailOrder.alipay?.qrCodeUrl" :src="detailOrder.alipay.qrCodeUrl" class="detail-dialog__qr" alt="">
        </template>

        <template v-if="detailOrder.payMode === 'wechatpay'">
          <div class="detail-dialog__row">
            <span>微信账号</span>
            <strong>{{ detailOrder.wechatPay?.wechat || '--' }}</strong>
          </div>
          <img v-if="detailOrder.wechatPay?.qrWeCodeUrl" :src="detailOrder.wechatPay.qrWeCodeUrl" class="detail-dialog__qr" alt="">
        </template>
      </div>

      <template #footer>
        <el-button v-if="canCancelDetail" class="ctc-dialog__danger-btn" type="danger" @click="cancelOrderClick">取消订单</el-button>
        <el-button v-if="canMarkPaid" class="ctc-dialog__primary-btn" type="warning" @click="payOrderClick">我已线下付款</el-button>
        <el-button class="ctc-dialog__ghost-btn" @click="closeDetailDialog">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import {
  cancelCtcOrder,
  createCtcOrder,
  fetchCtcDetail,
  fetchCtcOrders,
  fetchCtcQuote,
  fetchSecuritySetting,
  fetchUsdtWallet,
  markCtcOrderPaid,
  sendCtcCode
} from './ctc-api'
import {
  buildOrderPayload,
  calcSettlement,
  directionLabelOf,
  formatDateTime,
  formatMoney,
  getOrderStatusText,
  normalizeOrderPage,
  normalizeSecuritySetting,
  normalizeWalletBalance,
  payModeLabelOf
} from './ctc-utils'

const { t } = useI18n()
const router = useRouter()
const store = useStore()

const activeTab = ref('usdt')
const loading = ref(false)
const timer = ref(null)
const orderTimer = ref(null)
const pageNo = ref(1)
const pageSize = ref(10)
const total = ref(0)
const orders = ref([])
const quote = reactive({ buy: '0.00', sell: '0.00' })
const security = reactive({
  realVerified: false,
  fundsVerified: false,
  accountVerified: false
})
const tradeForm = reactive({
  payType: 'bank',
  receiveType: 'bank',
  buyAmount: null,
  sellAmount: null
})
const walletBalance = ref('0.00')
const verifyDialogVisible = ref(false)
const detailDialogVisible = ref(false)
const submitDirection = ref('buy')
const codeIsSending = ref(false)
const cooldown = ref(60)
const sendCodeText = ref('发送验证码')
const countdownSeconds = ref(0)
const verifyForm = reactive({
  code: '',
  fundpwd: ''
})
const detailOrder = reactive({})

const payTypeOptions = [
  { label: t('ctc.bank'), value: 'bank' },
  { label: t('ctc.alipay'), value: 'alipay' },
  { label: t('ctc.wechatpay'), value: 'wechatpay' }
]
const receiveTypeOptions = [{ label: t('ctc.bank'), value: 'bank' }]

const isLogin = computed(() => store.getters.isLogin)
const totalBuyMoney = computed(() => calcSettlement(quote.buy, tradeForm.buyAmount))
const totalSellMoney = computed(() => calcSettlement(quote.sell, tradeForm.sellAmount))
const detailStatusText = computed(() => getOrderStatusText(detailOrder))
const verifyDialogTitle = computed(() => submitDirection.value === 'buy' ? '确认买入订单' : '确认卖出订单')
const verifyDialogHint = computed(() => submitDirection.value === 'buy'
  ? '当前流程仅创建订单并校验账户安全，请在订单详情中按收款信息线下付款。'
  : '当前流程仅创建卖出订单并校验账户安全，平台会按订单状态向您的收款账户付款。')
const verifyDialogConfirmText = computed(() => submitDirection.value === 'buy' ? '确认创建买单' : '确认创建卖单')
const countdownText = computed(() => {
  if (countdownSeconds.value <= 0) {
    return ''
  }

  const minutes = `${Math.floor(countdownSeconds.value / 60)}`.padStart(2, '0')
  const seconds = `${countdownSeconds.value % 60}`.padStart(2, '0')

  return `${minutes}:${seconds}`
})
const canCancelDetail = computed(() => {
  if (!detailOrder.id) {
    return false
  }

  return (detailOrder.direction === 1 && detailOrder.status === 0) || (detailOrder.direction === 0 && detailOrder.status < 2)
})
const canMarkPaid = computed(() => detailOrder.direction === 0 && detailOrder.status === 1)

const handleRequestError = (error, fallbackMessage) => {
  const message = error?.response?.data?.message || error?.message || fallbackMessage
  ElNotification.error({ title: '提示', message })
}

const directionLabel = (direction) => directionLabelOf(direction)
const payModeLabel = (payMode) => payModeLabelOf(payMode)
const orderStatusText = (order) => getOrderStatusText(order)
const directionClass = (direction) => direction === 0 ? 'direction-buy' : 'direction-sell'

const closeVerifyDialog = () => {
  verifyDialogVisible.value = false
  verifyForm.code = ''
  verifyForm.fundpwd = ''
}

const closeDetailDialog = () => {
  detailDialogVisible.value = false
  if (orderTimer.value) {
    clearInterval(orderTimer.value)
    orderTimer.value = null
  }
}

const startSendCodeCooldown = () => {
  codeIsSending.value = true
  sendCodeText.value = `${cooldown.value}s 后重发`

  const timerId = setInterval(() => {
    cooldown.value -= 1
    sendCodeText.value = `${cooldown.value}s 后重发`

    if (cooldown.value <= 0) {
      clearInterval(timerId)
      cooldown.value = 60
      codeIsSending.value = false
      sendCodeText.value = '发送验证码'
    }
  }, 1000)
}

const refreshQuote = async () => {
  try {
    const response = await fetchCtcQuote()
    if (response.code === 0) {
      quote.buy = formatMoney(response.data?.buy)
      quote.sell = formatMoney(response.data?.sell)
    }
  } catch (error) {
    handleRequestError(error, '获取 C2C 报价失败')
  }
}

const refreshWallet = async () => {
  try {
    const response = await fetchUsdtWallet()
    if (response.code === 0) {
      walletBalance.value = normalizeWalletBalance(response.data)
    }
  } catch (error) {
    handleRequestError(error, '获取钱包余额失败')
  }
}

const refreshSecurity = async () => {
  try {
    const response = await fetchSecuritySetting()
    if (response.code === 0) {
      Object.assign(security, normalizeSecuritySetting(response.data))
    } else {
      ElNotification.error({ title: '提示', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '鑾峰彇瀹夊叏璁剧疆澶辫触')
  }
}

const refreshOrders = async () => {
  loading.value = true

  try {
    const response = await fetchCtcOrders({
      pageNo: pageNo.value,
      pageSize: pageSize.value
    })

    if (response.code === 0) {
      const page = normalizeOrderPage(response.data)
      orders.value = page.content
      total.value = page.totalElements
    } else {
      ElNotification.error({ title: '鎻愮ず', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '鑾峰彇璁㈠崟鍒楄〃澶辫触')
  } finally {
    loading.value = false
  }
}

const refreshLoginData = async () => {
  if (!isLogin.value) {
    orders.value = []
    total.value = 0
    walletBalance.value = '0.00'
    Object.assign(security, {
      realVerified: false,
      fundsVerified: false,
      accountVerified: false
    })
    return
  }

  await Promise.all([refreshSecurity(), refreshWallet(), refreshOrders()])
}

const loadDataPage = (page) => {
  pageNo.value = page
  refreshOrders()
}

const openVerifyDialog = async (direction) => {
  submitDirection.value = direction

  if (!isLogin.value) {
    ElMessage.error(t('common.logintip'))
    router.push('/login')
    return
  }

  await refreshSecurity()

  if (!security.realVerified) {
    ElMessage.error('请先完成实名认证')
    return
  }

  if (!security.fundsVerified) {
    ElMessage.error('请先设置资金密码')
    return
  }

  if (direction === 'buy') {
    if (!tradeForm.buyAmount || tradeForm.buyAmount < 50 || tradeForm.buyAmount > 50000) {
      ElMessage.error('请输入 50 到 50000 之间的买入数量')
      return
    }
  } else {
    if (!security.accountVerified) {
      ElMessage.error('请先完成收款方式设置')
      return
    }

    if (!tradeForm.sellAmount || tradeForm.sellAmount < 2 || tradeForm.sellAmount > 50000) {
      ElMessage.error('请输入 2 到 50000 之间的卖出数量')
      return
    }
  }

  closeVerifyDialog()
  verifyDialogVisible.value = true
}

const sendCode = async () => {
  try {
    const response = await sendCtcCode()
    if (response.code === 0) {
      startSendCodeCooldown()
      ElNotification.success({ title: '提示', message: response.message || '验证码已发送' })
    } else {
      ElNotification.error({ title: '提示', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '发送验证码失败')
  }
}

const setDetailCountdown = () => {
  if (orderTimer.value) {
    clearInterval(orderTimer.value)
    orderTimer.value = null
  }

  countdownSeconds.value = 0

  if (!detailOrder.currentTime) {
    return
  }

  const currentTimestamp = Math.floor(new Date(detailOrder.currentTime).getTime() / 1000)
  const baseTime = detailOrder.status === 1 ? detailOrder.confirmTime : detailOrder.createTime

  if (!baseTime) {
    return
  }

  const baseTimestamp = Math.floor(new Date(baseTime).getTime() / 1000)
  const passed = currentTimestamp - baseTimestamp

  if (passed >= 1800) {
    return
  }

  countdownSeconds.value = 1800 - passed
  orderTimer.value = setInterval(() => {
    countdownSeconds.value -= 1
    if (countdownSeconds.value <= 0) {
      clearInterval(orderTimer.value)
      orderTimer.value = null
    }
  }, 1000)
}

const applyDetailOrder = (order) => {
  Object.keys(detailOrder).forEach((key) => delete detailOrder[key])
  Object.assign(detailOrder, order)
  setDetailCountdown()
  detailDialogVisible.value = true
}

const showOrderDetail = async (orderId) => {
  try {
    const response = await fetchCtcDetail({ oid: orderId })
    if (response.code === 0) {
      applyDetailOrder(response.data)
    } else {
      ElNotification.error({ title: '提示', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '获取订单详情失败')
  }
}

const submitOrder = async () => {
  if (!verifyForm.code) {
    ElMessage.error('请填写短信验证码')
    return
  }

  if (!verifyForm.fundpwd) {
    ElMessage.error('请填写资金密码')
    return
  }

  const payload = buildOrderPayload({
    direction: submitDirection.value,
    quote,
    tradeForm,
    verifyForm
  })

  try {
    const response = await createCtcOrder(payload)
    if (response.code === 0) {
      closeVerifyDialog()
      await refreshOrders()
      applyDetailOrder(response.data)
    } else {
      ElNotification.error({ title: '提示', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '创建订单失败')
  }
}

const cancelOrder = async () => {
  try {
    const response = await cancelCtcOrder({ oid: detailOrder.id })
    if (response.code === 0) {
      await refreshOrders()
      applyDetailOrder(response.data)
    } else {
      ElNotification.error({ title: '提示', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '取消订单失败')
  }
}

const payOrder = async () => {
  try {
    const response = await markCtcOrderPaid({ oid: detailOrder.id })
    if (response.code === 0) {
      await refreshOrders()
      applyDetailOrder(response.data)
    } else {
      ElNotification.error({ title: '提示', message: response.message })
    }
  } catch (error) {
    handleRequestError(error, '标记已付款失败')
  }
}

const cancelOrderClick = () => {
  ElMessageBox.confirm('确定取消这笔订单吗？', '取消订单', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => cancelOrder()).catch(() => {})
}

const payOrderClick = () => {
  ElMessageBox.confirm('请确认您已经完成线下付款。恶意标记已付款会触发账户限制。', '我已线下付款', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => payOrder()).catch(() => {})
}

const init = async () => {
  store.commit('navigate', '/ctc')
  await refreshQuote()
  await refreshLoginData()
}

onMounted(async () => {
  await init()
  timer.value = setInterval(refreshQuote, 30000)
})

onBeforeUnmount(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }

  if (orderTimer.value) {
    clearInterval(orderTimer.value)
  }
})
</script>

<style scoped lang="scss">
.ctc-page {
  padding: 60px 0 50px;
  color: #fff;
}

.ctc-page__banner {
  display: block;
  width: 100%;
}

.ctc-page__container {
  padding: 0 5%;
}

.ctc-page__container > h1 {
  margin-top: -170px;
  padding: 50px 0 20px;
  font-size: 32px;
  letter-spacing: 3px;
  text-align: center;
}

.ctc-page__desc {
  margin: 0;
  text-align: center;
  letter-spacing: 1px;
}

.ctc-page__content {
  margin-top: 55px;
}

.ctc-page__tabs :deep(.el-tabs__nav-wrap::after) {
  background: #27313e;
}

.ctc-page__tabs :deep(.el-tabs__item) {
  color: #9aa4b6;
  font-size: 18px;
}

.ctc-page__tabs :deep(.el-tabs__item.is-active),
.ctc-page__tabs :deep(.el-tabs__item:hover) {
  color: #f0a70a;
}

.ctc-page__tabs :deep(.el-tabs__active-bar) {
  background: #f0a70a;
}

.ctc-page__hero {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 320px;
  gap: 16px;
  min-height: 470px;
}

.ctc-page__grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
}

.trade-card,
.notice-card,
.orders-card {
  border: 1px solid #27313e;
  background: #192330;
  box-shadow: 0 12px 30px rgba(10, 18, 32, 0.22);
}

.trade-card {
  display: flex;
  min-height: 455px;
  flex-direction: column;
  padding: 28px 28px 24px;
}

.trade-card__header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: 24px;
}

.trade-card__eyebrow {
  margin: 0 0 10px;
  color: #9aa4b6;
  font-size: 12px;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.trade-card__title {
  margin: 0;
  font-size: 28px;
  font-weight: 600;
}

.trade-card--buy .trade-card__title,
.trade-card--buy .trade-card__summary strong {
  color: #f15057;
}

.trade-card--sell .trade-card__title,
.trade-card--sell .trade-card__summary strong {
  color: #00b275;
}

.trade-card__badge {
  border: 1px solid #334056;
  padding: 6px 10px;
  color: #9aa4b6;
  font-size: 12px;
}

.trade-card__form {
  display: flex;
  flex: 1;
  flex-direction: column;
}

.trade-card__form :deep(.el-form-item) {
  margin-bottom: 18px;
}

.trade-card__form :deep(.el-form-item__label) {
  color: #9aa4b6;
}

.trade-card__form :deep(.el-input__wrapper),
.trade-card__form :deep(.el-input-number__wrapper),
.trade-card__form :deep(.el-select__wrapper) {
  border-radius: 0;
  box-shadow: 0 0 0 1px #27313e inset;
  background: #111925;
}

.trade-card__form :deep(.el-input__inner),
.trade-card__form :deep(.el-input-number__input),
.trade-card__form :deep(.el-select__selected-item) {
  color: #fff;
}

.trade-card__number,
.trade-card__select {
  width: 100%;
}

.trade-card__link {
  margin: -6px 0 16px;
  color: #4b8df8;
  font-size: 12px;
  text-align: right;
  text-decoration: none;
}

.trade-card__balance {
  margin: -6px 0 16px;
  color: #9aa4b6;
  font-size: 12px;
  text-align: right;
}

.trade-card__summary {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-top: 1px solid #27313e;
  margin-top: auto;
  padding-top: 20px;
}

.trade-card__summary span {
  color: #9aa4b6;
  font-size: 14px;
}

.trade-card__summary strong {
  font-size: 26px;
}

.trade-card__tip {
  margin: 8px 0 18px;
  color: #7f8a9b;
  font-size: 12px;
  line-height: 1.6;
  text-align: right;
}

.trade-card__tip--tight {
  margin-top: -6px;
  margin-bottom: 16px;
}

.trade-card__action {
  width: 100%;
  border: 0;
  border-radius: 0;
  color: #fff;
}

.trade-card__action--buy {
  background: #f15057;
}

.trade-card__action--buy:hover {
  background: #ff7278;
}

.trade-card__action--sell {
  background: #00b275;
}

.trade-card__action--sell:hover {
  background: #01ce88;
}

.notice-card {
  display: flex;
  flex-direction: column;
  min-height: 455px;
}

.notice-card__body {
  flex: 1;
  padding: 25px 30px;
  color: #bcbcbc;
  font-size: 12px;
  line-height: 26px;
}

.notice-card__title {
  margin: 0 0 10px;
  color: #fff;
  font-size: 18px;
  text-align: center;
}

.notice-card__more {
  float: right;
  color: #f0a70a;
  text-decoration: none;
}

.notice-card__footer {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 8px;
  border-top: 1px solid #27313e;
  padding: 12px 16px;
}

.notice-card__button {
  border: 1px solid #0074eb;
  padding: 8px 10px;
  color: #2a93ff;
  font-size: 13px;
  text-align: center;
  text-decoration: none;
}

.orders-card {
  margin-top: 18px;
  padding: 24px;
}

.orders-card__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 18px;
}

.orders-card__eyebrow {
  margin: 0 0 6px;
  color: #9aa4b6;
  font-size: 12px;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.orders-card__title {
  margin: 0;
  font-size: 22px;
}

.orders-card__table :deep(th.el-table__cell) {
  background: #1b2635;
  color: #9aa4b6;
}

.orders-card__table :deep(.el-table__row) {
  background: #192330;
  color: #fff;
}

.orders-card__table :deep(td.el-table__cell) {
  border-bottom-color: #27313e;
  background: #192330;
}

.orders-card__table :deep(.el-table__body tr:hover > td.el-table__cell) {
  background: #223043 !important;
  color: #fff;
}

.orders-card__table :deep(.el-table__body tr.hover-row > td.el-table__cell) {
  background: #223043 !important;
  color: #fff;
}

.orders-card__table :deep(.el-table__inner-wrapper::before) {
  background: #27313e;
}

.orders-card__pagination {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

.direction-buy {
  color: #f15057;
}

.direction-sell {
  color: #00b275;
}

:deep(.ctc-dialog) {
  border: 1px solid #27313e;
  border-radius: 0;
  background: #192330;
  box-shadow: 0 24px 60px rgba(4, 10, 18, 0.55);
}

:deep(.ctc-dialog .el-dialog__header),
:deep(.ctc-dialog__header) {
  margin: 0;
  padding: 20px 24px 16px;
  border-bottom: 1px solid #27313e;
  background: linear-gradient(180deg, #1f2b3b 0%, #192330 100%);
}

:deep(.ctc-dialog .el-dialog__headerbtn .el-dialog__close) {
  color: #9aa4b6;
}

:deep(.ctc-dialog .el-dialog__headerbtn:hover .el-dialog__close) {
  color: #f0a70a;
}

:deep(.ctc-dialog .el-dialog__body) {
  padding: 24px;
  color: #fff;
  background: #192330;
}

:deep(.ctc-dialog .el-dialog__footer) {
  padding: 0 24px 24px;
  background: #192330;
}

.ctc-dialog__title-wrap {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.ctc-dialog__title {
  margin: 0;
  color: #fff;
  font-size: 20px;
  font-weight: 600;
}

.ctc-dialog__ghost-btn,
.ctc-dialog__primary-btn,
.ctc-dialog__secondary-btn,
.ctc-dialog__danger-btn {
  border-radius: 0;
}

.ctc-dialog__ghost-btn {
  border-color: #334056;
  background: transparent;
  color: #d9e1ee;
}

.ctc-dialog__ghost-btn:hover {
  border-color: #4b5a73;
  color: #fff;
  background: #223043;
}

.ctc-dialog__primary-btn {
  min-width: 148px;
  background: #f0a70a;
  border-color: #f0a70a;
  color: #111925;
  font-weight: 600;
}

.ctc-dialog__primary-btn:hover {
  background: #ffbe34;
  border-color: #ffbe34;
  color: #111925;
}

.ctc-dialog__secondary-btn {
  border-color: #334056;
  background: #223043;
  color: #f0a70a;
}

.ctc-dialog__secondary-btn:hover {
  border-color: #4b5a73;
  background: #29384d;
  color: #ffbe34;
}

.ctc-dialog__danger-btn {
  background: #f15057;
  border-color: #f15057;
  color: #fff;
}

.ctc-dialog__danger-btn:hover {
  background: #ff7278;
  border-color: #ff7278;
}

.verify-dialog__hint {
  margin: 0 0 18px;
  padding: 12px 14px;
  border: 1px solid #334056;
  background: #111925;
  color: #9aa4b6;
  font-size: 12px;
  line-height: 1.7;
}

.verify-dialog :deep(.el-form-item__label) {
  color: #9aa4b6;
}

.verify-dialog :deep(.el-input__wrapper) {
  border-radius: 0;
  box-shadow: 0 0 0 1px #27313e inset;
  background: #111925;
}

.verify-dialog :deep(.el-input__wrapper:hover),
.verify-dialog :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #41546d inset;
  background: #111925;
}

.verify-dialog :deep(.el-input__inner) {
  background: transparent;
  color: #fff;
}

.verify-dialog :deep(.el-input__inner::placeholder) {
  color: #7f8a9b;
}

.verify-dialog :deep(.el-input__prefix),
.verify-dialog :deep(.el-input__suffix),
.verify-dialog :deep(.el-input__icon) {
  color: #9aa4b6;
}

.verify-dialog :deep(input:-webkit-autofill),
.verify-dialog :deep(input:-webkit-autofill:hover),
.verify-dialog :deep(input:-webkit-autofill:focus) {
  -webkit-text-fill-color: #fff;
  -webkit-box-shadow: 0 0 0 1000px #111925 inset;
  transition: background-color 9999s ease-in-out 0s;
}

.verify-dialog__row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) auto;
  gap: 12px;
}

.detail-dialog__status {
  margin: 0 0 16px;
  background: #f0a70a;
  padding: 8px 12px;
  color: #111925;
  text-align: center;
}

.detail-dialog__metrics {
  margin-bottom: 16px;
}

.detail-dialog__metrics :deep(.el-col) {
  border: 1px solid #27313e;
  background: #1b2635;
  padding: 14px;
}

.detail-dialog__metric-value {
  margin: 0 0 8px;
  font-size: 20px;
  font-weight: 600;
  text-align: center;
}

.detail-dialog__metric-label {
  margin: 0;
  color: #7f8a9b;
  font-size: 12px;
  text-align: center;
}

.detail-dialog__unit {
  font-size: 13px;
}

.detail-dialog__notice {
  margin-bottom: 16px;
  color: #d9e1ee;
  font-size: 13px;
  line-height: 1.8;
}

.detail-dialog__notice strong {
  color: #f0a70a;
}

.detail-dialog__notice-tag {
  display: inline-block;
  margin-right: 10px;
  padding: 2px 8px;
  border: 1px solid rgba(240, 167, 10, 0.35);
  background: rgba(240, 167, 10, 0.12);
  color: #f0a70a;
  font-size: 12px;
}

.detail-dialog__notice-tag--sell {
  border-color: rgba(0, 178, 117, 0.35);
  background: rgba(0, 178, 117, 0.12);
  color: #00b275;
}

.detail-dialog__countdown {
  margin-left: 12px;
  color: #f15057;
  font-weight: 600;
}

.detail-dialog__panel {
  border: 1px solid #27313e;
  background: #1b2635;
  padding: 18px;
}

.detail-dialog__panel-title {
  margin: 0 0 12px;
  color: #9aa4b6;
  font-size: 12px;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.detail-dialog__row {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  padding: 10px 0;
  border-bottom: 1px solid #27313e;
  color: #d9e1ee;
  font-size: 13px;
}

.detail-dialog__row:last-child {
  border-bottom: 0;
}

.detail-dialog__qr {
  display: block;
  width: 220px;
  max-width: 100%;
  margin: 16px auto 0;
  border: 1px solid #27313e;
}

@media screen and (max-width: 1280px) {
  .ctc-page__hero {
    grid-template-columns: 1fr;
  }

  .notice-card {
    min-height: auto;
  }
}

@media screen and (max-width: 768px) {
  .shoujiIf {
    display: none;
  }

  .ctc-page__content {
    margin-top: 16px;
  }

  .ctc-page__grid {
    grid-template-columns: 1fr;
  }

  .trade-card {
    min-height: auto;
  }

  .orders-card {
    padding: 16px;
  }
}
</style>

<template>
  <div class="content-wrap">
    <div class="order-workbench">
      <p class="breadcrumb">
        <router-link :to="breadcrumbEntryPath" class="breadcrumb-link">{{ breadcrumbEntryLabel }}</router-link>
        <span class="breadcrumb-separator">></span>
        <span class="breadcrumb-current">{{ text.orderDetail }}</span>
      </p>

      <section class="order-summary-bar" :data-status="statusTone">
        <div class="order-summary-bar__status">
          <span class="summary-kicker">{{ text.orderDesk }}</span>
          <h1>{{ statusText }}</h1>
          <p class="summary-subtitle">
            {{ text.orderNo }} {{ msg.orderSn || tradeId }}
            <span v-if="msg.otherSide"> | {{ text.counterparty }} {{ msg.otherSide }}</span>
          </p>
        </div>

        <div class="order-summary-bar__metrics">
          <div
            v-for="item in summaryMetrics"
            :key="item.label"
            class="summary-metric"
            :class="{ 'summary-metric--highlight': item.highlight }"
          >
            <span class="summary-metric__label">{{ item.label }}</span>
            <strong :class="{ 'reserve-time': item.highlight }">{{ item.value }}</strong>
          </div>
        </div>

        <div class="order-summary-bar__actions">
          <el-button v-if="hasAction('confirm-payment')" type="warning" @click="modal1 = true">
            {{ $t('otc.chat.orderstatus_1') }}
          </el-button>
          <el-button v-if="hasAction('cancel-order')" type="danger" @click="modal3 = true">
            {{ $t('otc.chat.orderstatus_4') }}
          </el-button>
          <el-button v-if="hasAction('release-order')" type="warning" @click="modal5 = true">
            {{ $t('otc.chat.orderstatus_3') }}
          </el-button>
          <el-button v-if="hasAction('appeal-order')" type="danger" plain @click="appearOrder">
            {{ tradeType === 1 ? $t('otc.chat.orderstatus_5') : $t('otc.chat.orderstatus_2') }}
          </el-button>
        </div>
      </section>

      <section class="payment-methods-band sidebar-card payment-methods-card" v-show="statusBtn !== 0">
        <div class="card-head">
          <h3>{{ text.paymentMethods }}</h3>
          <span class="card-head-subtitle">{{ text.visibleToBuyer }}</span>
        </div>
        <div class="payment-list payment-list--band">
          <div class="payment-card">
            <div class="payment-card__head">
              <i class="icons bankfor"></i>
              <span>{{ text.bank }}</span>
            </div>
            <template v-if="bankInfo">
              <div class="payment-card__info">
                <div class="payment-card__detail">
                  <strong>{{ payInfo != null ? payInfo.realName : '' }}</strong>
                  <span>{{ bankInfo.branch }} / {{ bankInfo.cardNo }}</span>
                </div>
              </div>
            </template>
            <p v-else class="payment-tip">{{ $t('otc.chat.tip1') }}</p>
          </div>

          <div class="payment-card">
            <div class="payment-card__head">
              <i class="icons alipay"></i>
              <span>{{ $t('otc.chat.zfb') }}</span>
            </div>
            <template v-if="alipay">
              <div class="payment-card__info">
                <div class="payment-card__detail">
                  <strong>{{ alipay.aliNo }}</strong>
                </div>
                <button v-if="alipay.qrCodeUrl" class="qr-link" @click="showQRCode(1)">
                  {{ $t('otc.chat.qrcode') }}
                </button>
              </div>
            </template>
            <p v-else class="payment-tip">{{ $t('otc.chat.tip2') }}</p>
          </div>

          <div class="payment-card">
            <div class="payment-card__head">
              <i class="icons wechat"></i>
              <span>{{ $t('otc.chat.wx') }}</span>
            </div>
            <template v-if="wechatPay">
              <div class="payment-card__info">
                <div class="payment-card__detail">
                  <strong>{{ wechatPay.wechat }}</strong>
                </div>
                <button v-if="wechatPay.qrWeCodeUrl" class="qr-link" @click="showQRCode(2)">
                  {{ $t('otc.chat.qrcode') }}
                </button>
              </div>
            </template>
            <p v-else class="payment-tip">{{ $t('otc.chat.tip3') }}</p>
          </div>
        </div>
      </section>

      <div class="workspace-grid">
        <main class="chat-main">
          <div class="chat-shell" :style="chatShellStyle">
            <Chatline :msg="msg" />
          </div>
        </main>

        <aside ref="sidebarRef" class="sidebar">
          <section class="sidebar-card counterpart-card">
            <div class="card-head">
              <h3>{{ tradeType === 0 ? $t('otc.chat.seller') : $t('otc.chat.buyer') }}</h3>
              <el-popover
                v-if="tradeType === 0 && msg.memberMobile"
                placement="top"
                :content="msg.memberMobile"
              >
                <template #reference>
                  <img src="../../assets/images/icon-tel.png" alt="" class="contact-icon" />
                </template>
              </el-popover>
            </div>
            <router-link :to="otherSidePath" class="counterpart-link">
              {{ msg.otherSide || '--' }}
            </router-link>
            <p class="counterpart-note">
              {{ tradeType === 0
                ? text.buyerGuide
                : text.sellerGuide }}
            </p>
          </section>

          <section class="sidebar-card overview-card">
            <div class="card-head">
              <h3>{{ text.orderSnapshot }}</h3>
              <span class="card-head-subtitle">{{ statusText }}</span>
            </div>
            <div class="overview-grid">
              <div v-for="item in overviewItems" :key="item.label" class="overview-item">
                <span class="overview-item__label">{{ item.label }}</span>
                <strong class="overview-item__value">{{ item.value }}</strong>
              </div>
            </div>
          </section>

        </aside>
      </div>
    </div>

    <el-dialog v-model="modal1" :title="$t('otc.chat.tip')" class="otc-action-dialog">
      <div class="dialog-panel">
        <p class="dialog-warning">{{ $t('otc.chat.msg1') }}</p>
      </div>
      <template #footer>
        <el-button @click="modal1 = false">{{ text.cancel }}</el-button>
        <el-button type="warning" :loading="confirming" @click="handleConfirmPay">{{ text.confirmPayment }}</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="modal3" :title="$t('otc.chat.tip')" class="otc-action-dialog">
      <div class="dialog-panel">
        <p class="dialog-warning">{{ $t('otc.chat.msg3') }}</p>
      </div>
      <template #footer>
        <el-button @click="modal3 = false">{{ text.keepOrder }}</el-button>
        <el-button type="danger" :loading="confirming" @click="handleCancelOrder()">{{ text.confirmCancel }}</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="modal4" :title="$t('otc.chat.tip')" class="otc-action-dialog" @close="onDialog4Close">
      <el-form :model="formItem" :label-width="80">
        <el-form-item :label="$t('otc.chat.comptype')">
          <el-select v-model="formItem.select">
            <el-option value="1" :label="$t('otc.chat.msg4')" />
            <el-option value="2" :label="$t('otc.chat.msg5')" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('otc.chat.compremark')">
          <el-input
            v-model="formItem.remark"
            type="textarea"
            :rows="2"
            :maxlength="200"
            :placeholder="$t('otc.chat.willcomp')"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="onDialog4Close">{{ text.cancel }}</el-button>
        <el-button type="warning" :loading="confirming" @click="handleAppealSubmit">{{ text.submitAppeal }}</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="modal5" :title="$t('otc.chat.tip')" class="otc-action-dialog">
      <div class="dialog-panel">
        <p class="dialog-warning">
          {{ $t('otc.chat.msg6') }}
        </p>
        <el-input
          v-model="fundpwd"
          type="password"
          :placeholder="$t('otc.chat.msg7')"
          class="fundpwd-input"
        />
      </div>
      <template #footer>
        <el-button @click="modal5 = false">{{ text.cancel }}</el-button>
        <el-button type="warning" :loading="confirming" @click="handleRelease">{{ text.releaseAsset }}</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="modal6" class="otc-action-dialog otc-qrcode-dialog">
      <template #header></template>
      <div class="qrcode-dialog">
        <img class="qrcode-image" :src="payCodeUrl" />
      </div>
      <template #footer></template>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, ref, nextTick, onMounted, onBeforeUnmount } from 'vue'
import {
  ElMessage,
  ElButton,
  ElDialog,
  ElForm,
  ElFormItem,
  ElSelect,
  ElOption,
  ElInput,
  ElPopover
} from 'element-plus'
import axios from 'axios'
import io from 'socket.io-client'
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router'
import Chatline from './Chatline.vue'
import { canAutoCancelOrder, getOrderStatusPresentation } from './chat-state'
import { buildOtcCheckUserPath, resolveTradeIdFromRoute } from './route-helpers'

const route = useRoute()
const { locale } = useI18n()

const host = ''
let socket = null

const reserveTime = ref('0:00')
const reserveInteval = ref(null)
const fundpwd = ref('')
const statusBtn = ref(0)
const tradeType = ref(0)
const confirming = ref(false)
const modal1 = ref(false)
const modal3 = ref(false)
const modal4 = ref(false)
const modal5 = ref(false)
const modal6 = ref(false)
const watching = ref(false)
const sidebarRef = ref(null)
const chatShellHeight = ref('auto')
let sidebarResizeObserver = null

const formItem = ref({
  select: '',
  remark: ''
})

const msg = ref({})
const payInfo = ref({})
const bankInfo = ref({})
const alipay = ref({})
const wechatPay = ref({})
const payCodeUrl = ref('')

const tradeId = computed(() => resolveTradeIdFromRoute(route))
const statusPresentation = computed(() => getOrderStatusPresentation({
  status: statusBtn.value,
  tradeType: tradeType.value
}))
const statusText = computed(() => statusPresentation.value.text)
const availableActions = computed(() => new Set(statusPresentation.value.actions))
const statusTone = computed(() => {
  if (statusBtn.value === 1) {
    return 'pending'
  }

  if (statusBtn.value === 2) {
    return tradeType.value === 1 ? 'release' : 'paid'
  }

  if (statusBtn.value === 3) {
    return 'completed'
  }

  if (statusBtn.value === 4) {
    return 'appeal'
  }

  return 'canceled'
})
const isZh = computed(() => String(locale.value || '').toLowerCase().startsWith('zh'))
const text = computed(() => isZh.value ? {
  orderDetail: '订单详情',
  orderDesk: 'OTC 订单工作台',
  orderNo: '订单号',
  counterparty: '交易对象',
  buyerGuide: '等待买家完成付款，再确认收款并放行资产。',
  sellerGuide: '确认收款后再放行数字资产。',
  orderSnapshot: '订单预览',
  paymentMethods: '收款方式',
  visibleToBuyer: '仅买家可见',
  bank: '银行卡',
  cancel: '取消',
  confirmPayment: '确认付款完成',
  keepOrder: '继续保留订单',
  confirmCancel: '确认取消交易',
  submitAppeal: '确认申诉',
  releaseAsset: '确认放行',
  myOrders: '我的订单',
  otcTrade: 'OTC交易',
  total: '订单总额',
  price: '单价',
  amount: '交易数量',
  countdown: '剩余时间',
  timeLimit: '支付时限',
  minutes: '分钟',
  paidNotice: '对方已完成付款，请尽快确认收款并放行。',
  canceledNotice: '对方已取消订单。',
  releasedNotice: '对方已完成放行，请注意查收资产。',
  appealedNotice: '对方已发起申诉。',
  requestFailed: '请求失败',
  fundPasswordRequired: '请输入资金密码',
  autoCanceled: '订单超时已自动取消',
  appealDelay: '付款完成 30 分钟后才允许申诉！',
  qrUnavailable: '当前收款方式未配置付款码'
} : {
  orderDetail: 'Order Detail',
  orderDesk: 'OTC Order Desk',
  orderNo: 'Order No.',
  counterparty: 'Counterparty',
  buyerGuide: 'Wait for the buyer to finish payment, then confirm receipt and release the asset.',
  sellerGuide: 'Confirm receipt before releasing the digital asset.',
  orderSnapshot: 'Order Snapshot',
  paymentMethods: 'Payment Methods',
  visibleToBuyer: 'Visible to buyer only',
  bank: 'Bank',
  cancel: 'Cancel',
  confirmPayment: 'Confirm Payment',
  keepOrder: 'Keep Order',
  confirmCancel: 'Confirm Cancel',
  submitAppeal: 'Submit Appeal',
  releaseAsset: 'Release Asset',
  myOrders: 'My Orders',
  otcTrade: 'OTC Trade',
  total: 'Total',
  price: 'Price',
  amount: 'Amount',
  countdown: 'Countdown',
  timeLimit: 'Time Limit',
  minutes: 'min',
  paidNotice: 'The counterparty marked the order as paid. Please confirm receipt and release promptly.',
  canceledNotice: 'The counterparty canceled this order.',
  releasedNotice: 'The counterparty has released the asset. Please check your account balance.',
  appealedNotice: 'The counterparty submitted an appeal.',
  requestFailed: 'Request failed',
  fundPasswordRequired: 'Please enter your fund password',
  autoCanceled: 'Order timed out and was canceled automatically.',
  appealDelay: 'Appeal is available 30 minutes after payment is completed.',
  qrUnavailable: 'No payment QR code is configured for this method.'
})
const otherSidePath = computed(() => buildOtcCheckUserPath(msg.value.otherSide || ''))
const breadcrumbEntryLabel = computed(() => route.query.source === 'order' ? text.value.myOrders : text.value.otcTrade)
const breadcrumbEntryPath = computed(() => {
  if (route.query.source === 'order') {
    return '/uc/order'
  }

  const unit = String(msg.value.unit || 'USDT').toLowerCase()
  return `/otc/trade/${unit}`
})
const summaryMetrics = computed(() => {
  const metrics = [
    { label: text.value.total, value: msg.value.money ? `${msg.value.money} CNY` : '--' },
    { label: text.value.price, value: msg.value.price ? `${msg.value.price} CNY` : '--' },
    { label: text.value.amount, value: msg.value.amount ? `${msg.value.amount} ${msg.value.unit || ''}` : '--' }
  ]

  if (statusBtn.value === 1) {
    metrics.push({
      label: text.value.countdown,
      value: reserveTime.value,
      highlight: true
    })
  }

  return metrics
})
const overviewItems = computed(() => ([
  { label: text.value.orderNo, value: msg.value.orderSn || tradeId.value || '--' },
  { label: text.value.counterparty, value: msg.value.otherSide || '--' },
  { label: text.value.price, value: msg.value.price ? `${msg.value.price} CNY` : '--' },
  { label: text.value.amount, value: msg.value.amount ? `${msg.value.amount} ${msg.value.unit || ''}` : '--' },
  { label: text.value.total, value: msg.value.money ? `${msg.value.money} CNY` : '--' },
  { label: text.value.timeLimit, value: msg.value.timeLimit ? `${msg.value.timeLimit} ${text.value.minutes}` : '--' }
]))
const chatShellStyle = computed(() => ({
  minHeight: chatShellHeight.value,
  height: chatShellHeight.value
}))

const orderRequestConfig = {
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    'x-auth-token': localStorage.getItem('TOKEN')
  }
}

const hasAction = (action) => availableActions.value.has(action)

const parseOrderTime = (value) => {
  const rawValue = String(value || '').trim()
  if (!rawValue) {
    return Number.NaN
  }

  const directTime = new Date(rawValue).getTime()
  if (Number.isFinite(directTime)) {
    return directTime
  }

  const legacyTime = new Date(rawValue.replace(/-/g, '/')).getTime()
  if (Number.isFinite(legacyTime)) {
    return legacyTime
  }

  return Number.NaN
}

const resolveQrPreviewUrl = (value) => {
  const rawValue = String(value || '').trim()
  if (!rawValue) {
    return ''
  }

  if (/^(https?:|data:|blob:)/i.test(rawValue)) {
    return rawValue
  }

  if (rawValue.startsWith('/')) {
    return rawValue
  }

  return `/${rawValue.replace(/^\.?\//, '')}`
}

const syncChatShellHeight = () => {
  if (typeof window === 'undefined' || window.innerWidth <= 1200) {
    chatShellHeight.value = 'auto'
    return
  }

  const sidebarEl = sidebarRef.value
  if (!sidebarEl) {
    chatShellHeight.value = 'auto'
    return
  }

  chatShellHeight.value = `${Math.ceil(sidebarEl.offsetHeight)}px`
}

const queueSyncChatShellHeight = () => {
  nextTick(() => {
    syncChatShellHeight()
  })
}

const observeSidebarHeight = () => {
  if (typeof ResizeObserver === 'undefined') {
    queueSyncChatShellHeight()
    return
  }

  sidebarResizeObserver?.disconnect()
  nextTick(() => {
    if (!sidebarRef.value) {
      return
    }

    sidebarResizeObserver = new ResizeObserver(() => {
      syncChatShellHeight()
    })
    sidebarResizeObserver.observe(sidebarRef.value)
    syncChatShellHeight()
  })
}

const closeReserveCountdown = () => {
  if (reserveInteval.value) {
    clearInterval(reserveInteval.value)
    reserveInteval.value = null
  }
}

const applyOrderStatus = (nextStatus) => {
  statusBtn.value = Number(nextStatus)
  msg.value = {
    ...msg.value,
    status: Number(nextStatus)
  }

  closeReserveCountdown()
  if (statusBtn.value === 1) {
    setReserveTime()
  }

  queueSyncChatShellHeight()
}

const sendOrderStatusNotice = (nextStatus) => {
  const contentMap = {
    2: text.value.paidNotice,
    0: text.value.canceledNotice,
    3: text.value.releasedNotice,
    4: text.value.appealedNotice
  }

  const content = contentMap[nextStatus]
  if (!content) {
    return
  }

  const jsonParam = {
    uidTo: msg.value.hisId,
    uidFrom: msg.value.myId,
    orderId: tradeId.value,
    content,
    messageType: 0,
    status: nextStatus
  }
  socket?.emit('/app/message/chat', JSON.stringify(jsonParam))
}

const submitOrderAction = async ({ path, payload, nextStatus, successMessage, onSuccess }) => {
  confirming.value = true

  try {
    const response = await axios.post(`${host}${path}`, payload, orderRequestConfig)
    const resp = response.data
    if (resp.code !== 0) {
      ElMessage.error(resp.message)
      return false
    }

    if (nextStatus !== undefined) {
      applyOrderStatus(nextStatus)
      sendOrderStatusNotice(nextStatus)
    }

    if (typeof onSuccess === 'function') {
      onSuccess(resp)
    }

    ElMessage.success(successMessage || resp.message)
    await getDetail()
    return true
  } catch {
    ElMessage.error(text.value.requestFailed)
    return false
  } finally {
    confirming.value = false
  }
}

const handleConfirmPay = async () => {
  await submitOrderAction({
    path: '/otc/order/pay',
    payload: { orderSn: tradeId.value },
    nextStatus: 2,
    onSuccess: () => {
      modal1.value = false
    }
  })
}

const handleCancelOrder = async ({ auto = false } = {}) => {
  await submitOrderAction({
    path: '/otc/order/cancel',
    payload: { orderSn: tradeId.value },
    nextStatus: 0,
    successMessage: auto ? text.value.autoCanceled : undefined,
    onSuccess: () => {
      modal3.value = false
    }
  })
}

const handleAppealSubmit = async () => {
  await submitOrderAction({
    path: '/otc/order/appeal',
    payload: {
      orderSn: tradeId.value,
      remark: formItem.value.remark
    },
    nextStatus: 4,
    onSuccess: () => {
      onDialog4Close()
    }
  })
}

const handleRelease = async () => {
  if (fundpwd.value === '') {
    ElMessage.error(text.value.fundPasswordRequired)
    return
  }

  await submitOrderAction({
    path: '/otc/order/release',
    payload: {
      orderSn: tradeId.value,
      jyPassword: fundpwd.value
    },
    nextStatus: 3,
    onSuccess: () => {
      modal5.value = false
      fundpwd.value = ''
    }
  })
}

const onDialog4Close = () => {
  modal4.value = false
  formItem.value = {
    select: '',
    remark: ''
  }
}

const appearOrder = () => {
  const payTimeValue = new Date(msg.value.payTime || '').getTime()
  if (Number.isFinite(payTimeValue) && parseInt((Date.now() - payTimeValue) / 1000, 10) < 1800) {
    ElMessage.info(text.value.appealDelay)
    return
  }

  modal4.value = true
}

const showQRCode = (type) => {
  const previewUrl = resolveQrPreviewUrl(type === 1 ? alipay.value?.qrCodeUrl : wechatPay.value?.qrWeCodeUrl)
  if (!previewUrl) {
    ElMessage.info(text.value.qrUnavailable)
    return
  }

  payCodeUrl.value = previewUrl
  modal6.value = true
}

const resetStatus = () => {
  closeReserveCountdown()
  if (canAutoCancelOrder({ status: statusBtn.value, tradeType: tradeType.value })) {
    handleCancelOrder({ auto: true })
  }
}

const getReserveTime = () => {
  const createTime = parseOrderTime(msg.value.createTime)
  const timeLimit = parseInt(msg.value.timeLimit, 10)

  if (!Number.isFinite(createTime) || !Number.isFinite(timeLimit)) {
    reserveTime.value = '0:00'
    return
  }

  const throughSeconds = parseInt((Date.now() - createTime) / 1000, 10)
  const reserveSeconds = timeLimit * 60 - throughSeconds
  const safeReserveSeconds = Math.max(reserveSeconds, 0)
  const minutes = parseInt(safeReserveSeconds / 60, 10)
  const seconds = parseInt(safeReserveSeconds % 60, 10)
  reserveTime.value = `${minutes}:${seconds >= 10 ? seconds : `0${seconds}`}`

  if (reserveSeconds <= 0) {
    resetStatus()
  }
}

const setReserveTime = () => {
  closeReserveCountdown()
  getReserveTime()
  reserveInteval.value = setInterval(getReserveTime, 1000)
}

const watchOrderStutusNotice = () => {
  socket.on(`/user/${msg.value.myId}/order-notice/${tradeId.value}`, (response) => {
    const orderNotice = JSON.parse(response)
    if (orderNotice.content) {
      ElMessage.success(orderNotice.content)
    }
    if (orderNotice.status !== undefined) {
      applyOrderStatus(orderNotice.status)
    }
  })
}

async function getDetail() {
  try {
    const response = await axios.post(`${host}/otc/order/detail`, {
      orderSn: tradeId.value
    }, orderRequestConfig)

    const resp = response.data
    if (resp.code !== 0) {
      ElMessage.error(resp.message)
      return
    }

    msg.value = resp.data
    payInfo.value = resp.data.payInfo

    if (payInfo.value === null) {
      bankInfo.value = null
      alipay.value = null
      wechatPay.value = null
    } else {
      bankInfo.value = resp.data.payInfo.bankInfo
      alipay.value = resp.data.payInfo.alipay
      wechatPay.value = resp.data.payInfo.wechatPay
    }

    tradeType.value = resp.data.type
    applyOrderStatus(resp.data.status)

    if (!watching.value) {
      watchOrderStutusNotice()
      watching.value = true
    }

    observeSidebarHeight()
  } catch {
    ElMessage.error(text.value.requestFailed)
  }
}

const initSocket = () => {
  socket = io(window.location.origin, {
    autoConnect: false,
    transports: ['websocket']
  })
  socket.connect()
}

onMounted(() => {
  initSocket()
  getDetail()
  window.addEventListener('resize', syncChatShellHeight)
})

onBeforeUnmount(() => {
  closeReserveCountdown()
  sidebarResizeObserver?.disconnect()
  window.removeEventListener('resize', syncChatShellHeight)
  socket?.disconnect()
})
</script>

<style scoped lang="scss">
:root {
  --color-primary: #f0ac19;
  --color-primary-hover: #f0bc47;
  --color-primary-light: rgba(240, 172, 25, 0.1);
  --color-danger: #f56c6c;
  --text-primary: #ffffff;
  --text-secondary: #aab6c7;
  --text-muted: #6b7a8f;
  --bg-secondary: #1a1f2e;
  --bg-card: linear-gradient(180deg, rgba(22, 30, 44, 0.98), rgba(12, 18, 29, 0.98));
  --bg-card-subtle: rgba(30, 41, 59, 0.6);
  --border-default: rgba(255, 255, 255, 0.08);
  --border-light: rgba(255, 255, 255, 0.06);
  --border-strong: rgba(255, 255, 255, 0.12);
  --shadow-md: 0 4px 12px rgba(0, 0, 0, 0.24);
  --shadow-lg: 0 8px 24px rgba(0, 0, 0, 0.28);
  --radius-sm: 8px;
  --radius-md: 12px;
  --radius-lg: 16px;
  --radius-xl: 20px;
  --transition-fast: 150ms ease;
  --transition-base: 200ms ease;
}

.content-wrap {
  min-height: 100vh;
  padding: 80px 0 36px;
  background:
    radial-gradient(circle at top, rgba(240, 172, 25, 0.08), transparent 26%),
    linear-gradient(180deg, #0b1118 0%, #101723 100%);
}

.order-workbench {
  width: min(1400px, calc(100vw - 48px));
  margin: 0 auto;
}

.breadcrumb {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  font-size: 14px;
}

.breadcrumb-link {
  color: var(--color-primary);
  text-decoration: none;
  transition: opacity var(--transition-fast);
}

.breadcrumb-link:hover {
  opacity: 0.8;
}

.breadcrumb-separator {
  margin: 0 10px;
  color: var(--text-muted);
}

.breadcrumb-current {
  color: var(--text-secondary);
}

.order-summary-bar {
  position: relative;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 18px;
  padding: 18px 22px;
  margin-bottom: 18px;
  border: 1px solid var(--border-strong);
  border-radius: 18px;
  background:
    radial-gradient(ellipse at top left, rgba(240, 172, 25, 0.1), transparent 38%),
    linear-gradient(135deg, rgba(24, 34, 49, 0.98), rgba(12, 18, 29, 0.98));
  box-shadow: var(--shadow-lg);
  overflow: hidden;
}

.order-summary-bar::before {
  content: '';
  position: absolute;
  inset: 0 auto auto 0;
  width: 220px;
  height: 220px;
  background: radial-gradient(circle, rgba(240, 172, 25, 0.16), transparent 72%);
  pointer-events: none;
}

.order-summary-bar[data-status='paid'] {
  background:
    radial-gradient(ellipse at top left, rgba(34, 197, 94, 0.14), transparent 42%),
    linear-gradient(135deg, rgba(24, 34, 49, 0.98), rgba(12, 18, 29, 0.98));
}

.order-summary-bar[data-status='release'] {
  background:
    radial-gradient(ellipse at top left, rgba(59, 130, 246, 0.16), transparent 42%),
    linear-gradient(135deg, rgba(24, 34, 49, 0.98), rgba(12, 18, 29, 0.98));
}

.order-summary-bar[data-status='completed'] {
  background:
    radial-gradient(ellipse at top left, rgba(16, 185, 129, 0.14), transparent 42%),
    linear-gradient(135deg, rgba(24, 34, 49, 0.98), rgba(12, 18, 29, 0.98));
}

.order-summary-bar[data-status='appeal'],
.order-summary-bar[data-status='canceled'] {
  background:
    radial-gradient(ellipse at top left, rgba(239, 68, 68, 0.14), transparent 42%),
    linear-gradient(135deg, rgba(24, 34, 49, 0.98), rgba(12, 18, 29, 0.98));
}

.order-summary-bar__status,
.order-summary-bar__metrics,
.order-summary-bar__actions {
  position: relative;
  z-index: 1;
}

.order-summary-bar__status {
  flex: 0.95 1 320px;
  min-width: 0;
}

.order-summary-bar__status h1 {
  margin: 8px 0 6px;
  color: var(--text-primary);
  font-size: 24px;
  font-weight: 700;
  line-height: 1.2;
  letter-spacing: -0.02em;
}

.summary-kicker {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 5px 10px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.06);
  color: #f6cd70;
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.summary-subtitle {
  margin: 0;
  color: var(--text-secondary);
  font-size: 13px;
  line-height: 1.5;
}

.order-summary-bar__metrics {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(132px, 1fr));
  gap: 10px;
  flex: 1.35 1 520px;
}

.summary-metric {
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-width: 0;
  gap: 4px;
  min-height: 68px;
  padding: 10px 12px;
  border-radius: 14px;
  background: linear-gradient(135deg, rgba(37, 49, 69, 0.58), rgba(19, 28, 42, 0.42));
  border: 1px solid var(--border-light);
  backdrop-filter: blur(6px);
  transition:
    transform var(--transition-fast),
    border-color var(--transition-base),
    background var(--transition-base);
}

.summary-metric:hover {
  transform: translateY(-2px);
  border-color: var(--border-default);
  background: linear-gradient(135deg, rgba(48, 62, 85, 0.76), rgba(23, 34, 49, 0.62));
}

.summary-metric--highlight {
  border-color: rgba(240, 172, 25, 0.26);
  background: linear-gradient(135deg, rgba(77, 46, 10, 0.56), rgba(35, 25, 12, 0.54));
}

.summary-metric__label {
  color: var(--text-muted);
  font-size: 13px;
  font-weight: 600;
  line-height: 1.2;
  word-break: break-word;
}

.summary-metric strong {
  color: var(--text-primary);
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 0;
  line-height: 1.35;
  word-break: break-word;
}

.order-summary-bar__actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-end;
  gap: 10px;
  flex: 0 1 250px;
}

.order-summary-bar__actions :deep(.el-button) {
  min-width: 116px;
  height: 40px;
  font-weight: 600;
  border-radius: 12px;
  transition: all var(--transition-fast);
}

.order-summary-bar__actions :deep(.el-button:hover) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(240, 172, 25, 0.3);
}

.reserve-time {
  color: #ff8a65 !important;
  font-weight: 600;
}

.payment-methods-band {
  display: block;
  width: 100%;
  margin: 0 0 24px;
  box-sizing: border-box;
}

.workspace-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 320px;
  gap: 24px;
  align-items: start;
}

.chat-main {
  min-width: 0;
  align-self: start;
}

.chat-shell {
  padding: 18px;
  border: 1px solid var(--border-default);
  border-radius: 24px;
  background: linear-gradient(180deg, rgba(16, 23, 35, 0.96), rgba(13, 19, 30, 0.98));
  box-shadow: var(--shadow-lg);
  box-sizing: border-box;
}

.sidebar {
  display: flex;
  flex-direction: column;
  gap: 18px;
  min-width: 0;
  align-self: start;
  position: relative;
  top: 0;
}

.sidebar-card {
  padding: 20px 20px 18px;
  border: 1px solid var(--border-default);
  border-radius: var(--radius-lg);
  background: var(--bg-card);
  box-shadow: var(--shadow-md);
  color: var(--text-primary);
  transition:
    box-shadow var(--transition-base),
    border-color var(--transition-base);
}

.sidebar-card:hover {
  box-shadow: var(--shadow-lg);
  border-color: var(--border-strong);
}

.card-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 18px;
  padding-bottom: 14px;
  border-bottom: 1px solid var(--border-light);
}

.card-head h3 {
  margin: 0;
  color: var(--text-primary);
  font-size: 16px;
  font-weight: 600;
  letter-spacing: -0.01em;
}

.card-head-subtitle {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 500;
}

.counterpart-link {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
  color: var(--color-primary);
  font-size: 20px;
  font-weight: 700;
  text-decoration: none;
  transition: opacity var(--transition-fast);
}

.counterpart-link:hover {
  opacity: 0.8;
}

.counterpart-note {
  margin: 0;
  color: var(--text-secondary);
  font-size: 13px;
  line-height: 1.6;
}

.contact-icon {
  width: 20px;
  height: 20px;
  cursor: pointer;
  transition: transform var(--transition-fast);
}

.contact-icon:hover {
  transform: scale(1.1);
}

.overview-grid {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.overview-item {
  display: grid;
  grid-template-columns: minmax(0, 92px) minmax(0, 1fr);
  gap: 12px;
  align-items: center;
  padding: 12px 14px;
  border-radius: 14px;
  background: linear-gradient(135deg, var(--bg-card-subtle), rgba(21, 31, 45, 0.44));
  border: 1px solid var(--border-light);
  transition:
    background var(--transition-base),
    border-color var(--transition-base),
    transform var(--transition-fast);
}

.overview-item:hover {
  background: rgba(51, 65, 85, 0.5);
  border-color: var(--border-default);
  transform: translateY(-1px);
}

.overview-item__label {
  color: var(--text-muted);
  font-size: 15px;
  font-weight: 600;
  line-height: 1.4;
}

.overview-item__value {
  color: var(--text-primary);
  font-size: 15px;
  font-weight: 600;
  text-align: right;
  word-break: break-word;
}

.payment-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.payment-list--band {
  display: grid;
  grid-template-columns: repeat(3, minmax(220px, 1fr));
  gap: 14px;
}

.payment-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 16px;
  border-radius: var(--radius-md);
  background: linear-gradient(135deg, rgba(30, 41, 59, 0.6), rgba(30, 41, 59, 0.4));
  border: 1px solid var(--border-light);
  transition:
    border-color var(--transition-base),
    background var(--transition-base),
    transform var(--transition-fast);
}

.payment-card:hover {
  border-color: var(--border-default);
  background: rgba(51, 65, 85, 0.5);
  transform: translateX(4px);
}

.payment-card__head {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 70px;
  flex-shrink: 0;
  color: var(--text-primary);
  font-weight: 600;
  font-size: 14px;
}

.payment-card__info {
  display: flex;
  align-items: center;
  gap: 16px;
  flex: 1;
  min-width: 0;
}

.payment-card__detail {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
  min-width: 0;
}

.payment-card strong {
  color: var(--text-primary);
  font-size: 14px;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.payment-card span,
.payment-tip {
  color: var(--text-secondary);
  font-size: 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.payment-tip {
  margin: 0;
  line-height: 1.6;
  word-break: break-all;
}

.qr-link {
  flex-shrink: 0;
  width: fit-content;
  padding: 8px 16px;
  border: 1px solid var(--color-primary);
  border-radius: var(--radius-sm);
  background: transparent;
  color: var(--color-primary);
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.qr-link:hover {
  background: var(--color-primary-light);
  border-color: var(--color-primary-hover);
  color: var(--color-primary-hover);
  transform: translateY(-1px);
}

.dialog-warning {
  margin: 0;
  color: var(--color-danger);
  font-weight: 600;
  font-size: 15px;
  line-height: 1.7;
}

.fundpwd-input {
  margin-top: 14px;
}

.dialog-panel {
  padding: 18px 20px;
  border: 1px solid rgba(240, 172, 25, 0.12);
  border-radius: 16px;
  background: linear-gradient(180deg, rgba(14, 20, 31, 0.96), rgba(10, 15, 24, 0.98));
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.03);
}

:deep(.el-overlay) {
  background: rgba(2, 6, 14, 0.72) !important;
}

:deep(.otc-action-dialog) {
  background-color: #121927 !important;
  background-image: none !important;
  border: 1px solid var(--border-strong);
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: 0 24px 60px rgba(0, 0, 0, 0.48);
}

:deep(.otc-action-dialog .el-dialog__title) {
  color: var(--text-primary);
}

:deep(.otc-action-dialog .el-dialog__header) {
  margin: 0;
  padding: 18px 22px 8px;
  background-color: #121927 !important;
}

:deep(.otc-action-dialog .el-dialog__body) {
  padding: 12px 22px 18px;
  background-color: #121927 !important;
  color: var(--text-secondary);
}

:deep(.otc-action-dialog .el-dialog__footer) {
  padding: 0 22px 22px;
  background-color: #121927 !important;
}

:deep(.otc-action-dialog .el-dialog__headerbtn),
:deep(.otc-action-dialog .el-dialog__close) {
  color: var(--text-secondary) !important;
}

:deep(.otc-action-dialog .el-dialog__headerbtn:hover),
:deep(.otc-action-dialog .el-dialog__close:hover) {
  color: var(--text-primary) !important;
}

.qrcode-dialog {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.qrcode-image {
  width: 220px;
  height: 220px;
  object-fit: contain;
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-lg);
}

.icons {
  display: inline-block;
  width: 20px;
  height: 20px;
  vertical-align: middle;
  background-size: 100% 100%;
  background-repeat: no-repeat;
}

.icons.alipay {
  background-image: url(../../assets/img/alipay.png);
}

.icons.wechat {
  background-image: url(../../assets/img/wechat.png);
}

.icons.bankfor {
  background-image: url(../../assets/img/bankcard.png);
}

@media (max-width: 1200px) {
  .order-summary-bar {
    flex-wrap: wrap;
    align-items: flex-start;
  }

  .order-summary-bar__metrics {
    order: 3;
    width: 100%;
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .order-summary-bar__actions {
    flex: 1 1 100%;
    justify-content: flex-start;
  }

  .workspace-grid {
    grid-template-columns: 1fr;
  }

  .sidebar {
    position: static;
  }

  .payment-list--band {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .content-wrap {
    padding-top: 72px;
  }

  .order-workbench {
    width: calc(100vw - 24px);
    padding: 0 8px;
  }

  .order-summary-bar {
    padding: 16px;
    gap: 14px;
  }

  .order-summary-bar__status h1 {
    font-size: 22px;
  }

  .order-summary-bar__metrics {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .order-summary-bar__actions :deep(.el-button) {
    min-width: 0;
    flex: 1 1 calc(50% - 6px);
  }

  .chat-shell {
    min-height: 0;
    padding: 16px;
  }

  .overview-item {
    grid-template-columns: 1fr;
    gap: 8px;
  }

  .overview-item__value {
    text-align: left;
  }

  .payment-card {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }

  .payment-card__head {
    width: auto;
  }

  .payment-card__info {
    width: 100%;
  }

  .qr-link {
    width: 100%;
    text-align: center;
  }
}

@media (max-width: 480px) {
  .order-summary-bar__metrics {
    grid-template-columns: 1fr;
  }

  .order-summary-bar__actions {
    width: 100%;
  }

  .order-summary-bar__actions :deep(.el-button) {
    width: 100%;
    flex-basis: 100%;
  }
}
</style>

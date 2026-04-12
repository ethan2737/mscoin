<template>
  <div class="content-wrap">
    <div class="container chat-in-box" id="List">
      <p class="breadcrumb">
        <router-link to="/uc/order" class="breadcrumb-link">{{ $t('otc.myorder') }}</router-link>
        <span class="breadcrumb-separator">></span>
        <span class="breadcrumb-current">订单详情</span>
      </p>
      <div class="row chat-in">
        <div class="col-left">
          <div class="leftmenu left-box chat-right">
            <div class="chat-right-in">
              <h6 class="h6-flex">
                <span v-if="tradeType === 0">{{ $t('otc.chat.seller') }}:</span>
                <span v-else>{{ $t('otc.chat.buyer') }}:</span>
                <router-link :to="otherSidePath">
                  {{ msg.otherSide }}
                </router-link>
              </h6>
              <el-popover v-if="tradeType === 0 && msg.memberMobile" placement="top" :content="msg.memberMobile">
                <template #reference>
                  <img src="../../assets/images/icon-tel.png" alt="" class="pop-tel-img" />
                </template>
              </el-popover>
              <h6>
                <span>{{ $t('otc.chat.exchangeamount') }}:</span>
                <span>{{ msg.money }} CNY</span>
              </h6>
              <div class="mt20" v-if="tradeType === 0">
                <h5>{{ $t('otc.chat.operatetip') }}:</h5>
                <div>
                  <p>1. {{ $t('otc.chat.operatetip_1') }} "<em>{{ $t('otc.chat.finishpayment') }}</em>" {{ $t('otc.chat.operatetip_1_1') }}</p>
                  <p>2. {{ $t('otc.chat.operatetip_1_2') }}</p>
                </div>
                <span><b>{{ $t('otc.chat.note') }}:</b></span><br>
                <span>{{ $t('otc.chat.notetip') }}</span><br>
              </div>
              <div class="mt20" v-else>
                <h5>{{ $t('otc.chat.operatetip') }}:</h5>
                <div>
                  <p>1. {{ $t('otc.chat.operatetip_2_1') }} "<em>{{ $t('otc.chat.confirmrelease') }}</em>" {{ $t('otc.chat.paydigital') }}</p>
                  <p>2. {{ $t('otc.chat.operatetip_2_2') }}</p>
                  <p>3. {{ $t('otc.chat.operatetip_2_3') }}</p>
                </div>
                <span><b>{{ $t('otc.chat.note') }}:</b></span><br>
                <span>{{ $t('otc.chat.notetip') }}</span><br>
              </div>
              <div class="bottom-btn">
                <div class="status-actions">
                  <h6 class="status-label">
                    {{ $t('otc.chat.orderstatus') }}:
                    <span>{{ statusText }}</span>
                  </h6>
                  <div v-if="hasAction('confirm-payment') || hasAction('cancel-order')">
                    <el-button v-if="hasAction('confirm-payment')" type="warning" @click="modal1 = true">{{ $t('otc.chat.orderstatus_1') }}</el-button>
                    <el-button v-if="hasAction('cancel-order')" type="danger" @click="modal3 = true">{{ $t('otc.chat.orderstatus_4') }}</el-button>
                  </div>
                  <div v-else-if="hasAction('release-order') || hasAction('appeal-order')">
                    <el-button v-if="hasAction('release-order')" type="warning" @click="modal5 = true">{{ $t('otc.chat.orderstatus_3') }}</el-button>
                    <el-button v-if="hasAction('appeal-order')" type="danger" plain @click="appearOrder">
                      {{ tradeType === 1 ? $t('otc.chat.orderstatus_5') : $t('otc.chat.orderstatus_2') }}
                    </el-button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col-right">
          <div class="rightbox">
            <div class="row chat-top chat-summary">
              <div class="col order-time">
                <h5>{{ statusText }}</h5>
                <div v-show="statusBtn === 1" class="reserve-time">{{ reserveTime }}</div>
              </div>
              <div class="col order-info">
                <h5>
                  <label class="order-name">{{ $t('otc.chat.order') }}</label>
                  <span>{{ msg.orderSn }}</span>
                </h5>
                <p>
                  {{ $t('otc.chat.and') }}
                  <router-link :to="otherSidePath">
                    {{ msg.otherSide }}
                  </router-link>
                  {{ $t('otc.chat.transition') }}
                </p>
              </div>
              <div class="col order-info">
                <h5>{{ msg.price }}</h5>
                <span>{{ $t('otc.chat.transprice') }}(CNY)</span>
              </div>
              <div class="col order-info">
                <h5>{{ msg.amount }}</h5>
                <span>{{ $t('otc.chat.transnum') }}({{ msg.unit }})</span>
              </div>
              <div class="col order-info">
                <h5>{{ msg.money }}</h5>
                <span>{{ $t('otc.chat.transmoney') }}(CNY)</span>
              </div>
            </div>
            <div class="row chat-top chat-payment" v-show="statusBtn !== 0">
              <div class="col order-info" v-if="bankInfo">
                <i class="icons bankfor"></i>
                <span>{{ payInfo != null ? payInfo.realName : '' }}</span>
                <p>{{ bankInfo.branch }}</p>
                <p>{{ bankInfo.cardNo }}</p>
              </div>
              <div class="col order-info" v-else>
                <i class="icons bankfor"></i>
                <pre></pre>
                <p class="payment-tip">{{ $t('otc.chat.tip1') }}</p>
              </div>
              <div class="col order-info" v-if="alipay">
                <i class="icons alipay"></i>
                <span>{{ $t('otc.chat.zfb') }}</span>
                <pre></pre>
                <p>{{ alipay.aliNo }}</p>
                <p v-if="alipay.qrCodeUrl">
                  <a @click="showQRCode(1)">{{ $t('otc.chat.qrcode') }}</a>
                </p>
              </div>
              <div class="col order-info" v-else>
                <i class="icons alipay"></i>
                <pre></pre>
                <p class="payment-tip">{{ $t('otc.chat.tip2') }}</p>
              </div>
              <div class="col order-info" v-if="wechatPay">
                <i class="icons wechat"></i>
                <span>{{ $t('otc.chat.wx') }}</span>
                <pre></pre>
                <p>{{ wechatPay.wechat }}</p>
                <p v-if="wechatPay.qrWeCodeUrl">
                  <a @click="showQRCode(2)">{{ $t('otc.chat.qrcode') }}</a>
                </p>
              </div>
              <div class="col order-info" v-else>
                <i class="icons wechat"></i>
                <pre></pre>
                <p class="payment-tip">{{ $t('otc.chat.tip3') }}</p>
              </div>
            </div>
            <Chatline :msg="msg" />
          </div>
        </div>
      </div>
    </div>
    <el-dialog v-model="modal1" :title="$t('otc.chat.tip')">
      <p class="dialog-warning">{{ $t('otc.chat.msg1') }}</p>
      <template #footer>
        <el-button @click="modal1 = false">取消</el-button>
        <el-button type="warning" :loading="confirming" @click="handleConfirmPay">确认付款完成</el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="modal3" :title="$t('otc.chat.tip')">
      <p class="dialog-warning">{{ $t('otc.chat.msg3') }}</p>
      <template #footer>
        <el-button @click="modal3 = false">继续保留订单</el-button>
        <el-button type="danger" :loading="confirming" @click="handleCancelOrder()">确认取消交易</el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="modal4" :title="$t('otc.chat.tip')" @close="onDialog4Close">
      <el-form :model="formItem" :label-width="80">
        <el-form-item :label="$t('otc.chat.comptype')">
          <el-select v-model="formItem.select">
            <el-option value="1" :label="$t('otc.chat.msg4')" />
            <el-option value="2" :label="$t('otc.chat.msg5')" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('otc.chat.compremark')">
          <el-input v-model="formItem.remark" type="textarea" :rows="2" :maxlength="200" :placeholder="$t('otc.chat.willcomp')" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="onDialog4Close">取消</el-button>
        <el-button type="warning" :loading="confirming" @click="handleAppealSubmit">确认申诉</el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="modal5" :title="$t('otc.chat.tip')">
      <p class="dialog-warning">
        {{ $t('otc.chat.msg6') }}<br />
        <el-input type="password" v-model="fundpwd" :placeholder="$t('otc.chat.msg7')" class="fundpwd-input" />
      </p>
      <template #footer>
        <el-button @click="modal5 = false">取消</el-button>
        <el-button type="warning" :loading="confirming" @click="handleRelease">确认放行</el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="modal6">
      <template #header></template>
      <div class="qrcode-dialog">
        <img class="qrcode-image" :src="payCodeUrl" />
      </div>
      <template #footer></template>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, ref, onMounted, onBeforeUnmount } from 'vue'
import { ElMessage, ElButton, ElDialog, ElForm, ElFormItem, ElSelect, ElOption, ElInput, ElPopover } from 'element-plus'
import axios from 'axios'
import io from 'socket.io-client'
import { useRoute } from 'vue-router'
import Chatline from './Chatline.vue'
import { canAutoCancelOrder, getOrderStatusPresentation } from './chat-state'
import { buildOtcCheckUserPath, resolveTradeIdFromRoute } from './route-helpers'

const route = useRoute()

const host = ''
let socket = null

const watching = ref(false)
const reserveTime = ref('60')
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
const otherSidePath = computed(() => buildOtcCheckUserPath(msg.value.otherSide || ''))

const orderRequestConfig = {
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    'x-auth-token': localStorage.getItem('TOKEN')
  }
}

const hasAction = (action) => availableActions.value.has(action)

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
}

const sendOrderStatusNotice = (nextStatus) => {
  const contentMap = {
    2: '对方已完成付款，请尽快确认收款并放行。',
    0: '对方已取消订单。',
    3: '对方已完成放行，请注意查收资产。',
    4: '对方已发起申诉。'
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
    ElMessage.error('请求失败')
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
    successMessage: auto ? '订单超时已自动取消' : undefined,
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
    ElMessage.error('请输入资金密码')
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
    ElMessage.info('付款完成 30 分钟后才允许申诉!')
    return
  }

  modal4.value = true
}

const showQRCode = (type) => {
  payCodeUrl.value = type === 1 ? alipay.value.qrCodeUrl : wechatPay.value.qrWeCodeUrl
  modal6.value = true
}

const resetStatus = () => {
  closeReserveCountdown()
  if (canAutoCancelOrder({ status: statusBtn.value, tradeType: tradeType.value })) {
    handleCancelOrder({ auto: true })
  }
}

const getReserveTime = () => {
  const createTime = new Date((msg.value.createTime || '').replace(/-/g, '/')).getTime()
  const throughSeconds = parseInt((Date.now() - createTime) / 1000, 10)
  const reserveSeconds = parseInt(msg.value.timeLimit, 10) * 60 - throughSeconds
  const minutes = parseInt(reserveSeconds / 60, 10)
  const seconds = parseInt(reserveSeconds % 60, 10)
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
  } catch {
    ElMessage.error('请求失败')
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
})

onBeforeUnmount(() => {
  closeReserveCountdown()
  socket?.disconnect()
})
</script>

<style scoped lang="scss">
.content-wrap {
  min-height: 600px;
  padding-top: 80px;
}

.breadcrumb {
  padding: 10px 0 10px 20px;
  font-size: 16px;
}

.breadcrumb-link {
  color: #f0a70a;
}

.breadcrumb-separator,
.breadcrumb-current {
  font-size: 14px;
  color: #c7d0dd;
  margin-left: 4px;
}

.container {
  width: 85%;
  margin: 0 auto;
  min-width: 1200px;
}

.row {
  display: flex;
  flex-wrap: wrap;

  .col-left {
    flex: 0 0 16.666667%;
    max-width: 16.666667%;
  }

  .col-right {
    flex: 0 0 83.333333%;
    max-width: 83.333333%;
  }
}

.chat-in-box {
  .chat-in {
    .chat-right {
      .chat-right-in {
        h6 {
          font-size: 14px;
          font-weight: 500;
          color: #fff;
          min-width: 195px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          margin-bottom: 8px;

          a {
            color: #f0a70a;
          }

          span {
            margin-left: 6px;
          }
        }

        p {
          color: #ccc;
          font-size: 12px;
          margin-bottom: 8px;
          line-height: 1.5;

          em {
            color: #f40a0a;
            font-style: normal;
          }
        }

        h5 {
          font-weight: 600;
          margin-bottom: 10px;
          color: #fff;
        }
      }
    }
  }
}

.leftmenu {
  margin-bottom: 60px;
  background-color: #192330;
  position: relative;
  min-height: 1px;
  padding: 50px 15px;
  text-align: left;
}

.mt20 {
  margin-top: 20px;
}

.h6-flex {
  display: flex;
  overflow: auto;
  min-width: auto;
  white-space: normal;

  > span {
    flex: 0 0 40px;
  }

  > a {
    flex: 1 1 100%;
    width: 100%;
  }
}

.status-actions {
  padding-top: 20px;
}

.status-label {
  font-weight: 600;
}

.reserve-time {
  color: #ed3f14;
  font-weight: 700;
}

.rightbox {
  background: #192330;
  margin-left: 20px;
  padding-bottom: 20px;
  margin-bottom: 40px;
}

.chat-top {
  padding: 30px 0;
  font-size: 14px;
}

.chat-summary,
.chat-payment {
  justify-content: space-between;
}

.chat-summary .col,
.chat-payment .col {
  flex: 1;
}

.chat-top {
  .order-time {
    h5 {
      font-size: 16px;
      line-height: 40px;
    }
  }

  .order-info {
    h5 {
      font-weight: 600;
      font-size: 14px;
    }

    p {
      a {
        color: #f0a70a;
      }
    }
  }
}

.payment-tip {
  color: rgb(145, 143, 143);
}

.dialog-warning {
  color: #f56c6c;
  font-weight: 700;
}

.fundpwd-input {
  margin-top: 10px;
}

.qrcode-dialog {
  text-align: center;
}

.qrcode-image {
  width: 200px;
}

.icons {
  height: 17px;
  width: 17px;
  display: inline-block;
  margin-top: -1px;
  background-size: 100% 100%;
  vertical-align: middle;

  &.alipay {
    background-image: url(../../assets/img/alipay.png);
  }

  &.wechat {
    background-image: url(../../assets/img/wechat.png);
  }

  &.bankfor {
    background-image: url(../../assets/img/bankcard.png);
  }
}
</style>

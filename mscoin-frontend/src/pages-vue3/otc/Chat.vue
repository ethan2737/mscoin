<template>
  <div class="content-wrap">
    <div class="container chat-in-box" id="List">
      <p style="padding: 10px 0 10px 20px;font-size: 16px;">
        <router-link to="/uc/order" style="color:#f0a70a;">{{ $t('otc.myorder') }}</router-link> > <span style="font-size:14px;">订单详情</span>
      </p>
      <div class="row chat-in">
        <div class="col-left">
          <div class="leftmenu left-box chat-right">
            <div class="chat-right-in">
              <h6 class="h6-flex">
                <span v-if="tradeType === 0">{{ $t('otc.chat.seller') }}:</span>
                <span v-else>{{ $t('otc.chat.buyer') }}:</span>
                <router-link :to="{ path: '/checkuser', query: { 'id': msg.otherSide }}">
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
                  <p>1、{{ $t('otc.chat.operatetip_1') }}"<em>{{ $t('otc.chat.finishpayment') }}</em>"。{{ $t('otc.chat.operatetip_1_1') }}。</p>
                  <p>2、{{ $t('otc.chat.operatetip_1_2') }}。</p>
                </div>
                <span><b>{{ $t('otc.chat.note') }}：</b></span><br>
                <span>{{ $t('otc.chat.notetip') }}</span><br>
              </div>
              <div class="mt20" v-else>
                <h5>{{ $t('otc.chat.operatetip') }}:</h5>
                <div>
                  <p>1、{{ $t('otc.chat.operatetip_2_1') }}"<em>{{ $t('otc.chat.confirmrelease') }}</em>"{{ $t('otc.chat.paydigital') }}！</p>
                  <p>2、{{ $t('otc.chat.operatetip_2_2') }}</p>
                  <p>3、{{ $t('otc.chat.operatetip_2_3') }}</p>
                </div>
                <span><b>{{ $t('otc.chat.note') }}：</b></span><br>
                <span>{{ $t('otc.chat.notetip') }}</span><br>
              </div>
              <div class="bottom-btn">
                <div style="padding-top:20px;">
                  <h6 style="font-weight: 600">{{ $t('otc.chat.orderstatus') }}:
                    <span>{{ statusText }}</span>
                  </h6>
                  <div v-show="statusBtn === 1 && tradeType === 0">
                    <el-button type="warning" @click="modal1 = true">{{ $t('otc.chat.orderstatus_1') }}</el-button>
                    <el-button type="error" @click="modal3 = true">{{ $t('otc.chat.orderstatus_4') }}</el-button>
                  </div>
                  <div v-show="statusBtn === 2 && tradeType === 0">
                    <el-button type="warning" @click="appearOrder">{{ $t('otc.chat.orderstatus_2') }}</el-button>
                    <el-button type="error" @click="modal3 = true">{{ $t('otc.chat.orderstatus_4') }}</el-button>
                  </div>
                  <div v-show="statusBtn === 2 && tradeType === 1">
                    <el-button type="warning" @click="modal5 = true">{{ $t('otc.chat.orderstatus_3') }}</el-button>
                    <el-button type="error" @click="appearOrder">{{ $t('otc.chat.orderstatus_5') }}</el-button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col-right">
          <div class="rightbox">
            <div class="row chat-top" style="display: flex; justify-content: space-between;">
              <div class="col order-time" style="flex: 1;">
                <h5>{{ statusText }}</h5>
                <div v-show="statusBtn === 1" class="reserve-time">{{ reserveTime }}</div>
              </div>
              <div class="col order-info" style="flex: 1;">
                <h5>
                  <label class="order-name">{{ $t('otc.chat.order') }}</label>
                  <span>{{ msg.orderSn }}</span>
                </h5>
                <p>
                  {{ $t('otc.chat.and') }}
                  <router-link :to="{ path: '/checkuser', query: { 'id': msg.otherSide }}">
                    {{ msg.otherSide }}
                  </router-link>
                  {{ $t('otc.chat.transition') }}
                </p>
              </div>
              <div class="col order-info" style="flex: 1;">
                <h5>{{ msg.price }}</h5>
                <span>{{ $t('otc.chat.transprice') }}(CNY)</span>
              </div>
              <div class="col order-info" style="flex: 1;">
                <h5>{{ msg.amount }}</h5>
                <span>{{ $t('otc.chat.transnum') }}({{ msg.unit }})</span>
              </div>
              <div class="col order-info" style="flex: 1;">
                <h5>{{ msg.money }}</h5>
                <span>{{ $t('otc.chat.transmoney') }}(CNY)</span>
              </div>
            </div>
            <div class="row chat-top" style="display: flex; justify-content: space-between;" v-show="statusBtn !== 0">
              <div class="col order-info" style="flex: 1;" v-if="bankInfo && bankInfo !== null">
                <i class="icons bankfor"></i>
                <span>{{ payInfo != null ? payInfo.realName : "" }} </span>
                <p>{{ bankInfo.branch }}</p>
                <p>{{ bankInfo.cardNo }}</p>
              </div>
              <div class="col order-info" style="flex: 1;" v-else>
                <i class="icons bankfor"></i>
                <pre></pre>
                <p style="color:rgb(145, 143, 143)">{{ $t('otc.chat.tip1') }}</p>
              </div>
              <div class="col order-info" style="flex: 1;" v-if="alipay && alipay !== null">
                <i class="icons alipay"></i>
                <span>{{ $t('otc.chat.zfb') }}</span>
                <pre></pre>
                <p>{{ alipay.aliNo }}</p>
                <p v-if="alipay && alipay !== null && alipay.qrCodeUrl !== null && alipay.qrCodeUrl !== ''">
                  <a @click="showQRCode(1)">{{ $t('otc.chat.qrcode') }}</a>
                </p>
              </div>
              <div class="col order-info" style="flex: 1;" v-else>
                <i class="icons alipay"></i>
                <pre></pre>
                <p style="color:rgb(145, 143, 143)">{{ $t('otc.chat.tip2') }}</p>
              </div>
              <div class="col order-info" style="flex: 1;" v-if="wechatPay && wechatPay !== null">
                <i class="icons wechat"></i>
                <span>{{ $t('otc.chat.wx') }}</span>
                <pre></pre>
                <p>{{ wechatPay.wechat }}</p>
                <p v-if="wechatPay && wechatPay !== null && wechatPay.qrWeCodeUrl !== null && wechatPay.qrWeCodeUrl !== ''">
                  <a @click="showQRCode(2)">{{ $t('otc.chat.qrcode') }}</a>
                </p>
              </div>
              <div class="col order-info" style="flex: 1;" v-else>
                <i class="icons wechat"></i>
                <pre></pre>
                <p style="color:rgb(145, 143, 143)">{{ $t('otc.chat.tip3') }}</p>
              </div>
            </div>
            <Chatline :msg="msg" />
          </div>
        </div>
      </div>
    </div>
    <el-dialog v-model="modal1" :title="$t('otc.chat.tip')" @close="onDialog1Close">
      <p style="color:red;font-weight: bold;">{{ $t('otc.chat.msg1') }}</p>
    </el-dialog>
    <el-dialog v-model="modal3" :title="$t('otc.chat.tip')" @close="onDialog3Close">
      <p style="color:red;font-weight: bold;">{{ $t('otc.chat.msg3') }}</p>
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
    </el-dialog>
    <el-dialog v-model="modal5" :title="$t('otc.chat.tip')" @close="onDialog5Close">
      <p style="color:red;font-weight: bold;">
        {{ $t('otc.chat.msg6') }}<br/>
        <el-input type="password" v-model="fundpwd" :placeholder="$t('otc.chat.msg7')" />
      </p>
    </el-dialog>
    <el-dialog v-model="modal6">
      <template #header></template>
      <div style="text-align:center">
        <img style="width: 200px;" :src="payCodeUrl" />
      </div>
      <template #footer></template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 聊天记录页面
 */
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { ElMessage, ElButton, ElDialog, ElForm, ElFormItem, ElSelect, ElOption, ElInput, ElPopover } from 'element-plus'
import axios from 'axios'
import { useRouter } from 'vue-router/composables'
import Chatline from './Chatline.vue'

const router = useRouter()

const host = 'http://localhost'
let socket = null

const watching = ref(false)
const reserveTime = ref('60')
const reserveInteval = ref(null)
const fundpwd = ref('')
const statusBtn = ref(0)
const tradeType = ref(0)
const payTime = ref('')
const statusText = ref('')
const modal1 = ref(false)
const modal2 = ref(false)
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

const ok1 = () => {
  axios.post(`${host}/otc/order/pay`, {
    orderSn: router.currentRoute.value.query.tradeId
  }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success(resp.message)
      sendOrderStatusNotice(1)
      getDetail()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const ok3 = () => {
  axios.post(`${host}/otc/order/cancel`, {
    orderSn: router.currentRoute.value.query.tradeId
  }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success(resp.message)
      sendOrderStatusNotice(3)
      getDetail()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const ok4 = () => {
  const params = {
    orderSn: router.currentRoute.value.query.tradeId,
    remark: formItem.value.remark
  }

  axios.post(`${host}/otc/order/appeal`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success(resp.message)
      sendOrderStatusNotice(4)
      getDetail()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const ok5 = () => {
  if (fundpwd.value === '') {
    ElMessage.error('请输入资金密码')
    return
  }

  const params = {
    orderSn: router.currentRoute.value.query.tradeId,
    jyPassword: fundpwd.value
  }

  axios.post(`${host}/otc/order/release`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success(resp.message)
      sendOrderStatusNotice(5)
      getDetail()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const onDialog1Close = () => {
  ok1()
}

const onDialog3Close = () => {
  ok3()
}

const onDialog4Close = () => {
  ok4()
}

const onDialog5Close = () => {
  ok5()
}

const appearOrder = () => {
  const nowTime = new Date().getTime()
  const payTimeVal = new Date(msg.value.payTime).getTime()
  if (parseInt((nowTime - payTimeVal) / 1000) < 1800) {
    ElMessage.info('付款完成 30 分钟后才允许申诉!')
    return
  }
  modal4.value = true
}

const showQRCode = (type) => {
  if (type === 1) {
    payCodeUrl.value = alipay.value.qrCodeUrl
  } else {
    payCodeUrl.value = wechatPay.value.qrWeCodeUrl
  }
  modal6.value = true
}

const setReserveTime = () => {
  getReserveTime()
  reserveInteval.value = setInterval(() => {
    getReserveTime()
  }, 1000)
}

const getReserveTime = () => {
  const d1 = new Date().getTime()
  const d2 = new Date(msg.value.createTime?.replace(/-/g, '/')).getTime()
  const throughSeconds = parseInt((d1 - d2) / 1000)
  const reserveSeconds = parseInt(msg.value.timeLimit) * 60 - throughSeconds
  const minutes = parseInt(reserveSeconds / 60)
  const seconds = parseInt(reserveSeconds % 60)
  reserveTime.value = minutes + ':' + (seconds >= 10 ? seconds : '0' + seconds)
  if (reserveSeconds <= 0) {
    resetStatus()
  }
}

const resetStatus = () => {
  if (reserveInteval.value) clearInterval(reserveInteval.value)
  statusBtn.value = 5
  ok3()
}

const sendOrderStatusNotice = (type) => {
  if (reserveInteval.value) clearInterval(reserveInteval.value)
  let content = ''
  if (type === 1) content = '对方已付款，请查收并确认放行!'
  else if (type === 3) content = '对方已取消订单!'
  else if (type === 4) content = '对方已申诉!'
  else if (type === 5) content = '对方已放行，请查收!'

  const jsonParam = {
    uidTo: msg.value.hisId,
    uidFrom: msg.value.myId,
    orderId: router.currentRoute.value.query.tradeId,
    content: content,
    messageType: 0
  }
  socket?.emit('/app/message/chat', JSON.stringify(jsonParam))
}

const watchOrderStutusNotice = () => {
  socket.on('/user/' + msg.value.myId + '/order-notice/' + router.currentRoute.value.query.tradeId, (response) => {
    if (reserveInteval.value) clearInterval(reserveInteval.value)
    const confirmPayMsg = JSON.parse(response)
    ElMessage.success(confirmPayMsg.content)
    statusBtn.value = confirmPayMsg.status
    if (confirmPayMsg.status === 1) {
      statusText.value = '待确认收款'
      setReserveTime()
    } else if (confirmPayMsg.status === 2) {
      statusText.value = '交易完成'
    } else if (confirmPayMsg.status === 3) {
      statusText.value = '已取消'
    } else if (confirmPayMsg.status === 4) {
      statusText.value = '已申诉'
    } else if (confirmPayMsg.status === 0) {
      statusText.value = '订单关闭'
    }
  })
}

const getDetail = () => {
  axios.post(`${host}/otc/order/detail`, {
    orderSn: router.currentRoute.value.query.tradeId
  }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      msg.value = resp.data
      payInfo.value = resp.data.payInfo
      if (payInfo.value === null) {
        bankInfo.value = alipay.value = wechatPay.value = null
      } else {
        bankInfo.value = resp.data.payInfo.bankInfo
        alipay.value = resp.data.payInfo.alipay
        wechatPay.value = resp.data.payInfo.wechatPay
      }

      if (!watching.value) {
        watchOrderStutusNotice()
        watching.value = true
      }

      statusBtn.value = resp.data.status
      tradeType.value = resp.data.type
      if (resp.data.status === 1) {
        statusText.value = '待确认收款'
        setReserveTime()
      } else if (resp.data.status === 2) {
        statusText.value = '交易完成'
      } else if (resp.data.status === 3) {
        statusText.value = '已取消'
      } else if (resp.data.status === 4) {
        statusText.value = '已申诉'
      } else if (resp.data.status === 0) {
        statusText.value = '订单关闭'
      }
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const initSocket = () => {
  socket = io('http://localhost')
}

onMounted(() => {
  initSocket()
  getDetail()
})

onBeforeUnmount(() => {
  if (reserveInteval.value) clearInterval(reserveInteval.value)
  socket?.disconnect()
})
</script>

<style scoped lang="scss">
.content-wrap {
  min-height: 600px;
  padding-top: 80px;
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
  padding: 50px 15px 50px 15px;
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

.pop-tel {
  position: absolute;
  top: 50px;
  right: 10px;
  width: 25px;
  height: 25px;
  z-index: 100;

  img {
    width: 100%;
    height: 100%;
  }
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

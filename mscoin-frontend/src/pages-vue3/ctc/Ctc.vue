<template>
  <div class="ctc">
    <img class="bannerimg shoujiIf" src="../../assets/images/ctc-bg.jpg" alt="">
    <div class="ctc_container">
      <h1 class="shoujiIf">{{$t('ctc.title')}}</h1>
      <p class="shoujiIf" style="letter-spacing: 1px;">{{$t('ctc.desc')}}</p>
      <div class="main">
        <el-tabs v-model="activeTab" style="width:100%;">
          <el-tab-pane name="all">
            <template #label>
              <span>USDT{{$t('ctc.trade')}}</span>
            </template>
            <div class="ctc-container">
              <div class="trade_wrap">
                <div class="trade_panel">
                  <div class="trade_bd_ctc">
                    <div class="panel panel_buy">
                      <div class="bd bd_limited">
                        <el-form>
                          <el-form-item class="buy-input">
                            <label class="before">{{$t('ctc.buyprice')}}</label>
                            <el-input v-model="buyPrice" disabled></el-input>
                            <label class="after" style="color: #45b854;">CNY</label>
                          </el-form-item>
                          <el-form-item class="trade-input">
                            <label class="before">{{$t('ctc.buynum')}}:</label>
                            <el-input-number style="width:70%;float:right;" v-model="buyAmount" size="large" :max="50000" :min="50" :placeholder="$t('ctc.input50tips')"></el-input-number>
                            <label class="after">USDT</label>
                          </el-form-item>
                          <p style="font-size: 12px;margin-top: -20px;text-align:right;margin-bottom: 10px;">&nbsp; </p>
                          <el-form-item>
                            <label class="before">{{$t('ctc.payType')}}:</label>
                            <el-select v-model="payType" style="width:70%;float:right;" size="large">
                              <el-option v-for="item in payTypeList" :value="item.value" :key="item.value">{{ item.label }}</el-option>
                            </el-select>
                          </el-form-item>
                          <div style="height: 30px;line-height: 30px;margin-top: -20px;margin-bottom: 5px;color: #0074eb;text-align:right;font-size:12px;">
                            <router-link to="/uc/account">{{$t("ctc.payset")}}</router-link>
                          </div>
                          <div class="total buy_total" style="min-height">
                            <div style="min-height: 40px;">
                              <div style="float:left;">{{$t('ctc.payamount')}}</div>
                              <div style="float:right;">
                                <span style="color: #45b854;font-size:24px;font-weight: 600;">{{totalBuyMoney}}</span>
                                <span style="font-size:14px;margin-left: 5px;color: #45b854;">CNY</span>
                              </div>
                            </div>
                            <div style="width: 100%;font-size:12px;text-align:right;">{{$t("ctc.moneyTips")}}</div>
                          </div>
                          <el-button style="padding-bottom: 10px;padding-top: 10px;" class="bg-green" size="large" @click="buyClick">{{$t("ctc.buyin")}} USDT</el-button>
                        </el-form>
                      </div>
                    </div>
                    <div class="panel panel_sell">
                      <div class="bd bd_limited">
                        <el-form ref="formValidateRef" :model="formValidate">
                          <el-form-item class="sell-input">
                            <label class="before">{{$t('ctc.sellprice')}}</label>
                            <el-input v-model="sellPrice" disabled></el-input>
                            <label class="after" style="color: #f2334f;">CNY</label>
                          </el-form-item>
                          <el-form-item class="trade-input">
                            <label class="before">{{$t('ctc.sellnum')}}:</label>
                            <el-input-number style="width:70%;float:right;" v-model="sellAmount" size="large" :max="50000" :min="2" :placeholder="$t('ctc.input2tips')"></el-input-number>
                            <label class="after">USDT</label>
                          </el-form-item>
                          <p style="font-size: 12px;margin-top: -20px;text-align:right;margin-bottom: 10px;">
                            <span>{{$t('ctc.avabalance')}}</span>:
                            <span>{{wallet.base || '--'}}</span><span style="margin-left: 5px;">USDT</span>
                          </p>
                          <el-form-item>
                            <label class="before">{{$t('ctc.receiveType')}}:</label>
                            <el-select v-model="receiveType" style="width:70%;float:right;" size="large">
                              <el-option v-for="item in receiveTypeList" :value="item.value" :key="item.value">{{ item.label }}</el-option>
                            </el-select>
                          </el-form-item>
                          <div style="height: 30px;line-height: 30px;margin-top: -20px;text-align:right;margin-bottom: 5px;font-size: 12px;">
                            {{$t("ctc.useselfaccount")}}
                          </div>
                          <div class="total buy_total">
                            <div style="min-height: 40px;">
                              <div style="float:left;">{{$t('ctc.getamount')}}</div>
                              <div style="float:right;color: #f2334f;">
                                <span style="color: #f2334f;font-size:24px;font-weight: 600;">{{totalSellMoney}}</span>
                                <span style="font-size:14px;margin-left: 5px;">CNY</span>
                              </div>
                            </div>
                            <div style="width: 100%;font-size:12px;text-align:right;">{{$t("ctc.moneyTips")}}</div>
                          </div>
                          <el-button style="padding-bottom: 10px;padding-top: 10px;" class="bg-red" size="large" @click="sellClick">{{$t("ctc.sell")}} USDT</el-button>
                        </el-form>
                      </div>
                    </div>
                  </div>

                  <div class="shoujiIf" style="float:right;width: 30%;height: 455px;font-size: 12px;color: #bcbcbc;">
                    <div style="padding: 25px 35px;width: 100%;height: 395px;overflow-y: auto; overflow-x:hidden;background-color:#192330;text-align:left;line-height: 26px;">
                      <p style="text-align:center;font-size: 18px;margin-bottom: 10px;">{{$t("ctc.notice")}}</p>
                      <p>{{$t("ctc.notice1")}}</p>
                      <p>{{$t("ctc.notice2")}}</p>
                      <p>{{$t("ctc.notice3")}}</p>
                      <p>{{$t("ctc.notice4")}}</p>
                      <p>{{$t("ctc.notice5")}}</p>
                      <router-link target="_blank" to="/helpdetail?cate=2&id=40&cateTitle=交易指南" style="float:right;">{{$t("ctc.moredetail")}}</router-link>
                    </div>
                    <div class="notice-bottom">
                      <router-link to="/uc/safe" class="notice-btn-left">{{$t("ctc.verifyset")}}</router-link>
                      <router-link to="/uc/account" class="notice-btn-right">{{$t("ctc.payset")}}</router-link>
                    </div>
                  </div>
                  <div></div>
                </div>
              </div>
            </div>
            <div class="table shoujiIf">
              <el-table :no-data-text="$t('common.nodata')" :columns="columns" :data="orders" :loading="loading"></el-table>
              <div class="page">
                <el-pagination :total="total" :page-size="pageSize" :current-page="pageNo" @current-change="loadDataPage"></el-pagination>
              </div>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>

    <el-dialog v-model="modal" width="450px">
      <template #title>
        {{$t("ctc.tip")}}
      </template>
      <el-form class="withdraw-form-inline" ref="formInlineRef" :model="formInline">
        <el-form-item prop="code">
          <el-input id="verifyCode" style="width:calc(100% - 105px)" type="text" autocomplete="off" v-model="formInline.code" :placeholder="$t('uc.regist.smscode')">
          </el-input>
          <input id="sendCode" style="width:100px;border-radius: 10px;height: 30px;line-height: 10px;position: relative;top: 2px;" @click="sendCode" type="button" :value="sendcodeValue" :disabled="codeIsSending">
          </input>
        </el-form-item>
        <el-form-item>
          <el-input id="fundPwd" type="password" autocomplete="off" v-model="formInline.fundpwd" :placeholder="$t('otc.chat.msg7')"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span style="margin-right:50px;cursor: pointer;" @click="cancel">取消</span>
        <span style="background:#f0ac19;cursor: pointer;color:#fff;width:80px;border-radius:30px;display:inline-block;text-align:center;height:30px;line-height: 30px;" @click="ok">确定</span>
      </template>
    </el-dialog>

    <el-dialog v-model="detailModal" title="订单详情" @close="closeDetail">
      <p class="ctc-order-status" v-if="detailOrder.direction == 0 && detailOrder.status == 0">订单状态：等待承兑商接单...</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 0 && detailOrder.status == 1">订单状态：承兑商已接单，等待您付款中</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 0 && detailOrder.status == 2">订单状态：已付款，等待承兑商放币</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 0 && detailOrder.status == 3">订单状态：已完成</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 0 && detailOrder.status == 4">订单状态：已取消 ({{detailOrder.cancelReason}})</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 1 && detailOrder.status == 0">订单状态：等待承兑商接单...</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 1 && detailOrder.status == 1">订单状态：承兑商已接单，正在付款中</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 1 && detailOrder.status == 2">订单状态：承兑商已付款，确认放币中</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 1 && detailOrder.status == 3">订单状态：已完成</p>
      <p class="ctc-order-status" v-if="detailOrder.direction == 1 && detailOrder.status == 4">订单状态：已取消 ({{detailOrder.cancelReason}})</p>
      <el-row style="background: #27384a;padding: 10px 0px;border-radius: 5px;">
        <el-col :span="8">
          <p v-if="detailOrder.direction == 0" class="item-title">买入</p>
          <p v-if="detailOrder.direction == 1" class="item-title">卖出</p>
          <p class="item-desc">订单类型</p>
        </el-col>
        <el-col :span="8">
          <p v-if="detailOrder.direction == 0" class="item-title">{{toFixed(detailOrder.amount, 2)}} <span class="unit">USDT</span></p>
          <p v-if="detailOrder.direction == 1" class="item-title">{{toFixed(detailOrder.amount, 2)}} <span class="unit">USDT</span></p>
          <p class="item-desc">交易数量</p>
        </el-col>
        <el-col :span="8">
          <p v-if="detailOrder.direction == 0" class="item-title green">{{toFixed(detailOrder.money, 2)}} <span class="unit">CNY</span></p>
          <p v-if="detailOrder.direction == 1" class="item-title red">{{toFixed(detailOrder.money, 2)}} <span class="unit">CNY</span></p>
          <p class="item-desc">交易总额</p>
        </el-col>
      </el-row>

      <div style="font-size: 12px;margin-top: 15px;" v-if="detailOrder.direction == 0">
        <el-icon style="color:rgb(183, 183, 183);margin-right:5px;font-size:14px;"><InfoFilled /></el-icon>请向以下收款账户汇款/转账：<span class="green" style="font-size: 20px;font-weight:bold;">{{toFixed(detailOrder.money, 2)}}</span> <span class="green">CNY</span>

        <div style="float:right;padding: 2px 10px;color:#FF0000;" v-if="orderCountdown > 0 && (detailOrder.status == 0 || detailOrder.status == 1)">
          <el-icon style="font-weight:bold;font-size:18px;margin-top:-5px;margin-right: 3px;"><Clock /></el-icon>
          <span style="font-size:16px;">{{formatCountdown(orderCountdown)}}</span>
        </div>
      </div>

      <div style="font-size: 12px;margin-top: 15px;" v-if="detailOrder.direction == 1">
        <el-icon style="color:rgb(183, 183, 183);margin-right:5px;font-size:14px;"><InfoFilled /></el-icon>你的以下账户将收到汇款/转账：<span class="red" style="font-size: 20px;font-weight:bold;">{{toFixed(detailOrder.money, 2)}}</span> <span class="red">CNY</span>
      </div>

      <el-row style="margin-top: 5px;background: #27384a;padding: 20px 0px;border-radius: 5px;">
        <el-col :span="24">
          <div style="float:left;margin-left:20px;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">账户实名:</span>
            <span style="font-size:14px;color:#FFF;">{{detailOrder.realName}}</span>
          </div>
          <div style="float:right;margin-right: 20px;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">收款方式:</span>
            <span style="font-size:14px;color:#FFF;" v-if="detailOrder.payMode == 'bank'">银行卡</span>
            <span style="font-size:14px;color:#FFF;" v-if="detailOrder.payMode == 'alipay'">支付宝</span>
            <span style="font-size:14px;color:#FFF;" v-if="detailOrder.payMode == 'wechatpay'">微信支付</span>
          </div>
        </el-col>
        <el-col :span="24" v-if="detailOrder.payMode == 'bank'" style="margin-top: 10px;text-align:left;">
          <div style="float:left;margin-left:20px;width: 100%;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">开户银行:</span>
            <span style="font-size:14px;color:#FFF;">{{detailOrder.bankInfo.bank}}</span>
          </div>
          <div style="float:left;margin-left:20px;width: 100%;margin-top: 10px;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">开户支行:</span>
            <span style="font-size:14px;color:#FFF;">{{detailOrder.bankInfo.branch}}</span>
          </div>
          <div style="float:left;margin-left:20px;width: 100%;margin-top: 10px;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">银行卡号:</span>
            <span style="font-size:16px;color:#FFF;letter-spacing: 3px;font-weight:bold;">{{detailOrder.bankInfo.cardNo}}</span>
          </div>
        </el-col>

        <el-col :span="24" v-if="detailOrder.payMode == 'alipay'" style="margin-top: 10px;text-align:left;">
          <div style="float:left;margin-left:20px;width: 100%;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">支付宝账号:</span>
            <span style="font-size:14px;color:#FFF;">{{detailOrder.alipay.aliNo}}</span>
          </div>
          <div style="float:left;margin-left:20px;width: 100%;margin-top: 10px;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">收款码:</span>
          </div>
          <div style="float:left;margin-left:20px;width: 100%;margin-top: 10px;text-align:center;">
            <img :src="detailOrder.alipay.qrCodeUrl" style="width: 300px;height:400px;" alt="">
          </div>
        </el-col>

        <el-col :span="24" v-if="detailOrder.payMode == 'wechatpay'" style="margin-top: 10px;text-align:left;">
          <div style="float:left;margin-left:20px;width: 100%;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">微信账号:</span>
            <span style="font-size:14px;color:#FFF;">{{detailOrder.wechatPay.wechat}}</span>
          </div>
          <div style="float:left;margin-left:20px;width: 100%;margin-top: 10px;">
            <span style="color:rgb(190, 190, 190);font-size:12px;">收款码:</span>
          </div>
          <div style="float:left;margin-left:20px;width: 100%;margin-top: 10px;text-align:center;">
            <img :src="detailOrder.wechatPay.qrWeCodeUrl" style="width: 300px;height:400px;" alt="">
          </div>
        </el-col>
      </el-row>

      <template #footer>
        <el-button v-if="(detailOrder.direction==1 && detailOrder.status==0) || (detailOrder.direction==0 && detailOrder.status < 2)" type="danger" size="large" @click="cancelOrderClick">撤消订单</el-button>
        <el-button v-if="detailOrder.direction==0 && detailOrder.status == 1" type="success" size="large" @click="payOrderClick">标记已付款</el-button>
        <el-button type="default" size="large" @click="closeDetail">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - CTC 交易页面
 */
import { ref, reactive, computed, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { ElMessage, ElNotification, ElTable, ElPagination, ElTabs, ElTabPane, ElForm, ElFormItem, ElInput, ElInputNumber, ElSelect, ElOption, ElButton, ElDialog, ElCol, ElRow, ElIcon } from 'element-plus'
import { InfoFilled, Clock } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import moment from 'moment'

const router = useRouter()
const store = useStore()

const host = ''

const activeTab = ref('all')
const formValidateRef = ref(null)
const formInlineRef = ref(null)

const payTypeList = ref([{
  label: '银行',
  value: 'bank'
},{
  label: '支付宝',
  value: 'alipay'
},{
  label: '微信支付',
  value: 'wechatpay'
}])

const receiveTypeList = ref([{
  label: '银行',
  value: 'bank'
}])

const countdown = ref(60)
const timer = ref(null)
const orderTimer = ref(null)
const orderCountdown = ref(0)
const direction = ref('buy')
const receiveType = ref('bank')
const payType = ref('bank')
const buyPrice = ref(7.00)
const buyAmount = ref(null)
const sellPrice = ref(7.00)
const sellAmount = ref(null)
const modal = ref(false)
const detailModal = ref(false)
const formInline = reactive({
  code: '',
  fundpwd: ''
})
const formValidate = reactive({})
const fundpwd = ref('')
const codeIsSending = ref(false)
const sendcodeValue = ref('发送验证码')
const loading = ref(false)
const pageSize = ref(10)
const pageNo = ref(1)
const total = ref(10)
const user = reactive({})
const userAccount = reactive({})
const orders = ref([])
const detailOrder = reactive({})
const wallet = reactive({
  base: 0
})

const lang = computed(() => store.state.lang)
const langPram = computed(() => {
  if (store.state.lang === '简体中文') return 'CN'
  if (store.state.lang === 'English') return 'EN'
  return 'CN'
})
const isLogin = computed(() => store.getters.isLogin)
const totalBuyMoney = computed(() => (buyPrice.value * buyAmount.value).toFixed(2))
const totalSellMoney = computed(() => (sellPrice.value * sellAmount.value).toFixed(2))

const dateFormat = (tick) => moment(tick).format('YYYY-MM-DD HH:mm:ss')

const toFixed = (value, scale) => {
  if (value !== undefined && value !== null && value !== '') {
    return Number(value).toFixed(scale)
  }
  return 0
}

const formatCountdown = (value) => {
  const m = parseInt(value / 60)
  const s = value % 60
  const mm = m < 10 ? '0' + m : m
  const ss = s < 10 ? '0' + s : s
  return mm + ':' + ss
}

const getC2cPrice = () => {
  axios.post(`${host}/market/ctc-usdt`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      buyPrice.value = res.data.data.buy
      sellPrice.value = res.data.data.sell
    }
  }).catch(() => {})
}

const getWallet = () => {
  axios.post(`${host}/uc/walletUSDT`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      wallet.base = res.data.data.balance || 0
    }
  }).catch(() => {})
}

const getAccountSecurity = () => {
  axios.post(`${host}/uc/approve/security/setting`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      Object.assign(user, res.data.data)
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const getAccount = () => {
  axios.post(`${host}/uc/approve/account/setting`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      Object.assign(userAccount, res.data.data)
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const getOrderList = () => {
  loading.value = true
  const params = {
    pageNo: pageNo.value,
    pageSize: pageSize.value
  }
  axios.post(`${host}/uc/ctc/page-query`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      if (res.data.data.content.length > 0) {
        total.value = res.data.data.totalElements
        orders.value = res.data.data.content
      }
    }
    loading.value = false
  }).catch(() => {
    loading.value = false
  })
}

const loadDataPage = (page) => {
  pageNo.value = page
  getOrderList()
}

const createOrder = () => {
  const params = {}
  if (direction.value === 'buy') {
    params.price = buyPrice.value
    params.amount = buyAmount.value
    params.payType = payType.value
    params.direction = 0
  } else {
    params.price = sellPrice.value
    params.amount = sellAmount.value
    params.payType = receiveType.value
    params.direction = 1
  }
  params.unit = 'USDT'
  params.fundpwd = formInline.fundpwd
  params.code = formInline.code

  axios.post(`${host}/uc/ctc/new-ctc-order`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      getOrderList()
      Object.assign(detailOrder, res.data.data)
      modal.value = false
      showDetailModal()
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const showDetailModal = () => {
  if (orderTimer.value) {
    clearInterval(orderTimer.value)
  }
  if (detailOrder.currentTime) {
    const currentT = parseInt(new Date(detailOrder.currentTime).getTime() / 1000)
    if (detailOrder.status === 0) {
      const endT = parseInt(new Date(detailOrder.createTime).getTime() / 1000)
      orderCountdown.value = currentT - endT
    } else if (detailOrder.status === 1) {
      const endT = parseInt(new Date(detailOrder.confirmTime).getTime() / 1000)
      orderCountdown.value = currentT - endT
    }

    if (orderCountdown.value < 30 * 60) {
      orderCountdown.value = 1800 - orderCountdown.value
      orderTimer.value = setInterval(() => {
        orderCountdown.value--
        if (orderCountdown.value < 1) {
          clearInterval(orderTimer.value)
        }
      }, 1000)
    }
  }
  detailModal.value = true
}

const cancel = () => {
  modal.value = false
  formInline.code = ''
  formInline.fundpwd = ''
}

const sendCode = () => {
  axios.post(`${host}/uc/mobile/ctc/code`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      settime()
      ElNotification.success({ title: '提示', message: res.data.message })
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const settime = () => {
  sendcodeValue.value = countdown.value + 's 后重新发送'
  codeIsSending.value = true
  const timercode = setInterval(() => {
    countdown.value--
    sendcodeValue.value = countdown.value + 's 后重新发送'
    if (countdown.value <= 0) {
      clearInterval(timercode)
      sendcodeValue.value = '发送验证码'
      countdown.value = 60
      codeIsSending.value = false
    }
  }, 1000)
}

const ok = () => {
  if (formInline.code === '') {
    ElMessage.error('请填写短信验证码')
    return
  }
  if (formInline.fundpwd === '') {
    ElMessage.error('请填写资金密码')
    return
  }
  createOrder()
}

const valid = (type) => {
  if (!isLogin.value) {
    ElMessage.error('请先登录！')
    return false
  }
  if (user.realVerified !== 1) {
    ElMessage.error('请先完成实名认证！')
    return false
  }
  if (user.fundsVerified !== 1) {
    ElMessage.error('请设置资产交易密码！')
    return false
  }
  if (type === 0) {
    if (!buyAmount.value || buyAmount.value < 50 || buyAmount.value > 50000) {
      ElMessage.error('请输入正确的买入数量')
      return false
    }
    return true
  } else {
    if (userAccount.bankVerified !== 1) {
      ElMessage.error('请先绑定银行卡')
      return false
    }
    if (!sellAmount.value || sellAmount.value < 2 || sellAmount.value > 50000) {
      ElMessage.error('请输入正确的卖出数量')
      return false
    }
    return true
  }
}

const buyClick = () => {
  direction.value = 'buy'
  if (valid(0)) {
    modal.value = true
    const timercode = setInterval(() => {
      if (countdown.value <= 0) {
        clearInterval(timercode)
        sendcodeValue.value = '发送验证码'
        codeIsSending.value = false
      }
    }, 1000)
  }
}

const sellClick = () => {
  direction.value = 'sell'
  if (valid(1)) {
    modal.value = true
    const timercode = setInterval(() => {
      if (countdown.value <= 0) {
        clearInterval(timercode)
        sendcodeValue.value = '发送验证码'
        codeIsSending.value = false
      }
    }, 1000)
  }
}

const cancelOrderClick = () => {
  ElMessageBox.confirm('您确定要取消该笔订单吗？', '确定取消订单吗', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    cancelOrder()
  }).catch(() => {})
}

const payOrderClick = () => {
  ElMessageBox.confirm('标记已付款前请确认您已付款！注意：对于恶意标记付款的账户，我们将对您的账户进行冻结等限制！', '确定您已付款吗？', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    payOrder()
  }).catch(() => {})
}

const closeDetail = () => {
  detailModal.value = false
}

const payOrder = () => {
  const params = { oid: detailOrder.id }
  axios.post(`${host}/uc/ctc/pay-ctc-order`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      getOrderList()
      Object.assign(detailOrder, res.data.data)
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const cancelOrder = () => {
  const params = { oid: detailOrder.id }
  axios.post(`${host}/uc/ctc/cancel-ctc-order`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      getOrderList()
      Object.assign(detailOrder, res.data.data)
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const init = () => {
  store.commit('navigate', 'nav-ctc')
  getC2cPrice()
  if (isLogin.value) {
    getOrderList()
    getAccount()
    getAccountSecurity()
    getWallet()
  }
}

onMounted(() => {
  init()
  timer.value = setInterval(() => {
    getC2cPrice()
  }, 30000)
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

<style>
.ctc .item-title {
  font-size: 20px;
  text-align: center;
  font-weight: bold;
  color: rgb(188, 188, 188);
}
.ctc .red {
  color: #f2334f;
}
.ctc .green {
  color: #45b854;
}
.ctc .item-title .unit {
  font-size: 14px;
}
.ctc .item-desc {
  font-size: 12px;
  text-align: center;
  color: #7c7f82;
}
.ctc .notice-bottom {
  margin-top: 5px;
  height: 55px;
  background-color: #192330;
  padding-top: 12px;
  color: rgb(42, 147, 255);
}
.ctc .notice-btn-left {
  height: 30px;
  line-height: 30px;
  width: 42%;
  margin-left: 5%;
  float: left;
  border-radius: 3px;
  border: 1px solid rgb(0, 116, 235);
}
.ctc .notice-btn-left:hover {
  cursor: pointer;
}
.ctc #sendCode {
  position: absolute;
  border: none;
  background: none;
  top: 6px;
  outline: none;
  right: 0;
  width: 30%;
  color: #f0ac19;
  cursor: pointer;
  height: 20px;
  line-height: 20px;
  border-left: 1px solid #dddee1;
}
.ctc .notice-btn-right {
  height: 30px;
  line-height: 30px;
  width: 42%;
  margin-right: 5%;
  float: right;
  border-radius: 3px;
  border: 1px solid rgb(0, 116, 235);
}
.ctc .notice-btn-right:hover {
  cursor: pointer;
}
.ctc .el-tabs__nav {
  border-bottom: 1px solid #323c53;
  font-size: 18px;
}
.ctc .el-tabs__item:hover {
  color: #f0a70a;
}
.ctc .el-tabs__item.is-active {
  color: #f0a70a;
  font-size: 18px;
}
.ctc .el-tabs__active-bar {
  background-color: #f0a70a;
}
.ctc .buy_total {
  border-top: 1px solid #323c53;
  padding-top: 30px;
  margin-bottom: 30px;
}
.ctc .trade_bd_ctc {
  width: 70%;
}
.ctc .trade_bd_ctc .panel {
  position: relative;
  z-index: 2;
  float: left;
  width: 49%;
  height: 455px;
  margin-top: 0;
  margin-right: 0;
  border: 0 solid transparent;
  padding-top: 35px;
}
.ctc .trade_panel {
  background: transparent !important;
}
.ctc .trade_panel .panel .hd {
  line-height: 20px;
  height: 20px;
  border-bottom: 1px solid #1F2943;
  margin-bottom: 5px;
}
.ctc .trade_panel .panel .hd span {
  padding-left: 0;
  font-size: 12px;
  margin: 0 3px;
  float: right;
}
.ctc-order-status {
  text-align: center;
  margin-bottom: 15px;
  background: #f0a70a;
  padding: 5px 0px;
  border-radius: 2px;
  color: #000000;
}
.ctc .trade_panel .panel .hd b {
  padding-left: 0;
  font-size: 12px;
  color: #7A98F7;
  float: right;
}
.ctc .trade_panel .panel.panel_buy {
  padding-right: 35px;
  padding-left: 35px;
  background: #192330;
}
.ctc .trade_panel .panel.panel_sell {
  padding-right: 35px;
  padding-left: 35px;
  background: #192330;
  margin-left: 5px;
}
.ctc .trade_wrap .buy-input .el-input__inner {
  color: #45b854;
  font-weight: bold;
  font-size: 22px;
  height: 40px;
}
.ctc .trade_wrap .sell-input .el-input__inner {
  color: #f2334f;
  font-weight: bold;
  font-size: 22px;
  height: 40px;
}
.ctc .trade_wrap .trade-input .el-input__inner {
  border: 1px solid #27313e;
  color: #fff;
  height: 40px;
  text-indent: 25px;
  border-radius: 0;
}
.ctc .trade_wrap .el-input-wrapper {
  outline: none;
}
.ctc .trade_wrap .el-input:focus,
.ctc .trade_wrap .el-input:hover {
  box-shadow: none;
  outline: none;
}
.ctc .trade_wrap .el-input-number-input:focus,
.ctc .trade_wrap .el-input-number-input:hover {
  box-shadow: none;
  border-color: #41546d;
  outline: none;
}
.ctc .trade_wrap .el-form-item-content input {
  padding-left: 50px;
  text-align: right;
  padding-right: 75px;
  font-size: 20px;
  line-height: 30px;
}
.ctc .trade_wrap .el-form-item-content input::placeholder {
  font-size: 12px;
  color: #515a6e;
  margin-bottom: 10px;
}
.ctc .trade_wrap .el-form-item-content label.before {
  position: absolute;
  top: 4px;
  left: 10px;
  color: #7c7f82;
  z-index: 2;
  font-size: 14px;
}
.ctc .trade_wrap .el-form-item-content label.after {
  position: absolute;
  top: 4px;
  right: 25px;
  color: #7c7f82;
  font-size: 14px;
}
.trade_bd_ctc .el-button {
  width: 100%;
  border: 0;
  color: #fff;
}
.trade_bd_ctc .el-button.bg-red {
  background-color: #f15057;
}
.trade_bd_ctc .el-button.bg-red:hover {
  background-color: #ff7278;
}
.trade_bd_ctc .el-button.bg-green {
  background-color: #00b275;
}
.trade_bd_ctc .el-button.bg-green:hover {
  background-color: #01ce88;
}
.trade_bd_ctc .el-button.bg-gray {
  background-color: #35475b;
  cursor: not-allowed;
  color: #9fabb5;
}
.trade_bd_ctc .el-button.bg-gray:hover {
  color: #9fabb5 !important;
}
.ctc .total {
  min-height: 90px;
}
@media screen and (max-width: 768px) {
  .shoujiIf {
    display: none;
  }
  .ctc .main {
    margin-top: 10px !important;
  }
  .ctc .trade_bd_ctc {
    width: 100%;
  }
  .ctc .trade_bd_ctc .panel {
    width: 100%;
  }
  .ctc .trade_panel .panel.panel_sell {
    margin-left: 0px;
    margin-top: 15px;
  }
  .ctc .trade_wrap .el-form-item-content input {
    padding-left: 0px;
  }
}
</style>
<style lang="scss" scoped>
.ctc {
  height: 100%;
  background-size: cover;
  position: relative;
  overflow: hidden;
  padding-bottom: 50px;
  padding-top: 60px;
  color: #fff;
}
.ctc .bannerimg {
  display: block;
  width: 100%;
}
.ctc_container {
  padding: 0 5%;
  text-align: center;
  height: 100%;
  > h1 {
    margin-top: -170px;
    font-size: 32px;
    line-height: 1;
    padding: 50px 0 20px 0;
    letter-spacing: 3px;
  }
}
.ctc .main {
  margin-top: 55px;
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  flex-wrap: wrap;
}
.ctc-container {
  min-height: 470px;
}
.bottom-panel {
  border-top: 1px solid rgb(237, 237, 237);
  margin-top: 15px;
  .bottom {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    span {
      font-size: 12px;
      color: #a7a7a7;
      margin-top: 15px;
    }
    button, a {
      margin-top: 11px;
    }
  }
}
.right {
  float: right;
}
.left {
  float: left;
}
.gray {
  color: #a7a7a7;
}
</style>

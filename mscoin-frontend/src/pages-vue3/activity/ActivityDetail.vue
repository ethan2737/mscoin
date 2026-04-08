<template>
  <div class="activity">
    <div class="activity_container">
      <el-row :gutter="20" style="min-height: 600px;">
        <el-col :xs="24" :sm="24" :md="16" :lg="16" style="min-height: 600px;margin-bottom:10px;">
          <div class="left-container">
            <div class="title">
              <img :src="activityDetail.smallImageUrl" alt="">
              <div class="title-text">
                <h1>{{activityDetail.title}}</h1>
                <p>{{activityDetail.detail}}</p>
              </div>
            </div>
            <div class="content-wrap">
              <div class="content">
                <div class="content-text" v-html="activityDetail.content"></div>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :md="8" :lg="8">
          <div class="right-container">
            <p class="base">{{$t('activity.baseinfo')}}</p>
            <div class="progress">
              <div class="progress-text" style="margin-bottom: -3px;">
                <span class="gray">{{$t('activity.progress')}}</span>
                <span class="gray">{{$t('activity.totalsupply')}}</span>
              </div>
              <el-progress :percent="fixedScale(activityDetail.progress, 2)" status="active" style="width: 100%;" :stroke-width="5" :show-text="false" />
              <div class="progress-text">
                <span>{{fixedScale(activityDetail.progress, 2)}}%</span>
                <span>{{fixedScale(activityDetail.totalSupply, activityDetail.amountScale)}} {{activityDetail.unit}}</span>
              </div>
            </div>
            <div class="info">
              <div class="info-item">
                <p class="title gray">{{$t('activity.status')}}</p>
                <p class="value" v-if="activityDetail.step == 0">{{$t('activity.prepare')}}</p>
                <p class="value" v-if="activityDetail.step == 1">{{$t('activity.ongoing')}}</p>
                <p class="value" v-if="activityDetail.step == 2">{{$t('activity.tokendistribute')}}</p>
                <p class="value" v-if="activityDetail.step == 3">{{$t('activity.completed')}}</p>
              </div>
              <div class="info-item">
                <p class="title gray">{{$t('activity.activitytype')}}</p>
                <p class="value" style="color:#F90;" v-if="activityDetail.type==1">
                  {{$t('activity.activitytype1')}}
                  <router-link to="/helpdetail?cate=1&id=33&cateTitle=常见问题" target="_blank">{{$t('activity.ruledetail')}}</router-link>
                </p>
                <p class="value" style="color:#F90;" v-if="activityDetail.type==2">
                  {{$t('activity.activitytype2')}}
                  <router-link to="/helpdetail?cate=1&id=32&cateTitle=常见问题" target="_blank">{{$t('activity.ruledetail')}}</router-link>
                </p>
                <p class="value" style="color:#F90;" v-if="activityDetail.type==3">
                  {{$t('activity.activitytype3')}}
                  <router-link to="/helpdetail?cate=1&id=36&cateTitle=常见问题" target="_blank">{{$t('activity.ruledetail')}}</router-link>
                </p>
                <p class="value" style="color:#F90;" v-if="activityDetail.type==4">
                  {{$t('activity.activitytype4')}}
                  <router-link to="/helpdetail?cate=1&id=37&cateTitle=常见问题" target="_blank">{{$t('activity.ruledetail')}}</router-link>
                </p>
                <p class="value" style="color:#F90;" v-if="activityDetail.type==5">{{$t('activity.activitytype5')}}</p>
              </div>
              <div class="info-item">
                <p class="title gray">{{$t('activity.totalsupply')}}</p>
                <p class="value">{{fixedScale(activityDetail.totalSupply, activityDetail.amountScale)}}<span style="font-size:12px;margin-left:5px;">{{activityDetail.unit}}</span></p>
              </div>
              <div class="info-item" v-if="activityDetail.type != 3">
                <p class="title gray">{{$t('activity.publishprice')}}</p>
                <p class="value">{{fixedScale(activityDetail.price, activityDetail.priceScale)}}<span style="font-size:12px;margin-left:5px;">{{activityDetail.acceptUnit}}</span></p>
              </div>
              <div class="info-item" v-if="activityDetail.type == 4 && activityDetail.discount==1">
                <p class="title gray">{{$t('activity.accountdiscount')}}</p>
                <span v-for="(item, key) in activityDetail.discounts" v-show="item.memberLevelId == myLevel" :key="key" style="margin-right:3px;">
                  V{{item.memberLevelId}}(<span style="color:#F90;">{{item.discount}}{{$t('activity.discount')}}</span>)
                </span>
              </div>
              <div class="info-item" v-if="activityDetail.type == 4 && activityDetail.power>0 && activityDetail.powerNum > 0">
                <p class="title gray">{{$t('activity.addpower')}}</p>
                <p class="value">{{activityDetail.power}}(KB/S)</p>
              </div>
              <div class="info-item" v-if="activityDetail.type != 5">
                <p class="title gray">{{$t('activity.activitycoin')}}</p>
                <p class="value">{{activityDetail.unit}}</p>
              </div>
              <div class="info-item">
                <p class="title gray" v-if="activityDetail.type!=3">{{$t('activity.acceptcoin')}}</p>
                <p class="title gray" v-if="activityDetail.type==3">{{$t('activity.holdcoin')}}</p>
                <p class="value">{{activityDetail.acceptUnit}}</p>
              </div>
              <div class="info-item">
                <p class="title gray">{{$t('activity.limittimes')}}</p>
                <p class="value" v-if="activityDetail.limitTimes > 0">{{activityDetail.limitTimes}}次</p>
                <p class="value" v-else>{{$t('activity.unlimited')}}</p>
              </div>
              <div class="info-item">
                <p class="title gray">{{$t('activity.limitamount')}}</p>
                <p class="value" v-if="activityDetail.minLimitAmout > 0 && activityDetail.maxLimitAmout > 0">
                  {{fixedScale(activityDetail.minLimitAmout, activityDetail.amountScale)}} ~ {{fixedScale(activityDetail.maxLimitAmout, activityDetail.amountScale)}}
                </p>
                <p class="value" v-if="activityDetail.minLimitAmout > 0 && activityDetail.maxLimitAmout == 0">
                  ≥ {{fixedScale(activityDetail.minLimitAmout, activityDetail.amountScale)}}
                </p>
                <p class="value" v-if="activityDetail.minLimitAmout == 0 && activityDetail.maxLimitAmout > 0">
                  ≤ {{fixedScale(activityDetail.maxLimitAmout, activityDetail.amountScale)}}
                </p>
                <p class="value" v-if="activityDetail.minLimitAmout == 0 && activityDetail.maxLimitAmout == 0">不限</p>
              </div>
              <div class="info-item" v-if="activityDetail.leveloneCount > 0">
                <p class="title gray">{{$t('activity.leveloneCount')}}</p>
                <p class="value" v-if="activityDetail.leveloneCount > 0">{{activityDetail.leveloneCount}}</p>
                <p class="value" v-else>{{$t('activity.unlimited')}}</p>
              </div>
              <div class="info-item">
                <p class="title gray">{{$t('activity.starttime')}}</p>
                <p class="value">{{dateFormat(activityDetail.startTime)}}</p>
              </div>
              <div class="info-item">
                <p class="title gray">{{$t('activity.endtime')}}</p>
                <p class="value">{{dateFormat(activityDetail.endTime)}}</p>
              </div>
            </div>
            <!-- 仅自由申购类型显示已兑/剩余数量 -->
            <div class="tips" v-if="activityDetail.type == 4 || activityDetail.type == 5">
              <div class="left-tip">
                <p class="title gray">{{$t('activity.alreadyamount')}}</p>
                <p>{{fixedScale(activityDetail.tradedAmount, activityDetail.amountScale)}} {{activityDetail.unit}}</p>
              </div>
              <div class="right-tip">
                <p class="title gray">{{$t('activity.leftamount')}}</p>
                <p>{{fixedScale(activityDetail.totalSupply - activityDetail.tradedAmount, activityDetail.amountScale)}} {{activityDetail.unit}}</p>
              </div>
            </div>
            <!-- 持仓瓜分类型显示冻结参与者 -->
            <div class="tips flexcolumn" v-if="activityDetail.type == 3">
              <div class="tipsline1">
                <div class="left-tip">
                  <p class="title gray">{{$t('activity.myalreadyholdamount')}}</p>
                  <p>{{fixedScale(myHoldAmount, activityDetail.amountScale)}} {{activityDetail.acceptUnit}}</p>
                </div>
                <div class="right-tip">
                  <p class="title gray">{{$t('activity.alreadyholdamount')}}</p>
                  <p>{{fixedScale(activityDetail.freezeAmount, activityDetail.amountScale)}} {{activityDetail.acceptUnit}}</p>
                </div>
              </div>
              <div class="tipsline2" v-if="isLogin">
                {{$t('activity.currentdivided')}}：{{holdPercent(myHoldAmount, activityDetail.freezeAmount, activityDetail.totalSupply)}} {{activityDetail.unit}}
              </div>
            </div>
            <p v-if="activityDetail.type == 3" class="hold-tips">*{{$t('activity.holdtips')}}</p>
            <div class="do-activity" v-if="isLogin && (activityDetail.type==3 || activityDetail.type==4 || activityDetail.type==5)">
              <div class="do-left">
                <p class="p-title" v-if="activityDetail.type==3">{{$t('activity.inputholdamount')}}</p>
                <p class="p-title" v-if="activityDetail.type==4">{{$t('activity.inputamount')}}</p>
                <p class="p-title" v-if="activityDetail.type==5">{{$t('activity.inputminingamount')}}</p>
                <el-input
                  v-model="attendAmount"
                  type="number"
                  placeholder="0.00000"
                  style="width: 90%;"
                >
                  <template #append v-if="activityDetail.type==3">{{activityDetail.acceptUnit}}</template>
                  <template #append v-else>{{activityDetail.unit}}</template>
                </el-input>
              </div>
              <div class="do-right">
                <p class="p-title">{{$t('activity.mybalance')}}</p>
                <div class="p-value">{{fixedScale(mybalance, activityDetail.priceScale)}} <span>{{activityDetail.acceptUnit}}</span></div>
              </div>
            </div>
            <div class="bottom">
              <div v-if="(activityDetail.type==3 || activityDetail.type==4 || activityDetail.type==5)">
                <el-button
                  type="warning"
                  size="large"
                  style="width: 100%;"
                  v-if="activityDetail.step==1 && isActivityInDate"
                  @click="apply"
                >
                  {{$t('activity.attend')}}
                </el-button>
                <el-button
                  type="warning"
                  size="large"
                  style="width: 100%;"
                  disabled
                  v-else
                >
                  {{$t('activity.attend')}}
                </el-button>
              </div>
              <div v-else>
                <el-button
                  type="warning"
                  size="large"
                  style="width: 100%;"
                  :disabled="screenWidth < 768"
                  v-if="activityDetail.step==1 && isActivityInDate"
                  @click="apply"
                >
                  <span>{{screenWidth > 768 ? $t('activity.attend') : $t('activity.login')}}</span>
                </el-button>
                <el-button
                  type="warning"
                  size="large"
                  style="width: 100%;"
                  disabled
                  v-else
                >
                  {{screenWidth > 768 ? $t('activity.attend') : $t('activity.login')}}
                </el-button>
              </div>
            </div>
            <p class="mobile-tip">{{$t('activity.tipsmobile1')}}</p>
            <img v-if="activityDetail.step==3" src="../../assets/images/completed.png" style="position: absolute;top:-10px;right:15px;width: 110px;" alt="">
          </div>
          <div class="memo">
            <p style="font-size:14px;margin-bottom:10px;">{{$t('activity.attention')}}：</p>
            <p>{{$t('activity.attentiontxt1')}}</p>
            <p>{{$t('activity.attentiontxt2')}}</p>
            <p>{{$t('activity.attentiontxt3')}}</p><br>
            <p>{{$t('activity.attentiontxt4')}}</p>
          </div>

          <router-link class="more-activity" to="/lab/">{{$t('activity.moreactivity')}}</router-link>
          <div class="show-qrcode" style="text-align: right;padding: 25px 10px;text-align:center;background:#FFF;margin-top: 10px;">
            <canvas ref="qrcodeCanvas"></canvas>
            <div style="width:100%;text-align:center;color:#000;">{{$t("cms.scanforshare")}}</div>
          </div>
        </el-col>
      </el-row>
    </div>

    <!-- modal -->
    <el-dialog v-model="modal2" width="360px" :show-close="false">
      <template #header>
        <div style="color:#f60;text-align:center;">
          <el-icon size="20" color="#00b5f6;"><InfoFilled /></el-icon>
          <span>{{$t('uc.finance.withdraw.safevalidate')}}</span>
        </div>
      </template>
      <div style="text-align:center">
        <el-form ref="formValidateAddr" :model="formValidateAddr" :rules="ruleValidate" :label-width="85">
          <!-- 手机号 -->
          <el-form-item :label="$t('uc.finance.withdraw.telno')" prop="mobileNo" v-show="validPhone">
            <el-input disabled size="large" v-model="formValidateAddr.mobileNo"></el-input>
          </el-form-item>
          <!-- 手机验证码 -->
          <el-form-item :label="$t('uc.finance.withdraw.smscode')" prop="vailCode2" v-show="validPhone">
            <el-input v-model="formValidateAddr.vailCode2" size="large">
              <template #append>
                <div class="timebox">
                  <el-button @click="send(2)" :disabled="disbtn">
                    <span v-if="sendMsgDisabled2">{{time2+$t('uc.finance.withdraw.second')}}</span>
                    <span v-if="!sendMsgDisabled2">{{$t('uc.finance.withdraw.clickget')}}</span>
                  </el-button>
                </div>
              </template>
            </el-input>
          </el-form-item>
          <!-- 邮箱 -->
          <el-form-item :label="$t('uc.finance.withdraw.email')" prop="email" v-show="validEmail">
            <el-input disabled v-model="formValidateAddr.email" size="large"></el-input>
          </el-form-item>
          <!-- 邮箱验证码 -->
          <el-form-item :label="$t('uc.finance.withdraw.emailcode')" prop="vailCode1" v-show="validEmail">
            <el-input v-model="formValidateAddr.vailCode1" size="large">
              <template #append>
                <div class="timebox">
                  <el-button @click="send(1)" :disabled="disbtn">
                    <span v-if="sendMsgDisabled1">{{time1+$t('uc.finance.withdraw.second')}}</span>
                    <span v-if="!sendMsgDisabled1">{{$t('uc.finance.withdraw.clickget')}}</span>
                  </el-button>
                </div>
              </template>
            </el-input>
          </el-form-item>
        </el-form>
      </div>
      <template #footer>
        <el-button type="primary" size="large" style="width: 100%;" @click="handleSubmit('formValidateAddr')">
          {{$t('activity.submit')}}
        </el-button>
      </template>
    </el-dialog>

    <div class="app_bottom">
      <div class="left_logo">
        <img style="float:left;" src="../../assets/images/applogo.png" alt="">
        <div style="float:left;height: 40px;line-height:40px;color:#000;">{{$t("cms.downloadslogan")}}</div>
      </div>
      <div class="right_btn_wrap"><router-link to="/app" class="right_btn">{{$t("cms.download")}}</router-link></div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 活动详情页面
 */
import { ref, reactive, computed, onMounted, nextTick } from 'vue'
import { ElMessage, ElRow, ElCol, ElProgress, ElInput, ElButton, ElForm, ElFormItem, ElDialog, ElIcon } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import axios from 'axios'
import { useRouter, useRoute, useStore } from 'vue-router/composables'
import QRCode from 'qrcode'

const router = useRouter()
const route = useRoute()
const store = useStore()

const host = 'http://localhost'

// 状态
const mybalance = ref(0)
const modal2 = ref(false)
const sendMsgDisabled1 = ref(false)
const sendMsgDisabled2 = ref(false)
const time1 = ref(60)
const time2 = ref(60)
const attendAmount = ref(0)
const myHoldAmount = ref(0)
const screenWidth = ref(window.screen.width)
const disbtn = ref(false)
const validEmail = ref(false)
const validPhone = ref(false)
const myRecordList = ref([])
const isActivityInDate = ref(false)
const myLevel = ref('')
const qrcodeCanvas = ref(null)

const activityDetail = reactive({
  title: '-',
  detail: '-',
  status: '1',
  step: '0',
  type: 1,
  startTime: '',
  endTime: '',
  totalSupply: 0,
  price: 0,
  priceScale: 0,
  unit: '-',
  acceptUnit: '-',
  amountScale: 0,
  maxLimitAmout: 0,
  minLimitAmout: 0,
  limitTimes: 0,
  activityLink: '-',
  noticeLink: '-',
  settings: '',
  smallImageUrl: '-',
  bannerImageUrl: '-',
  tradedAmount: 0,
  progress: 0,
  leveloneCount: 0,
  content: 'NONE',
  discount: 0,
  discounts: [],
  power: 0,
  powerNum: 0,
  freezeAmount: 0
})

const queryId = ref(0)

const formValidateAddr = reactive({
  mobileNo: '',
  vailCode2: '',
  email: '',
  vailCode1: ''
})

const ruleValidate = reactive({
  mobileNo: [{ required: false, message: '手机号不能为空', trigger: 'blur' }],
  vailCode2: [{ required: false, message: '验证码不能为空', trigger: 'change' }],
  email: [{ required: false, type: 'email', message: '邮箱格式不正确', trigger: 'blur' }],
  vailCode1: [{ required: false, message: '验证码不能为空', trigger: 'change' }]
})

// 计算属性
const lang = computed(() => store.state.lang)
const langPram = computed(() => store.state.lang === '简体中文' ? 'CN' : 'EN')
const isLogin = computed(() => store.getters.isLogin)

// 工具函数
const dateFormat = (tick) => {
  const moment = require('moment')
  return moment(tick).format('YYYY-MM-DD HH:mm:ss')
}

const fixedScale = (value, scale) => {
  if (value !== undefined && value !== null && value !== '') {
    return Number(value).toFixed(scale)
  }
  return 0
}

const holdPercent = (value, totalHold, totalSupply) => {
  if (value == 0) {
    return 0
  } else if (value > 0) {
    return (value / totalHold) * totalSupply
  }
  return 0
}

const generateQrcode = async () => {
  const value = host + '/lab/detail/' + queryId.value
  if (qrcodeCanvas.value) {
    await QRCode.toCanvas(qrcodeCanvas.value, value, { width: 200 })
  }
}

// 方法
const init = () => {
  store.commit('navigate', 'nav-activity')
  getData()
}

const getData = () => {
  const param = { id: route.params.id }
  axios.post(`${host}/uc/activity/detail`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(res => {
    if (res.data.code == 0) {
      Object.assign(activityDetail, res.data.data)
      if (isLogin.value) {
        getWallet()
        getMyRecords()
        getMember()
      }
      const currentDateStr = dateFormat(new Date().getTime())
      if (currentDateStr >= activityDetail.startTime && currentDateStr <= activityDetail.endTime) {
        isActivityInDate.value = true
      }
      nextTick(() => {
        generateQrcode()
      })
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const apply = () => {
  if (!isLogin.value) {
    ElMessage.error({ message: '请先登录', duration: 3000 })
    setTimeout(() => {
      router.push('/login')
    }, 3000)
  } else {
    attendClick()
  }
}

const attendClick = () => {
  const interval = setInterval(() => {
    if (time2.value <= 0) {
      sendMsgDisabled2.value = false
      clearInterval(interval)
      disbtn.value = false
    }
  }, 1000)

  if (!attendAmount.value) {
    if (activityDetail.type == 3) {
      ElMessage.error('请输入持仓数量')
    } else if (activityDetail.type == 4) {
      ElMessage.error('请输入数量')
    }
    return
  }

  if (parseFloat(attendAmount.value) == 0) {
    return
  }

  if (activityDetail.type == 5) {
    if (!/(^[1-9]\d*$)/.test(attendAmount.value)) {
      ElMessage.error('请输入整数')
      return
    }
  }

  var temAlreadyAttendAmount = 0
  var temAlreadyAttendTimes = myRecordList.value.length

  for (var i = 0; i < myRecordList.value.length; i++) {
    if (activityDetail.type == 3) {
      temAlreadyAttendAmount += myRecordList.value[i].freezeAmount
    } else if (activityDetail.type == 4) {
      temAlreadyAttendAmount += myRecordList.value[i].amount
    }
  }

  if (activityDetail.minLimitAmout > 0 && attendAmount.value < activityDetail.minLimitAmout) {
    ElMessage.error('最小起兑量不足')
    return
  }

  if (activityDetail.maxLimitAmout > 0 && ((temAlreadyAttendAmount + parseFloat(attendAmount.value)) > activityDetail.maxLimitAmout)) {
    ElMessage.error('超过最大兑换量')
    return
  }

  if (activityDetail.limitTimes > 0 && (temAlreadyAttendTimes + 1) > activityDetail.limitTimes) {
    ElMessage.error('超过限购次数')
    return
  }

  modal2.value = true
}

const handleSubmit = (name) => {
  submit(name)
}

const submit = (name) => {
  const param = {
    amount: attendAmount.value,
    activityId: route.params.id
  }
  if (validPhone.value) {
    param.aims = formValidateAddr.mobileNo
    param.code = formValidateAddr.vailCode2
  } else {
    param.aims = formValidateAddr.email
    param.code = formValidateAddr.vailCode1
  }

  axios.post(`${host}${'/uc/activity/attend'}`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      ElMessage.success(resp.message)
      modal2.value = false
      init()
      attendAmount.value = 0
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('提交失败')
  })
}

const getWallet = () => {
  axios.post(`${host}${'/uc/wallet/'}${activityDetail.acceptUnit}`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    mybalance.value = resp.data.balance || 0
  }).catch(() => {})
}

const getMyRecords = () => {
  const param = { activityId: route.params.id }
  axios.post(`${host}${'/uc/member/activity'}`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    myRecordList.value = resp.data || []
    myHoldAmount.value = 0
    if (activityDetail.type == 3) {
      for (var i = 0; i < myRecordList.value.length; i++) {
        myHoldAmount.value += myRecordList.value[i].freezeAmount
      }
    }
  }).catch(() => {})
}

const getMember = () => {
  axios.post(`${host}/uc/approve/security/setting`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      myLevel.value = resp.data.level.id
      if (resp.data.mobilePhone) {
        formValidateAddr.mobileNo = resp.data.mobilePhone
        validPhone.value = true
        validEmail.value = false
      } else {
        formValidateAddr.email = resp.data.email
        validPhone.value = false
        validEmail.value = true
      }
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {})
}

const send = (index) => {
  disbtn.value = true
  if (index == 1) {
    if (formValidateAddr.email) {
      axios.post(`${host}/uc/activity/code`, {}, {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json',
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      }).then(response => {
        const resp = response.data
        if (resp.code == 0) {
          ElMessage.success(resp.message)
          sendMsgDisabled1.value = true
          const interval = setInterval(() => {
            if (time1.value-- <= 0) {
              time1.value = 60
              sendMsgDisabled1.value = false
              clearInterval(interval)
              disbtn.value = false
            }
          }, 1000)
        } else {
          ElMessage.error(resp.message)
          disbtn.value = false
        }
      }).catch(() => {
        disbtn.value = false
      })
    } else {
      disbtn.value = false
    }
  } else if (index == 2) {
    if (formValidateAddr.mobileNo) {
      axios.post(`${host}/uc/mobile/activity/code`, {}, {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json',
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      }).then(response => {
        const resp = response.data
        if (resp.code == 0) {
          ElMessage.success(resp.message)
          sendMsgDisabled2.value = true
          const interval = setInterval(() => {
            if (time2.value-- <= 0) {
              time2.value = 60
              sendMsgDisabled2.value = false
              clearInterval(interval)
              disbtn.value = false
            }
          }, 1000)
        } else {
          ElMessage.error(resp.message)
          disbtn.value = false
        }
      }).catch(() => {
        disbtn.value = false
      })
    } else {
      disbtn.value = false
    }
  }
}

onMounted(() => {
  queryId.value = route.params.id
  init()
})
</script>

<style>
.ivu-progress-bg {
  border-radius: 0 !important;
  background-color: #ff8100;
  max-width: 100%;
}

.mobile-tip {
  text-align: center;
  color: #FF0000;
  margin-top: 10px;
  display: none;
  font-size: 10px;
}

.app_bottom {
  display: none;
  z-index: 999;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 40px;
  background: rgba(242, 246, 250, 1) !important;
}

.app_bottom .left_logo img {
  height: 30px;
  margin-top: 5px;
  border-radius: 5px;
  margin-left: 5px;
  margin-right: 5px;
}

.app_bottom .right_btn_wrap {
  float: right;
  height: 40px;
  line-height: 40px;
  margin-right: 5px;
}

.app_bottom .right_btn {
  color: #FFF;
  border-radius: 3px;
  padding: 0px 10px;
  line-height: 26px;
  height: 26px;
  display: block;
  margin-top: 7px;
  background: linear-gradient(200deg, #ff9900, #ffbe2b, #ffa500);
}

@media screen and (max-width: 768px) {
  .activity {
    padding-top: 45px !important;
  }

  .activity_container {
    padding: 0 3% !important;
  }

  .mobile-tip {
    display: block !important;
  }

  .app_bottom {
    display: block !important;
  }

  .activity .activity_container .left-container .title img {
    height: 50px !important;
    width: 70px !important;
    margin-top: 10px !important;
  }

  .activity .activity_container .left-container .title .title-text {
    padding-left: 10px !important;
  }

  .activity .activity_container .left-container .title .title-text h1 {
    font-size: 18px !important;
  }

  .activity .activity_container .left-container {
    min-height: 50px !important;
  }

  .activity .activity_container .right-container {
    min-height: 50px !important;
  }

  .more-activity {
    display: block !important;
  }
}
</style>

<style lang="scss" scoped>
.activity {
  background: rgba(242, 246, 250, 1) !important;
  height: 100%;
  background-size: cover;
  position: relative;
  overflow: hidden;
  padding-bottom: 30px;
  padding-top: 60px;
  color: #fff;

  .activity_container {
    padding: 0 12%;
    text-align: center;
    height: 100%;
    min-height: 600px;
    margin-top: 20px;

    .left-container {
      width: 100%;
      background: #FFF;
      min-height: 1600px;
      padding: 20px 20px;
      box-shadow: 0 0 25px #DDD;

      .title {
        padding-bottom: 20px;
        border-bottom: 1px solid #ededed;
        display: flex;
        flex-direction: row;
        justify-content: flex-start;

        img {
          height: 80px;
          width: 100px;
        }

        .title-text {
          padding-left: 20px;
          text-align: left;

          h1 {
            margin-top: 5px;
            font-size: 22px;
            font-weight: normal;
            color: #31363e !important;
          }

          p {
            margin-top: 10px;
            font-size: 12px;
            color: #a7a7a7;
          }
        }
      }

      .content-wrap {
        flex: 1 1 100%;
        width: 100%;
        background: #FFF;
        min-height: 600px;

        .content {
          padding: 20px 10px;
          color: #74777a;
          letter-spacing: 1px;
          font-size: 14px;
          text-align: left;
          line-height: 25px;

          .content-text {
            p {
              text-align: justify;

              &::after {
                width: 100%;
                content: '';
                height: 0;
              }
            }

            :deep(.image-desc) {
              img {
                width: 100% !important;
              }
            }
          }
        }
      }
    }

    .right-container {
      width: 100%;
      background: #FFF;
      min-height: 300px;
      padding: 20px 20px;
      color: #31363e !important;
      box-shadow: 0 0 25px #DDD;

      p.base {
        text-align: left;
        font-size: 18px;
        margin-bottom: 35px;
        color: #5d5d5d;
      }

      .progress {
        .progress-text {
          width: 100%;
          display: flex;
          flex-direction: row;
          justify-content: space-between;

          span {
            font-size: 12px;
          }
        }
      }

      .info {
        margin-top: 35px;

        .info-item {
          text-align: left;
          margin-bottom: 20px;

          .title {
            font-size: 13px;
            margin-bottom: 5px;
          }

          .value {
            font-size: 14px;
            letter-spacing: 1px;

            a {
              float: right;
              font-size: 12px;
              margin-top: 5px;
            }
          }
        }
      }

      .flexcolumn {
        display: flex;
        flex-direction: column !important;
      }

      p.hold-tips {
        font-size: 12px;
        text-align: right;
        margin-top: 10px;
        color: #F90;
      }

      .tips {
        background: rgba(240, 245, 249, 1) !important;
        padding: 15px 15px;
        border-radius: 3px;
        margin-top: 35px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;

        .left-tip {
          text-align: left;

          p.title {
            font-size: 12px;
          }

          p {
            font-size: 13px;
          }
        }

        .right-tip {
          text-align: right;

          p.title {
            font-size: 12px;
          }

          p {
            font-size: 13px;
          }
        }

        .tipsline1 {
          display: flex;
          flex-direction: row;
          justify-content: space-between;
          width: 100%;
        }

        .tipsline2 {
          padding-top: 10px;
          font-size: 12px;
          width: 100%;
          display: flex;
          flex-direction: row;
          justify-content: center;
          width: 100%;
          text-align: center;
          border-top: 1px solid #cedae3;
          margin-top: 10px;
        }
      }

      .bottom {
        width: 100%;
        margin-top: 30px;

        .el-button--warning[disabled] {
          background-color: #8c8c8c;
          border-color: #dcdee2;
        }

        button {
          border-radius: 0;
          padding: 10px 5px;
        }
      }

      .t-header {
        font-size: 12px;
        color: #727272;
        padding-bottom: 3px;
        background: #f0f5f9 !important;
        padding: 5px 0;
        margin-top: -20px;
      }

      .do-activity {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        margin-top: 30px;

        .do-left {
          width: 50%;
          text-align: left;

          .p-title {
            font-size: 12px;
            color: #aeb2b9 !important;
          }
        }

        .do-right {
          width: 50%;
          text-align: right;

          .p-title {
            font-size: 12px;
            color: #aeb2b9 !important;
          }

          .p-value {
            margin-top: 2px;
            line-height: 30px;
            height: 30px;
            letter-spacing: 1px;
            color: #000;
            padding-right: 5px;
            background: #ffd7a6;
            background-image: repeating-linear-gradient(60deg, hsla(0, 0%, 100%, .1), hsla(0, 0%, 100%, .1) 10px, transparent 0, transparent 20px);

            span {
              font-size: 12px;
            }
          }
        }
      }
    }

    .memo {
      text-align: left;
      font-size: 12px;
      color: #959595 !important;
      background: #fffded;
      padding: 20px 20px 30px 20px;
      border-top: 1px dashed #CCC;

      p {
        margin-top: 5px;
      }
    }

    .more-activity {
      width: 100%;
      margin-top: 20px;
      border-radius: 3px;
      background-color: #FFF;
      padding: 8px 0 8px 0;
      display: none;
      color: #f0a70a;
    }
  }
}

.gray {
  color: #aeb2b9 !important;
}

.do-activity {
  .el-input__inner,
  .el-input-number__input,
  .el-input-number {
    font-size: 14px;
    background-color: #FFF;
    color: #000;
    border-color: #FFF;
    border: 1px solid #666;
    border-radius: 0;

    &:hover {
      border-color: #FFF;
      border: 1px solid #666;
    }

    &:focus {
      border-color: #FFF;
      border: 1px solid #666;
      box-shadow: none;
    }
  }

  .el-input-group__append {
    background-color: #ffffff;
    border: 1px solid #666;
    border-left: none;
    border-radius: 0;
    font-size: 14px;
    color: #5f5f5f;
  }

  .el-input[disabled]:hover {
    border-color: #27313e;
  }
}

.timebox {
  display: flex;
  align-items: center;
}
</style>

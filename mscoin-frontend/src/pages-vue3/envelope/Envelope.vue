<template>
  <div class="envelope">
    <div v-if="hasInviteUser" style="width:80%;height: 35px;padding: 5px 0 5px 0px;border-radius: 35px;background-color:rgb(157, 0, 0);margin-left:10%;text-align:center;display: flex;flex-direction:row;justify-content:center;margin-bottom:10px;">
      <img style="width: 25px; height: 25px;margin-right: 10px;border-radius: 25px;" :src="inviteUserAvatar" alt="">
      <div style="height: 30px;line-height:30px;color: #EEE;">{{inviteUserId}} 邀您领红包！</div>
    </div>
    <div style="position:absolute;top: 40px;width: 100%;text-align:center;padding-top: 50px;">
      <div v-if="hasInviteUser" style="width:100%;height:35px;"></div>
      <p style="text-align:center;padding: 15px 0px;font-size:14px;letter-spacing:1px;width:100%;color:#000;">{{envelopeInfo.name}}</p>
      <p style="text-align:center;padding: 15px 0px;font-size:36px;letter-spacing:1px;width:100%;color:#fb272a;font-weight:bold;">{{envelopeInfo.totalAmount}} {{envelopeInfo.unit}}</p>
      <p style="text-align:center;padding: 15px 0px;font-size:14px;letter-spacing:1px;width:100%;color:rgb(227, 205, 187);">已领取：{{envelopeInfo.receiveCount}}/{{envelopeInfo.count}}</p>
      <p style="text-align:center;padding: 20px 0px;font-size:14px;letter-spacing:1px;width:100%;color:#555;margin-top: 40px;">
        <img style="width: 100px;height:100px;border-radius:60px;" :src="envelopeInfo.logo" alt="">
      </p>
    </div>
    <img style="width: 100%;" :src="envelopeInfo.bgImage" alt="">
    <p style="margin-top: -80px;margin-bottom: 80px;text-align:center;color: rgb(255, 136, 79);font-size:13px;">该红包通过 MSZLU.COM 发出</p>
    <div class="input-panel" v-if="!hasReceived && envelopeInfo.state == 0">
      <div style="color: rgb(177, 177, 177);font-size: 16px;margin: 10px 0 20px 0;" v-html="envelopeInfo.detail"></div>
      <el-form ref="formInlineRef" inline>
        <el-form-item prop="user">
          <el-input type="text" v-model="formInline.phone" placeholder="请输入手机号码"></el-input>
        </el-form-item>

        <el-form-item prop="code">
          <el-input type="text" v-model="formInline.verifyCode" :placeholder="$t('uc.regist.smscode')">
          </el-input>
          <input id="sendCode" @click="sendCode" type="button" shape="circle" :value="sendcodeValue" :disabled='codedisabled'>
          </input>
        </el-form-item>
        <el-form-item>
          <el-button class="register_btn" @click="handleSubmit">领取红包</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="envelope-result" v-if="hasReceived && envelopeInfo.state == 0">
      <p style="font-size:14px;text-align:center;color: #999;margin-top: 5px;">恭喜您！</p>
      <p style="text-align:center;font-size: 30px;color: rgb(251, 39, 42);font-weight:bold;margin: 10px 0;">{{receiveAmount}} BTC</p>
      <el-button v-if="envelopeInfo.invite == 1" class="register_btn" @click="inviteMore">邀请好友增加领取次数</el-button>
      <p v-if="envelopeInfo.invite ==1" style="font-size:14px;text-align:center;color: #999;margin-top:15px;">邀请一个好友增加一次领取机会</p>
      <p style="text-align:center;">
        <router-link style="font-size:14px;text-align:center;color: #085fff;margin-top:15px;text-decoration:underline;" to="/app">下载 APP | 查看资产</router-link>
      </p>
    </div>
    <div class="envelope-result" v-if="envelopeInfo.state == 1">
      <p style="font-size:20px;text-align:center;color: #999;margin-top: 5px;">红包已领完！</p>
    </div>
    <div class="envelope-result" v-if="envelopeInfo.state == 2">
      <p style="font-size:20px;text-align:center;color: #999;margin-top: 5px;">红包已过期！</p>
    </div>

    <div class="record-list">
      <div class="title">领取记录</div>

      <div class="content">
        <p v-if="envelopeDetailList.length == 0" style="color: #999;margin: 20px 0;">暂无领取</p>
        <div class="item clearfix" v-for="item in envelopeDetailList" :key="item.id">
          <div class="phone">{{item.userIdentify}}</div><div class="amount">{{item.amount}} {{envelopeInfo.unit}}</div>
        </div>
      </div>
    </div>
    <p style="text-align:center;margin-top: 25px;margin-bottom: 20px;">
    <router-link style="font-size:14px;text-align:center;color: #EEE;margin-top:15px;text-decoration:underline;" to="/app">© MSZLU.COM | 下载 APP</router-link>
    </p>

    <div class="loading-overlay" v-if="spinShow">
      <el-icon class="is-loading"><Loading /></el-icon>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 红包信封页面
 */
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElForm, ElFormItem, ElInput, ElButton, ElIcon } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const store = useStore()

const host = ''

const spinShow = ref(false)
const hasReceived = ref(false)
const inviteLink = ref('https://www.MSZLU.COM/envelope/')
const envelopeNo = ref('')
const country = ref('')
const sendcodeValue = ref('发送验证码')
const codedisabled = ref(false)
const receiveAmount = ref('0.00')
const promotionCode = ref('')
const hasInviteUser = ref(false)
const inviteUserId = ref('***********')
const inviteUserAvatar = ref('https://kicks.oss-cn-shanghai.aliyuncs.com/defaultavatar.png')

const envelopeInfo = reactive({
  id: 0,
  name: '**********',
  totalAmount: '0.00',
  unit: '',
  receiveCount: 0,
  count: 0,
  invite: 0,
  type: 0,
  logo: 'https://kicks.oss-cn-shanghai.aliyuncs.com/applogo.png',
  bgImage: 'https://kicks.oss-cn-shanghai.aliyuncs.com/redenvelope.png',
  state: 0,
  detail: '留给未来一个暴富的可能'
})

const formInline = reactive({
  verifyCode: '',
  phone: '',
  promotionCode: '',
  envelopeNo: ''
})

const queryDetail = reactive({
  envelopeId: '',
  pageNo: 0,
  pageSize: 30
})

const envelopeDetailList = ref([])
const countdown = ref(60)

const lang = computed(() => store.state.lang)
const langPram = computed(() => {
  if (store.state.lang === '简体中文') return 'CN'
  if (store.state.lang === 'English') return 'EN'
  return 'CN'
})

const sendCode = () => {
  const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
  if (!reg.test(formInline.phone)) {
    ElMessage.error('请输入合法的手机号')
  } else {
    const params = {
      phone: formInline.phone,
      country: '中国',
      envelopeId: envelopeInfo.id
    }
    if (envelopeInfo.id === 0 || envelopeInfo.id === null || envelopeInfo.id === undefined) {
      ElMessage.error('红包不存在')
      return
    }
    axios.post(`${host}/uc/redenvelope/code`, params, {
      headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
    }).then(res => {
      if (res.data.code === 0) {
        ElMessage.success(res.data.message)
        settime()
      } else {
        ElMessage.error(res.data.message)
      }
    }).catch(() => {})
  }
}

const settime = () => {
  sendcodeValue.value = '重新发送 (' + countdown.value + ')'
  codedisabled.value = true
  const timercode = setInterval(() => {
    countdown.value--
    sendcodeValue.value = '重新发送 (' + countdown.value + ')'
    if (countdown.value <= 0) {
      clearInterval(timercode)
      codedisabled.value = false
      sendcodeValue.value = '发送验证码'
      countdown.value = 60
    }
  }, 1000)
}

const handleSubmit = () => {
  if (formInline.verifyCode === '') {
    ElMessage.error('请输入验证码')
    return
  }

  axios.post(`${host}/uc/redenvelope/receive`, formInline, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      promotionCode.value = res.data.data.promotionCode
      receiveAmount.value = res.data.data.amount
      hasReceived.value = true
      getEnvelopeDetailList()
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const inviteMore = () => {
  router.replace('/envelope/' + formInline.envelopeNo + '?code=' + promotionCode.value)
  ElMessageBox.confirm('点击"确定"将生成您的专属红包邀请页面。</p><br><p>使用方法：进入专属邀请页面后，分享该页面给微信好友即可', '跳转刷新提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    formInline.promotionCode = promotionCode.value
    hasReceived.value = false
    getEnvelope()
  }).catch(() => {})
}

const getEnvelope = () => {
  const param = {
    envelopeNo: formInline.envelopeNo,
    code: formInline.promotionCode
  }

  spinShow.value = true
  axios.post(`${host}/uc/redenvelope/query`, param, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    spinShow.value = false
    if (res.data.code === 0) {
      envelopeInfo.id = res.data.data.id
      envelopeInfo.name = res.data.data.name
      envelopeInfo.totalAmount = res.data.data.totalAmount
      envelopeInfo.receiveCount = res.data.data.receiveCount
      envelopeInfo.count = res.data.data.count
      envelopeInfo.detail = res.data.data.detail
      envelopeInfo.unit = res.data.data.unit
      envelopeInfo.invite = res.data.data.invite
      envelopeInfo.type = res.data.data.type
      envelopeInfo.state = res.data.data.state
      if (res.data.data.inviteUser !== null && res.data.data.inviteUser !== '' && res.data.data.inviteUser !== undefined) {
        hasInviteUser.value = true
        inviteUserId.value = res.data.data.inviteUser
        if (res.data.data.inviteUserAvatar !== null) {
          inviteUserAvatar.value = res.data.data.inviteUserAvatar
        }
      }

      if (res.data.data.logoImage !== null && res.data.data.logoImage !== '') {
        envelopeInfo.logo = res.data.data.logoImage
      }
      if (res.data.data.bgImage !== null && res.data.data.bgImage !== '') {
        envelopeInfo.bgImage = res.data.data.bgImage
      }

      document.title = '【' + envelopeInfo.totalAmount + ' ' + envelopeInfo.unit + '】' + envelopeInfo.name + ' — MSZLU.COM 交易所'
      getEnvelopeDetailList()
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const getEnvelopeDetailList = () => {
  queryDetail.envelopeId = envelopeInfo.id
  axios.post(`${host}/uc/redenvelope/query-detail`, queryDetail, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      envelopeDetailList.value = res.data.data.content
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const init = () => {
  store.state.HeaderActiveName = '1-8'
  store.commit('navigate', 'nav-envelope')

  const eNo = route.params.eno

  if (eNo === undefined) {
    router.push('/app')
  } else {
    formInline.envelopeNo = eNo
    if (route.query.code !== undefined && route.query.code !== '' && route.query.code !== null) {
      formInline.promotionCode = route.query.code
    }
  }

  getEnvelope()
}

onMounted(() => {
  init()
})
</script>

<style lang="scss" scoped>
.envelope {
  padding-top: 60px;
  position: relative;

  .loading-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(255, 255, 255, 0.8);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
  }

  .envelope-result {
    width: 90%;
    margin-left: 5%;
    background: #FFF;
    border-radius: 5px;
    position: relative;
    padding: 15px 20px 25px 20px;

    .register_btn {
      width: 100%;
      background-color: #f0ac19;
      outline: none;
      border-color: #f0ac19;
      color: #fff;
      border-radius: 5px;
      font-size: 16px;
      margin-top: 10px;

      &:focus {
        box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
      }
    }
  }

  .input-panel {
    width: 90%;
    margin-left: 5%;
    background: #FFF;
    border-radius: 5px;
    position: relative;
    padding: 15px 20px 5px 20px;

    .el-form-item {
      margin-bottom: 15px;

      .el-form-item-content {
        .register_btn {
          width: 100%;
          background-color: #f0ac19;
          outline: none;
          border-color: #f0ac19;
          color: #fff;
          border-radius: 5px;
          font-size: 16px;
          margin-top: 10px;

          &:focus {
            box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
          }
        }

        .el-input__inner {
          background-color: #EDEDED !important;
          border: none !important;
          color: #333;
        }

        .el-input:hover {
          border: none;
        }

        #sendCode {
          position: absolute;
          border: none;
          background: transparent;
          top: 1px;
          outline: none;
          right: 0;
          width: 40%;
          color: #f05e19;
          cursor: pointer;
        }
      }
    }
  }

  .record-list {
    margin-top: 30px;
    text-align: center;

    .title {
      position: relative;
      width: 40%;
      margin-left: 30%;
      font-size: 16px;

      &:before {
        position: absolute;
        left: -30%;
        top: 12px;
        content: '';
        height: 1px;
        width: 50%;
        display: block;
        background-color: #FFF;
      }

      &:after {
        position: absolute;
        right: -30%;
        top: 12px;
        content: '';
        height: 1px;
        width: 50%;
        display: block;
        background-color: #FFF;
      }
    }

    .content {
      background-color: #FFF;
      width: 90%;
      border-radius: 5px;
      margin-left: 5%;
      margin-top: 20px;
      padding: 10px 5px;

      .item {
        color: #888;
        padding: 8px 10px;

        .phone {
          float: left;
          letter-spacing: 1px;
        }

        .amount {
          float: right;
        }
      }
    }
  }
}

.clearfix:after {
  visibility: hidden;
  display: block;
  font-size: 0;
  content: ' ';
  clear: both;
  height: 0;
}
</style>

<template>
  <div class="login_form mob-login">
    <div class="login_right">
      <div style="color: #F90;margin-bottom: 60px;margin-top: 60px;">
        <h1 style="border-left: 5px solid #F90;padding-left: 10px;line-height:30px;height:30px;">MSZLU.COM</h1>
        <p style="padding-left:15px;letter-spacing: 6px;">{{$t('footer.gsmc')}}</p>
        <div style="margin-left:5px;letter-spacing: 2px;margin-top: 10px;color: rgb(226, 226, 227);font-size:13px;padding: 5px 10px;">安全 ● 诚实 ● 公平 ● 热情 ● 开放</div>
      </div>
      <el-form v-if="allowRegister" ref="formInlineRef" :model="formInline" :rules="ruleInline" inline>
        <el-form-item prop="username" style="display:none;">
          <el-input type="text" v-model="formInline.username" :placeholder="$t('uc.regist.username')"></el-input>
        </el-form-item>
        <el-form-item prop="user">
          <el-input type="text" v-model="formInline.user" :placeholder="key">
            <template #prepend>
              <el-select v-model="country" style="width: 65px;border-bottom: 1px solid #27313e;">
                <el-option value="中国" label="+86"><span>+86</span><span style="margin-left:10px;color:#ccc">中国</span></el-option>
                <el-option value="新加坡" label="+65"><span>+65</span><span style="margin-left:10px;color:#ccc">新加坡</span></el-option>
                <el-option value="韩国" label="+82"><span>+82</span><span style="margin-left:10px;color:#ccc">韩国</span></el-option>
                <el-option value="日本" label="+81"><span>+81</span><span style="margin-left:10px;color:#ccc">日本</span></el-option>
                <el-option value="泰国" label="+66"><span>+66</span><span style="margin-left:10px;color:#ccc">泰国</span></el-option>
                <el-option value="俄罗斯" label="+7"><span>+7</span><span style="margin-left:10px;color:#ccc">俄罗斯</span></el-option>
                <el-option value="英国" label="+44"><span>+44</span><span style="margin-left:10px;color:#ccc">英国</span></el-option>
                <el-option value="越南" label="+84"><span>+84</span><span style="margin-left:10px;color:#ccc">越南</span></el-option>
                <el-option value="印度" label="+91"><span>+91</span><span style="margin-left:10px;color:#ccc">印度</span></el-option>
                <el-option value="意大利" label="+39"><span>+39</span><span style="margin-left:10px;color:#ccc">意大利</span></el-option>
                <el-option value="香港" label="+852"><span>+852</span><span style="margin-left:10px;color:#ccc">香港</span></el-option>
                <el-option value="马来西亚" label="+60"><span>+60</span><span style="margin-left:10px;color:#ccc">马来西亚</span></el-option>
                <el-option value="台湾省" label="+886"><span>+886</span><span style="margin-left:10px;color:#ccc">台湾省</span></el-option>
                <el-option value="土耳其" label="+90"><span>+90</span><span style="margin-left:10px;color:#ccc">土耳其</span></el-option>
              </el-select>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item prop="code" v-show="showCode">
          <el-input type="text" v-model="formInline.code" :placeholder="$t('uc.regist.smscode')"></el-input>
          <input id="sendCode" @click="sendCode" type="button" shape="circle" :value="sendcodeValue" :disabled='codedisabled'>
          </input>
        </el-form-item>
        <el-form-item prop="password" class="password">
          <el-input type="password" v-model="formInline.password" :placeholder="$t('uc.regist.pwd')"></el-input>
        </el-form-item>
        <el-form-item prop="repassword" class="password">
          <el-input type="password" v-model="formInline.repassword" :placeholder="$t('uc.regist.repwd')"></el-input>
        </el-form-item>
        <el-form-item prop="promotion">
          <el-input type="text" v-model="formInline.promotion">
            <template #prepend>{{$t('uc.regist.promotion')}} :</template>
          </el-input>
        </el-form-item>
        <div class="check-agree" style="">
          <label>
            <el-checkbox v-model="agree">{{$t('uc.regist.agreement')}}</el-checkbox>
          </label>
          <a v-if="lang=='简体中文'" href="/helpdetail?cate=1&id=5&cateTitle=常见问题" target="_blank" style="">《{{$t('uc.regist.userprotocol')}}》</a>
          <a v-if="lang=='English'" href="/helpdetail?cate=1&id=35&cateTitle=Privacy Policy" target="_blank" style="">《{{$t('uc.regist.userprotocol')}}》</a>
        </div>
        <el-form-item>
          <el-button class="register_btn" @click="handleSubmit" :disabled="registing">{{$t('uc.regist.regist')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="section" id="page4">
      <ul>
        <li>
          <div><img src="../../assets/images/feature_safe.png" alt=""></div>
          <p class="title">{{$t('description.title1')}}</p>
          <p>{{$t('description.message1')}}</p>
        </li>
        <li>
          <div><img src="../../assets/images/feature_fast.png" alt=""></div>
          <p class="title">{{$t('description.title2')}}</p>
          <p>{{$t('description.message2')}}</p>
        </li>
        <li>
          <div><img src="../../assets/images/feature_global.png" alt=""></div>
          <p class="title">{{$t('description.title3')}}</p>
          <p>{{$t('description.message3')}}</p>
        </li>
        <li>
          <div><img src="../../assets/images/feature_choose.png" alt=""></div>
          <p class="title">{{$t('description.title4')}}</p>
          <p>{{$t('description.message4')}}</p>
        </li>
      </ul>
    </div>
    <div class="app_bottom_reg">
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
 * Vue 3 迁移 - 手机注册页面
 */
import { ref, reactive, computed, watch, onMounted, nextTick } from 'vue'
import { ElMessage, ElNotification, ElForm, ElFormItem, ElInput, ElSelect, ElOption, ElButton, ElCheckbox } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const store = useStore()

const host = ''

const country = ref('中国')
const codedisabled = ref(false)
const sendcodeValue = ref('发送验证码')
const isRegister = ref(false)
const ticket = ref('')
const randStr = ref('')
const registing = ref(false)
const captchaObj = ref(null)
const modal1 = ref(false)
const _captchaResult = ref(null)
const agree = ref(true)
const allowRegister = ref(true)
const areas = ref([])
const changeActive = ref(0)
const showCode = ref(true)
const countdown = ref(60)

const formInline = reactive({
  username: '',
  user: '',
  code: '',
  areaCode: '',
  password: '',
  repassword: '',
  promotion: ''
})

const ruleInline = reactive({})

const key = ref('')
const code = ref('')

const lang = computed(() => store.state.lang)
const isLogin = computed(() => store.getters.isLogin)

const validateUser = (rule, value, callback) => {
  if (changeActive.value === 0) {
    const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
    if (value === '') {
      callback(new Error('请输入手机号'))
    } else if (!reg.test(formInline.user)) {
      callback(new Error('手机号格式错误'))
    } else {
      callback()
    }
  } else {
    const reg = /^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$/
    if (value === '') {
      callback(new Error('请输入邮箱'))
    } else if (!reg.test(formInline.user)) {
      callback(new Error('邮箱格式错误'))
    } else {
      callback()
    }
  }
}

const validateRepassword = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请确认密码'))
  } else if (value !== formInline.password) {
    callback(new Error('两次密码不一致'))
  } else {
    callback()
  }
}

const initRules = () => {
  ruleInline.user = [{ validator: validateUser, trigger: 'blur' }]
  ruleInline.code = [
    {
      required: true,
      message: '请输入短信验证码',
      trigger: 'blur'
    }
  ]
  ruleInline.password = [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    },
    {
      type: 'string',
      min: 6,
      message: '密码至少 6 位',
      trigger: 'blur'
    }
  ]
  ruleInline.repassword = [{ validator: validateRepassword, trigger: 'blur' }]
}

watch(changeActive, (val) => {
  if (val === 0) {
    showCode.value = true
    key.value = '手机号'
    ruleInline.code = [
      {
        required: true,
        message: '请输入短信验证码',
        trigger: 'blur'
      }
    ]
  } else {
    showCode.value = false
    key.value = '邮箱'
    ruleInline.code = []
  }
})

watch(lang, () => {
  updateLangData()
})

const updateLangData = () => {
  if (changeActive.value === 0) {
    key.value = '手机号'
  } else {
    key.value = '邮箱'
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

const sendCode = () => {
  const mobilePhone = formInline.user
  const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
  if (mobilePhone === '' || !reg.test(mobilePhone)) {
    ElMessage.error('请输入正确的手机号')
    return
  } else {
    initGtCaptcha()
  }
}

const initGtCaptcha = () => {
  const self = this
  const captcha1 = { show: () => {} }
  captcha1.show()
  success()
}

const success = () => {
  const params = {
    phone: formInline.user,
    country: '中国'
  }
  const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
  if (reg.test(params.phone)) {
    axios.post(`${host}/uc/mobile/code`, params, {
      headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
    }).then(res => {
      if (res.data.code === 0) {
        ElNotification.success({ title: '提示', message: res.data.message })
        settime()
      } else {
        ElNotification.error({ title: '提示', message: res.data.message })
      }
    }).catch(() => {})
  } else {
    ElNotification.error({ title: '提示', message: '手机号格式错误' })
  }
}

const handleSubmit = () => {
  if (!agree.value) {
    ElNotification.error({ title: '提示', message: '请同意用户协议' })
    return
  }

  if (changeActive.value === 1) {
    openValidateModal()
  } else {
    if (isRegister.value) {
      registing.value = true
      const params = {
        phone: formInline.user,
        username: formInline.username + formInline.user,
        password: formInline.password,
        promotion: formInline.promotion,
        code: formInline.code,
        country: country.value,
        ticket: ticket.value,
        randStr: randStr.value
      }

      axios.post(`${host}/uc/register/phone`, params, {
        headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
      }).then(res => {
        registing.value = false
        if (res.data.code === 0) {
          ElNotification.success({ title: '提示', message: res.data.message })
          setTimeout(() => {
            router.push('/app')
          }, 3000)
        } else {
          ElNotification.error({ title: '提示', message: res.data.message })
        }
      }).catch(() => {
        registing.value = false
      })
    } else {
      ElNotification.error({ title: '提示', message: '请输入正确的验证码' })
    }
  }
}

const openValidateModal = () => {
  ElMessage.info('邮箱注册功能暂未开放')
}

const actives = (index) => {
  changeActive.value = index
}

const init = () => {
  store.commit('navigate', 'nav-other')
  store.state.HeaderActiveName = '0'
  if (isLogin.value) {
    router.push('/')
  }
  document.title = (lang.value === '简体中文' ? '新用户注册 - ' : 'New Register - ') + 'MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
}

onMounted(() => {
  window.scrollTo(0, 0)
  init()
  initRules()
  actives(changeActive.value)
  if (route.query.code !== undefined && route.query.code !== '' && route.query.code !== null) {
    formInline.promotion = route.query.code
  } else {
    formInline.promotion = ''
  }
})
</script>

<style scoped lang="scss">
#page4 {
  background: transparent;
  padding: 80px 0 80px 0;
  ul {
    width: 99%;
    margin: 0 auto;
    li {
      flex: 0 0 25%;
      display: inline-block;
      width: 100%;
      padding: 0 15px;
      div {
        width: 130px;
        height: 130px;
        border-radius: 50%;
        vertical-align: middle;
        text-align: center;
        margin: 0 auto;
        img {
          height: 125px;
          margin-top: 8px;
        }
      }
      p {
        font-size: 14px;
        margin: 20px 0;
        text-align: center;
        color: #828ea1;
      }
      p.title {
        color: #fff;
        font-size: 18px;
        font-weight: 400;
      }
    }
  }
}

.login_form {
  position: relative;
  .login_right {
    padding: 20px 30px;
    background: transparent;
    width: 100%;
    padding-top: 60px;
    border-radius: 5px;
    .tel-title {
      color: #fff;
      text-align: left;
      margin-bottom: 30px;
    }
    .el-form-item {
      .el-form-item-content {
        .register_btn {
          width: 100%;
          background-color: #f0ac19;
          outline: none;
          border-color: #f0ac19;
          color: #fff;
          border-radius: 5px;
          font-size: 18px;
          margin-top: 20px;
          &:focus {
            box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
          }
        }
        .el-input__inner {
          border: 1px solid red;
        }
        #sendCode {
          position: absolute;
          border: 1px solid #0b1520;
          background: transparent;
          top: -10px;
          outline: none;
          right: 0;
          width: 30%;
          color: #f0ac19;
          cursor: pointer;
        }
      }
    }
    .check-agree {
      color: #979797;
      display: inline-block;
      line-height: 30px;
      font-size: 12px;
      cursor: default;
      a {
        color: #f0ac19;
        margin-left: -10px;
      }
    }
  }
}

.login_title {
  text-align: center;
  height: 80px;
  font-size: 25px;
}

#captcha {
  width: 100%;
  display: inline-block;
}

.show {
  display: block;
}

.hide {
  display: none;
}

#notice {
  color: red;
}

#wait {
  text-align: left;
  color: #666;
  margin: 0;
}

.tel-title {
  font-size: 25px;
}

.login_left {
  display: none;
}
</style>

<style lang="scss">
.login_form {
  .login_right {
    .el-form-item {
      .el-form-item-content {
        .el-input__inner {
          border: none;
          border-bottom: 1px solid #27313e;
          font-size: 14px;
          background: transparent;
          border-radius: 0;
          &:focus {
            border: none;
            border-bottom: 1px solid #27313e;
            box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
          }
        }
      }
    }
    .check-agree {
      .el-checkbox__input {
        &:focus {
          border: none;
          outline: none;
          box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
        }
      }
      .is-checked {
        .el-checkbox__inner {
          border: 1px solid #f0ac19;
          background-color: #f0ac19;
        }
      }
      .el-checkbox.is-indeterminate {
        .el-checkbox__inner {
          background: transparent;
        }
      }
    }
  }
}
</style>

<style>
.app_bottom_reg {
  z-index: 999;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 40px;
  background: rgba(242,246,250,1) !important;
}
.app_bottom_reg .left_logo img {
  height: 30px;
  margin-top: 5px;
  border-radius: 5px;
  margin-left: 5px;
  margin-right: 5px;
}
.app_bottom_reg .right_btn_wrap {
  float: right;
  height: 40px;
  line-height: 40px;
  margin-right: 5px;
}
.app_bottom_reg .right_btn {
  color: #FFF;
  border-radius: 3px;
  padding: 0px 10px;
  line-height: 26px;
  height: 26px;
  display: block;
  margin-top: 7px;
  background: linear-gradient(200deg, #ff9900, #ffbe2b, #ffa500);
}
</style>

<template>
  <div class="login_form">
    <div class="login_right">
      <el-form ref="formInlineRef" :model="formInline" :rules="ruleInline" inline>
        <el-form-item style="text-align:center;">
          <div class="tel-title">{{$t('uc.login.getlostpwd')}}</div>
        </el-form-item>
        <el-form-item prop="user">
          <el-input type="text" v-model="formInline.user" :placeholder="$t('uc.login.usertip')"></el-input>
        </el-form-item>
        <el-form-item prop="code">
          <el-input type="text" v-model="formInline.code" :placeholder="$t('uc.regist.smscode')"></el-input>
          <input id="sendCode" @click="sendCode" type="button" :value="sendcodeValue" :disabled="codedisabled">
          </input>
        </el-form-item>
        <el-form-item prop="password" class="password">
          <el-input type="password" v-model="formInline.password" :placeholder="$t('uc.forget.newpwd')"></el-input>
        </el-form-item>
        <el-form-item prop="repassword" class="password">
          <el-input type="password" v-model="formInline.repassword" :placeholder="$t('uc.forget.confirmpwd')"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="login_btn" @click="handleSubmit">{{$t('uc.forget.save')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 找回密码页面
 */
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElNotification, ElForm, ElFormItem, ElInput, ElButton } from 'element-plus'
import axios from 'axios'
import { useRouter, useStore } from 'vue-router/composables'

const router = useRouter()
const store = useStore()

const host = 'http://localhost'

const codedisabled = ref(false)
const sendcodeValue = ref('发送验证码')
const captchaObj = ref(null)
const modal1 = ref(false)
const _captchaResult = ref(null)
const changeActive = ref(0)
const countdown = ref(60)

const formInline = reactive({
  user: '',
  code: '',
  password: '',
  repassword: ''
})

const ruleInline = reactive({})

const isLogin = computed(() => store.getters.isLogin)

const validateUser = (rule, value, callback) => {
  const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
  if (value === '') {
    callback(new Error('请输入手机号'))
  } else if (!reg.test(formInline.user)) {
    callback(new Error('手机号格式错误'))
  } else {
    callback()
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
    { required: true, message: '请输入验证码', trigger: 'blur' }
  ]
  ruleInline.password = [
    {
      required: true,
      message: '请输入新密码',
      trigger: 'blur'
    },
    {
      type: 'string',
      min: 6,
      message: '密码至少 6 位',
      trigger: 'blur'
    }
  ]
  ruleInline.repassword = [
    { validator: validateRepassword, trigger: 'blur' },
    {
      type: 'string',
      min: 6,
      message: '密码至少 6 位',
      trigger: 'blur'
    }
  ]
}

const settime = () => {
  sendcodeValue.value = countdown.value
  codedisabled.value = true
  const timercode = setInterval(() => {
    countdown.value--
    sendcodeValue.value = countdown.value
    if (countdown.value <= 0) {
      clearInterval(timercode)
      codedisabled.value = false
      sendcodeValue.value = '发送验证码'
      countdown.value = 60
    }
  }, 1000)
}

const sendCode = () => {
  const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
  const tel = formInline.user
  const flagtel = reg.test(tel)
  if (flagtel) {
    afterValidate()
  } else {
    ElMessage.error('请填写正确的手机号')
  }
}

const afterValidate = () => {
  modal1.value = false
  const params = {
    account: formInline.user
  }
  axios.post(`${host}/uc/mobile/reset/code`, params, {
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

const handleSubmit = () => {
  const params = {
    account: formInline.user,
    code: formInline.code,
    mode: 0,
    password: formInline.password
  }
  axios.post(`${host}/uc/reset/login/password`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      ElNotification.success({ title: '提示', message: '重置成功' })
      router.push('/login')
    } else {
      ElNotification.error({ title: '提示', message: res.data.message })
    }
  }).catch(() => {})
}

const init = () => {
  if (isLogin.value) {
    router.push('/')
  } else {
    store.state.HeaderActiveName = '1-4'
  }
}

onMounted(() => {
  init()
  initRules()
})
</script>

<style scoped lang="scss">
.login_form {
  background: #0b1520 url(../../assets/images/login_bg.png) no-repeat center center;
  height: 760px;
  position: relative;
  overflow: hidden;
  .login_right {
    padding: 20px 30px;
    position: absolute;
    background: #17212e;
    width: 350px;
    height: auto;
    left: 50%;
    top: 50%;
    margin-left: -175px;
    margin-top: -205px;
    border-radius: 5px;
    border-top: 4px solid #f0ac19;
    .tel-title {
      font-size: 25px;
      color: #fff;
    }
    .el-form-item {
      .el-form-item-content {
        .login_btn {
          width: 100%;
          outline: none;
          font-size: 18px;
          border-radius: 30px;
        }
        .el-input__inner {
          border: 1px solid red;
        }
        #sendCode {
          position: absolute;
          border: 1px solid #f0ac19;
          background: transparent;
          top: -10px;
          outline: none;
          right: 0;
          width: 30%;
          color: #d5851d;
          cursor: pointer;
        }
      }
    }
  }
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
  }
}
</style>

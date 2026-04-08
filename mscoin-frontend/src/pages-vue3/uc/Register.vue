<template>
  <div class="login_form">
    <div style="width:100%;text-align:center;font-size:24px;color:#fff;position:absolute;top:80px;">
      <marquee width=100% behavior=scroll direction=left align=middle>
        <span v-for="(item, index) in myRandomly" :key="index" style="margin-right:40%;color:#F3BB01;font-size:20px;">
          {{ welcome }}{{ item.telNum }}  {{ become }}{{ item.memberLevelName }}{{ receiveRewards }}
        </span>
      </marquee>
    </div>
    <div class="login_right">
      <el-form v-if="allowRegister" ref="formInline" :model="formInline" :rules="ruleInline" inline>
        <el-form-item style="text-align:center;">
          <div class="tel-title">{{ registerTitle }}</div>
        </el-form-item>
        <el-form-item prop="username" style="display:none;">
          <el-input type="text" v-model="formInline.username" :placeholder="usernameTip" />
        </el-form-item>
        <el-form-item prop="user">
          <el-input type="text" v-model="formInline.user" :placeholder="userTip">
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

        <el-form-item prop="code" v-if="showCode">
          <el-input type="text" v-model="formInline.code" :placeholder="smsCodeTip" />
          <button id="sendCode" @click="sendCode" :disabled="codedisabled" class="send-code-btn">
            {{ sendcodeValue }}
          </button>
        </el-form-item>
        <el-form-item prop="password" class="password">
          <el-input type="password" v-model="formInline.password" :placeholder="pwdTip" />
        </el-form-item>
        <el-form-item prop="repassword" class="password">
          <el-input type="password" v-model="formInline.repassword" :placeholder="rePwdTip" />
        </el-form-item>
        <el-form-item prop="promotion">
          <el-input type="text" v-model="formInline.promotion">
            <template #prepend>{{ promotionCode }}</template>
          </el-input>
        </el-form-item>
        <div id="VAPTCHAContainer" style="width: 300px;height: 36px;"/>
        <div class="check-agree">
          <el-checkbox v-model="agree">{{ userAgreement }}</el-checkbox>
          <a :href="agreementLink" target="_blank" class="agreement-link">《{{ userProtocol }}》</a>
        </div>
        <el-form-item>
          <el-button class="register_btn" @click="handleSubmit" :loading="registing">{{ registerTitle }}</el-button>
        </el-form-item>
      </el-form>
      <el-alert v-else type="warning" :closable="false">
        <template #title>Coming soon!</template>
        <template #description>MSZLU.COM will open register soon</template>
      </el-alert>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 注册页面
 * 迁移点：
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. 使用 Element Plus 组件替代 iView 组件
 * 4. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 */
import { ref, reactive, computed, inject, onMounted, watch } from 'vue'
import { ElMessage, ElNotification } from 'element-plus'
import { vaptcha } from '../../extend/vaptcha'
import axios from 'axios'

// Vuex 3.x 和 Vue Router 3.x 兼容方案
const store = inject('store')
const router = inject('router')

// 文本常量
const welcome = '欢迎'
const become = '成为'
const receiveRewards = '并收到奖励'
const registerTitle = '注册'
const usernameTip = '用户名'
const userTip = '请输入手机号'
const smsCodeTip = '短信验证码'
const pwdTip = '请输入密码'
const rePwdTip = '确认密码'
const promotionCode = '邀请码 :'
const userAgreement = '我已阅读并同意'
const userProtocol = '用户协议'
const sendcodeText = '发送验证码'
const resendcodeText = '重新发送 ('

// 表单数据
const country = ref('中国')
const codedisabled = ref(false)
const sendcodeValue = ref(sendcodeText)
const agree = ref(true)
const allowRegister = ref(true)
const changeActive = ref(0)
const showCode = ref(true)
const countdown = ref(60)
const registing = ref(false)
const captchaObj = ref(null)
const myRandomly = ref([])

const formInline = reactive({
  username: '',
  user: '',
  code: '',
  areaCode: '',
  password: '',
  repassword: '',
  promotion: ''
})

// 表单验证规则
const validateUser = (rule, value, callback) => {
  if (changeActive.value === 0) {
    const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
    if (value === '') {
      callback(new Error('请输入手机号'))
    } else if (!reg.test(formInline.user)) {
      callback(new Error('请输入正确的手机号'))
    } else {
      callback()
    }
  } else {
    const reg = /^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$/
    if (value === '') {
      callback(new Error('请输入邮箱'))
    } else if (!reg.test(formInline.user)) {
      callback(new Error('请输入正确的邮箱'))
    } else {
      callback()
    }
  }
}

const validateRepassword = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请确认密码'))
  } else if (value !== formInline.password) {
    callback(new Error('两次输入的密码不一致'))
  } else {
    callback()
  }
}

const ruleInline = reactive({
  user: [{ validator: validateUser, trigger: 'blur' }],
  code: [
    {
      required: true,
      message: '请输入短信验证码',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    },
    {
      type: 'string',
      min: 6,
      message: '密码长度至少 6 位',
      trigger: 'blur'
    }
  ],
  repassword: [{ validator: validateRepassword, trigger: 'blur' }]
})

// 计算属性
const agreementLink = computed(() => {
  return store.state.lang === '简体中文'
    ? '/helpdetail?cate=1&id=5&cateTitle=常见问题'
    : '/helpdetail?cate=1&id=35&cateTitle=Privacy Policy'
})

// 获取随机推荐用户
const getRandomly = () => {
  axios.post(store.state.host + store.state.api.uc.getRandomly)
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        myRandomly.value = resp.data
      }
    })
}

// 初始化验证码
const initCaptcha = () => {
  vaptcha().then((vaptcha) => {
    vaptcha({
      vid: '63fec1c3507890ee2e7f9dd1',
      mode: 'click',
      scene: 2,
      container: '#VAPTCHAContainer',
      area: 'auto'
    }).then(function (VAPTCHAObj) {
      VAPTCHAObj.render()
      VAPTCHAObj.listen('pass', function () {
        captchaObj.value = {
          server: VAPTCHAObj.server,
          token: VAPTCHAObj.token
        }
      })
    })
  })
}

// 发送验证码
const settime = () => {
  sendcodeValue.value = resendcodeText + countdown.value + ')'
  codedisabled.value = true

  const timer = setInterval(() => {
    countdown.value--
    sendcodeValue.value = resendcodeText + countdown.value + ')'
    if (countdown.value <= 0) {
      clearInterval(timer)
      codedisabled.value = false
      sendcodeValue.value = sendcodeText
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
  }

  const params = {
    phone: formInline.user,
    country: '中国'
  }

  settime()

  axios.post(store.state.host + '/uc/mobile/code', JSON.stringify(params))
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElNotification.success({
          title: '提示',
          message: resp.message
        })
      } else {
        ElNotification.error({
          title: '提示',
          message: resp.message
        })
      }
    })
    .catch(() => {
      ElNotification.error({
        title: '提示',
        message: '发送失败'
      })
    })
}

// 提交注册
const handleSubmit = () => {
  if (captchaObj.value === null) {
    ElMessage.error('请先完成验证')
    return false
  }

  // 模拟表单验证通过
  if (agree.value !== true) {
    ElMessage.error('请同意用户协议')
    return
  }

  if (changeActive.value === 1) {
    // 邮箱注册逻辑（暂未实现）
    return
  }

  if (registing.value) return

  registing.value = true

  const params = {
    phone: formInline.user,
    username: formInline.username + formInline.user,
    password: formInline.password,
    promotion: formInline.promotion,
    code: formInline.code,
    country: country.value,
    captcha: captchaObj.value
  }

  axios.post(store.state.host + '/uc/register/phone', JSON.stringify(params))
    .then(response => {
      registing.value = false
      const resp = response.data
      if (resp.code === 0) {
        ElNotification.success({
          title: '提示',
          message: '注册成功!'
        })
        setTimeout(() => {
          router.push('/login')
        }, 3000)
      } else {
        ElNotification.error({
          title: '提示',
          message: resp.message
        })
      }
    })
    .catch(() => {
      registing.value = false
      ElNotification.error({
        title: '提示',
        message: '注册失败'
      })
    })
}

// 监听语言变化
watch(() => store.state.lang, () => {
  // 语言变化时更新文本
})

// 生命周期
onMounted(() => {
  window.scrollTo(0, 0)
  store.commit('navigate', 'nav-other')
  store.state.HeaderActiveName = '0'
  getRandomly()
  initCaptcha()

  // 处理 URL 参数中的邀请码
  const query = router.currentRoute.value.query
  if (query.code) {
    formInline.promotion = query.code
  }
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
    height: 585px;
    left: 50%;
    top: 50%;
    margin-left: -175px;
    margin-top: -255px;
    border-top: 4px solid #f0ac19;
    border-radius: 5px;

    .tel-title {
      color: #fff;
      font-size: 25px;
    }

    .check-agree {
      color: #979797;
      display: inline-block;
      line-height: 30px;
      font-size: 12px;
      cursor: default;
      margin-bottom: 10px;

      .agreement-link {
        color: #f0ac19;
        margin-left: -10px;
      }
    }

    :deep(.el-form-item) {
      .el-form-item__content {
        .register_btn.el-button {
          width: 100%;
          background-color: #f0ac19;
          border-color: #f0ac19;
          color: #fff;
          border-radius: 5px;
          font-size: 18px;
          margin-top: 20px;

          &:hover, &:focus {
            background-color: #e5a316;
            border-color: #e5a316;
          }
        }

        .send-code-btn {
          position: absolute;
          border: 1px solid #f0ac19;
          background: transparent;
          top: -10px;
          outline: none;
          right: 0;
          width: 30%;
          color: #f0ac19;
          cursor: pointer;
          height: 38px;

          &:disabled {
            cursor: not-allowed;
            opacity: 0.6;
          }
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
// 全局样式
.el-form.el-form--inline {
  .el-form-item {
    .el-form-item__content {
      .el-input-wrapper {
        .el-input {
          .el-input__inner {
            border: none;
            border-bottom: 1px solid #27313e;
            font-size: 14px;
            background: transparent;
            border-radius: 0;

            &:focus {
              border: none;
              border-bottom: 1px solid #27313e;
            }
          }
        }
      }
    }
  }

  .el-checkbox {
    .el-checkbox__input {
      .el-checkbox__inner {
        background: transparent;
        border: 1px solid #979797;

        &:hover {
          border-color: #f0ac19;
        }
      }

      &.is-checked {
        .el-checkbox__inner {
          background-color: #f0ac19;
          border-color: #f0ac19;
        }
      }
    }
  }
}

.el-select {
  .el-input__inner {
    background-color: transparent;
    color: #fff;
  }

  .el-input-group__prepend {
    background-color: #17212e;
    border-bottom: 1px solid #27313e;
    color: #fff;
  }
}

.el-dropdown-menu__item {
  background-color: #1c2a32;
  color: #fff;
}
</style>

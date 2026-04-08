<template>
  <div class="login_form">
    <div class="login_right">
      <el-form ref="formInline" :model="formInline" :rules="ruleInline" inline>
        <div class="login_title">{{ loginTitle }}</div>
        <el-form-item prop="user">
          <el-input
            name="user"
            type="text"
            v-model="formInline.user"
            :placeholder="userTip"
            class="user"
          >
            <template #prepend>
              <el-select v-model="country" style="width: 65px">
                <el-option value="+86" label="+86"><span>+86</span><span style="margin-left:10px;color:#ccc">中国</span></el-option>
                <el-option value="+65" label="+65"><span>+65</span><span style="margin-left:10px;color:#ccc">新加坡</span></el-option>
                <el-option value="+82" label="+82"><span>+82</span><span style="margin-left:10px;color:#ccc">韩国</span></el-option>
                <el-option value="+81" label="+81"><span>+81</span><span style="margin-left:10px;color:#ccc">日本</span></el-option>
                <el-option value="+66" label="+66"><span>+66</span><span style="margin-left:10px;color:#ccc">泰国</span></el-option>
                <el-option value="+7" label="+7"><span>+7</span><span style="margin-left:10px;color:#ccc">俄罗斯</span></el-option>
                <el-option value="+44" label="+44"><span>+44</span><span style="margin-left:10px;color:#ccc">英国</span></el-option>
                <el-option value="+84" label="+84"><span>+84</span><span style="margin-left:10px;color:#ccc">越南</span></el-option>
                <el-option value="+91" label="+91"><span>+91</span><span style="margin-left:10px;color:#ccc">印度</span></el-option>
                <el-option value="+39" label="+39"><span>+39</span><span style="margin-left:10px;color:#ccc">意大利</span></el-option>
                <el-option value="+852" label="+852"><span>+852</span><span style="margin-left:10px;color:#ccc">香港</span></el-option>
                <el-option value="+60" label="+60"><span>+60</span><span style="margin-left:10px;color:#ccc">马来西亚</span></el-option>
                <el-option value="+886" label="+886"><span>+886</span><span style="margin-left:10px;color:#ccc">台湾省</span></el-option>
                <el-option value="+90" label="+90"><span>+90</span><span style="margin-left:10px;color:#ccc">土耳其</span></el-option>
              </el-select>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password" class="password">
          <el-input
            type="password"
            v-model="formInline.password"
            :placeholder="pwdTip"
            @keyup.enter="doLogin"
          />
        </el-form-item>
        <div id="VAPTCHAContainer" style="width: 300px;height: 36px;"/>
        <p id="notice" class="hide">{{ validateMsg }}</p>
        <p style="height:30px;margin-top:15px;">
          <router-link to="/findPwd" style="color:#979797;float:right;padding-right:10px;font-size:12px;">
            {{ forgetPwd }}
          </router-link>
        </p>
        <el-form-item style="margin-bottom:15px;">
          <el-button class="login_btn" @click="doLogin" :loading="logining">{{ loginTitle }}</el-button>
        </el-form-item>
        <div class='to_register'>
          <span>{{ noAccount }}</span>
          <router-link to="/register">{{ goRegister }}</router-link>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 登录页面
 * 迁移点：
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. 使用 Element Plus 组件替代 iView 组件
 * 4. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 */
import { ref, reactive, computed, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { vaptcha } from '../../extend/vaptcha'

// Vuex 3.x 和 Vue Router 3.x 兼容方案
const store = inject('store')
const router = inject('router')

// 文本常量（替代 i18n）
const loginTitle = '登录'
const userTip = '请输入手机号/邮箱'
const pwdTip = '请输入密码'
const validateMsg = '请先完成验证码验证'
const forgetPwd = '忘记密码？'
const noAccount = '没有账号？'
const goRegister = '立即注册'

// 表单数据
const country = ref('+86')
const logining = ref(false)
const captchaObj = ref(null)

const formInline = reactive({
  user: '',
  password: ''
})

// 表单验证规则
const validateUser = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入手机号或邮箱'))
  } else {
    callback()
  }
}

const validatePassword = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else if (value.length < 6) {
    callback(new Error('密码长度至少 6 位'))
  } else {
    callback()
  }
}

const ruleInline = reactive({
  user: [{ validator: validateUser, trigger: 'blur' }],
  password: [
    { validator: validatePassword, trigger: 'blur' }
  ]
})

// 表单引用
const formRef = ref(null)

// 初始化验证码
const initCaptcha = () => {
  vaptcha().then((vaptcha) => {
    vaptcha({
      vid: '63fec1c3507890ee2e7f9dd1',
      mode: 'click',
      scene: 1,
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

// 登录处理
const doLogin = () => {
  const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
  const tel = formInline.user
  const flagtel = reg.test(tel)
  const flagpass = formInline.password.length >= 6

  if (!flagtel || !flagpass) {
    ElMessage.error('请填写完整的信息')
    return false
  }

  if (captchaObj.value === null) {
    ElMessage.error('请先完成验证')
    return false
  }

  handleSubmit()
}

const handleSubmit = () => {
  const params = {
    username: formInline.user,
    password: formInline.password,
    captcha: captchaObj.value
  }

  // 使用 axios 替代 this.$http
  import('axios').then(({ default: axios }) => {
    const state = store.state
    axios.post(state.host + state.api.uc.login, JSON.stringify(params))
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('登录成功')
          store.commit('setMember', resp.data)
          localStorage.setItem('TOKEN', resp.data.token)
          router.push('/uc/safe')
        } else {
          ElMessage.error(resp.message)
        }
      })
      .catch(error => {
        ElMessage.error('登录失败：' + error.message)
      })
  })
}

// 生命周期
onMounted(() => {
  store.commit('navigate', 'nav-other')
  store.state.HeaderActiveName = '0'

  // 检查是否已登录
  if (store.getters.isLogin) {
    router.push('/uc/safe')
  } else {
    initCaptcha()
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
    padding: 20px 30px 20px 30px;
    position: absolute;
    background: #17212e;
    width: 350px;
    height: 366px;
    left: 50%;
    top: 50%;
    margin-left: -175px;
    margin-top: -165px;
    border-top: 4px solid #f0ac19;
    border-radius: 5px;

    .login_title {
      height: 70px;
      color: #fff;
      text-align: center;
      font-size: 25px;
    }

    :deep(.el-form-item) {
      .el-form-item__content {
        .login_btn.el-button {
          width: 100%;
          background-color: #f0ac19;
          border-color: #f0ac19;
          color: #fff;
          font-size: 18px;
          border-radius: 5px;

          &:hover, &:focus {
            background-color: #e5a316;
            border-color: #e5a316;
          }
        }
      }
    }
  }

  .to_register {
    overflow: hidden;
    font-size: 12px;

    span {
      float: left;
    }

    a {
      float: right;
      color: #f0ac19;
    }
  }
}

.hide {
  display: none;
}

#notice {
  color: red;
}

:deep(.el-input) {
  .el-input__inner {
    background-color: transparent;
    font-size: 14px;
    border: none;
    border-bottom: 1px solid #27313e;
    border-radius: 0;

    &:focus {
      border-bottom: 1px solid #27313e;
    }
  }

  .el-input-group__prepend {
    background-color: #17212e;
    border-bottom: 1px solid #27313e;
    color: #fff;
  }
}

:deep(.el-select) {
  .el-input__inner {
    background-color: transparent;
    color: #fff;
  }
}

:deep(.el-dropdown-menu__item) {
  background-color: #1c2a32;
  color: #fff;
}
</style>

<style lang="scss">
// 全局样式覆盖
.el-form.el-form--inline {
  .el-form-item {
    .el-form-item__content {
      .el-input-wrapper {
        .el-input {
          .el-input__inner {
            background-color: transparent;
            color: #fff;
          }
        }
      }
    }
  }
}

.el-select-dropdown {
  background: #1c2a32;

  .el-select-dropdown__item {
    color: #fff;

    &.selected {
      color: #f0ac19;
    }
  }
}
</style>

<template>
  <div class="register-page">
    <div class="register-page__ticker" v-if="rewardMessages.length">
      <div class="register-page__ticker-track">
        <span v-for="item in rewardMessages" :key="item.id" class="register-page__ticker-item">
          {{ item.message }}
        </span>
      </div>
    </div>

    <div class="register-card">
      <div class="register-card__header">
        <h1>注册</h1>
        <p>创建账户后即可继续体验平台服务</p>
      </div>

      <el-form ref="formRef" :model="form" :rules="rules" label-position="top" autocomplete="off">
        <input class="register-autofill-trap" type="text" name="username" autocomplete="username" tabindex="-1">
        <input class="register-autofill-trap" type="password" name="current-password" autocomplete="current-password" tabindex="-1">
        <el-form-item class="register-form-item" prop="user">
          <el-input
            ref="userInputRef"
            v-model="form.user"
            name="register-phone"
            autocomplete="off"
            placeholder="手机号">
            <template #prepend>
              <el-select v-model="countryCode" style="width: 92px">
                <el-option
                  v-for="option in countryOptions"
                  :key="option.value"
                  :label="option.label"
                  :value="option.value"
                />
              </el-select>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item class="register-form-item" prop="code">
          <div class="register-card__code-row">
            <el-input
              ref="codeInputRef"
              v-model="form.code"
              name="register-sms-code"
              autocomplete="one-time-code"
              placeholder="短信验证码" />
            <el-button :disabled="codedisabled" @click="sendCode">
              {{ sendcodeValue }}
            </el-button>
          </div>
        </el-form-item>

        <el-form-item class="register-form-item" prop="password">
          <el-input
            ref="passwordInputRef"
            v-model="form.password"
            type="password"
            name="register-password"
            autocomplete="new-password"
            show-password
            placeholder="密码" />
        </el-form-item>

        <el-form-item class="register-form-item" prop="repassword">
          <el-input
            ref="repasswordInputRef"
            v-model="form.repassword"
            type="password"
            name="register-password-confirm"
            autocomplete="new-password"
            show-password
            placeholder="确认密码" />
        </el-form-item>

        <el-form-item class="register-form-item" prop="promotion">
          <el-input v-model="form.promotion" placeholder="邀请码（选填）">
            <template #prepend>邀请码</template>
          </el-input>
        </el-form-item>

        <div :class="['captcha-box', { 'captcha-passed': captchaPassed }]">
          <div class="captcha-box__title">人机验证</div>
          <p class="captcha-box__desc">当前使用前端降级验证，后续可切回正式验证码服务。</p>
          <el-button
            class="captcha-box__button"
            :type="captchaPassed ? 'success' : 'default'"
            @click="markCaptchaPassed"
          >
            {{ captchaPassed ? '验证已完成' : '点击完成验证' }}
          </el-button>
        </div>

        <div class="register-card__agreement">
          <el-checkbox v-model="agree">我已阅读并同意</el-checkbox>
          <a :href="agreementLink" target="_blank">《用户协议》</a>
        </div>

        <el-form-item class="register-form-item">
          <el-button class="register-card__submit" type="warning" :loading="registing" @click="handleSubmit">
            注册
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { computed, inject, nextTick, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElNotification } from 'element-plus'
import axios from 'axios'
import { buildApiUrl, useRuntimeContract } from '../../config/runtime-vue3'

const store = inject('store')
const router = inject('router')
const runtime = useRuntimeContract()
const formRef = ref(null)
const userInputRef = ref(null)
const codeInputRef = ref(null)
const passwordInputRef = ref(null)
const repasswordInputRef = ref(null)

const countryCode = ref('+86')
const codedisabled = ref(false)
const sendcodeValue = ref('发送验证码')
const agree = ref(true)
const registing = ref(false)
const captchaPassed = ref(false)
const rewardMessages = ref([])
const countdown = ref(60)

const countryOptions = [
  { value: '+86', label: '+86 中国' },
  { value: '+65', label: '+65 新加坡' },
  { value: '+81', label: '+81 日本' },
  { value: '+82', label: '+82 韩国' },
  { value: '+852', label: '+852 中国香港' },
  { value: '+886', label: '+886 中国台湾' }
]

const form = reactive({
  username: '',
  user: '',
  code: '',
  password: '',
  repassword: '',
  promotion: ''
})

const rules = {
  user: [
    {
      validator: (_rule, value, callback) => {
        if (!value) {
          callback(new Error('请输入手机号'))
          return
        }
        const reg = /^[1][3-9][0-9]{9}$/
        if (!reg.test(value)) {
          callback(new Error('手机号格式不正确'))
          return
        }
        callback()
      },
      trigger: 'blur'
    }
  ],
  code: [{ required: true, message: '请输入短信验证码', trigger: 'blur' }],
  password: [
    {
      validator: (_rule, value, callback) => {
        if (!value) {
          callback(new Error('请输入密码'))
          return
        }
        if (value.length < 6) {
          callback(new Error('密码长度至少 6 位'))
          return
        }
        callback()
      },
      trigger: 'blur'
    }
  ],
  repassword: [
    {
      validator: (_rule, value, callback) => {
        if (!value) {
          callback(new Error('请再次输入密码'))
          return
        }
        if (value !== form.password) {
          callback(new Error('两次输入的密码不一致'))
          return
        }
        callback()
      },
      trigger: 'blur'
    }
  ]
}

const agreementLink = computed(() => {
  return store.state.lang === 'English'
    ? '/helpdetail?cate=1&id=35&cateTitle=Privacy Policy'
    : '/helpdetail?cate=1&id=5&cateTitle=常见问题'
})

const markCaptchaPassed = () => {
  captchaPassed.value = true
}

const clearRegisterAutofill = async () => {
  form.user = ''
  form.code = ''
  form.password = ''
  form.repassword = ''

  await nextTick()

  ;[userInputRef, codeInputRef, passwordInputRef, repasswordInputRef].forEach((inputRef) => {
    if (inputRef.value?.input) {
      inputRef.value.input.value = ''
    }
  })
}

const settime = () => {
  sendcodeValue.value = `重新发送(${countdown.value})`
  codedisabled.value = true

  const timer = setInterval(() => {
    countdown.value -= 1
    sendcodeValue.value = `重新发送(${countdown.value})`
    if (countdown.value <= 0) {
      clearInterval(timer)
      codedisabled.value = false
      sendcodeValue.value = '发送验证码'
      countdown.value = 60
    }
  }, 1000)
}

const getRandomly = async () => {
  try {
    const response = await axios.post(buildApiUrl(runtime.api.uc.getRandomly))
    const resp = response.data
    if (resp.code === 0 && Array.isArray(resp.data)) {
      rewardMessages.value = resp.data.slice(0, 6).map((item, index) => ({
        id: `${item.telNum ?? 'user'}-${index}`,
        message: `欢迎 ${item.telNum ?? '新用户'} 成为 ${item.memberLevelName ?? '会员'} 并领取奖励`
      }))
    }
  } catch {
    rewardMessages.value = []
  }
}

const sendCode = async () => {
  const valid = await formRef.value?.validateField('user').catch(() => false)
  if (valid === false) return

  if (!captchaPassed.value) {
    ElMessage.error('请先完成人机验证')
    return
  }

  const params = {
    phone: form.user,
    country: '中国'
  }

  try {
    const response = await axios.post(buildApiUrl('/uc/mobile/code'), params)
    const resp = response.data
    if (resp.code === 0) {
      settime()
      ElNotification.success({ title: '提示', message: resp.message || '验证码已发送' })
      return
    }
    ElNotification.error({ title: '提示', message: resp.message || '发送失败' })
  } catch (error) {
    ElNotification.error({ title: '提示', message: error?.response?.data?.message || '发送验证码失败' })
  }
}

const handleSubmit = async () => {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  if (!captchaPassed.value) {
    ElMessage.error('请先完成人机验证')
    return
  }

  if (!agree.value) {
    ElMessage.error('请先阅读并同意用户协议')
    return
  }

  if (registing.value) return
  registing.value = true

  const params = {
    phone: form.user,
    username: form.username || form.user,
    password: form.password,
    promotion: form.promotion,
    code: form.code,
    country: countryCode.value,
    captcha: {
      server: 'local-fallback',
      token: 'passed'
    }
  }

  try {
    const response = await axios.post(buildApiUrl('/uc/register/phone'), params)
    const resp = response.data
    if (resp.code === 0) {
      ElNotification.success({ title: '提示', message: resp.message || '注册成功' })
      setTimeout(() => router.push('/login'), 1200)
      return
    }
    ElNotification.error({ title: '提示', message: resp.message || '注册失败' })
  } catch (error) {
    ElNotification.error({ title: '提示', message: error?.response?.data?.message || '注册请求失败' })
  } finally {
    registing.value = false
  }
}

onMounted(() => {
  window.scrollTo(0, 0)
  store.commit('navigate', 'nav-other')
  store.state.HeaderActiveName = '0'
  const query = router.currentRoute.value.query
  if (query.code) {
    form.promotion = query.code
  }
  getRandomly()
  void clearRegisterAutofill()
  window.setTimeout(() => {
    void clearRegisterAutofill()
  }, 300)
})
</script>

<style scoped lang="scss">
.register-page {
  min-height: 760px;
  background: #0b1520 url(../../assets/images/login_bg.png) no-repeat center center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 16px 60px;
}

.register-page__ticker {
  max-width: 980px;
  margin: 0 auto 24px;
  overflow: hidden;
  border: 1px solid rgba(240, 167, 10, 0.24);
  border-radius: 999px;
  background: rgba(243, 187, 1, 0.08);
}

.register-page__ticker-track {
  display: flex;
  gap: 36px;
  padding: 12px 18px;
  white-space: nowrap;
  overflow: auto hidden;
  scrollbar-width: none;
}

.register-page__ticker-track::-webkit-scrollbar {
  display: none;
}

.register-page__ticker-item {
  color: #f3bb01;
  font-size: 14px;
}

.register-card {
  width: 100%;
  max-width: 380px;
  margin: 0 auto;
  padding: 28px 30px;
  background: #17212e;
  border-top: 4px solid #f0ac19;
  border-radius: 6px;
  box-shadow: 0 18px 60px rgba(0, 0, 0, 0.32);
}

.register-card__header {
  margin-bottom: 18px;
  text-align: center;

  h1 {
    margin: 0 0 8px;
    color: #fff;
    font-size: 28px;
    font-weight: 600;
  }

  p {
    margin: 0;
    color: #8d99ae;
    font-size: 13px;
  }
}

.register-card__code-row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 128px;
  gap: 10px;
  width: 100%;
  align-items: stretch;

  .el-button {
    height: 42px;
    background-color: #27313e;
    border-color: #27313e;
    color: #fff;
    border-radius: 6px;
    transition: all 0.2s ease;

    &:hover {
      background-color: #2f3a47;
      color: #f0a70a;
    }

    &:disabled {
      background-color: #1f2833;
      color: #6b7280;
      cursor: not-allowed;
    }
  }
}

.captcha-box {
  margin: 10px 0 18px;
  padding: 14px 16px;
  border: 1px solid #27313e;
  border-radius: 6px;
  background: rgba(11, 21, 32, 0.45);
  transition: border-color 0.2s ease, background-color 0.2s ease;

  &.captcha-passed {
    border-color: rgba(102, 187, 106, 0.3);
    background: rgba(102, 187, 106, 0.08);
  }
}

.captcha-box__title {
  color: #fff;
  font-size: 14px;
  font-weight: 600;
}

.captcha-box__desc {
  margin: 6px 0 12px;
  color: #8d99ae;
  font-size: 12px;
  line-height: 1.5;
}

.captcha-box__button {
  width: 100%;
  height: 40px;
  font-size: 14px;
  font-weight: 500;

  /* 未完成态 - 强调可点击（覆盖 Element Plus default 类型） */
  &.el-button--default,
  &:not(.el-button--success) {
    background-color: #f0ac19 !important;
    border-color: #f0ac19 !important;
    color: #0b1520 !important;

    &:hover {
      background-color: #ffb83d !important;
      border-color: #ffb83d !important;
      color: #0b1520 !important;
    }
  }

  /* 已完成态 - 明确绿色成功状态 */
  &.el-button--success {
    background-color: #66bb6a !important;
    border-color: #66bb6a !important;
    color: #fff !important;
    cursor: default;

    &:hover {
      background-color: #81c784 !important;
      border-color: #81c784 !important;
      color: #fff !important;
    }
  }
}

.register-card__agreement {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
  color: #a6b2c6;
  font-size: 13px;

  :deep(.el-checkbox) {
    height: auto;
    margin: 0;
    cursor: pointer;
  }

  :deep(.el-checkbox__input) {
    cursor: pointer;
  }

  a {
    color: #f0ac19;
    cursor: pointer;

    &:hover {
      color: #f0a70a;
      text-decoration: underline;
    }
  }
}

.register-autofill-trap {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  clip-path: inset(50%);
  border: 0;
  pointer-events: none;
}

.register-card__submit {
  width: 100%;
  height: 42px;
}

.register-form-item {
  margin-bottom: 20px;
}

/* 输入框和选择器暗色样式 */
.register-form-item :deep(.el-input__wrapper),
.register-form-item :deep(.el-select__wrapper) {
  background: #101923 !important;
  box-shadow: inset 0 0 0 1px #314153 !important;
  border-radius: 6px;
  transition: box-shadow 0.2s ease, background-color 0.2s ease;
}

/* 组合控件 prepend 样式（区号选择框、邀请码前缀） */
.register-form-item :deep(.el-input-group__prepend) {
  background: #101923;
  box-shadow: inset 0 0 0 1px #314153;
  border: 0;
  border-radius: 6px 0 0 6px;
  transition: box-shadow 0.2s ease, background-color 0.2s ease;
  color: #a6b2c6;
  font-size: 13px;
}

/* hover 态 */
.register-form-item :deep(.el-input__wrapper:hover),
.register-form-item :deep(.el-select__wrapper:hover) {
  box-shadow: inset 0 0 0 1px #4a627b !important;
}

.register-form-item :deep(.el-input-group__prepend:hover) {
  box-shadow: inset 0 0 0 1px #4a627b;
}

/* focus 态 */
.register-form-item :deep(.el-input__wrapper.is-focus),
.register-form-item :deep(.el-select__wrapper.is-focused),
.register-form-item :deep(.el-select__wrapper.is-focus) {
  background: #132131 !important;
  box-shadow: inset 0 0 0 1px #f0ac19 !important;
}

/* 当输入框聚焦时，同步改变 prepend 背景色 - 使用 class 联动而非:has() */
.register-form-item :deep(.el-input-group--focus .el-input-group__prepend),
.register-form-item :deep(.el-input-group.is-focus .el-input-group__prepend),
.register-form-item :deep(.el-input-group:focus-within .el-input-group__prepend) {
  background: #132131 !important;
  box-shadow: inset 0 0 0 1px #f0ac19 !important;
}

.register-form-item :deep(.el-input-group__prepend:focus-within) {
  background: #132131 !important;
  box-shadow: inset 0 0 0 1px #f0ac19 !important;
}

/* error 态 */
.register-form-item.is-error :deep(.el-input__wrapper),
.register-form-item.is-error :deep(.el-select__wrapper) {
  box-shadow: inset 0 0 0 1px #f56c6c !important;
}

.register-form-item.is-error :deep(.el-input-group__prepend) {
  box-shadow: inset 0 0 0 1px #f56c6c !important;
}

.register-form-item :deep(.el-input-group__prepend .el-select__wrapper) {
  box-shadow: none;
  background: transparent;
  border-radius: 0;
}

/* 全局 prepend 文字颜色 */
:deep(.el-input-group__prepend) {
  color: #fff;
}

/* 输入框文字和占位符颜色 */
:deep(.el-input__inner) {
  color: #fff;
}

:deep(.el-input__inner::placeholder) {
  color: #70839a;
}

/* 密码输入框特殊处理 */
:deep(input[type="password"]) {
  color: #fff;
  -webkit-text-fill-color: #fff;
}

/* 修复浏览器自动填充白底问题 */
:deep(input:-webkit-autofill),
:deep(input:-webkit-autofill:hover),
:deep(input:-webkit-autofill:focus),
:deep(input:-webkit-autofill:active) {
  -webkit-box-shadow: 0 0 0 40px #101923 inset !important;
  -webkit-text-fill-color: #fff !important;
  background-color: #101923 !important;
  color: #fff !important;
  transition: background-color 5000s ease-in-out 0s;
}

:deep(.el-form-item__label) {
  color: #a6b2c6;
}

:deep(.el-button:not(.register-card__submit)) {
  background-color: #27313e;
  border-color: #27313e;
  color: #fff;

  &:hover {
    color: #f0a70a;
  }

  &:disabled {
    background-color: #1f2833;
    color: #6b7280;
  }
}
</style>

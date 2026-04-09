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

      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item prop="user">
          <el-input v-model="form.user" placeholder="手机号">
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

        <el-form-item prop="code">
          <div class="register-card__code-row">
            <el-input v-model="form.code" placeholder="短信验证码" />
            <el-button :disabled="codedisabled" @click="sendCode">
              {{ sendcodeValue }}
            </el-button>
          </div>
        </el-form-item>

        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" show-password placeholder="密码" />
        </el-form-item>

        <el-form-item prop="repassword">
          <el-input v-model="form.repassword" type="password" show-password placeholder="确认密码" />
        </el-form-item>

        <el-form-item prop="promotion">
          <el-input v-model="form.promotion" placeholder="邀请码（选填）">
            <template #prepend>邀请码</template>
          </el-input>
        </el-form-item>

        <div class="captcha-box">
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

        <el-form-item>
          <el-button class="register-card__submit" type="warning" :loading="registing" @click="handleSubmit">
            注册
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { computed, inject, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElNotification } from 'element-plus'
import axios from 'axios'
import { buildApiUrl, useRuntimeContract } from '../../config/runtime-vue3'

const store = inject('store')
const router = inject('router')
const runtime = useRuntimeContract()
const formRef = ref(null)

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
    const response = await axios.post(buildApiUrl('/uc/mobile/code'), JSON.stringify(params))
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
      mode: 'fallback',
      passed: true
    }
  }

  try {
    const response = await axios.post(buildApiUrl('/uc/register/phone'), JSON.stringify(params))
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
})
</script>

<style scoped lang="scss">
.register-page {
  min-height: 760px;
  background: #0b1520 url(../../assets/images/login_bg.png) no-repeat center center;
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
  grid-template-columns: minmax(0, 1fr) 112px;
  gap: 10px;
  width: 100%;
}

.captcha-box {
  margin: 10px 0 18px;
  padding: 14px 16px;
  border: 1px solid #27313e;
  border-radius: 6px;
  background: rgba(11, 21, 32, 0.45);
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
}

.register-card__agreement {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
  color: #a6b2c6;
  font-size: 12px;

  a {
    color: #f0ac19;
  }
}

.register-card__submit {
  width: 100%;
  height: 42px;
}

:deep(.el-input__wrapper),
:deep(.el-select__wrapper) {
  background: transparent;
  box-shadow: none;
}

:deep(.el-input-group__prepend) {
  background: #17212e;
  border: 0;
  border-bottom: 1px solid #27313e;
  color: #fff;
  box-shadow: none;
}

:deep(.el-input__inner) {
  color: #fff;
}

:deep(.el-form-item__label) {
  color: #a6b2c6;
}
</style>

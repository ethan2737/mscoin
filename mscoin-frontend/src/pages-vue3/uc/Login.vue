<template>
  <div class="auth-page">
    <div class="auth-card">
      <div class="auth-card__header">
        <h1>登录</h1>
        <p>使用手机号或邮箱进入账户</p>
      </div>

      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item class="auth-form-item" prop="user">
          <el-input v-model="form.user" placeholder="手机号或邮箱">
            <template #prepend>
              <el-select v-model="countryCode" style="width: 84px">
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

        <el-form-item class="auth-form-item" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            show-password
            placeholder="密码"
            @keyup.enter="doLogin"
          />
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

        <div class="auth-card__links">
          <router-link to="/findPwd">忘记密码？</router-link>
        </div>

        <el-form-item>
          <el-button class="auth-card__submit" type="warning" :loading="logining" @click="doLogin">
            登录
          </el-button>
        </el-form-item>
      </el-form>

      <div class="auth-card__footer">
        <span>没有账号？</span>
        <router-link to="/register">立即注册</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
import { buildApiUrl, useRuntimeContract } from '../../config/runtime-vue3'
import { establishAuthenticatedSession, hasAuthenticatedSession } from '../../utils/auth-session'

const store = inject('store')
const router = inject('router')
const runtime = useRuntimeContract()
const getReturnUrl = () => router.currentRoute.value.query.returnUrl || router.currentRoute.value.params.returnUrl || '/uc/safe'

const formRef = ref(null)
const countryCode = ref('+86')
const captchaPassed = ref(false)
const logining = ref(false)

const countryOptions = [
  { value: '+86', label: '+86 中国' },
  { value: '+65', label: '+65 新加坡' },
  { value: '+81', label: '+81 日本' },
  { value: '+82', label: '+82 韩国' },
  { value: '+852', label: '+852 中国香港' },
  { value: '+886', label: '+886 中国台湾' }
]

const form = reactive({
  user: '',
  password: ''
})

const rules = {
  user: [
    {
      validator: (_rule, value, callback) => {
        if (!value) {
          callback(new Error('请输入手机号或邮箱'))
          return
        }
        callback()
      },
      trigger: 'blur'
    }
  ],
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
  ]
}

const markCaptchaPassed = () => {
  captchaPassed.value = true
}

const doLogin = async () => {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  if (!captchaPassed.value) {
    ElMessage.error('请先完成人机验证')
    return
  }

  logining.value = true
  const params = {
    username: form.user,
    password: form.password,
    captcha: {
      mode: 'fallback',
      passed: true
    }
  }

  try {
    const response = await axios.post(buildApiUrl(runtime.api.uc.login), params)
    const resp = response.data
    if (resp.code === 0) {
      await establishAuthenticatedSession({
        loginData: resp.data,
        axiosInstance: axios,
        memberInfoUrl: buildApiUrl(runtime.api.uc.memberInfo),
        storage: localStorage,
        tokenKey: runtime.tokenKey,
        memberKey: runtime.memberKey,
        store
      })
      ElMessage.success('登录成功')
      router.push(getReturnUrl())
      return
    }
    ElMessage.error(resp.message || '登录失败')
  } catch (error) {
    ElMessage.error(error?.response?.data?.message || '登录请求失败')
  } finally {
    logining.value = false
  }
}

onMounted(() => {
  store.commit('navigate', 'nav-other')
  store.state.HeaderActiveName = '0'
  store.commit('recoveryMember')
  if (hasAuthenticatedSession({
    storage: localStorage,
    tokenKey: runtime.tokenKey,
    memberKey: runtime.memberKey,
    store
  })) {
    router.push(getReturnUrl())
  }
})
</script>

<style scoped lang="scss">
.auth-page {
  min-height: 760px;
  background: #0b1520 url(../../assets/images/login_bg.png) no-repeat center center;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px 16px;
}

.auth-card {
  width: 100%;
  max-width: 380px;
  padding: 28px 30px;
  background: #17212e;
  border-top: 4px solid #f0ac19;
  border-radius: 6px;
  box-shadow: 0 18px 60px rgba(0, 0, 0, 0.32);
}

.auth-card__header {
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

.auth-card__links {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 12px;

  a {
    color: #a6b2c6;
    font-size: 12px;
  }
}

.auth-card__submit {
  width: 100%;
  height: 42px;
}

.auth-card__footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #8d99ae;
  font-size: 12px;

  a {
    color: #f0ac19;
  }
}

.auth-form-item {
  margin-bottom: 20px;
}

.auth-form-item :deep(.el-input__wrapper),
.auth-form-item :deep(.el-select__wrapper) {
  background: #101923;
  box-shadow: inset 0 0 0 1px #314153;
  border-radius: 6px;
  transition: box-shadow 0.2s ease, background-color 0.2s ease;
}

.auth-form-item :deep(.el-input-group__prepend) {
  background: #101923;
  box-shadow: inset 0 0 0 1px #314153;
  border: 0;
  border-radius: 6px 0 0 6px;
  transition: box-shadow 0.2s ease, background-color 0.2s ease;
}

.auth-form-item :deep(.el-input__wrapper:hover),
.auth-form-item :deep(.el-select__wrapper:hover) {
  box-shadow: inset 0 0 0 1px #4a627b;
}

.auth-form-item :deep(.el-input-group__prepend:hover) {
  box-shadow: inset 0 0 0 1px #4a627b;
}

.auth-form-item :deep(.el-input__wrapper.is-focus),
.auth-form-item :deep(.el-select__wrapper.is-focused),
.auth-form-item :deep(.el-select__wrapper.is-focus) {
  background: #132131;
  box-shadow: inset 0 0 0 1px #f0ac19;
}

.auth-form-item :deep(.el-input-group__prepend:focus-within) {
  background: #132131;
  box-shadow: inset 0 0 0 1px #f0ac19;
}

.auth-form-item.is-error :deep(.el-input__wrapper),
.auth-form-item.is-error :deep(.el-select__wrapper) {
  box-shadow: inset 0 0 0 1px #f56c6c;
}

.auth-form-item.is-error :deep(.el-input-group__prepend) {
  box-shadow: inset 0 0 0 1px #f56c6c;
}

:deep(.el-input-group__prepend) {
  color: #fff;
}

.auth-form-item :deep(.el-input-group__prepend .el-select__wrapper) {
  box-shadow: none;
  background: transparent;
  border-radius: 0;
}

:deep(.el-input__inner) {
  color: #fff;
}

:deep(.el-input__inner::placeholder) {
  color: #70839a;
}

:deep(.el-form-item__label) {
  color: #a6b2c6;
}
</style>

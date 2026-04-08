<template>
  <div class="nav-rights safe">
    <div class="nav-right padding-right-clear">
      <div class="padding-right-clear padding-left-clear rightarea user account-box">
        <div class="rightarea-con">
          <div class="user-top-icon">
            <div class="user-icons">
              <div class="user-face user-avatar-public">
                <span class="user-avatar-in">V{{ user.level.id }}</span>
              </div>
              <div class="user-name" style="line-height:52px">
                <span style="line-height:52px">{{ user.username }}</span>
              </div>
            </div>
            <el-row class="user-right">
              <el-col :xs="24" :sm="24" :md="8" :lg="8">
                <i class="m3"></i>
                <div class="itp" v-if="user.realVerified==0&&user.phoneVerified==0&&user.fundsVerified==0">{{ safeLevelLow }}</div>
                <div class="itp" v-else-if="user.realVerified==1&&user.phoneVerified==1&&user.fundsVerified==1">{{ safeLevelHigh }}</div>
                <div class="itp" v-else>{{ safeLevelMedium }}</div>
              </el-col>
            </el-row>
          </div>
          <section class="accountContent">
            <div class="account-in">
              <!-- 实名认证 -->
              <div class="account-item">
                <div class="account-item-in">
                  <el-icon size="18" color="#00b5f6"><credit-card /></el-icon>
                  <span class="card-number">{{ verifiedTitle }}</span>
                  <p v-if="user.realVerified==1" class="bankInfo" style="color: #fff;font-size: 13px;">{{ user.realName }}</p>
                  <p v-else-if="user.realVerified==0&&user.realAuditing==0&&user.realNameRejectReason!=null" class="bankInfo" style="color: #f0a70a;font-size: 13px;">
                    审核未通过{{ user.realNameRejectReason?"："+user.realNameRejectReason:""}}，请重试。
                  </p>
                  <p v-else class="bankInfo" style="color: #828ea1;font-size: 13px;">
                    {{ verifiedTip }}
                  </p>
                  <span v-if="user.realVerified==1">{{ verifyPass }}</span>
                  <span v-else-if="user.realAuditing==1">{{ binding }}</span>
                  <a class="btn" @click="showItem(6)" v-else-if="user.realVerified==0&&user.realAuditing==0&&user.realNameRejectReason!=null" :title="user.realNameRejectReason">{{ bindRetry }}</a>
                  <a v-else class="btn" @click="showItem(6)">{{ notVerified }}</a>
                </div>
                <div class="account-detail" v-show="choseItem==6">
                  <div class="detail-list" style="width: 100%;">
                    <el-form ref="formValidate6" :model="formValidate6" :rules="ruleValidate" label-width="85px" style="text-align:center;">
                      <el-form-item :label="realNameLabel" prop="realName">
                        <el-input v-model="formValidate6.realName" size="large"></el-input>
                      </el-form-item>
                      <el-form-item :label="idCardLabel" prop="idCard">
                        <el-input v-model="formValidate6.idCard" size="large"></el-input>
                      </el-form-item>
                      <div class="sfz_bottom">
                        <el-row>
                          <el-col :xs="24" :lg="8">
                            <input type="hidden" name="imgPreview" :value="imgPreview" />
                            <div class="idcard-title">{{ uploadPositive }}</div>
                            <img id="frontCardImg" style="width: 180px;height: 120px;" :src="frontCardImg">
                            <div class="acc_sc">
                              <el-upload ref="upload1" :before-upload="beforeUpload" :on-success="frontHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :show-file-list="false">
                                <el-button icon="Upload">{{ uploadBtn }}</el-button>
                              </el-upload>
                            </div>
                          </el-col>
                          <el-col :xs="24" :lg="8">
                            <input type="hidden" name="imgNext" :value="imgNext" />
                            <div class="idcard-title">{{ uploadNegative }}</div>
                            <img id="backCardImg" style="width: 180px;height: 120px;" :src="backCardImg">
                            <div class="acc_sc">
                              <el-upload ref="upload2" :before-upload="beforeUpload" :on-success="backHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :show-file-list="false">
                                <el-button icon="Upload">{{ uploadBtn }}</el-button>
                              </el-upload>
                            </div>
                          </el-col>
                          <el-col :xs="24" :lg="8">
                            <input type="hidden" name="imgLast" :value="imgLast" />
                            <div class="idcard-title">{{ uploadHand }}</div>
                            <img id="handCardImg" style="width: 180px;height: 120px;" :src="handCardImg">
                            <div class="acc_sc">
                              <el-upload ref="upload3" :before-upload="beforeUpload" :on-success="handHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :show-file-list="false">
                                <el-button icon="Upload">{{ uploadBtn }}</el-button>
                              </el-upload>
                            </div>
                          </el-col>
                        </el-row>
                        <div class="idcard-desc">
                          <p>{{ idcardVerifymsg1 }}</p>
                          <p>{{ idcardVerifymsg2 }}</p>
                          <p>{{ idcardVerifymsg3 }}</p>
                        </div>
                      </div>
                      <el-form-item style="text-align:center;">
                        <el-button type="warning" @click="handleSubmit('formValidate6')">{{ saveBtn }}</el-button>
                        <el-button @click="handleReset('formValidate6')" style="margin-left: 8px">{{ resetBtn }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                </div>
              </div>

              <!-- 手机号绑定 -->
              <div class="account-item">
                <div class="account-item-in">
                  <el-icon size="20" color="#00b5f6"><cellphone /></el-icon>
                  <span class="card-number">{{ phoneTitle }}</span>
                  <p v-if="user.phoneVerified==1" class="bankInfo" style="color: #fff;font-size: 13px;">
                    {{ user.mobilePhone }}
                  </p>
                  <p v-else class="bankInfo" style="color: #828ea1;font-size: 13px;">
                    {{ bindPhone }}
                  </p>
                  <span v-if="user.phoneVerified==1">{{ binded }}</span>
                  <a v-else class="btn" @click="showItem(3)">{{ bind }}</a>
                </div>
                <div class="account-detail" v-show="choseItem==3">
                  <div class="detail-list">
                    <el-form ref="formValidate3" :model="formValidate3" :rules="ruleValidate" label-width="110px">
                      <el-form-item :label="phoneLabel" prop="mobile">
                        <el-input v-model="formValidate3.mobile" size="large"></el-input>
                      </el-form-item>
                      <el-form-item :label="loginPwdLabel" prop="password">
                        <el-input v-model="formValidate3.password" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="phoneCodeLabel" prop="vailCode2">
                        <el-input v-model="formValidate3.vailCode2" size="large">
                          <template #append>
                            <div class="timebox">
                              <el-button @click="send(2)" :disabled="sendMsgDisabled2">
                                <span v-if="sendMsgDisabled2">{{ time2 }}{{ second }}</span>
                                <span v-if="!sendMsgDisabled2">{{ clickGet }}</span>
                              </el-button>
                            </div>
                          </template>
                        </el-input>
                      </el-form-item>
                      <el-form-item>
                        <el-button type="warning" @click="handleSubmit('formValidate3')">{{ saveBtn }}</el-button>
                        <el-button @click="handleReset('formValidate3')" style="margin-left: 8px">{{ resetBtn }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                </div>
              </div>

              <!-- 登录密码 -->
              <div class="account-item">
                <div class="account-item-in">
                  <el-icon size="20" color="#00b5f6"><lock /></el-icon>
                  <span class="card-number">{{ loginPwdTitle }}</span>
                  <p class="bankInfo" style="color: #828ea1;font-size: 13px;">
                    {{ loginTip }}
                  </p>
                  <a class="btn" v-if="user.phoneVerified==0" @click="noPhone">{{ edit }}</a>
                  <a class="btn" v-else @click="showItem(4)">{{ edit }}</a>
                </div>
                <div class="account-detail" v-show="choseItem==4">
                  <div class="detail-list">
                    <el-form ref="formValidate4" :model="formValidate4" :rules="ruleValidate" label-width="95px">
                      <el-form-item :label="oldPwdLabel" prop="oldPw">
                        <el-input v-model="formValidate4.oldPw" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="newPwdLabel" prop="newPw">
                        <el-input v-model="formValidate4.newPw" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="confirmNewPwdLabel" prop="newPwConfirm">
                        <el-input v-model="formValidate4.newPwConfirm" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="phoneCodeLabel" prop="vailCode3">
                        <el-input v-model="formValidate4.vailCode3" size="large">
                          <template #append>
                            <div class="timebox">
                              <el-button @click="send(3)" :disabled="sendMsgDisabled3">
                                <span v-if="sendMsgDisabled3">{{ time3 }}{{ second }}</span>
                                <span v-if="!sendMsgDisabled3">{{ clickGet }}</span>
                              </el-button>
                            </div>
                          </template>
                        </el-input>
                      </el-form-item>
                      <el-form-item>
                        <el-button type="warning" @click="handleSubmit('formValidate4')">{{ saveBtn }}</el-button>
                        <el-button @click="handleReset('formValidate4')" style="margin-left: 8px">{{ resetBtn }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                </div>
              </div>

              <!-- 资金密码 -->
              <div class="account-item">
                <div class="account-item-in">
                  <el-icon size="20" color="#00b5f6"><money /></el-icon>
                  <span class="card-number">{{ fundPwdTitle }}</span>
                  <p class="bankInfo" style="color: #828ea1;font-size: 13px;">
                    {{ fundTip }}
                  </p>
                  <a class="btn" v-if="user.phoneVerified==0" @click="noPhone">{{ set }}</a>
                  <a class="btn" v-else-if="user.fundsVerified==0" @click="showItem(5)">{{ set }}</a>
                  <a class="btn" v-else @click="showItemFundpwd()">{{ edit }}</a>
                </div>
                <div class="account-detail" v-show="choseItem==5">
                  <!-- 设置资金密码 -->
                  <div class="detail-list" v-show="user.fundsVerified!=1">
                    <el-form ref="formValidate7" :model="formValidate7" :rules="ruleValidate" label-width="85px">
                      <el-form-item :label="fundPwdLabel" prop="pw7">
                        <el-input v-model="formValidate7.pw7" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="confirmPwdLabel" prop="pw7Confirm">
                        <el-input v-model="formValidate7.pw7Confirm" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item>
                        <el-button type="warning" @click="handleSubmit('formValidate7')">{{ saveBtn }}</el-button>
                        <el-button @click="handleReset('formValidate7')" style="margin-left: 8px">{{ resetBtn }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                  <!-- 修改资金密码 -->
                  <div class="detail-list" v-show="user.fundsVerified==1 && !fGetBackFundpwd">
                    <el-form ref="formValidate5" :model="formValidate5" :rules="ruleValidate" label-width="95px">
                      <el-form-item :label="oldFundPwdLabel" prop="oldPw">
                        <el-input v-model="formValidate5.oldPw" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="newFundPwdLabel" prop="newMPw">
                        <el-input v-model="formValidate5.newMPw" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="confirmNewPwdLabel" prop="newMPwConfirm">
                        <el-input v-model="formValidate5.newMPwConfirm" size="large" type="password"></el-input>
                      </el-form-item>
                      <p style="text-align:right;">
                        <a @click="handleReset('formValidate8');fGetBackFundpwd=!fGetBackFundpwd" style="color:#f0ac19;">{{ forgetPwd }}</a>
                      </p>
                      <el-form-item>
                        <el-button type="warning" @click="handleSubmit('formValidate5')">{{ saveBtn }}</el-button>
                        <el-button @click="handleReset('formValidate5')" style="margin-left: 8px">{{ resetBtn }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                  <!-- 找回资金密码 -->
                  <div class="detail-list" v-show="user.fundsVerified==1 && fGetBackFundpwd">
                    <el-form ref="formValidate8" :model="formValidate8" :rules="ruleValidate" label-width="85px">
                      <el-form-item :label="newFundPwdLabel" prop="newMPw8">
                        <el-input v-model="formValidate8.newMPw8" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="confirmNewPwdLabel" prop="newMPwConfirm8">
                        <el-input v-model="formValidate8.newMPwConfirm8" size="large" type="password"></el-input>
                      </el-form-item>
                      <el-form-item :label="phoneCodeLabel" prop="vailCode5">
                        <el-input v-model="formValidate8.vailCode5" size="large">
                          <template #append>
                            <div class="timebox">
                              <el-button @click="send(5)" :disabled="sendMsgDisabled5">
                                <span v-if="sendMsgDisabled5">{{ time5 }}{{ second }}</span>
                                <span v-if="!sendMsgDisabled5">{{ clickGet }}</span>
                              </el-button>
                            </div>
                          </template>
                        </el-input>
                      </el-form-item>
                      <el-form-item>
                        <el-button type="warning" @click="handleSubmit('formValidate8')">{{ saveBtn }}</el-button>
                        <el-button @click="handleReset('formValidate8')" style="margin-left: 8px">{{ resetBtn }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 安全设置页面
 * 迁移点：
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. 使用 Element Plus 组件替代 iView 组件
 * 4. 使用 inject('store') 兼容 Vuex 3.x
 */
import { ref, reactive, computed, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { CreditCard, Cellphone, Lock, Money, Upload } from '@element-plus/icons-vue'
import axios from 'axios'

// Vuex 3.x 兼容方案
const store = inject('store')

// 文本常量
const safeLevelLow = '安全等级低'
const safeLevelHigh = '安全等级高'
const safeLevelMedium = '安全等级中'
const verifiedTitle = '实名认证'
const verifiedTip = '完成实名认证，提高账户安全性'
const verifyPass = '已认证'
const binding = '审核中'
const bindRetry = '重新认证'
const notVerified = '未认证'
const realNameLabel = '真实姓名'
const idCardLabel = '身份证号'
const uploadPositive = '上传身份证正面'
const uploadNegative = '上传身份证反面'
const uploadHand = '上传手持身份证'
const uploadBtn = '上传'
const idcardVerifymsg1 = '请确保身份证在有效期内'
const idcardVerifymsg2 = '请确保身份证信息清晰可见'
const idcardVerifymsg3 = '请确保照片与身份证本人一致'
const saveBtn = '保存'
const resetBtn = '重置'
const phoneTitle = '手机号'
const bindPhone = '绑定手机号，提高账户安全性'
const binded = '已绑定'
const bind = '绑定'
const edit = '修改'
const set = '设置'
const loginPwdTitle = '登录密码'
const loginTip = '定期修改登录密码，提高账户安全性'
const fundPwdTitle = '资金密码'
const fundTip = '用于提现、转账等资金操作'
const phoneLabel = '手机号'
const loginPwdLabel = '登录密码'
const phoneCodeLabel = '手机验证码'
const oldPwdLabel = '旧密码'
const newPwdLabel = '新密码'
const confirmNewPwdLabel = '确认新密码'
const fundPwdLabel = '资金密码'
const confirmPwdLabel = '确认密码'
const oldFundPwdLabel = '旧资金密码'
const newFundPwdLabel = '新资金密码'
const forgetPwd = '忘记密码？'
const clickGet = '点击获取'
const second = '秒'

// 表单数据
const fGetBackFundpwd = ref(false)
const imgPreview = ref('')
const imgNext = ref('')
const imgLast = ref('')
const frontCardImg = ref(new URL('../../assets/images/frontCardImg.png', import.meta.url).href)
const backCardImg = ref(new URL('../../assets/images/backCardImg.png', import.meta.url).href)
const handCardImg = ref(new URL('../../assets/images/HandCardImg.png', import.meta.url).href)
const choseItem = ref(0)
const time1 = ref(60)
const time2 = ref(60)
const time3 = ref(60)
const time5 = ref(60)
const sendMsgDisabled1 = ref(false)
const sendMsgDisabled2 = ref(false)
const sendMsgDisabled3 = ref(false)
const sendMsgDisabled5 = ref(false)

const uploadHeaders = reactive({ 'x-auth-token': localStorage.getItem('TOKEN') })
const uploadUrl = store.state.host + '/uc/upload/oss/image'

const user = ref({
  level: { id: 0, name: '' },
  realVerified: 0,
  realAuditing: 0,
  phoneVerified: 0,
  fundsVerified: 0,
  mobilePhone: '',
  realName: '',
  realNameRejectReason: null
})

const formValidate2 = reactive({ mail: '', vailCode1: '', password: '' })
const formValidate3 = reactive({ mobile: '', vailCode2: '', password: '' })
const formValidate4 = reactive({ oldPw: '', newPw: '', newPwConfirm: '', vailCode3: '' })
const formValidate5 = reactive({ oldPw: '', newMPw: '', newMPwConfirm: '' })
const formValidate6 = reactive({ realName: '', idCard: '' })
const formValidate7 = reactive({ pw7: '', pw7Confirm: '' })
const formValidate8 = reactive({ newMPw8: '', newMPwConfirm8: '', vailCode5: '' })

// 表单验证规则
const ruleValidate = reactive({
  mail: [
    { required: true, type: 'email', message: '请输入正确的邮箱', trigger: 'blur' }
  ],
  vailCode1: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  mobile: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
  vailCode2: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  vailCode3: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  vailCode5: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  password: [
    { required: true, type: 'string', min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  oldPw: [
    { required: true, type: 'string', min: 6, message: '请输入旧密码', trigger: 'blur' }
  ],
  newPw: [
    { required: true, type: 'string', min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  newPwConfirm: [
    { required: true, type: 'string', min: 6, message: '请确认新密码', trigger: 'blur' }
  ],
  newMPw: [
    { required: true, type: 'string', min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  newMPwConfirm: [
    { required: true, type: 'string', min: 6, message: '请确认密码', trigger: 'blur' }
  ],
  pw7: [
    { required: true, type: 'string', min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  pw7Confirm: [
    { required: true, type: 'string', min: 6, message: '请确认密码', trigger: 'blur' }
  ],
  newMPw8: [
    { required: true, type: 'string', min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  newMPwConfirm8: [
    { required: true, type: 'string', min: 6, message: '请确认密码', trigger: 'blur' }
  ],
  realName: [
    { required: true, message: '请输入真实姓名', trigger: 'blur' }
  ],
  idCard: [
    { required: true, message: '请输入身份证号', trigger: 'blur' }
  ]
})

// 上传前验证
const beforeUpload = (file) => {
  if (file && file.size >= 1024000 * 8) {
    ElMessage.error('上传图片大小不能超过 8M')
    return false
  }
  return true
}

// 图片上传成功处理
const frontHandleSuccess = (res, file, fileList) => {
  if (res.code === 0) {
    frontCardImg.value = imgPreview.value = res.data
  } else {
    ElMessage.error(res.message)
  }
}

const backHandleSuccess = (res, file, fileList) => {
  if (res.code === 0) {
    backCardImg.value = imgNext.value = res.data
  } else {
    ElMessage.error(res.message)
  }
}

const handHandleSuccess = (res, file, fileList) => {
  if (res.code === 0) {
    handCardImg.value = imgLast.value = res.data
  } else {
    ElMessage.error(res.message)
  }
}

// 显示表单项
const showItem = (index) => {
  choseItem.value = index
}

const showItemFundpwd = () => {
  fGetBackFundpwd.value = false
  handleReset('formValidate5')
  showItem(5)
}

// 没有手机号提示
const noPhone = () => {
  ElMessage.info('请先绑定手机号')
  showItem(3)
}

// 发送验证码
const send = (index) => {
  if (index === 1) {
    if (formValidate2.mail) {
      axios.post(store.state.host + '/uc/bind/email/code', { email: formValidate2.mail })
        .then(response => {
          const resp = response.data
          if (resp.code === 0) {
            sendMsgDisabled1.value = true
            const interval = setInterval(() => {
              if (time1.value-- <= 0) {
                time1.value = 60
                sendMsgDisabled1.value = false
                clearInterval(interval)
              }
            }, 1000)
          } else {
            ElMessage.error(resp.message)
          }
        })
    }
  } else if (index === 2) {
    if (formValidate3.mobile) {
      axios.post(store.state.host + '/uc/mobile/bind/code', { phone: formValidate3.mobile })
        .then(response => {
          const resp = response.data
          if (resp.code === 0) {
            sendMsgDisabled2.value = true
            const interval = setInterval(() => {
              if (time2.value-- <= 0) {
                time2.value = 60
                sendMsgDisabled2.value = false
                clearInterval(interval)
              }
            }, 1000)
          } else {
            ElMessage.error(resp.message)
          }
        })
    }
  } else if (index === 3) {
    axios.post(store.state.host + '/uc/mobile/update/password/code')
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          sendMsgDisabled3.value = true
          const interval = setInterval(() => {
            if (time3.value-- <= 0) {
              time3.value = 60
              sendMsgDisabled3.value = false
              clearInterval(interval)
            }
          }, 1000)
        } else {
          ElMessage.error(resp.message)
        }
      })
  } else if (index === 5) {
    axios.post(store.state.host + '/uc/mobile/transaction/code')
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          sendMsgDisabled5.value = true
          const interval = setInterval(() => {
            if (time5.value-- <= 0) {
              time5.value = 60
              sendMsgDisabled5.value = false
              clearInterval(interval)
            }
          }, 1000)
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
}

// 表单重置
const handleReset = (name) => {
  const form = formRefs.value[name]
  if (form) form.resetFields()
}

// 表单引用
const formRefs = ref({})

// 表单提交
const submit = (name) => {
  // 实名认证
  if (name === 'formValidate6') {
    if (imgPreview.value === '') {
      ElMessage.error('请上传身份证正面照')
      return false
    }
    if (imgNext.value === '') {
      ElMessage.error('请上传身份证反面照')
      return false
    }
    if (imgLast.value === '') {
      ElMessage.error('请上传手持身份证照')
      return false
    }
    const param = {
      realName: formValidate6.realName,
      idCard: formValidate6.idCard,
      idCardFront: imgPreview.value,
      idCardBack: imgNext.value,
      handHeldIdCard: imgLast.value
    }
    axios.post(store.state.host + '/uc/approve/real/name', param)
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('保存成功')
          getMember()
          choseItem.value = 0
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
  // 手机认证
  if (name === 'formValidate3') {
    const param = {
      phone: formValidate3.mobile,
      code: formValidate3.vailCode2,
      password: formValidate3.password
    }
    axios.post(store.state.host + '/uc/approve/bind/phone', param)
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('保存成功')
          getMember()
          choseItem.value = 0
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
  // 登录密码
  if (name === 'formValidate4') {
    const param = {
      oldPassword: formValidate4.oldPw,
      newPassword: formValidate4.newPw,
      code: formValidate4.vailCode3
    }
    axios.post(store.state.host + '/uc/approve/update/password', param)
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('保存成功')
          getMember()
          choseItem.value = 0
          store.commit('removeMember')
          localStorage.removeItem('TOKEN')
          setTimeout(() => {
            const router = inject('router')
            router.push('/login')
          }, 2000)
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
  // 修改资金密码
  if (name === 'formValidate5') {
    const param = {
      oldPassword: formValidate5.oldPw,
      newPassword: formValidate5.newMPw
    }
    axios.post(store.state.host + '/uc/approve/update/transaction/password', param)
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('保存成功')
          handleReset('formValidate5')
          getMember()
          choseItem.value = 0
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
  // 设置资金密码
  if (name === 'formValidate7') {
    const param = {
      jyPassword: formValidate7.pw7
    }
    axios.post(store.state.host + '/uc/approve/transaction/password', param)
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('保存成功')
          getMember()
          choseItem.value = 0
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
  // 找回资金密码
  if (name === 'formValidate8') {
    const param = {
      newPassword: formValidate8.newMPw8,
      code: formValidate8.vailCode5
    }
    axios.post(store.state.host + '/uc/approve/reset/transaction/password', param)
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success('保存成功')
          fGetBackFundpwd.value = false
          handleReset('formValidate5')
          getMember()
          choseItem.value = 0
        } else {
          ElMessage.error(resp.message)
        }
      })
  }
}

const handleSubmit = (name) => {
  const form = formRefs.value[name]
  if (!form) return
  form.validate(valid => {
    if (valid) {
      submit(name)
    } else {
      ElMessage.error('保存失败')
    }
  })
}

// 获取会员信息
const getMember = () => {
  axios.post(store.state.host + '/uc/approve/security/setting')
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        user.value = resp.data
      } else {
        ElMessage.error('请先登录')
      }
    })
}

// 生命周期
onMounted(() => {
  getMember()
})
</script>

<style scoped lang="scss">
button.el-button {
  &:focus {
    box-shadow: 0 0 0 2px rgba(45, 140, 240, 0);
  }
}

button.el-button.el-button--primary {
  box-shadow: 0 0 0 2px rgba(45, 140, 240, 0);
}

.nav-right {
  padding-left: 15px;

  .user .user-top-icon {
    height: 80px;
    border-bottom: 1px dashed #27313e;
    position: relative;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding: 0 50px;
  }
}

.uploadimgtip {
  position: relative;
  top: -20px;
  color: #f0a70a;
}

.account-box .account-in .account-item .account-detail {
  padding: 30px 0;
  margin: 6px 0;
}

.account-box .account-in .account-item .account-detail .detail-list {
  width: 40%;
  margin: 0 auto;
}

.sfz_bottom {
  height: 250px;
}

.btn_marginLeft {
  margin-left: -85px;
}

@media screen and (max-width: 1200px) {
  .account-box .account-in .account-item .account-detail .detail-list {
    width: 90%;
    margin: 0 auto;
  }

  .btn_marginLeft {
    margin-left: 0px;
  }

  .sfz_bottom {
    height: auto;
  }

  :deep(.el-form) {
    .el-form-item {
      .el-form-item__content {
        margin-left: 0px !important;
      }
    }
  }
}

.account-box .account-in .account-item .account-detail .detail-list .input-control {
  margin-bottom: 10px;
  height: 45px;
}

.detail-list .input-control .el-input-group__prepend {
  width: 63px;
}

.detail-list .input-control .el-input {
  height: 45px;
}

.account-box .account-in .account-item {
  margin-bottom: 10px;
}

.account-box .account-in .account-item .account-item-in {
  display: flex;
  align-items: center;
  padding: 15px 30px 15px 50px;
  box-shadow: 0 1px 0 0 rgba(69, 112, 128, 0.06);
  font-size: 14px;
  color: #fff;
}

.account-box .account-in .account-item .account-item-in .icons {
  height: 17px;
  width: 17px;
  display: inline-block;
  margin-top: -1px;
  background-size: 100% 100%;
}

.account-box .account-in .account-item .account-item-in .yesImg {
  background-image: url(../../assets/img/overicon.png);
}

.icons.noImg {
  background-image: url(../../assets/img/noicon.png);
}

.account-box .account-in .account-item .account-item-in .card-number {
  width: 142px;
  height: 40px;
  margin-right: 15px;
  border-right: 1px dashed #27313e;
  padding: 0 15px;
  line-height: 40px;
  text-align: left;
  display: inline-block;
}

.account-box .account-in .account-item .account-item-in .bankInfo {
  width: 70%;
  text-align: left;
}

.account-box .account-in .account-item .account-item-in .btn {
  padding: 8px 10px;
  cursor: pointer;
  color: #f0a70a;
}

.tips-g {
  color: #8994a3;
  font-size: 12px;
}

.gr {
  color: #b4b4b4;
}

.m1 {
  display: inline-block;
  width: 25px;
  height: 25px;
  background: url(../../assets/img/m1.png);
  background-size: 100% 100%;
  vertical-align: middle;
  margin-right: 5px;
}

.m2 {
  display: inline-block;
  width: 25px;
  height: 25px;
  background: url(../../assets/img/m2.png);
  background-size: 100% 100%;
  vertical-align: middle;
  margin-right: 5px;
}

.m3 {
  display: inline-block;
  width: 25px;
  height: 25px;
  background: url(../../assets/img/m3.png);
  background-size: 100% 100%;
  vertical-align: middle;
  margin-right: 5px;
}

.itp {
  display: inline-block;
}

.user-right {
  width: 80%;
}

.rightarea {
  padding-left: 15px !important;
  padding-right: 15px !important;
  margin-bottom: 60px !important;
}

.rightarea .rightarea-top {
  line-height: 75px;
  border-bottom: #f1f1f1 solid 1px;
}

.rightarea .trade-process {
  line-height: 30px;
  padding: 0 15px;
  background: #f1f1f1;
  display: inline-block;
  position: relative;
  margin-right: 20px;
}

.rightarea .trade-process.active {
  color: #eb6f6c;
  background: #f9f5eb;
}

.rightarea .trade-process .icon {
  background: #fff;
  border-radius: 20px;
  height: 20px;
  width: 20px;
  display: inline-block;
  line-height: 20px;
  text-align: center;
  margin-right: 10px;
}

.rightarea .trade-process .arrow {
  position: absolute;
  top: 10px;
  right: -5px;
  width: 0;
  height: 0;
  border-top: 5px solid transparent;
  border-bottom: 5px solid transparent;
  border-left: 5px solid #f1f1f1;
}

.rightarea .trade-process.active .arrow {
  border-left: 5px solid #f9f5eb;
}

.rightarea .rightarea-tabs {
  border: none;
}

.rightarea .rightarea-tabs li > a {
  width: 100%;
  height: 100%;
  padding: 0;
  margin-right: 0;
  font-size: 14px;
  color: #646464;
  border-radius: 0;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
}

.rightarea .rightarea-tabs li > a:hover {
  background-color: #fcfbfb;
}

.rightarea .rightarea-tabs li {
  width: 125px;
  height: 40px;
  position: relative;
  margin: -1px 0 0 -1px;
  border: 1px solid #f1f1f1;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.rightarea .rightarea-tabs li.active {
  background-color: #fcfbfb;
}

.rightarea .rightarea-tabs li:last-child {
  border-right: 1px solid #f1f1f1;
}

.rightarea .rightarea-tabs li.active > a,
.rightarea .rightarea-tabs li:hover > a {
  color: #da2e22;
  border: none;
}

.rightarea .panel-tips {
  border: 3px solid #fdfaf3;
  color: #9e9e9e;
  font-size: 12px;
}

.rightarea .panel-tips .panel-header {
  background: #fdfaf3;
  line-height: 40px;
  margin-bottom: 15px;
}

.rightarea .panel-tips .panel-title {
  font-size: 16px;
}

.rightarea .recordtitle {
  cursor: pointer;
}

.user .top-icon {
  width: 75px;
  height: 75px;
  display: inline-block;
}

.user .user-info {
  padding: 0px 10px;
  background-color: #fff;
  color: #fff;
}

.user .user-info .user-info-top {
  border-bottom: 1px dashed #27313e;
  padding-bottom: 20px;
}

.user .user-info h5 {
  font-size: 18px;
}

.user .user-info h5 .iconfont {
  font-size: 20px;
  color: #e24a64;
  margin-left: 10px;
}

.user-avatar {
  display: flex;
  justify-content: space-between;
}

.user-icons {
  display: flex;
  align-self: center;
  width: 300px;
}

.user-icons .icons-in {
  height: 45px;
  width: 45px;
  border-radius: 50%;
  background-color: #00b5f6;
  display: flex;
  justify-content: center;
  align-items: center;
  align-self: center;
}

.user-icons .icons-in em {
  color: white;
  font-size: 20px;
  font-style: normal;
}

.user-icons .user-name {
  margin-left: 10px;
  display: flex;
  justify-content: flex-start;
  flex-direction: column;
}

.user-icons .user-name span {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  width: 225px;
  height: 52px;
  overflow: hidden;
  font-size: 14px;
  white-space: nowrap;
  text-overflow: ellipsis;
  text-overflow: ellipsis;
}

.user-top-icon .trade-info {
  width: 420px;
  padding-left: 30px;
  display: flex;
}

.user-top-icon .trade-info .item {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.user-top-icon .trade-info .item.capital {
  width: 60%;
}

.user-icons .user-name span.uid {
  color: #8994a3;
  font-size: 12px;
}

.circle-info {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  height: 80px;
  width: 80px;
  border-radius: 50%;
  border: 2px solid #ebeff5;
  margin-left: 14px;
}

.circle-info span {
  font-weight: bolder;
}

.circle-info p {
  color: #8994a3;
  margin: 0;
  padding: 0;
}

.circle-info p.count {
  color: #fff;
  font-size: 14px;
  font-weight: 600;
}

.user-avatar-public {
  background: #df9a00;
  border-radius: 50%;
  height: 52px;
  width: 52px;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 1px 5px 0 rgba(71, 78, 114, 0.24);
  position: relative;
}

.user-avatar-public > .user-avatar-in {
  background: #f0a70a;
  border-radius: 50%;
  height: 42px;
  width: 42px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
}

.router-link-active {
  color: red;
}

.account-item-in .el-icon {
  color: #f0a70a !important;
}

.btn {
  color: #f0a70a;
}

.el-button.el-button--primary {
  background-color: #f0a70a;
  border-color: #f0a70a;
}
</style>

<style lang="scss">
li.el-upload-list__file.is-finish {
  &:hover {
    span {
      color: #f0a70a;
    }
  }
}

.idcard-title {
  font-size: 13px;
  margin-bottom: 10px;
}

.acc_sc {
  margin-top: 10px;
}

.idcard-desc {
  padding: 10px 150px 50px 150px;

  > p {
    font-size: 13px;
    margin-bottom: 10px;
    text-align: left;
    color: #828ea1;
  }
}

@media screen and (max-width: 1200px) {
  .safe .nav-right .user .user-top-icon {
    padding: 0 0 !important;
  }

  .idcard-desc {
    padding: 0;

    > p {
      font-size: 13px;
      margin-bottom: 10px;
      text-align: left;
      color: #828ea1;
    }
  }

  .safe .account-box .account-in .account-item .account-item-in {
    padding: 15px 0px 15px 0px;
  }

  .safe .account-box .account-in .account-item .account-item-in .bankInfo {
    width: 50% !important;
  }

  .safe .account-box .account-in .account-item .account-item-in .card-number {
    width: 100px !important;
  }

  .safe .user-icons .user-name span {
    width: 100px !important;
  }
}
</style>

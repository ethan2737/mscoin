<template>
  <div class="nav-rights uc_account">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear record account-box">
        <section class="trade-group merchant-top">
          <i class="merchant-icon tips"></i>
          <span class="tips-word">{{ pageTitle }}</span>
          <span class="tips-g">{{ pageTip }}</span>
        </section>
        <section class="accountContent">
          <div class="account-in">
            <!-- 银行卡 -->
            <div class="account-item">
              <div class="account-item-in">
                <span class="card-number">{{ bankCardTitle }}</span>
                <p v-if="user.bankVerified === 1" class="bankInfo" style="color: #fff;">
                  {{ user.bankInfo?.cardNo }}
                </p>
                <p v-else class="bankInfo">{{ bankCardTip }}</p>
                <a class="btn" v-if="user.bankVerified === 1" @click="showItem(1)">{{ modifyBtn }}</a>
                <a class="btn" v-else @click="showItem(1)">{{ bindBtn }}</a>
              </div>
              <div class="account-detail" v-show="choseItem === 1">
                <div class="detail-list">
                  <el-form ref="formValidate1" :model="formValidate1" :rules="ruleValidate" label-width="85px">
                    <el-form-item :label="nameLabel" prop="name">
                      <el-input disabled v-model="formValidate1.name" />
                    </el-form-item>
                    <el-form-item :label="bankAccountLabel" prop="bankName">
                      <el-select v-model="formValidate1.bankName" style="width: 100%;">
                        <el-option v-for="item in bankNameList" :key="item.value" :label="item.label" :value="item.value" />
                      </el-select>
                    </el-form-item>
                    <el-form-item :label="bankBranchLabel" prop="bankBranch">
                      <el-input v-model="formValidate1.bankBranch" />
                    </el-form-item>
                    <el-form-item :label="bankNoLabel" prop="bankNo">
                      <el-input v-model="formValidate1.bankNo" type="text" />
                    </el-form-item>
                    <el-form-item :label="confirmBankNoLabel" prop="bankNoConfirm">
                      <el-input v-model="formValidate1.bankNoConfirm" type="text" />
                    </el-form-item>
                    <el-form-item :label="fundPwdLabel" prop="password">
                      <el-input v-model="formValidate1.password" type="password" />
                    </el-form-item>
                    <el-form-item>
                      <el-button type="warning" @click="handleSubmit('formValidate1')">{{ saveBtn }}</el-button>
                    </el-form-item>
                  </el-form>
                </div>
              </div>
            </div>

            <!-- 支付宝 -->
            <div class="account-item">
              <div class="account-item-in">
                <span class="card-number">{{ alipayTitle }}</span>
                <p v-if="user.aliVerified === 1" class="bankInfo" style="color: #fff;">
                  {{ user.alipay?.aliNo }}
                </p>
                <p v-else class="bankInfo">{{ alipayTip }}</p>
                <a class="btn" v-if="user.aliVerified === 1" @click="showItem(2)">{{ modifyBtn }}</a>
                <a class="btn" v-else @click="showItem(2)">{{ bindBtn }}</a>
              </div>
              <div class="account-detail" v-show="choseItem === 2">
                <div class="detail-list">
                  <el-form ref="formValidate2" :model="formValidate2" :rules="ruleValidate" label-width="95px">
                    <el-row>
                      <el-col :xs="24" :lg="8">
                        <input type="hidden" name="aliPreview" :value="aliPreview" />
                        <img v-if="aliImg" :alt="imgTip" style="width: 200px;height: 200px;" :src="aliImg" />
                        <img v-else :alt="imgTip" style="width: 200px;height: 200px;border:1px solid #ccc;border-radius: 10px;" :src="placeholderImg" />
                        <div class="acc_sc">
                          <el-upload ref="upload1" :before-upload="beforeUpload" :on-success="aliHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :show-file-list="false">
                            <el-button icon="Upload">{{ uploadBtn }}IMG</el-button>
                          </el-upload>
                        </div>
                      </el-col>
                      <el-col :xs="24" :lg="16">
                        <el-form-item :label="nameLabel" prop="name">
                          <el-input disabled v-model="formValidate2.name" />
                        </el-form-item>
                        <el-form-item :label="alipayAccountLabel" prop="alipay">
                          <el-input v-model="formValidate2.alipay" />
                        </el-form-item>
                        <el-form-item :label="fundPwdLabel" prop="password">
                          <el-input v-model="formValidate2.password" type="password" />
                        </el-form-item>
                        <el-form-item>
                          <el-button type="warning" @click="handleSubmit('formValidate2')">{{ saveBtn }}</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
                  </el-form>
                </div>
              </div>
            </div>

            <!-- 微信支付 -->
            <div class="account-item">
              <div class="account-item-in">
                <span class="card-number">{{ wechatTitle }}</span>
                <p v-if="user.wechatVerified === 1" class="bankInfo" style="color: #fff;">
                  {{ user.wechatPay?.wechat }}
                </p>
                <p v-else class="bankInfo">{{ wechatTip }}</p>
                <a class="btn" v-if="user.wechatVerified === 1" @click="showItem(3)">{{ modifyBtn }}</a>
                <a class="btn" v-else @click="showItem(3)">{{ bindBtn }}</a>
              </div>
              <div class="account-detail" v-show="choseItem === 3">
                <div class="detail-list">
                  <el-form ref="formValidate3" :model="formValidate3" :rules="ruleValidate" label-width="85px">
                    <el-row>
                      <el-col :xs="24" :lg="8">
                        <input type="hidden" name="wePreview" :value="wePreview" />
                        <img v-if="weImg" :alt="imgTip" style="width: 200px;height: 200px;" :src="weImg" />
                        <img v-else :alt="imgTip" style="width: 200px;height: 200px;border:1px solid #ccc;border-radius: 10px;" :src="placeholderImg" />
                        <div class="acc_sc">
                          <el-upload ref="upload2" :before-upload="beforeUpload" :on-success="weHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :show-file-list="false">
                            <el-button icon="Upload">{{ uploadBtn }}</el-button>
                          </el-upload>
                        </div>
                      </el-col>
                      <el-col :xs="24" :lg="16">
                        <el-form-item :label="nameLabel" prop="name">
                          <el-input disabled v-model="formValidate3.name" />
                        </el-form-item>
                        <el-form-item :label="wechatAccountLabel" prop="wechat">
                          <el-input v-model="formValidate3.wechat" />
                        </el-form-item>
                        <el-form-item :label="fundPwdLabel" prop="password">
                          <el-input v-model="formValidate3.password" type="password" />
                        </el-form-item>
                        <el-form-item>
                          <el-button type="warning" @click="handleSubmit('formValidate3')">{{ saveBtn }}</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
                  </el-form>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 账户设置页面
 * 迁移点：
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. 使用 Element Plus 组件替代 iView 组件
 * 4. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 */
import { ref, reactive, inject, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Upload } from '@element-plus/icons-vue'
import axios from 'axios'

// Vuex 3.x 和 Vue Router 3.x 兼容方案
const store = inject('store')
const router = inject('router')

// 文本常量
const pageTitle = '账户设置'
const pageTip = '绑定收款方式，用于提现'
const bankCardTitle = '银行卡'
const bankCardTip = '绑定银行卡'
const alipayTitle = '支付宝'
const alipayTip = '绑定支付宝账号'
const wechatTitle = '微信支付'
const wechatTip = '绑定微信收款码'
const modifyBtn = '修改'
const bindBtn = '绑定'
const saveBtn = '保存'
const nameLabel = '真实姓名'
const bankAccountLabel = '开户银行'
const bankBranchLabel = '支行名称'
const bankNoLabel = '银行卡号'
const confirmBankNoLabel = '确认卡号'
const fundPwdLabel = '资金密码'
const alipayAccountLabel = '支付宝账号'
const wechatAccountLabel = '微信账号'
const uploadBtn = '上传'
const imgTip = '上传凭证图片'

// 图片上传
const uploadHeaders = reactive({ 'x-auth-token': localStorage.getItem('TOKEN') })
const uploadUrl = store.state.host + '/uc/upload/oss/image'
const placeholderImg = new URL('../../assets/images/upload_placeholder.png', import.meta.url).href

const aliImg = ref('')
const aliPreview = ref('')
const weImg = ref('')
const wePreview = ref('')
const choseItem = ref(0)

const user = reactive({
  bankVerified: 0,
  aliVerified: 0,
  wechatVerified: 0,
  bankInfo: null,
  alipay: null,
  wechatPay: null,
  realName: ''
})

const formValidate1 = reactive({
  name: '',
  password: '',
  bankName: '',
  bankBranch: '',
  bankNo: '',
  bankNoConfirm: ''
})

const formValidate2 = reactive({
  name: '',
  alipay: '',
  password: ''
})

const formValidate3 = reactive({
  name: '',
  wechat: '',
  password: ''
})

const bankNameList = reactive([
  { value: '中国工商银行', label: '中国工商银行' },
  { value: '中国农业银行', label: '中国农业银行' },
  { value: '中国建设银行', label: '中国建设银行' },
  { value: '中国邮政储蓄银行', label: '中国邮政储蓄银行' },
  { value: '招商银行', label: '招商银行' },
  { value: '中国银行', label: '中国银行' },
  { value: '交通银行', label: '交通银行' },
  { value: '中信银行', label: '中信银行' },
  { value: '华夏银行', label: '华夏银行' },
  { value: '中国民生银行', label: '中国民生银行' },
  { value: '广发银行', label: '广发银行' },
  { value: '平安银行', label: '平安银行' },
  { value: '兴业银行', label: '兴业银行' },
  { value: '上海浦东发展银行', label: '上海浦东发展银行' },
  { value: '浙商银行', label: '浙商银行' },
  { value: '渤海银行', label: '渤海银行' },
  { value: '恒丰银行', label: '恒丰银行' },
  { value: '花旗银行', label: '花旗银行' },
  { value: '渣打银行', label: '渣打银行' },
  { value: '汇丰银行', label: '汇丰银行' },
  { value: '中国光大银行', label: '中国光大银行' },
  { value: '上海银行', label: '上海银行' },
  { value: '江苏银行', label: '江苏银行' },
  { value: '重庆银行', label: '重庆银行' },
  { value: '天津银行', label: '天津银行' },
  { value: '厦门银行', label: '厦门银行' },
  { value: '城市商业银行', label: '城市商业银行' },
  { value: '农村商业银行', label: '农村商业银行' },
  { value: '徽商银行', label: '徽商银行' }
])

// 表单验证规则
const validateBankNo = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入银行卡号'))
  } else if (!/([0-9]){6,18}/.test(value)) {
    callback(new Error('银行卡号格式不正确'))
  } else {
    callback()
  }
}

const validateBankNoConfirm = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请确认银行卡号'))
  } else if (!/([0-9]){6,18}/.test(value)) {
    callback(new Error('银行卡号格式不正确'))
  } else if (value !== formValidate1.bankNo) {
    callback(new Error('两次输入的卡号不一致'))
  } else {
    callback()
  }
}

const ruleValidate = reactive({
  name: [
    { required: true, message: '请输入真实姓名', trigger: 'blur' }
  ],
  bankName: [
    { required: true, message: '请选择开户银行', trigger: 'change' }
  ],
  bankBranch: [
    { required: true, message: '请输入支行名称', trigger: 'blur' }
  ],
  bankNo: [
    { required: true, message: '请输入银行卡号', trigger: 'blur' },
    { validator: validateBankNo, trigger: 'blur' }
  ],
  bankNoConfirm: [
    { required: true, message: '请确认银行卡号', trigger: 'blur' },
    { validator: validateBankNoConfirm, trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入资金密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  alipay: [
    { required: true, message: '请输入支付宝账号', trigger: 'blur' }
  ],
  wechat: [
    { required: true, message: '请输入微信账号', trigger: 'blur' }
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
const aliHandleSuccess = (res, file, fileList) => {
  if (res.code === 0) {
    aliImg.value = aliPreview.value = res.data
  } else {
    ElMessage.error(res.message)
  }
}

const weHandleSuccess = (res, file, fileList) => {
  if (res.code === 0) {
    weImg.value = wePreview.value = res.data
  } else {
    ElMessage.error(res.message)
  }
}

// 显示表单项
const showItem = (index) => {
  choseItem.value = index
}

// 表单引用
const formRefs = ref({})

// 表单提交
const submit = (name) => {
  // 银行卡
  if (name === 'formValidate1') {
    const param = {
      bank: formValidate1.bankName,
      branch: formValidate1.bankBranch,
      jyPassword: formValidate1.password,
      realName: formValidate1.name,
      cardNo: formValidate1.bankNo
    }

    const url = user.bankVerified === 1
      ? store.state.host + '/uc/approve/update/bank'
      : store.state.host + '/uc/approve/bind/bank'

    axios.post(url, param).then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('保存成功')
        getAccount()
        choseItem.value = 0
      } else {
        ElMessage.error(resp.message)
      }
    })
  }

  // 支付宝
  if (name === 'formValidate2') {
    const param = {
      ali: formValidate2.alipay,
      jyPassword: formValidate2.password,
      realName: formValidate2.name,
      qrCodeUrl: aliPreview.value
    }

    const url = user.aliVerified === 1
      ? store.state.host + '/uc/approve/update/ali'
      : store.state.host + '/uc/approve/bind/ali'

    axios.post(url, param).then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('保存成功')
        getAccount()
        choseItem.value = 0
      } else {
        ElMessage.error(resp.message)
      }
    })
  }

  // 微信
  if (name === 'formValidate3') {
    const param = {
      wechat: formValidate3.wechat,
      jyPassword: formValidate3.password,
      realName: formValidate3.name,
      qrCodeUrl: wePreview.value
    }

    const url = user.wechatVerified === 1
      ? store.state.host + '/uc/approve/update/wechat'
      : store.state.host + '/uc/approve/bind/wechat'

    axios.post(url, param).then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('保存成功')
        getAccount()
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

// 获取账户信息
const getAccount = () => {
  axios.post(store.state.host + '/uc/approve/account/setting').then(response => {
    const resp = response.data
    if (resp.code === 0) {
      user.bankVerified = resp.data.bankVerified
      user.aliVerified = resp.data.aliVerified
      user.wechatVerified = resp.data.wechatVerified
      user.bankInfo = resp.data.bankInfo
      user.alipay = resp.data.alipay
      user.wechatPay = resp.data.wechatPay
      user.realName = resp.data.realName

      formValidate1.name = formValidate2.name = formValidate3.name = user.realName
      formValidate1.bankName = user.bankInfo ? user.bankInfo.bank : ''
      formValidate1.bankBranch = user.bankInfo ? user.bankInfo.branch : ''
      formValidate1.bankNo = user.bankInfo ? user.bankInfo.cardNo : ''
      formValidate2.alipay = user.alipay ? user.alipay.aliNo : ''
      formValidate3.wechat = user.wechatPay ? user.wechatPay.wechat : ''
      aliImg.value = aliPreview.value = user.alipay ? user.alipay.qrCodeUrl : ''
      weImg.value = wePreview.value = user.wechatPay ? user.wechatPay.qrWeCodeUrl : ''
    } else {
      ElMessage.error(resp.message)
      router.push('/uc/safe')
    }
  })
}

// 生命周期
onMounted(() => {
  getAccount()
})
</script>

<style scoped lang="scss">
.account-box .account-in .account-item .account-detail {
  padding: 30px 0;
  margin: 6px 0;
}

.account-box .account-in .account-item .account-detail .detail-list {
  width: 40%;
  width: 80%;
  margin: 0 auto;
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
  color: rgb(130, 142, 161);
  font-size: 13px;
}

.account-box .account-in .account-item .account-item-in .btn {
  padding: 8px 10px;
  cursor: pointer;
  color: #f0a70a;
}

.tips-g {
  color: rgb(130, 142, 161);
  font-size: 12px;
}

.table-inner {
  position: relative;
  text-align: left;
  border-radius: 3px;
  padding: 23px 20px 20px;
}

.acb-p1 {
  font-size: 18px;
  font-weight: 600;
  line-height: 50px;
}

.acb-p2 {
  font-size: 14px;
  line-height: 24px;
}

.action-inner {
  width: 100%;
  display: table;
}

.action-inner .inner-box {
  display: table-cell;
  width: 100%;
}

.action-box .title .copy {
  user-select: text;
}

.action-box .title a.link-copy {
  font-size: 14px;
  margin-left: 20px;
}

.hb-night a {
  text-decoration: none;
  color: #f0a70a;
  transition: all .2s ease-in-out;
  cursor: pointer;
}

.action-box .title a.link-qrcode {
  margin-left: 20px;
  font-size: 14px;
  position: relative;
}

.action-box .subtitle {
  font-size: 12px;
  margin-top: 30px;
}

.action-content {
  width: 100%;
  margin-top: 30px;
  overflow: hidden;
  display: table;
}

.action-box .title {
  margin-top: 20px;
  font-size: 20px;
  user-select: none;
}

.action-box .title .show-qrcode {
  position: absolute;
  top: -100px;
  left: 40px;
  padding: 10px;
}

.action-inner .inner-box.deposit-address {
  width: 80%;
}

p.describe {
  font-size: 16px;
  font-weight: 600;
}

.merchant-top {
  height: 50px;
  display: flex;
  align-items: center;
  padding: 0 15px;
}

.trade-group {
  margin-bottom: 20px;
  font-size: 14px;
}

.merchant-icon {
  display: inline-block;
  margin-left: 4px;
  background-size: 100% 100%;
}

.merchant-top .tips-word {
  flex-grow: 2;
  text-align: left;
}

.merchant-icon.tips {
  width: 4px;
  height: 22px;
  margin-right: 10px;
  background: #f0a70a;
}

.bill_box {
  width: 100%;
  height: auto;
  overflow: hidden;
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

.rightarea .rightarea-con {
  padding-top: 30px;
  padding-bottom: 125px;
}

.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 15px;
}

.order_box {
  width: 100%;
  background: #fff;
  height: 56px;
  line-height: 56px;
  margin-bottom: 20px;
  border-bottom: 2px solid #ccf2ff;
  position: relative;
  text-align: left;
}

.order_box a {
  color: #8994A3;
  font-size: 16px;
  padding: 0 30px;
  cursor: pointer;
  text-decoration: none;
  text-align: center;
  line-height: 54px;
  display: inline-block;
}

.order_box .active {
  border-bottom: 2px solid #f0a70a;
}

.order_box .search {
  position: absolute;
  width: 300px;
  height: 32px;
  top: 12px;
  right: 0;
  display: flex;
}

.el-button.el-button--warning {
  background-color: #f0a70a;
  border-color: #f0a70a;
}

@media screen and (max-width: 768px) {
  .uc_account .nav-right {
    padding: 0 0 !important;
  }

  .uc_account .account-box .account-in .account-item .account-item-in .card-number {
    padding: 0px 5px !important;
  }

  .uc_account .merchant-top {
    padding: 0 0 !important;
  }

  .account-box .account-in .account-item .account-detail {
    padding: 0 0;
    margin: 6px 0;
  }

  .acc_sc {
    margin-bottom: 20px;
  }

  .shoujiBtn {
    margin-left: -90px !important;
  }
}
</style>

<style lang="scss">
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
</style>

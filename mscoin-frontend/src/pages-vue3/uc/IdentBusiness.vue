<template>
  <div class="identbusiness" style="padding: 81px;padding-top: 80px;">
    <div class="content">
      <div style="width: 80%;margin: 0 auto;margin-bottom: 60px;">
        <div class="ident-title" v-if="certStatus === 0">
          <h3>{{$t('uc.identity.apply')}}</h3>
          <p style="font-size: 14px;margin-top: 10px"></p>
        </div>
        <div class="ident-title" v-else-if="certStatus == 1">
          <h3>{{$t('uc.identity.tijiao')}}</h3>
        </div>
        <div class="ident-title" v-else-if="certStatus == 2">
          <h3>{{$t('uc.identity.tijiaosuc')}}</h3>
        </div>
        <div class="ident-title" v-else-if="certStatus == 3">
          <h3>{{$t("uc.identity.tijiaofail")}}</h3>
        </div>
        <div class="ident-title" v-else-if="certStatus == 5">
          <h3>{{$t("uc.identity.zhuxiaotijiao")}}</h3>
        </div>
        <div class="ident-title" v-else-if="certStatus == 6">
          <h3>{{$t("uc.identity.shenhefail")}}</h3>
        </div>
        <div class="ident-title" v-else-if="certStatus == 7">
          <h3>{{$t("uc.identity.shenhesuc")}}</h3>
        </div>
        <ElSteps class="apply-step" :current="certStatus == 2 ? 3 : certStatus == 3 ? 2 : certStatus" :status="certStatus == 3 ? 'error' :'finish'" v-if="certStatus != 0 && certStatus != 5 && certStatus != 6 && certStatus != 7">
          <ElStep :title="prepare"></ElStep>
          <ElStep :title="review"></ElStep>
          <ElStep :title="certStatus == 1 || certStatus == 0 ? result : certStatus == 2 ? certified : shenheshibai"></ElStep>
        </ElSteps>
        <ElSteps class="apply-step" :current="certStatus == 5 ? 1 : certStatus == 6 ? 2 : 3" :status="certStatus == 6 ? 'error':'finish'" v-if="certStatus == 5 || certStatus == 6 || certStatus == 7">
          <ElStep :title="shangjiazhuxiao"></ElStep>
          <ElStep :title="tijiaoshenqing"></ElStep>
          <ElStep :title="certStatus == 5 ? result : certStatus == 6 ? shenheshibai : passed"></ElStep>
        </ElSteps>

        <div v-if="certStatus == 6" style="width: 500px;margin: 0 auto;text-align: center;">
          <el-button type="warning" style="width: 120px;background:#f0ac19;border-color:#f0ac19" @click="modal_return=true" long size="large">{{$t("uc.identity.shenagain")}}</el-button>
          <div class="fail-reason" style="margin-top: 50px;font-size: 16px;">
            <el-icon style="color:red;font-size:16px;vertical-align: middle;"><InfoFilled /></el-icon>
            <span style="margin-left: 10px;">{{$t('uc.identity.yuanyin')}}:{{refuseReason}}</span>
          </div>
        </div>

        <div v-if="certStatus == 7" style="width: 500px;margin: 0 auto;text-align: center;">
          <el-button type="warning" style="width: 120px;background:#f0ac19;border-color:#f0ac19" @click="modal_read=true" long size="large">{{$t("uc.identity.sheqinggain")}}</el-button>
        </div>

        <div v-if="certStatus == 3" style="width: 500px;margin: 0 auto;text-align: center;">
          <el-button type="warning" style="width: 120px;background:#f0ac19;border-color:#f0ac19" @click="modal_read=true" long size="large">{{$t("uc.identity.shenagain")}}</el-button>
          <div class="fail-reason" style="margin-top: 50px;font-size: 16px;">
            <el-icon style="color:red;font-size:16px;vertical-align: middle;"><InfoFilled /></el-icon>
            <span style="margin-left: 10px;">{{$t("uc.identity.reason")}}:{{certReason}}</span>
          </div>
        </div>

        <div v-else-if="certStatus == 2" style="width: 500px;margin: 0 auto;text-align: center;">
          <el-button type="warning" style="width: 120px;background:#f0ac19;border-color:#f0ac19" @click="publishAd" long size="large">{{$t('nav.fabu')}}</el-button>
          <div style="margin-top: 30px;font-size: 16px;text-align: center;">
            <a @click="returnAdit" style="color: #aaa;">{{$t("uc.identity.shenqingtuibao")}}</a>
          </div>
        </div>
      </div>
      <div class="ipshang" :class="certStatus != 0 ? 'applying' : ''">
        <div class="ident-title" v-if="certStatus == 3">
          <h3 style="font-size: 20px">{{$t('uc.identity.apply')}}</h3>
          <p style="font-size: 14px;margin-top: 10px">{{$t('uc.identity.become')}}</p>
        </div>
        <div class="ident-title" v-else-if="certStatus == 2">
          <h3>{{$t("uc.identity.getquan")}}</h3>
        </div>
        <el-row style="margin-top:40px;">
          <el-col :span="8">
            <div class="business-function">
              <img alt="" src="../../assets/images/business_show.png" width="300px">
              <p style="padding: 20px 0;font-weight: 600;font-size: 18px">{{$t('uc.identity.seat')}}</p>
              <span style="font-size: 14px;overflow:hidden; overflow: hidden;text-overflow:ellipsis;display: block;white-space:nowrap;color:#999;">{{$t("uc.identity.zhusnhu")}}</span>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="business-function">
              <img alt="" src="../../assets/images/business_service.png" width="300px">
              <p style="padding: 20px 0;font-weight: 600;font-size: 18px">{{$t('uc.identity.service')}}</p>
              <span style="font-size: 14px;color:#999;">{{$t("uc.identity.service")}}</span>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="business-function">
              <img alt="" src="../../assets/images/business_fee.png" width="300px">
              <p style="padding: 20px 0;font-weight: 600;font-size: 18px">{{$t('uc.identity.lowfee')}}</p>
              <span style="font-size: 14px;color:#999;">{{$t("uc.identity.lowfee")}}</span>
            </div>
          </el-col>
        </el-row>
        <div v-show="certStatus === 0" style="text-align: center;font-size: 16px;margin-top:50px">
          <el-checkbox v-model="single"></el-checkbox>
          <span>{{$t("uc.identity.read")}}</span>
          <router-link target="_blank" to="/helpdetail?cate=1&id=11&cateTitle=常见问题" class="cur" style="color:#f0a70a">{{$t('uc.identity.agreement')}}</router-link>
        </div>
        <div v-show="certStatus === 0" class="sq">
          <el-button @click="apply" style="background:#f0a70a;color:#fff;outline:none;">{{$t("uc.identity.lijishenqing")}}</el-button>
        </div>
      </div>
    </div>

    <el-dialog v-model="modal_read">
      <template #title>
        <span class="tit">{{$t('uc.identity.second.line')}}</span>
      </template>
      <div class="apply-note">
        <h3 style="padding-top: 10px;">{{$t('uc.identity.second.step1')}}</h3>
        <p>{{$t('uc.identity.second.step1c1')}}<br>{{$t('uc.identity.second.step1c2')}}</p>
        <h3>{{$t('uc.identity.second.step2')}}</h3>
        <p>{{$t('uc.identity.second.step2c')}}</p>
        <h3>{{$t('uc.identity.second.step3')}}</h3>
        <p>{{$t('uc.identity.second.stepc')}}</p>
        <div style="text-align: left;padding: 30px 0;">
          <el-checkbox v-model="agreeFrozen"></el-checkbox> {{$t('uc.identity.second.agree')}}
          <span>
            <font color="#f0a70a">{{auditText}}</font>{{$t('uc.identity.second.agreec')}}</span>
        </div>
        <el-button @click="apply2" long style="font-size: 16px;background:#f0a70a;color:#fff;border:1px solid #f0a70a;">{{$t('uc.identity.second.shenqingchngweishangjia')}}</el-button>
      </div>
    </el-dialog>

    <el-dialog v-model="modal_apply">
      <div class="apply-content">
        <div class="apply-title">
          <h3>{{$t("uc.identity.tijiaoziliao")}}</h3>
          <p>{{$t("uc.identity.place")}}</p>
        </div>
        <el-form class="apply-form" :model="apply_form" label-position="top">
          <el-form-item :label="phone">
            <el-input type="text" v-model="apply_form.telno" :placeholder="noEmpty"></el-input>
          </el-form-item>
          <el-form-item :label="wechat">
            <el-input type="text" v-model="apply_form.wechat" :placeholder="noEmpty"></el-input>
          </el-form-item>
          <el-form-item :label="qq">
            <el-input type="text" v-model="apply_form.qq" :placeholder="noEmpty"></el-input>
          </el-form-item>
          <el-row>
            <el-col :span="8">
              <el-form-item :label="bizhong">
                <el-select v-model="apply_form.coinSymbol" :placeholder="select" @change="onCoinChange">
                  <el-option v-for="(item,index) in auditCurrency" :value="item.coin.unit" :key="index"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <span>&nbsp;</span>
            </el-col>
            <el-col :span="8">
              <el-form-item :label="shuliang">
                <span>{{ apply_form.amount }}</span>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
              <el-upload ref="upload1" :on-success="assetHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :on-remove="assetRemove">
                <span style="line-height: 100px;font-size: 50px;color:#ccc;">+</span>
                <img v-show="assetImg" class="previewImg" :src="assetImg" alt="">
              </el-upload>
              <span>{{$t("uc.identity.gerenzichan")}}</span>
            </el-col>
            <el-col :span="8">
              <span>&nbsp;</span>
            </el-col>
            <el-col :span="8">
              <el-upload ref="upload2" :on-success="tradeHandleSuccess" :headers="uploadHeaders" :action="uploadUrl" :on-remove="tradeRemove">
                <span style="line-height: 100px;font-size: 50px;color:#ccc;">+</span>
                <img v-show="tradeImg" class="previewImg" :src="tradeImg" alt="">
              </el-upload>
              <span>{{$t("uc.identity.shuzizichan")}}</span>
            </el-col>
          </el-row>
          <el-form-item style="margin-top: 20px;">
            <el-button style="width:100%;background:#f0a70a;color:#fff;border:1px solid #f0a70a;" type="info" @click="apply3" :disabled="applyBtn">{{$t("uc.identity.lijishenqing")}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-dialog>

    <el-dialog v-model="modal_return" @close="returnAudit">
      <template #title>
        <p style="text-align: center;">{{$t("uc.identity.tips")}}</p>
      </template>
      <p style="text-align: center;font-size: 14px;">{{$t("uc.identity.wufachexiao")}}</p>
      <p style="text-align: center;font-size: 14px;">{{$t("uc.identity.suredo")}}</p>
      <el-input v-model="returnReason" type="textarea" :placeholder="placeholder" :rows="4"></el-input>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 企业认证页面
 */
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { ElMessage, ElNotification, ElButton, ElCheckbox, ElInput, ElSelect, ElOption, ElForm, ElFormItem, ElRow, ElCol, ElUpload, ElDialog, ElIcon, ElSteps, ElStep } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const router = useRouter()
const store = useStore()

const host = ''

const noEmpty = ref('必填')
const review = ref('准备资料')
const prepare = ref('提交审核')
const result = ref('等待结果')
const certified = ref('已认证')
const shenheshibai = ref('审核失败')
const shangjiazhuxiao = ref('商家注销')
const tijiaoshenqing = ref('提交申请')
const passed = ref('审核通过')
const placeholder = ref('请填写说明')
const select = ref('请选择')
const phone = ref('手机号')
const qq = ref('QQ 号')
const wechat = ref('微信号')
const bizhong = ref('币种')
const shuliang = ref('数量')

const single = ref(false)
const isShowShang = ref(true)
const isShowMailt = ref(false)
const isShowSubmitted = ref(false)
const isShowSuccess = ref(false)
const activeStepIndex = ref(0)
const certStatus = ref(0)
const certReason = ref('')
const auditCurrency = ref('')
const auditText = ref('')
const modal_read = ref(false)
const modal_return = ref(false)
const agreeFrozen = ref(false)
const modal_apply = ref(false)
const applyBtn = ref(false)

const apply_form = reactive({
  telno: '',
  wechat: '',
  qq: '',
  coinSymbol: '',
  amount: '',
  assetData: '',
  tradeData: ''
})

const assetImg = ref('')
const tradeImg = ref('')
const uploadHeaders = computed(() => ({ 'x-auth-token': localStorage.getItem('TOKEN') }))
const uploadUrl = ref(host + '/uc/upload/oss/image')
const returnReason = ref('')
const refuseReason = ref('')

const lang = computed(() => store.state.lang)

const islogin = () => {
  axios.post(`${host}/uc/approve/security/setting`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      if (!res.data.realName) {
        ElMessage.warning('请先完成实名认证')
        router.push('/uc/safe')
      } else if (res.data.phoneVerified === 0) {
        ElMessage.warning('请先绑定手机号')
        router.push('/uc/safe')
      } else if (res.data.fundsVerified === 0) {
        ElMessage.warning('请先设置资金密码')
        router.push('/uc/safe')
      }
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const publishAd = () => {
  router.push('/uc/ad/create')
}

const returnAdit = () => {
  modal_return.value = true
}

const returnAudit = () => {
  const params = { detail: returnReason.value }
  axios.post(`${host}/uc/approve/cancel/business`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      ElMessage.success('提交成功!')
      modal_return.value = false
      getSetting()
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const getAudiCoin = (symbol) => {
  for (let i = 0; i < auditCurrency.value.length; i++) {
    if (symbol === auditCurrency.value[i].coin.unit) {
      return auditCurrency.value[i]
    }
  }
  return null
}

const onCoinChange = (value) => {
  const coin = getAudiCoin(value)
  if (coin !== null) {
    apply_form.amount = coin.amount
  }
}

const getSetting = () => {
  axios.get(`${host}/uc/approve/business/setting`, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      const certifiedBusinessStatus = res.data.data.certifiedBusinessStatus
      activeStepIndex.value = certifiedBusinessStatus
      certStatus.value = certifiedBusinessStatus
      certReason.value = res.data.data.detail
      refuseReason.value = res.data.data.reason
    }
  }).catch(() => {})
}

const assetHandleSuccess = (res, file, fileList) => {
  apply_form.assetData = res.data
  assetImg.value = res.data
}

const tradeHandleSuccess = (res, file, fileList) => {
  apply_form.tradeData = res.data
  tradeImg.value = res.data
}

const assetRemove = () => {
  apply_form.assetData = ''
  assetImg.value = ''
}

const tradeRemove = () => {
  apply_form.tradeData = ''
  tradeImg.value = ''
}

const getAuthFound = () => {
  axios.get(`${host}/uc/approve/business-auth-deposit/list`, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      auditCurrency.value = res.data.data
      let tempText = ''
      for (let i = 0; i < res.data.data.length; i++) {
        if (i === 0) {
          apply_form.coinSymbol = res.data.data[i].coin.unit
          apply_form.amount = res.data.data[i].amount
        }
        tempText += res.data.data[i].amount + '个' + res.data.data[i].coin.unit
        if (i < res.data.data.length - 1) tempText += '或'
      }
      auditText.value = tempText
    }
  }).catch(() => {})
}

const apply = () => {
  if (!single.value) {
    ElMessage.warning('请同意协议')
    return
  }
  modal_read.value = true
}

const apply2 = () => {
  if (!agreeFrozen.value) {
    ElMessage.warning('请同意冻结相应数量的币')
    return
  }
  modal_read.value = false
  modal_apply.value = true
}

const apply3 = () => {
  if (!apply_form.telno) {
    ElMessage.error('请填写手机号')
    return
  }
  if (!apply_form.wechat) {
    ElMessage.error('请填写微信号')
    return
  }
  if (!apply_form.qq) {
    ElMessage.error('请填写 QQ 号')
    return
  }
  if (!apply_form.assetData) {
    ElMessage.error('请上传资产证明')
    return
  }
  if (!apply_form.tradeData) {
    ElMessage.error('请上传交易证明')
    return
  }

  const params = {
    businessAuthDepositId: getAudiCoin(apply_form.coinSymbol).id,
    json: JSON.stringify(apply_form)
  }
  axios.post(`${host}/uc/approve/certified/business/apply`, params, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      ElMessage.success('提交成功!')
      modal_apply.value = false
      certStatus.value = 1
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const updateLang = () => {
  prepare.value = '提交审核'
  review.value = '准备资料'
  result.value = '等待结果'
  certified.value = '已认证'
  shenheshibai.value = '审核失败'
  shangjiazhuxiao.value = '商家注销'
  tijiaoshenqing.value = '提交申请'
  passed.value = '审核通过'
  phone.value = '手机号'
  qq.value = 'QQ 号'
  wechat.value = '微信号'
  bizhong.value = '币种'
  shuliang.value = '数量'
}

watch(lang, () => {
  updateLang()
})

onMounted(() => {
  islogin()
  getSetting()
  getAuthFound()
})
</script>

<style scoped>
.previewImg {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
}
.content {
  width: 1200px;
  margin: 0 auto;
  padding-top: 30px;
  height: 100%;
  background: #192330;
}
.ipshang {
  overflow: hidden;
}
.ipshang.applying {
  background: #1f2833;
  padding: 40px 0;
}
.sq {
  width: 1200px;
  margin-top: 50px;
  text-align: center;
  margin-bottom: 50px;
}
.sq button {
  height: 50px;
  font-size: 18px;
  width: 450px;
}
.tit {
  font-size: 16px;
  line-height: 25px;
  border-left: 5px solid #f0a70a;
  padding-left: 15px;
}
.business-function {
  width: 300px;
  margin: 0 auto;
  border: none;
  background-color: #27313e;
  padding-bottom: 20px;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
}
.fail-reason {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>

<style>
.el-form-item {
  margin-bottom: 24px;
}
.identbusiness .el-steps .el-step__head {
  background-color: #192330;
}
.identbusiness .el-steps .el-step__title {
  background-color: #192330;
  color: #fff;
}
.el-button--primary {
  background: #f0a70a;
  border: 1px solid #f0a70a;
}
.el-button--primary:hover {
  background: #f0a70a;
  border: 1px solid #f0a70a;
}
.el-checkbox__input.is-checked .el-checkbox__inner {
  background-color: #f0a70a !important;
  border: 1px solid #f0a70a !important;
}
</style>

<template>
  <div class="my_ad_container">
    <div class="contbox">
      <div class="" v-loading="isSpinShow">
        <div class="send-box">
          <div class="title-box">
            <h6 class="titles">{{$t('otc.publishad.createad')}}
              <i class="iconfont icon-hongjiantou"></i>
            </h6>
            <p>{{$t('otc.publishad.msg1')}}
              <router-link to="/otc/trade/usdt">{{$t('otc.publishad.tradead')}}</router-link>
              。
            </p>
            <p>{{$t('otc.publishad.msg2')}}{{$t('otc.publishad.msg3')}}。</p>
            <p>{{$t('otc.publishad.msg4')}}
              <router-link to="/uc/ad">{{$t('otc.publishad.myad')}}</router-link>
              。
            </p>
          </div>
          <div class="formbox send-form">
            <el-form :model="form" :rules="ruleValidate" :label-width="90" ref="formRef">
              <el-form-item :label="$t('otc.publishad.iwant')" prop="advertiseType">
                <el-radio-group v-model="form.advertiseType">
                  <el-radio label="1" :disabled="isId">{{$t('otc.publishad.sellonline')}}</el-radio>
                  <el-radio label="0" :disabled="isId">{{$t('otc.publishad.buyonline')}}</el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.exchangecoin')" prop="coin">
                <el-select v-model="form.coin" :disabled="isId" @change="changeCoin">
                  <el-option v-for="(item, index) in coinList" :value="item.id" :key="index">{{item.unit}}</el-option>
                </el-select>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.country')" prop="country">
                <el-select v-model="form.country" @change="onAreaChange">
                  <el-option v-for="(area,index) in areas" :value="area.zhName" :key="index">{{area.zhName}}</el-option>
                </el-select>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.currency')" prop="rmb">
                <el-input v-model="form.rmb" disabled placeholder=""></el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.openfixedprice')">
                <el-switch v-model="form.fixed" inline-prompt active-text="{{$t('otc.publishad.open')}}" inactive-text="{{$t('otc.publishad.close')}}"></el-switch>
              </el-form-item>
              <p class="msg" v-show="form.fixed">{{$t('otc.publishad.usetip')}}</p>
              <el-form-item :label="$t('otc.publishad.premiseprice')" prop="premisePrice" v-show="!form.fixed">
                <el-input v-model="form.premisePrice" :placeholder="$t('otc.publishad.premisepricetip')">
                  <template #append>%</template>
                </el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.fixedprice')" prop="fixedPrice" v-show="form.fixed">
                <el-input v-model="form.fixedPrice" :placeholder="$t('otc.publishad.fixedpricetip')">
                  <template #append>{{form.rmb}}</template>
                </el-input>
              </el-form-item>
              <p class="msg">{{$t('otc.publishad.marketprice')}}：
                <span class="cankao">{{cankao}}</span>
              </p>
              <p class="msg" v-show="!form.fixed">{{$t('otc.publishad.marketpricetip')}}{{wantstyle}}。</p>
              <div class="el-form-item">
                <label class="el-form-item-label" style="width: 90px;">{{$t('otc.publishad.exchangeprice')}}</label>
                <div class="el-form-item-content" style="margin-left: 90px;">
                  <div class="el-input-wrapper el-input-type" id="price">
                    {{price}}&nbsp;CNY/{{symbol}}
                  </div>
                </div>
              </div>
              <p class="msg">{{$t('otc.publishad.formual')}}:（Bitstamp+Bitfinex+Coinbase）/ 3 *{{gongshi.toFixed(4) }}</p>
              <el-form-item :label="wantstyle+$t('otc.publishad.num')" prop="number">
                <el-input v-model="form.number" :placeholder="$t('otc.publishad.num_text1')+wantstyle+$t('otc.publishad.num_text2')"></el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.exchangeperiod')" prop="timeLimit">
                <el-input v-model="form.timeLimit" :placeholder="$t('otc.publishad.exchangeperiod_text1')+'('+wantTime+$t('otc.publishad.minute')+')'">
                  <template #append>{{$t('otc.publishad.minute')}}</template>
                </el-input>
              </el-form-item>
              <p class="msg">{{$t('otc.publishad.tip1')}} </p>
              <router-link to="/uc/account" style="padding-left: 90px;color:#f0ac19;">{{$t('otc.publishad.tip2')}}</router-link>
              <el-form-item :label="$t('otc.publishad.paymode')" prop="payMode">
                <el-select v-model="form.payMode" multiple>
                  <el-option v-for="(item,index) in payModeList" :value="item.value" :key="item.value" :disabled="item.isOpen">{{ item.label }}</el-option>
                </el-select>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.minlimit')" prop="minLimit">
                <el-input v-model="form.minLimit" :placeholder="$t('otc.publishad.tip3')">
                  <template #append>CNY</template>
                </el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.maxlimit')" prop="maxLimit">
                <el-input v-model="form.maxLimit" :placeholder="$t('otc.publishad.tip4')">
                  <template #append>CNY</template>
                </el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.remark')" prop="remark">
                <el-input v-model="form.remark" type="textarea" :rows="4" :maxlength="500" :placeholder="$t('otc.publishad.tip5')"></el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.openautoreply')">
                <el-switch v-model="form.autoReply" inline-prompt active-text="{{$t('otc.publishad.open')}}" inactive-text="{{$t('otc.publishad.close')}}"></el-switch>
              </el-form-item>
              <p class="msg">{{$t('otc.publishad.msg5')}}</p>
              <el-form-item :label="$t('otc.publishad.autoreply')" prop="autoword" v-show="form.autoReply">
                <el-input v-model="form.autoword" type="textarea" :rows="4" :maxlength="500" :placeholder="$t('otc.publishad.autoreplytip')"></el-input>
              </el-form-item>
              <el-form-item :label="$t('otc.publishad.fundpwd')" prop="priceW">
                <el-input v-model="form.priceW" :placeholder="$t('otc.publishad.fundpwdtip')" type="password" show-password></el-input>
              </el-form-item>
              <el-form-item>
                <el-button style="background:#f0a70a;color:#fff;border:1px solid #f0a70a;" long @click="handleSubmit('form')" :disabled="disAllowBtn">{{$t('otc.publishad.submit')}}</el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 广告发布页面
 */
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { ElMessage, ElForm, ElFormItem, ElInput, ElSelect, ElOption, ElRadioGroup, ElRadio, ElSwitch, ElButton } from 'element-plus'
import axios from 'axios'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const host = ''

const isSpinShow = ref(true)
const isId = ref(false)
const disAllowBtn = ref(false)
const formRef = ref(null)

const coinList = ref([])
const balance = ref(100)
const cankao = ref('')
const price = ref('')
const symbol = ref('')
const gongshi = ref(1)
const wantstyle = ref('')
const wantTime = ref('15-120')
const areas = ref([])

const form = reactive({
  advertiseType: '1',
  coin: '1',
  country: 'china',
  rmb: '',
  fixed: false,
  premisePrice: '',
  fixedPrice: '',
  number: '',
  timeLimit: '',
  payMode: [],
  minLimit: '',
  maxLimit: '',
  autoReply: false,
  remark: '',
  priceW: '',
  autoword: ''
})

const payModeList = ref([
  {
    value: '支付宝',
    label: '支付宝',
    isOpen: true
  },
  {
    value: '微信',
    label: '微信',
    isOpen: true
  },
  {
    value: '银行卡',
    label: '银行卡',
    isOpen: true
  }
])

// Validation rules
const numberCheck = (rule, value, callback) => {
  if (value === '' || value == 0) {
    callback(new Error('请输入有效的数量'))
  } else if (Number.isNaN(Number(value))) {
    callback(new Error('请输入有效的数量'))
  } else {
    callback()
  }
}

const premisePriceCheck = (rule, value, callback) => {
  if (form.fixed == false) {
    if (!value || value == 0) {
      return callback(new Error('请输入有效的溢价率'))
    } else if (Number.isNaN(Number(value))) {
      callback(new Error(''))
    } else if (value < 0 || value > 20) {
      callback(new Error('溢价率必须在 0-20 之间'))
    } else {
      callback()
    }
  } else {
    callback()
  }
}

const fixedPCheck = (rule, value, callback) => {
  if (form.fixed == true) {
    if (value === '') {
      callback(new Error('请输入有效的固定价格'))
    } else if (!/^[0-9]+(.[0-9]{2})?/.test(value)) {
      callback(new Error('请输入有效的固定价格'))
    } else {
      callback()
    }
  } else {
    callback()
  }
}

const maxCheck = (rule, value, callback) => {
  let priceNow = (price.value + '').replace(/,/g, '').replace(/[^\d|.]/g, '') - 0
  priceNow = round(mul(priceNow, form.number), 8)

  if (!value || value == 0) {
    return callback(new Error('请输入有效的最大限额'))
  } else if (!/^[0-9]+(.[0-9]{2})?$/.test(value)) {
    callback(new Error('请输入有效的数字'))
  } else if (parseFloat(value) < parseFloat(form.minLimit)) {
    callback(new Error('最大限额不能小于最小限额'))
  } else if (parseFloat(value) > parseFloat(priceNow)) {
    callback(new Error('最大限额不能超过 ' + priceNow + 'CNY！'))
  } else {
    callback()
  }
}

const minCheck = (rule, value, callback) => {
  if (!value || value == 0) {
    return callback(new Error('请输入有效的最小限额'))
  } else if (!/^\d+$/.test(value)) {
    callback(new Error('请输入有效的数字'))
  } else if (parseFloat(value) < 100) {
    callback(new Error('最小限额不能低于 100'))
  }
  if (form.maxLimit && form.maxLimit != 0 && parseFloat(value) > form.maxLimit - 0) {
    callback(new Error('最小限额不能大于最大限额'))
  } else {
    callback()
  }
}

const timeLimitCheck = (rule, value, callback) => {
  if (value === '' || !/^\d+$/.test(value)) {
    callback(new Error('请输入有效的交易期限'))
  }
  if (form.advertiseType == '1') {
    if (value < 15 || value > 120) {
      callback(new Error('出售交易期限必须在 15-120 分钟之间'))
    } else {
      callback()
    }
  }
  if (form.advertiseType == '0') {
    if (value < 10 || value > 30) {
      callback(new Error('买入交易期限必须在 10-30 分钟之间'))
    } else {
      callback()
    }
  } else {
    callback()
  }
}

const ruleValidate = reactive({
  advertiseType: [{ required: true, trigger: 'change' }],
  coin: [
    {
      required: true,
      message: '请选择币种',
      trigger: 'change'
    }
  ],
  country: [
    {
      required: true,
      message: '请选择国家',
      trigger: 'change'
    }
  ],
  rmb: [
    {
      required: true,
      message: '请输入法币',
      trigger: 'change'
    }
  ],
  premisePrice: [
    {
      validator: premisePriceCheck,
      message: '请输入有效的溢价率',
      trigger: 'blur'
    }
  ],
  fixedPrice: [
    {
      validator: fixedPCheck,
      message: '请输入有效的固定价格',
      trigger: 'blur'
    }
  ],
  number: [
    {
      required: true,
      validator: numberCheck,
      message: '请输入有效的数量',
      trigger: 'blur'
    }
  ],
  timeLimit: [
    {
      required: true,
      validator: timeLimitCheck,
      message: '请输入有效的交易期限',
      trigger: 'blur'
    }
  ],
  payMode: [
    {
      required: true,
      type: 'array',
      min: 1,
      message: '请至少选择一种支付方式',
      trigger: 'change'
    }
  ],
  minLimit: [{ required: true, validator: minCheck, trigger: 'blur' }],
  maxLimit: [{ required: true, validator: maxCheck, trigger: 'blur' }],
  priceW: [
    {
      required: true,
      message: '请输入资金密码',
      trigger: 'blur'
    }
  ]
})

const getAvatarInitial = (name) => {
  return (name + '').replace(/^\s+|\s+$/g, '').slice(0, 1)
}

const strpro = (str) => {
  let newStr = str
  str = str.slice(1)
  const re = /[\D\d]*/g
  const str2 = str.replace(re, function (s) {
    let result = ''
    for (let i = 0; i < s.length; i++) {
      result += '*'
    }
    return result
  })
  return newStr.slice(0, 1) + str2
}

const mul = (a, b) => {
  var c = 0,
    d = a.toString(),
    e = b.toString()
  try {
    c += d.split('.')[1].length
  } catch (f) {}
  try {
    c += e.split('.')[1].length
  } catch (f) {}
  return (
    Number(d.replace('.', '')) *
    Number(e.replace('.', '')) /
    Math.pow(10, c)
  )
}

const div = (a, b) => {
  var c,
    d,
    e = 0,
    f = 0
  try {
    e = a.toString().split('.')[1].length
  } catch (g) {}
  try {
    f = b.toString().split('.')[1].length
  } catch (g) {}
  return (
    (c = Number(a.toString().replace('.', ''))),
    (d = Number(b.toString().replace('.', ''))),
    mul(c / d, Math.pow(10, f - e))
  )
}

const round = (v, e) => {
  var t = 1
  for (; e > 0; t *= 10, e--)
  for (; e < 0; t /= 10, e++)
  return Math.round(v * t) / t
}

const changeCoin = () => {
  let coinItem = getCoin(form.coin)
  if (coinItem != null) {
    cankao.value = coinItem.marketPrice + ''
    let lv = (1 + form.premisePrice / 100).toFixed(4)
    let cankoNew = cankao.value.replace(/,/g, '').replace(/[^\d|.]/g, '') - 0
    price.value = (cankoNew * lv).toLocaleString()
    symbol.value = coinItem.unit
  }
}

const handleSubmit = (name) => {
  formRef.value.validate(valid => {
    if (valid) {
      disAllowBtn.value = true
      var params = {}
      params['price'] = (price.value + '').replace(/[^\d|.]/g, '') - 0
      params['advertiseType'] = form.advertiseType
      params['coin.id'] = form.coin
      params['minLimit'] = form.minLimit
      params['maxLimit'] = form.maxLimit
      params['timeLimit'] = form.timeLimit
      params['country'] = form.country
      if (form.fixed == true) {
        params['priceType'] = 0
      } else if (form.fixed == false) {
        params['priceType'] = 1
      }
      params['premiseRate'] = form.premisePrice
      params['remark'] = form.remark
      params['number'] = form.number
      params['pay'] = form.payMode
      params['jyPassword'] = form.priceW
      if (form.autoReply == true) {
        params['auto'] = 1
      } else if (form.autoReply == false) {
        params['auto'] = 0
      }
      params['autoword'] = form.autoword

      var isIdparams = {}
      isIdparams['id'] = route.query.id
      isIdparams['advertiseType'] = form.advertiseType
      isIdparams['price'] = (price.value + '').replace(/[^\d|.]/g, '') - 0
      isIdparams['coin.id'] = form.coin
      isIdparams['minLimit'] = form.minLimit
      isIdparams['maxLimit'] = form.maxLimit
      isIdparams['timeLimit'] = form.timeLimit
      isIdparams['country'] = form.country
      if (form.fixed == true) {
        isIdparams['priceType'] = 0
      } else if (form.fixed == false) {
        isIdparams['priceType'] = 1
      }
      isIdparams['premiseRate'] = form.premisePrice
      isIdparams['remark'] = form.remark
      isIdparams['number'] = form.number
      isIdparams['pay'] = form.payMode
      isIdparams['jyPassword'] = form.priceW
      if (form.autoReply == true) {
        isIdparams['auto'] = 1
      } else if (form.autoReply == false) {
        isIdparams['auto'] = 0
      }
      isIdparams['autoword'] = form.autoword

      if (isId.value) {
        axios.post(`${host}/otc/advertise/update`, isIdparams, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
            'x-auth-token': localStorage.getItem('TOKEN')
          }
        }).then(response => {
          const resp = response.data
          if (resp.code == 0) {
            ElMessage.success(resp.message)
            setTimeout(() => {
              router.push('/uc/ad')
            }, 3000)
          } else {
            ElMessage.error(resp.message)
            disAllowBtn.value = false
          }
        }).catch(() => {
          ElMessage.error('请求失败')
          disAllowBtn.value = false
        })
      } else {
        axios.post(`${host}/otc/advertise/create`, params, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
            'x-auth-token': localStorage.getItem('TOKEN')
          }
        }).then(response => {
          const resp = response.data
          if (resp.code == 0) {
            ElMessage.success(resp.message)
            setTimeout(() => {
              router.push('/uc/ad')
            }, 3000)
          } else {
            ElMessage.error(resp.message)
            disAllowBtn.value = false
          }
        }).catch(() => {
          ElMessage.error('请求失败')
          disAllowBtn.value = false
        })
      }
    } else {
      disAllowBtn.value = false
      ElMessage.error('提交失败，请检查表单填写')
    }
  })
}

const onAreaChange = (value) => {
  for (var i = 0; i < areas.value.length; i++) {
    if (areas.value[i].zhName == value) {
      form.rmb = areas.value[i].localCurrency
    }
  }
}

const findCoin = (coinId) => {
  for (let i = 0; i < coinList.value.length; i++) {
    if (coinList.value[i].id == coinId) {
      return coinList.value[i].unit
    }
  }
}

const getCoin = (coinId) => {
  for (let i = 0; i < coinList.value.length; i++) {
    if (coinList.value[i].id == coinId) {
      return coinList.value[i]
    }
  }
}

const getAreas = () => {
  axios.post(`${host}/uc/support/country`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    areas.value = resp.data
    form.country = areas.value[0].zhName
    form.rmb = areas.value[0].localCurrency
  }).catch(() => {
    ElMessage.error('获取国家列表失败')
  })
}

const getMember = () => {
  axios.get(`${host}/uc/identification`, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    const certifiedBusinessStatus = resp.data.certifiedBusinessStatus
    if (certifiedBusinessStatus == 2) {
      getAccount()
    } else {
      ElMessage.warning('请先申请商家认证!')
      router.push('/identbusiness')
    }
  }).catch(() => {
    ElMessage.error('获取用户信息失败')
  })
}

const getAccount = () => {
  axios.post(`${host}/uc/approve/account/setting`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      if (resp.data.bankVerified == 0 && resp.data.aliVerified == 0 && resp.data.wechatVerified == 0) {
        ElMessage.success('请先绑定至少一种支付方式!')
        router.push('/uc/account')
      }
      if (resp.data.aliVerified == 1) {
        payModeList.value[0].isOpen = false
      }
      if (resp.data.wechatVerified == 1) {
        payModeList.value[1].isOpen = false
      }
      if (resp.data.bankVerified == 1) {
        payModeList.value[2].isOpen = false
      }
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('获取账户信息失败')
  })
}

const getDetailAd = () => {
  isId.value = true
  axios.post(`${host}/otc/advertise/detail`, { id: route.query.id }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      form.coin = resp.data.coinId + ''
      form.country = resp.data.country.zhName
      cankao.value = resp.data.marketPrice + ''
      if (resp.data.priceType == 0) {
        form.fixed = true
        form.fixedPrice = resp.data.price
      } else if (resp.data.priceType == 1) {
        form.fixed = false
      }
      price.value = resp.data.price
      symbol.value = resp.data.coinUnit
      form.advertiseType = resp.data.advertiseType + ''
      form.minLimit = resp.data.minLimit
      form.maxLimit = resp.data.maxLimit
      form.remark = resp.data.remark
      form.timeLimit = resp.data.timeLimit
      form.premisePrice = resp.data.premiseRate
      form.payMode = (resp.data.payMode + '').split(',')
      form.number = resp.data.number
      if (resp.data.auto == 1) {
        form.autoReply = true
      } else if (resp.data.auto == 0) {
        form.autoReply = false
      }
      form.autoword = resp.data.autoword
    } else {
      ElMessage.error(resp.message)
    }
    isSpinShow.value = false
  }).catch(() => {
    ElMessage.error('获取广告详情失败')
    isSpinShow.value = false
  })
}

const getCoins = () => {
  axios.post(`${host}/otc/coin/all`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      coinList.value = resp.data
      form.coin = resp.data[0].id
      cankao.value = resp.data[0].marketPrice + ''
      let cankoNew = cankao.value.replace(/,/g, '').replace(/[^\d|.]/g, '') - 0
      let lv = (1 + form.premisePrice / 100).toFixed(4)
      price.value = (cankoNew * lv).toLocaleString()
      symbol.value = resp.data[0].unit
      if (route.query.id) {
        getDetailAd()
      } else {
        isSpinShow.value = false
      }
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('获取币种列表失败')
    isSpinShow.value = false
  })
}

// Watchers
const wantHistory = computed(() => form.advertiseType)
const premisePriceHistory = computed(() => form.premisePrice)
const fixedPriceHistory = computed(() => form.fixedPrice)

watch(wantHistory, (newValue, oldValue) => {
  if (newValue == '1') {
    wantstyle.value = '卖出'
    wantTime.value = '15-120'
  } else {
    wantstyle.value = '买入'
    wantTime.value = '10-30'
  }
})

watch(premisePriceHistory, (newValue, oldValue) => {
  let lv = (1 + newValue / 100).toFixed(4)
  let cankoNew = cankao.value.replace(/,/g, '').replace(/[^\d|.]/g, '') - 0
  price.value = round(mul(cankoNew, lv), 2).toLocaleString()
  gongshi.value = 1 + newValue / 100
})

watch(fixedPriceHistory, (newValue, oldValue) => {
  price.value = (newValue - 0).toLocaleString()
})

// Initialize
onMounted(() => {
  wantstyle.value = '卖出'
  getAreas()
  getCoins()
})
</script>

<style scoped lang="scss">
.my_ad_container {
  width: 80%;
  float: right;
}

.cankao {
  color: #e24a64;
}

.contbox {
  position: relative;
}

#price {
  font-size: 18px;
  color: #e24a64;
}

.send-box .send-form .msg {
  padding-left: 90px;
  margin-bottom: 10px;
  position: relative;
  top: -4px;
}

.formbox {
  width: 50%;
  padding-top: 30px;
}

.send-box {
  color: #fff;
  padding: 32px;
}

.title-box {
  border-bottom: 1px dashed #ccc;
  padding-bottom: 30px;
  text-align: left;
  padding-left: 18px;
}

.title-box .titles {
  font-size: 18px;
  font-weight: normal;
  color: #fff;
  margin-bottom: 15px;
}

.title-box p {
  line-height: 2;
}

.title-box p a {
  color: #f0a70a;
}

.order-table {
  margin-top: 20px;
}

.content-wrap {
  min-height: 750px;
}

.container {
  margin: 0 auto;
}

.my_ad_container :deep(.el-loading-spinner) {
  top: 200px;
}
</style>

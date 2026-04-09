<template>
  <div class="activity">
    <img style="banner" src="../../assets/images/crowdfunding_banner.png" alt="">
    <div class="banner_text">
      <p>{{$t('crowdfunding.totalAmount')}}</p>
      <p>
        <span class="money">{{totalMoney}}</span>
        <br class="br">
        <span class="money_span">{{$t('crowdfunding.totalAmount2')}}</span>
      </p>
      <div class="banner_content">
        <p>{{totalTimes}}<br class="br"><span>{{$t('crowdfunding.donations')}}</span></p>
        <p>{{totalProject}}<br class="br"><span>{{$t('crowdfunding.projectsGetHelp')}}</span></p>
      </div>
    </div>
    <div class="tab">
      <el-tabs v-model="activeTab" class="tabs1">
        <el-tab-pane :label="$t('crowdfunding.crowdfunding')" name="name1">
          <div class="tab_all" @click="syadd_show()">{{$t('crowdfunding.seeMore')}} ></div>
          <div class="swiper-container">
            <div class="swiper-wrapper">
              <div class="swiper-slide swiper-slide1" v-for="(item, index) in list.slice(0, 4)" :key="index">
                <div @click="xs_details_show(item)">
                  <img :src="item.picUrl" alt=""/>
                  <div class="swiper_text">
                    <div class="swiper_title">
                      <p>{{item.fundingTitle}}</p>
                      <p>{{dateFormat(item.addTime)}}—{{dateFormat(item.endTime)}}</p>
                    </div>
                    <el-row class="swiper_content">
                        <el-col :span="8">
                          <p>{{item.targetAmount}}</p>
                          <p>{{$t('crowdfunding.targetAmount')}}(USDT)</p>
                        </el-col>
                        <el-col :span="8">
                          <p>{{item.amountReceived}}</p>
                          <p>{{$t('crowdfunding.amountRaised')}}(USDT)</p>
                        </el-col>
                        <el-col :span="8">
                          <p>{{item.times}}</p>
                          <p>{{$t('crowdfunding.helpTimes')}}({{$t('crowdfunding.times')}})</p>
                        </el-col>
                    </el-row>
                  </div>
                </div>
              </div>
            </div>
            <div class="swiper-pagination"></div>
            <div class="swiper-button-prev"></div>
            <div class="swiper-button-next"></div>
          </div>
          <div class="tab_all2" @click="syadd_show()">{{$t('crowdfunding.seeMore')}}</div>
          <div class="btn" v-if="ifxsBtn != ''" @click="maskClick()">
              <span>{{$t('crowdfunding.initiateCrowdfunding')}}</span>
          </div>
        </el-tab-pane>
        <el-tab-pane :label="$t('crowdfunding.publicWelfare')" name="name2">
          <div class="tab_all" @click="syadd_show2()">{{$t('crowdfunding.seeMore')}} ></div>
          <div class="swiper-container">
            <div class="swiper-wrapper">
              <div class="swiper-slide swiper-slide2" v-for="(item, index) in enablewelfare.slice(0, 4)" :key="index">
                <div @click="xx_detailsClick(item)">
                  <img :src="item.picUrl" alt=""/>
                  <div class="swiper_text">
                    <div class="swiper_title">
                      <p>{{item.fundingTitle}}</p>
                      <p>{{dateFormat(item.passTime)}}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="swiper-pagination"></div>
            <div class="swiper-button-prev"></div>
            <div class="swiper-button-next"></div>
          </div>
          <div class="tab_all2" @click="syadd_show2()">{{$t('crowdfunding.seeMore')}}</div>
          <div class="btn" v-if="ifxxBtn != ''" @click="xx_apply()">
              <span>{{$t('crowdfunding.applySubsidy')}}</span>
          </div>
        </el-tab-pane>
      </el-tabs>
      <div class="bottom">
        <img src="../../assets/images/crowdfunding_text.png" alt=""/>
      </div>
      <div class="bottom_text">
        <p>{{$t('crowdfunding.loveReminder')}}</p>
        <p>{{$t('crowdfunding.loveReminder2')}}</p>
      </div>
    </div>
    <!--线上全部-->
    <div class="Initiate_mask4" v-if="all_show">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><el-icon v-if="screenWidth > 1200" size="28"><Document /></el-icon>
          <el-icon v-if="screenWidth < 1200" size="22"><Document /></el-icon>{{$t('crowdfunding.allCrowdfunding')}}</p>
          <img @click="allClick()" src="../../assets/images/down.png" alt="" />
        </div>
        <div class="window_tk_content" style="max-height: 570px; overflow-y: scroll;">
          <div class="all_xs"  @click="xs_details_show2(item)" v-for="(item, index) in list" :key="index">
            <div class="all_xs_img">
              <img :src="item.picUrl" alt=""/>
            </div>
            <div class="all_xs_text">
              <div class="all_xs_text_left">
                <p class="all_xs_text_title">{{item.fundingTitle}}</p>
                <div class="all_xs_text_money">
                  <div class="all_xs_text_content">
                        <div>
                          <p>{{$t('crowdfunding.amountRaised')}}<span>(USDT)</span></p>
                          <p>{{item.amountReceived}}</p>
                        </div>
                        <div>
                          <p class="all_xs_text_money">======</p>
                        </div>
                        <div>
                          <p>{{$t('crowdfunding.targetAmount')}}<span>(USDT)</span></p>
                          <p>{{item.targetAmount}}</p>
                        </div>
                    </div>
                </div>
              </div>
              <div class="all_xs_text_right">
                <div class="all_xs_right_top">
                  <span>{{item.times}}</span>
                  <span>{{$t('crowdfunding.helpTimes')}}</span>
                </div>
                <div class="all_xs_right_bottom">
                  {{dateFormat(item.passTime)}}—{{dateFormat(item.endTime)}}
                </div>
              </div>
              <div style="clear:both;"></div>
            </div>
            <div style="clear:both;"></div>
          </div>
        </div>
      </div>
    </div>
    <!--线下全部-->
    <div class="Initiate_mask4" v-if="all_show2">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><el-icon v-if="screenWidth > 1200" size="28"><Document /></el-icon>
          <el-icon v-if="screenWidth < 1200" size="22"><Document /></el-icon>{{$t('crowdfunding.allPublicWelfare')}}</p>
          <img @click="allClick2()" src="../../assets/images/down.png" alt="" />
        </div>
        <div class="window_tk_content" style="max-height: 570px; overflow-y: scroll;" >
          <div class="all_xs" @click="xx_detailsClick2(item)" v-for="(item, index) in enablewelfare" :key="index">
            <div class="all_xs_img">
              <img :src="item.picUrl" alt=""/>
            </div>
            <div class="all_xs_text">
              <div class="all_xs_text_left">
                <p class="all_xs_text_title">{{item.fundingTitle}}</p>
                <div class="all_xs_text_money">
                  <div class="all_xs_text_content">
                        <div>
                          <p>{{$t('crowdfunding.auditAmount')}}<span>(USDT)</span></p>
                          <p>{{item.amountReceived}}</p>
                        </div>
                        <div>
                          <p class="all_xs_text_money">======</p>
                        </div>
                        <div>
                          <p>{{$t('crowdfunding.applicationAmount')}}<span>(USDT)</span></p>
                          <p>{{item.applyAmount}}</p>
                        </div>
                    </div>
                </div>
              </div>
              <div class="all_xs_text_right">
                <div class="all_xs_right_top">
                  {{dateFormat(item.passTime)}}
                </div>
              </div>
              <div style="clear:both;"></div>
            </div>
            <div style="clear:both;"></div>
          </div>
        </div>
      </div>
    </div>
    <div class="mask" v-if="mask"></div>
    <!--选择捐款类型-->
    <div class="Initiate_mask" v-if="Initiate_select">
      <div class="Initiate_crowdfunding">
        <div class="select_title">
          <p>{{$t('crowdfunding.fundraisingType')}}</p>
          <img @click="maskClick2()" src="../../assets/images/down.png" alt="" />
        </div>
        <el-row :gutter="10">
            <el-col :xs="{ span: 8 }" :lg="{ span: 6}">
              <div class="window" @click="selectClick1()">
                <img src="../../assets/images/crowdfunding_tk1.png" alt="" />
                <p>{{$t('crowdfunding.wish')}}</p>
              </div>
            </el-col>
            <el-col :xs="{ span: 8 }" :lg="{ span: 6, offset: 3 }">
              <div class="window" @click="selectClick2()">
                <img class="window_img_yl" src="../../assets/images/crowdfunding_tk2.png" alt="" />
                <p>{{$t('crowdfunding.medical')}}</p>
              </div>
            </el-col>
            <el-col :xs="{ span: 8 }" :lg="{ span: 6, offset: 3 }">
              <div class="window" @click="selectClick3()">
                <img src="../../assets/images/crowdfunding_tk3.png" alt="" />
                <p>{{$t('crowdfunding.pioneer')}}</p>
              </div>
            </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 众筹首页
 */
import { ref, reactive, computed, onMounted, nextTick } from 'vue'
import { ElMessage, ElTabs, ElTabPane, ElRow, ElCol, ElIcon } from 'element-plus'
import { Document } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import Swiper from 'swiper'
import moment from 'moment'

const router = useRouter()
const store = useStore()

const host = ''

const activeTab = ref('name1')
const uploadHeaders = computed(() => ({ 'x-auth-token': localStorage.getItem('TOKEN') }))
const uploadUrl = `${host}/uc/upload/oss/image`

const money = ref('')
const totalMoney = ref('')
const totalProject = ref('')
const totalTimes = ref('')
const isLogin = ref(false)
const mask = ref(false)
const all = ref(false)
const all_show = ref(false)
const all_show2 = ref(false)
const Initiate_select = ref(false)
const Initiate_xy = ref(false)
const Initiate_yl = ref(false)
const Initiate_yl2 = ref(false)
const Initiate_cy = ref(false)
const Initiate_xx = ref(false)
const xx_details_show = ref(false)
const xinashang_details_show = ref(false)
const JK_details_show = ref(false)
const PL_details_show = ref(false)
const screenWidth = ref(document.body.clientWidth)
const level = ref('')
const item = ref({})
const xx_level = ref('')
const code = ref('')
const ifxsBtn = ref('')
const ifxxBtn = ref('')
const zje = ref('')
const list = ref([])
const xs_details = ref({})
const xs_details_img = ref({})
const xx_details = ref({})
const recordList = ref([])
const commentList = ref([])
const enablewelfare = ref([])
const getWelfareDetail = ref([])

const ylDetails = reactive({
  targetAmount: '',
  fundingTitle: '',
  fundingUse: '',
  plan: '',
  story: '',
  picUrl: '',
  promise: '',
  patientName: '',
  patientDocumentType: '1',
  patientDocumentNumber: '',
  visitInformation: '',
  crowdPics: { pics: [] },
  sponsor: '',
  documentType: '1',
  documentNumber: '',
  houseProperty: '无房产',
  carProperty: '无车产',
  income: '',
  property: '',
  debt: '',
  otherFundType: '1'
})

const xyDetails = reactive({
  targetAmount: '',
  fundingTitle: '',
  instructions: '',
  picUrl: '',
  sponsor: '',
  documentType: '1',
  documentNumber: '',
  crowdPics: { pics: [] }
})

const xxDetails = reactive({
  applyAmount: '',
  fundingTitle: '',
  instructions: '',
  picUrl: '',
  sponsor: '',
  documentType: '1',
  documentNumber: '',
  material: '',
  crowdPics: { pics: [] }
})

const fundingMoney = ref('')
const jkMoney = ref('')
const common = ref('')

const lang = computed(() => store.state.lang)
const langPram = computed(() => {
  if (store.state.lang === '简体中文') return 'CN'
  if (store.state.lang === 'English') return 'EN'
  return 'CN'
})

const dateFormat = (tick) => moment(tick).format('YYYY-MM-DD')

const initSwiper = () => {
  nextTick(() => {
    const slides = document.querySelectorAll('.swiper-slide')
    if (slides.length === 0) {
      return
    }

    if (screenWidth.value > 1200) {
      new Swiper('.swiper-container', {
        autoplay: { delay: 3000, disableOnInteraction: false },
        loop: true,
        speed: 1000,
        centeredSlides: true,
        paginationClickable: true,
        slidesPerView: 3,
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev'
        },
        onInit: (swiper) => {
          if (swiper.slides && swiper.slides[3]) {
            swiper.slides[3].className = 'swiper-slide swiper-slide-active'
          }
        },
        breakpoints: { 668: { slidesPerView: 3 } }
      })
    } else {
      new Swiper('.swiper-container', {
        autoplay: { delay: 3000, disableOnInteraction: false },
        loop: true,
        speed: 1000,
        centeredSlides: true,
        paginationClickable: true,
        slidesPerView: 1,
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev'
        },
        onInit: (swiper) => {
          if (swiper.slides && swiper.slides[1]) {
            swiper.slides[1].className = 'swiper-slide swiper-slide-active'
          }
        },
        breakpoints: { 668: { slidesPerView: 1 } }
      })
    }
  })
}

const fmoney = () => {
  const n = 2
  let s = parseFloat((totalMoney.value + '').replace(/[^\d\.-]/g, '')).toFixed(n) + ''
  const l = s.split('.')[0].split('').reverse()
  const r = s.split('.')[1]
  let t = ''
  for (let i = 0; i < l.length; i++) {
    t += l[i] + ((i + 1) % 3 === 0 && (i + 1) !== l.length ? ',' : '')
  }
  totalMoney.value = t.split('').reverse().join('') + '.' + r
}

const fmoney2 = () => {
  const n = 2
  let s = parseFloat((totalProject.value + '').replace(/[^\d\.-]/g, '')).toFixed(n) + ''
  const l = s.split('.')[0].split('').reverse()
  const r = s.split('.')[1]
  let t = ''
  for (let i = 0; i < l.length; i++) {
    t += l[i] + ((i + 1) % 3 === 0 && (i + 1) !== l.length ? ',' : '')
  }
  totalProject.value = t.split('').reverse().join('')
}

const fmoney3 = () => {
  const n = 2
  let s = parseFloat((totalTimes.value + '').replace(/[^\d\.-]/g, '')).toFixed(n) + ''
  const l = s.split('.')[0].split('').reverse()
  const r = s.split('.')[1]
  let t = ''
  for (let i = 0; i < l.length; i++) {
    t += l[i] + ((i + 1) % 3 === 0 && (i + 1) !== l.length ? ',' : '')
  }
  totalTimes.value = t.split('').reverse().join('')
}

const info = () => {
  if (isLogin.value) {
    getLevelRight()
  }

  axios.get(`${host}/uc/crowdfunding/totalAmount`, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      totalMoney.value = res.data.data.totalMoney
      totalProject.value = res.data.data.totalProject
      totalTimes.value = res.data.data.totalTimes
      fmoney()
      fmoney2()
      fmoney3()
    }
  }).catch(() => {})

  axios.get(`${host}/uc/crowdfunding/crowdfundingList`, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      res.data.data.funding.forEach(v => list.value.push(v))
      res.data.data.medicalfunding.forEach(v => list.value.push(v))
      initSwiper()
    }
  }).catch(() => {})

  axios.get(`${host}/uc/crowdfunding/offlineList`, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      enablewelfare.value = res.data.data
    }
  }).catch(() => {})
}

const getLevelRight = () => {
  axios.post(`${host}/uc/crowdfunding/level`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      ifxsBtn.value = res.data.data.crowdRight
      ifxxBtn.value = res.data.data.welfareRight
      level.value = res.data.data.crowdRight * 10000
      xx_level.value = res.data.data.welfareRight * 10000
    }
  }).catch(() => {})
}

const syadd_show = () => { all_show.value = true }
const allClick = () => { all_show.value = false }
const syadd_show2 = () => { all_show2.value = true }
const allClick2 = () => { all_show2.value = false }

const maskClick = () => {
  mask.value = true
  Initiate_select.value = true
}

const maskClick2 = () => {
  mask.value = false
  Initiate_select.value = false
}

const selectClick1 = () => {
  Initiate_select.value = false
  Initiate_xy.value = true
}

const selectClick2 = () => {
  Initiate_select.value = false
  Initiate_yl.value = true
}

const selectClick3 = () => {
  Initiate_select.value = false
  Initiate_cy.value = true
}

const xs_details_show = (itemData) => {
  item.value = itemData
  mask.value = true
  xinashang_details_show.value = true
  xs_details.value = itemData
  donateList(itemData)
}

const xs_details_show2 = (itemData) => {
  item.value = itemData
  mask.value = true
  xinashang_details_show.value = true
  all_show.value = false
  all.value = true
  xs_details.value = itemData
  donateList(itemData)
}

const xx_detailsClick = (itemData) => {
  mask.value = true
  xx_details_show.value = true
  xx_details.value = itemData

  axios.post(`${host}/uc/crowdfunding/getWelfareDetail`, { id: itemData.id }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      getWelfareDetail.value = res.data.data.picsUrl
    }
  }).catch(() => {})
}

const xx_detailsClick2 = (itemData) => {
  mask.value = true
  xx_details_show.value = true
  all_show2.value = false
  all.value = true
  xx_details.value = itemData
}

const donateList = (itemData) => {
  axios.post(`${host}/uc/crowdfunding/donateList`, { id: itemData.id, type: itemData.type }, {
    emulateJSON: false,
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      recordList.value = res.data.data.recordList
      commentList.value = res.data.data.fundingCommonList
      xs_details.value = res.data.data.funding
      xs_details_img.value = res.data.data.picsUrl
    }
  }).catch(() => {})
}

const init = () => {
  store.commit('navigate', 'nav-crowdfunding')
}

onMounted(() => {
  init()
  isLogin.value = store.getters.isLogin
  info()
})
</script>

<style lang="scss" scoped>
.all_xs {
  width: 90%;
  margin: 0 auto;
  height: auto;
  padding-top: 30px;
  border-bottom: 1px solid #ccc;
}
.all_xs_img {
  float: left;
}
.all_xs_img img {
  width: 300px;
  height: 150px;
  border-radius: 10px;
}
.all_xs_text {
  float: right;
  width: calc(100% - 300px);
  color: #000;
  border-radius: 10px;
}
.all_xs_text_left {
  float: left;
  margin-left: 20px;
}
.all_xs_text_title {
  width: 300px;
  font-size: 22px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.all_xs_text_content {
  background: #f5f5f5;
  border-radius: 10px;
  padding: 0px 10px;
  display: flex;
}
.all_xs_text_content > div {
  width: auto;
  text-align: center;
}
.all_xs_text_content p {
  line-height: 30px;
  min-width: 100px;
  padding: 5px 10px;
  text-align: center;
}
.all_xs_text_content p:nth-child(2) {
  color: red;
  font-size: 20px;
  font-weight: bold;
}
.all_xs_text_content p:nth-child(1) {
  color: #000;
  font-size: 16px;
  white-space: nowrap;
}
.all_xs_text_money {
  position: relative;
  top: 25px;
  color: #FE6700;
}
.all_xs_text_right {
  float: right;
  position: relative;
}
.all_xs_right_top {
  font-size: 16px;
  height: 40px;
  line-height: 40px;
}
.all_xs_right_top span:nth-child(1) {
  color: red;
  font-size: 20px;
}
.all_xs_right_top span:nth-child(2) {
  color: #666;
  font-size: 14px;
  position: relative;
  top: 0px;
}
.all_xs_right_bottom {
  margin-top: 80px;
  font-size: 14px;
}

.activity {
  width: 100%;
  background: rgba(242, 246, 250, 1) !important;
  height: 100%;
  background-size: cover;
  position: relative;
  overflow: hidden;
  padding-bottom: 20px;
  padding-top: 60px;
  color: #fff;
}
.activity .banner {
  width: 100%;
}
.banner_text {
  position: absolute;
  top: 60px;
  width: 100%;
  text-align: center;
  color: #fff;
  font-weight: 500;
  font-size: 28px;
  margin-top: 40px;
}
.money {
  font-size: 64px;
  color: red;
  letter-spacing: 14px;
}
.money_span {
  font-size: 33px;
  color: #fff;
  font-weight: 400;
  position: relative;
  top: -10px;
}
.banner_content {
  width: 1200px;
  margin: 0 auto;
}
.banner_content p {
  font-size: 20px;
  letter-spacing: 4px;
}
.banner_content p span {
  font-size: 18px;
  letter-spacing: 1px;
  position: relative;
  top: -1px;
}
.banner_content p:nth-child(1) {
  float: left;
  font-weight: 100;
}
.banner_content p:nth-child(2) {
  float: right;
  font-weight: 100;
}
.br {
  display: none;
}
.tab {
  width: 1200px;
  margin: 0 auto;
  font-weight: 500;
}
.tab_all {
  position: relative;
  top: -42px;
  float: right;
  color: #A1A1A1;
  cursor: pointer;
}
.tab_all2 {
  width: 100%;
  text-align: center;
  font-size: 18px;
  margin-top: 10px;
  display: none;
}
.btn {
  margin: 0 auto;
  margin-top: 10px;
  width: 240px;
  height: 60px;
  line-height: 60px;
  font-size: 23px;
  color: #fff;
  text-align: center;
  border-radius: 50px;
  background-image: linear-gradient(to bottom, #FF9113, #F16219);
}
.btn span {
  cursor: pointer;
}
.bottom {
  width: 1200px;
  text-align: center;
  margin-top: 42px;
}
.bottom_text {
  width: 1200px;
  text-align: left;
  color: #2987FE;
}
.mask {
  position: absolute;
  left: 0px;
  top: 0px;
  background: rgba(0, 0, 0, 0.5);
  width: 100%;
  height: 100%;
  z-index: 1;
}
.Initiate_mask {
  position: fixed;
  top: 20%;
  width: 100%;
  z-index: 2;
}
.Initiate_mask2 {
  position: relative;
  margin-top: -45%;
  width: 100%;
  z-index: 2;
}
.Initiate_mask3 {
  position: fixed;
  top: 20%;
  width: 100%;
  z-index: 2;
}
.Initiate_mask4 {
  position: fixed;
  top: 100px;
  width: 100%;
  z-index: 2;
}
.Initiate_crowdfunding {
  margin: 0 auto;
  width: 950px;
  height: 500px;
}
.select_title {
  width: 100%;
  height: 50px;
  display: flex;
  margin-bottom: 90px;
}
.select_title p {
  width: calc(100% - 26px);
  text-align: center;
  font-size: 32px;
}
.select_title img {
  width: 26px;
  height: 26px;
  position: relative;
  top: 12px;
  cursor: pointer;
}
.window {
  height: 320px;
}
.window_img_yl {
  margin-top: 70px;
}
.window img {
  margin-top: 90px;
}
.window p {
  color: #0071FE;
  cursor: pointer;
}
.window p:nth-child(2) {
  font-size: 24px;
  font-weight: bold;
  margin-top: 10px;
}
.window p:nth-child(3) {
  font-size: 16px;
}
.window_tk {
  position: relative;
  margin: 0 auto;
  width: 1110px;
  height: auto;
  background: #fff;
  border-radius: 20px;
  z-index: 2;
}
.window_tk2 {
  position: relative;
  margin: 0 auto;
  width: 700px;
  height: auto;
  background: #fff;
  border-radius: 20px;
  box-shadow: 0px 0px 4px #ccc;
  z-index: 2;
}
.window_tk_title {
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
  width: 100%;
  height: 173px;
  display: flex;
  background-image: linear-gradient(to bottom, #2987FE, #FFFFFF);
}
.window_tk_title p {
  width: calc(100% - 86px);
  text-align: left;
  font-size: 22px;
  margin-left: 30px;
  line-height: 50px;
}
.window_tk_title img {
  width: 26px;
  height: 26px;
  position: relative;
  top: 12px;
  cursor: pointer;
}
.window_tk_content {
  position: relative;
  top: -120px;
  min-height: 300px;
  border-radius: 20px;
  width: 1056px;
  margin: 0 auto;
  background: #fff;
}
.window_tk_content2 {
  position: relative;
  top: -120px;
  min-height: 200px;
  border-radius: 20px;
  width: 656px;
  margin: 0 auto;
  background: #fff;
}
.window_btn {
  cursor: pointer;
  position: relative;
  top: -40px;
  margin: 0 auto;
  width: 240px;
  height: 60px;
  line-height: 60px;
  font-size: 23px;
  color: #fff;
  text-align: center;
  border-radius: 50px;
  background-image: linear-gradient(to bottom, #0071FE, #0071FE);
}
.window_btn2 {
  cursor: pointer;
  position: relative;
  top: -40px;
  margin: 0 auto;
  width: 240px;
  height: 60px;
  line-height: 60px;
  font-size: 23px;
  color: #fff;
  text-align: center;
  border-radius: 50px;
  background-image: linear-gradient(to bottom, #FF9113, #F16219);
}
.window_jk_title {
  position: relative;
  top: -40px;
  margin: 0 auto;
  width: 240px;
  height: 60px;
  line-height: 60px;
  font-size: 16px;
  color: red;
  text-align: center;
  border-radius: 50px;
}
.window_tk_nr {
  width: 92%;
  margin: 0 auto;
  padding-top: 20px;
}
.window_tk_nr_width {
  width: 300px;
}
.window_tk_nr_p {
  color: #000;
  padding-bottom: 10px;
  position: relative;
  left: -20px;
}
.img_text {
  position: relative;
  color: #aaa;
  left: 140px;
  top: -30px;
  font-size: 12px;
}
.details {
  margin-left: 10px;
}
.details > p {
  color: #000;
  font-size: 14px;
  padding: 10px;
  position: relative;
  left: -20px;
}
.details > p span {
  color: #0071FE;
}
.window_xx_name {
  width: 95%;
  margin: 0 auto;
  padding-top: 20px;
  color: #000;
  display: flex;
}
.window_tx {
  width: 46px;
  height: 46px;
  border-radius: 25px;
}
.window_xx_name p {
  height: 46px;
  line-height: 46px;
  font-size: 22px;
  margin-left: 20px;
}
.window_xx_name p:nth-child(3) {
  font-size: 18px !important;
  color: #A1A1A1;
}
.window_xx_title {
  color: #000;
  width: 95%;
  margin: 0 auto;
  margin-top: 23px;
  height: 30px;
  line-height: 30px;
  display: flex;
}
.window_xx_title p:nth-child(1) {
  width: calc(100% - 220px);
  font-size: 28px;
  font-weight: bold;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.window_xx_title p:nth-child(2) {
  width: 220px;
  font-size: 18px;
  color: #a1a1a1;
}
.window_xx_content {
  width: 95%;
  margin: 0 auto;
  margin-top: 50px;
}
.window_xx_content p {
  text-indent: 35px;
  line-height: 30px;
  font-size: 18px;
  color: #373737;
}
.window_xx_img img {
  margin-top: 20px;
  width: 200px;
  height: 200px;
}
.window_xx_name2 {
  width: 95%;
  margin: 0 auto;
  padding-top: 20px;
  color: #000;
}
.window_tx_left {
  float: left;
  display: flex;
}
.window_tx_left p {
  height: 46px;
  line-height: 46px;
  font-size: 22px;
  margin-left: 20px;
}
.window_tx_left p:nth-child(3) {
  font-size: 18px !important;
  color: #A1A1A1;
}
.window_tx_right {
  width: 165px;
  height: 44px;
  line-height: 44px;
  text-align: center;
  background: #F65F47;
  border-radius: 50px;
  color: #fff;
  font-size: 22px;
  float: right;
}
.window_xs_content {
  width: 95%;
  margin: 0 auto;
  margin-top: 20px;
  background: #F8F8F8;
  border-radius: 10px;
  height: auto;
}
.window_col {
  text-align: center;
  color: #000;
  margin-top: 10px;
}
.window_col p:nth-child(1) {
  color: #FE6700;
  font-size: 32px;
}
.window_col p:nth-child(2) {
  color: #707070;
  font-size: 20px;
}
.window_col2 {
  text-align: center;
  height: 120px;
  line-height: 120px;
  color: #000;
  margin-top: 60px;
  background: #fff;
  box-shadow: 0px 0px 4px #ccc;
  border-radius: 15px;
  font-size: 20px;
  border: 2px solid #fff;
  cursor: pointer;
}
.window_col2:hover {
  border: 2px solid #0071FE;
}
.window_col3 {
  border: 2px solid #0071FE;
}
.window_xs_lable {
  width: 95%;
  margin: 0 auto;
  margin-top: 20px;
  background: #fff;
  border-radius: 10px;
  height: auto;
}
.window_xs_lable_title {
  width: 95%;
  margin: 0 auto;
  padding: 20px 0;
}
.window_xs_lable_title .left {
  float: left;
  color: #5E5E5E;
  font-size: 22px;
}
.window_xs_lable_title .right {
  float: right;
  color: #A1A1A1;
  font-size: 16px;
}
.window_xs_lable_content {
  width: 95%;
  margin: 0 auto;
  color: #000;
}
.window_xs_lable_content > div {
  width: 100%;
  display: flex;
  height: 50px;
  line-height: 50px;
  font-size: 20px;
}
.window_xs_lable_content > div > p {
  width: 33.33%;
  text-align: left;
}
.window_xs_lable_content > div > p:nth-child(2) {
  color: #F16219;
  text-align: center;
}
.window_xs_lable_content > div > p:nth-child(3) {
  text-align: right;
}
.window_xs_pl {
  width: 100%;
  display: flex;
}
.window_xs_pl p {
  width: 50%;
}
.window_xs_pl p:nth-child(1) {
  text-align: left;
  font-size: 20px;
  line-height: 30px;
}
.window_xs_pl_img {
  width: 30px;
  height: 30px;
  border-radius: 50px;
  position: relative;
  top: 8px;
}
.window_xs_pl p:nth-child(2) {
  text-align: right;
  font-size: 18px;
  position: relative;
  top: 8px;
}
.window_xs_pl_con {
  width: 100%;
  line-height: 30px;
  border-bottom: 1px solid #ccc;
  padding: 20px 0;
  text-indent: 20px;
}
.window_jkjg {
  width: 500px;
  margin: 0 auto;
}
.swiper-container {
  min-height: 300px;
}
.swiper-wrapper {
  margin: 20px 0px;
}
.swiper-slide {
  -webkit-transition: transform 1.0s;
  -moz-transition: transform 1.0s;
  -ms-transition: transform 1.0s;
  -o-transition: transform 1.0s;
  -webkit-transform: scale(0.7);
  transform: scale(0.7);
  cursor: pointer;
}
.swiper-slide-active,
.swiper-slide-duplicate-active {
  -webkit-transform: scale(1);
  transform: scale(1);
}
.swiper-slide a {
  background: #fff;
  padding: 10px;
  display: block;
  border-radius: 14px;
}
.swiper-slide > div {
  width: auto;
  box-shadow: 0px 0px 4px #ccc;
  border-radius: 14px;
  height: 300px;
}
.swiper-slide img {
  width: 100%;
  border-top-left-radius: 14px;
  border-top-right-radius: 14px;
  display: block;
  height: 194px;
}
.swiper-slide2 img {
  width: 100%;
  border-top-left-radius: 14px;
  border-top-right-radius: 14px;
  display: block;
  height: 254px;
}
.swiper-button-next,
.swiper-button-prev {
  outline: none;
  opacity: 0.6;
  width: 25px;
  height: 40px;
}
.swiper_title {
  width: 90%;
  margin: 0 auto;
  display: flex;
  margin-top: 5px;
}
.swiper_title p {
  width: 50%;
}
.swiper_title p:nth-child(1) {
  color: #373737;
  font-size: 20px;
  font-weight: 600;
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.swiper_title p:nth-child(2) {
  color: #373737;
  font-size: 12px;
  position: relative;
  text-align: right;
  top: 6px;
}
.swiper_content {
  margin-top: 10px;
}
.swiper_content > div p:nth-child(1) {
  color: #FE6700;
  font-size: 20px;
}
.swiper_content > div p:nth-child(2) {
  color: #707070;
  font-size: 13px;
}

/* tabs 样式 */
:deep(.tabs1 .el-tabs__nav) {
  border-bottom: none;
}
:deep(.tabs1 .el-tabs__item) {
  font-size: 26px;
}
:deep(.tabs2 .el-tabs__nav) {
  .el-tabs__item {
    font-size: 26px;
  }
}

/* 移动端适配 */
@media screen and (max-width: 1200px) {
  .swiper-container {
    min-height: 250px;
  }
  :deep(.tabs1 .el-tabs__item) {
    font-size: 22px;
    width: 50%;
    text-align: center;
    margin-left: 5px;
  }
  .activity {
    width: 100%;
    background: rgba(242, 246, 250, 1) !important;
    height: 100%;
    background-size: cover;
    position: relative;
    overflow: hidden;
    padding-bottom: 20px;
    padding-top: 45px;
    color: #fff;
  }
  .activity .banner {
    width: 100%;
    height: 400px;
  }
  .banner_text {
    position: absolute;
    top: 60px;
    width: 100%;
    text-align: center;
    color: #fff;
    font-weight: 500;
    font-size: 24px;
    margin-top: 0px;
  }
  .br {
    display: block;
  }
  .money {
    font-size: 36px;
    color: red;
    letter-spacing: 3px;
  }
  .money_span {
    font-size: 24px;
    color: #fff;
    font-weight: 400;
    position: relative;
    top: 0px;
  }
  .banner_content {
    width: 100%;
    margin: 0 auto;
    margin-top: 10px;
  }
  .banner_content p {
    font-size: 20px;
    letter-spacing: 1px;
    width: 50%;
    text-align: center;
  }
  .banner_content p span {
    font-size: 18px;
    letter-spacing: 1px;
    position: relative;
    top: -1px;
  }
  .banner_content p:nth-child(1) {
    float: left;
    font-weight: 600;
  }
  .banner_content p:nth-child(2) {
    float: right;
    font-weight: 600;
  }
  .tab {
    width: 100%;
    font-weight: 500;
  }
  .tab_all {
    display: none;
  }
  .swiper-wrapper {
    margin: 0px 0px;
  }
  .swiper-button-prev,
  .swiper-button-next {
    display: none;
  }
  .swiper-slide2 img {
    width: 100%;
    border-top-left-radius: 14px;
    border-top-right-radius: 14px;
    display: block;
    height: 224px;
  }
  .swiper-slide2 > div {
    width: auto;
    box-shadow: 0px 0px 4px #ccc;
    border-radius: 14px;
    height: 300px;
  }
  .swiper-slide2 .swiper_title {
    width: 95%;
    display: flex;
    margin-top: 25px;
  }
  .btn {
    margin: 0 auto;
    margin-top: 20px;
    width: 190px;
    height: 45px;
    line-height: 45px;
    font-size: 18px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #FF9113, #F16219);
  }
  .btn span {
    cursor: pointer;
  }
  .bottom {
    width: 100%;
    text-align: center;
    margin-top: 42px;
  }
  .bottom img {
    width: 95%;
    margin: 0 auto;
  }
  .bottom_text {
    width: 95%;
    margin: 0 auto;
    text-align: left;
    color: #2987FE;
  }
  .tab_all2 {
    width: 100%;
    text-align: center;
    font-size: 18px;
    margin-top: 10px;
    display: block;
  }
  .Initiate_mask4 {
    position: fixed;
    top: 40px;
    width: 100%;
    z-index: 2;
  }
  .all_xs_img {
    float: left;
    width: 100%;
  }
  .all_xs_img img {
    width: 100%;
    height: 160px;
    border-radius: 10px;
  }
  .all_xs_text {
    float: none;
    width: auto;
    color: #000;
    border-radius: 10px;
  }
  .all_xs_text_left {
    float: none;
    margin-left: 0px;
  }
  .all_xs_text_title {
    width: 300px;
    font-size: 20px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .all_xs_text_content {
    background: #f5f5f5;
    border-radius: 10px;
    padding: 0px 0px;
    display: flex;
  }
  .all_xs_text_content > div {
    width: 50%;
    text-align: center;
  }
  .all_xs_text_content > div:nth-child(2) {
    display: none;
  }
  .all_xs_text_content p {
    line-height: 30px;
    min-width: 100px;
    padding: 5px 10px;
    text-align: center;
  }
  .all_xs_text_content p:nth-child(2) {
    color: red;
    font-size: 20px;
    font-weight: bold;
  }
  .all_xs_text_content p:nth-child(1) {
    color: #000;
    font-size: 16px;
    white-space: nowrap;
  }
  .all_xs_text_money {
    position: relative;
    top: 5px;
    color: #FE6700;
  }
  .all_xs_text_right {
    float: none;
    position: relative;
  }
  .all_xs_right_top {
    font-size: 16px;
    height: 40px;
    line-height: 40px;
    float: left;
  }
  .all_xs_right_top span:nth-child(1) {
    color: red;
    font-size: 20px;
  }
  .all_xs_right_top span:nth-child(2) {
    color: #666;
    font-size: 14px;
    position: relative;
    top: 0px;
  }
  .all_xs_right_bottom {
    float: right;
    margin-top: 12px;
    font-size: 14px;
  }
  .Initiate_mask {
    position: fixed;
    top: 30%;
    width: 100%;
    z-index: 2;
  }
  .Initiate_crowdfunding {
    margin: 0 auto;
    width: 95%;
    height: 300px;
  }
  .select_title {
    width: 100%;
    height: 50px;
    display: flex;
    margin-bottom: 10px;
  }
  .select_title p {
    width: calc(100% - 26px);
    text-align: center;
    font-size: 22px;
  }
  .select_title img {
    width: 22px;
    height: 22px;
    position: relative;
    top: 6px;
    cursor: pointer;
  }
  .window {
    height: 120px;
  }
  .window img {
    margin-top: 10px;
    width: 30px;
    height: 30px;
  }
  .window_img_yl {
    margin-top: 10px;
  }
  .window p {
    color: #0071FE;
    cursor: pointer;
  }
  .window p:nth-child(2) {
    font-size: 14px;
    font-weight: bold;
    margin-top: 10px;
  }
  .window p:nth-child(3) {
    font-size: 12px;
  }
  .Initiate_mask2 {
    position: relative;
    margin-top: -840px;
    width: 100%;
    z-index: 2;
  }
  .window_tk {
    position: relative;
    margin: 0 auto;
    width: 100%;
    height: auto;
    background: #fff;
    border-radius: 0px;
    z-index: 2;
  }
  .window_tk_title {
    border-top-left-radius: 0px;
    border-top-right-radius: 0px;
    width: 100%;
    height: 173px;
    display: flex;
    background-image: linear-gradient(to bottom, #2987FE, #FFFFFF);
  }
  .window_tk_title p {
    width: calc(100% - 46px);
    text-align: left;
    font-size: 18px;
    margin-left: 10px;
    line-height: 50px;
  }
  .window_tk_title img {
    width: 18px;
    height: 18px;
    position: relative;
    top: 16px;
    cursor: pointer;
  }
  .window_tk_content {
    position: relative;
    top: -120px;
    min-height: 300px;
    border-radius: 0px;
    width: 100%;
    margin: 0 auto;
    background: #fff;
  }
  .window_btn {
    cursor: pointer;
    position: relative;
    top: -40px;
    margin: 0 auto;
    width: 140px;
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #0071FE, #0071FE);
  }
  .window_btn2 {
    cursor: pointer;
    position: relative;
    top: -40px;
    margin: 0 auto;
    width: 140px;
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #FF9113, #F16219);
  }
  .window_tk_nr {
    width: 92%;
    margin: 0 auto;
    padding-top: 20px;
  }
  .window_tk_nr_width {
    width: 300px;
  }
  .window_tk_nr_p {
    color: #000;
    padding-bottom: 10px;
    position: relative;
    left: -5px;
  }
  .img_text {
    position: relative;
    color: #aaa;
    left: 0px;
    top: 8px;
    font-size: 12px;
  }
  .details {
    margin-left: 10px;
  }
  .details > p {
    color: #000;
    font-size: 14px;
    padding: 10px;
    position: relative;
    left: -20px;
  }
  .details > p span {
    color: #0071FE;
  }
  .window_xx_name {
    width: 95%;
    margin: 0 auto;
    padding-top: 20px;
    color: #000;
    display: flex;
  }
  .window_tx {
    width: 40px;
    height: 40px;
    border-radius: 25px;
  }
  .window_xx_name p {
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    margin-left: 20px;
  }
  .window_xx_name p:nth-child(3) {
    font-size: 14px !important;
    color: #A1A1A1;
  }
  .window_xx_title {
    color: #000;
    width: 95%;
    margin: 0 auto;
    margin: 10px 15px;
    height: 45px;
    line-height: 30px;
    display: block;
  }
  .window_xx_title p:nth-child(1) {
    width: 100%;
    font-size: 16px;
    font-weight: bold;
  }
  .window_xx_title p:nth-child(2) {
    width: 200px;
    font-size: 14px;
    color: #a1a1a1;
  }
  .window_xx_content {
    width: 95%;
    margin: 0 auto;
    margin-top: 50px;
  }
  .window_xx_content p {
    text-indent: 35px;
    line-height: 30px;
    font-size: 14px;
    color: #373737;
  }
  .window_xx_img img {
    margin-top: 20px;
    width: 150px;
    height: 150px;
  }
  .window_xx_name2 {
    width: 95%;
    margin: 0 auto;
    padding-top: 20px;
    color: #000;
  }
  .window_tx_left {
    float: left;
    display: flex;
  }
  .window_tx_left p {
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    margin-left: 10px;
  }
  .window_tx_left p:nth-child(3) {
    font-size: 14px !important;
    color: #A1A1A1;
  }
  .window_tx_right {
    width: 105px;
    margin-top: 3px;
    height: 34px;
    line-height: 34px;
    text-align: center;
    background: #F65F47;
    border-radius: 50px;
    color: #fff;
    font-size: 14px;
    float: right;
  }
  .window_xs_content {
    width: 95%;
    margin: 0 auto;
    margin-top: 20px;
    background: #F8F8F8;
    border-radius: 10px;
    height: auto;
  }
  .window_col {
    text-align: center;
    color: #000;
    margin-top: 10px;
  }
  .window_col p:nth-child(1) {
    color: #FE6700;
    font-size: 16px;
  }
  .window_col p:nth-child(2) {
    color: #707070;
    font-size: 12px;
  }
  .window_xs_lable {
    width: 95%;
    margin: 0 auto;
    margin-top: 12px;
    background: #fff;
    border-radius: 10px;
    height: auto;
  }
  .window_xs_lable_title {
    width: 95%;
    margin: 0 auto;
    padding: 10px 0;
  }
  .window_xs_lable_title .left {
    float: left;
    color: #5E5E5E;
    font-size: 14px;
  }
  .window_xs_lable_title .right {
    float: right;
    color: #A1A1A1;
    font-size: 12px;
    margin-top: 2px;
  }
  .window_xs_lable_content {
    width: 95%;
    margin: 0 auto;
    color: #000;
  }
  .window_xs_lable_content > div {
    width: 100%;
    display: flex;
    height: 40px;
    line-height: 40px;
    font-size: 14px;
  }
  .window_xs_lable_content > div > p {
    width: 33.33%;
    text-align: left;
  }
  .window_xs_lable_content > div > p:nth-child(1) {
    width: 40%;
    text-align: left;
  }
  .window_xs_lable_content > div > p:nth-child(2) {
    width: 30%;
    color: #F16219;
    text-align: center;
  }
  .window_xs_lable_content > div > p:nth-child(3) {
    width: 30%;
    text-align: right;
  }
  .window_xs_pl {
    width: 100%;
    display: flex;
  }
  .window_xs_pl p {
    width: 50%;
  }
  .window_xs_pl p:nth-child(1) {
    text-align: left;
    font-size: 14px;
    line-height: 25px;
  }
  .window_xs_pl_img {
    width: 25px;
    height: 25px;
    border-radius: 50px;
    position: relative;
    top: 8px;
  }
  .window_xs_pl p:nth-child(2) {
    text-align: right;
    font-size: 14px;
    line-height: 25px;
  }
  .window_xs_pl_con {
    width: 100%;
    line-height: 30px;
    border-bottom: 1px solid #ccc;
    padding: 10px 0;
    text-indent: 20px;
  }
  .Initiate_mask3 {
    position: fixed;
    top: 15%;
    width: 100%;
    z-index: 2;
  }
  .window_tk2 {
    position: relative;
    margin: 0 auto;
    width: 95%;
    height: auto;
    background: #fff;
    border-radius: 0px;
    box-shadow: 0px 0px 4px #ccc;
    z-index: 2;
  }
  .window_tk_content2 {
    position: relative;
    top: -120px;
    min-height: 200px;
    border-radius: 20px;
    width: 90%;
    margin: 0 auto;
    background: #fff;
  }
  .window_col2 {
    text-align: center;
    height: 100px;
    line-height: 100px;
    color: #000;
    margin-top: 60px;
    background: #fff;
    box-shadow: 0px 0px 4px #ccc;
    border-radius: 15px;
    font-size: 14px;
    border: 2px solid #fff;
    cursor: pointer;
  }
  .window_col2:hover {
    border: 2px solid #0071FE;
  }
  .window_jkjg {
    width: 80%;
    margin: 0 auto;
  }
}

@media screen and (min-width: 1300px) {
  .Initiate_mask2 {
    margin-top: -70%;
  }
}
@media screen and (min-width: 1400px) {
  .Initiate_mask2 {
    margin-top: -65%;
  }
}
@media screen and (min-width: 1500px) {
  .Initiate_mask2 {
    margin-top: -60%;
  }
}
@media screen and (min-width: 1600px) {
  .Initiate_mask2 {
    margin-top: -55%;
  }
}
@media screen and (min-width: 1900px) {
  .Initiate_mask2 {
    margin-top: -45%;
  }
}
</style>

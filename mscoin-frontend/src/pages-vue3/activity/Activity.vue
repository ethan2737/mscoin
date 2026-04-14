<template>
  <div class="activity">
    <img class="bannerimg banner-pc" src="../../assets/images/activity-bg_pro.jpg">
    <div class="activity_container">
      <div class="main">
        <el-tabs v-model="currentCate" @tab-change="tabChange">
          <el-tab-pane :label="$t('activity.all')" name="all">
            <div class="activity-container">
              <div class="tips-line" v-if="activityList.all.loaded && activityList.all.total == 0">
                <img src="../../assets/images/emptydata.png" alt="无数据">
              </div>
              <div class="tips-line" v-if="!activityList.all.loaded">
                <el-spinner size="large" />
              </div>
              <div class="activity-item" v-for="(item, index) in activityList.all.items" :key="item.id || index">
                <div class="activity-type" v-if="item.type==1">{{$t('activity.activitytype1')}}</div>
                <div class="activity-type" v-if="item.type==2">{{$t('activity.activitytype2')}}</div>
                <div class="activity-type" v-if="item.type==3">{{$t('activity.activitytype3')}}</div>
                <div class="activity-type" v-if="item.type==4">{{$t('activity.activitytype4')}}</div>
                <div class="activity-type" v-if="item.type==5">{{$t('activity.activitytype5')}}</div>
                <div class="row-container">
                  <div class="col-img"><img :src="item.smallImageUrl" alt=""></div>
                  <div class="col-content">
                    <div class="title">
                      <span>{{item.title}}</span>
                      <div v-if="item.step==0" class="step0">{{$t('activity.prepare')}}</div>
                      <div v-if="item.step==1" class="step1">{{$t('activity.ongoing')}}</div>
                      <div v-if="item.step==2" class="step2">{{$t('activity.tokendistribute')}}</div>
                      <div v-if="item.step==3" class="step3">{{$t('activity.completed')}}</div>
                    </div>
                    <div style="width: 100%;padding-top: 10px;">
                      <div class="row-inner">
                        <div class="col-progress">
                          <div class="progress-text" style="margin-bottom: -3px;">
                            <span class="gray">{{$t('activity.progress')}}</span>
                            <span class="gray">{{$t('activity.totalsupply')}}</span>
                          </div>
                          <el-progress :percentage="item.progress" status="active" :stroke-width="5" :show-text="false" />
                          <div class="progress-text">
                            <span>{{fixedScale(item.progress, 2)}}%</span>
                            <span>{{fixedScale(item.totalSupply, item.amountScale)}} {{item.unit}}</span>
                          </div>
                          <div class="progress-time-wrap">
                            <p class="progress-time">{{$t('activity.starttime')}}：{{dateFormat(item.startTime)}}</p>
                            <p class="progress-time">{{$t('activity.endtime')}}：{{dateFormat(item.endTime)}}</p>
                          </div>
                        </div>
                      </div>
                      <div class="bottom-panel">
                        <div class="bottom">
                          <span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span>
                          <router-link v-if="item.step==1" class="el-button el-button--warning" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                        <div class="bottom-mobile">
                          <p><span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span></p>
                          <router-link v-if="item.step==1" class="el-button el-button--warning el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="page" v-if="activityList.all.total > 0">
              <el-pagination
                :total="activityList.all.total"
                :page-size="activityList.all.pageSize"
                :current-page="activityList.all.pageNo"
                layout="prev, pager, next"
                @current-change="(page) => pageChange(page)"
              />
            </div>
          </el-tab-pane>
          <el-tab-pane :label="$t('activity.prepare')" name="step0">
            <div class="activity-container">
              <div class="tips-line" v-if="activityList.step0.loaded && activityList.step0.total == 0">
                <img src="../../assets/images/emptydata.png" alt="无数据">
              </div>
              <div class="tips-line" v-if="!activityList.step0.loaded">
                <el-spinner size="large" />
              </div>
              <div class="activity-item" v-for="(item, index) in activityList.step0.items" :key="item.id || index">
                <div class="activity-type" v-if="item.type==1">{{$t('activity.activitytype1')}}</div>
                <div class="activity-type" v-if="item.type==2">{{$t('activity.activitytype2')}}</div>
                <div class="activity-type" v-if="item.type==3">{{$t('activity.activitytype3')}}</div>
                <div class="activity-type" v-if="item.type==4">{{$t('activity.activitytype4')}}</div>
                <div class="activity-type" v-if="item.type==5">{{$t('activity.activitytype5')}}</div>
                <div class="row-container">
                  <div class="col-img"><img :src="item.smallImageUrl" alt=""></div>
                  <div class="col-content">
                    <div class="title">
                      <span>{{item.title}}</span>
                      <div v-if="item.step==0" class="step0">{{$t('activity.prepare')}}</div>
                      <div v-if="item.step==1" class="step1">{{$t('activity.ongoing')}}</div>
                      <div v-if="item.step==2" class="step2">{{$t('activity.tokendistribute')}}</div>
                      <div v-if="item.step==3" class="step3">{{$t('activity.completed')}}</div>
                    </div>
                    <div style="width: 100%;padding-top: 10px;">
                      <div class="row-inner">
                        <div class="col-progress">
                          <div class="progress-text" style="margin-bottom: -3px;">
                            <span class="gray">{{$t('activity.progress')}}</span>
                            <span class="gray">{{$t('activity.totalsupply')}}</span>
                          </div>
                          <el-progress :percentage="item.progress" status="active" :stroke-width="5" :show-text="false" />
                          <div class="progress-text">
                            <span>{{fixedScale(item.progress, 2)}}%</span>
                            <span>{{fixedScale(item.totalSupply, item.amountScale)}} {{item.unit}}</span>
                          </div>
                          <div class="progress-time-wrap">
                            <p class="progress-time">{{$t('activity.starttime')}}：{{dateFormat(item.startTime)}}</p>
                            <p class="progress-time">{{$t('activity.endtime')}}：{{dateFormat(item.endTime)}}</p>
                          </div>
                        </div>
                      </div>
                      <div class="bottom-panel">
                        <div class="bottom">
                          <span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span>
                          <router-link v-if="item.step==1" class="el-button el-button--warning" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                        <div class="bottom-mobile">
                          <p><span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span></p>
                          <router-link v-if="item.step==1" class="el-button el-button--warning el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="page" v-if="activityList.step0.total > 0">
              <el-pagination
                :total="activityList.step0.total"
                :page-size="activityList.step0.pageSize"
                :current-page="activityList.step0.pageNo"
                layout="prev, pager, next"
                @current-change="(page) => pageChange(page)"
              />
            </div>
          </el-tab-pane>
          <el-tab-pane :label="$t('activity.ongoing')" name="step1">
            <div class="activity-container">
              <div class="tips-line" v-if="activityList.step1.loaded && activityList.step1.total == 0">
                <img src="../../assets/images/emptydata.png" alt="无数据">
              </div>
              <div class="tips-line" v-if="!activityList.step1.loaded">
                <el-spinner size="large" />
              </div>
              <div class="activity-item" v-for="(item, index) in activityList.step1.items" :key="item.id || index">
                <div class="activity-type" v-if="item.type==1">{{$t('activity.activitytype1')}}</div>
                <div class="activity-type" v-if="item.type==2">{{$t('activity.activitytype2')}}</div>
                <div class="activity-type" v-if="item.type==3">{{$t('activity.activitytype3')}}</div>
                <div class="activity-type" v-if="item.type==4">{{$t('activity.activitytype4')}}</div>
                <div class="activity-type" v-if="item.type==5">{{$t('activity.activitytype5')}}</div>
                <div class="row-container">
                  <div class="col-img"><img :src="item.smallImageUrl" alt=""></div>
                  <div class="col-content">
                    <div class="title">
                      <span>{{item.title}}</span>
                      <div v-if="item.step==0" class="step0">{{$t('activity.prepare')}}</div>
                      <div v-if="item.step==1" class="step1">{{$t('activity.ongoing')}}</div>
                      <div v-if="item.step==2" class="step2">{{$t('activity.tokendistribute')}}</div>
                      <div v-if="item.step==3" class="step3">{{$t('activity.completed')}}</div>
                    </div>
                    <div style="width: 100%;padding-top: 10px;">
                      <div class="row-inner">
                        <div class="col-progress">
                          <div class="progress-text" style="margin-bottom: -3px;">
                            <span class="gray">{{$t('activity.progress')}}</span>
                            <span class="gray">{{$t('activity.totalsupply')}}</span>
                          </div>
                          <el-progress :percentage="item.progress" status="active" :stroke-width="5" :show-text="false" />
                          <div class="progress-text">
                            <span>{{fixedScale(item.progress, 2)}}%</span>
                            <span>{{fixedScale(item.totalSupply, item.amountScale)}} {{item.unit}}</span>
                          </div>
                          <div class="progress-time-wrap">
                            <p class="progress-time">{{$t('activity.starttime')}}：{{dateFormat(item.startTime)}}</p>
                            <p class="progress-time">{{$t('activity.endtime')}}：{{dateFormat(item.endTime)}}</p>
                          </div>
                        </div>
                      </div>
                      <div class="bottom-panel">
                        <div class="bottom">
                          <span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span>
                          <router-link v-if="item.step==1" class="el-button el-button--warning" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                        <div class="bottom-mobile">
                          <p><span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span></p>
                          <router-link v-if="item.step==1" class="el-button el-button--warning el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="page" v-if="activityList.step1.total > 0">
              <el-pagination
                :total="activityList.step1.total"
                :page-size="activityList.step1.pageSize"
                :current-page="activityList.step1.pageNo"
                layout="prev, pager, next"
                @current-change="(page) => pageChange(page)"
              />
            </div>
          </el-tab-pane>
          <el-tab-pane :label="$t('activity.distributing')" name="step2">
            <div class="activity-container">
              <div class="tips-line" v-if="activityList.step2.loaded && activityList.step2.total == 0">
                <img src="../../assets/images/emptydata.png" alt="无数据">
              </div>
              <div class="tips-line" v-if="!activityList.step2.loaded">
                <el-spinner size="large" />
              </div>
              <div class="activity-item" v-for="(item, index) in activityList.step2.items" :key="item.id || index">
                <div class="activity-type" v-if="item.type==1">{{$t('activity.activitytype1')}}</div>
                <div class="activity-type" v-if="item.type==2">{{$t('activity.activitytype2')}}</div>
                <div class="activity-type" v-if="item.type==3">{{$t('activity.activitytype3')}}</div>
                <div class="activity-type" v-if="item.type==4">{{$t('activity.activitytype4')}}</div>
                <div class="activity-type" v-if="item.type==5">{{$t('activity.activitytype5')}}</div>
                <div class="row-container">
                  <div class="col-img"><img :src="item.smallImageUrl" alt=""></div>
                  <div class="col-content">
                    <div class="title">
                      <span>{{item.title}}</span>
                      <div v-if="item.step==0" class="step0">{{$t('activity.prepare')}}</div>
                      <div v-if="item.step==1" class="step1">{{$t('activity.ongoing')}}</div>
                      <div v-if="item.step==2" class="step2">{{$t('activity.tokendistribute')}}</div>
                      <div v-if="item.step==3" class="step3">{{$t('activity.completed')}}</div>
                    </div>
                    <div style="width: 100%;padding-top: 10px;">
                      <div class="row-inner">
                        <div class="col-progress">
                          <div class="progress-text" style="margin-bottom: -3px;">
                            <span class="gray">{{$t('activity.progress')}}</span>
                            <span class="gray">{{$t('activity.totalsupply')}}</span>
                          </div>
                          <el-progress :percentage="item.progress" status="active" :stroke-width="5" :show-text="false" />
                          <div class="progress-text">
                            <span>{{fixedScale(item.progress, 2)}}%</span>
                            <span>{{fixedScale(item.totalSupply, item.amountScale)}} {{item.unit}}</span>
                          </div>
                          <div class="progress-time-wrap">
                            <p class="progress-time">{{$t('activity.starttime')}}：{{dateFormat(item.startTime)}}</p>
                            <p class="progress-time">{{$t('activity.endtime')}}：{{dateFormat(item.endTime)}}</p>
                          </div>
                        </div>
                      </div>
                      <div class="bottom-panel">
                        <div class="bottom">
                          <span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span>
                          <router-link v-if="item.step==1" class="el-button el-button--warning" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                        <div class="bottom-mobile">
                          <p><span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span></p>
                          <router-link v-if="item.step==1" class="el-button el-button--warning el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="page" v-if="activityList.step2.total > 0">
              <el-pagination
                :total="activityList.step2.total"
                :page-size="activityList.step2.pageSize"
                :current-page="activityList.step2.pageNo"
                layout="prev, pager, next"
                @current-change="(page) => pageChange(page)"
              />
            </div>
          </el-tab-pane>
          <el-tab-pane :label="$t('activity.completed')" name="step3">
            <div class="activity-container">
              <div class="tips-line" v-if="activityList.step3.loaded && activityList.step3.total == 0">
                <img src="../../assets/images/emptydata.png" alt="无数据">
              </div>
              <div class="tips-line" v-if="!activityList.step3.loaded">
                <el-spinner size="large" />
              </div>
              <div class="activity-item" v-for="(item, index) in activityList.step3.items" :key="item.id || index">
                <div class="activity-type" v-if="item.type==1">{{$t('activity.activitytype1')}}</div>
                <div class="activity-type" v-if="item.type==2">{{$t('activity.activitytype2')}}</div>
                <div class="activity-type" v-if="item.type==3">{{$t('activity.activitytype3')}}</div>
                <div class="activity-type" v-if="item.type==4">{{$t('activity.activitytype4')}}</div>
                <div class="activity-type" v-if="item.type==5">{{$t('activity.activitytype5')}}</div>
                <div class="row-container">
                  <div class="col-img"><img :src="item.smallImageUrl" alt=""></div>
                  <div class="col-content">
                    <div class="title">
                      <span>{{item.title}}</span>
                      <div v-if="item.step==0" class="step0">{{$t('activity.prepare')}}</div>
                      <div v-if="item.step==1" class="step1">{{$t('activity.ongoing')}}</div>
                      <div v-if="item.step==2" class="step2">{{$t('activity.tokendistribute')}}</div>
                      <div v-if="item.step==3" class="step3">{{$t('activity.completed')}}</div>
                    </div>
                    <div style="width: 100%;padding-top: 10px;">
                      <div class="row-inner">
                        <div class="col-progress">
                          <div class="progress-text" style="margin-bottom: -3px;">
                            <span class="gray">{{$t('activity.progress')}}</span>
                            <span class="gray">{{$t('activity.totalsupply')}}</span>
                          </div>
                          <el-progress :percentage="item.progress" status="active" :stroke-width="5" :show-text="false" />
                          <div class="progress-text">
                            <span>{{fixedScale(item.progress, 2)}}%</span>
                            <span>{{fixedScale(item.totalSupply, item.amountScale)}} {{item.unit}}</span>
                          </div>
                          <div class="progress-time-wrap">
                            <p class="progress-time">{{$t('activity.starttime')}}：{{dateFormat(item.startTime)}}</p>
                            <p class="progress-time">{{$t('activity.endtime')}}：{{dateFormat(item.endTime)}}</p>
                          </div>
                        </div>
                      </div>
                      <div class="bottom-panel">
                        <div class="bottom">
                          <span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span>
                          <router-link v-if="item.step==1" class="el-button el-button--warning" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary" :to="'/lab/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                        <div class="bottom-mobile">
                          <p><span><el-icon><InfoFilled /></el-icon> {{item.detail}}</span></p>
                          <router-link v-if="item.step==1" class="el-button el-button--warning el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                          <router-link v-else class="el-button el-button--primary el-button--long" :to="'/lab/detail/'+ item.id">{{$t('activity.viewdetail')}}</router-link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="page" v-if="activityList.step3.total > 0">
              <el-pagination
                :total="activityList.step3.total"
                :page-size="activityList.step3.pageSize"
                :current-page="activityList.step3.pageNo"
                layout="prev, pager, next"
                @current-change="(page) => pageChange(page)"
              />
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
    <div class="app_bottom">
      <div class="left_logo">
        <img style="float:left;" src="../../assets/images/applogo.png" alt="">
        <div style="float:left;height: 40px;line-height:40px;color:#000;">{{$t("cms.downloadslogan")}}</div>
      </div>
      <div class="right_btn_wrap"><router-link to="/app" class="right_btn">{{$t("cms.download")}}</router-link></div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 活动首页
 */
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElTabs, ElTabPane, ElProgress, ElPagination, ElIcon } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vuex'
import moment from 'moment'

const store = useStore()

const host = ''

const currentCate = ref('all')

const activityList = reactive({
  all: { loaded: false, pageSize: 10, pageNo: 1, total: 0, items: [] },
  step0: { loaded: false, pageSize: 10, pageNo: 1, total: 0, items: [] },
  step1: { loaded: false, pageSize: 10, pageNo: 1, total: 0, items: [] },
  step2: { loaded: false, pageSize: 10, pageNo: 1, total: 0, items: [] },
  step3: { loaded: false, pageSize: 10, pageNo: 1, total: 0, items: [] }
})

const lang = computed(() => store.state.lang)

const langPram = computed(() => {
  if (store.state.lang === '简体中文') return 'CN'
  if (store.state.lang === 'English') return 'EN'
  return 'CN'
})

const dateFormat = (tick) => {
  return moment(tick).format('YYYY-MM-DD HH:mm:ss')
}

const fixedScale = (value, scale) => {
  return value.toFixed(scale)
}

const init = () => {
    store.commit('navigate', '/activity')
}

const detail = (aId) => {
  router.push('/lab/detail/' + aId)
}

const pageChange = (page) => {
  if (currentCate.value === 'all') {
    activityList.all.pageNo = page
    getData(-1)
  }
  if (currentCate.value === 'step0') {
    activityList.step0.pageNo = page
    getData(0)
  }
  if (currentCate.value === 'step1') {
    activityList.step1.pageNo = page
    getData(1)
  }
  if (currentCate.value === 'step2') {
    activityList.step2.pageNo = page
    getData(2)
  }
  if (currentCate.value === 'step3') {
    activityList.step3.pageNo = page
    getData(3)
  }
}

const tabChange = (name) => {
  currentCate.value = name
  if (name === 'all' && !activityList.all.loaded) getData(-1)
  if (name === 'step0' && !activityList.step0.loaded) getData(0)
  if (name === 'step1' && !activityList.step1.loaded) getData(1)
  if (name === 'step2' && !activityList.step2.loaded) getData(2)
  if (name === 'step3' && !activityList.step3.loaded) getData(3)
}

const getData = (step = -1) => {
  const param = { step, pageSize: 10 }
  if (step === -1) param.pageNo = activityList.all.pageNo
  if (step === 0) param.pageNo = activityList.step0.pageNo
  if (step === 1) param.pageNo = activityList.step1.pageNo
  if (step === 2) param.pageNo = activityList.step2.pageNo
  if (step === 3) param.pageNo = activityList.step3.pageNo

  axios.post(`${host}/uc/activity/page-query`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(res => {
    if (res.status === 200 && res.data.code === 0) {
      const aList = res.data.data.content
      if (step === -1) {
        activityList.all.loaded = true
        activityList.all.items = aList
        activityList.all.total = res.data.data.totalElements
      } else if (step === 0) {
        activityList.step0.loaded = true
        activityList.step0.items = aList
        activityList.step0.total = res.data.data.totalElements
      } else if (step === 1) {
        activityList.step1.loaded = true
        activityList.step1.items = aList
        activityList.step1.total = res.data.data.totalElements
      } else if (step === 2) {
        activityList.step2.loaded = true
        activityList.step2.items = aList
        activityList.step2.total = res.data.data.totalElements
      } else if (step === 3) {
        activityList.step3.loaded = true
        activityList.step3.items = aList
        activityList.step3.total = res.data.data.totalElements
      }
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

onMounted(() => {
  init()
  getData()
})
</script>

<style scoped>
/* Tabs 样式完全覆盖 Element Plus 默认蓝色 */
.activity :deep(.el-tabs-bar) {
  border-bottom: 1px solid #27313e;
}

/* 强制设置所有 tab 项为白色 */
.activity :deep(.el-tabs-nav .el-tabs-tab),
.activity :deep(.el-tabs__item),
.activity :deep(.el-tabs--top .el-tabs__item) {
  color: #fff !important;
}

/* hover 状态 */
.activity :deep(.el-tabs-nav .el-tabs-tab:hover),
.activity :deep(.el-tabs__item:hover) {
  color: #fff !important;
}

/* 选中状态 */
.activity :deep(.el-tabs-nav .el-tabs-tab.is-active),
.activity :deep(.el-tabs__item.is-active) {
  color: #f0a70a !important;
}

/* 底部激活条 */
.activity :deep(.el-tabs__active-bar) {
  background-color: #f0a70a !important;
}

.app_bottom {
  display: none;
  z-index: 999;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 40px;
  background: #17212e !important;
  border-top: 1px solid #1f2833;
}

.app_bottom .left_logo img {
  height: 30px;
  margin-top: 5px;
  border-radius: 5px;
  margin-left: 5px;
  margin-right: 5px;
}

.app_bottom .right_btn_wrap {
  float: right;
  height: 40px;
  line-height: 40px;
  margin-right: 5px;
}

.app_bottom .right_btn {
  color: #FFF;
  border-radius: 3px;
  padding: 0px 10px;
  line-height: 26px;
  height: 26px;
  display: block;
  margin-top: 7px;
  background: linear-gradient(200deg, #ff9900, #ffbe2b, #ffa500);
}

@media screen and (max-width: 768px) {
  .activity {
    padding-top: 45px !important;
  }
  .activity_container {
    padding: 0 2% !important;
  }
  .activity .main {
    margin-top: 20px !important;
  }
  .progress-time {
    text-align: left !important;
    padding-right: 0px !important;
    margin-top: 10px !important;
  }
  .activity-item .title {
    text-align: left !important;
  }
  .activity-item .title div {
    display: none !important;
  }
  .activity-item .activity-type {
    display: none !important;
  }
  .banner-pc {
    display: none !important;
  }
  .banner-mobile {
    display: block !important;
  }
  .activity_container > h1 {
    font-size: 20px !important;
    margin-top: -170px !important;
  }
  .activity-item img {
    width: 100px !important;
    height: 100px !important;
  }
  .activity-item .title span {
    font-size: 18px !important;
  }
  .app_bottom {
    display: block !important;
  }
  .bottom-panel .bottom {
    display: none !important;
  }
  .bottom-panel .bottom-mobile {
    display: block !important;
  }
  .bottom-panel .bottom-mobile p {
    margin: 10px 0 10px 0;
  }
  .bottom-panel .bottom-mobile p span {
    font-size: 12px;
    color: #999;
    margin-top: 15px;
  }
}

.banner-pc {
  display: block;
}

.banner-mobile {
  display: none;
}
</style>

<style lang="scss">
/* 全局样式，强制覆盖 Element Plus tabs 默认颜色 */
.activity .el-tabs__item {
  color: #fff !important;

  &:hover {
    color: #fff !important;
  }

  &.is-active {
    color: #f0a70a !important;
  }
}
</style>

<style lang="scss" scoped>
.activity {
  background: #0b1520;
  min-height: 100vh;
  position: relative;
  overflow: hidden;
  padding-bottom: 50px;
  padding-top: 60px;
  color: #fff;
}

.activity .bannerimg {
  width: 100%;
}

.activity_container {
  padding: 0 2%;
  text-align: center;
  height: 100%;
  min-height: 600px;

  > h1 {
    margin-top: -200px;
    font-size: 32px;
    line-height: 1;
    padding: 50px 0 20px 0;
    letter-spacing: 3px;
  }
}

.activity .main {
  margin-top: 40px;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
}

.activity-container {
  width: 100%;
  min-height: 200px;
}

.tips-line {
  height: 100px;
  line-height: 100px;
  text-align: center;
  position: relative;
  display: inline-block;

  img {
    width: 100px;
    margin-top: 50px;
  }

  .el-spinner {
    margin-top: 50px;
  }
}

.activity-item {
  position: relative;
  overflow: visible;
  padding: 15px 20px;
  margin: 25px auto;
  max-width: 1200px;
  width: calc(100% - 40px);
  min-height: 200px;
  background: #17212e;
  border-radius: 8px;
  border: 1px solid #1f2833;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  box-sizing: border-box;

  &:hover {
    box-shadow: 0 0 25px rgba(240, 167, 10, 0.3);
    transform: scale(1.01, 1.01);
    border-color: #f0a70a;
  }

  .activity-type {
    position: absolute;
    width: 140px;
    height: 25px;
    line-height: 25px;
    background-color: #000;
    color: white;
    text-align: center;
    transform: rotate(-40deg);
    top: 23px;
    left: -30px;
    z-index: 99;
    box-shadow: 1px 1px 4px #000;
  }

  .title {
    width: 100%;
    padding-top: 10px;
    display: flex;
    flex-direction: row;
    align-items: center;
    flex-wrap: wrap;

    span {
      font-size: 22px;
      color: #fff !important;
    }

    div {
      height: 30px;
      line-height: 30px;
      padding: 0 15px 0 25px;
      border-top-left-radius: 15px;
      border-bottom-left-radius: 15px;
      border-top-right-radius: 3px;
      border-bottom-right-radius: 3px;
      position: relative;
      white-space: nowrap;

      &:before {
        content: "●";
        position: absolute;
        top: -1px;
        left: 5px;
      }
    }

    div.step0 {
      margin-left: 15px;
      color: #FFF;
      border: 1px solid #FFF;
      background: #5bacff;
      background-image: repeating-linear-gradient(60deg, hsla(0,0%,100%,.1), hsla(0,0%,100%,.1) 10px, transparent 0, transparent 20px);
    }

    div.step1 {
      margin-left: 15px;
      color: #ffffff;
      border: 1px solid #19be6b;
      background: #19be6b;
      background-image: repeating-linear-gradient(60deg, hsla(0,0%,100%,.1), hsla(0,0%,100%,.1) 10px, transparent 0, transparent 20px);
    }

    div.step2 {
      margin-left: 15px;
      color: #FFFFFF;
      border: 1px solid #f0a70a;
      background: #f0a70a;
      background-image: repeating-linear-gradient(60deg, hsla(0,0%,100%,.1), hsla(0,0%,100%,.1) 10px, transparent 0, transparent 20px);
    }

    div.step3 {
      margin-left: 15px;
      color: #ffffff;
      border: 1px solid #828282;
      background: #828282;
      background-image: repeating-linear-gradient(60deg, hsla(0,0%,100%,.1), hsla(0,0%,100%,.1) 10px, transparent 0, transparent 20px);
    }
  }

  img {
    width: 160px;
    height: 160px;
    border-radius: 8px;
  }
}

.row-container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 15px;
}

.col-img {
  flex: 0 0 auto;
  width: 160px;
  flex-shrink: 0;
}

.col-content {
  flex: 1;
  padding-left: 15px;
  min-width: 0;
}

.row-inner {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 10px;
}

.col-progress {
  flex: 1;
  min-width: 200px;
}

.progress-time-wrap {
  margin-top: 15px;
}

.progress-text {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;

  span {
    font-size: 12px;
    color: #ccc;
  }
}

.bottom-panel {
  border-top: 1px solid #1f2833;
  margin-top: 15px;

  .bottom-mobile {
    display: none;
  }

  .bottom {
    display: flex;
    flex-direction: row;
    justify-content: space-between;

    span {
      font-size: 12px;
      color: #ccc;
      margin-top: 15px;
    }

    button, a {
      margin-top: 11px;
    }

    a.el-button--primary {
      background: #f0a70a;
      border-color: #f0a70a;
      color: #fff;

      &:hover {
        background: #ffb319;
        border-color: #ffb319;
      }
    }
  }
}

.progress-time {
  font-size: 13px;
  letter-spacing: 1px;
  text-align: left;
  margin: 8px 0 0 0;
  color: #ccc;
}

.right {
  float: right;
}

.left {
  float: left;
}

.gray {
  color: #ccc;
}
</style>

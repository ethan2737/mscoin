<template>
  <div class="notice">
    <div class="notice-list" style="margin-top: 100px; color: #696969;">
      <div class="main">
        <h2>{{$t("cms.noticecenter")}}</h2>
        <div class="list">
          <div :class="item.id==queryId?'active' : 'item1'" v-for="item in FAQList" :key="item.id" @click="noticedeail(item.id)">
            <span class="text">{{item.title}}</span>
          </div>
        </div>
        <div class="page">
          <el-pagination
            :total="totalNum"
            :page-size="pageSize"
            :current-page="pageNo"
            layout="prev, pager, next"
            @current-change="loadDataPage"
            size="small"
          />
        </div>
      </div>
    </div>
    <div class="content-wrap">
      <div v-show="hasContent">
        <div class="header">
          <h2>{{data.info.title}}</h2>
          <span>{{data.info.createTime}}</span>
        </div>
        <div class="content">
          <div class="content-text" v-html="data.info.content"></div>
          <div class="show-qrcode" style="text-align: right;padding: 15px 10px;border-radius:10px;background:#FFF;margin-top: 30px;">
            <div style="text-align: right;">{{$t("cms.scanforshare")}}</div>
          </div>
        </div>
        <div class="nav-bottom">
          <div class="left" v-if="data.back">
            <router-link class="link" :to="'../announcement/'+data.back.id">
              < {{$t("cms.prevnotice")}} <p style="color:#f0a70a;">{{data.back.title}}</p>
            </router-link>
          </div>
          <div class="right" v-if="data.next">
            <router-link class="link" :to="'../announcement/'+data.next.id">{{$t("cms.nextnotice")}} >
              <p style="color:#f0a70a;">{{data.next.title}}</p>
            </router-link>
          </div>
        </div>
      </div>
      <div v-show="!hasContent">
        <p style="font-size: 30px;text-align:center;margin-top: 15%;">暂无内容</p>
        <p style="text-align:center;font-size: 12px;margin-top: 10px;color: #828ea1;">{{$t("cms.notexist")}}</p>
      </div>
      <div v-if="spinShow" class="loading-overlay">加载中...</div>
    </div>
    <div class="bottom-list">
      <p style="font-size: 18px;margin: 15px 0;">最新公告</p>
      <div class="notice-item" v-for="item in FAQList" :key="item.id" @click="noticedeail(item.id)">
        <span class="text">[{{subTime(item.createTime)}}] {{item.title}}</span>
      </div>
    </div>
    <div class="app_bottom">
      <div class="left_logo">
        <img style="float:left;" src="../../assets/images/applogo.png"></img>
        <div style="float:left;height: 40px;line-height:40px;color:#000;">{{$t("cms.downloadslogan")}}</div>
      </div>
      <div class="right_btn_wrap"><router-link to="/app" class="right_btn">{{$t("cms.download")}}</router-link></div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - CMS 公告详情页面
 */
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { ElMessage, ElPagination } from 'element-plus'
import axios from 'axios'
import { useRouter, useRoute, useStore } from 'vue-router/composables'

const router = useRouter()
const route = useRoute()
const store = useStore()

const host = 'http://localhost'

const data = reactive({ info: {} })
const queryId = ref('')
const initLang = ref(store.state.lang)
const hasContent = ref(true)
const pageNo = ref(1)
const pageSize = ref(10)
const totalNum = ref(0)
const FAQList = ref([])
const spinShow = ref(false)

const lang = computed(() => {
  if (store.state.lang !== initLang.value) {
    router.back()
  }
  return store.state.lang
})

const langPram = computed(() => {
  if (store.state.lang === '简体中文') {
    return 'CN'
  }
  if (store.state.lang === 'English') {
    return 'EN'
  }
  return 'CN'
})

const subTime = (str) => {
  if (!str) return ''
  return str.substring(5, 10)
}

const loadNoticeInfo = () => {
  let id = route.params.id
  if (id == null || id == '' || id == 0) {
    if (FAQList.value.length > 0) {
      id = FAQList.value[0].id
    } else {
      id = 0
    }
  }
  queryId.value = id
  spinShow.value = true

  const param = { id: id, lang: langPram.value }
  axios.post(`${host}/uc/announcement/more`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const result = response.data
    if (result.code == 0) {
      const responseData = result.data
      Object.assign(data, responseData)
      hasContent.value = true
      spinShow.value = false
    } else {
      hasContent.value = false
      spinShow.value = false
    }
  }).catch(() => {
    hasContent.value = false
    spinShow.value = false
    ElMessage.error('请求失败')
  })
}

const fetchData = () => {
  const id = route.params.id
  if (id == null || id == '') {
    return
  }
  loadNoticeInfo()
}

const noticedeail = (id) => {
  router.push({ path: '/announcement/' + id })
}

const loadDataPage = (pageIndex) => {
  const param = {
    pageNo: pageIndex,
    pageSize: pageSize.value,
    lang: langPram.value
  }
  axios.post(`${host}/uc/ancillary/announcement`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      if (resp.data.content.length == 0) return
      FAQList.value = resp.data.content
      totalNum.value = resp.data.totalElements
      loadNoticeInfo()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

watch(() => route.params, () => {
  fetchData()
}, { immediate: true })

onMounted(() => {
  store.commit('navigate', 'nav-service')
  loadDataPage(pageNo.value)
})
</script>

<style scoped>
.main {
  color: #696969;
}

.notice {
  padding-bottom: 20px;
  overflow-x: hidden;
  padding: 0px 180px;
  min-height: 800px;
  background: rgba(242, 246, 250, 1) !important;
}

.bottom-list {
  display: none;
  padding: 10px 20px 20px 20px;
  color: #696969;
  background-color: #FFF;
}

.bottom-list .notice-item span {
  display: block;
  margin: 10px 0;
  color: #f0a70a;
  font-size: 14px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
  overflow: hidden;
}

@media screen and (max-width: 768px) {
  .notice {
    padding: 0px 0px;
    position: relative;
  }
  .show-qrcode {
    display: none;
  }
  .bottom-list {
    display: block !important;
  }
  .notice .content-wrap {
    margin-top: 45px;
    padding: 30px 20px 50px 20px;
    margin-bottom: 0px;
  }
  .notice-list {
    display: none;
  }
  .notice .content-wrap .header h2 {
    text-align: left;
  }
  .notice .content-wrap .content {
    padding-bottom: 60px;
  }
  .notice .nav-bottom {
    padding: 30px 0px 30px 0px;
  }
  .notice .nav-bottom .link > p {
    display: none;
  }
  .app_bottom {
    display: block !important;
  }
}

.content-wrap {
  width: 1200px;
  margin: 0 auto;
}

.header {
  text-align: right;
  margin-bottom: 10px;
  padding: 10px 0;
  border-bottom: 1px solid #ebebeb;
}

.header h2 {
  text-align: center;
  margin-bottom: 10px;
  font-weight: normal;
  color: #696969;
}

.header span {
  padding: 0 10px;
  color: #999;
}

.loading-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(255, 255, 255, 0.82);
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
  color: #696969;
}

.app_bottom {
  display: none;
  z-index: 999;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 40px;
  background: rgba(242, 246, 250, 1) !important;
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
</style>

<style lang="scss">
.notice {
  .main {
    h2 {
      text-align: left;
      font-size: 18px;
      font-weight: normal;
    }

    .list {
      margin-top: 20px;

      .item1 {
        width: 100%;
        text-align: left;
        padding: 15px 15px;
        color: rgba(130, 142, 161, 1);

        &:hover {
          cursor: pointer;
          color: #f0a70a;
        }
      }

      .active {
        background: #FFF;
        color: #f0a70a;
        width: 100%;
        text-align: left;
        padding: 15px 15px;
        border-radius: 3px;

        &:hover {
          cursor: pointer;
          color: #df9800;
        }
      }
    }

    .page {
      text-align: right;
      margin-top: 10px;

      :deep(.el-pagination) {
        .btn-prev,
        .btn-next {
          background: transparent !important;
          color: #000;
          border: none;
        }

        .el-pager li {
          background-color: transparent !important;
          color: #000;
          border: none;
        }
      }
    }
  }

  .content-wrap {
    text-align: left;
    position: relative;
    width: 100%;
    min-height: 562px;
    background-color: #FFF;
    color: #000;
    margin-top: 100px;
    border-radius: 3px;
    padding: 50px 50px;
    margin-bottom: 50px;
    padding-bottom: 110px;

    .link {
      font-size: 14px;
      color: #f0a70a;
    }

    .header {
      margin-bottom: 40px;
      padding-bottom: 10px;

      h2 {
        font-size: 20px;
      }
    }

    .content {
      color: #74777a;
      margin-bottom: 80px;
      letter-spacing: 1px;
      line-height: 25px;

      .content-text {
        p {
          text-align: justify;

          &::after {
            width: 100%;
            content: '';
            height: 0;
          }
        }
      }

      span {
        .MsoNormal {
          font-weight: 100;
          line-height: 30px;
        }

        .p1 {
          .s1 {
            b {
              font-weight: 400;

              span {
                font-size: 18px !important;
              }
            }
          }
        }

        .p2 {
          span {
            font-size: 12pt !important;
            line-height: 30px;
          }
        }

        .p3 {
          span {
            font-size: 12pt !important;
            color: #fff !important;
            line-height: 30px;
          }
        }
      }

      .image-desc {
        .uploaded-img {
          width: 100%;
        }

        .image-caption {
          display: none !important;
        }
      }
    }
  }

  .nav-bottom {
    position: absolute;
    left: 0;
    bottom: 0;
    width: 90%;
    margin-left: 5%;
    padding: 50px 0px 50px 0px;
    border-top: 1px solid #ebebeb;
    border-bottom: 1px solid #ebebeb;

    .link > p {
      color: #fff;
      line-height: 1;
      margin-top: 10px;
    }

    .left {
      text-align: left;
      float: left;
    }

    .right {
      text-align: right;
      float: right;
    }
  }
}
</style>

<template>
  <div class="notice">
    <div class="notice-list" style="margin-top: 100px; color: #696969;">
      <div class="main">
        <h2>{{ $t('cms.noticecenter') }}</h2>
        <div class="list">
          <div
            v-for="item in noticeList"
            :key="item.id"
            :class="item.id === currentNoticeId ? 'active' : 'item1'"
            @click="openNotice(item.id)"
          >
            <span class="text">{{ item.title }}</span>
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
      <div v-if="isInitialLoading" class="loading-state">
        <div class="loading-title">加载中...</div>
      </div>

      <template v-else-if="hasContent">
        <div class="header">
          <h2>{{ noticeData.info.title }}</h2>
          <span>{{ noticeData.info.createTime }}</span>
        </div>
        <div class="content">
          <div class="content-text" v-html="noticeData.info.content"></div>
          <div class="show-qrcode" style="text-align: right; padding: 15px 10px; border-radius: 10px; background: #FFF; margin-top: 30px;">
            <div style="text-align: right;">{{ $t('cms.scanforshare') }}</div>
          </div>
        </div>
        <div class="nav-bottom">
          <div class="left" v-if="noticeData.back">
            <router-link class="link" :to="`/notice/item/${noticeData.back.id}`">
              &lt; {{ $t('cms.prevnotice') }}
              <p style="color: #f0a70a;">{{ noticeData.back.title }}</p>
            </router-link>
          </div>
          <div class="right" v-if="noticeData.next">
            <router-link class="link" :to="`/notice/item/${noticeData.next.id}`">
              {{ $t('cms.nextnotice') }} &gt;
              <p style="color: #f0a70a;">{{ noticeData.next.title }}</p>
            </router-link>
          </div>
        </div>
      </template>

      <div v-else>
        <p style="font-size: 30px; text-align: center; margin-top: 15%;">暂无内容</p>
        <p style="text-align: center; font-size: 12px; margin-top: 10px; color: #828ea1;">{{ $t('cms.notexist') }}</p>
      </div>
    </div>

    <div class="bottom-list">
      <p style="font-size: 18px; margin: 15px 0;">最新公告</p>
      <div v-for="item in noticeList" :key="item.id" class="notice-item" @click="openNotice(item.id)">
        <span class="text">[{{ subTime(item.createTime) }}] {{ item.title }}</span>
      </div>
    </div>

    <div class="app_bottom">
      <div class="left_logo">
        <img style="float: left;" src="../../assets/images/applogo.png" alt="">
        <div style="float: left; height: 40px; line-height: 40px; color: #000;">{{ $t('cms.downloadslogan') }}</div>
      </div>
      <div class="right_btn_wrap">
        <router-link to="/app" class="right_btn">{{ $t('cms.download') }}</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { ElMessage, ElPagination } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const store = useStore()

const host = ''
const pageSize = 10

const noticeData = reactive({ info: {} })
const noticeList = ref([])
const currentNoticeId = ref('')
const pageNo = ref(1)
const totalNum = ref(0)
const hasContent = ref(true)
const isInitialLoading = ref(true)
const initLang = ref(store.state.lang)

const langParam = computed(() => (store.state.lang === 'English' ? 'EN' : 'CN'))

watch(
  () => store.state.lang,
  (value) => {
    if (value !== initLang.value) {
      router.back()
    }
  }
)

watch(
  () => route.params.id,
  () => {
    if (noticeList.value.length > 0) {
      void loadNoticeInfo()
    }
  }
)

function subTime(str) {
  return str ? str.substring(5, 10) : ''
}

function openNotice(id) {
  router.push({ path: `/notice/item/${id}` })
}

async function loadNoticeInfo() {
  let id = route.params.id
  if (!id || id === '0') {
    id = noticeList.value[0]?.id || 0
  }

  currentNoticeId.value = id
  hasContent.value = true

  try {
    const response = await axios.post(
      `${host}/uc/announcement/more`,
      { id, lang: langParam.value },
      {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json',
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      }
    )
    const result = response.data
    if (result.code === 0 && result.data?.info) {
      Object.assign(noticeData, result.data)
      hasContent.value = true
    } else {
      hasContent.value = false
    }
  } catch {
    hasContent.value = false
    ElMessage.error('请求失败')
  } finally {
    isInitialLoading.value = false
  }
}

async function loadDataPage(pageIndex = 1) {
  pageNo.value = pageIndex
  isInitialLoading.value = noticeList.value.length === 0

  try {
    const response = await axios.post(
      `${host}/uc/ancillary/announcement`,
      {
        pageNo: pageIndex,
        pageSize,
        lang: langParam.value
      },
      {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json',
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      }
    )
    const result = response.data
    if (result.code === 0) {
      noticeList.value = result.data?.content || []
      totalNum.value = result.data?.totalElements || 0
      if (noticeList.value.length > 0) {
        await loadNoticeInfo()
      } else {
        hasContent.value = false
        isInitialLoading.value = false
      }
      return
    }
    ElMessage.error(result.message || '请求失败')
    hasContent.value = false
  } catch {
    hasContent.value = false
    ElMessage.error('请求失败')
  } finally {
    if (noticeList.value.length === 0) {
      isInitialLoading.value = false
    }
  }
}

onMounted(() => {
  store.commit('navigate', 'nav-service')
  void loadDataPage(1)
})
</script>

<style scoped>
.main {
  color: #696969;
}

.notice {
  padding-bottom: 20px;
  overflow-x: hidden;
  padding: 0 180px;
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

.content-wrap {
  width: 1200px;
  margin: 0 auto;
  text-align: left;
  position: relative;
  width: 100%;
  min-height: 562px;
  background-color: #FFF;
  color: #000;
  margin-top: 100px;
  border-radius: 3px;
  padding: 50px 50px 110px;
  margin-bottom: 50px;
}

.loading-state {
  min-height: 360px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.loading-title {
  font-size: 16px;
  color: #696969;
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
  padding: 0 10px;
  line-height: 26px;
  height: 26px;
  display: block;
  margin-top: 7px;
  background: linear-gradient(200deg, #ff9900, #ffbe2b, #ffa500);
}

@media screen and (max-width: 768px) {
  .notice {
    padding: 0;
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
    margin-bottom: 0;
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
    padding: 30px 0;
  }

  .notice .nav-bottom .link > p {
    display: none;
  }

  .app_bottom {
    display: block !important;
  }
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
        padding: 15px;
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
        padding: 15px;
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

      .p1 .s1 b span,
      .p2 span {
        font-size: 12pt !important;
        line-height: 30px;
      }

      .p3 span {
        font-size: 12pt !important;
        color: #fff !important;
        line-height: 30px;
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

  .nav-bottom {
    position: absolute;
    left: 0;
    bottom: 0;
    width: 90%;
    margin-left: 5%;
    padding: 50px 0;
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

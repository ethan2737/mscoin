<template>
  <div class="notice">
    <div class="banner">
      <span>{{$t("header.service")}}</span>
    </div>
    <div class="main">
      <div class="list">
        <div class="item" v-for="item in FAQList" :key="item.id" @click="noticedeail(item.id)">
          <img v-show="item.isTop==0" class="iconimg" src="../../assets/images/icon-top.png" alt="">
          <span class="text">{{item.title}}</span>
          <span class="time">{{item.createTime}}</span>
        </div>
      </div>
      <div class="page">
        <el-pagination
          :total="totalNum"
          :page-size="pageSize"
          :current-page="pageNo"
          layout="prev, pager, next"
          @current-change="loadDataPage"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - CMS 公告列表页面
 */
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElPagination } from 'element-plus'
import axios from 'axios'
import { useRouter, useStore } from 'vue-router/composables'

const router = useRouter()
const store = useStore()

const host = 'http://localhost'

const pageNo = ref(1)
const pageSize = ref(10)
const totalNum = ref(0)
const FAQList = ref([])

const lang = computed(() => store.state.lang)

const langPram = computed(() => {
  if (store.state.lang === '简体中文') {
    return 'CN'
  }
  if (store.state.lang === 'English') {
    return 'EN'
  }
  return 'CN'
})

const init = () => {
  store.state.HeaderActiveName = '1-7'
  store.commit('navigate', 'nav-service')
  loadDataPage(pageNo.value)
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
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const noticedeail = (id) => {
  router.push({ path: '/notice/index', query: { id: id } })
}

const titleLang = (str) => {
  const reg = new RegExp('[\\u4E00-\\u9FFF]+', 'g')
  if (reg.test(str)) {
    return '简体中文'
  } else {
    return 'English'
  }
}

onMounted(() => {
  init()
})
</script>

<style lang="scss" scoped>
.notice {
  .banner {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 320px;
    background: linear-gradient(to right, #001a40, #000109);
    background-size: 100% 100%;
    color: #fff;
    font-size: 40px;
  }

  .main {
    width: 70%;
    margin: 0 auto;
    background-color: #192330;
    color: #fff;
    margin-top: -50px;
    border-radius: 6px;
    padding: 50px 100px;
    margin-bottom: 50px;

    .list {
      font-size: 14px;

      .item {
        line-height: 50px;
        height: 50px;
        border-bottom: 1px solid #27313e;
        cursor: pointer;

        .iconimg {
          width: 35px;
          vertical-align: sub;
          margin-left: 10px;
          padding-bottom: 5px;
        }

        .time {
          float: right;
          color: #999;
          font-size: 14px;
        }

        .text {
          display: inline-block;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          width: 70%;

          &:hover {
            color: #f0a70a;
          }
        }
      }
    }

    .page {
      text-align: right;
      margin-top: 20px;

      :deep(.el-pagination) {
        .btn-prev,
        .btn-next {
          background-color: transparent;
          color: #fff;
          border: none;
        }

        .el-pager li {
          background-color: transparent;
          color: #fff;
          border: none;

          &.is-active {
            color: #f0a70a;
          }
        }
      }
    }
  }
}
</style>

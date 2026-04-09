<template>
  <div class="helplist">
    <div class="route-wrap">
      <router-link to="/help">{{$t("header.helpcenter")}}</router-link>
      <span>></span>
      <span>{{cateTitle}}</span>
    </div>
    <div class="container">
      <h1>{{cateTitle}}</h1>
      <div class="list">
        <router-link class="item" v-for="(item,index) in list" :key="index" :to="{path:'/help/detail',query:{cate:cate,id:item.id,cateTitle:cateTitle}}">
          <span class="text">{{item.title}}</span>
          <span class="time">{{item.createTime}}</span>
        </router-link>
      </div>
      <div class="page">
        <el-pagination
          :total="total"
          :page-size="pageSize"
          :current-page="pageNo"
          layout="prev, pager, next"
          @current-change="pageChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - CMS 帮助列表页面
 */
import { ref, computed, onMounted, watch } from 'vue'
import { ElMessage, ElPagination } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'

const store = useStore()
const route = useRoute()
const router = useRouter()

const host = ''

const cate = ref(0)
const pageNo = ref(1)
const pageSize = ref(10)
const total = ref(0)
const list = ref([])
const cateTitle = ref('')

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
  store.commit('navigate', 'nav-uc')
}

const pageChange = (data) => {
  pageNo.value = data
  getData()
}

const getData = () => {
  const params = {
    pageNo: pageNo.value,
    pageSize: pageSize.value,
    cate: cate.value,
    lang: langPram.value
  }
  axios.post(`${host}/uc/ancillary/more/help/page`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(res => {
    if (res.status == 200 && res.data.code == 0) {
      list.value = res.data.data.content
      total.value = res.data.data.totalElements
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

watch(() => route.query, (query) => {
  const { cate: queryCate, cateTitle: queryTitle } = query || {}
  cate.value = queryCate
  cateTitle.value = queryTitle
  getData()
}, { immediate: true, deep: true })

onMounted(() => {
  init()
})
</script>

<style lang="scss">
.helplist {
  width: 100%;
  margin: 0 auto;
  padding: 80px 20%;
  background: #FFF;
  color: #000;
  min-height: 800px;

  .container {
    > h1 {
      text-align: center;
      margin: 30px 0 20px 0;
    }

    .page {
      text-align: right;
      margin-top: 10px;

      .el-pagination {
        background: transparent;

        .btn-prev,
        .btn-next {
          background-color: transparent;
          color: #000;
          border: none;
        }

        .el-pager li {
          background-color: transparent;
          color: #000;
          border: none;

          &.is-active {
            color: #f0a70a;
          }
        }
      }
    }
  }
}

.list {
  font-size: 16px;

  .item {
    color: #464646;
    display: block;
    line-height: 40px;
    border-bottom: 1px solid #ebebeb;
    cursor: pointer;

    .iconimg {
      width: 14px;
      vertical-align: sub;
      margin-left: 20px;
    }

    .time {
      float: right;
      color: #999;
      font-size: 14px;
    }

    &:hover {
      color: #f0a70a;
    }
  }
}

.route-wrap {
  font-size: 14px;

  a {
    color: #f1ab15;
  }

  span {
    color: #f1ab15;
  }
}

@media screen and (max-width: 768px) {
  .helplist {
    width: 100%;
    margin: 0 auto;
    padding: 80px 5%;
    background: #FFF;
    color: #000;

    .time {
      display: none;
    }
  }
}
</style>

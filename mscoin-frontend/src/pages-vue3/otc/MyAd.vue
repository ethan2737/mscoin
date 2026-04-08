<template>
  <div class="nav-rights">
    <div class="my_ad_box">
      <div class="add_ad">
        <el-button type="warning" @click="publish">
          <el-icon><Plus /></el-icon>
          {{$t('otc.myad.post')}}
        </el-button>
      </div>
      <el-alert type="info" :closable="false">{{$t('otc.myad.alert')}}</el-alert>
      <div class="order-table">
        <el-table :data="tableAdv" v-loading="loading" border>
          <el-table-column :label="$t('otc.myad.no')" prop="id" width="55" align="center" />
          <el-table-column :label="$t('otc.myad.type')" width="90" align="center">
            <template #default="{ row }">
              {{ row.advertiseType == 0 ? $t('otc.myad.buy') : $t('otc.myad.sell') }}
            </template>
          </el-table-column>
          <el-table-column :label="$t('otc.myad.limit')" width="100" align="center">
            <template #default="{ row }">
              {{ row.minLimit }} ~ {{ row.maxLimit }}
            </template>
          </el-table-column>
          <el-table-column :label="$t('otc.myad.remain')" prop="remainAmount" width="90" align="center" />
          <el-table-column :label="$t('otc.myad.coin')" prop="coinUnit" width="100" align="center" />
          <el-table-column :label="$t('otc.myad.created')" prop="createTime" width="160" align="center" />
          <el-table-column :label="$t('otc.myad.operate')" width="180" align="center">
            <template #default="{ row, index }">
              <el-button
                size="small"
                style="margin-right: 8px;"
                @click="handleUpdate(row)"
              >
                {{$t('otc.myad.update')}}
              </el-button>
              <el-button
                size="small"
                style="margin-right: 8px;"
                @click="handleShelf(row)"
              >
                {{ row.status == 0 ? $t('otc.myad.dropoff') : $t('otc.myad.shelf') }}
              </el-button>
              <el-button
                size="small"
                type="danger"
                @click="handleDelete(row, index)"
              >
                {{$t('otc.myad.delete')}}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <div style="margin: 10px; overflow: hidden" id="pages">
          <div style="float: right;">
            <el-pagination
              v-if="totalElement > 0"
              layout="prev, pager, next"
              :page-size="pageNumber"
              :total="totalElement"
              :current-page="currentPage"
              @current-change="changePage"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 我的广告页面
 */
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { ElMessage, ElMessageBox, ElButton, ElTable, ElTableColumn, ElPagination, ElAlert, ElIcon } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import axios from 'axios'
import { useRouter, useStore } from 'vue-router/composables'

const router = useRouter()
const store = useStore()

const host = 'http://localhost'

const loginmsg = ref('请先登录')
const dataCount = ref(0)
const tableAdv = ref([])
const loading = ref(true)
const totalElement = ref(0)
const pageNumber = ref(10)
const currentPage = ref(1)

const updateLangData = () => {
  // 语言切换动时更新表格列标题（Element Plus 自动响应式）
}

const publish = () => {
  router.push('/uc/ad/create')
}

const handleUpdate = (row) => {
  if (row.status == 0) {
    ElMessage.error('广告已下架，无法编辑')
  } else {
    router.push('/uc/ad/update?id=' + row.id)
  }
}

const handleShelf = (row) => {
  const canshu = { id: row.id }
  const url = row.status == 0
    ? `${host}/otc/advertise/on/shelves`
    : `${host}/otc/advertise/off/shelves`

  axios.post(url, canshu, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      ElMessage.success(resp.message)
      getAd()
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
}

const handleDelete = (row, index) => {
  if (row.status == 1) {
    ElMessageBox.confirm('确定要删除该广告吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      const canshu = { id: row.id }
      axios.post(`${host}/otc/advertise/delete`, canshu, {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json',
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      }).then(response => {
        const resp = response.data
        if (resp.code == 0) {
          ElMessage.success(resp.message)
          tableAdv.value.splice(index, 1)
        } else {
          ElMessage.error(resp.message)
        }
      }).catch(() => {
        ElMessage.error('请求失败')
      })
    }).catch(() => {
      // 用户取消
    })
  } else {
    ElMessage.error('下架广告后才可以删除！')
  }
}

const changePage = (page) => {
  currentPage.value = page
  // 如需分页加载数据，在此处添加逻辑
}

const getAd = () => {
  axios.post(`${host}/otc/advertise/all`, {}, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      tableAdv.value = resp.data.content || []
      for (let i = 0; i < tableAdv.value.length; i++) {
        if (tableAdv.value[i].coin) {
          tableAdv.value[i].coinUnit = tableAdv.value[i].coin.unit
        }
      }
      loading.value = false
      totalElement.value = resp.data.totalElements || 0
    } else {
      ElMessage.error(loginmsg.value)
      loading.value = false
    }
  }).catch(() => {
    ElMessage.error(loginmsg.value)
    loading.value = false
  })
}

// Watch for lang changes
const lang = computed(() => store.state.lang)
watch(lang, () => {
  updateLangData()
})

onMounted(() => {
  getAd()
})
</script>

<style scoped lang="scss">
.nav-rights {
  padding: 0 0 0 20px;

  .my_ad_box {
    .add_ad {
      margin-bottom: 20px;

      .el-button {
        background: #f0a70a;
        color: #fff;
        border-color: #f0a70a;

        &:hover {
          border-color: #f0a70a;
          opacity: 0.9;
        }
      }
    }

    .el-alert.el-alert--info {
      border: none;
      background-color: #192330;
      text-align: center;
    }
  }
}

.order-table {
  :deep(.el-table) {
    background: transparent;

    th.el-table__cell {
      background: #27313e;
      color: #fff;
    }

    td.el-table__cell {
      background: transparent;
      color: #fff;
    }

    .el-button--small {
      border-radius: 10px;

      &.el-button--default {
        border: 1px solid #00b275;
        background-color: transparent;
        color: #00b275;
      }

      &.el-button--warning {
        border: 1px solid #f0ac19;
        background-color: transparent;
        color: #f0ac19;
      }

      &.el-button--danger {
        border: 1px solid #f15057;
        background-color: transparent;
        color: #f15057;
      }
    }
  }
}
</style>

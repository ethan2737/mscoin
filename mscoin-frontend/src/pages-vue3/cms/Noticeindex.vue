<template>
  <div class="nav-right notice">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear">
        <section class="trade-group merchant-top">
          <i class="merchant-icon tips"></i>
          <span class="tips-word">平台公告</span>
        </section>
        <section class="noticeBox">
          <el-table :data="tableData1" :show-header="false">
            <el-table-column prop="name" label=" ">
              <template #default="{ row }">
                <router-link :to="`/notice/detail?id=${row.id}`">{{ row.name }}</router-link>
              </template>
            </el-table-column>
            <el-table-column prop="payNums" label=" " align="right" />
          </el-table>
          <div style="margin: 10px;overflow: hidden">
            <div style="float: right;">
              <el-pagination
                :total="total"
                :current-page="currentPage"
                layout="prev, pager, next"
                @current-change="changePage"
              />
            </div>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - CMS 公告索引页面
 */
import { ref, onMounted } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElPagination } from 'element-plus'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const host = ''

const sasa = ref(false)
const tableData1 = ref([])
const total = ref(100)
const currentPage = ref(1)

const mockTableData1 = (index) => {
  const data = []
  const page = index || 1
  axios.post(`${host}/otc/advertise/page`, { pageNo: page }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code == 0) {
      for (let i = 0; i < 10; i++) {
        if (resp.data.content[i]) {
          data.push({
            name: resp.data.content[i].id,
            payNums: '2018-01-11 15:36:07'
          })
        }
      }
      tableData1.value = data
    } else if (resp.code == 4000) {
      ElMessage.success('请先登录')
      router.push('/login')
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    ElMessage.error('请求失败')
  })
  return data
}

const changePage = (index) => {
  mockTableData1(index)
}

onMounted(() => {
  mockTableData1(1)
})
</script>

<style>
.notice .el-table:after {
  width: 0;
}

.notice .el-table:before {
  height: 0;
  background: none !important;
}

.notice .noticeBox .el-table-wrapper {
  border: 0;
}

.notice .el-table {
  background: transparent;
}

.notice .el-table .el-table__header-wrapper {
  display: none;
}

.notice .el-table .el-table__body tr {
  background: transparent;
}

.notice .el-table .el-table__body td {
  background: transparent;
  color: #fff;
  border-color: #27313e;
}

.notice .trade-group {
  height: 50px;
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-align: center;
  -ms-flex-align: center;
  align-items: center;
  background: #fff;
  padding: 0 15px;
  color: #fff;
  margin-bottom: 20px;
  font-size: 14px;
}

.notice .merchant-icon.tips {
  display: inline-block;
  margin-left: 4px;
  width: 4px;
  height: 22px;
  margin-right: 10px;
  background: #f0a70a;
}

.notice .merchant-top .tips-word {
  -webkit-box-flex: 2;
  -ms-flex-positive: 2;
  flex-grow: 2;
  text-align: left;
}

.notice .bill_box {
  width: 100%;
  height: auto;
  overflow: hidden;
}

.notice .rightarea {
  background: #fff;
  padding-left: 15px !important;
  padding-right: 15px !important;
  margin-bottom: 60px !important;
  padding-top: 20px;
}

.notice .nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 15px;
}

.notice .order_box {
  width: 100%;
  background: #fff;
  height: 56px;
  line-height: 56px;
  margin-bottom: 20px;
  border-bottom: 2px solid #ccf2ff;
  position: relative;
  text-align: left;
}

.notice .order_box a {
  color: #0B0D1B;
  font-size: 16px;
  padding: 0 30px;
  cursor: pointer;
  text-decoration: none;
  text-align: center;
  line-height: 54px;
  display: inline-block;
}

.notice .order_box .active {
  border-bottom: 2px solid #00b5f6;
}

.notice .order_box .search {
  position: absolute;
  width: 300px;
  height: 32px;
  top: 12px;
  right: 0;
  display: flex;
}

.notice .el-pagination {
  background: transparent;
}

.notice .el-pagination .btn-prev,
.notice .el-pagination .btn-next {
  background: transparent;
  color: #fff;
  border: none;
}

.notice .el-pagination .el-pager li {
  background: transparent;
  color: #fff;
  border: none;
}

.notice .el-pagination .el-pager li.is-active {
  color: #f0a70a;
}
</style>

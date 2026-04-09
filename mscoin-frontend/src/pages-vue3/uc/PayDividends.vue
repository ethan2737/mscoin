<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear record">
        <div class="col-xs-12 rightarea-con">
          <div class="trade_accumulative">
            <div class="trade_accumulative_return">
              {{ $t('uc.finance.paydividende.money_holding') }}&nbsp;&nbsp;{{ Number(accumulative_return).toFixed(8) }}
            </div>
          </div>
          <div class="order-table">
            <el-table :data="tableRecord" :no-data-text="$t('common.nodata')" v-loading="loading" border>
              <el-table-column prop="haveTime" :label="$t('uc.finance.paydividende.datehodld')" align="center" />
              <el-table-column :label="$t('uc.finance.paydividende.paydividends')" align="center">
                <template #default="{ row }">
                  <span :title="row.memBouns">
                    {{ toFloor(row.memBouns) }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="arriveTime" :label="$t('uc.finance.paydividende.account_date')" align="center" />
            </el-table>
            <div style="margin: 10px; overflow: hidden">
              <div style="float: right;">
                <el-pagination
                  layout="total, prev, pager, next"
                  :total="total"
                  :page-size="pageSize"
                  :current-page="page + 1"
                  @current-change="changePage"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 分红页面
 */
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const store = useStore()
const router = useRouter()

const host = ''
const api = {
  uc: {
    paydividends: '/uc/asset/dividend/list'
  }
}

const loading = ref(false)
const accumulative_return = ref('0')
const accumulat_return = ref('0')
const pageSize = ref(10)
const page = ref(0)
const total = ref(0)
const tableRecord = ref([])
const memberId = ref(0)

const toFloor = (number, scale = 8) => {
  if (Number(number) === 0) return 0
  const str = number + ''
  if (str.indexOf('e') > -1 || str.indexOf('E') > -1) {
    const num = Number(number).toFixed(scale + 1)
    return (num + '').substring(0, (num + '').length - 1)
  } else if (str.indexOf('.') > -1) {
    if (scale === 0) return str.substring(0, str.indexOf('.'))
    return str.substring(0, str.indexOf('.') + scale + 1)
  }
  return str
}

const init = () => {
  if (!store.getters.isLogin) {
    router.push('/login')
    return
  }
  memberId.value = store.getters.member.id
  if (memberId.value && memberId.value !== 0) {
    getTableData()
  }
}

const getTableData = (pageNo = 1, size = 10) => {
  loading.value = true
  axios.post(`${host}${api.uc.paydividends}`, {
    memberId: memberId.value,
    pageNo,
    pageSize: size
  }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    loading.value = false
    const resp = response.data
    if (resp.code === 0) {
      const data = resp.data
      accumulative_return.value = data.amount || 0
      total.value = parseInt(resp.totalElement)
      tableRecord.value = data.bonus
    } else {
      ElMessage.error(resp.message)
    }
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const changePage = (pageIndex) => {
  getTableData(pageIndex)
}

onMounted(() => {
  init()
})
</script>

<style scoped lang="scss">
.bill_box {
  width: 100%;
  height: auto;
  overflow: hidden;
}

.rightarea {
  background: #fff;
  padding-left: 15px !important;
  padding-right: 15px !important;
  margin-bottom: 60px !important;

  .rightarea-top {
    line-height: 75px;
    border-bottom: #f1f1f1 solid 1px;
  }

  .rightarea-con {
    padding-top: 30px;
    padding-bottom: 20px;
  }
}

.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 15px;
}

.order-table {
  margin-top: 20px;
}

.trade_accumulative {
  height: auto;
  overflow: hidden;
  font-size: 18px;
  font-weight: 600;
  text-align: left;
  border-bottom: 1px solid #f5f5f5;
  padding-bottom: 20px;
  margin-bottom: 20px;

  .trade_accumulative_return {
    width: 50%;
    float: left;
  }

  .trade_accumulat_return {
    width: 50%;
    float: left;
  }
}
</style>

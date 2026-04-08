<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear">
        <section class="trade-group merchant-top">
          <i class="merchant-icon tips"></i>
          <span class="tips-word">{{ $t('uc.mining.title') }}</span>
        </section>
        <div class="shaow">
          <div class="money_table">
            <el-table :columns="tableColumns" :data="orderList" v-loading="loading" border>
              <el-table-column prop="activityName" :label="$t('uc.activity.column1')" align="center" width="300">
                <template #default="{ row }">
                  <a :href="rootHost + '/lab/detail/' + row.activityId" target="_blank" style="color: #f0a70a;">
                    {{ row.activityName }}
                  </a>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.activity.column2')" align="center">
                <template #default="{ row }">
                  {{ formatType(row.type) }}
                </template>
              </el-table-column>
              <el-table-column prop="amount" :label="$t('uc.activity.column3')" align="center" />
              <el-table-column prop="baseSymbol" :label="$t('uc.activity.column4')" align="center" />
              <el-table-column prop="coinSymbol" :label="$t('uc.activity.column5')" align="center" />
              <el-table-column :label="$t('uc.activity.column6')" align="center">
                <template #default="{ row }">
                  <span :style="{ color: getStatusColor(row.type, row.state) }">
                    {{ formatStatus(row.type, row.state) }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('uc.activity.column7')" align="center">
                <template #default="{ row }">
                  {{ row.turnover }} {{ row.baseSymbol }}
                </template>
              </el-table-column>
              <el-table-column prop="createTime" :label="$t('uc.activity.column8')" align="center" width="140" />
            </el-table>
            <div class="page">
              <el-pagination
                layout="prev, pager, next"
                :total="total"
                :page-size="pageSize"
                :current-page="pageNo"
                @current-change="loadDataPage"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 创新订单页面
 */
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElTable, ElTableColumn, ElPagination } from 'element-plus'
import axios from 'axios'
import { useStore } from 'vue-router/composables'

const store = useStore()

const host = 'http://localhost'
const rootHost = 'http://localhost'
const api = {
  uc: {
    myInnovationOrderList: '/uc/activity/myOrders'
  }
}

const loginmsg = ref('登录提示')
const total = ref(0)
const pageSize = ref(10)
const loading = ref(true)
const pageNo = ref(1)
const orderList = ref([])

const formatType = (type) => {
  const typeMap = {
    1: '首发抢购',
    2: '首发分摊',
    3: '持仓瓜分',
    4: '自由认购',
    5: '云矿机认购'
  }
  return typeMap[type] || '未知'
}

const formatStatus = (type, state) => {
  if (type === 5) {
    const statusMap = {
      1: '未部署',
      2: '已部署',
      3: '已撤销'
    }
    return statusMap[state] || '临时'
  } else {
    const statusMap = {
      1: '待成交',
      2: '已成交',
      3: '已撤销'
    }
    return statusMap[state] || '临时'
  }
}

const getStatusColor = (type, state) => {
  if (state === 1) return '#F3BB01'
  if (state === 2) return '#05DF3F'
  if (state === 3) return '#E80000'
  return '#A7A5A1'
}

const getMyOrderList = () => {
  const params = {
    pageNo: pageNo.value,
    pageSize: pageSize.value
  }

  axios.post(`${host}${api.uc.myInnovationOrderList}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      orderList.value = resp.data.content
    } else {
      ElMessage.error(loginmsg.value)
    }
    loading.value = false
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const loadDataPage = (data) => {
  pageNo.value = data
  getMyOrderList()
}

onMounted(() => {
  getMyOrderList()
})

// 表格列定义使用 computed 以支持国际化
const tableColumns = computed(() => {
  return [
    {
      label: '活动名称',
      key: 'activityName',
      align: 'center',
      width: 300
    },
    {
      label: '类型',
      key: 'type',
      align: 'center'
    },
    {
      label: '数量',
      key: 'amount',
      align: 'center'
    },
    {
      label: '基础币种',
      key: 'baseSymbol',
      align: 'center'
    },
    {
      label: '挖矿币种',
      key: 'coinSymbol',
      align: 'center'
    },
    {
      label: '状态',
      key: 'state',
      align: 'center'
    },
    {
      label: '成交额',
      key: 'turnover',
      align: 'center'
    },
    {
      label: '创建时间',
      key: 'createTime',
      align: 'center',
      width: 140
    }
  ]
})
</script>

<style lang="scss">
.nav-right {
  .rightarea.bill_box {
    .shaow {
      padding: 5px;
    }
    .money_table {
      .search {
        width: 200px;
        margin-bottom: 10px;
      }
      .el-table {
        background: #27313e;
        color: #fff;

        th.el-table__cell {
          background: #27313e;
          color: #fff;
        }

        td.el-table__cell {
          background: transparent;
          color: #fff;
          padding: 10px 10px;

          .el-button {
            background: transparent;
            height: 25px;
            padding: 0;
            border-radius: 0;

            span {
              display: inline-block;
              line-height: 20px;
              font-size: 12px;
              padding: 0 15px;
              letter-spacing: 1px;
            }
          }

          .el-button--info {
            border: 1px solid #f0ac19;
            span {
              color: #f0ac19;
            }
          }

          .el-button--danger {
            border: 1px solid #f15057;
            span {
              color: #f15057;
            }
          }

          .el-button--primary {
            border: 1px solid #00b275;
            span {
              color: #00b275;
            }
          }

          .el-button--default {
            border: 1px solid #2c384f;
            background: #222c3e;
            span {
              color: #54637a;
            }
          }
        }
      }
    }
  }
}
</style>

<style scoped lang="scss">
.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 0 0 15px;

  .rightarea.bill_box {
    padding-left: 15px;
    width: 100%;
    height: auto;
    overflow: hidden;
  }
}

.header-btn {
  float: right;
  padding: 5px 15px;
  border: 1px solid #f0ac19;
  color: #f0ac19;
  margin-left: 20px;

  &:hover {
    background: #f0ac19;
    color: #000;
    cursor: pointer;
  }
}

.page {
  text-align: right;
  margin-top: 20px;
}
</style>

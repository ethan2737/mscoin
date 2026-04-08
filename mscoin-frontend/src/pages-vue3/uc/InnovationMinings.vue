<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear">
        <section class="trade-group merchant-top">
          <i class="merchant-icon tips"></i>
          <span class="tips-word">{{ $t('uc.mining.title') }}</span>
        </section>
        <div class="shaow">
          <div class="money_table mining-list">
            <div v-if="miningList.length === 0" style="text-align: center; margin-top: 30px; color: #fff;">
              {{ $t('uc.mining.empty') }}
            </div>
            <div
              v-for="item in miningList"
              :key="item.activityId"
              style="width: 94%; margin: 5px auto;"
            >
              <el-card style="position: relative; background: #2f3e59; border: none;">
                <div style="width: 100%; min-height: 58px; display: flex; align-items: center;">
                  <img
                    :src="item.image"
                    style="width: 50px; height: 50px; border-radius: 50%;"
                  />
                  <div style="text-align: left; margin-left: 15px; flex: 1;">
                    <h2 style="color: #fff; font-size: 16px; margin: 0;">{{ item.title }}</h2>
                    <p>
                      <span
                        v-if="item.miningStatus === 0"
                        style="font-size: 12px; padding: 2px 8px; border-radius: 10px; background: #FF0000; color: #fff;"
                      >
                        {{ $t('uc.mining.status0') }}
                      </span>
                      <span
                        v-if="item.miningStatus === 1"
                        style="font-size: 12px; padding: 2px 8px; border-radius: 10px; background: #00b275; color: #fff;"
                      >
                        {{ $t('uc.mining.status1') }}
                      </span>
                      <span
                        v-if="item.miningStatus === 2"
                        style="font-size: 12px; padding: 2px 8px; border-radius: 10px; background: #888; color: #fff;"
                      >
                        {{ $t('uc.mining.status2') }}
                      </span>
                    </p>
                  </div>
                  <el-button
                    v-if="item.canFetchProfit > 0"
                    type="warning"
                    @click="fetchProfit(item)"
                    style="box-shadow: 0px 0px 4px #f29100;"
                  >
                    {{ $t('uc.mining.pending') }}{{ item.canFetchProfit }}
                  </el-button>
                </div>
                <div style="width: 100%; padding: 5px 10px; background: #000; border-radius: 5px; margin-top: 10px;">
                  <table class="config-table" style="width: 100%;">
                    <tr>
                      <td style="color: #828ea1; font-size: 12px; text-align: left;">{{ $t('uc.mining.miningUnit') }}</td>
                      <td style="color: #EEE; font-size: 12px; text-align: right; padding-right: 10px;">{{ item.miningUnit }}</td>
                      <td style="color: #828ea1; font-size: 12px; text-align: left; padding-left: 10px;">{{ $t('uc.mining.miningProfit') }}</td>
                      <td style="color: #EEE; font-size: 12px; text-align: right;">{{ item.totalProfit }}</td>
                    </tr>
                    <tr>
                      <td style="color: #828ea1; font-size: 12px; text-align: left;">{{ $t('uc.mining.miningDays') }}</td>
                      <td style="color: #EEE; font-size: 12px; text-align: right; padding-right: 10px;">
                        {{ item.miningTimes }}<span style="color: #828ea1;">{{ $t('uc.mining.times') }}</span>
                      </td>
                      <td style="color: #828ea1; font-size: 12px; text-align: left; padding-left: 10px;">{{ $t('uc.mining.minedgedDays') }}</td>
                      <td style="color: #EEE; font-size: 12px; text-align: right;">
                        {{ item.minedgedTimes }}<span style="color: #828ea1;">{{ $t('uc.mining.times') }}</span>
                      </td>
                    </tr>
                    <tr>
                      <td style="color: #828ea1; font-size: 12px; text-align: left;">{{ $t('uc.mining.miningDaysProfit') }}</td>
                      <td style="color: #EEE; font-size: 12px; text-align: right; padding-right: 10px;">
                        {{ item.miningDaysprofit }} {{ item.miningUnit }}/{{ item.miningDays }}
                        <span v-if="item.period === 0" style="color: #828ea1;">{{ $t('uc.mining.day') }}</span>
                        <span v-if="item.period === 1" style="color: #828ea1;">{{ $t('uc.mining.week') }}</span>
                        <span v-if="item.period === 2" style="color: #828ea1;">{{ $t('uc.mining.month') }}</span>
                        <span v-if="item.period === 3" style="color: #828ea1;">{{ $t('uc.mining.year') }}</span>
                      </td>
                      <td style="color: #828ea1; font-size: 12px; text-align: left; padding-left: 10px;">{{ $t('uc.mining.miningCurrentDaysProfit') }}</td>
                      <td style="color: #EEE; font-size: 12px; text-align: right;">
                        {{ item.currentDaysprofit }} {{ item.miningUnit }}/{{ item.miningDays }}
                        <span v-if="item.period === 0" style="color: #828ea1;">{{ $t('uc.mining.day') }}</span>
                        <span v-if="item.period === 1" style="color: #828ea1;">{{ $t('uc.mining.week') }}</span>
                        <span v-if="item.period === 2" style="color: #828ea1;">{{ $t('uc.mining.month') }}</span>
                        <span v-if="item.period === 3" style="color: #828ea1;">{{ $t('uc.mining.year') }}</span>
                      </td>
                    </tr>
                  </table>
                </div>
                <div style="font-size: 12px; padding-top: 10px; text-align: left; color: #828ea1;" v-if="item.miningInvite > 0">
                  <el-icon><Info-Filled /></el-icon>
                  {{ $t('uc.mining.invitetip1') }}{{ item.miningInvite * 100 | percentFun }}{{ $t('uc.mining.invitetip2') }}
                  <span v-if="item.miningInvitelimit !== 0">{{ item.miningInvitelimit * 100 | percentFun }}%</span>
                  <span v-if="item.miningInvitelimit === 0">-{{ $t('uc.mining.invitetip3') }}</span>
                </div>
                <div style="font-size: 12px; padding-top: 10px; padding-bottom: 10px; text-align: left; color: #828ea1;">
                  <span>{{ $t('uc.mining.reminder') }}<router-link to="/crowdfunding" style="color: #f0a70a;">{{ $t('uc.mining.btn') }}</router-link></span>
                </div>
                <template #footer>
                  <el-button long type="warning" :to="'/mining/detail/' + item.activityId" style="width: 100%;">
                    {{ $t('activity.invite2') }}
                  </el-button>
                </template>
              </el-card>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 创新挖矿页面
 */
import { ref } from 'vue'
import { ElMessage, ElCard, ElButton, ElIcon } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vue-router/composables'

const store = useStore()

const host = 'http://localhost'
const api = {
  uc: {
    myInnovationMinings: '/uc/activity/myMinings',
    myMiningsFetchProfit: '/uc/activity/fetchProfit'
  }
}

const loginmsg = ref('登录提示')
const total = ref(0)
const pageSize = ref(10)
const loading = ref(true)
const pageNo = ref(1)
const miningList = ref([])

const percentFun = (value) => {
  const tem = value * 100
  return tem.toFixed(0)
}

const fetchProfit = (item) => {
  loading.value = true
  item.canFetchProfit = 0
  const params = { id: item.id }

  axios.post(`${host}${api.uc.myMiningsFetchProfit}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      getMyMiningList()
      ElMessage.success('领取成功')
    } else {
      ElMessage.error(resp.message)
    }
    loading.value = false
  }).catch(() => {
    loading.value = false
    ElMessage.error('请求失败')
  })
}

const getMyMiningList = () => {
  const params = {
    pageNo: pageNo.value,
    pageSize: pageSize.value
  }

  axios.post(`${host}${api.uc.myInnovationMinings}`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      miningList.value = resp.data.content
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
  getMyMiningList()
}

// Expose methods for pagination
defineExpose({
  loadDataPage,
  getMyMiningList
})

// Initialize on mount
getMyMiningList()
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
      .el-card {
        background: #2f3e59;
        border: none !important;

        .el-card__header {
          background: #27313e;
          color: #fff;
        }

        .el-card__body {
          color: #fff;

          .el-button {
            background: transparent;
            height: 25px;
            padding: 0 0px;
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
    padding-right: 15px;
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

.mining-list {
  background: transparent;
}

.config-table {
  width: 100%;

  tr {
    td {
      color: #828ea1;
      font-size: 12px;

      &:nth-child(1) {
        text-align: left;
      }

      &:nth-child(2) {
        text-align: right;
        padding-right: 10px;
        color: #EEE;
      }

      &:nth-child(3) {
        text-align: left;
        padding-left: 10px;
      }

      &:nth-child(4) {
        text-align: right;
        color: #EEE;
      }
    }
  }
}

.merchant-top {
  height: 50px;
  display: flex;
  align-items: center;
  padding: 0 15px;
  margin-bottom: 20px;
  font-size: 14px;
}

.merchant-icon {
  display: inline-block;
  margin-left: 4px;
  background-size: 100% 100%;
}

.merchant-top .tips-word {
  flex-grow: 2;
  text-align: left;
}

.merchant-icon.tips {
  width: 4px;
  height: 22px;
  margin-right: 10px;
  background: #f0a70a;
}

.bill_box {
  width: 100%;
  height: auto;
  overflow: hidden;
}

.shoujiBtn {
  width: 100%;
  height: 95px;
}

@media screen and (max-width: 1200px) {
  .shouji {
    font-size: 15px;
  }
  .shoujiBtn {
    width: 100%;
    height: auto;
  }
}
</style>

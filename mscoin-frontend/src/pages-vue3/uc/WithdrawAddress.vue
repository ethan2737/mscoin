<template>
  <div class="nav-rights">
    <div class="nav-right">
      <div class="bill_box_address">
        <section class="trade-group merchant-tops">
          <h1 class="tips-word1">{{ $t('uc.finance.withdraw.addressmanager') }}</h1>
        </section>
        <section>
          <div class="table-inner">
            <div class="action-inner">
              <div class="inner-left">
                <p class="describe">{{ $t('uc.finance.withdraw.symbol') }}</p>
                <el-select v-model="coinType" style="width: 100px; margin-top: 10px;" size="large" @change="refresh">
                  <el-option v-for="item in coinList" :key="item" :value="item">{{ item }}</el-option>
                </el-select>
              </div>
              <div class="inner-box deposit-address mt25">
                <p class="describe">{{ $t('uc.finance.withdraw.address') }}</p>
                <el-input v-model="withdrawAddr" style="width: 90%; margin-top: 10px;" size="large" />
              </div>
              <div class="mt25">
                <p class="describe">{{ $t('uc.finance.withdraw.remark') }} / Memo</p>
                <el-input v-model="remark" style="width: 100%; margin-top: 10px;" size="large" />
              </div>
            </div>
            <div class="btnbox">
              <el-button type="warning" size="large" @click="addAddr">
                {{ $t('uc.finance.withdraw.add') }}
              </el-button>
            </div>
            <div class="action-content">
              <div class="action-body">
                <p class="acb-p1 describe">{{ $t('uc.finance.withdraw.addresslist') }}</p>
                <div class="order-table">
                  <el-table :data="dataRecharge" border style="width: 100%">
                    <el-table-column prop="unit" :label="$t('uc.finance.withdraw.symbol')" align="center" />
                    <el-table-column prop="address" :label="$t('uc.finance.withdraw.address')" align="center" />
                    <el-table-column prop="remark" :label="$t('uc.finance.withdraw.remark')" align="center" />
                    <el-table-column :label="$t('uc.finance.withdraw.operate')" align="center" width="150">
                      <template #default="{ row }">
                        <el-button
                          type="danger"
                          size="small"
                          style="border-radius: 10px;"
                          @click="delAddr(row.id)"
                        >
                          {{ $t('uc.finance.withdraw.delete') }}
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                  <div class="pagination-container">
                    <el-pagination
                      :total="dataCount"
                      :current-page="currentPage"
                      :page-size="10"
                      @current-change="changePage"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>

    <!-- 验证对话框 -->
    <el-dialog v-model="modal2" :title="$t('uc.finance.withdraw.safevalidate')" width="360px">
      <el-form ref="formValidateAddr" :model="formValidateAddr" :rules="ruleValidate" label-width="85px">
        <el-form-item
          v-if="validPhone"
          :label="$t('uc.finance.withdraw.telno')"
          prop="mobileNo"
        >
          <el-input v-model="formValidateAddr.mobileNo" size="large" disabled />
        </el-form-item>
        <el-form-item
          v-if="validPhone"
          :label="$t('uc.finance.withdraw.smscode')"
          prop="vailCode2"
        >
          <el-input v-model="formValidateAddr.vailCode2" size="large">
            <template #append>
              <el-button @click="send(2)" :disabled="disbtn">
                <span v-if="sendMsgDisabled2">{{ time2 }}{{ $t('uc.finance.withdraw.second') }}</span>
                <span v-else>{{ $t('uc.finance.withdraw.clickget') }}</span>
              </el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item
          v-if="validEmail"
          :label="$t('uc.finance.withdraw.email')"
          prop="email"
        >
          <el-input v-model="formValidateAddr.email" size="large" disabled />
        </el-form-item>
        <el-form-item
          v-if="validEmail"
          :label="$t('uc.finance.withdraw.emailcode')"
          prop="vailCode1"
        >
          <el-input v-model="formValidateAddr.vailCode1" size="large">
            <template #append>
              <el-button @click="send(1)" :disabled="disbtn">
                <span v-if="sendMsgDisabled1">{{ time1 }}{{ $t('uc.finance.withdraw.second') }}</span>
                <span v-else>{{ $t('uc.finance.withdraw.clickget') }}</span>
              </el-button>
            </template>
          </el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button type="primary" size="large" style="width: 100%;" @click="handleSubmit('formValidateAddr')">
          {{ $t('uc.finance.withdraw.save') }}
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 提现地址管理页面
 */
import { ref, reactive, inject, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'

const router = inject('router')

const host = ''

const modal2 = ref(false)
const disbtn = ref(false)
const sendMsgDisabled1 = ref(false)
const sendMsgDisabled2 = ref(false)
const time1 = ref(60)
const time2 = ref(60)
const withdrawAddr = ref('')
const remark = ref('')
const coinType = ref('')
const validEmail = ref(false)
const validPhone = ref(false)
const coinList = ref([])
const dataRecharge = ref([])
const dataCount = ref(0)
const currentPage = ref(1)
const loading = ref(true)

const formValidateAddr = reactive({
  mobileNo: '',
  vailCode2: '',
  email: '',
  vailCode1: ''
})

const ruleValidate = reactive({
  mobileNo: [
    { required: true, message: $t('uc.finance.withdraw.telerr'), trigger: 'blur' }
  ],
  vailCode2: [
    { required: true, message: $t('uc.finance.withdraw.codeerr'), trigger: 'change' }
  ],
  email: [
    { required: true, type: 'email', message: $t('uc.finance.withdraw.emailerr'), trigger: 'blur' }
  ],
  vailCode1: [
    { required: true, message: $t('uc.finance.withdraw.codeerr'), trigger: 'change' }
  ]
})

let intervalTimer = null

const refresh = () => {
  withdrawAddr.value = ''
  remark.value = ''
  getList(0, 10)
}

const getMember = () => {
  axios.post(`${host}/uc/approve/security/setting`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      if (resp.data.mobilePhone) {
        formValidateAddr.mobileNo = resp.data.mobilePhone
        validPhone.value = true
        validEmail.value = false
      } else {
        formValidateAddr.email = resp.data.email
        validPhone.value = false
        validEmail.value = true
      }
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
  })
}

const getCoin = () => {
  axios.post(`${host}/uc/withdraw/support/coin`, {}, {
    withCredentials: true,
    headers: {
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      for (let i = 0; i < resp.data.length; i++) {
        coinList.value.push(resp.data[i])
      }
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
  })
}

const getList = (pageNo, pageSize) => {
  axios.post(`${host}/uc/withdraw/address/page`, { pageNo, pageSize }, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0 && resp.data.content) {
      dataRecharge.value = resp.data.content
      dataCount.value = resp.data.totalElement
    } else {
      ElMessage.error(resp.message)
    }
    loading.value = false
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
    loading.value = false
  })
}

const send = (index) => {
  disbtn.value = true
  if (index === 1) {
    if (formValidateAddr.email) {
      axios.post(`${host}/uc/add/address/code`, {}, {
        withCredentials: true,
        headers: {
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      })
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success(resp.message)
          sendMsgDisabled1.value = true
          intervalTimer = setInterval(() => {
            if (time1.value-- <= 0) {
              time1.value = 60
              sendMsgDisabled1.value = false
              clearInterval(intervalTimer)
              disbtn.value = false
            }
          }, 1000)
        } else {
          ElMessage.error(resp.message)
          disbtn.value = false
        }
      })
      .catch(() => {
        ElMessage.error($t('common.logintip'))
        disbtn.value = false
      })
    } else {
      ElMessage.error($t('uc.finance.withdraw.emailerr'))
      disbtn.value = false
    }
  } else if (index === 2) {
    if (formValidateAddr.mobileNo) {
      axios.post(`${host}/uc/mobile/add/address/code`, {}, {
        withCredentials: true,
        headers: {
          'x-auth-token': localStorage.getItem('TOKEN')
        }
      })
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElMessage.success(resp.message)
          sendMsgDisabled2.value = true
          intervalTimer = setInterval(() => {
            if (time2.value-- <= 0) {
              time2.value = 60
              sendMsgDisabled2.value = false
              clearInterval(intervalTimer)
              disbtn.value = false
            }
          }, 1000)
        } else {
          ElMessage.error(resp.message)
          disbtn.value = false
        }
      })
      .catch(() => {
        ElMessage.error($t('common.logintip'))
        disbtn.value = false
      })
    } else {
      ElMessage.error($t('uc.finance.withdraw.telerr'))
      disbtn.value = false
    }
  }
}

const addAddr = () => {
  if (!coinType.value) {
    ElMessage.warning($t('uc.finance.withdraw.symboltip'))
  } else if (!withdrawAddr.value) {
    ElMessage.warning($t('uc.finance.withdraw.addresstip'))
  } else if (!remark.value) {
    ElMessage.warning($t('uc.finance.withdraw.remarktip'))
  } else {
    modal2.value = true
  }
}

const changePage = (index) => {
  currentPage.value = index
  getList(index, 10)
}

const delAddr = (id) => {
  ElMessageBox.confirm($t('common.delete'), $t('common.tip'), {
    confirmButtonText: $t('common.confirm'),
    cancelButtonText: $t('common.cancel'),
    type: 'warning'
  })
  .then(() => {
    axios.post(`${host}/uc/withdraw/address/delete`, { id }, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json',
        'x-auth-token': localStorage.getItem('TOKEN')
      }
    })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success(resp.message)
        refresh()
      } else {
        ElMessage.error(resp.message)
      }
    })
    .catch(() => {
      ElMessage.error($t('common.logintip'))
    })
  })
  .catch(() => {})
}

const handleSubmit = (formName) => {
  submit(formName)
}

const submit = (formName) => {
  const param = {
    address: withdrawAddr.value,
    unit: coinType.value,
    remark: remark.value
  }

  if (validPhone.value) {
    param.aims = formValidateAddr.mobileNo
    param.code = formValidateAddr.vailCode2
  } else {
    param.aims = formValidateAddr.email
    param.code = formValidateAddr.vailCode1
  }

  axios.post(`${host}/uc/withdraw/address/add`, param, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  })
  .then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElMessage.success($t('uc.finance.withdraw.savemsg2'))
      formValidateAddr.vailCode2 = ''
      refresh()
      modal2.value = false
    } else {
      ElMessage.error(resp.message)
    }
  })
  .catch(() => {
    ElMessage.error($t('common.logintip'))
  })
}

onMounted(() => {
  coinType.value = router.currentRoute?.query?.name || ''
  getMember()
  getList(0, 10)
  getCoin()
})
</script>

<style scoped lang="scss">
.nav-rights {
  .nav-right {
    height: auto;
    overflow: hidden;
    padding: 0 15px;

    .bill_box_address {
      section.trade-group.merchant-tops {
        .tips-word1 {
          margin-bottom: 20px;
          text-align: left;
          font-weight: normal;
          margin-left: 30px;
        }
      }

      .table-inner {
        .action-inner {
          display: table;
          padding: 0 30px;
          width: 100%;

          .inner-left {
            display: table-cell;
            width: 15%;
          }

          .inner-box.deposit-address {
            display: table-cell;
            width: 45%;
          }

          .mt25 {
            display: table-cell;
            width: 43%;
          }
        }

        .btnbox {
          text-align: right;
          padding: 25px 30px;
        }

        .action-content {
          padding: 0 30px;

          .acb-p1 {
            font-size: 16px;
            font-weight: 600;
            line-height: 50px;
          }
        }
      }
    }
  }
}

.order-table {
  margin-top: 20px;

  .pagination-container {
    margin: 10px;
    overflow: hidden;
    float: right;
  }
}

p.describe {
  font-size: 16px;
  font-weight: 600;
}
</style>

<template>
  <div class="login_form app-download">
    <div class="login_right">
      <div style="color: #000;margin-bottom: 60px;padding-top: 160px;width: 100%;text-align:center;z-index: 10;">
        <img src="../../assets/images/applogo.png" style="width: 22%;border-radius: 15px;" alt="">
        <p style="font-size:18px;">MSCOIN</p>
        <p style="font-size:12px;margin-top: 10px;color:#888;">最新版本：v{{version}}</p>
        <p style="font-size:10px;margin-top: 10px;color:#888;letter-spacing: 1px;">发布时间：{{publishTime}}</p>
        <p style="font-size:18px;margin-top: 45px;">
            <span style="border: 1px solid #F90; padding: 8px 30px;border-radius: 5px;background-color:#F90;color:#FFF;" @click="downloadClick">
            <el-icon style="font-size: 24px;margin-right: 5px;vertical-align: middle;"><Download /></el-icon>立即下载</span>
        </p>
        <p style="font-size:12px;margin-top: 20px;color:#888;"></p>
      </div>
    </div>

    <div class="section" id="page4">
      <ul>
        <li>
          <div><img src="https://kicks.oss-cn-shanghai.aliyuncs.com/2019/download1.png" alt=""></div>
          <p class="title">极致体验</p>
          <p>精心优化的界面显示，体验流畅的操作响应</p>
        </li>
        <li>
          <div><img src="https://kicks.oss-cn-shanghai.aliyuncs.com/2019/download2.png" alt=""></div>
          <p class="title">币种行情</p>
          <p>支持 MACD、KDJ、RSI、BOLL 等多种专业指标</p>
        </li>
        <li>
          <div><img src="https://kicks.oss-cn-shanghai.aliyuncs.com/2019/download5.png" alt=""></div>
          <p class="title">币币交易</p>
          <p>支持限价委托与市价委托两种方式</p>
        </li>
        <li>
          <div><img src="https://kicks.oss-cn-shanghai.aliyuncs.com/2019/download3.png" alt=""></div>
          <p class="title">法币交易</p>
          <p>优质承兑商，保证资金通道顺畅无阻</p>
        </li>
        <li>
          <div><img src="https://kicks.oss-cn-shanghai.aliyuncs.com/2019/download4.png" alt=""></div>
          <p class="title">资产中心</p>
          <p>随时随地关注资产变化，极速充值/提现</p>
        </li>
      </ul>
    </div>

    <div class="cover" id="cover" @click="coverClick"><img src="https://kicks.oss-cn-shanghai.aliyuncs.com/2019/appdowncover.png" alt=""></div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 应用下载页面
 */
import { ref, computed, onMounted } from 'vue'
import { ElIcon } from 'element-plus'
import { Download } from '@element-plus/icons-vue'
import axios from 'axios'
import { useStore } from 'vuex'

const store = useStore()

const host = ''

const version = ref('1.0.0')
const publishTime = ref('2020/11/05 12:32:00')
const device = ref('ios')

const lang = computed(() => store.state.lang)
const isLogin = computed(() => store.getters.isLogin)

const detect = () => {
  const agent = navigator.userAgent.toLowerCase()
  const android = agent.indexOf('android')
  const iphone = agent.indexOf('iphone')
  const ipad = agent.indexOf('ipad')
  if (android !== -1) {
    return 'android'
  }
  if (iphone !== -1 || ipad !== -1) {
    return 'ios'
  }
  return 'ios'
}

const isWeiXin = () => {
  const ua = window.navigator.userAgent.toLowerCase()
  return ua.match(/MicroMessenger/i) === 'micromessenger'
}

const downloadClick = () => {
  if (isWeiXin()) {
    document.getElementById('cover').style.display = 'block'
  } else {
    location.href = 'https://kicks.oss-cn-shanghai.aliyuncs.com/apk/kickstarterkick.apk'
  }
}

const coverClick = () => {
  document.getElementById('cover').style.display = 'none'
}

const getVersion = () => {
  axios.post(`${host}/uc/ancillary/system/app/version/0`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      version.value = res.data.data.version
      publishTime.value = res.data.data.publishTime
    }
  }).catch(() => {})
}

const init = () => {
  location.href = 'http://www.91guanzhu.com/app/RYQW'
  document.title = (lang.value === '简体中文' ? 'APP 下载 - ' : 'APP Download - ') + 'MSCOIN | 全球比特币交易平台 | 全球数字货币交易平台'
  getVersion()
  device.value = detect()
}

onMounted(() => {
  window.scrollTo(0, 0)
  init()
})
</script>

<style scoped lang="scss">
#page4 {
  background: transparent;
  padding: 20px 0 80px 0;
  ul {
    width: 99%;
    margin: 0 auto;
    li {
      flex: 0 0 25%;
      display: inline-block;
      width: 100%;
      padding: 0 15px;
      div {
        border-radius: 50%;
        vertical-align: middle;
        text-align: center;
        margin: 0 auto;
        img {
          width: 80%;
          margin-top: 28px;
        }
      }
      p {
        font-size: 14px;
        margin: 10px 0;
        text-align: center;
        color: #828ea1;
      }
      p.title {
        color: #000;
        font-size: 18px;
        font-weight: 400;
      }
    }
  }
}
</style>

<style lang="scss">
.app-download {
  background: #f2f6fa !important;
}

.login_form {
  .cover {
    width: 100%;
    height: 100%;
    position: fixed;
    top: 0;
    right: 0;
    background: rgba(0,0,0,0.8);
    z-index: 9999;
    display: none;
  }
  .cover img {
    width: 100%;
  }
  .login_right {
    form.ivu-form.ivu-form-label-right.ivu-form-inline {
      text-align: center;
      .ivu-form-item {
        .ivu-form-item-content {
          .ivu-input-wrapper.ivu-input-type {
            .ivu-input {
              border: none;
              border-bottom: 1px solid #27313e;
              font-size: 14px;
              background: transparent;
              border-radius: 0;
              &:focus {
                border: none;
                border-bottom: 1px solid #27313e;
                box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
              }
            }
          }
        }
      }
      .check-agree {
        .ivu-checkbox-wrapper {
          .ivu-checkbox-input {
            &:focus {
              border: none;
              outline: none;
              box-shadow: 2px 2px 5px transparent, -2px -2px 4px transparent;
            }
          }
        }
        .ivu-checkbox-wrapper.ivu-checkbox-wrapper-checked {
          .ivu-checkbox.ivu-checkbox-checked {
            .ivu-checkbox-inner {
              border: 1px solid #f0ac19;
              background-color: #f0ac19;
            }
          }
        }
        .ivu-checkbox-wrapper.ivu-checkbox-default {
          .ivu-checkbox {
            .ivu-checkbox-inner {
              background: transparent;
            }
          }
        }
      }
    }
  }
}
</style>

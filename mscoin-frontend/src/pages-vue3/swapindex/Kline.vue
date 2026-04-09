<template>
  <div class="container swap" :class="skin">
    <div class="main">
      <div
        style="
          display: flex;
          flex: 100%;
          flex-wrap: wrap;
          justify-content: space-between;
        "
      >
        <div class="center">
          <div class="symbol">
            <div class="item" style="margin-left: 10px">
              <span class="coin">{{ currentCoin.name }}</span>
              <el-popover
                trigger="hover"
                placement="bottom-start"
                width="300"
              >
                <template #reference>
                  <el-icon
                    class="info-icon"
                    style="color: #546886; margin-left: 5px; cursor: pointer;"
                  >
                    <InfoFilled />
                  </el-icon>
                </template>
                <div class="api">
                  <div class="coin-info">{{ coinInfo.information }}</div>
                  <p style="text-align: right; margin-top: -10px">
                    <a :href="coinInfo.infolink" target="_blank">{{
                      $t("swap.moredetail")
                    }}</a>
                  </p>
                </div>
              </el-popover>
            </div>
            <div class="item">
              <span class="text">{{ $t("service.NewPrice") }}</span>
              <span
                class="num"
                :class="{
                  buy: currentCoin.change > 0,
                  sell: currentCoin.change < 0,
                }"
                >{{ toFixed(currentCoin.close, 4) }}</span
              >
              <span class="price-cny"
                >≈ ￥{{ toFixed(currentCoin.usdRate * CNYRate, 2) }}</span
              >
            </div>
            <div class="item">
              <span class="text">{{ $t("service.Change") }}</span>
              <span
                class="num"
                :class="{
                  buy: currentCoin.change > 0,
                  sell: currentCoin.change < 0,
                }"
                >{{ currentCoin.rose }}</span
              >
            </div>
            <div class="item">
              <span class="text">{{ $t("service.high") }}</span>
              <span class="num">{{ toFixed(currentCoin.high, 4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t("service.low") }}</span>
              <span class="num">{{ toFixed(currentCoin.low, 4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t("service.ExchangeNum") }}</span>
              <span class="num"
                >{{ toFixed(currentCoin.volume, 4) }}
                {{ currentCoin.coin }}</span
              >
            </div>
          </div>
          <div class="imgtable">
            <div class="handler">
              <span
                @click="changeImgTable('k')"
                :class="{ active: currentImgTable === 'k' }"
                >{{ $t("swap.kline") }}</span
              >
              <span
                @click="changeImgTable('s')"
                :class="{ active: currentImgTable === 's' }"
                >{{ $t("swap.depth") }}</span
              >
            </div>
            <div
              id="swap_kline_container"
              :class="{ hidden: currentImgTable === 's' }"
            ></div>
            <DepthGraph
              :class="{ hidden: currentImgTable === 'k' }"
              ref="depthGraph"
            ></DepthGraph>
          </div>
        </div>
      </div>
    </div>

    <!-- 弹出框：调整时间 -->
    <el-dialog
      v-model="form.timeModal"
      :title="$t('swap.time')"
      width="400px"
      class="vertical-center-modal"
    >
      <div class="leverage">
        <el-slider
          v-model="form.timeAdjustVal2"
          @input="xxlever2(form.timeAdjustVal2)"
          :marks="times"
          :show-tooltip="false"
          :min="1"
          :max="6"
        ></el-slider>
      </div>
      <template #footer>
        <el-button type="default" size="large" @click="form.timeModal = false">{{
          $t("common.close")
        }}</el-button>
        <el-button type="primary" size="large" @click="adjustTime()">{{
          $t("common.ok")
        }}</el-button>
      </template>
    </el-dialog>

    <!-- 弹出框：调整杠杆 -->
    <el-dialog
      v-model="form.leverageModal"
      :title="$t('swap.modifyLeverage')"
      width="400px"
      class="vertical-center-modal"
    >
      <div class="leverage">
        <el-slider
          v-model="form.leverageAdjustVal2"
          @input="xxlever(form.leverageAdjustVal2)"
          :marks="leverages"
          :show-tooltip="false"
          :min="1"
          :max="5"
        ></el-slider>
      </div>
      <template #footer>
        <el-button
          type="default"
          size="large"
          @click="form.leverageModal = false"
          >{{ $t("common.close") }}</el-button
        >
        <el-button type="primary" size="large" @click="adjustLeverage()">{{
          $t("common.ok")
        }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 合约 K 线页面
 */
import { ref, reactive, computed, onMounted, onBeforeUnmount, inject } from 'vue'
import { ElMessage, ElPopover, ElIcon, ElSlider, ElButton, ElDialog, ElLoading } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import DepthGraph from '../../components/exchange/DepthGraph.vue'
import axios from 'axios'
import $ from 'jquery'
import moment from 'moment'

const store = inject('store')
const router = inject('router')

const host = ''
const api = {
  swap: {
    symbolInfo: '/swap/symbol-info',
    kline: '/market/swap/kline',
    plate: '/market/swap/plate',
    currentPosition: '/uc/swap/position/current',
    currentOrder: '/uc/swap/order/current',
    historyOrder: '/uc/swap/order/history',
    wallet: '/uc/swap/wallet',
    leverage: '/uc/swap/leverage',
    leverageAdjust: '/uc/swap/leverage/adjust',
    timeAdjust: '/uc/swap/time/adjust'
  }
}

const contractId = ref(null)
const skin = ref('night')
const currentImgTable = ref('k')
const CNYRate = ref(null)
const defaultPath = ref(1)
const basecion = ref('usdt')
const coinScale = ref(6)
const baseCoinScale = ref(4)
const takerFee = ref(0.001)
const makerFee = ref(0.001)
const searchKey = ref('')
const coinInfo = ref({})
const fullTrade = ref({})

const currentCoin = reactive({
  id: '',
  base: '',
  coin: '',
  symbol: '',
  perUsdt: 100000,
  close: 0,
  name: '',
  change: 0,
  rose: '0.00%',
  high: 0,
  low: 0,
  volume: 0,
  type: 'ALWAYS',
  coinScale: 6,
  baseCoinScale: 4,
  usdRate: 0
})

const leverages = {
  1: '1x',
  2: '2x',
  3: '5x',
  4: '10x',
  5: '20x'
}

const times = {
  1: '10S',
  2: '30S',
  3: '60S',
  4: '120S',
  5: '360S',
  6: '3000S'
}

const form = reactive({
  timeModal: false,
  leverageModal: false,
  timeAdjustVal2: 1,
  leverageAdjustVal2: 1,
  time: 10,
  leverage: 5,
  patterns: '1',
  type: '2',
  entrustType: 1,
  market: false,
  triggerPrice: null,
  limitPrice: null,
  limitAmount: 1,
  limitPercent: 0
})

const wallet = reactive({
  base: 0.0,
  frozen: 0.0,
  profit: 0.0,
  kick: 0.0
})

const currentPosition = reactive({
  columns: [],
  rows: []
})

const currentOrder = reactive({
  columns: [],
  columns2: [],
  rows: []
})

const historyOrder = reactive({
  columns: [],
  columns2: [],
  rows: []
})

const plate = reactive({
  columns: [],
  askRows: [],
  bidRows: [],
  maxPostion: 10
})

const trade = reactive({
  columns: [],
  rows: []
})

const isLogin = computed(() => store?.getters.isLogin || false)
const lang = computed(() => store?.state.lang || '简体中文')

const positionNum = computed(() => {
  let num = 0
  let buyNum = 0
  let sellNum = 0
  currentPosition.rows.forEach((e) => {
    if (e.direction === 'BUY') {
      buyNum += e.availablePosition
    } else {
      sellNum += e.availablePosition
    }
    num += e.position
  })
  return num
})

const maxOpen = computed(() => {
  const balance = wallet.base
  const perUsdt = currentCoin.perUsdt
  const leverage = form.leverage
  if (balance <= 0 || !perUsdt || !leverage) {
    return 0
  }
  return parseInt((balance * leverage) / perUsdt)
})

let ws = null
let heartbeatTimer = null
let chartWidget = null

const toFixed = (num, digits) => {
  if (num === undefined || num === null) return '0.00'
  return parseFloat(num).toFixed(digits)
}

const toFloor = (num, digits = 8) => {
  if (num === undefined || num === null) return '0.00000000'
  const pow = Math.pow(10, digits)
  return Math.floor(num * pow) / pow
}

const dateFormat = (date, format = 'YYYY-MM-DD HH:mm:ss') => {
  return moment(date).format(format)
}

const timeFormat = (time) => {
  return moment(time).format('HH:mm:ss')
}

const changeImgTable = (str) => {
  currentImgTable.value = str
}

const getCNYRate = () => {
  axios.post(`${host}/market/rate/usd/cny`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      CNYRate.value = res.data.data
    }
  }).catch(() => {})
}

const getCoinInfo = () => {
  axios.post(`${host}/uc/ancillary/contract/info`, {
    coinSymbol: currentCoin.coin
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      coinInfo.value = res.data.data
      currentCoin.name = res.data.data.name
    }
  }).catch(() => {})
}

const getSymbol = () => {
  axios.post(`${host}${api.swap.symbolInfo}`, {
    id: currentCoin.id
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      const resp = res.data.data
      basecion.value = resp.baseSymbol.toLowerCase()
      currentCoin.symbol = resp.coinSymbol + '/' + resp.baseSymbol
      currentCoin.coin = resp.coinSymbol
      currentCoin.type = resp.type
      currentCoin.base = resp.baseSymbol
      currentCoin.coinScale = resp.coinScale
      currentCoin.baseCoinScale = resp.baseCoinScale
      currentCoin.perUsdt = resp.shareNumber
      baseCoinScale.value = resp.baseCoinScale
      coinScale.value = resp.coinScale
      takerFee.value = resp.takerFee
      makerFee.value = resp.makerFee

      store?.commit('setSkin', skin.value)
      initKline()
      startWebsocket()
    }
  }).catch(() => {})
}

const initKline = () => {
  // TradingView K 线图初始化
  const widgetConfig = {
    symbol: currentCoin.symbol,
    interval: '1',
    container_id: 'swap_kline_container',
    width: 800,
    height: 500,
    autosize: true,
    timezone: 'Asia/Shanghai',
    theme: 'dark',
    locale: 'zh_CN',
    library_path: '/charting_library/',
    datafeed: new Datafeeds.UDF({
      baseUrl: host,
    }),
    custom_css_url: '/dark.css',
    disabled_features: [
      'use_localstorage',
      'save_chart_properties_to_local_storage',
    ],
    enabled_features: ['study_templates'],
    studies_overrides: {
      'volume.volume.color.0': '#00b66c',
      'volume.volume.color.1': '#ff5252',
    }
  }

  chartWidget = new window.TradingView.widget(widgetConfig)
}

const startWebsocket = () => {
  const wsUrl = 'wss://api.gateex.io/ws-swap'
  ws = new WebSocket(wsUrl)

  ws.onopen = () => {
    subscribeMarket()
    startHeartbeat()
  }

  ws.onmessage = (event) => {
    const data = JSON.parse(event.data)
    if (data.channel === 'market') {
      updateMarket(data.data)
    } else if (data.channel === 'plate') {
      updatePlate(data.data)
    } else if (data.channel === 'trade') {
      updateTrade(data.data)
    }
  }

  ws.onclose = () => {
    stopHeartbeat()
    setTimeout(() => {
      startWebsocket()
    }, 5000)
  }

  ws.onerror = (error) => {
    console.error('WebSocket error:', error)
  }
}

const subscribeMarket = () => {
  if (ws && ws.readyState === WebSocket.OPEN) {
    ws.send(JSON.stringify({
      action: 'subscribe',
      channel: 'market',
      symbol: currentCoin.id
    }))
  }
}

const startHeartbeat = () => {
  heartbeatTimer = setInterval(() => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: 'heartbeat' }))
    }
  }, 30000)
}

const stopHeartbeat = () => {
  if (heartbeatTimer) {
    clearInterval(heartbeatTimer)
    heartbeatTimer = null
  }
}

const updateMarket = (data) => {
  currentCoin.close = data.close
  currentCoin.change = data.change
  currentCoin.rose = (data.change * 100).toFixed(2) + '%'
  currentCoin.high = data.high
  currentCoin.low = data.low
  currentCoin.volume = data.volume
  currentCoin.usdRate = data.usdRate
}

const updatePlate = (data) => {
  plate.askRows = data.asks || []
  plate.bidRows = data.bids || []
}

const updateTrade = (data) => {
  trade.rows = data.trades || []
}

const getPlate = (type = 'all') => {
  axios.post(`${host}${api.swap.plate}`, {
    symbol: currentCoin.id,
    type: type
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      updatePlate(res.data.data)
    }
  }).catch(() => {})
}

const getPlateFull = () => {
  axios.post(`${host}${api.swap.plate}`, {
    symbol: currentCoin.id
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      updatePlate(res.data.data)
    }
  }).catch(() => {})
}

const getCurrentPosition = () => {
  axios.post(`${host}${api.swap.currentPosition}`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      currentPosition.rows = res.data.data || []
    }
  }).catch(() => {})
}

const getCurrentOrder = () => {
  axios.post(`${host}${api.swap.currentOrder}`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      currentOrder.rows = res.data.data || []
    }
  }).catch(() => {})
}

const getHistoryOrder = () => {
  axios.post(`${host}${api.swap.historyOrder}`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      historyOrder.rows = res.data.data || []
    }
  }).catch(() => {})
}

const getWallet = () => {
  axios.post(`${host}${api.swap.wallet}`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      const data = res.data.data
      wallet.base = data.base || 0
      wallet.frozen = data.frozen || 0
      wallet.profit = data.profit || 0
    }
  }).catch(() => {})
}

const getLeverage = () => {
  axios.post(`${host}${api.swap.leverage}`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      form.leverage = res.data.data.leverage
      form.time = res.data.data.time
      form.patterns = res.data.data.patterns
    }
  }).catch(() => {})
}

const adjustLeverage = () => {
  const leverageMap = { 1: 1, 2: 2, 3: 5, 4: 10, 5: 20 }
  const leverage = leverageMap[form.leverageAdjustVal2]

  axios.post(`${host}${api.swap.leverageAdjust}`, {
    leverage: leverage
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      form.leverage = leverage
      form.leverageModal = false
      ElMessage.success('杠杆调整成功')
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const adjustTime = () => {
  const timeMap = { 1: 10, 2: 30, 3: 60, 4: 120, 5: 360, 6: 3000 }
  const time = timeMap[form.timeAdjustVal2]

  axios.post(`${host}${api.swap.timeAdjust}`, {
    time: time
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      form.time = time
      form.timeModal = false
      ElMessage.success('时间调整成功')
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const xxlever = (val) => {
  form.leverageAdjustVal2 = val
}

const xxlever2 = (val) => {
  form.timeAdjustVal2 = val
}

const init = () => {
  const id = router.currentRoute?.value?.params?.pair

  if (!id) {
    router.push('/swapindex/' + defaultPath.value)
  }

  ElLoading.service({ fullscreen: true })

  axios.post(`${host}${api.swap.symbolInfo}`, { id: id || defaultPath.value }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(response => {
    const resp = response.data
    if (!resp || resp.code !== 0) return

    basecion.value = resp.data.baseSymbol.toLowerCase()
    currentCoin.symbol = resp.data.coinSymbol + '/' + resp.data.baseSymbol
    currentCoin.coin = resp.data.coinSymbol
    currentCoin.id = resp.data.id || id || defaultPath.value
    currentCoin.type = resp.data.type
    currentCoin.base = resp.data.baseSymbol

    store?.commit('setSkin', skin.value)

    currentCoin.coinScale = resp.data.coinScale
    currentCoin.baseCoinScale = resp.data.baseCoinScale
    currentCoin.perUsdt = resp.data.shareNumber
    baseCoinScale.value = resp.data.baseCoinScale
    coinScale.value = resp.data.coinScale
    takerFee.value = resp.data.takerFee
    makerFee.value = resp.data.makerFee
    currentCoin.enableMarketBuy = resp.data.enableMarketBuy
    currentCoin.enableMarketSell = resp.data.enableMarketSell
    currentCoin.exchangeable = resp.data.exchangeable

    getCNYRate()
    getCoinInfo()
    getSymbol()
    getPlateFull()

    if (isLogin.value) {
      getWallet()
      getCurrentPosition()
      getCurrentOrder()
      getHistoryOrder()
      getLeverage()
    }

    ElLoading.service({ fullscreen: true }).close()
  }).catch(() => {
    ElLoading.service({ fullscreen: true }).close()
  })
}

onMounted(() => {
  init()
})

onBeforeUnmount(() => {
  stopHeartbeat()
  if (ws) {
    ws.close()
  }
  if (chartWidget) {
    chartWidget.remove()
  }
})
</script>

<style scoped lang="scss">
.container.swap {
  min-height: calc(100vh - 60px);
  background: #0b1520;
  padding: 20px;

  .main {
    max-width: 1400px;
    margin: 0 auto;
  }

  .center {
    width: 100%;

    .symbol {
      display: flex;
      align-items: center;
      padding: 15px 0;
      border-bottom: 1px solid #27313e;
      flex-wrap: wrap;

      .item {
        margin-left: 20px;
        display: flex;
        align-items: center;

        .coin {
          font-size: 18px;
          font-weight: 600;
          color: #fff;
        }

        .text {
          color: #61688a;
          font-size: 13px;
          margin-right: 8px;
        }

        .num {
          font-size: 16px;
          font-weight: 600;
          color: #fff;
        }

        .num.buy {
          color: #00c853;
        }

        .num.sell {
          color: #ff5252;
        }

        .price-cny {
          font-size: 12px;
          color: #61688a;
          margin-left: 8px;
        }
      }
    }

    .imgtable {
      margin-top: 10px;

      .handler {
        display: flex;
        margin-bottom: 10px;

        span {
          padding: 8px 20px;
          cursor: pointer;
          color: #61688a;
          border-bottom: 2px solid transparent;

          &.active {
            color: #f0a70a;
            border-bottom-color: #f0a70a;
          }

          &:hover {
            color: #fff;
          }
        }
      }

      #swap_kline_container {
        height: 500px;
        background: #17212e;
        border-radius: 8px;
        overflow: hidden;
      }

      .hidden {
        display: none;
      }
    }
  }

  .leverage {
    padding: 20px 10px;

    :deep(.el-slider) {
      .el-slider__marks-text {
        color: #61688a;
      }
    }
  }
}

.vertical-center-modal {
  :deep(.el-dialog) {
    margin-top: 10vh !important;
  }
}

.api {
  .coin-info {
    color: #333;
    font-size: 14px;
    line-height: 1.6;
  }
}

.info-icon {
  &:hover {
    color: #f0a70a;
  }
}
</style>

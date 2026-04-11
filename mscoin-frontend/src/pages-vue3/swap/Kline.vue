<template>
  <div class="container swap" :class="skin">
    <div class="main">
      <div style="display: flex; flex: 100%; flex-wrap: wrap; justify-content: space-between;">
        <div class="center">
          <!-- 币种信息 -->
          <div class="symbol">
            <div class="item" style="margin-left: 10px;">
              <span class="coin">{{ currentCoin.name }}</span>
              <el-popover trigger="hover" :width="300" placement="bottom-start">
                <template #reference>
                  <el-icon style="color: #546886; margin-left: 5px;"><InfoFilled /></el-icon>
                </template>
                <div class="api">
                  <div class="coin-info">{{ coinInfo.information }}</div>
                  <p style="text-align: right; margin-top: 10px;">
                    <a :href="coinInfo.infolink" target="_blank">{{ $t('swap.moredetail') }}</a>
                  </p>
                </div>
              </el-popover>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.NewPrice') }}</span>
              <span class="num" :class="{ buy: currentCoin.change > 0, sell: currentCoin.change < 0 }">
                {{ currentCoin.close?.toFixed(4) }}
              </span>
              <span class="price-cny">≈ ￥{{ (currentCoin.usdRate * CNYRate)?.toFixed(2) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.Change') }}</span>
              <span class="num" :class="{ buy: currentCoin.change > 0, sell: currentCoin.change < 0 }">
                {{ currentCoin.rose }}
              </span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.high') }}</span>
              <span class="num">{{ currentCoin.high?.toFixed(4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.low') }}</span>
              <span class="num">{{ currentCoin.low?.toFixed(4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.ExchangeNum') }}</span>
              <span class="num">{{ currentCoin.volume?.toFixed(4) }} {{ currentCoin.coin }}</span>
            </div>
          </div>

          <!-- K 线图和深度图切换 -->
          <div class="imgtable">
            <div class="handler">
              <span @click="changeImgTable('k')" :class="{ active: currentImgTable === 'k' }">
                {{ $t('swap.kline') }}
              </span>
              <span @click="changeImgTable('s')" :class="{ active: currentImgTable === 's' }">
                {{ $t('swap.depth') }}
              </span>
            </div>
            <div id="swap_kline_container" :class="{ hidden: currentImgTable === 's' }"></div>
            <DepthGraph :class="{ hidden: currentImgTable === 'k' }" ref="depthGraphRef" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - K 线图表组件 (合约交易)
 * 迁移点:
 * 1. 使用 <script setup> 语法和 Composition API
 * 2. Element Plus 组件替代 iView 组件
 * 3. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 * 4. TradingView 图表集成保持不变
 * 5. axios 替代 vue-resource
 */
import { ref, reactive, inject, onMounted, onBeforeUnmount, nextTick, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import axios from 'axios'
import moment from 'moment'
import { useRuntimeContract } from '../../config/runtime-vue3'

// 导入深度图组件
import DepthGraph from '../exchange/DepthGraph.vue'

// 导入 TradingView datafeed
import Datafeeds from '../../assets/js/charting_library/datafeed/swaptrade.js'

const store = inject('store')
const router = inject('router')
const { api } = useRuntimeContract()

const host = ''

// 响应式数据
const skin = ref('night')
const currentImgTable = ref('k')
const CNYRate = ref(null)
const datafeed = ref(null)
const symbol = ref('')
const memberId = ref(0)
const isLogin = ref(false)
const lang = ref('zh')

// 当前币种
const currentCoin = reactive({
  id: '',
  name: '',
  coin: '',
  base: '',
  symbol: '',
  close: 0,
  price: 0,
  rose: '',
  high: 0,
  low: 0,
  volume: 0,
  turnover: 0,
  usdRate: 0,
  change: 0,
  type: 'ALWAYS',
  coinScale: 6,
  baseCoinScale: 4,
  perUsdt: 100000
})

const coinInfo = reactive({
  information: '',
  infolink: ''
})

const coins = reactive({
  _map: [],
  USDT: [],
  BTC: [],
  ETH: [],
  USDT2: [],
  favor: []
})

// 深度图引用
const depthGraphRef = ref(null)

// TradingView widget 引用
let tvWidget = null

// 工具函数
const toFixed = (value, scale = 4) => {
  if (value === null || value === undefined) return 0
  return Number(value).toFixed(scale)
}

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

const dateFormat = (time) => {
  return moment(time).format('YYYY-MM-DD HH:mm:ss')
}

const changeImgTable = (str) => {
  currentImgTable.value = str
}

// 获取汇率
const getCNYRate = () => {
  axios.post(`${host}/market/exchange-rate/usd/cny`)
    .then(response => {
      const resp = response.data
      CNYRate.value = resp.data
    })
}

// 获取币种列表
const getSymbol = (type = 1) => {
  axios.post(`${host}${api.swap.thumb}`, { type })
    .then(response => {
      const resp = response.data
      // 先清空已有数据
      for (let i = 0; i < resp.length; i++) {
        const coin = resp[i]
        coin.base = resp[i].symbol.split('/')[1]
        coins[coin.base] = []
        coins[coin.base + '2'] = []
        coins._map = []
      }
      for (let i = 0; i < resp.length; i++) {
        const coin = resp[i]
        const base = coin.base
        coin.isFavor = coin.isFavor === undefined ? false : coin.isFavor
        coins[base].push(coin)
        coins[base + '2'].push(coin)
        coins._map[coin.symbol] = coin
        if (coin.symbol === symbol.value) {
          currentCoin.price = coin.close
          currentCoin.id = coin.id
        }
      }
    })
}

// 获取币种信息
const getCoinInfo = () => {
  axios.post(`${host}${api.market.coinInfo}`, { unit: currentCoin.coin })
    .then(response => {
      const resp = response.data
      if (resp) {
        coinInfo.information = resp.introduction
        coinInfo.infolink = resp.website
      }
    })
}

// 获取币种详情
const getSymbolInfo = () => {
  axios.post(`${host}${api.swap.symbolInfo}`, { id: currentCoin.id })
    .then(response => {
      const resp = response.data
      if (resp) {
        currentCoin.coinScale = resp.coinScale
        currentCoin.baseCoinScale = resp.baseCoinScale
        currentCoin.perUsdt = resp.shareNumber
      }
    })
}

// 初始化 K 线图表
const getKline = () => {
  const config = {
    autosize: true,
    height: 500,
    fullscreen: false,
    symbol: symbol.value,
    interval: '60',
    timezone: 'Asia/Shanghai',
    toolbar_bg: '#192330',
    container_id: 'swap_kline_container',
    datafeed: datafeed.value,
    library_path:
      process.env.NODE_ENV === 'production'
        ? '/assets/charting_library/'
        : '/src/assets/js/charting_library/',
    locale: getLocale(),
    debug: false,
    drawings_access: {
      type: 'black',
      tools: [{ name: 'Regression Trend' }]
    },
    disabled_features: [
      'header_resolutions',
      'timeframes_toolbar',
      'header_symbol_search',
      'header_chart_type',
      'header_compare',
      'header_undo_redo',
      'header_screenshot',
      'header_saveload',
      'use_localstorage_for_settings',
      'left_toolbar',
      'volume_force_overlay'
    ],
    enabled_features: [
      'hide_last_na_study_output',
      'move_logo_to_main_pane'
    ],
    custom_css_url: 'bundles/common.css',
    supported_resolutions: ['1', '5', '15', '30', '60', '1D', '1W', '1M'],
    charts_storage_url: 'http://saveload.tradingview.com',
    charts_storage_api_version: '1.1',
    client_id: 'tradingview.com',
    user_id: 'public_user_id',
    overrides: {
      'paneProperties.background': '#192330',
      'paneProperties.vertGridProperties.color': 'rgba(0,0,0,.1)',
      'paneProperties.horzGridProperties.color': 'rgba(0,0,0,.1)',
      'scalesProperties.textColor': '#61688A',
      'mainSeriesProperties.candleStyle.upColor': '#00b275',
      'mainSeriesProperties.candleStyle.downColor': '#f15057',
      'mainSeriesProperties.candleStyle.drawBorder': false,
      'mainSeriesProperties.candleStyle.wickUpColor': '#589065',
      'mainSeriesProperties.candleStyle.wickDownColor': '#AE4E54',
      'paneProperties.legendProperties.showLegend': false,
      'mainSeriesProperties.areaStyle.color1': 'rgba(71, 78, 112, 0.5)',
      'mainSeriesProperties.areaStyle.color2': 'rgba(71, 78, 112, 0.5)',
      'mainSeriesProperties.areaStyle.linecolor': '#9194a4',
      volumePaneSize: 'small'
    },
    time_frames: [
      { text: '1min', resolution: '1', description: 'realtime', title: '实时' },
      { text: '1min', resolution: '1', description: '1min' },
      { text: '5min', resolution: '5', description: '5min' },
      { text: '15min', resolution: '15', description: '15min' },
      { text: '30min', resolution: '30', description: '30min' },
      { text: '1hour', resolution: '60', description: '1hour' },
      { text: '4hour', resolution: '240', description: '4hour' },
      { text: '1day', resolution: '1D', description: '1day' },
      { text: '1week', resolution: '1W', description: '1week' },
      { text: '1mon', resolution: '1M', description: '1mon' }
    ]
  }

  if (skin.value === 'day') {
    config.toolbar_bg = '#fff'
    config.custom_css_url = 'bundles/common_day.css'
    config.overrides['paneProperties.background'] = '#fff'
    config.overrides['mainSeriesProperties.candleStyle.upColor'] = '#a6d3a5'
    config.overrides['mainSeriesProperties.candleStyle.downColor'] = '#ffa5a6'
  }

  // 动态加载 TradingView
  const script = document.createElement('script')
  script.src = process.env.NODE_ENV === 'production'
    ? '/assets/charting_library/charting_library.min.js'
    : '/src/assets/js/charting_library/charting_library.min.js'
  script.onload = () => {
    tvWidget = window.tvWidget = new TradingView.widget(config)
    tvWidget.onChartReady(function() {
      tvWidget.chart().executeActionById('drawingToolbarAction')
      tvWidget.chart().createStudy('Moving Average', false, false, [5], null, { 'plot.color': '#EDEDED' })
      tvWidget.chart().createStudy('Moving Average', false, false, [10], null, { 'plot.color': '#ffe000' })
      tvWidget.chart().createStudy('Moving Average', false, false, [30], null, { 'plot.color': '#ce00ff' })
      tvWidget.chart().createStudy('Moving Average', false, false, [60], null, { 'plot.color': '#00adff' })

      // 创建时间周期按钮
      createPeriodButtons(tvWidget)
    })
  }
  document.head.appendChild(script)
}

// 创建时间周期按钮
const createPeriodButtons = (widget) => {
  // 分时按钮
  widget.createButton()
    .attr('title', 'realtime')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(3)
      widget.setSymbol('', '1')
    })
    .append('<span>分时</span>')

  // M1
  widget.createButton()
    .attr('title', 'M1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '1')
    })
    .append('<span>M1</span>')

  // M5
  widget.createButton()
    .attr('title', 'M5')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '5')
    })
    .append('<span>M5</span>')

  // M15
  widget.createButton()
    .attr('title', 'M15')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '15')
    })
    .append('<span>M15</span>')

  // M30
  widget.createButton()
    .attr('title', 'M30')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '30')
    })
    .append('<span>M30</span>')

  // H1
  widget.createButton()
    .attr('title', 'H1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '60')
    })
    .append('<span>H1</span>')
    .addClass('selected')

  // D1
  widget.createButton()
    .attr('title', 'D1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '1D')
    })
    .append('<span>D1</span>')

  // W1
  widget.createButton()
    .attr('title', 'W1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '1W')
    })
    .append('<span>W1</span>')
}

const getLocale = () => {
  return lang.value === 'zh' ? 'zh' : 'en'
}

// 初始化
const init = () => {
  const params = router.currentRoute.value.params.pair || 'btc_usdt'
  const base = params.split('_')[1]
  const coin = params.toUpperCase().split('_')[0]

  symbol.value = `${coin}/${base}`
  currentCoin.symbol = symbol.value
  currentCoin.coin = coin
  currentCoin.base = base

  store.commit('setSkin', skin.value)

  getCNYRate()
  getSymbolInfo()
  getSymbol()
  getCoinInfo()

  // 初始化 datafeed
  datafeed.value = new Datafeeds.UDF.UdfDatafeed('/market', {})

  // 延迟初始化图表，等待 DOM 渲染
  nextTick(() => {
    getKline()
  })
}

// 监听语言变化
watch(() => store.state.lang, (newLang) => {
  lang.value = newLang
})

// 生命周期
onMounted(() => {
  isLogin.value = !!localStorage.getItem('TOKEN')
  memberId.value = store.getters?.member?.id || 0
  lang.value = store.state.lang || 'zh'
  init()
})

onBeforeUnmount(() => {
  // 清理 TradingView
  if (tvWidget) {
    tvWidget.remove()
    tvWidget = null
    window.tvWidget = null
  }
})
</script>

<style scoped lang="scss">
.container.swap {
  min-height: calc(100vh - 200px);
  background-color: #0b1520;
  color: #fff;

  .main {
    width: 99%;
    margin-left: 0.5%;
    max-width: 100%;
    overflow-x: hidden;

    .center {
      width: 100%;
      min-width: 0;
      overflow: hidden;

      .symbol {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        padding: 15px 0;
        border-bottom: 1px solid #27313e;

        .item {
          margin: 0 15px;

          .coin {
            font-size: 18px;
            font-weight: 700;
            color: #fff;
          }

          .text {
            font-size: 12px;
            color: #76839b;
          }

          .num {
            display: block;
            font-size: 18px;
            margin-top: 5px;
          }

          .buy {
            color: #00b050;
          }

          .sell {
            color: #f0a70a;
          }

          .price-cny {
            font-size: 12px;
            color: #76839b;
          }
        }
      }

      .imgtable {
        margin-top: 20px;

        .handler {
          display: flex;
          margin-bottom: 10px;

          span {
            padding: 8px 20px;
            cursor: pointer;
            border-bottom: 2px solid transparent;
            color: #76839b;

            &.active {
              color: #f0a70a;
              border-bottom-color: #f0a70a;
            }

            &:hover {
              color: #f0a70a;
            }
          }
        }

        #swap_kline_container {
          height: 500px;
          width: 100%;
        }

        .hidden {
          display: none;
        }
      }
    }
  }
}

// 白天模式
.container.swap.day {
  background-color: #f5f5f5;

  .center {
    .symbol {
      border-bottom-color: #e8e8e8;

      .coin {
        color: #333;
      }

      .text {
        color: #999;
      }
    }
  }
}

// TradingView 按钮样式
:deep(.chart-widget-toolbar-button) {
  &.selected {
    background-color: #f0a70a !important;
    color: #000 !important;
  }
}
</style>

<template>
  <div class="index-page">
    <div class="market-section">
      <div class="announcement-bar">
        <el-icon class="notice-icon"><Bell /></el-icon>
        <span class="notice-text">{{ announcement }}</span>
      </div>

      <div class="search-bar">
        <el-input
          v-model="searchText"
          placeholder="搜索币种"
          :prefix-icon="Search"
          clearable
        />
      </div>

      <div class="tabs-wrap">
        <el-tabs v-model="activeZone">
          <el-tab-pane label="USDT 区" name="usdt" />
          <el-tab-pane label="BTC 区" name="btc" />
          <el-tab-pane label="ETH 区" name="eth" />
          <el-tab-pane label="我的收藏" name="favor" />
        </el-tabs>
      </div>

      <div class="market-table">
        <table class="coin-table">
          <thead>
            <tr>
              <th width="50">收藏</th>
              <th>币种</th>
              <th>最新价</th>
              <th>涨跌幅</th>
              <th>最高价</th>
              <th>最低价</th>
              <th>24h 涨跌</th>
              <th>24h 成交量</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="coin in filteredCoins" :key="coin.symbol">
              <td>
                <el-icon
                  class="favor-icon"
                  :class="{ active: coin.isFavor }"
                  @click="toggleFavor(coin)"
                >
                  <StarFilled v-if="coin.isFavor" />
                  <Star v-else />
                </el-icon>
              </td>
              <td>
                <div class="coin-info">
                  <span class="coin-name">{{ coin.name }}</span>
                  <span class="coin-symbol">{{ coin.symbol }}</span>
                </div>
              </td>
              <td>
                <div class="price-info">
                  <div class="price">{{ formatPrice(coin.price) }}</div>
                  <div class="price-cny">≈ {{ formatPrice(coin.priceCNY) }} CNY</div>
                </div>
              </td>
              <td>
                <span :class="coin.rise >= 0 ? 'rise' : 'fall'">
                  {{ coin.rise >= 0 ? '+' : '' }}{{ formatPercent(coin.rise) }}
                </span>
              </td>
              <td>{{ formatPrice(coin.high) }}</td>
              <td>{{ formatPrice(coin.low) }}</td>
              <td>
                <span :class="coin.change24h >= 0 ? 'rise' : 'fall'">
                  {{ formatSignedPrice(coin.change24h) }}
                </span>
              </td>
              <td>{{ formatVolume(coin.volume) }}</td>
              <td>
                <div class="action-buttons">
                  <router-link :to="`/exchange/${coin.href}`" class="trade-btn">
                    去交易
                  </router-link>
                </div>
              </td>
            </tr>
          </tbody>
        </table>

        <div v-if="filteredCoins.length === 0" class="empty-result">
          暂无相关币种
        </div>
      </div>
    </div>

    <div class="features-section" id="page2">
      <el-row :gutter="20">
        <el-col :span="6" v-for="feature in features" :key="feature.title">
          <div class="feature-card">
            <img :src="feature.image" :alt="feature.title">
            <p class="title">{{ feature.title }}</p>
            <p class="desc">{{ feature.description }}</p>
          </div>
        </el-col>
      </el-row>
    </div>

    <div class="ads-section" id="page4">
      <el-carousel v-if="picList.length > 0" height="400px">
        <el-carousel-item v-for="pic in picList" :key="pic.id">
          <img :src="pic.image" class="banner-image" alt="banner">
        </el-carousel-item>
      </el-carousel>
    </div>

    <div class="agent-section" id="page6">
      <div class="agent-panel">
        <h3>代理专区</h3>
        <el-row :gutter="20">
          <el-col :span="6">
            <div class="agent-item">
              <el-icon class="agent-icon"><User /></el-icon>
              <p>推广返佣</p>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="agent-item">
              <el-icon class="agent-icon"><Money /></el-icon>
              <p>手续费分成</p>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="agent-item">
              <el-icon class="agent-icon"><TrendCharts /></el-icon>
              <p>团队管理</p>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="agent-item">
              <el-icon class="agent-icon"><Service /></el-icon>
              <p>专属客服</p>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, ref } from 'vue'
import {
  ElCarousel,
  ElCarouselItem,
  ElCol,
  ElIcon,
  ElInput,
  ElMessage,
  ElRow,
  ElTabPane,
  ElTabs
} from 'element-plus'
import { Bell, Money, Search, Service, Star, StarFilled, TrendCharts, User } from '@element-plus/icons-vue'
import axios from 'axios'
import io from 'socket.io-client'
import { useRouter } from 'vue-router'
import api from '@/config/api'
import { mapThumbToCoinViewModel, mergeRealtimeThumb } from '@/pages-vue3/index/market'
import bannerBg from '@/assets/images/bannerbg.png'
import featureChoose from '@/assets/images/feature_choose.png'
import featureFast from '@/assets/images/feature_fast.png'
import featureGlobal from '@/assets/images/feature_global.png'
import featureSafe from '@/assets/images/feature_safe.png'

const router = useRouter()

const host = ''
const defaultAnnouncement = '欢迎来到 MSCON 交易平台！'
const fallbackCoins = [
  {
    symbol: 'BTC/USDT',
    name: 'BTC',
    base: 'USDT',
    href: 'btc_usdt',
    price: 68234.12,
    priceCNY: 493216.88,
    rise: 2.36,
    high: 68980.55,
    low: 67112.4,
    change24h: 1575.88,
    volume: 18452.23,
    isFavor: false
  },
  {
    symbol: 'ETH/USDT',
    name: 'ETH',
    base: 'USDT',
    href: 'eth_usdt',
    price: 3521.48,
    priceCNY: 25472.9,
    rise: 1.42,
    high: 3588.66,
    low: 3460.25,
    change24h: 49.28,
    volume: 92840.56,
    isFavor: false
  },
  {
    symbol: 'SOL/USDT',
    name: 'SOL',
    base: 'USDT',
    href: 'sol_usdt',
    price: 172.36,
    priceCNY: 1246.58,
    rise: -0.85,
    high: 175.92,
    low: 169.48,
    change24h: -1.47,
    volume: 265430.11,
    isFavor: false
  }
]
const fallbackAds = [
  {
    id: 'local-banner',
    image: bannerBg
  }
]

const searchText = ref('')
const activeZone = ref('usdt')
const announcement = ref(defaultAnnouncement)
const coinList = ref([])
const favorList = ref([])
const picList = ref([])

const features = [
  {
    title: '极速撮合',
    description: '撮合延迟更低，交易体验更稳定。',
    image: featureFast
  },
  {
    title: '安全可靠',
    description: '冷热钱包隔离，多重签名保护资产。',
    image: featureSafe
  },
  {
    title: '全球服务',
    description: '支持多地区用户访问与交易。',
    image: featureGlobal
  },
  {
    title: '专业团队',
    description: '持续优化撮合、风控与服务体验。',
    image: featureChoose
  }
]

const allCoins = computed(() => {
  if (activeZone.value === 'favor') {
    return coinList.value.filter((coin) => coin.isFavor)
  }
  const zoneMap = {
    usdt: 'USDT',
    btc: 'BTC',
    eth: 'ETH'
  }
  return coinList.value.filter((coin) => coin.base === zoneMap[activeZone.value])
})

const filteredCoins = computed(() => {
  if (!searchText.value) return allCoins.value
  const keyword = searchText.value.toLowerCase()
  return allCoins.value.filter((coin) =>
    coin.name.toLowerCase().includes(keyword) ||
    coin.symbol.toLowerCase().includes(keyword)
  )
})

const formatPrice = (price) => {
  if (price === undefined || price === null || Number.isNaN(Number(price))) return '0.00'
  const num = Number(price)
  if (num >= 1000) return num.toFixed(2)
  if (num >= 1) return num.toFixed(4)
  if (num >= 0.01) return num.toFixed(6)
  return num.toFixed(8)
}

const formatPercent = (value) => {
  if (value === undefined || value === null || Number.isNaN(Number(value))) return '0.00%'
  return `${Number(value).toFixed(2)}%`
}

const formatVolume = (volume) => {
  if (volume === undefined || volume === null || Number.isNaN(Number(volume))) return '0'
  return Number(volume).toLocaleString()
}

const formatSignedPrice = (value) => {
  if (value === undefined || value === null || Number.isNaN(Number(value))) return '0.00'
  const amount = Number(value)
  const sign = amount > 0 ? '+' : amount < 0 ? '-' : ''
  return `${sign}${formatPrice(Math.abs(amount))}`
}

const mapCoin = (coin) => {
  return mapThumbToCoinViewModel(coin, favorList.value)
}

const setFallbackCoinList = () => {
  if (coinList.value.length > 0) return
  coinList.value = fallbackCoins.map((coin) => ({
    ...coin,
    isFavor: favorList.value.includes(coin.symbol)
  }))
}

const updateCoinFavorState = () => {
  coinList.value = coinList.value.map((coin) => ({
    ...coin,
    isFavor: favorList.value.includes(coin.symbol)
  }))
}

const authHeaders = () => {
  const token = localStorage.getItem('TOKEN')
  return token ? { 'x-auth-token': token } : {}
}

const toggleFavor = async (coin) => {
  const token = localStorage.getItem('TOKEN')
  if (!token) {
    ElMessage.warning('请先登录')
    router.push('/login')
    return
  }

  const request = coin.isFavor ? api.exchange.favorDelete : api.exchange.favorAdd

  try {
    const response = await axios.post(host + request, { symbol: coin.symbol }, {
      headers: authHeaders()
    })

    if (response.data.code === 0) {
      coin.isFavor = !coin.isFavor
      if (coin.isFavor) {
        favorList.value = Array.from(new Set([...favorList.value, coin.symbol]))
      } else {
        favorList.value = favorList.value.filter((symbol) => symbol !== coin.symbol)
      }
      ElMessage.success(coin.isFavor ? '已添加收藏' : '已取消收藏')
      updateCoinFavorState()
      return
    }

    ElMessage.error(response.data.message || '收藏操作失败')
  } catch {
    ElMessage.error('收藏操作失败')
  }
}

const loadFavorList = async () => {
  const token = localStorage.getItem('TOKEN')
  if (!token) return

  try {
    const response = await axios.post(host + api.exchange.favorFind, {}, {
      headers: authHeaders()
    })
    const payload = response.data?.data ?? response.data ?? []
    if (Array.isArray(payload)) {
      favorList.value = payload.map((item) => item.symbol ?? item).filter(Boolean)
      updateCoinFavorState()
    }
  } catch {
    favorList.value = []
  }
}

const loadCoinList = async () => {
  try {
    const response = await axios.post(host + api.market.thumbTrend, {}, {
      headers: authHeaders()
    })
    if (response.data.code === 0) {
      coinList.value = (response.data.data || []).map(mapCoin)
      return
    }
  } catch {
    if (import.meta.env.DEV) {
      setFallbackCoinList()
    }
  }
}

const loadAds = async () => {
  try {
    const response = await axios.post(`${host}/uc/ancillary/system/advertise`, {
      sysAdvertiseLocation: 1,
      lang: 'CN'
    }, {
      headers: authHeaders()
    })
    if (response.data.code === 0) {
      picList.value = (response.data.data || []).map((item) => ({
        ...item,
        image: item.url || item.image || bannerBg
      }))
      return
    }
  } catch {
    if (import.meta.env.DEV) {
      picList.value = fallbackAds
    }
  }
}

const loadAnnouncement = async () => {
  try {
    const response = await axios.post(host + api.uc.announcement, {
      pageNo: 1,
      pageSize: 1,
      lang: 'CN'
    }, {
      headers: authHeaders()
    })
    if (response.data.code === 0 && response.data.data.content?.length > 0) {
      announcement.value = response.data.data.content[0].title
      return
    }
  } catch {
    announcement.value = defaultAnnouncement
  }
}

let socket = null
let reconnectTimer = null

const connectSocket = () => {
  if (socket) return

  socket = io(window.location.origin, {
    path: '/socket.io',
    transports: ['websocket', 'polling'],
    reconnection: false
  })

  socket.on('/topic/market/thumb', (message) => {
    try {
      updatePrice(JSON.parse(message))
    } catch (error) {
      console.error('Socket message parse error:', error)
    }
  })

  socket.on('disconnect', () => {
    socket = null
    if (!reconnectTimer) {
      reconnectTimer = setTimeout(() => {
        reconnectTimer = null
        connectSocket()
      }, 5000)
    }
  })
}

const updatePrice = (payload) => {
  const index = coinList.value.findIndex((coin) => coin.symbol === payload.symbol)
  if (index === -1) return

  const nextCoin = mapCoin(mergeRealtimeThumb(coinList.value[index], payload))
  coinList.value.splice(index, 1, nextCoin)
}

onMounted(async () => {
  window.scrollTo(0, 0)
  await loadFavorList()
  await loadCoinList()
  await loadAds()
  await loadAnnouncement()
  connectSocket()
})

onBeforeUnmount(() => {
  if (socket) {
    socket.close()
    socket = null
  }
  if (reconnectTimer) {
    clearTimeout(reconnectTimer)
    reconnectTimer = null
  }
})
</script>

<style scoped lang="scss">
.index-page {
  min-height: 100vh;
  background: #0b1520;
}

.market-section {
  padding: 80px 0;
  background: #17212e;

  .announcement-bar {
    display: flex;
    align-items: center;
    padding: 15px 30px;
    background: #1f2833;
    border-radius: 8px;
    margin-bottom: 20px;

    .notice-icon {
      color: #f0a70a;
      font-size: 18px;
      margin-right: 10px;
    }

    .notice-text {
      color: #fff;
      font-size: 14px;
    }
  }

  .search-bar {
    margin-bottom: 20px;

    .el-input__wrapper {
      background: #1f2833;
      border: 1px solid #27313e;

      .el-input__inner {
        color: #fff;
      }
    }
  }

  .tabs-wrap {
    margin-bottom: 30px;

    :deep(.el-tabs__header) {
      .el-tabs__item {
        color: #999;

        &.is-active {
          color: #f0a70a;
        }
      }
    }
  }

  .market-table {
    overflow-x: auto;

    .coin-table {
      width: 100%;
      border-collapse: collapse;

      th {
        padding: 15px 10px;
        text-align: left;
        color: #999;
        font-weight: 400;
        font-size: 14px;
        border-bottom: 1px solid #27313e;
      }

      td {
        padding: 20px 10px;
        color: #fff;
        border-bottom: 1px solid #1f2833;

        .favor-icon {
          cursor: pointer;
          font-size: 20px;
          color: #666;

          &.active {
            color: #f0a70a;
          }
        }

        .coin-info {
          display: flex;
          flex-direction: column;

          .coin-name {
            font-weight: 600;
            margin-bottom: 4px;
          }

          .coin-symbol {
            font-size: 12px;
            color: #999;
          }
        }

        .price-info {
          .price {
            font-size: 16px;
            font-weight: 600;
          }

          .price-cny {
            font-size: 12px;
            color: #999;
          }
        }

        .rise {
          color: #00c853;
        }

        .fall {
          color: #ff5252;
        }

        .action-buttons {
          .trade-btn {
            display: inline-block;
            padding: 6px 20px;
            background: #f0a70a;
            color: #fff;
            border-radius: 4px;
            text-decoration: none;

            &:hover {
              background: #ffb319;
            }
          }
        }
      }
    }

    .empty-result {
      text-align: center;
      padding: 50px;
      color: #999;
    }
  }
}

.features-section {
  padding: 80px 30px;
  background: #0b1520;

  .feature-card {
    text-align: center;
    padding: 30px 20px;
    background: #17212e;
    border-radius: 10px;
    margin-bottom: 20px;

    img {
      width: 100px;
      height: 100px;
      margin-bottom: 20px;
      object-fit: contain;
    }

    .title {
      color: #fff;
      font-size: 18px;
      margin-bottom: 10px;
    }

    .desc {
      color: #999;
      font-size: 14px;
      line-height: 1.7;
    }
  }
}

.ads-section {
  padding: 40px 30px;
  background: #17212e;

  .banner-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  :deep(.el-carousel__container) {
    border-radius: 10px;
    overflow: hidden;
  }
}

.agent-section {
  padding: 80px 30px;
  background: #0b1520;

  .agent-panel {
    max-width: 1200px;
    margin: 0 auto;
    background: #17212e;
    padding: 40px 30px;
    border-radius: 10px;

    h3 {
      color: #fff;
      text-align: center;
      margin-bottom: 40px;
      font-size: 24px;
    }
  }

  .agent-item {
    text-align: center;
    padding: 30px 20px;
    background: #1f2833;
    border-radius: 8px;
    margin-bottom: 20px;

    .agent-icon {
      font-size: 48px;
      color: #f0a70a;
      margin-bottom: 15px;
    }

    p {
      color: #fff;
      font-size: 16px;
    }
  }
}
</style>

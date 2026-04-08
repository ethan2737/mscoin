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
          prefix-icon="Search"
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
                  <Star v-if="coin.isFavor" />
                  <StarFilled v-else />
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
                  <div class="price-cny">≈{{ formatPrice(coin.priceCNY) }} CNY</div>
                </div>
              </td>
              <td>
                <span :class="coin.rise >= 0 ? 'rise' : 'fall'">
                  {{ coin.rise >= 0 ? '+' : '' }}{{ coin.rise }}%
                </span>
              </td>
              <td>{{ formatPrice(coin.high) }}</td>
              <td>{{ formatPrice(coin.low) }}</td>
              <td>
                <span :class="coin.change24h >= 0 ? 'rise' : 'fall'">
                  {{ coin.change24h >= 0 ? '+' : '' }}{{ coin.change24h }}%
                </span>
              </td>
              <td>{{ formatVolume(coin.volume) }}</td>
              <td>
                <div class="action-buttons">
                  <router-link :to="'/ctc?symbol=' + coin.symbol" class="trade-btn">
                    买入
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
            <img :src="feature.image" alt="">
            <p class="title">{{ feature.title }}</p>
            <p class="desc">{{ feature.description }}</p>
          </div>
        </el-col>
      </el-row>
    </div>

    <div class="ads-section" id="page4">
      <el-carousel v-if="picList.length > 0" height="400px">
        <el-carousel-item v-for="pic in picList" :key="pic.id">
          <img :src="pic.image" style="width: 100%; height: 100%; object-fit: cover;" alt="">
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
              <p>手续费分红</p>
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
/**
 * Vue 3 迁移 - 首页
 */
import { ref, reactive, computed, onMounted, onBeforeUnmount } from 'vue'
import { ElMessage, ElInput, ElTabs, ElTabPane, ElIcon, ElCarousel, ElCarouselItem, ElRow, ElCol } from 'element-plus'
import { Star, StarFilled, Bell, Search, User, Money, TrendCharts, Service } from '@element-plus/icons-vue'
import axios from 'axios'
import { useRouter, useStore } from 'vue-router/composables'

const router = useRouter()
const store = useStore()

const host = 'http://localhost'

const searchText = ref('')
const activeZone = ref('usdt')
const announcement = ref('欢迎来到 MSCON 交易平台！')
const coinList = ref([])
const favorList = ref([])
const picList = ref([])

const features = [
  {
    title: '极速交易',
    description: '匹配速度快，交易延时低',
    image: 'https://kicks.oss-cn-shanghai.aliyuncs.com/2019/feature1.png'
  },
  {
    title: '安全无忧',
    description: '冷热钱包隔离，多重签名',
    image: 'https://kicks.oss-cn-shanghai.aliyuncs.com/2019/feature2.png'
  },
  {
    title: '全球服务',
    description: '支持多个国家和地区',
    image: 'https://kicks.oss-cn-shanghai.aliyuncs.com/2019/feature3.png'
  },
  {
    title: '专业团队',
    description: '来自全球顶尖金融机构',
    image: 'https://kicks.oss-cn-shanghai.aliyuncs.com/2019/feature4.png'
  }
]

const lang = computed(() => store.state.lang)

const allCoins = computed(() => {
  if (activeZone.value === 'favor') {
    return coinList.value.filter(c => c.isFavor)
  }
  const zoneMap = {
    usdt: '_USDT',
    btc: '_BTC',
    eth: '_ETH'
  }
  const zone = zoneMap[activeZone.value]
  return coinList.value.filter(c => c.symbol.endsWith(zone))
})

const filteredCoins = computed(() => {
  if (!searchText.value) return allCoins.value
  const text = searchText.value.toLowerCase()
  return allCoins.value.filter(c =>
    c.name.toLowerCase().includes(text) ||
    c.symbol.toLowerCase().includes(text)
  )
})

const formatPrice = (price) => {
  if (price === undefined || price === null) return '0.00'
  const num = parseFloat(price)
  if (num >= 1000) return num.toFixed(2)
  if (num >= 1) return num.toFixed(4)
  if (num >= 0.01) return num.toFixed(6)
  return num.toFixed(8)
}

const formatVolume = (volume) => {
  if (volume === undefined || volume === null) return '0'
  return parseFloat(volume).toLocaleString()
}

const toggleFavor = (coin) => {
  const token = localStorage.getItem('TOKEN')
  if (!token) {
    ElMessage.warning('请先登录')
    router.push('/login')
    return
  }

  const params = { symbol: coin.symbol }
  axios.post(`${host}/uc/collect/coin`, params, {
    headers: { 'x-auth-token': token }
  }).then(res => {
    if (res.data.code === 0) {
      coin.isFavor = !coin.isFavor
      ElMessage.success(coin.isFavor ? '已添加收藏' : '已取消收藏')
    } else {
      ElMessage.error(res.data.message)
    }
  }).catch(() => {})
}

const loadCoinList = () => {
  axios.post(`${host}/market/coin/all`, {}, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      coinList.value = res.data.data.map(coin => ({
        ...coin,
        isFavor: favorList.value.includes(coin.symbol)
      }))
    }
  }).catch(() => {})
}

const loadFavorList = () => {
  const token = localStorage.getItem('TOKEN')
  if (!token) return

  axios.post(`${host}/uc/collect/list`, {}, {
    headers: { 'x-auth-token': token }
  }).then(res => {
    if (res.data.code === 0) {
      favorList.value = res.data.data || []
    }
  }).catch(() => {})
}

const loadAds = () => {
  axios.post(`${host}/uc/ancillary/advertise/list`, { type: 1 }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0) {
      picList.value = res.data.data || []
    }
  }).catch(() => {})
}

const loadAnnouncement = () => {
  axios.post(`${host}/uc/ancillary/help/list`, {
    cateId: 4,
    pageNo: 1,
    pageSize: 1
  }, {
    headers: { 'x-auth-token': localStorage.getItem('TOKEN') }
  }).then(res => {
    if (res.data.code === 0 && res.data.data.content.length > 0) {
      announcement.value = res.data.data.content[0].title
    }
  }).catch(() => {})
}

let socket = null
let heartbeatTimer = null

const connectSocket = () => {
  socket = new WebSocket('ws://localhost:9100/ws')

  socket.onopen = () => {
    console.log('WebSocket connected')
    startHeartbeat()
  }

  socket.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data)
      if (data.type === 'price') {
        updatePrice(data.data)
      }
    } catch (e) {
      console.error('Socket message parse error:', e)
    }
  }

  socket.onclose = () => {
    console.log('WebSocket closed')
    stopHeartbeat()
    setTimeout(() => {
      connectSocket()
    }, 5000)
  }

  socket.onerror = (error) => {
    console.error('WebSocket error:', error)
  }
}

const startHeartbeat = () => {
  heartbeatTimer = setInterval(() => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ type: 'heartbeat' }))
    }
  }, 30000)
}

const stopHeartbeat = () => {
  if (heartbeatTimer) {
    clearInterval(heartbeatTimer)
    heartbeatTimer = null
  }
}

const updatePrice = (data) => {
  const index = coinList.value.findIndex(c => c.symbol === data.symbol)
  if (index !== -1) {
    const coin = coinList.value[index]
    coin.price = data.price
    coin.priceCNY = data.priceCNY
    coin.rise = data.rise
    coin.high = data.high
    coin.low = data.low
    coin.change24h = data.change24h
    coin.volume = data.volume
  }
}

onMounted(() => {
  window.scrollTo(0, 0)
  loadCoinList()
  loadFavorList()
  loadAds()
  loadAnnouncement()
  connectSocket()
})

onBeforeUnmount(() => {
  stopHeartbeat()
  if (socket) {
    socket.close()
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
    }

    .title {
      color: #fff;
      font-size: 18px;
      margin-bottom: 10px;
    }

    .desc {
      color: #999;
      font-size: 14px;
    }
  }
}

.ads-section {
  padding: 40px 30px;
  background: #17212e;

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

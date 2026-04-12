<template>
  <div ref="containerRef" class="depth-chart" @mouseleave="clearHover">
    <canvas ref="canvasRef" class="depth-chart__canvas"></canvas>
    <canvas ref="coverCanvasRef" class="depth-chart__canvas"></canvas>
    <div
      v-if="hoverState.visible"
      class="depth-chart__tooltip"
      :style="{ left: `${hoverState.left}px`, top: `${hoverState.top}px` }"
    >
      <div class="depth-chart__tooltip-title" :class="`depth-chart__tooltip-title--${hoverState.side}`">
        {{ hoverState.sideLabel }}
      </div>
      <div class="depth-chart__tooltip-row">
        价格 ({{ hoverState.priceUnit }}): {{ hoverState.priceText }}
      </div>
      <div class="depth-chart__tooltip-row">
        累计数量 ({{ hoverState.amountUnit }}): {{ hoverState.amountText }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { nextTick, onMounted, reactive, ref, watch } from 'vue'

const CHART_PADDING = Object.freeze({
  top: 18,
  right: 90,
  bottom: 54,
  left: 12
})

const props = defineProps({
  values: {
    type: Object,
    default: () => ({
      bid: { items: [], highestPrice: 0, lowestPrice: 0 },
      ask: { items: [], highestPrice: 0, lowestPrice: 0 },
      skin: 'night',
      priceUnit: '',
      amountUnit: '',
      priceScale: 4,
      amountScale: 4
    })
  }
})

const containerRef = ref(null)
const canvasRef = ref(null)
const coverCanvasRef = ref(null)
const contextRef = ref(null)
const coverContextRef = ref(null)
const pWidth = ref(900)
const pHeight = ref(300)
const xObj = ref({})
const yObj = ref({})
const buyPoints = ref([])
const sellPoints = ref([])
const hoverState = reactive({
  visible: false,
  left: 0,
  top: 0,
  side: 'buy',
  sideLabel: '买盘 Bid',
  priceUnit: '--',
  amountUnit: '--',
  priceText: '--',
  amountText: '--'
})

const getChartBounds = () => {
  const right = Math.max(pWidth.value - CHART_PADDING.right, CHART_PADDING.left + 120)
  const bottom = Math.max(pHeight.value - CHART_PADDING.bottom, CHART_PADDING.top + 60)

  return {
    left: CHART_PADDING.left,
    right,
    top: CHART_PADDING.top,
    bottom,
    middle: CHART_PADDING.left + (right - CHART_PADDING.left) / 2,
    width: right - CHART_PADDING.left,
    height: bottom - CHART_PADDING.top
  }
}

const getPriceScale = () => Number.isFinite(Number(props.values?.priceScale))
  ? Number(props.values.priceScale)
  : 4

const getAmountScale = () => Number.isFinite(Number(props.values?.amountScale))
  ? Number(props.values.amountScale)
  : 4

const formatDepthValue = (value, scale) => {
  const numericValue = Number(value)
  if (!Number.isFinite(numericValue)) return '--'
  return numericValue.toFixed(Math.max(scale, 0))
}

const safePlate = () => {
  const bid = props.values?.bid ?? { items: [], highestPrice: 0, lowestPrice: 0 }
  const ask = props.values?.ask ?? { items: [], highestPrice: 0, lowestPrice: 0 }

  return {
    bid: {
      highestPrice: Number(bid.highestPrice) || 0,
      lowestPrice: Number(bid.lowestPrice) || 0,
      items: Array.isArray(bid.items)
        ? bid.items.map(item => ({
          price: Number(item.price) || 0,
          amount: Number(item.amount) || 0
        }))
        : []
    },
    ask: {
      highestPrice: Number(ask.highestPrice) || 0,
      lowestPrice: Number(ask.lowestPrice) || 0,
      items: Array.isArray(ask.items)
        ? ask.items.map(item => ({
          price: Number(item.price) || 0,
          amount: Number(item.amount) || 0
        }))
        : []
    },
    skin: props.values?.skin ?? 'night',
    priceUnit: props.values?.priceUnit ?? '',
    amountUnit: props.values?.amountUnit ?? '',
    priceScale: getPriceScale(),
    amountScale: getAmountScale()
  }
}

const resetCanvas = () => {
  if (!canvasRef.value || !coverCanvasRef.value || !containerRef.value) return false

  const width = Math.floor(containerRef.value.offsetWidth || 0)
  const height = Math.floor(containerRef.value.offsetHeight || 0)
  if (!width || !height) return false

  pWidth.value = width
  pHeight.value = height
  canvasRef.value.width = width
  coverCanvasRef.value.width = width
  canvasRef.value.height = height
  coverCanvasRef.value.height = height
  contextRef.value = canvasRef.value.getContext('2d')
  coverContextRef.value = coverCanvasRef.value.getContext('2d')
  return true
}

const clearHoverCanvas = () => {
  if (!coverContextRef.value || !coverCanvasRef.value) return
  coverContextRef.value.clearRect(0, 0, coverCanvasRef.value.width, coverCanvasRef.value.height)
}

const clearHover = () => {
  hoverState.visible = false
  clearHoverCanvas()
}

const convertTotal = (plate) => {
  for (let i = 1; i < plate.ask.items.length; i++) {
    plate.ask.items[i].amount += plate.ask.items[i - 1].amount
  }
  for (let i = 1; i < plate.bid.items.length; i++) {
    plate.bid.items[i].amount += plate.bid.items[i - 1].amount
  }
}

const initPoints = (plate) => {
  const buyItems = plate.bid.items
  const sellItems = plate.ask.items
  const bounds = getChartBounds()

  if (!buyItems.length && !sellItems.length) {
    buyPoints.value = []
    sellPoints.value = []
    return
  }

  xObj.value.min = plate.bid.lowestPrice ?? buyItems[0]?.price ?? 0
  xObj.value.max = plate.ask.highestPrice ?? sellItems[sellItems.length - 1]?.price ?? xObj.value.min
  yObj.value.min = 0

  const bidValue = buyItems.length ? buyItems[buyItems.length - 1].amount : 0
  const askValue = sellItems.length ? sellItems[sellItems.length - 1].amount : 0
  yObj.value.max = Math.max(bidValue, askValue, 1) * 1.1

  const yScale = bounds.height / Math.max(yObj.value.max - yObj.value.min, 1)
  const halfWidth = bounds.width / 2

  buyPoints.value = []
  if (buyItems.length) {
    const buyScale = halfWidth / Math.max(buyItems.length - 1, 1)
    for (let i = buyItems.length - 1; i >= 0; i--) {
      const x = bounds.left + (buyItems.length - 1 - i) * buyScale
      const y = Math.floor(bounds.bottom - (buyItems[i].amount - yObj.value.min) * yScale)
      buyPoints.value.push({
        x,
        y,
        amount: buyItems[i].amount,
        price: buyItems[i].price
      })
    }
  }

  sellPoints.value = []
  const sellArr = sellItems.length
    ? [{ price: plate.bid.highestPrice ?? buyItems[0]?.price ?? 0, amount: 0 }, ...sellItems]
    : []

  if (sellArr.length) {
    const sellScale = halfWidth / Math.max(sellArr.length - 1, 1)
    for (let i = 0; i < sellArr.length; i++) {
      const x = bounds.middle + i * sellScale
      const y = Math.floor(bounds.bottom - (sellArr[i].amount - yObj.value.min) * yScale)
      sellPoints.value.push({
        x,
        y,
        amount: sellArr[i].amount,
        price: sellArr[i].price
      })
    }
  }
}

const drawBuy = (plate) => {
  if (!contextRef.value || !buyPoints.value.length) return
  const bounds = getChartBounds()

  contextRef.value.beginPath()
  contextRef.value.moveTo(bounds.left, bounds.bottom)
  for (const point of buyPoints.value) {
    contextRef.value.lineTo(point.x, point.y)
  }
  contextRef.value.lineTo(bounds.middle, bounds.bottom)
  contextRef.value.fillStyle = plate.skin === 'day' ? 'rgba(0,178,117,.18)' : 'rgba(0, 255, 136, 0.18)'
  contextRef.value.fill()

  contextRef.value.beginPath()
  for (const [index, point] of buyPoints.value.entries()) {
    if (index === 0) {
      contextRef.value.moveTo(point.x, point.y)
    } else {
      contextRef.value.lineTo(point.x, point.y)
    }
  }
  contextRef.value.strokeStyle = plate.skin === 'day' ? '#00b275' : '#00ff88'
  contextRef.value.lineWidth = 2
  contextRef.value.stroke()
}

const drawSell = (plate) => {
  if (!contextRef.value || !sellPoints.value.length) return
  const bounds = getChartBounds()
  const lastY = sellPoints.value[sellPoints.value.length - 1]?.y ?? bounds.bottom

  contextRef.value.beginPath()
  contextRef.value.moveTo(bounds.middle, bounds.bottom)
  for (const point of sellPoints.value) {
    contextRef.value.lineTo(point.x, point.y)
  }
  contextRef.value.lineTo(bounds.right, lastY)
  contextRef.value.lineTo(bounds.right, bounds.bottom)
  contextRef.value.fillStyle = plate.skin === 'day' ? 'rgba(241,80,87,.18)' : 'rgba(255, 95, 81, 0.18)'
  contextRef.value.fill()

  contextRef.value.beginPath()
  for (const [index, point] of sellPoints.value.entries()) {
    if (index === 0) {
      contextRef.value.moveTo(point.x, point.y)
    } else {
      contextRef.value.lineTo(point.x, point.y)
    }
  }
  contextRef.value.strokeStyle = plate.skin === 'day' ? '#f15057' : '#ff5f51'
  contextRef.value.lineWidth = 2
  contextRef.value.stroke()
}

const buildTickIndexes = (length) => {
  if (length <= 1) return [0]
  const indexes = [0, Math.floor((length - 1) / 2), length - 1]
  return [...new Set(indexes)].filter(index => index >= 0)
}

const drawAxis = (plate) => {
  if (!contextRef.value) return
  const bounds = getChartBounds()
  const axisColor = plate.skin === 'day' ? '#8792a5' : '#61688a'
  const priceUnit = props.values?.priceUnit || '--'
  const amountUnit = props.values?.amountUnit || '--'

  contextRef.value.save()
  contextRef.value.strokeStyle = axisColor
  contextRef.value.lineWidth = 1
  contextRef.value.fillStyle = axisColor
  contextRef.value.font = '11px Microsoft YaHei'

  contextRef.value.beginPath()
  contextRef.value.moveTo(bounds.left, bounds.bottom)
  contextRef.value.lineTo(bounds.right, bounds.bottom)
  contextRef.value.lineTo(bounds.right, bounds.top)
  contextRef.value.stroke()

  contextRef.value.setLineDash([4, 4])
  contextRef.value.beginPath()
  contextRef.value.moveTo(bounds.middle, bounds.top)
  contextRef.value.lineTo(bounds.middle, bounds.bottom)
  contextRef.value.stroke()
  contextRef.value.setLineDash([])

  for (const index of buildTickIndexes(buyPoints.value.length)) {
    const point = buyPoints.value[index]
    if (!point) continue
    contextRef.value.beginPath()
    contextRef.value.moveTo(point.x, bounds.bottom)
    contextRef.value.lineTo(point.x, bounds.bottom + 4)
    contextRef.value.stroke()
    contextRef.value.fillText(
      formatDepthValue(point.price, plate.priceScale),
      Math.max(bounds.left, point.x - 18),
      bounds.bottom + 18
    )
  }

  for (const index of buildTickIndexes(sellPoints.value.length)) {
    const point = sellPoints.value[index]
    if (!point) continue
    contextRef.value.beginPath()
    contextRef.value.moveTo(point.x, bounds.bottom)
    contextRef.value.lineTo(point.x, bounds.bottom + 4)
    contextRef.value.stroke()
    contextRef.value.fillText(
      formatDepthValue(point.price, plate.priceScale),
      Math.min(bounds.right - 42, point.x - 18),
      bounds.bottom + 18
    )
  }

  const yTickCount = 4
  const yStep = (bounds.bottom - bounds.top) / yTickCount
  for (let i = 1; i <= yTickCount; i++) {
    const y = bounds.bottom - yStep * i
    const amount = yObj.value.max * (i / yTickCount)
    contextRef.value.beginPath()
    contextRef.value.moveTo(bounds.right, y)
    contextRef.value.lineTo(bounds.right + 4, y)
    contextRef.value.stroke()
    contextRef.value.fillText(
      formatDepthValue(amount, plate.amountScale),
      bounds.right + 8,
      y + 4
    )
  }

  contextRef.value.font = '12px Microsoft YaHei'
  contextRef.value.fillText(`价格 (${priceUnit})`, bounds.middle - 28, pHeight.value - 10)
  contextRef.value.fillText(`累计数量 (${amountUnit})`, bounds.right - 4, bounds.top - 4)
  contextRef.value.restore()
}

const drawLabels = (plate) => {
  if (!contextRef.value) return
  const bounds = getChartBounds()

  contextRef.value.save()
  contextRef.value.font = 'bold 13px Microsoft YaHei'
  contextRef.value.fillStyle = plate.skin === 'day' ? '#00b275' : '#00ff88'
  contextRef.value.fillText('买盘 Bid', bounds.left, bounds.top - 2)
  contextRef.value.fillStyle = plate.skin === 'day' ? '#f15057' : '#ff5f51'
  contextRef.value.fillText('卖盘 Ask', bounds.right - 64, bounds.top - 2)
  contextRef.value.restore()
}

const getNearestPoint = (points, x) => {
  if (!points.length) return null
  let nearestPoint = null
  let minDistance = Infinity

  for (const point of points) {
    const distance = Math.abs(point.x - x)
    if (distance < minDistance) {
      minDistance = distance
      nearestPoint = point
    }
  }

  return minDistance <= 24 ? nearestPoint : null
}

const renderHoverMarker = (point) => {
  clearHoverCanvas()
  if (!coverContextRef.value || !point) return

  coverContextRef.value.beginPath()
  coverContextRef.value.arc(point.x, point.y, 9, 0, 2 * Math.PI)
  coverContextRef.value.fillStyle = 'rgba(53, 64, 103, 0.5)'
  coverContextRef.value.fill()
  coverContextRef.value.closePath()

  coverContextRef.value.beginPath()
  coverContextRef.value.arc(point.x, point.y, 4, 0, 2 * Math.PI)
  coverContextRef.value.fillStyle = '#7a98f7'
  coverContextRef.value.fill()
  coverContextRef.value.closePath()
}

const initHoverEvent = () => {
  if (!coverCanvasRef.value) return

  coverCanvasRef.value.onmousemove = (event) => {
    const bounds = getChartBounds()
    const isBuySide = event.offsetX <= bounds.middle
    const points = isBuySide ? buyPoints.value : sellPoints.value
    const point = getNearestPoint(points, event.offsetX)

    if (!point) {
      clearHover()
      return
    }

    renderHoverMarker(point)

    hoverState.visible = true
    hoverState.side = isBuySide ? 'buy' : 'sell'
    hoverState.sideLabel = isBuySide ? '买盘 Bid' : '卖盘 Ask'
    hoverState.priceUnit = props.values?.priceUnit || '--'
    hoverState.amountUnit = props.values?.amountUnit || '--'
    hoverState.priceText = formatDepthValue(point.price, getPriceScale())
    hoverState.amountText = formatDepthValue(point.amount, getAmountScale())
    hoverState.left = Math.min(event.offsetX + 12, Math.max(12, pWidth.value - 188))
    hoverState.top = Math.max(12, Math.min(point.y - 84, pHeight.value - 90))
  }

  coverCanvasRef.value.onmouseleave = clearHover
}

const draw = async (inputPlate = props.values) => {
  await nextTick()
  if (!resetCanvas()) return

  const plate = {
    ...safePlate(),
    ...inputPlate,
    bid: {
      ...safePlate().bid,
      ...(inputPlate?.bid ?? {})
    },
    ask: {
      ...safePlate().ask,
      ...(inputPlate?.ask ?? {})
    },
    priceUnit: inputPlate?.priceUnit ?? props.values?.priceUnit ?? '',
    amountUnit: inputPlate?.amountUnit ?? props.values?.amountUnit ?? '',
    priceScale: Number(inputPlate?.priceScale ?? props.values?.priceScale ?? 4) || 4,
    amountScale: Number(inputPlate?.amountScale ?? props.values?.amountScale ?? 4) || 4
  }

  plate.bid.items = Array.isArray(plate.bid.items) ? plate.bid.items.slice(0, 100).map(item => ({ ...item })) : []
  plate.ask.items = Array.isArray(plate.ask.items) ? plate.ask.items.slice(0, 100).map(item => ({ ...item })) : []

  clearHover()
  contextRef.value?.clearRect(0, 0, pWidth.value, pHeight.value)

  if (!plate.bid.items.length && !plate.ask.items.length) {
    return
  }

  convertTotal(plate)
  initPoints(plate)
  drawBuy(plate)
  drawSell(plate)
  drawAxis(plate)
  drawLabels(plate)
}

watch(
  () => props.values,
  value => {
    if (value) {
      draw(value)
    }
  },
  { deep: true }
)

onMounted(async () => {
  await nextTick()
  if (!resetCanvas()) return
  initHoverEvent()
  draw(props.values)
})

defineExpose({
  draw
})
</script>

<style scoped>
.depth-chart {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.depth-chart__canvas {
  position: absolute;
  left: 0;
  top: 0;
}

.depth-chart__tooltip {
  position: absolute;
  min-width: 172px;
  padding: 10px 12px;
  border: 1px solid rgba(122, 152, 247, 0.3);
  border-radius: 6px;
  background: rgba(18, 27, 40, 0.94);
  color: #d7deed;
  pointer-events: none;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.28);
}

.depth-chart__tooltip-title {
  margin-bottom: 6px;
  font-size: 13px;
  font-weight: 700;
}

.depth-chart__tooltip-title--buy {
  color: #00ff88;
}

.depth-chart__tooltip-title--sell {
  color: #ff5f51;
}

.depth-chart__tooltip-row {
  font-size: 12px;
  line-height: 1.6;
}
</style>

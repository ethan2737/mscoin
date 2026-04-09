<template>
  <div ref="containerRef" class="depth-chart">
    <canvas ref="canvasRef" class="depth-chart__canvas"></canvas>
    <canvas ref="coverCanvasRef" class="depth-chart__canvas"></canvas>
  </div>
</template>

<script setup>
import { nextTick, onMounted, ref, watch } from 'vue'

const props = defineProps({
  values: {
    type: Object,
    default: () => ({
      bid: { items: [], highestPrice: 0, lowestPrice: 0 },
      ask: { items: [], highestPrice: 0, lowestPrice: 0 },
      skin: 'night'
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

const safePlate = () => {
  const bid = props.values?.bid ?? { items: [], highestPrice: 0, lowestPrice: 0 }
  const ask = props.values?.ask ?? { items: [], highestPrice: 0, lowestPrice: 0 }
  return {
    bid: {
      highestPrice: bid.highestPrice ?? 0,
      lowestPrice: bid.lowestPrice ?? 0,
      items: Array.isArray(bid.items) ? bid.items.map((item) => ({ ...item })) : []
    },
    ask: {
      highestPrice: ask.highestPrice ?? 0,
      lowestPrice: ask.lowestPrice ?? 0,
      items: Array.isArray(ask.items) ? ask.items.map((item) => ({ ...item })) : []
    },
    skin: props.values?.skin ?? 'night'
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

const initPoints = (plate) => {
  const buy = plate.bid
  const sell = plate.ask
  const buyItems = buy.items
  const sellItems = sell.items

  if (!buyItems.length && !sellItems.length) {
    buyPoints.value = []
    sellPoints.value = []
    return
  }

  xObj.value.min = buy.lowestPrice ?? buyItems[0]?.price ?? 0
  xObj.value.max = sell.highestPrice ?? sellItems[sellItems.length - 1]?.price ?? xObj.value.min
  yObj.value.min = 0

  const bidValue = buyItems.length ? buyItems[buyItems.length - 1].amount : 0
  const askValue = sellItems.length ? sellItems[sellItems.length - 1].amount : 0
  yObj.value.max = Math.max(bidValue, askValue, 1) * 1.2

  const usableWidth = Math.max(pWidth.value - 50, 1)
  const usableHeight = Math.max(pHeight.value - 50, 1)
  const yScale = usableHeight / Math.max(yObj.value.max - yObj.value.min, 1)

  buyPoints.value = []
  if (buyItems.length) {
    const buyScale = usableWidth / 2 / Math.max(buyItems.length - 1, 1)
    for (let i = buyItems.length - 1; i >= 0; i--) {
      const x = (buyItems.length - 1 - i) * buyScale
      const y = Math.floor(usableHeight - (buyItems[i].amount - yObj.value.min) * yScale)
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
    ? [{ price: buy.highestPrice ?? buyItems[0]?.price ?? 0, amount: 0 }, ...sellItems]
    : []
  if (sellArr.length) {
    const sellScale = usableWidth / 2 / Math.max(sellArr.length - 1, 1)
    for (let i = 0; i < sellArr.length; i++) {
      const x = i * sellScale + usableWidth / 2
      const y = Math.floor(usableHeight - (sellArr[i].amount - yObj.value.min) * yScale)
      sellPoints.value.push({
        x,
        y,
        amount: sellArr[i].amount,
        price: sellArr[i].price
      })
    }
  }
}

const convertTotal = (plate) => {
  for (let i = 1; i < plate.ask.items.length; i++) {
    plate.ask.items[i].amount += plate.ask.items[i - 1].amount
  }
  for (let i = 1; i < plate.bid.items.length; i++) {
    plate.bid.items[i].amount += plate.bid.items[i - 1].amount
  }
}

const drawBuy = (plate) => {
  if (!contextRef.value || !buyPoints.value.length) return
  contextRef.value.beginPath()
  contextRef.value.moveTo(0, pHeight.value)
  for (const point of buyPoints.value) {
    contextRef.value.lineTo(point.x, point.y)
  }
  contextRef.value.lineTo((pWidth.value - 50) / 2, pHeight.value)
  contextRef.value.fillStyle = plate.skin === 'day' ? 'rgba(0,178,117,.5)' : '#243235'
  contextRef.value.fill()
}

const drawSell = (plate) => {
  if (!contextRef.value || !sellPoints.value.length) return
  contextRef.value.beginPath()
  contextRef.value.moveTo((pWidth.value - 50) / 2, pHeight.value)
  for (const point of sellPoints.value) {
    contextRef.value.lineTo(point.x, point.y)
  }
  const lastY = sellPoints.value[sellPoints.value.length - 1]?.y ?? pHeight.value
  contextRef.value.lineTo(pWidth.value - 5, lastY)
  contextRef.value.lineTo(pWidth.value - 5, pHeight.value)
  contextRef.value.fillStyle = plate.skin === 'day' ? 'rgba(241,80,87,.5)' : '#392231'
  contextRef.value.fill()
}

const getPoint = (points, x) => points.find((point) => Math.abs(point.x - x) <= 2) ?? null

const initHoverEvent = () => {
  if (!coverCanvasRef.value || !coverContextRef.value) return
  coverCanvasRef.value.onmousemove = (event) => {
    const points = event.offsetX <= (pWidth.value - 50) / 2 ? buyPoints.value : sellPoints.value
    const point = getPoint(points, event.offsetX)
    coverCanvasRef.value.height = pHeight.value
    if (!point) return

    coverContextRef.value.beginPath()
    coverContextRef.value.arc(event.offsetX, point.y, 10, 0, 2 * Math.PI)
    coverContextRef.value.fillStyle = '#354067'
    coverContextRef.value.fill()
    coverContextRef.value.closePath()

    coverContextRef.value.beginPath()
    coverContextRef.value.arc(event.offsetX, point.y, 5, 0, 2 * Math.PI)
    coverContextRef.value.fillStyle = '#7a98f7'
    coverContextRef.value.fill()
    coverContextRef.value.closePath()
  }
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
    }
  }

  plate.bid.items = Array.isArray(plate.bid.items) ? plate.bid.items.slice(0, 100).map((item) => ({ ...item })) : []
  plate.ask.items = Array.isArray(plate.ask.items) ? plate.ask.items.slice(0, 100).map((item) => ({ ...item })) : []

  if (!plate.bid.items.length && !plate.ask.items.length) {
    return
  }

  convertTotal(plate)
  initPoints(plate)
  canvasRef.value.height = pHeight.value
  drawBuy(plate)
  drawSell(plate)
}

watch(
  () => props.values,
  (value) => {
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
}

.depth-chart__canvas {
  position: absolute;
  left: 0;
  top: 0;
}
</style>

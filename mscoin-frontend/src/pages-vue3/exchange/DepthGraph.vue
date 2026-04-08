<template>
  <div id="depth_chart" style="position: relative;width: 100%;height: 100%;">
    <canvas id="depthGraph" style="position:absolute;left:0"></canvas>
    <canvas id="depthGraph_cover" style="position:absolute;left:0"></canvas>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 深度图组件
 * 迁移点：
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. Canvas API 保持不变
 * 4. props 使用 defineProps
 * 5. emit 使用 defineEmits
 */
import { ref, onMounted, watch } from 'vue'

// Props
const props = defineProps({
  values: {
    type: Object,
    required: true
  }
})

// Emits
const emit = defineEmits(['update:values'])

// 响应式数据
const xObj = ref({})
const yObj = ref({})
const canvas = ref(null)
const context = ref(null)
const canvas_cover = ref(null)
const context_cover = ref(null)
const pWidth = ref(900)
const pHeight = ref(300)
const buyPoints = ref([])
const sellPoints = ref([])

// 初始化 Canvas
const initCanvas = () => {
  canvas.value = document.getElementById('depthGraph')
  canvas_cover.value = document.getElementById('depthGraph_cover')
  canvas.value.width = canvas_cover.value.width = pWidth.value = Math.floor(
    document.getElementById('depth_chart').offsetWidth
  )
  canvas.value.height = canvas_cover.value.height = pHeight.value = Math.floor(
    document.getElementById('depth_chart').offsetHeight
  )
  context.value = canvas.value.getContext('2d')
  context_cover.value = canvas_cover.value.getContext('2d')
  initHoverEvent()
}

// 初始化点位
const initPoints = () => {
  const buy = props.values.bid
  const sell = props.values.ask

  if (buy !== undefined && sell !== undefined) {
    xObj.value.min = buy.lowestPrice
    xObj.value.max = sell.highestPrice
    yObj.value.min = 0

    let bidValue = null
    if (buy.items.length >= 1) {
      bidValue = buy.items[buy.items.length - 1].amount
    }
    if (buy.items.length === 0) {
      bidValue = 0
    }

    let askValue = null
    if (sell.items.length >= 1) {
      askValue = sell.items[sell.items.length - 1].amount
    }
    if (sell.items.length === 0) {
      askValue = 0
    }

    yObj.value.max = Math.max(bidValue, askValue)
    yObj.value.max = yObj.value.max * 1.2 // y 轴最大值略大于最大委托量

    const buyArr = buy.items
    const xScale = Math.floor((pWidth.value - 50) / (2 * (buyArr.length - 1))) // 买入 x 轴等分
    const yScale = parseFloat(
      (pHeight.value - 50) / (yObj.value.max - yObj.value.min)
    ).toFixed(4)

    buyPoints.value = []
    for (let i = buyArr.length - 1; i >= 0; i--) {
      const x = (buyArr.length - 1 - i) * xScale
      const y = Math.floor(
        pHeight.value - 50 - (buyArr[i].amount - yObj.value.min) * yScale
      )

      buyPoints.value.push({
        x: x,
        y: y,
        amount: buyArr[i].amount,
        price: buyArr[i].price
      })
    }

    sellPoints.value = []
    const sellArr = sell.items
    sellArr.unshift({ price: buy.highestPrice, amount: 0 })

    const sellXScale = Math.floor((pWidth.value - 50) / (2 * (sellArr.length - 1)))
    for (let i = 0; i < sellArr.length; i++) {
      const x = i * sellXScale + Math.floor((pWidth.value - 50) / 2)
      const y = Math.floor(
        pHeight.value - 50 - (sellArr[i].amount - yObj.value.min) * yScale
      )
      sellPoints.value.push({
        x: x,
        y: y,
        amount: sellArr[i].amount,
        price: sellArr[i].price
      })
    }
  }
}

// 转换总额
const convertTotal = () => {
  for (let i = 1; i < props.values.ask.items.length; i++) {
    props.values.ask.items[i].amount += props.values.ask.items[i - 1].amount
  }

  for (let i = 1; i < props.values.bid.items.length; i++) {
    props.values.bid.items[i].amount += props.values.bid.items[i - 1].amount
  }
}

// 绘制深度图
const draw = (plate) => {
  // 限制数据量
  if (plate.bid.items.length > 100) {
    plate.bid.items = plate.bid.items.slice(0, 100)
  }
  if (plate.ask.items.length > 100) {
    plate.ask.items = plate.ask.items.slice(0, 100)
  }

  convertTotal()
  initPoints()
  canvas.value.height = pHeight.value
  drawBuy()
  drawSell()
}

// 绘制买入深度
const drawBuy = () => {
  context.value.beginPath()
  context.value.moveTo(0, pHeight.value - 0)
  for (let i = 0; i < buyPoints.value.length; i++) {
    context.value.lineTo(buyPoints.value[i].x, buyPoints.value[i].y)
  }
  context.value.lineTo((pWidth.value - 50) / 2, pHeight.value - 0)

  if (props.values.skin === 'day') {
    context.value.fillStyle = 'rgba(0,178,117,.5)'
  } else {
    context.value.fillStyle = '#243235'
  }
  context.value.fill()
}

// 绘制卖出深度
const drawSell = () => {
  context.value.beginPath()
  context.value.moveTo((pWidth.value - 50) / 2, pHeight.value - 0)
  for (let i = 0; i < sellPoints.value.length; i++) {
    context.value.lineTo(sellPoints.value[i].x, sellPoints.value[i].y)
  }
  context.value.lineTo(
    pWidth.value - 5,
    sellPoints.value[sellPoints.value.length - 1].y
  )
  context.value.lineTo(pWidth.value - 5, pHeight.value - 0)

  if (props.values.skin === 'day') {
    context.value.fillStyle = 'rgba(241,80,87,.5)'
  } else {
    context.value.fillStyle = '#392231'
  }
  context.value.fill()
}

// 绘制坐标轴
const drawAxis = () => {
  context.value.beginPath()
  context.value.moveTo(0, pHeight.value - 50)
  context.value.lineTo(pWidth.value - 50, pHeight.value - 50)
  context.value.lineTo(pWidth.value - 50, 0)
  context.value.strokeStyle = '#61688a'
  context.value.stroke()
  context.value.closePath()

  context.value.beginPath()
  context.value.moveTo(pWidth.value - 50, pHeight.value - 50)
  context.value.lineTo(pWidth.value - 50 + 5, pHeight.value - 50)
  context.value.strokeStyle = '#61688a'
  context.value.stroke()
  context.value.closePath()

  context.value.fillStyle = '#61688a'
  context.value.font = '12px Adobe Ming Std'
  context.value.fillText('0', pWidth.value - 50 + 10, pHeight.value - 50 + 5)

  if (buyPoints.value.length > 50) {
    buyPoints.value.length = 50
  }

  for (let i = 0; i < buyPoints.value.length; i++) {
    if (i % 10) {
      context.value.beginPath()
      context.value.moveTo(buyPoints.value[i].x, pHeight.value - 50)
      context.value.lineTo(buyPoints.value[i].x, pHeight.value - 45)
      context.value.strokeStyle = '#61688a'
      context.value.stroke()
      context.value.closePath()

      context.value.fillStyle = '#61688a'
      context.value.font = '12px Adobe Ming Std'
      context.value.fillText(
        buyPoints.value[i].price,
        buyPoints.value[i].x - 10,
        pHeight.value - 30
      )
    }
  }

  for (let i = 0; i < sellPoints.value.length; i++) {
    if (i % 10) {
      context.value.beginPath()
      context.value.moveTo(sellPoints.value[i].x, pHeight.value - 50)
      context.value.lineTo(sellPoints.value[i].x, pHeight.value - 45)
      context.value.strokeStyle = '#61688a'
      context.value.stroke()
      context.value.closePath()

      context.value.fillStyle = '#61688a'
      context.value.font = '12px Adobe Ming Std'
      context.value.fillText(
        sellPoints.value[i].price,
        sellPoints.value[i].x - 10,
        pHeight.value - 30
      )
    }
  }

  const yInterval = Math.floor((pHeight.value - 100) / 5)
  for (let i = 1; i <= 5; i++) {
    context.value.beginPath()
    context.value.moveTo(
      pWidth.value - 50,
      pHeight.value - 50 - i * yInterval
    )
    context.value.lineTo(
      pWidth.value - 50 + 5,
      pHeight.value - 50 - i * yInterval
    )
    context.value.strokeStyle = '#61688a'
    context.value.stroke()
    context.value.closePath()

    context.value.fillStyle = '#61688a'
    context.value.font = '12px Adobe Ming Std'
    const yVal = yObj.value.max * i / 6
    context.value.fillText(
      yVal,
      pWidth.value - 50 + 10,
      pHeight.value - 50 - i * yInterval + 5
    )
  }
}

// 初始化悬停事件
const initHoverEvent = () => {
  const opts = {
    pWidth: pWidth.value,
    pHeight: pHeight.value,
    buyPoints: buyPoints.value,
    sellPoints: sellPoints.value,
    canvas_cover: canvas_cover.value,
    context_cover: context_cover.value
  }

  canvas_cover.value.addEventListener('mousemove', function(e) {
    let arr = 0
    if (e.offsetX <= (opts.pWidth - 50) / 2) {
      arr = getPoint(opts.buyPoints, e.offsetX)
    } else {
      arr = getPoint(opts.sellPoints, e.offsetX)
    }

    if (arr !== null) {
      // 清空画板
      opts.canvas_cover.height = 500

      // 画圆
      opts.context_cover.beginPath()
      opts.context_cover.arc(e.offsetX, arr.y, 10, 0, 2 * Math.PI)
      opts.context_cover.fillStyle = '#354067'
      opts.context_cover.fill()
      opts.context_cover.closePath()

      opts.context_cover.beginPath()
      opts.context_cover.arc(e.offsetX, arr.y, 5, 0, 2 * Math.PI)
      opts.context_cover.fillStyle = '#7a98f7'
      opts.context_cover.fill()
      opts.context_cover.closePath()

      // 画矩形框
      let rx = 0
      if (e.offsetX <= 155) {
        rx = 5
      } else {
        rx = e.offsetX - 150
      }
      let ry = 0
      if (arr.y <= 100) {
        ry = arr.y + 20
      } else if (arr.y + 100 > opts.pHeight - 50) {
        ry = arr.y - 100
      } else {
        ry = arr.y - 100
      }

      opts.context_cover.beginPath()
      opts.context_cover.rect(rx, ry, 180, 80)
      opts.context_cover.fillStyle = '#262a42'
      opts.context_cover.fill()
      opts.context_cover.closePath()

      // 填充文本
      opts.context_cover.beginPath()
      opts.context_cover.fillStyle = '#ddd'
      opts.context_cover.font = '14px Microsoft YaHei'
      opts.context_cover.fillText('委托价', rx + 20, ry + 30)
      opts.context_cover.fillText(arr.price, rx + 70, ry + 30)
      opts.context_cover.fillText('累计', rx + 20, ry + 60)
      opts.context_cover.fillText(arr.amount.toFixed(8), rx + 70, ry + 60)
      opts.context_cover.closePath()
    }
  }, false)
}

// 获取点位
const getPoint = (arr, x) => {
  for (let i = 0; i < arr.length; i++) {
    if (x === arr[i].x) {
      return arr[i]
    }
  }
  return null
}

// 监听 props.values 变化
watch(() => props.values, (newValue, oldValue) => {
  initPoints()
  if (newValue) {
    draw(newValue)
  }
}, { deep: true })

// 生命周期
onMounted(() => {
  initCanvas()
})

// 暴露 draw 方法给父组件
defineExpose({
  draw
})
</script>

<style scoped>
</style>

<template>
  <div>
    <p>
      <div class="l-title" v-if="state === 0">{{ $t('exchange.publishcounttxt0') }}</div>
      <div class="l-title" v-if="state === 1">{{ $t('exchange.publishcounttxt1') }}</div>
      <div class="l-title" v-if="state === 2">{{ $t('exchange.publishcounttxt2') }}</div>
      <div class="l-title" v-if="state === 3">{{ $t('exchange.publishcounttxt3') }}</div>
    </p>
    <p style="margin-top: 15px;">
      <span class="num">{{ day }}</span><span class="txt">{{ $t('exchange.dateTimeDay') }}</span>
      <span class="num">{{ hour }}</span><span class="txt">{{ $t('exchange.dateTimeHour') }}</span>
      <span class="num">{{ minutes }}</span><span class="txt">{{ $t('exchange.dateTimeMinutes') }}</span>
      <span class="num">{{ seconds }}</span><span class="txt">{{ $t('exchange.dateTimeSeconds') }}</span>
    </p>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 倒计时组件
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. props 使用 defineProps
 * 4. emit 使用 defineEmits
 * 5. 定时器清理逻辑保持
 */
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'

// Props
const props = defineProps({
  countDownBgColor: String,
  publishState: Number,
  publishType: String,
  currentTime: Number,
  startTime: String,
  endTime: String,
  clearTime: String,
  showPublishMask: Function,
  hidePublishMask: Function,
  hideCountDown: Function
})

// Emits
const emit = defineEmits(['update:countDownBgColor', 'update:publishState'])

// 响应式数据
const currentT = ref(0)
const day = ref('00')
const hour = ref('00')
const minutes = ref('00')
const seconds = ref('00')
const timer = ref(null)
const flag = ref(false)
const state = ref(0) // 0:活动未开始 1：活动已开始但未结束 2：活动已结束但未清盘结束 3：活动已清盘结束

// 时间倒计时逻辑
const timeDown = () => {
  let endTime = 0

  const nowTime = parseInt(currentT.value + props.currentTime)
  const startT = parseInt(new Date(props.startTime).getTime() / 1000)
  const endT = parseInt(new Date(props.endTime).getTime() / 1000)
  const clearT = parseInt(new Date(props.clearTime).getTime() / 1000)

  if (nowTime <= clearT) {
    state.value = 2
    endTime = clearT
  }
  if (nowTime <= endT) {
    state.value = 1
    endTime = endT
  }
  if (nowTime <= startT) {
    state.value = 0
    endTime = startT
  }

  const leftTime = endTime - nowTime
  day.value = formate(parseInt(leftTime / (24 * 60 * 60)))
  hour.value = formate(parseInt(leftTime / (60 * 60) % 24))
  minutes.value = formate(parseInt(leftTime / 60 % 60))
  seconds.value = formate(parseInt(leftTime % 60))

  if (state.value === 0) {
    // 活动未开始
    if (props.showPublishMask) props.showPublishMask()
    emit('update:countDownBgColor', '#003478')
    emit('update:publishState', 0)
  }
  if (state.value === 1) {
    // 活动开始中
    if (props.hidePublishMask) props.hidePublishMask()
    emit('update:countDownBgColor', '#094802')
    emit('update:publishState', 1)
  }
  if (state.value === 2) {
    // 清盘中
    if (props.showPublishMask) props.showPublishMask()
    emit('update:countDownBgColor', '#5b0000')
    emit('update:publishState', 2)
  }
  if (state.value === 2 && leftTime <= 0) {
    // 清盘结束，正常交易
    if (props.hideCountDown) props.hideCountDown()
    if (props.hidePublishMask) props.hidePublishMask()
    flag.value = true
    state.value = 3
    emit('update:publishState', 3)
  }
  currentT.value = currentT.value + 1
}

// 格式化时间
const formate = (time) => {
  if (time >= 10) {
    return time
  } else {
    return `0${time}`
  }
}

// 监听 props 变化
watch(() => [props.startTime, props.endTime, props.clearTime, props.currentTime], () => {
  timeDown()
})

// 生命周期
onMounted(() => {
  timer.value = setInterval(() => {
    if (flag.value === true) {
      clearInterval(timer.value)
    } else {
      timeDown()
    }
  }, 1000)
})

onBeforeUnmount(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }
})
</script>

<style scoped>
.num {
  font-size: 20px;
  background: linear-gradient(0deg, #df9000, #ffb100);
  padding: 2px 4px;
  border-radius: 2px;
  color: #000;
  text-shadow: 1px 1px 1px #908e8e;
}

.txt {
  font-size: 12px;
}
</style>

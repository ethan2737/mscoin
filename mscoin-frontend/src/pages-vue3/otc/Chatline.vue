<template>
  <section class="chat-card">
    <div class="chat-card__header">
      <div class="chat-card__title">
        <h3>{{ text.orderChat }}</h3>
        <p>{{ text.orderChatDesc }}</p>
      </div>
      <button
        v-show="showHistoryButton"
        class="history-button"
        @click="getBefore"
      >
        <el-icon class="clock"><Clock /></el-icon>
        <span>{{ $t('otc.chatline.loadmore') }}</span>
      </button>
    </div>

    <div class="chat-card__banner">
      <div class="chat-card__system">
        <p class="system-box" v-show="msg.status === 1">{{ $t('otc.chatline.status_1') }}</p>
        <p class="system-box" v-show="msg.status === 2">{{ $t('otc.chatline.status_2') }}</p>
        <p class="system-box" v-show="msg.status === 3">{{ $t('otc.chatline.status_3') }}</p>
        <p class="system-box" v-show="msg.status === 4">{{ $t('otc.chatline.status_4') }}</p>
        <p class="system-box" v-show="msg.status === 0">{{ $t('otc.chatline.status_5') }}</p>
      </div>

      <div class="chat-card__notice">
        <el-icon><Warning /></el-icon>
        <span>{{ $t('otc.chatline.warning') }}</span>
      </div>
    </div>

    <div class="chat-card__messages" id="scrollChat">
      <template v-for="item in msgLists" :key="item.id || item.sendTimeStr || item.timeNow">
        <div class="message-row" v-show="item.uidFrom !== msg.myId">
          <div class="message-row__avatar">
            <div class="user-face user-avatar-public">
              <span class="user-avatar-in" v-if="!item.fromAvatar">{{ msgnameS }}</span>
              <img v-else class="chat-avatar-image" :src="item.fromAvatar" />
            </div>
          </div>
          <div class="message-row__body">
            <p class="message-row__name">{{ msg.otherSide }}</p>
            <div class="message-bubble">
              <span>{{ item.content }}</span>
            </div>
            <span class="message-row__time">{{ item.sendTimeStr }}</span>
          </div>
        </div>

        <div class="message-row message-row--self" v-show="item.uidFrom === msg.myId">
          <div class="message-row__body">
            <p class="message-row__name">{{ user?.username }}</p>
            <div class="message-bubble message-bubble--self">
              <span>{{ item.content }}</span>
            </div>
            <span class="message-row__time">{{ item.timeNow || item.sendTimeStr }}</span>
          </div>
          <div class="message-row__avatar">
            <div class="user-face user-avatar-public">
              <span class="user-avatar-in" v-if="!item.fromAvatar">{{ usernameS }}</span>
              <img v-else class="chat-avatar-image" :src="item.fromAvatar" />
            </div>
          </div>
        </div>
      </template>
    </div>

    <div class="audio-wrap">
      <audio id="noticeMusic" :src="audioSrc">{{ text.audioUnsupported }}</audio>
    </div>

    <div class="chat-card__composer">
      <el-input
        v-model="mytext"
        @keyup.enter="sendName"
        autocomplete="off"
        :placeholder="text.messagePlaceholder"
        class="msg-input"
      />
      <button class="send-btn" @click="sendName">
        <el-icon :size="20" color="#0f1a26"><Position /></el-icon>
      </button>
    </div>

    <div class="chat-card__toggles">
      <el-checkbox v-model="fOpenNotice" @change="handleNoticeClick">{{ text.desktopNotice }}</el-checkbox>
      <el-checkbox v-model="fOpenAudio" @change="handleAudioClick">{{ text.audioAlert }}</el-checkbox>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount, watch } from 'vue'
import { ElMessage, ElInput, ElCheckbox, ElIcon } from 'element-plus'
import { Clock, Warning, Position } from '@element-plus/icons-vue'
import axios from 'axios'
import io from 'socket.io-client'
import { useI18n } from 'vue-i18n'

const props = defineProps({
  msg: {
    type: Object,
    required: true
  }
})
const { locale } = useI18n()

const host = ''
let socket = null

const audioSrc = ref('/assets/audio/notice.wav')
const fOpenAudio = ref(true)
const fOpenNotice = ref(false)
const currentPage = ref(1)
const totalPage = ref(1)
const showMore = ref(true)
const mytext = ref('')
const msgLists = ref([])

const msgnameS = computed(() => (props.msg?.otherSide + '').slice(0, 1))
const user = computed(() => JSON.parse(localStorage.getItem('MEMBER') || '{}'))
const usernameS = computed(() => (user.value?.username + '').slice(0, 1))
const orderId = computed(() => props.msg?.orderId || '')
const showHistoryButton = computed(() => showMore.value && currentPage.value <= totalPage.value)
const isZh = computed(() => String(locale.value || '').toLowerCase().startsWith('zh'))
const text = computed(() => isZh.value ? {
  orderChat: '订单沟通',
  orderChatDesc: '保持沟通顺畅，按顺序完成付款、确认收款和放行。',
  audioUnsupported: '您的浏览器不支持 audio 标签。',
  messagePlaceholder: '请输入聊天内容...',
  desktopNotice: '开启桌面消息提醒',
  audioAlert: '开启声音消息提醒',
  allowNotice: '请先允许浏览器通知权限。',
  noticeBlocked: '桌面通知已被浏览器拦截。',
  noticeEnabled: '桌面通知已开启。',
  noticeUnsupported: '当前浏览器不支持通知功能。',
  noticeDisabled: '桌面通知已关闭。',
  audioEnabled: '声音提醒已开启。',
  audioDisabled: '声音提醒已关闭。',
  messageRequired: '请输入聊天内容。'
} : {
  orderChat: 'Order Chat',
  orderChatDesc: 'Keep the conversation clear and finish payment, receipt confirmation, and release in sequence.',
  audioUnsupported: 'Your browser does not support audio.',
  messagePlaceholder: 'Type a message...',
  desktopNotice: 'Desktop notice',
  audioAlert: 'Audio alert',
  allowNotice: 'Please allow browser notifications first.',
  noticeBlocked: 'Desktop notifications are blocked in your browser settings.',
  noticeEnabled: 'Desktop notifications enabled.',
  noticeUnsupported: 'This browser does not support notifications.',
  noticeDisabled: 'Desktop notifications disabled.',
  audioEnabled: 'Audio alert enabled.',
  audioDisabled: 'Audio alert disabled.',
  messageRequired: 'Please enter a message.'
})

const scrollToBottom = () => {
  const div = document.getElementById('scrollChat')
  if (div) {
    div.scrollTop = div.scrollHeight
  }
}

const playAudio = () => {
  const audio = document.getElementById('noticeMusic')
  if (audio !== null) {
    audio.play()
  }
}

const handleNoticeClick = () => {
  if (fOpenNotice.value) {
    if (window.Notification) {
      if (Notification.permission === 'default') {
        ElMessage.info(text.value.allowNotice)
        Notification.requestPermission().then((result) => {
          if (result === 'denied') {
            ElMessage.info(text.value.noticeBlocked)
          }
        })
      } else if (Notification.permission === 'denied') {
        ElMessage.info(text.value.noticeBlocked)
      } else {
        ElMessage.info(text.value.noticeEnabled)
      }
    } else {
      ElMessage.info(text.value.noticeUnsupported)
    }
  } else {
    ElMessage.info(text.value.noticeDisabled)
  }
}

const handleAudioClick = () => {
  if (fOpenAudio.value) {
    ElMessage.info(text.value.audioEnabled)
  } else {
    ElMessage.info(text.value.audioDisabled)
  }
}

const connect = () => {
  socket.on('/user/' + props.msg.myId + '/' + orderId.value, (response) => {
    const otheritem = JSON.parse(response)
    msgLists.value.push(otheritem)
    if (fOpenNotice.value && window.Notification && Notification.permission === 'granted') {
      const notification = new Notification(props.msg.otherSide + ':', {
        body: otheritem.content,
        icon: 'https://bihuo-ex.oss-ap-southeast-1.aliyuncs.com/FAA55D97ED0370F08273C3A94F765C22.png'
      })
      notification.onclick = () => {
        notification.close()
      }
    }
    if (fOpenAudio.value) {
      playAudio()
    }
    setTimeout(scrollToBottom, 100)
  })
}

const sendName = () => {
  if (mytext.value) {
    const content = mytext.value
    const jsonParam = {
      uidTo: props.msg.hisId,
      uidFrom: props.msg.myId,
      orderId: orderId.value,
      nameFrom: user.value.username,
      nameTo: props.msg.otherSide,
      content
    }

    socket.emit('/app/message/chat', JSON.stringify(jsonParam))
    const myitem = { ...jsonParam, timeNow: CurentTime() }
    msgLists.value.push(myitem)
    setTimeout(scrollToBottom, 100)
    mytext.value = ''
  } else {
    ElMessage.info(text.value.messageRequired)
  }
}

const CurentTime = () => {
  const now = new Date()
  const year = now.getFullYear()
  const month = now.getMonth() + 1
  const day = now.getDate()
  const hh = now.getHours()
  const mm = now.getMinutes()

  let clock = year + '-'
  if (month < 10) clock += '0'
  clock += month + '-'
  if (day < 10) clock += '0'
  clock += day + ' '
  if (hh < 10) clock += '0'
  clock += hh + ':'
  if (mm < 10) clock += '0'
  clock += mm
  return clock
}

const getBefore = () => {
  const params = {
    orderId: orderId.value,
    Page: currentPage.value
  }

  axios.post(`${host}/chat/getHistoryMessage`, params, {
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json',
      'x-auth-token': localStorage.getItem('TOKEN')
    }
  }).then(response => {
    const resp = response.data
    totalPage.value = resp.totalPage
    if (resp.data) {
      if (currentPage.value < resp.totalPage || currentPage.value === resp.totalPage) {
        showMore.value = true
        for (let i = 0; i < resp.data.length; i++) {
          msgLists.value.unshift(resp.data[i])
        }
        currentPage.value = currentPage.value + 1
      } else {
        showMore.value = false
      }
    } else {
      showMore.value = false
    }
  }).catch(() => {
    console.error('failed to fetch chat history')
  })
}

const initSocket = () => {
  socket = io(window.location.origin, {
    autoConnect: false,
    transports: ['websocket']
  })
  socket.connect()
  socket.on('connect', () => {
    console.log('socket connected')
  })
}

watch(() => props.msg, () => {
  if (props.msg && !socket) {
    initSocket()
    connect()
  }
}, { immediate: true })

onMounted(() => {
  getBefore()
  setTimeout(scrollToBottom, 100)
})

onBeforeUnmount(() => {
  socket?.disconnect()
})
</script>

<style scoped lang="scss">
.chat-card {
  display: flex;
  flex-direction: column;
  height: 100%;
  min-height: 0;
  border-radius: 20px;
  overflow: hidden;
  background: linear-gradient(180deg, rgba(9, 14, 22, 0.74), rgba(15, 22, 33, 0.92));
  border: 1px solid rgba(255, 255, 255, 0.05);
}

.chat-card__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 18px 20px 16px;
}

.chat-card__title h3 {
  margin: 0 0 6px;
  color: #fff;
  font-size: 22px;
}

.chat-card__title p {
  margin: 0;
  color: #96a2b4;
  font-size: 13px;
  line-height: 1.5;
}

.history-button {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 14px;
  border: 1px solid rgba(240, 172, 25, 0.24);
  border-radius: 12px;
  background: rgba(240, 172, 25, 0.08);
  color: #f0ac19;
  cursor: pointer;
  transition: transform 150ms ease, border-color 200ms ease, background 200ms ease;
}

.history-button:hover {
  transform: translateY(-1px);
  border-color: rgba(240, 172, 25, 0.42);
  background: rgba(240, 172, 25, 0.14);
}

.chat-card__banner {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding: 0 20px 16px;
}

.chat-card__system {
  width: 100%;
  min-width: 0;
}

.system-box {
  margin: 0;
  padding: 12px 16px;
  border-radius: 14px;
  background: linear-gradient(135deg, rgba(33, 43, 57, 0.96), rgba(23, 31, 43, 0.92));
  color: #d5dbe4;
  text-align: center;
}

.chat-card__notice {
  display: flex;
  box-sizing: border-box;
  align-items: center;
  gap: 8px;
  width: 100%;
  max-width: 100%;
  padding: 12px 14px;
  border: 1px solid rgba(237, 37, 37, 0.26);
  border-radius: 14px;
  background: rgba(83, 20, 20, 0.22);
  color: #ff7a7a;
  font-size: 13px;
  line-height: 1.5;
}

.chat-card__notice span {
  min-width: 0;
  word-break: break-word;
}

.chat-card__messages {
  flex: 1;
  min-height: 0;
  margin: 0 12px;
  padding: 8px 8px 20px;
  overflow-y: auto;
}

.message-row {
  display: flex;
  gap: 12px;
  margin-bottom: 18px;
  animation: slideIn 0.3s ease-out;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.message-row--self {
  justify-content: flex-end;
}

.message-row--self .message-row__body {
  align-items: flex-end;
}

.message-row__avatar {
  flex: 0 0 auto;
}

.message-row__body {
  display: flex;
  flex-direction: column;
  max-width: min(72%, 620px);
}

.message-row__name {
  margin: 0 0 6px;
  color: #94a0b2;
  font-size: 12px;
}

.message-bubble {
  padding: 14px 18px;
  border-radius: 18px 18px 18px 8px;
  background: linear-gradient(135deg, #edf1f6, #e5e9ef);
  color: #253243;
  line-height: 1.65;
  word-break: break-word;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.08);
  transition: transform 150ms ease;
}

.message-bubble:hover {
  transform: translateY(-1px);
}

.message-bubble--self {
  border-radius: 18px 18px 8px 18px;
  background: linear-gradient(135deg, #f0ac19, #f5b43e);
  color: #0f1a26;
  box-shadow: 0 8px 18px rgba(240, 172, 25, 0.24);
}

.message-row__time {
  margin-top: 6px;
  color: #7d8897;
  font-size: 12px;
}

.chat-card__composer {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 52px;
  gap: 10px;
  padding: 12px 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
  background: rgba(8, 12, 19, 0.36);
}

.chat-card__toggles {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  padding: 12px 20px 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.04);
  color: #d7ddea;
}

.audio-wrap {
  display: none;
}

.send-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 52px;
  height: 44px;
  border: none;
  border-radius: 14px;
  background: linear-gradient(135deg, #f0ac19, #f5b43e);
  cursor: pointer;
  transition: transform 150ms ease, box-shadow 200ms ease;
}

.send-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 8px 18px rgba(240, 172, 25, 0.28);
}

.chat-avatar-image {
  width: 42px;
  height: 42px;
  object-fit: cover;
  border-radius: 50%;
}

.user-avatar-public {
  background: #fff;
  height: 52px;
  width: 52px;
  box-shadow: 0 1px 5px 0 rgba(71, 78, 114, 0.24);
  position: relative;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.user-avatar-in {
  background: #f0a70a;
  height: 42px;
  width: 42px;
  color: #fff;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.chat-card__messages::-webkit-scrollbar {
  width: 6px;
}

.chat-card__messages::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
  border-radius: 3px;
}

.chat-card__messages::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.12);
  border-radius: 3px;
}

.chat-card__messages::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.18);
}

:deep(.msg-input .el-input__wrapper) {
  min-height: 44px;
  padding: 0 12px;
  background: #111b28;
  box-shadow: inset 0 0 0 1px #27313e;
  transition: box-shadow 200ms ease;
}

:deep(.msg-input .el-input__wrapper:hover),
:deep(.msg-input .el-input__wrapper.is-focus) {
  box-shadow: inset 0 0 0 1px #4c617d;
}

:deep(.msg-input .el-input__inner) {
  color: #fff;
  font-size: 14px;
}

:deep(.msg-input .el-input__inner::placeholder) {
  color: #738196;
}

:deep(.el-checkbox) {
  color: #d7ddea;
}

@media (max-width: 768px) {
  .chat-card {
    min-height: 0;
  }

  .chat-card__header {
    flex-direction: column;
    align-items: flex-start;
  }

  .chat-card__messages {
    min-height: 240px;
    padding-bottom: 16px;
  }

  .message-row__body {
    max-width: calc(100% - 52px);
  }

  .chat-card__composer {
    grid-template-columns: 1fr 48px;
    gap: 8px;
  }

  .chat-card__toggles {
    flex-direction: column;
    gap: 10px;
  }

  .send-btn {
    width: 48px;
    height: 42px;
  }
}
</style>

<template>
  <div>
    <div class="chat-line">
      <div class="scroll-chat" id="scrollChat">
        <div class="chat-container-box" id="sysbox">
          <div class="system-box">
            <p class="msg-content" v-show="msg.status === 1">{{ $t('otc.chatline.status_1') }}</p>
            <p class="msg-content" v-show="msg.status === 2">{{ $t('otc.chatline.status_2') }}</p>
            <p class="msg-content" v-show="msg.status === 3">{{ $t('otc.chatline.status_3') }}</p>
            <p class="msg-content" v-show="msg.status === 4">{{ $t('otc.chatline.status_4') }}</p>
            <p class="msg-content" v-show="msg.status === 0">{{ $t('otc.chatline.status_5') }}</p>
          </div>
        </div>
        <h5 class="more" v-show="currentPage < totalPage || currentPage === totalPage">
          <el-icon class="clock"><Clock /></el-icon>
          <span @click="getBefore">{{ $t('otc.chatline.loadmore') }}</span>
        </h5>
        <template v-for="item in msgLists" :key="item.id || item.sendTimeStr">
          <div class="chat-container-box" v-show="item.uidFrom !== msg.myId">
            <div class="user-chat-box">
              <div class="user-avatar-box">
                <div class="avatar-box">
                  <div class="user-face user-avatar-public">
                    <span class="user-avatar-in" v-if="!item.fromAvatar">{{ msgnameS }}</span>
                    <img v-else style="width: 42px;height:42px" :src="item.fromAvatar" />
                  </div>
                </div>
              </div>
              <div class="user-content-box">
                <p class="user-name">{{ msg.otherSide }}</p>
                <div class="chat-info">
                  <div class="user-desc">
                    <p class="icon"></p>
                    <em class="input-in">{{ item.content }}</em>
                  </div>
                  <span class="times">{{ item.sendTimeStr }}</span>
                </div>
              </div>
            </div>
          </div>
          <div class="chat-container-box" v-show="item.uidFrom === msg.myId">
            <div class="user-chat-box self-chat-box">
              <div class="user-content-box">
                <p class="user-name">{{ user?.username }}</p>
                <div class="chat-info">
                  <div class="user-desc self-desc">
                    <p class="icon"></p>
                    <em class="input-in">{{ item.content }}</em>
                  </div>
                  <span class="times">{{ item.timeNow || item.sendTimeStr }}</span>
                </div>
              </div>
              <div class="user-avatar-box">
                <div class="avatar-box">
                  <div class="user-face user-avatar-public">
                    <span class="user-avatar-in" v-if="!item.fromAvatar">{{ usernameS }}</span>
                    <img v-else style="width: 42px;height:42px" :src="item.fromAvatar" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </template>
      </div>
      <h5 class="spe_show">
        <el-icon><Warning /></el-icon>
        {{ $t('otc.chatline.warning') }}
      </h5>
      <div class="audio-wrap">
        <audio id="noticeMusic" :src="audioSrc">您的浏览器不支持 audio 标签。</audio>
      </div>
      <div class="send-msg-box">
        <div class="img-btn" style="cursor: default;background-color:#192330;">
        </div>
        <el-input
          v-model="mytext"
          @keyup.enter="sendName"
          autocomplete="off"
          placeholder="请输入聊天内容..."
          class="msg-input"
        />
        <button class="send-btn" @click="sendName">
          <el-icon :size="28" color="#f0ac19"><Position /></el-icon>
        </button>
        <div class="msg-notice">
          <el-checkbox v-model="fOpenNotice" @change="handleNoticeClick">开启桌面消息提醒</el-checkbox>
          <el-checkbox v-model="fOpenAudio" @change="handleAudioClick">开启声音消息提醒</el-checkbox>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 聊天输入组件
 */
import { ref, computed, onMounted, onBeforeUnmount, watch } from 'vue'
import { ElMessage, ElInput, ElCheckbox, ElIcon } from 'element-plus'
import { Clock, Warning, Position } from '@element-plus/icons-vue'
import axios from 'axios'
import io from 'socket.io-client'

const props = defineProps({
  msg: {
    type: Object,
    required: true
  }
})

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
const timeNow = ref('')

const msgnameS = computed(() => {
  return (props.msg?.otherSide + '').slice(0, 1)
})

const user = computed(() => {
  return JSON.parse(localStorage.getItem('MEMBER') || '{}')
})

const usernameS = computed(() => {
  return (user.value?.username + '').slice(0, 1)
})

const orderId = computed(() => {
  return props.msg?.orderId || ''
})

const scrollToBottom = () => {
  const div = document.getElementById('scrollChat')
  if (div) {
    div.scrollTop = div.scrollHeight
  }
}

const playAudio = () => {
  const audio = document.getElementById('noticeMusic')
  if (audio !== null) audio.play()
}

const handleNoticeClick = () => {
  if (fOpenNotice.value) {
    if (window.Notification) {
      if (Notification.permission === 'default') {
        ElMessage.info('请点击允许进行开启!')
        Notification.requestPermission().then((result) => {
          if (result === 'denied') {
            ElMessage.info('您已屏蔽消息提醒，如需开通，请查看帮助!')
          }
        })
      } else if (Notification.permission === 'denied') {
        ElMessage.info('您已屏蔽消息提醒，如需开通，请查看帮助!')
      } else {
        ElMessage.info('您已开启桌面消息提醒!')
      }
    } else {
      ElMessage.info('您的浏览器不支持该功能')
    }
  } else {
    ElMessage.info('您已关闭桌面消息提醒!')
  }
}

const handleAudioClick = () => {
  if (fOpenAudio.value) {
    ElMessage.info('您已开启声音消息提醒!')
  } else {
    ElMessage.info('您已关闭声音消息提醒!')
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
      content: content
    }

    socket.emit('/app/message/chat', JSON.stringify(jsonParam))
    const myitem = { ...jsonParam, timeNow: CurentTime() }
    msgLists.value.push(myitem)
    setTimeout(scrollToBottom, 100)
    mytext.value = ''
  } else {
    ElMessage.info('请输入聊天内容')
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
    console.error('获取历史消息失败')
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
.msg-notice {
  width: 20%;

  > label {
    margin-top: 8px;
  }
}

.chat-in-box .chat-in .chat-line .spe_show {
  color: #ed2525;
  padding: 16px;
  margin-right: 3px;
  border: 1px solid #27313e;
}

.chat-in-box .chat-in .chat-line {
  padding: 100px 24px 100px 24px;
  box-sizing: border-box;
  border: 1px solid #27313e;
  width: 95%;
  margin: 0 auto;
  position: relative;
  min-height: 650px;
  transform: translate(0, 0);
}

#sysbox {
  position: fixed;
  top: 15px;
  left: 0;
  right: 0;
}

.chat-in-box .chat-in .chat-line .scroll-chat {
  display: flex;
  flex-direction: column;
  height: 557px;
  overflow-y: auto;
  padding-right: 10px;
}

.chat-in-box .chat-in .chat-line .scroll-chat .more {
  text-align: center;
  color: #f0a70a;

  span {
    cursor: pointer;
    padding: 6px 20px 6px 0;
  }
}

.chat-container-box {
  margin-bottom: 20px;
  position: relative;

  h5 {
    text-align: center;
    font-size: 14px;
    padding: 10px 0;
  }

  .user-chat-box {
    .user-avatar-box,
    .user-content-box {
      float: left;
    }

    .user-content-box {
      max-width: 700px;

      .user-name {
        padding-left: 16px;
        font-size: 12px;
        color: #8994a3;
        margin-bottom: 4px;
      }

      .chat-info {
        .user-desc {
          background-color: #f1f1f4;
          padding: 10px 16px;
          border-radius: 6px;
          margin-bottom: 0;
          margin-left: 16px;
          position: relative;
          font-size: 14px;
          word-break: break-all;

          .icon {
            height: 0;
            width: 0;
            border: 8px solid transparent;
            border-right: 14px solid #f1f1f4;
            position: absolute;
            top: 6px;
            left: -20px;
          }

          em {
            font-style: normal;
            font-size: 12px;
          }
        }

        .times {
          padding-left: 16px;
          padding-right: 16px;
          margin-top: 6px;
          display: block;
          text-align: right;
          color: #c3c1c1;
        }
      }
    }
  }

  .system-box {
    padding: 16px 24px;
    background-color: #27313e;
    border-radius: 7px;
    margin: 0 auto;
    margin-bottom: 20px;
    font-size: 16px;
    color: #ccc;
    max-width: 500px;
  }

  .user-chat-box.self-chat-box {
    .user-avatar-box,
    .user-content-box {
      float: right;
    }

    .user-content-box {
      .user-name {
        text-align: right;
        padding-right: 16px;
      }

      .chat-info {
        .user-desc.self-desc {
          margin-left: 0;
          margin-right: 16px;
          background-color: #f0a70a;
          color: #fff;

          .icon {
            height: 0;
            width: 0;
            border: 8px solid transparent;
            border-left: 14px solid #f0a70a;
            position: absolute;
            top: 6px;
            right: -18px;
            left: unset;
          }
        }
      }
    }
  }
}

.chat-in-box .chat-in .chat-line .send-msg-box {
  display: flex;
  flex: 1;
  width: 94%;
  position: absolute;
  bottom: 21px;
  left: 25px;
  box-shadow: 0 0 3px 1px rgba(0, 0, 0, 0.1);
  border: 1px solid #27313e;

  .img-btn {
    position: relative;
    overflow: hidden;
    background-color: transparent;
    width: 10%;
    min-width: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    background: #ebeff5;
  }

  .msg-input {
    flex: 1;
    border: 1px solid transparent;
    border-left: none;
    background-color: transparent;

    :deep(.el-input__inner) {
      color: #fff;
      height: 60px;
      background-color: transparent;
    }
  }

  .send-btn {
    height: 60px;
    border: none;
    background-color: transparent;
    color: white;
    width: 12%;
    outline: none;
    font-size: 18px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    background: #27313e;
  }
}

.avatar-box {
  display: flex;
  align-items: center;
  flex-direction: column;
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

  > .user-avatar-in {
    background: #f0a70a;
    height: 42px;
    width: 42px;
    color: #fff;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>

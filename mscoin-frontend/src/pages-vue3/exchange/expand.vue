<template>
  <div>
    <ElRow class="expand-row" :class="skin">
      <ElCol span="8">
        <span class="expand-key">{{ $t('exchange.expand.time') }} </span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-key">{{ $t('exchange.expand.price') }} </span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-key">{{ $t('exchange.expand.amount') }} </span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-key">成交额</span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-key">{{ $t('exchange.expand.fee') }}</span>
      </ElCol>
    </ElRow>
    <ElRow v-for="(row, index) in rows" :key="index">
      <ElCol span="8">
        <span class="expand-value">{{ row.time | dateFormat }}</span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-value">{{ row.price | toFloor }}</span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-value">{{ row.amount | toFloor }}</span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-value">{{ row.turnover | toFloor }}</span>
      </ElCol>
      <ElCol span="4">
        <span class="expand-value">{{ row.fee | toFloor }}</span>
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 委托详情展开组件
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. props 使用 defineProps
 */
import { ElRow, ElCol } from 'element-plus'

// Props
defineProps({
  skin: String,
  rows: Array
})
</script>

<style scoped>
.expand-row {
  margin-bottom: 8px;
}

:deep(td.el-table__expanded-cell) {
  padding: 10px 50px;
}

.expand-row .el-col {
  line-height: 20px;
}

.expand-key {
  color: #546886;
  font-size: 12px;
}

.expand-value {
  color: #fff;
  font-size: 13px;
}
</style>

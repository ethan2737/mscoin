<template>
  <div>
    <Row class="expand-row" :class="skin">
      <Col span="8">
        <span class="expand-key">{{ $t('exchange.expand.time') }} </span>
      </Col>
      <Col span="4">
        <span class="expand-key">{{ $t('exchange.expand.price') }} </span>
      </Col>
      <Col span="4">
        <span class="expand-key">{{ $t('exchange.expand.amount') }} </span>
      </Col>
      <Col span="4">
        <span class="expand-key">成交额</span>
      </Col>
      <Col span="4">
        <span class="expand-key">{{ $t('exchange.expand.fee') }}</span>
      </Col>
    </Row>
    <Row v-for="(row, index) in rows" :key="index">
      <Col span="8">
        <span class="expand-value">{{ row.time | dateFormat }}</span>
      </Col>
      <Col span="4">
        <span class="expand-value">{{ row.price | toFloor }}</span>
      </Col>
      <Col span="4">
        <span class="expand-value">{{ row.amount | toFloor }}</span>
      </Col>
      <Col span="4">
        <span class="expand-value">{{ row.turnover | toFloor }}</span>
      </Col>
      <Col span="4">
        <span class="expand-value">{{ row.fee | toFloor }}</span>
      </Col>
    </Row>
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
import { Row, Col } from 'element-plus'

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

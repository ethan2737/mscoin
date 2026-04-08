<template>
  <div class="trade-expand">
    <div class="trade-detail">
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.time') }}：</span>
        <span class="value">{{ row.createTime }}</span>
      </div>
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.tradepair') }}：</span>
        <span class="value">{{ row.coinName }}</span>
      </div>
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.direction') }}：</span>
        <span class="value" :class="row.type === 1 ? 'buy-color' : 'sell-color'">
          {{ row.type === 1 ? $t('otc.trade.buy') : $t('otc.trade.sell') }}
        </span>
      </div>
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.price') }}：</span>
        <span class="value">{{ row.price }}</span>
      </div>
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.entrust') }}：</span>
        <span class="value">{{ row.amount }}</span>
      </div>
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.dealt') }}：</span>
        <span class="value">{{ row.dealAmount }}</span>
      </div>
      <div class="detail-row">
        <span class="label">{{ $t('otc.trade.fee') }}：</span>
        <span class="value">{{ row.fee }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 委托展开组件
 */
import { computed } from 'vue'

const props = defineProps({
  row: {
    type: Object,
    required: true
  }
})
</script>

<style scoped lang="scss">
.trade-expand {
  padding: 10px 20px;
  background: #1a1a2e;

  .trade-detail {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;

    .detail-row {
      display: flex;
      align-items: center;
      min-width: 200px;

      .label {
        color: #828ea1;
        font-size: 14px;
        margin-right: 10px;
      }

      .value {
        color: #fff;
        font-size: 14px;

        &.buy-color {
          color: #2eb794;
        }

        &.sell-color {
          color: #f6465d;
        }
      }
    }
  }
}
</style>

const STATUS_LABELS = {
  unpaid: '待付款',
  waitingRelease: '待放行',
  waitingReceipt: '待确认收款',
  finished: '已完成',
  canceled: '已取消',
  appealed: '申诉中'
}

export function getOrderStatusPresentation({ status, tradeType }) {
  if (status === 1) {
    return {
      text: STATUS_LABELS.unpaid,
      actions: tradeType === 0 ? ['confirm-payment', 'cancel-order'] : []
    }
  }

  if (status === 2) {
    return {
      text: tradeType === 1 ? STATUS_LABELS.waitingReceipt : STATUS_LABELS.waitingRelease,
      actions: tradeType === 1 ? ['release-order', 'appeal-order'] : ['appeal-order']
    }
  }

  if (status === 3) {
    return {
      text: STATUS_LABELS.finished,
      actions: []
    }
  }

  if (status === 4) {
    return {
      text: STATUS_LABELS.appealed,
      actions: []
    }
  }

  return {
    text: STATUS_LABELS.canceled,
    actions: []
  }
}

export function canAutoCancelOrder({ status, tradeType }) {
  return status === 1 && tradeType === 0
}

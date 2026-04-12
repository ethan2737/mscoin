import moment from 'moment'

export const isVerifiedFlag = (value) => value === true || value === 1 || value === '1' || value === 'true'

export const normalizeSecuritySetting = (value = {}) => ({
  ...value,
  realVerified: isVerifiedFlag(value.realVerified),
  fundsVerified: isVerifiedFlag(value.fundsVerified),
  accountVerified: isVerifiedFlag(value.accountVerified)
})

export const normalizeWalletBalance = (value = {}) => {
  const raw = Number(value.balance ?? value.base ?? 0)
  return Number.isFinite(raw) ? raw.toFixed(2) : '0.00'
}

export const normalizeOrderPage = (value = {}) => ({
  content: Array.isArray(value.content) ? value.content : [],
  totalElements: Number(value.totalElements ?? value.total ?? 0)
})

export const calcSettlement = (price, amount) => {
  const priceNumber = Number(price)
  const amountNumber = Number(amount)

  if (!Number.isFinite(priceNumber) || !Number.isFinite(amountNumber)) {
    return '0.00'
  }

  return (priceNumber * amountNumber).toFixed(2)
}

export const buildOrderPayload = ({ direction, quote, tradeForm, verifyForm }) => {
  const isBuy = direction === 'buy'

  return {
    price: Number(isBuy ? quote.buy : quote.sell),
    amount: Number(isBuy ? tradeForm.buyAmount : tradeForm.sellAmount),
    payType: isBuy ? tradeForm.payType : tradeForm.receiveType,
    direction: isBuy ? 0 : 1,
    unit: 'USDT',
    fundpwd: verifyForm.fundpwd,
    code: verifyForm.code
  }
}

export const directionLabelOf = (direction) => direction === 0 ? '买入' : direction === 1 ? '卖出' : '--'

export const payModeLabelOf = (payMode) => {
  if (payMode === 'bank') {
    return '银行卡'
  }

  if (payMode === 'alipay') {
    return '支付宝'
  }

  if (payMode === 'wechatpay') {
    return '微信'
  }

  return '--'
}

export const getOrderStatusText = (order = {}) => {
  if (order.status === 0) {
    return '等待承兑商接单'
  }

  if (order.status === 1 && order.direction === 0) {
    return '已接单，等待付款'
  }

  if (order.status === 1 && order.direction === 1) {
    return '已接单，等待平台付款'
  }

  if (order.status === 2) {
    return '已付款，等待放币'
  }

  if (order.status === 3) {
    return '交易完成'
  }

  if (order.status === 4) {
    return order.cancelReason ? `已取消（${order.cancelReason}）` : '已取消'
  }

  return '--'
}

export const formatDateTime = (value) => {
  if (!value) {
    return '--'
  }

  return moment(value).format('YYYY-MM-DD HH:mm:ss')
}

export const formatMoney = (value) => {
  const amount = Number(value)
  return Number.isFinite(amount) ? amount.toFixed(2) : '--'
}

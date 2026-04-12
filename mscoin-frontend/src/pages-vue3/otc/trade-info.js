export function toOtcNumber(value, fallback = 0) {
  const parsed = Number.parseFloat(String(value ?? '').replace(/,/g, '').trim())
  return Number.isFinite(parsed) ? parsed : fallback
}

export function roundTo(value, scale = 8) {
  const factor = 10 ** scale
  return Math.round(toOtcNumber(value) * factor) / factor
}

export function calculateCryptoAmount(money, price) {
  const safePrice = toOtcNumber(price)
  if (safePrice <= 0) {
    return 0
  }

  return roundTo(toOtcNumber(money) / safePrice, 8)
}

export function calculateFiatAmount(amount, price) {
  return roundTo(toOtcNumber(amount) * toOtcNumber(price), 2)
}

export function buildOtcOrderPayload({ advertiseId, coinId, price, money, amount }) {
  return {
    id: toOtcNumber(advertiseId),
    coinId: toOtcNumber(coinId),
    price: toOtcNumber(price),
    money: toOtcNumber(money),
    amount: toOtcNumber(amount)
  }
}

export function resolveTradeAction(advertiseType) {
  return Number(advertiseType) === 1
    ? { orderPath: '/otc/order/buy', actionKey: 'buyin', buttonKey: 'confirmbuyin' }
    : { orderPath: '/otc/order/sell', actionKey: 'sellout', buttonKey: 'confirmsellout' }
}

function toNumber(value, fallback = 0) {
  const parsed = Number(value)
  return Number.isNaN(parsed) ? fallback : parsed
}

function normalizeRisePercent(coin) {
  if (coin.chg !== undefined && coin.chg !== null) {
    return Number((toNumber(coin.chg) * 100).toFixed(2))
  }

  const fallback = coin.rise ?? coin.change24h ?? 0
  return Number(toNumber(fallback).toFixed(2))
}

export function mapThumbToCoinViewModel(coin, favorSymbols = []) {
  const symbol = coin.symbol || ''
  const [name = symbol, base = ''] = symbol.split('/')
  const favorSet = favorSymbols instanceof Set ? favorSymbols : new Set(favorSymbols)
  const price = toNumber(coin.close ?? coin.price ?? 0)

  return {
    ...coin,
    symbol,
    name,
    base,
    href: `${name}_${base}`.toLowerCase(),
    price,
    close: price,
    priceCNY: toNumber(coin.usdRate ?? coin.priceCNY ?? 0),
    rise: normalizeRisePercent(coin),
    high: toNumber(coin.high ?? 0),
    low: toNumber(coin.low ?? 0),
    change24h: toNumber(coin.change ?? coin.change24h ?? 0),
    volume: toNumber(coin.volume ?? coin.turnover ?? 0),
    isFavor: favorSet.has(symbol)
  }
}

export function mergeRealtimeThumb(currentCoin, payload) {
  if (!currentCoin) {
    return payload
  }

  const incomingLastDayClose = toNumber(payload.lastDayClose, 0)
  if (incomingLastDayClose > 0) {
    return {
      ...currentCoin,
      ...payload
    }
  }

  return {
    ...currentCoin,
    ...payload,
    high: currentCoin.high,
    low: currentCoin.low,
    volume: currentCoin.volume,
    turnover: currentCoin.turnover,
    lastDayClose: currentCoin.lastDayClose,
    change: currentCoin.change,
    chg: currentCoin.chg
  }
}

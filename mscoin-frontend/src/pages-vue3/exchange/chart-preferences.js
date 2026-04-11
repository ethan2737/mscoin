export const shouldUseAreaChartForSymbol = (symbolInfo = {}) => {
  const close = Number(symbolInfo.close)
  return symbolInfo.base === 'ETH' &&
    Number(symbolInfo.baseCoinScale) >= 8 &&
    close > 0.9 &&
    close < 1.1
}

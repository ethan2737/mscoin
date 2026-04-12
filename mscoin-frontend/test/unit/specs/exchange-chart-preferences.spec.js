const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const moduleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/exchange/chart-preferences.js')
).href

test('uses area chart for ultra-stable eth quoted pairs with very high quote precision', async () => {
  const { shouldUseAreaChartForSymbol } = await import(moduleUrl)

  assert.equal(shouldUseAreaChartForSymbol({
    base: 'ETH',
    baseCoinScale: 8,
    close: 0.9999
  }), true)

  assert.equal(shouldUseAreaChartForSymbol({
    base: 'ETH',
    baseCoinScale: 8,
    close: 0.9998
  }), true)
})

test('keeps candle chart for regular volatile pairs', async () => {
  const { shouldUseAreaChartForSymbol } = await import(moduleUrl)

  assert.equal(shouldUseAreaChartForSymbol({
    base: 'USDT',
    baseCoinScale: 2,
    close: 72676.8
  }), false)

  assert.equal(shouldUseAreaChartForSymbol({
    base: 'ETH',
    baseCoinScale: 6,
    close: 12.56
  }), false)
})

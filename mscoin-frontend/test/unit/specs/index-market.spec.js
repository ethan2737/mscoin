const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const marketModuleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/index/market.js')
).href

test('maps rise percent and 24h amount change from different fields', async () => {
  const { mapThumbToCoinViewModel } = await import(marketModuleUrl)

  const coin = mapThumbToCoinViewModel({
    symbol: 'BTC/USDT',
    close: 71558.4,
    usdRate: 71558.4,
    chg: 0.0123,
    change: 874.16,
    high: 73130,
    low: 71558.4,
    volume: 278310676.25523883
  }, ['BTC/USDT'])

  assert.equal(coin.rise, 1.23)
  assert.equal(coin.change24h, 874.16)
  assert.equal(coin.isFavor, true)
})

test('falls back to legacy change24h and rise fields for local placeholder data', async () => {
  const { mapThumbToCoinViewModel } = await import(marketModuleUrl)

  const coin = mapThumbToCoinViewModel({
    symbol: 'SOL/USDT',
    price: 82.92,
    priceCNY: 603.12,
    rise: -0.85,
    change24h: -0.71,
    volume: 52746423.17244762
  }, [])

  assert.equal(coin.rise, -0.85)
  assert.equal(coin.change24h, -0.71)
  assert.equal(coin.isFavor, false)
})

test('preserves existing 24h metrics when realtime payload carries empty summary fields', async () => {
  const { mergeRealtimeThumb } = await import(marketModuleUrl)

  const merged = mergeRealtimeThumb({
    symbol: 'BTC/USDT',
    close: 71629.9,
    change: 350,
    chg: 0.00491022,
    volume: 410167440.15642124,
    lastDayClose: 71279.9,
    high: 72350,
    low: 70531.6
  }, {
    symbol: 'BTC/USDT',
    close: 71680.1,
    change: 0,
    chg: 0,
    volume: 0,
    lastDayClose: 0
  })

  assert.equal(merged.close, 71680.1)
  assert.equal(merged.change, 350)
  assert.equal(merged.chg, 0.00491022)
  assert.equal(merged.volume, 410167440.15642124)
  assert.equal(merged.lastDayClose, 71279.9)
})

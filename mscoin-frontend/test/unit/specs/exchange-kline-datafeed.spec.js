const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const datafeedModuleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/assets/js/charting_library/datafeed/bitrade.js')
).href

test('uses trading pair precision as pricescale fallback when explicit scale is missing', async () => {
  const { resolvePriceScale, resolvePricePrecision } = await import(datafeedModuleUrl)

  assert.equal(resolvePriceScale(undefined, { baseCoinScale: 4 }), 10000)
  assert.equal(resolvePriceScale(null, { priceScale: 5 }), 100000)
  assert.equal(resolvePriceScale(undefined, {}), 100)
  assert.equal(resolvePricePrecision(undefined, { baseCoinScale: 8 }), 8)
})

test('normalizes realtime kline payload from both camelCase and pascal case fields', async () => {
  const { normalizeRealtimeBarPayload } = await import(datafeedModuleUrl)

  assert.deepEqual(
    normalizeRealtimeBarPayload({
      time: 1712894400000,
      openPrice: 0.9999,
      highestPrice: 1.0001,
      lowestPrice: 0.9988,
      closePrice: 0.9995,
      volume: 12.34
    }),
    {
      time: 1712894400000,
      open: 0.9999,
      high: 1.0001,
      low: 0.9988,
      close: 0.9995,
      volume: 12.34
    }
  )

  assert.deepEqual(
    normalizeRealtimeBarPayload({
      Time: 1712894460000,
      OpenPrice: 1.01,
      HighestPrice: 1.03,
      LowestPrice: 1,
      ClosePrice: 1.02,
      Volume: 45.6
    }),
    {
      time: 1712894460000,
      open: 1.01,
      high: 1.03,
      low: 1,
      close: 1.02,
      volume: 45.6
    }
  )
})

test('builds a new realtime bar when trade data crosses into the next resolution bucket', async () => {
  const { updateBarFromTrade } = await import(datafeedModuleUrl)

  const updatedBar = updateBarFromTrade(
    {
      time: 1712894400000,
      open: 100,
      high: 105,
      low: 99,
      close: 103,
      volume: 10
    },
    {
      price: 107,
      amount: 2,
      time: 1712894700000
    },
    '5'
  )

  assert.deepEqual(updatedBar, {
    time: 1712894700000,
    open: 103,
    high: 107,
    low: 103,
    close: 107,
    volume: 2
  })
})

test('updates the current realtime bar inside the same resolution bucket', async () => {
  const { updateBarFromTrade } = await import(datafeedModuleUrl)

  const updatedBar = updateBarFromTrade(
    {
      time: 1712894400000,
      open: 100,
      high: 105,
      low: 99,
      close: 103,
      volume: 10
    },
    {
      price: 98,
      amount: 1.5,
      time: 1712894550000
    },
    '5'
  )

  assert.deepEqual(updatedBar, {
    time: 1712894400000,
    open: 100,
    high: 105,
    low: 98,
    close: 98,
    volume: 11.5
  })
})

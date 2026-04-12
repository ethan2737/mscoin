const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const moduleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/exchange/order-utils.js')
).href

test('toOrderNumber converts string input into numeric order values', async () => {
  const { toOrderNumber } = await import(moduleUrl)

  assert.equal(toOrderNumber('1'), 1)
  assert.equal(toOrderNumber('0.0924'), 0.0924)
  assert.equal(toOrderNumber(''), 0)
  assert.equal(toOrderNumber('abc'), 0)
})

test('buildOrderPayload normalizes limit payload numbers before submit', async () => {
  const { buildOrderPayload } = await import(moduleUrl)

  assert.deepEqual(
    buildOrderPayload({
      symbol: '1INCH/USDT',
      direction: 'BUY',
      type: 'LIMIT_PRICE',
      price: '0.0924',
      amount: '1',
      useDiscount: 0
    }),
    {
      symbol: '1INCH/USDT',
      direction: 'BUY',
      type: 'LIMIT_PRICE',
      price: 0.0924,
      amount: 1,
      useDiscount: 0
    }
  )
})

test('buildOrderPayload normalizes market payload amount before submit', async () => {
  const { buildOrderPayload } = await import(moduleUrl)

  assert.deepEqual(
    buildOrderPayload({
      symbol: '1INCH/USDT',
      direction: 'SELL',
      type: 'MARKET_PRICE',
      amount: '2.5'
    }),
    {
      symbol: '1INCH/USDT',
      direction: 'SELL',
      type: 'MARKET_PRICE',
      amount: 2.5
    }
  )
})

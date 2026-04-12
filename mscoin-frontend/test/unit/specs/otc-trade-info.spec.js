const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const moduleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/otc/trade-info.js')
).href

test('calculateCryptoAmount and calculateFiatAmount keep OTC quote conversion stable', async () => {
  const { calculateCryptoAmount, calculateFiatAmount } = await import(moduleUrl)

  assert.equal(calculateCryptoAmount('100', '7.12'), 14.04494382)
  assert.equal(calculateCryptoAmount('', '7.12'), 0)
  assert.equal(calculateFiatAmount('14.04494382', '7.12'), 100)
})

test('buildOtcOrderPayload normalizes numeric fields before submit', async () => {
  const { buildOtcOrderPayload } = await import(moduleUrl)

  assert.deepEqual(
    buildOtcOrderPayload({
      advertiseId: '9001',
      coinId: '1',
      price: '7.15',
      money: '200.50',
      amount: '28.04195804'
    }),
    {
      id: 9001,
      coinId: 1,
      price: 7.15,
      money: 200.5,
      amount: 28.04195804
    }
  )
})

test('resolveTradeAction maps advertise type to the expected OTC order endpoint', async () => {
  const { resolveTradeAction } = await import(moduleUrl)

  assert.deepEqual(resolveTradeAction(1), {
    orderPath: '/otc/order/buy',
    actionKey: 'buyin',
    buttonKey: 'confirmbuyin'
  })
  assert.deepEqual(resolveTradeAction(0), {
    orderPath: '/otc/order/sell',
    actionKey: 'sellout',
    buttonKey: 'confirmsellout'
  })
})

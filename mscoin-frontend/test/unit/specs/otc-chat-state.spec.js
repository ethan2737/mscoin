const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const moduleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/otc/chat-state.js')
).href

test('OTC chat status labels match the real order lifecycle', async () => {
  const { getOrderStatusPresentation } = await import(moduleUrl)

  assert.deepEqual(getOrderStatusPresentation({ status: 1, tradeType: 0 }), {
    text: '待付款',
    actions: ['confirm-payment', 'cancel-order']
  })
  assert.deepEqual(getOrderStatusPresentation({ status: 2, tradeType: 0 }), {
    text: '待放行',
    actions: ['appeal-order']
  })
  assert.deepEqual(getOrderStatusPresentation({ status: 2, tradeType: 1 }), {
    text: '待确认收款',
    actions: ['release-order', 'appeal-order']
  })
  assert.deepEqual(getOrderStatusPresentation({ status: 3, tradeType: 0 }), {
    text: '已完成',
    actions: []
  })
  assert.deepEqual(getOrderStatusPresentation({ status: 0, tradeType: 0 }), {
    text: '已取消',
    actions: []
  })
  assert.deepEqual(getOrderStatusPresentation({ status: 4, tradeType: 1 }), {
    text: '申诉中',
    actions: []
  })
})

test('OTC chat does not expose cancel action after payment is marked complete', async () => {
  const { canAutoCancelOrder } = await import(moduleUrl)

  assert.equal(canAutoCancelOrder({ status: 1, tradeType: 0 }), true)
  assert.equal(canAutoCancelOrder({ status: 2, tradeType: 0 }), false)
  assert.equal(canAutoCancelOrder({ status: 2, tradeType: 1 }), false)
})

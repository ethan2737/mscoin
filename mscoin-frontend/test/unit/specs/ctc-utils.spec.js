const test = require('node:test')
const assert = require('node:assert/strict')
const path = require('path')
const { pathToFileURL } = require('url')

const utilsModuleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/ctc/ctc-utils.js')
).href

test('normalizeSecuritySetting兼容字符串和数字认证标记', async () => {
  const { normalizeSecuritySetting } = await import(utilsModuleUrl)

  const result = normalizeSecuritySetting({
    realVerified: 'true',
    fundsVerified: 1,
    accountVerified: 'false'
  })

  assert.equal(result.realVerified, true)
  assert.equal(result.fundsVerified, true)
  assert.equal(result.accountVerified, false)
})

test('buildOrderPayload按买卖方向构建下单参数', async () => {
  const { buildOrderPayload } = await import(utilsModuleUrl)

  const buyPayload = buildOrderPayload({
    direction: 'buy',
    quote: { buy: '7.18', sell: '7.11' },
    tradeForm: { buyAmount: 100, sellAmount: 50, payType: 'alipay', receiveType: 'bank' },
    verifyForm: { code: '123456', fundpwd: '654321' }
  })
  const sellPayload = buildOrderPayload({
    direction: 'sell',
    quote: { buy: '7.18', sell: '7.11' },
    tradeForm: { buyAmount: 100, sellAmount: 50, payType: 'alipay', receiveType: 'bank' },
    verifyForm: { code: '123456', fundpwd: '654321' }
  })

  assert.deepEqual(buyPayload, {
    price: 7.18,
    amount: 100,
    payType: 'alipay',
    direction: 0,
    unit: 'USDT',
    fundpwd: '654321',
    code: '123456'
  })
  assert.deepEqual(sellPayload, {
    price: 7.11,
    amount: 50,
    payType: 'bank',
    direction: 1,
    unit: 'USDT',
    fundpwd: '654321',
    code: '123456'
  })
})

test('getOrderStatusText输出正确的订单状态文案', async () => {
  const { getOrderStatusText } = await import(utilsModuleUrl)

  assert.equal(getOrderStatusText({ direction: 0, status: 1 }), '已接单，等待付款')
  assert.equal(getOrderStatusText({ direction: 1, status: 1 }), '已接单，等待平台付款')
  assert.equal(getOrderStatusText({ status: 4, cancelReason: '用户取消' }), '已取消（用户取消）')
})

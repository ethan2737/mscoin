const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const helperModuleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/otc/route-helpers.js')
).href

test('OTC helper builds namespaced detail routes instead of legacy global chat routes', async () => {
  const { buildOtcChatPath, buildOtcCheckUserPath } = await import(helperModuleUrl)

  assert.equal(buildOtcChatPath('OTC00005002'), '/otc/chat/OTC00005002')
  assert.equal(buildOtcCheckUserPath('merchant_001'), '/otc/checkuser/merchant_001')
})

test('OTC related pages keep the OTC top navigation active', async () => {
  const { resolveTopNavByPath } = await import(helperModuleUrl)

  assert.equal(resolveTopNavByPath('/otc/trade/usdt'), '/otc/trade/usdt')
  assert.equal(resolveTopNavByPath('/otc/chat/OTC00005002'), '/otc/trade/usdt')
  assert.equal(resolveTopNavByPath('/otc/checkuser/merchant_001'), '/otc/trade/usdt')
  assert.equal(resolveTopNavByPath('/uc/order'), '/otc/trade/usdt')
  assert.equal(resolveTopNavByPath('/uc/ident/business'), '/otc/trade/usdt')
})

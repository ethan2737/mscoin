const test = require('node:test')
const assert = require('node:assert/strict')
const path = require('node:path')
const { pathToFileURL } = require('node:url')

function resolveModule(relativePath) {
  return path.resolve(__dirname, '../../..', relativePath)
}

async function importFresh(relativePath) {
  const fileUrl = pathToFileURL(resolveModule(relativePath)).href
  return import(`${fileUrl}?t=${Date.now()}`)
}

test('永续本地 mock 默认关闭', async () => {
  delete process.env.MSCOIN_ENABLE_PERPETUAL_MOCKS

  const mod = await importFresh('dev/localAcceptanceMocks.mjs')
  const state = mod.createLocalAcceptanceMockState()

  assert.equal(mod.shouldEnablePerpetualMocks(), false)
  assert.equal(
    mod.resolvePerpetualAcceptanceMock({ method: 'POST', path: '/uc/contract/current', body: {} }, state),
    null
  )
})

test('永续相关 uc 路由必须直连 swap-api', async () => {
  const mod = await importFresh('vite.config.mjs')
  const proxy = mod.default.server.proxy

  assert.equal(proxy['/swap'].target, 'http://127.0.0.1:8086')
  assert.equal(proxy['/uc/contract'].target, 'http://127.0.0.1:8086')
  assert.equal(proxy['/uc/contract-wallet'].target, 'http://127.0.0.1:8086')
  assert.equal(proxy['/uc/asset/contract-transaction/all'].target, 'http://127.0.0.1:8086')
  assert.equal(proxy['/uc'].target, 'http://127.0.0.1:8888')
})

const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const mockModuleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../dev/localAcceptanceMocks.mjs')
).href

test('homepage advertise mock returns non-empty carousel data', async () => {
  const { createLocalAcceptanceMockState, resolveLocalAcceptanceMock } = await import(mockModuleUrl)

  const response = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/ancillary/system/advertise'
  }, createLocalAcceptanceMockState())

  assert.equal(response.code, 0)
  assert.equal(response.data.length, 3)
  assert.equal(response.data[0].title, 'BTC/USDT 交易入口')
})

test('favor add and delete mocks keep in-memory selected symbols in sync', async () => {
  const { createLocalAcceptanceMockState, resolveLocalAcceptanceMock } = await import(mockModuleUrl)

  const state = createLocalAcceptanceMockState()

  resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/exchange/favor/add',
    body: { symbol: 'BTC/USDT' }
  }, state)

  let response = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/exchange/favor/find'
  }, state)

  assert.deepEqual(response.data, [{ symbol: 'BTC/USDT' }])

  resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/exchange/favor/delete',
    body: { symbol: 'BTC/USDT' }
  }, state)

  response = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/exchange/favor/find'
  }, state)

  assert.deepEqual(response.data, [])
})

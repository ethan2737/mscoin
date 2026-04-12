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

test('ctc mock supports quote query and order lifecycle', async () => {
  const { createLocalAcceptanceMockState, resolveLocalAcceptanceMock } = await import(mockModuleUrl)

  const state = createLocalAcceptanceMockState()
  const quote = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/market/ctc-usdt'
  }, state)

  assert.equal(quote.code, 0)
  assert.equal(typeof quote.data.buy, 'number')

  const created = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/ctc/new-ctc-order',
    body: {
      price: 7.18,
      amount: 100,
      payType: 'bank',
      direction: 0,
      unit: 'USDT',
      fundpwd: '654321',
      code: '123456'
    }
  }, state)

  assert.equal(created.code, 0)
  assert.equal(created.data.orderSn.startsWith('CTC'), true)

  const page = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/ctc/page-query',
    body: { pageNo: 1, pageSize: 10 }
  }, state)

  assert.equal(page.data.totalElements, 1)
  assert.equal(page.data.content[0].status, 1)

  const paid = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/ctc/pay-ctc-order',
    body: { oid: created.data.id }
  }, state)

  assert.equal(paid.data.status, 2)

  const canceled = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/ctc/cancel-ctc-order',
    body: { oid: created.data.id }
  }, state)

  assert.equal(canceled.data.status, 4)
  assert.equal(canceled.data.cancelReason, '用户取消')
})

test('ctc dependency mocks expose verified security state and account bindings', async () => {
  const { createLocalAcceptanceMockState, resolveLocalAcceptanceMock } = await import(mockModuleUrl)

  const state = createLocalAcceptanceMockState()

  const security = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/approve/security/setting'
  }, state)
  const account = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/approve/account/setting'
  }, state)
  const wallet = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/asset/wallet/USDT'
  }, state)

  assert.equal(security.code, 0)
  assert.equal(security.data.realVerified, 1)
  assert.equal(security.data.fundsVerified, 1)
  assert.equal(security.data.accountVerified, 1)
  assert.equal(account.code, 0)
  assert.equal(account.data.bankVerified, 1)
  assert.equal(account.data.realName, '本地测试用户')
  assert.equal(wallet.code, 0)
  assert.equal(wallet.data.balance, 1250.32)
})

test('account and safe mocks can persist development changes', async () => {
  const { createLocalAcceptanceMockState, resolveLocalAcceptanceMock } = await import(mockModuleUrl)

  const state = createLocalAcceptanceMockState()

  const setFundPassword = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/approve/update/transaction/password',
    body: { newPassword: '123456' }
  }, state)
  const bindBank = resolveLocalAcceptanceMock({
    method: 'POST',
    path: '/uc/approve/update/bank',
    body: {
      realName: '开发联调用户',
      bank: '中国建设银行',
      branch: '深圳南山支行',
      cardNo: '6227000012345678901'
    }
  }, state)

  assert.equal(setFundPassword.code, 0)
  assert.equal(state.security.fundsVerified, 1)
  assert.equal(bindBank.code, 0)
  assert.equal(state.account.bankInfo.bank, '中国建设银行')
  assert.equal(state.security.accountVerified, 1)
})

test('shared mock state is auto-upgraded when old hot-reload state misses ctc dependencies', async () => {
  const { localAcceptanceMockPlugin } = await import(mockModuleUrl)

  globalThis.__MSCOIN_LOCAL_ACCEPTANCE_STATE__ = {
    favorSymbols: new Set(),
    homepageAdvertisements: []
  }

  const plugin = localAcceptanceMockPlugin()
  assert.equal(plugin.name, 'local-acceptance-mocks')

  const upgraded = globalThis.__MSCOIN_LOCAL_ACCEPTANCE_STATE__
  assert.equal(Boolean(upgraded.security), true)
  assert.equal(Boolean(upgraded.account), true)
  assert.equal(Boolean(upgraded.wallet), true)
  assert.equal(upgraded.security.realVerified, 1)
  assert.equal(upgraded.account.bankVerified, 1)
})

const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const moduleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/exchange/wallet.js')
).href

test('pickWalletBalances reads balances by coin unit from wallet list', async () => {
  const { pickWalletBalances } = await import(moduleUrl)

  const balances = pickWalletBalances({
    wallets: [
      { balance: 12.5, coin: { unit: 'USDT' } },
      { balance: 3.25, coin: { unit: 'BTC' } }
    ],
    baseSymbol: 'USDT',
    coinSymbol: 'BTC'
  })

  assert.deepEqual(balances, {
    base: 12.5,
    coin: 3.25
  })
})

test('pickWalletBalances falls back to zero when wallet entry is missing', async () => {
  const { pickWalletBalances } = await import(moduleUrl)

  const balances = pickWalletBalances({
    wallets: [
      { balance: 8, coinName: 'ETH' }
    ],
    baseSymbol: 'USDT',
    coinSymbol: 'BTC'
  })

  assert.deepEqual(balances, {
    base: 0,
    coin: 0
  })
})

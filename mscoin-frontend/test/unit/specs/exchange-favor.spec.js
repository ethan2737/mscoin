const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')
const fs = require('fs')

const favorModuleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/exchange/favor.js')
).href

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('applyFavorState keeps coin map, row state and favor list in sync without duplicates', async () => {
  const { applyFavorState } = await import(favorModuleUrl)

  const trackedCoin = { symbol: 'BTC/USDT', isFavor: false }
  const otherCoin = { symbol: 'ETH/USDT', isFavor: false }
  const row = { symbol: 'BTC/USDT', isFavor: false }
  const coins = {
    _map: {
      'BTC/USDT': trackedCoin,
      'ETH/USDT': otherCoin
    },
    favor: []
  }

  const added = applyFavorState({
    coins,
    row,
    symbol: 'BTC/USDT',
    isFavor: true,
    currentCoinSymbol: 'BTC/USDT'
  })

  assert.equal(row.isFavor, true)
  assert.equal(trackedCoin.isFavor, true)
  assert.deepEqual(coins.favor.map(item => item.symbol), ['BTC/USDT'])
  assert.equal(added.currentCoinIsFavor, true)

  applyFavorState({
    coins,
    row,
    symbol: 'BTC/USDT',
    isFavor: true,
    currentCoinSymbol: 'ETH/USDT'
  })

  assert.deepEqual(coins.favor.map(item => item.symbol), ['BTC/USDT'])

  const removed = applyFavorState({
    coins,
    row,
    symbol: 'BTC/USDT',
    isFavor: false,
    currentCoinSymbol: 'BTC/USDT'
  })

  assert.equal(row.isFavor, false)
  assert.equal(trackedCoin.isFavor, false)
  assert.deepEqual(coins.favor, [])
  assert.equal(removed.currentCoinIsFavor, false)
})

test('favor messages use custom wording and selected rows render with filled stars', () => {
  const exchangeSource = readSource('../../../src/pages-vue3/exchange/Exchange.vue')
  const favorSource = readSource('../../../src/pages-vue3/exchange/favor.js')

  assert.match(favorSource, /已添加自选/)
  assert.match(favorSource, /已取消自选/)
  assert.match(exchangeSource, /import\s+\{\s*Star,\s*StarFilled,\s*InfoFilled\s*\}\s+from\s+'@element-plus\/icons-vue'/)
  assert.match(exchangeSource, /<StarFilled v-if="row\.isFavor" \/>/)
  assert.match(exchangeSource, /<StarFilled v-if="currentCoinIsFavor" \/>/)
})

test('usdt btc and eth lists all expose the favorite toggle entry', () => {
  const exchangeSource = readSource('../../../src/pages-vue3/exchange/Exchange.vue')

  const toggleMatches = exchangeSource.match(/@click\.stop="toggleFavor\(row\)"/g) || []
  const rowFilledStarMatches = exchangeSource.match(/<StarFilled v-if="row\.isFavor" \/>/g) || []

  assert.equal(toggleMatches.length, 3)
  assert.equal(rowFilledStarMatches.length, 3)
})

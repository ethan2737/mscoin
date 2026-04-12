const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')
const { pathToFileURL } = require('url')

const moduleUrl = pathToFileURL(
  path.resolve(__dirname, '../../../src/pages-vue3/exchange/plate-utils.js')
).href

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('sell rows keep real prices near the latest price and pad the top with placeholders', async () => {
  const { buildSellPlateRows } = await import(moduleUrl)

  const rows = buildSellPlateRows([
    { price: 100.1, amount: 2 },
    { price: 100.2, amount: 3 }
  ], 5)

  assert.equal(rows.length, 5)
  assert.equal(rows[0].displayPrice, '--')
  assert.equal(rows[1].displayPrice, '--')
  assert.equal(rows[2].displayPrice, '--')
  assert.equal(rows[3].price, 100.2)
  assert.equal(rows[4].price, 100.1)
  assert.equal(rows[3].totalAmount, 5)
  assert.equal(rows[4].totalAmount, 2)
})

test('buy rows keep real prices near the latest price and pad the bottom with placeholders', async () => {
  const { buildBuyPlateRows } = await import(moduleUrl)

  const rows = buildBuyPlateRows([
    { price: 99.9, amount: 4 },
    { price: 99.8, amount: 1 }
  ], 5)

  assert.equal(rows.length, 5)
  assert.equal(rows[0].price, 99.9)
  assert.equal(rows[1].price, 99.8)
  assert.equal(rows[0].totalAmount, 4)
  assert.equal(rows[1].totalAmount, 5)
  assert.equal(rows[2].displayPrice, '--')
  assert.equal(rows[4].displayPrice, '--')
})

test('placeholder rows are not selectable for price autofill', async () => {
  const { createPlaceholderRow, getPlateRowPrice } = await import(moduleUrl)

  assert.equal(getPlateRowPrice(createPlaceholderRow('BUY', 1)), null)
  assert.equal(getPlateRowPrice({ price: 88.12 }), 88.12)
})

test('exchange plate uses row click handlers and explicit labels around the latest price', () => {
  const exchangeSource = readSource('../../../src/pages-vue3/exchange/Exchange.vue')

  assert.match(exchangeSource, /@row-click="buyPlate"/)
  assert.match(exchangeSource, /@row-click="sellPlate"/)
  assert.match(exchangeSource, /\$t\('exchange\.sellplate'\)/)
  assert.match(exchangeSource, /\$t\('exchange\.lastprice'\)/)
  assert.match(exchangeSource, /\$t\('exchange\.buyplate'\)/)
  assert.match(exchangeSource, /<div class="trade-wrap__header">最新成交<\/div>/)
  assert.match(exchangeSource, /\.trade-wrap\s*\{[\s\S]*?flex:\s*0 0 176px;/)
  assert.match(exchangeSource, /\.trade-wrap__header\s*\{[\s\S]*?height:\s*40px;/)
  assert.match(exchangeSource, /\.trade-wrap__header\s*\{[\s\S]*?background-color:\s*#192330;/)
  assert.match(exchangeSource, /\.trade-wrap__header\s*\{[\s\S]*?border-bottom:\s*1px solid #27313e;/)
  assert.match(exchangeSource, /\.plate-nowprice\s*\{[\s\S]*?min-height:\s*48px;/)
})

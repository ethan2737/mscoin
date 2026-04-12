const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('exchange page binds depth graph to reactive plate state instead of a constant empty payload', () => {
  const exchangeSource = readSource('../../../src/pages-vue3/exchange/Exchange.vue')

  assert.match(exchangeSource, /:values="depthPlate"/)
  assert.doesNotMatch(exchangeSource, /:values="\{ bid: \{ items: \[\] \}, ask: \{ items: \[\] \}, skin \}"/)
})

test('exchange page ignores stale depth responses after a pair switch', () => {
  const exchangeSource = readSource('../../../src/pages-vue3/exchange/Exchange.vue')

  assert.match(exchangeSource, /let depthPlateRequestSeq = 0/)
  assert.match(exchangeSource, /const requestSeq = \+\+depthPlateRequestSeq/)
  assert.match(exchangeSource, /if \(requestSeq !== depthPlateRequestSeq \|\| requestSymbol !== currentCoin\.symbol\) return/)
})

test('depth graph labels include explicit price and amount units', () => {
  const depthGraphSource = readSource('../../../src/pages-vue3/exchange/DepthGraph.vue')

  assert.match(depthGraphSource, /价格/)
  assert.match(depthGraphSource, /累计数量/)
  assert.match(depthGraphSource, /props\.values\?\.priceUnit/)
  assert.match(depthGraphSource, /props\.values\?\.amountUnit/)
})

test('chart type switcher lives above the chart canvas instead of overlaying the plotting area', () => {
  const exchangeSource = readSource('../../../src/pages-vue3/exchange/Exchange.vue')

  assert.match(exchangeSource, /<div class="chart-panel">[\s\S]*?<div class="handler">/)
  assert.match(exchangeSource, /\.chart-panel\s*\{/)
  assert.match(exchangeSource, /\.imgtable\s*\{[\s\S]*?\.handler\s*\{[\s\S]*?justify-content:\s*flex-end;/)
  assert.match(exchangeSource, /\.imgtable\s*\{[\s\S]*?\.handler\s*\{[\s\S]*?height:\s*40px;/)
  assert.match(exchangeSource, /\.imgtable\s*\{[\s\S]*?\.handler\s*\{[\s\S]*?border-top:\s*1px solid #223047;/)
})

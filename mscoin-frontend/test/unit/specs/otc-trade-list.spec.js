const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('OTC trade list renders merchant names in plain text instead of masking them', () => {
  const source = readSource('../../../src/pages-vue3/otc/Trade.vue')

  assert.doesNotMatch(source, /strpro\(row\.memberName\)/)
  assert.match(source, /\{\{\s*row\.memberName\s*\}\}/)
})

test('OTC trade list uses dark board styling instead of default white element table surface', () => {
  const source = readSource('../../../src/pages-vue3/otc/Trade.vue')

  assert.match(source, /class="[^"]*otc-offer-board/)
  assert.match(source, /\.otc-offer-board/)
  assert.match(source, /--el-table-tr-bg-color:/)
  assert.match(source, /--el-bg-color:/)
})

test('OTC trade tabs and actions keep readable spacing and restrained controls in dark mode', () => {
  const source = readSource('../../../src/pages-vue3/otc/Trade.vue')

  assert.match(source, /\.el-tabs__nav-wrap\s*\{[\s\S]*padding:\s*0 18px;/)
  assert.match(source, /\.merchant-link\s*\{[\s\S]*background:\s*transparent;/)
  assert.match(source, /\.trade-action\s*\{[\s\S]*border-radius:\s*6px;/)
})

test('OTC trade visual hierarchy keeps tabs flat and merchant identity compact', () => {
  const source = readSource('../../../src/pages-vue3/otc/Trade.vue')

  assert.match(source, /\.el-tabs__item\s*\{[\s\S]*background:\s*transparent;/)
  assert.match(source, /\.merchant-avatar-image\s*\{[\s\S]*width:\s*32px;/)
  assert.match(source, /\.merchant-badge\s*\{[\s\S]*height:\s*14px;/)
  assert.match(source, /\.trade-action\s*\{[\s\S]*border-radius:\s*6px;/)
})

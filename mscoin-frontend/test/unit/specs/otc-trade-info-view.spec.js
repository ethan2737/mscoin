const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('OTC trade info shows plain merchant identity and persistent range hints', () => {
  const source = readSource('../../../src/pages-vue3/otc/TradeInfo.vue')

  assert.doesNotMatch(source, /maskedUsername/)
  assert.match(source, /\{\{\s*advertiseDetail\.username\s*\}\}/)
  assert.match(source, /\.avatar-shell\s*\{[\s\S]*width:\s*48px;/)
  assert.match(source, /\{\{\s*moneyFieldTip\s*\}\}/)
})

test('OTC trade info keeps both amount inputs on the same dark border style', () => {
  const source = readSource('../../../src/pages-vue3/otc/TradeInfo.vue')

  assert.match(source, /--el-input-border-color:\s*#27313e;/)
  assert.match(source, /--el-input-hover-border-color:\s*#27313e;/)
  assert.match(source, /--el-input-focus-border-color:\s*#27313e;/)
  assert.match(source, /el-input__wrapper\.is-focus/)
  assert.match(source, /box-shadow:\s*inset 0 0 0 1px #27313e/)
  assert.match(source, /el-input-group__prepend/)
})

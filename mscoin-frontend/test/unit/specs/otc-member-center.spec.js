const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('member center exposes OTC menu entries for merchant auth ads and orders', () => {
  const source = readSource('../../../src/pages-vue3/uc/MemberCenter.vue')

  assert.match(source, /<el-sub-menu index="4">/)
  assert.match(source, /router-link to="\/uc\/ident\/business">商家认证<\/router-link>/)
  assert.match(source, /router-link to="\/uc\/ad"/)
  assert.match(source, /router-link to="\/uc\/order"/)
  assert.match(source, /'\/uc\/ident\/business': '4-1'/)
  assert.match(source, /'\/uc\/order': '4-3'/)
})

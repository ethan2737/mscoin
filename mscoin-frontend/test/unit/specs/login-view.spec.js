const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

test('login inputs keep explicit dark-theme default and focus styles', () => {
  const filePath = path.resolve(__dirname, '../../../src/pages-vue3/uc/Login.vue')
  const source = fs.readFileSync(filePath, 'utf8')

  assert.match(source, /class="auth-form-item"/)
  assert.match(source, /\.auth-form-item :deep\(\.el-input__wrapper\)/)
  assert.match(source, /\.auth-form-item :deep\(\.el-input-group__prepend\)/)
  assert.match(source, /\.auth-form-item :deep\(\.el-input-group__prepend:focus-within\)/)
  assert.match(source, /\.auth-form-item :deep\(\.el-input__wrapper\.is-focus\)/)
  assert.match(source, /box-shadow: inset 0 0 0 1px #f0ac19;/)
})

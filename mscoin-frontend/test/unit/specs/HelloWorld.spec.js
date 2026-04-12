const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('app shell navigation uses route paths as active menu keys', () => {
  const appSource = readSource('../../../src/App.vue')

  assert.match(appSource, /:default-active="activeNav"/)
  assert.match(appSource, /@select="handleNavSelect"/)
  assert.match(appSource, /import \{ resolveTopNavByPath \} from '\.\/pages-vue3\/otc\/route-helpers'/)
  assert.match(appSource, /store\?\.commit\('navigate', resolveTopNavByPath\(path\)\)/)
})

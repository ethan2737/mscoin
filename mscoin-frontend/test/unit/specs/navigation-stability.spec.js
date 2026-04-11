const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('app router view recreates page instances on route changes', () => {
  const appSource = readSource('../../../src/App.vue')

  assert.match(appSource, /const routeViewKey = computed\(\(\) => \{/)
  assert.match(appSource, /<router-view v-if="isRouterAlive" :key="routeViewKey"><\/router-view>/)
})

test('swap page reacts to pair changes and disposes realtime resources on leave', () => {
  const swapSource = readSource('../../../src/pages-vue3/swapindex/Swapindex.vue')

  assert.match(swapSource, /const socketRef = ref\(null\)/)
  assert.match(swapSource, /const disconnectSocket = \(\) => \{/)
  assert.match(swapSource, /socketRef\.value\.removeAllListeners\(\)/)
  assert.match(swapSource, /const destroyKline = \(\) => \{/)
  assert.match(swapSource, /watch\(\(\) => route\.params\.pair, \(pair, previousPair\) => \{/)
  assert.match(swapSource, /onBeforeUnmount\(\(\) => \{/)
})

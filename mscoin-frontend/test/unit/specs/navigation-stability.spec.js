const fs = require('fs')
const path = require('path')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

describe('navigation stability regression coverage', () => {
  test('app router view recreates page instances on route changes', () => {
    const appSource = readSource('../../../src/App.vue')

    expect(appSource).toContain('const routeViewKey = computed(() => {')
    expect(appSource).toContain("<router-view v-if=\"isRouterAlive\" :key=\"routeViewKey\"></router-view>")
  })

  test('swap page reacts to pair changes and disposes realtime resources on leave', () => {
    const swapSource = readSource('../../../src/pages-vue3/swapindex/Swapindex.vue')

    expect(swapSource).toContain('const socketRef = ref(null)')
    expect(swapSource).toContain('const disconnectSocket = () => {')
    expect(swapSource).toContain('socketRef.value.removeAllListeners()')
    expect(swapSource).toContain('const destroyKline = () => {')
    expect(swapSource).toContain("watch(() => route.params.pair, (pair, previousPair) => {")
    expect(swapSource).toContain('onBeforeUnmount(() => {')
  })
})

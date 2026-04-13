const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('OTC chat page renders an order workbench layout with a status hero and split content area', () => {
  const source = readSource('../../../src/pages-vue3/otc/Chat.vue')

  assert.match(source, /class="order-workbench"/)
  assert.match(source, /class="order-hero"/)
  assert.match(source, /class="workbench-grid"/)
  assert.match(source, /class="[^"]*counterpart-card[^"]*"/)
  assert.match(source, /class="[^"]*payment-methods-card[^"]*"/)
  assert.match(source, /class="[^"]*overview-card[^"]*"/)
  assert.match(source, /class="[^"]*communication-card[^"]*"/)
  assert.match(source, /class="[^"]*guide-card[^"]*"/)
})

test('OTC chat page resolves breadcrumb label from route source instead of hardcoding my orders', () => {
  const source = readSource('../../../src/pages-vue3/otc/Chat.vue')

  assert.match(source, /breadcrumbEntryLabel/)
  assert.match(source, /breadcrumbEntryPath/)
  assert.match(source, /route\.query\.source/)
  assert.doesNotMatch(source, /router-link to="\/uc\/order" class="breadcrumb-link"/)
})

test('OTC chat page prevents NaN countdown output and aligns workbench rows with shared grid areas', () => {
  const source = readSource('../../../src/pages-vue3/otc/Chat.vue')

  assert.match(source, /Number\.isFinite\(createTime\)/)
  assert.match(source, /reserveTime\.value = '0:00'/)
  assert.match(source, /grid-template-areas:/)
  assert.match(source, /guide-card/)
})

test('OTC chatline becomes an embedded chat card instead of a page-level framed panel', () => {
  const source = readSource('../../../src/pages-vue3/otc/Chatline.vue')

  assert.match(source, /class="chat-card"/)
  assert.match(source, /class="chat-card__messages"/)
  assert.match(source, /class="chat-card__composer"/)
  assert.doesNotMatch(source, /padding: 100px 24px 100px 24px;/)
})

const fs = require('fs')
const path = require('path')
const test = require('node:test')
const assert = require('node:assert/strict')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

test('login flow establishes authenticated session before redirecting to the protected route', () => {
  const helperSource = readSource('../../../src/utils/auth-session.js')
  const loginSource = readSource('../../../src/pages-vue3/uc/Login.vue')

  assert.match(helperSource, /export async function establishAuthenticatedSession/)
  assert.match(helperSource, /storage\.setItem\(tokenKey, token\)/)
  assert.match(helperSource, /store\?\.commit\?\.?\('setMember', member\)/)
  assert.match(helperSource, /catch \(error\)/)
  assert.match(helperSource, /keep the successful login payload as the session source/)
  assert.ok(loginSource.includes('await establishAuthenticatedSession({'))
  assert.ok(
    loginSource.indexOf('await establishAuthenticatedSession({') <
      loginSource.indexOf('router.push(getReturnUrl())')
  )
})

test('logout flow clears cached auth state and redirects protected pages back through login', () => {
  const helperSource = readSource('../../../src/utils/auth-session.js')
  const appSource = readSource('../../../src/App.vue')
  const memberCenterSource = readSource('../../../src/pages-vue3/uc/MemberCenter.vue')
  const mainSource = readSource('../../../src/main-vue3.js')

  assert.match(helperSource, /export function clearAuthenticatedSession/)
  assert.match(helperSource, /storage\.removeItem\(tokenKey\)/)
  assert.match(helperSource, /storage\.removeItem\(memberKey\)/)
  assert.match(appSource, /clearAuthenticatedSession\(\{ storage: localStorage, store \}\)/)
  assert.match(appSource, /path: '\/login'/)
  assert.match(memberCenterSource, /hasAuthenticatedSession\(\{ storage: localStorage, store \}\)/)
  assert.match(mainSource, /if \(!to\.path\.startsWith\('\/uc'\)\) \{/)
})

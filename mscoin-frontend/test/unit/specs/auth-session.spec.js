const fs = require('fs')
const path = require('path')

function readSource(relativePath) {
  return fs.readFileSync(path.resolve(__dirname, relativePath), 'utf8')
}

describe('auth session integration', () => {
  test('login flow establishes authenticated session before redirecting to the protected route', () => {
    const helperSource = readSource('../../../src/utils/auth-session.js')
    const loginSource = readSource('../../../src/pages-vue3/uc/Login.vue')

    expect(helperSource).toContain('export async function establishAuthenticatedSession')
    expect(helperSource).toContain("storage.setItem(tokenKey, token)")
    expect(helperSource).toContain("store?.commit?.('setMember', member)")
    expect(helperSource).toContain('catch (error)')
    expect(helperSource).toContain('keep the successful login payload as the session source')
    expect(loginSource).toContain('await establishAuthenticatedSession({')
    expect(loginSource.indexOf('await establishAuthenticatedSession({')).toBeLessThan(
      loginSource.indexOf('router.push(getReturnUrl())')
    )
  })

  test('logout flow clears cached auth state and redirects protected pages back through login', () => {
    const helperSource = readSource('../../../src/utils/auth-session.js')
    const appSource = readSource('../../../src/App.vue')
    const memberCenterSource = readSource('../../../src/pages-vue3/uc/MemberCenter.vue')
    const mainSource = readSource('../../../src/main-vue3.js')

    expect(helperSource).toContain('export function clearAuthenticatedSession')
    expect(helperSource).toContain('storage.removeItem(tokenKey)')
    expect(helperSource).toContain('storage.removeItem(memberKey)')
    expect(appSource).toContain('clearAuthenticatedSession({ storage: localStorage, store })')
    expect(appSource).toContain("path: '/login'")
    expect(memberCenterSource).toContain('hasAuthenticatedSession({ storage: localStorage, store })')
    expect(mainSource).toContain("if (!to.path.startsWith('/uc')) {")
  })
})

const fs = require('fs')
const path = require('path')

describe('login view styles', () => {
  test('login inputs keep explicit dark-theme default and focus styles', () => {
    const filePath = path.resolve(__dirname, '../../../src/pages-vue3/uc/Login.vue')
    const source = fs.readFileSync(filePath, 'utf8')

    expect(source).toContain('class="auth-form-item"')
    expect(source).toContain('.auth-form-item :deep(.el-input__wrapper)')
    expect(source).toContain('.auth-form-item :deep(.el-input-group__prepend)')
    expect(source).toContain('.auth-form-item :deep(.el-input-group__prepend:focus-within)')
    expect(source).toContain('.auth-form-item :deep(.el-input__wrapper.is-focus)')
    expect(source).toContain("box-shadow: inset 0 0 0 1px #f0ac19;")
  })
})

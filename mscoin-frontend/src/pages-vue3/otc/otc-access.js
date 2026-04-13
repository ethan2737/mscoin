import axios from 'axios'
import { normalizeSecuritySetting } from '../ctc/ctc-utils'

function getAuthHeaders() {
  return {
    'Content-Type': 'application/json',
    'x-auth-token': localStorage.getItem('TOKEN')
  }
}

export async function fetchOtcSecuritySetting(host = '') {
  const response = await axios.post(`${host}/uc/approve/security/setting`, {}, {
    withCredentials: true,
    headers: getAuthHeaders()
  })

  const resp = response.data
  if (resp.code !== 0) {
    throw new Error(resp.message || '获取安全设置失败')
  }

  return normalizeSecuritySetting(resp.data || {})
}

export function hasRealNameVerification(security = {}, member = {}) {
  return security.realVerified === true || member.realVerified === 1 || member.realVerified === true || Boolean(security.realName || member.realName)
}

export async function ensureOtcTradeAccess({
  host = '',
  member = {},
  onMissingVerification
} = {}) {
  const security = await fetchOtcSecuritySetting(host)

  if (!hasRealNameVerification(security, member)) {
    onMissingVerification?.('请先完成实名认证')
    return { allowed: false, security }
  }

  return { allowed: true, security }
}

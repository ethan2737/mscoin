import Api from './api'

const compatApi = Object.freeze({
  ...Api,
  ucenter: {
    ...Api.uc,
    wallet: Api.uc.wallet
  },
  exchange: {
    ...Api.exchange,
    currentOrder: Api.exchange.current,
    historyOrder: Api.exchange.history,
    cancel: Api.exchange.orderCancel,
    order: Api.exchange.orderAdd
  }
})

export const runtimeContract = Object.freeze({
  host: '',
  wshost: '',
  api: compatApi,
  tokenKey: 'TOKEN',
  memberKey: 'MEMBER'
})

export function useRuntimeContract() {
  return runtimeContract
}

export function buildApiUrl(path = '') {
  return `${runtimeContract.host}${path}`
}

export function getAuthToken() {
  return localStorage.getItem(runtimeContract.tokenKey) || ''
}

export function getAuthHeaders() {
  const token = getAuthToken()
  return token ? { 'x-auth-token': token } : {}
}

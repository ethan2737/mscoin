import axios from 'axios'
import { buildApiUrl, getAuthHeaders } from '@config/runtime-vue3'

const post = async (path, payload = {}) => {
  const response = await axios.post(buildApiUrl(path), payload, {
    headers: getAuthHeaders()
  })

  return response.data
}

export const fetchCtcQuote = () => post('/market/ctc-usdt')
export const fetchUsdtWallet = () => post('/uc/asset/wallet/USDT')
export const fetchSecuritySetting = () => post('/uc/approve/security/setting')
export const fetchCtcOrders = (payload) => post('/uc/ctc/page-query', payload)
export const fetchCtcDetail = (payload) => post('/uc/ctc/detail', payload)
export const createCtcOrder = (payload) => post('/uc/ctc/new-ctc-order', payload)
export const sendCtcCode = () => post('/uc/mobile/ctc/code')
export const markCtcOrderPaid = (payload) => post('/uc/ctc/pay-ctc-order', payload)
export const cancelCtcOrder = (payload) => post('/uc/ctc/cancel-ctc-order', payload)

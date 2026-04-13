const OTC_TOP_NAV_PATH = '/otc/trade/usdt'

const OTC_RELATED_PREFIXES = [
  '/otc',
  '/chat',
  '/checkuser'
]

const OTC_RELATED_UC_PATHS = new Set([
  '/uc/order',
  '/uc/ad',
  '/uc/ad/create',
  '/uc/ad/update',
  '/uc/ident/business'
])

function normalizePathSegment(value) {
  return encodeURIComponent(String(value ?? '').trim())
}

export function buildOtcChatPath(tradeId, options = {}) {
  const basePath = `/otc/chat/${normalizePathSegment(tradeId)}`
  const source = String(options.source || '').trim()

  if (!source) {
    return basePath
  }

  return `${basePath}?source=${normalizePathSegment(source)}`
}

export function buildOtcCheckUserPath(memberId) {
  return `/otc/checkuser/${normalizePathSegment(memberId)}`
}

export function buildOtcTradeInfoPath(tradeId) {
  return `/otc/tradeInfo?tradeId=${normalizePathSegment(tradeId)}`
}

export function resolveTradeIdFromRoute(route) {
  return String(route?.params?.id || route?.query?.tradeId || '').trim()
}

export function resolveMemberIdFromRoute(route) {
  return String(route?.params?.id || route?.query?.id || '').trim()
}

export function resolveTopNavByPath(path) {
  if (!path) {
    return '/'
  }

  if (OTC_RELATED_UC_PATHS.has(path) || OTC_RELATED_PREFIXES.some((prefix) => path === prefix || path.startsWith(`${prefix}/`))) {
    return OTC_TOP_NAV_PATH
  }

  if (path === '/' || path === '/index') {
    return '/'
  }

  if (path === '/exchange' || path.startsWith('/exchange/')) {
    return '/exchange'
  }

  if (path === '/ctc' || path.startsWith('/ctc/')) {
    return '/ctc'
  }

  if (path === '/swapindex' || path.startsWith('/swapindex/1')) {
    return '/swapindex/1'
  }

  if (path.startsWith('/swapindex/2')) {
    return '/swapindex/2'
  }

  if (path === '/activity' || path.startsWith('/activity/') || path === '/lab' || path.startsWith('/lab/')) {
    return '/activity'
  }

  if (path === '/mining' || path.startsWith('/mining/')) {
    return '/mining'
  }

  if (path === '/crowdfunding' || path.startsWith('/crowdfunding/')) {
    return '/crowdfunding'
  }

  return '/'
}

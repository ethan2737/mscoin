function sendJson(res, statusCode, payload) {
  res.statusCode = statusCode
  res.setHeader('Content-Type', 'application/json; charset=utf-8')
  res.end(JSON.stringify(payload))
}

function ok(data) {
  return { code: 0, message: 'success', data }
}

function pathOf(req) {
  return (req.url || '').split('?')[0]
}

async function readJsonBody(req) {
  const chunks = []

  for await (const chunk of req) {
    chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk))
  }

  if (chunks.length === 0) {
    return {}
  }

  try {
    return JSON.parse(Buffer.concat(chunks).toString('utf8'))
  } catch {
    return {}
  }
}

const otcCoins = [{ unit: 'USDT' }]

const otcAdverts = [
  {
    advertiseId: 9001,
    advertiseType: 1,
    memberName: 'merchant-alpha',
    avatar: '',
    level: 2,
    transactions: 128,
    payMode: '银行卡 / 支付宝',
    remainAmount: 5000,
    minLimit: 100,
    maxLimit: 5000,
    price: 7.15
  },
  {
    advertiseId: 9002,
    advertiseType: 0,
    memberName: 'merchant-beta',
    avatar: '',
    level: 2,
    transactions: 86,
    payMode: '银行卡',
    remainAmount: 3200,
    minLimit: 200,
    maxLimit: 3200,
    price: 7.08
  }
]

const miningItems = [
  {
    id: 7001,
    miningName: 'USDT 活期',
    title: '本地联调矿池基线',
    smallImageUrl: '/src/assets/images/activity-bg_pro.jpg',
    miningUnit: 'USDT',
    limitTimes: 0,
    unit: '稳定收益模拟池',
    miningTimes: 30,
    miningDaysprofit: '1.68',
    miningDays: 1,
    miningPeriod: 0,
    price: '100.00',
    acceptUnit: 'USDT'
  }
]

const crowdfundingProjects = [
  {
    id: 8001,
    type: 'wish',
    fundingTitle: '本地联调爱心项目',
    picUrl: '/src/assets/images/crowdfunding_banner.png',
    targetAmount: '5000',
    amountReceived: '1250',
    times: 18,
    addTime: '2026-04-01 00:00:00',
    passTime: '2026-04-02 00:00:00',
    endTime: '2026-05-01 00:00:00',
    sponsor: '本地公益发起人'
  }
]

const offlineWelfare = [
  {
    id: 8101,
    fundingTitle: '线下公益物资补助',
    picUrl: '/src/assets/images/crowdfunding_banner.png',
    amountReceived: '860',
    applyAmount: '1000',
    passTime: '2026-04-03 00:00:00',
    sponsor: '公益服务中心'
  }
]

const swapThumb = [
  {
    symbol: 'BTC/USDT',
    name: 'BTC',
    close: 0,
    rose: '0.00%',
    chg: 0,
    high: 0,
    low: 0,
    volume: 0,
    usdRate: 7.2
  }
]

const emptyPlate = {
  ask: { items: [] },
  bid: { items: [] }
}

const createCtcCounterpartyInfo = (payMode) => ({
  bankInfo: {
    bank: '招商银行',
    branch: '深圳科技园支行',
    cardNo: '6222028800012345678'
  },
  alipay: {
    aliNo: 'merchant-alpha@mscoin.test',
    qrCodeUrl: '/src/assets/images/lang-zh.png'
  },
  wechatPay: {
    wechat: 'merchant-alpha-wechat',
    qrWeCodeUrl: '/src/assets/images/lang-en.png'
  },
  payMode
})

const createMockSecurityProfile = () => ({
  username: 'local-tester',
  level: { id: 1 },
  realVerified: 1,
  realAuditing: 0,
  phoneVerified: 1,
  fundsVerified: 1,
  accountVerified: 1,
  bankVerified: 1,
  aliVerified: 1,
  wechatVerified: 1,
  mobilePhone: '13800000000',
  realName: '本地测试用户',
  realNameRejectReason: null
})

const createMockAccountProfile = () => ({
  bankVerified: 1,
  aliVerified: 1,
  wechatVerified: 1,
  realName: '本地测试用户',
  bankInfo: {
    bank: '招商银行',
    branch: '深圳科技园支行',
    cardNo: '6222028800012345678'
  },
  alipay: {
    aliNo: 'local-tester@mscoin.test',
    qrCodeUrl: '/src/assets/images/lang-zh.png'
  },
  wechatPay: {
    wechat: 'local-tester-wechat',
    qrWeCodeUrl: '/src/assets/images/lang-en.png'
  }
})

const resolveUploadedImageUrl = (kind = 'proof') => kind === 'wechat'
  ? '/src/assets/images/lang-en.png'
  : '/src/assets/images/lang-zh.png'

const syncSecurityAccountFlags = (state) => {
  const accountVerified = state.account.bankVerified === 1 || state.account.aliVerified === 1 || state.account.wechatVerified === 1
  state.security.accountVerified = accountVerified ? 1 : 0
  state.security.bankVerified = state.account.bankVerified
  state.security.aliVerified = state.account.aliVerified
  state.security.wechatVerified = state.account.wechatVerified
}

const createCtcOrder = (state, body = {}) => {
  const now = new Date().toISOString()
  const id = state.ctc.nextOrderId++
  const price = Number(body.price || state.ctc.quote.buy)
  const amount = Number(body.amount || 0)
  const money = Number((price * amount).toFixed(2))
  const payMode = body.payType || 'bank'
  const direction = Number(body.direction || 0)
  const baseOrder = {
    id,
    orderSn: `CTC${String(id).padStart(8, '0')}`,
    createTime: now,
    confirmTime: now,
    currentTime: now,
    direction,
    amount,
    price,
    money,
    status: 1,
    realName: '本地测试用户',
    cancelReason: '',
    unit: body.unit || 'USDT'
  }

  const order = {
    ...baseOrder,
    ...createCtcCounterpartyInfo(payMode)
  }

  state.ctc.orders.unshift(order)
  return order
}

export function createLocalAcceptanceMockState() {
  return {
    favorSymbols: new Set(),
    security: createMockSecurityProfile(),
    account: createMockAccountProfile(),
    wallet: {
      unit: 'USDT',
      balance: 1250.32,
      frozenBalance: 0,
      address: ''
    },
    ctc: {
      quote: {
        buy: 7.18,
        sell: 7.11
      },
      nextOrderId: 1001,
      orders: []
    },
    homepageAdvertisements: [
      {
        id: 11001,
        name: '本地联调轮播-币币交易',
        url: '/src/assets/images/bannerbg.png',
        image: '/src/assets/images/bannerbg.png',
        linkUrl: '/#/exchange/BTC_USDT',
        title: 'BTC/USDT 交易入口'
      },
      {
        id: 11002,
        name: '本地联调轮播-众筹公益',
        url: '/src/assets/images/crowdfunding_banner.png',
        image: '/src/assets/images/crowdfunding_banner.png',
        linkUrl: '/#/crowdfunding',
        title: '众筹与公益创新实验室'
      },
      {
        id: 11003,
        name: '本地联调轮播-矿机理财',
        url: '/src/assets/images/activity-bg_pro.jpg',
        image: '/src/assets/images/activity-bg_pro.jpg',
        linkUrl: '/#/mining',
        title: '矿机专区体验页'
      }
    ]
  }
}

function ensureLocalAcceptanceMockStateShape(state) {
  if (!state.security) {
    state.security = createMockSecurityProfile()
  }

  if (!state.account) {
    state.account = createMockAccountProfile()
  }

  if (!state.wallet) {
    state.wallet = {
      unit: 'USDT',
      balance: 1250.32,
      frozenBalance: 0,
      address: ''
    }
  }

  if (!state.ctc) {
    state.ctc = {
      quote: {
        buy: 7.18,
        sell: 7.11
      },
      nextOrderId: 1001,
      orders: []
    }
  }

  if (!Array.isArray(state.ctc.orders)) {
    state.ctc.orders = []
  }

  if (!state.ctc.quote) {
    state.ctc.quote = { buy: 7.18, sell: 7.11 }
  }

  if (typeof state.ctc.nextOrderId !== 'number') {
    state.ctc.nextOrderId = 1001
  }

  if (!Array.isArray(state.homepageAdvertisements)) {
    state.homepageAdvertisements = createLocalAcceptanceMockState().homepageAdvertisements
  }

  syncSecurityAccountFlags(state)
  return state
}

function getSharedLocalAcceptanceMockState() {
  if (!globalThis.__MSCOIN_LOCAL_ACCEPTANCE_STATE__) {
    globalThis.__MSCOIN_LOCAL_ACCEPTANCE_STATE__ = createLocalAcceptanceMockState()
  }

  return ensureLocalAcceptanceMockStateShape(globalThis.__MSCOIN_LOCAL_ACCEPTANCE_STATE__)
}

export function resolveLocalAcceptanceMock({ method, path, body = {} }, state) {
  if (path === '/uc/ancillary/system/advertise') {
    return ok(state.homepageAdvertisements)
  }

  if (method === 'POST' && path === '/uc/approve/security/setting') {
    return ok(state.security)
  }

  if (method === 'POST' && path === '/uc/approve/account/setting') {
    return ok(state.account)
  }

  if (method === 'POST' && path === '/uc/asset/wallet/USDT') {
    return ok(state.wallet)
  }

  if (method === 'POST' && path === '/market/ctc-usdt') {
    return ok(state.ctc.quote)
  }

  if (method === 'POST' && (
    path === '/uc/mobile/ctc/code' ||
    path === '/uc/mobile/bind/code' ||
    path === '/uc/mobile/update/password/code' ||
    path === '/uc/mobile/transaction/code'
  )) {
    return { code: 0, message: '验证码已发送', data: null }
  }

  if (method === 'POST' && path === '/uc/upload/oss/image') {
    const kind = String(body.kind || body.type || '').toLowerCase()
    return ok(resolveUploadedImageUrl(kind))
  }

  if (method === 'POST' && path === '/uc/approve/real/name') {
    state.security.realVerified = 1
    state.security.realAuditing = 0
    state.security.realName = body.realName || state.security.realName
    state.account.realName = state.security.realName
    return { code: 0, message: '实名认证提交成功', data: null }
  }

  if (method === 'POST' && path === '/uc/approve/bind/phone') {
    state.security.phoneVerified = 1
    state.security.mobilePhone = body.phone || state.security.mobilePhone
    return { code: 0, message: '手机号绑定成功', data: null }
  }

  if (method === 'POST' && path === '/uc/approve/update/password') {
    return { code: 0, message: '登录密码修改成功', data: null }
  }

  if (method === 'POST' && path === '/uc/approve/update/transaction/password') {
    state.security.fundsVerified = 1
    return { code: 0, message: '资金密码设置成功', data: null }
  }

  if (method === 'POST' && (path === '/uc/approve/bind/bank' || path === '/uc/approve/update/bank')) {
    state.account.bankVerified = 1
    state.account.realName = body.realName || state.account.realName
    state.account.bankInfo = {
      bank: body.bank || state.account.bankInfo?.bank || '招商银行',
      branch: body.branch || state.account.bankInfo?.branch || '深圳科技园支行',
      cardNo: body.cardNo || state.account.bankInfo?.cardNo || '6222028800012345678'
    }
    syncSecurityAccountFlags(state)
    return { code: 0, message: '银行卡保存成功', data: null }
  }

  if (method === 'POST' && (path === '/uc/approve/bind/ali' || path === '/uc/approve/update/ali')) {
    state.account.aliVerified = 1
    state.account.realName = body.realName || state.account.realName
    state.account.alipay = {
      aliNo: body.ali || state.account.alipay?.aliNo || 'local-tester@mscoin.test',
      qrCodeUrl: body.qrCodeUrl || state.account.alipay?.qrCodeUrl || resolveUploadedImageUrl('alipay')
    }
    syncSecurityAccountFlags(state)
    return { code: 0, message: '支付宝保存成功', data: null }
  }

  if (method === 'POST' && (path === '/uc/approve/bind/wechat' || path === '/uc/approve/update/wechat')) {
    state.account.wechatVerified = 1
    state.account.realName = body.realName || state.account.realName
    state.account.wechatPay = {
      wechat: body.wechat || state.account.wechatPay?.wechat || 'local-tester-wechat',
      qrWeCodeUrl: body.qrCodeUrl || state.account.wechatPay?.qrWeCodeUrl || resolveUploadedImageUrl('wechat')
    }
    syncSecurityAccountFlags(state)
    return { code: 0, message: '微信收款方式保存成功', data: null }
  }

  if (method === 'POST' && path === '/uc/ctc/page-query') {
    const pageNo = Number(body.pageNo || 1)
    const pageSize = Number(body.pageSize || 10)
    const start = Math.max(pageNo - 1, 0) * pageSize
    const content = state.ctc.orders.slice(start, start + pageSize)
    return ok({
      content,
      totalElements: state.ctc.orders.length
    })
  }

  if (method === 'POST' && path === '/uc/ctc/detail') {
    const order = state.ctc.orders.find((item) => item.id === Number(body.oid))
    if (!order) {
      return { code: 1, message: '订单不存在', data: null }
    }

    return ok({
      ...order,
      currentTime: new Date().toISOString()
    })
  }

  if (method === 'POST' && path === '/uc/ctc/new-ctc-order') {
    if (!body.amount || Number(body.amount) <= 0) {
      return { code: 1, message: '数量必须大于 0', data: null }
    }

    return ok(createCtcOrder(state, body))
  }

  if (method === 'POST' && path === '/uc/ctc/pay-ctc-order') {
    const order = state.ctc.orders.find((item) => item.id === Number(body.oid))
    if (!order) {
      return { code: 1, message: '订单不存在', data: null }
    }

    order.status = 2
    order.currentTime = new Date().toISOString()
    return ok(order)
  }

  if (method === 'POST' && path === '/uc/ctc/cancel-ctc-order') {
    const order = state.ctc.orders.find((item) => item.id === Number(body.oid))
    if (!order) {
      return { code: 1, message: '订单不存在', data: null }
    }

    order.status = 4
    order.cancelReason = '用户取消'
    order.currentTime = new Date().toISOString()
    return ok(order)
  }

  if (method === 'POST' && path === '/exchange/favor/find') {
    return ok(Array.from(state.favorSymbols).map((symbol) => ({ symbol })))
  }

  if (method === 'POST' && path === '/exchange/favor/add') {
    if (body.symbol) {
      state.favorSymbols.add(body.symbol)
    }
    return ok(null)
  }

  if (method === 'POST' && path === '/exchange/favor/delete') {
    if (body.symbol) {
      state.favorSymbols.delete(body.symbol)
    }
    return ok(null)
  }

  return null
}

export function localAcceptanceMockPlugin() {
  const state = getSharedLocalAcceptanceMockState()

  return {
    name: 'local-acceptance-mocks',
    configureServer(server) {
      server.middlewares.use(async (req, res, next) => {
        const method = (req.method || 'GET').toUpperCase()
        const path = pathOf(req)
        const requestBodyNeeded = method === 'POST' && (
          path === '/exchange/favor/add' ||
          path === '/exchange/favor/delete' ||
          path === '/uc/approve/real/name' ||
          path === '/uc/approve/bind/phone' ||
          path === '/uc/approve/update/password' ||
          path === '/uc/approve/update/transaction/password' ||
          path === '/uc/approve/bind/bank' ||
          path === '/uc/approve/update/bank' ||
          path === '/uc/approve/bind/ali' ||
          path === '/uc/approve/update/ali' ||
          path === '/uc/approve/bind/wechat' ||
          path === '/uc/approve/update/wechat' ||
          path === '/uc/upload/oss/image' ||
          path === '/uc/ctc/page-query' ||
          path === '/uc/ctc/detail' ||
          path === '/uc/ctc/new-ctc-order' ||
          path === '/uc/ctc/pay-ctc-order' ||
          path === '/uc/ctc/cancel-ctc-order'
        )
        const body = requestBodyNeeded ? await readJsonBody(req) : {}
        const resolved = resolveLocalAcceptanceMock({ method, path, body }, state)

        if (resolved) {
          return sendJson(res, 200, resolved)
        }

        if (method === 'POST' && path === '/otc/coin') {
          return sendJson(res, 200, ok(otcCoins))
        }

        if (method === 'POST' && path === '/otc/advertise/list') {
          return sendJson(res, 200, ok({ context: otcAdverts, totalElement: otcAdverts.length }))
        }

        if (method === 'POST' && path === '/uc/mining/page-query') {
          return sendJson(res, 200, ok({ content: miningItems, totalElements: miningItems.length }))
        }

        if (method === 'POST' && path === '/uc/mining/detail') {
          return sendJson(res, 200, ok(miningItems[0]))
        }

        if (method === 'GET' && path === '/swap/contract-coin/info/1') {
          return sendJson(res, 200, ok({
            id: 1,
            baseSymbol: 'USDT',
            coinSymbol: 'BTC',
            type: 'ALWAYS',
            coinScale: 4,
            baseCoinScale: 4,
            shareNumber: 0.001,
            takerFee: 0.0005,
            makerFee: 0.0002,
            enableMarketBuy: 1,
            enableMarketSell: 1,
            exchangeable: 1
          }))
        }

        if (method === 'POST' && path === '/swap/symbol-thumb') {
          return sendJson(res, 200, swapThumb)
        }

        if (method === 'POST' && path === '/swap/exchange-plate-mini') {
          return sendJson(res, 200, emptyPlate)
        }

        if (method === 'POST' && path === '/swap/exchange-plate-full') {
          return sendJson(res, 200, emptyPlate)
        }

        if (method === 'POST' && path === '/swap/latest-trade') {
          return sendJson(res, 200, [])
        }

        if ((method === 'GET' || method === 'POST') && path === '/swap/history') {
          return sendJson(res, 200, { s: 'no_data' })
        }

        if (method === 'POST' && path === '/swap/order/position') {
          return sendJson(res, 200, ok([]))
        }

        if (method === 'POST' && path === '/swap/order-entrust/current') {
          return sendJson(res, 200, { content: [], totalElements: 0, number: 0 })
        }

        if (method === 'POST' && path === '/swap/order-entrust/history') {
          return sendJson(res, 200, { content: [], totalElements: 0, number: 0 })
        }

        if (method === 'GET' && path === '/swap/contract-leverage') {
          return sendJson(res, 200, ok({ leverage: 10 }))
        }

        if (method === 'POST' && path === '/uc/asset/contract-wallet/USDT') {
          return sendJson(res, 200, ok({ base: 0, frozen: 0, profit: 0 }))
        }

        if (method === 'POST' && path === '/uc/asset/wallet/KICK') {
          return sendJson(res, 200, ok({ balance: 0 }))
        }

        if (method === 'POST' && path === '/uc/contract/current') {
          return sendJson(res, 200, { content: [], totalElements: 0 })
        }

        if (method === 'POST' && path === '/uc/contract/history') {
          return sendJson(res, 200, { content: [], totalElements: 0 })
        }

        if (method === 'POST' && path === '/uc/contract/coin/thumb') {
          return sendJson(res, 200, [{ id: 1, name: 'BTC/USDT' }])
        }

        if (method === 'GET' && path === '/uc/crowdfunding/totalAmount') {
          return sendJson(res, 200, ok({ totalMoney: '12850', totalProject: '12', totalTimes: '56' }))
        }

        if (method === 'GET' && path === '/uc/crowdfunding/crowdfundingList') {
          return sendJson(res, 200, ok({ funding: crowdfundingProjects, medicalfunding: [] }))
        }

        if (method === 'GET' && path === '/uc/crowdfunding/offlineList') {
          return sendJson(res, 200, ok(offlineWelfare))
        }

        if (method === 'POST' && path === '/uc/crowdfunding/level') {
          return sendJson(res, 200, ok({ crowdRight: 1, welfareRight: 1 }))
        }

        if (method === 'POST' && path === '/uc/crowdfunding/getWelfareDetail') {
          return sendJson(res, 200, ok(['/src/assets/images/crowdfunding_text.png']))
        }

        if (method === 'POST' && path === '/uc/crowdfunding/donateList') {
          return sendJson(res, 200, ok({
            recordList: [
              { memberName: 'local-tester', amount: '100', createTime: '2026-04-09 15:00:00' }
            ],
            fundingCommonList: [
              { memberName: 'local-tester', comments: '开发环境验收基线', createTime: '2026-04-09 15:00:00' }
            ],
            funding: crowdfundingProjects[0],
            picsUrl: ['/src/assets/images/crowdfunding_text.png']
          }))
        }

        if (method === 'POST' && path === '/uc/activity/myMinings') {
          return sendJson(res, 200, ok({ content: [] }))
        }

        if (method === 'POST' && path === '/uc/activity/fetchProfit') {
          return sendJson(res, 200, ok(null))
        }

        if (method === 'POST' && path === '/uc/miningorder/my-minings') {
          return sendJson(res, 200, ok({ content: [] }))
        }

        if (method === 'POST' && path === '/uc/miningorder/fetch-profit') {
          return sendJson(res, 200, ok(null))
        }

        next()
      })
    }
  }
}

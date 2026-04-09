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

export function localAcceptanceMockPlugin() {
  return {
    name: 'local-acceptance-mocks',
    configureServer(server) {
      server.middlewares.use((req, res, next) => {
        const method = (req.method || 'GET').toUpperCase()
        const path = pathOf(req)

        if (path === '/uc/ancillary/system/advertise') {
          return sendJson(res, 200, ok([]))
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

        if (method === 'POST' && path === '/exchange/favor/find') {
          return sendJson(res, 200, ok([]))
        }

        next()
      })
    }
  }
}

// Vue 3 路由配置
// 导入已迁移的 Vue 3 组件
import Login from '../pages-vue3/uc/Login.vue'
import Register from '../pages-vue3/uc/Register.vue'
import MemberCenter from '../pages-vue3/uc/MemberCenter.vue'
import Safe from '../pages-vue3/uc/Safe.vue'
import Account from '../pages-vue3/uc/Account.vue'
import Recharge from '../pages-vue3/uc/Recharge.vue'
import Withdraw from '../pages-vue3/uc/Withdraw.vue'
import Record from '../pages-vue3/uc/Record.vue'
import EntrustCurrent from '../pages-vue3/uc/EntrustCurrent.vue'
import EntrustHistory from '../pages-vue3/uc/EntrustHistory.vue'
import Exchange from '../pages-vue3/exchange/Exchange.vue'
import Swapindex from '../pages-vue3/swapindex/Swapindex.vue'
import Kline from '../pages-vue3/swapindex/Kline.vue'
import WithdrawAddress from '../pages-vue3/uc/WithdrawAddress.vue'
import MoneyIndex from '../pages-vue3/uc/MoneyIndex.vue'
import myorder from '../pages-vue3/uc/myorder.vue'
import ContractMoneyIndex from '../pages-vue3/uc/ContractMoneyIndex.vue'
import ContractRecord from '../pages-vue3/uc/ContractRecord.vue'
import TradeExpand from '../pages-vue3/uc/TradeExpand.vue'
import MinTrade from '../pages-vue3/uc/MinTrade.vue'
import InvitingMin from '../pages-vue3/uc/InvitingMin.vue'
import PayDividends from '../pages-vue3/uc/PayDividends.vue'
import InnovationOrders from '../pages-vue3/uc/InnovationOrders.vue'
import InnovationMinings from '../pages-vue3/uc/InnovationMinings.vue'
import crowdfundingList from '../pages-vue3/uc/crowdfunding-list.vue'
import contractEntrustCurrent from '../pages-vue3/uc/contract/EntrustCurrent.vue'
import contractEntrustHistory from '../pages-vue3/uc/contract/EntrustHistory.vue'
import OtcMain from '../pages-vue3/otc/Main.vue'
import OtcTrade from '../pages-vue3/otc/Trade.vue'
import OtcTradeInfo from '../pages-vue3/otc/TradeInfo.vue'
import OtcCheckUser from '../pages-vue3/otc/CheckUser.vue'
import OtcChat from '../pages-vue3/otc/Chat.vue'
import OtcChatline from '../pages-vue3/otc/Chatline.vue'
import OtcAdPublish from '../pages-vue3/otc/AdPublish.vue'
import OtcMyAd from '../pages-vue3/otc/MyAd.vue'
import OtcCarousel from '../pages-vue3/otc/carousel.vue'
import Help from '../pages-vue3/cms/Help.vue'
import HelpList from '../pages-vue3/cms/HelpList.vue'
import HelpDetail from '../pages-vue3/cms/HelpDetail.vue'
import Notice from '../pages-vue3/cms/Notice.vue'
import NoticeItem from '../pages-vue3/cms/NoticeItem.vue'
import Noticeindex from '../pages-vue3/cms/Noticeindex.vue'
import AboutUs from '../pages-vue3/cms/AboutUs.vue'
import Activity from '../pages-vue3/activity/Activity.vue'
import ActivityDetail from '../pages-vue3/activity/ActivityDetail.vue'
import Partner from '../pages-vue3/activity/Partner.vue'
import Bzb from '../pages-vue3/activity/Bzb.vue'
import Mining from '../pages-vue3/mining/Mining.vue'
import MiningDetail from '../pages-vue3/mining/MiningDetail.vue'
import crowdfundingIndex from '../pages-vue3/crowdfunding/index.vue'
import Ctc from '../pages-vue3/ctc/Ctc.vue'
import Envelope from '../pages-vue3/envelope/Envelope.vue'
import AppDownload from '../pages-vue3/uc/AppDownload.vue'
import MobileRegister from '../pages-vue3/uc/MobileRegister.vue'
import FindPwd from '../pages-vue3/uc/FindPwd.vue'
import IdentBusiness from '../pages-vue3/uc/IdentBusiness.vue'
import Index from '../pages-vue3/index/Index.vue'
import DepthGraph from '../pages-vue3/exchange/DepthGraph.vue'
import BZCountDown from '../pages-vue3/exchange/BZCountDown.vue'
import expandRow from '../pages-vue3/exchange/expand.vue'

export default [
  // 首页
  { path: '/', redirect: '/index' },
  { path: '/index', component: Index },
  { path: '/invite', redirect: '/partner' },

  // 用户认证
  { path: '/login', component: Login },
  { path: '/login/returnUrl/:returnUrl', component: Login },
  { path: '/register', component: Register },
  { path: '/reg', component: Register },
  { path: '/mobile/register', component: MobileRegister },
  { path: '/findpwd', alias: '/findPwd', component: FindPwd },

  // 应用下载
  { path: '/app', component: AppDownload },

  // 红包信封
  { path: '/envelope/:eno', component: Envelope },

  // CTC 交易
  { path: '/ctc', component: Ctc },

  // 挖矿
  { path: '/mining', component: Mining },
  { path: '/mining/:id', component: MiningDetail },

  // 众筹
  { path: '/crowdfunding', component: crowdfundingIndex },
  { path: '/uc/crowdfunding/list', component: crowdfundingList },

  // 活动
  { path: '/activity', alias: '/lab', component: Activity },
  { path: '/activity/:id', alias: '/lab/detail/:id', component: ActivityDetail },
  { path: '/partner', component: Partner },
  { path: '/bzb', component: Bzb },

  // CMS
  { path: '/help', component: Help },
  { path: '/help/list', alias: '/helplist', component: HelpList },
  { path: '/help/detail', alias: '/helpdetail', component: HelpDetail },
  { path: '/notice', component: Notice },
  { path: '/notice/item/:id?', alias: ['/announcement/:id', '/announcement'], component: NoticeItem },
  { path: '/notice/index', component: Noticeindex },
  { path: '/about', alias: '/about-us', component: AboutUs },

  // OTC 交易
  { path: '/otc', component: OtcMain },
  { path: '/otc/trade', redirect: '/otc/trade/usdt' },
  { path: '/otc/trade/:pathMatch(.*)*', component: OtcTrade },
  { path: '/otc/tradeInfo', component: OtcTradeInfo },
  { path: '/checkuser', component: OtcCheckUser },
  { path: '/chat', component: OtcChat },
  { path: '/otc/checkuser/:id', redirect: to => `/checkuser?id=${to.params.id}` },
  { path: '/otc/chat/:id', redirect: to => `/chat?tradeId=${to.params.id}` },
  { path: '/otc/ad/publish', component: OtcAdPublish },
  { path: '/otc/myad', component: OtcMyAd },

  // 交易所
  { path: '/exchange', component: Exchange },
  { path: '/exchange/:pair', component: Exchange, name: 'ExchangePair' },

  // 合约交易
  { path: '/swapindex', component: Swapindex },
  { path: '/swapindex/:pair', component: Swapindex, name: 'SwapPair' },
  { path: '/kline/:pair', component: Kline, name: 'Kline' },

  // 用户中心
  {
    path: '/uc',
    component: MemberCenter,
    children: [
      { path: '', component: Safe },
      { path: 'safe', component: Safe },
      { path: 'account', component: Account },
      { path: 'money', component: MoneyIndex },
      { path: 'contract-money', component: ContractMoneyIndex },
      { path: 'recharge', component: Recharge },
      { path: 'withdraw', component: Withdraw },
      { path: 'withdraw/address', component: WithdrawAddress },
      { path: 'record', component: Record },
      { path: 'contract-record', component: ContractRecord },
      { path: 'ad', component: OtcMyAd },
      { path: 'ad/create', component: OtcAdPublish },
      { path: 'ad/update', component: OtcAdPublish },
      { path: 'order', component: myorder },
      { path: 'entrust/current', component: EntrustCurrent },
      { path: 'entrust/history', component: EntrustHistory },
      { path: 'contract/entrust/current', component: contractEntrustCurrent },
      { path: 'contract/entrust/history', component: contractEntrustHistory },
      { path: 'trade/expand', component: TradeExpand },
      { path: 'min/trade', component: MinTrade },
      { path: 'inviting/min', component: InvitingMin },
      { path: 'pay/dividends', component: PayDividends },
      { path: 'innovation/orders', alias: 'innovation/myorders', component: InnovationOrders },
      { path: 'innovation/minings', alias: 'innovation/myminings', component: InnovationMinings },
      { path: 'ident/business', component: IdentBusiness }
    ]
  },

  // 404 - 重定向到首页
  { path: '/:pathMatch(.*)*', redirect: '/index' }
]

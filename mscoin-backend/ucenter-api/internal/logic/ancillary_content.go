package logic

type helpCategorySeed struct {
	Cate    int64
	TitleCN string
	TitleEN string
}

type contentSeed struct {
	ID         int64
	Cate       int64
	TitleCN    string
	TitleEN    string
	ContentCN  string
	ContentEN  string
	CreateTime string
	IsTop      int64
}

var helpCategorySeeds = []helpCategorySeed{
	{Cate: 1, TitleCN: "新手指南", TitleEN: "Getting Started"},
	{Cate: 2, TitleCN: "账户安全", TitleEN: "Account Security"},
}

var helpArticleSeeds = []contentSeed{
	{
		ID:         1001,
		Cate:       1,
		TitleCN:    "如何完成本地登录与进入交易页",
		TitleEN:    "How to sign in locally and enter spot trading",
		ContentCN:  "<p>使用本地测试账号登录后，进入现货页选择交易对，即可在开发环境里完成完整的下单验证。</p><p>当前联调基线已经恢复到可登录、可看行情、可提交现货委托的状态。</p>",
		ContentEN:  "<p>Sign in with the local test account and open the spot exchange page to validate the order flow in development.</p><p>The local baseline now supports sign-in, market visibility, and spot order submission.</p>",
		CreateTime: "2026-04-09 10:00:00",
	},
	{
		ID:         1002,
		Cate:       1,
		TitleCN:    "开发环境下如何查看帮助与公告内容",
		TitleEN:    "How to verify help and announcement content locally",
		ContentCN:  "<p>帮助中心和公告中心现在由本地后端直接提供内容，不再依赖缺失的线上接口。</p><p>你可以通过帮助首页、帮助列表和公告详情页验证内容链路是否正常。</p>",
		ContentEN:  "<p>The help center and announcement pages are now backed by local backend content instead of missing remote contracts.</p><p>Validate the flow via help home, help list, and announcement detail pages.</p>",
		CreateTime: "2026-04-09 10:05:00",
	},
	{
		ID:         1003,
		Cate:       1,
		TitleCN:    "如何验证现货委托是否写入当前与历史列表",
		TitleEN:    "How to verify spot orders appear in current and history views",
		ContentCN:  "<p>提交限价买单后，当前委托与历史委托都会读取同一条本地订单记录。</p><p>当前委托按交易中状态过滤，历史委托展示该用户该交易对下的所有订单。</p>",
		ContentEN:  "<p>After submitting a limit buy order, both current and history entrust views read from the same local order record.</p><p>Current entrust filters trading orders, while history shows all orders for the selected symbol.</p>",
		CreateTime: "2026-04-09 10:10:00",
	},
	{
		ID:         2001,
		Cate:       2,
		TitleCN:    "本地测试账号的安全设置说明",
		TitleEN:    "Security settings for the local test account",
		ContentCN:  "<p>本地测试账号用于开发环境联调，不可用于任何生产环境。</p><p>如果需要重新验证认证态页面，请使用帮助文档中记录的测试账号重新登录。</p>",
		ContentEN:  "<p>The local test account is only for development validation and must not be used in production.</p><p>Use the documented test credentials whenever you need to re-validate authenticated pages.</p>",
		CreateTime: "2026-04-09 10:15:00",
	},
	{
		ID:         2002,
		Cate:       2,
		TitleCN:    "验证码降级模式的使用范围",
		TitleEN:    "Scope of the local captcha fallback mode",
		ContentCN:  "<p>本地登录页已经支持 fallback captcha 契约，仅在开发环境中用于打通真实登录链路。</p><p>该降级模式不会替换生产验证码实现。</p>",
		ContentEN:  "<p>The login page supports a fallback captcha contract for local development to keep the real sign-in flow working.</p><p>This fallback does not replace the production captcha implementation.</p>",
		CreateTime: "2026-04-09 10:20:00",
	},
}

var announcementSeeds = []contentSeed{
	{
		ID:         3001,
		TitleCN:    "开发环境联调基线已恢复",
		TitleEN:    "Development baseline restored",
		ContentCN:  "<p>当前开发环境已经恢复了 Vue 3 前端、登录链路、本地市场基线与现货委托基础能力。</p><p>本公告用于确认本地联调可以继续围绕真实业务流程推进。</p>",
		ContentEN:  "<p>The development environment has restored the Vue 3 frontend, sign-in flow, local market baseline, and the core spot entrust capability.</p><p>This notice confirms that end-to-end local validation can proceed against the real business flow.</p>",
		CreateTime: "2026-04-09 11:00:00",
		IsTop:      1,
	},
	{
		ID:         3002,
		TitleCN:    "帮助中心与公告中心改为本地后端提供内容",
		TitleEN:    "Help and announcement content now served locally",
		ContentCN:  "<p>为避免开发环境继续依赖缺失的线上内容接口，帮助与公告内容已经收口到本地后端。</p><p>联调时请优先验证路由、接口和页面渲染是否一致。</p>",
		ContentEN:  "<p>To avoid relying on missing remote content contracts, help and announcement content is now served by the local backend.</p><p>During validation, focus on route, API, and page-rendering consistency.</p>",
		CreateTime: "2026-04-09 11:10:00",
		IsTop:      0,
	},
	{
		ID:         3003,
		TitleCN:    "现货交易闭环进入最终验证阶段",
		TitleEN:    "Spot trading closure moved to final validation",
		ContentCN:  "<p>本地测试账号、钱包基线和订单表基础设施已经准备完毕。</p><p>接下来请在现货页提交真实委托，并分别在当前委托和历史委托中验证可见性。</p>",
		ContentEN:  "<p>The local test account, wallet baseline, and order-table foundation are ready.</p><p>Next, submit a real spot order in the exchange page and verify visibility in both current and history entrust views.</p>",
		CreateTime: "2026-04-09 11:20:00",
		IsTop:      0,
	},
}

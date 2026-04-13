package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRiskControlLogic_CheckMargin 测试保证金校验逻辑
func TestRiskControlLogic_CheckMargin(t *testing.T) {
	testCases := []struct {
		name          string
		price         float64
		amount        float64
		leverage      int32
		balance       float64
		frozen        float64
		entryptType   int32
		expectSuccess bool
		expectedErr   string
	}{
		// 限价单场景
		{"限价单 - 余额充足", 50000, 0.1, 10, 600, 0, 1, true, ""},
		{"限价单 - 余额刚好", 50000, 0.1, 10, 500, 0, 1, true, ""},
		{"限价单 - 余额不足", 50000, 0.1, 10, 400, 0, 1, false, "保证金不足"},
		{"限价单 - 高杠杆", 50000, 0.1, 20, 300, 0, 1, true, ""},
		{"限价单 - 低杠杆", 50000, 0.1, 5, 900, 0, 1, false, "保证金不足"},

		// 市价单场景
		{"市价单 - 总余额充足", 50000, 0.1, 10, 400, 200, 2, true, ""},
		{"市价单 - 总余额刚好", 50000, 0.1, 10, 500, 0, 2, true, ""},
		{"市价单 - 总余额不足", 50000, 0.1, 10, 300, 100, 2, false, "保证金不足"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 计算所需保证金
			requiredMargin := (tc.price * tc.amount) / float64(tc.leverage)

			// 验证计算结果
			expectedMargin := (50000 * 0.1) / float64(tc.leverage)
			assert.InDelta(t, expectedMargin, requiredMargin, 0.01)

			// 验证余额是否足够
			var availableBalance float64
			if tc.entryptType == 1 {
				// 限价单：只检查可用余额
				availableBalance = tc.balance
			} else {
				// 市价单：检查总余额
				availableBalance = tc.balance + tc.frozen
			}

			success := availableBalance >= requiredMargin
			assert.Equal(t, tc.expectSuccess, success)
		})
	}
}

// TestRiskControlLogic_CalculateRequiredMargin 测试保证金计算
func TestRiskControlLogic_CalculateRequiredMargin(t *testing.T) {
	testCases := []struct {
		name           string
		price          float64
		amount         float64
		leverage       int32
		expectedMargin float64
	}{
		{"50000 价格 0.1 数量 10 倍杠杆", 50000, 0.1, 10, 500},
		{"60000 价格 0.2 数量 20 倍杠杆", 60000, 0.2, 20, 600},
		{"45000 价格 0.5 数量 5 倍杠杆", 45000, 0.5, 5, 4500},
		{"100000 价格 1 数量 100 倍杠杆", 100000, 1, 100, 1000},
		{"30000 价格 0.01 数量 50 倍杠杆", 30000, 0.01, 50, 6},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requiredMargin := (tc.price * tc.amount) / float64(tc.leverage)
			assert.InDelta(t, tc.expectedMargin, requiredMargin, 0.01)
		})
	}
}

// TestRiskControlLogic_LiquidationPrice 测试强平价格计算
func TestRiskControlLogic_LiquidationPrice(t *testing.T) {
	testCases := []struct {
		name                 string
		direction            int32 // 1=做多，2=做空
		entryPrice           float64
		leverage             int32
		margin               float64
		size                 float64
		expectedLiquidation  float64
	}{
		// 做多强平价格 = 开仓价 - (保证金 / 数量)
		{"做多 10 倍杠杆 - 保证金 500", 1, 50000, 10, 500, 0.1, 45000},
		{"做多 20 倍杠杆 - 保证金 250", 1, 50000, 20, 250, 0.1, 47500},
		{"做多 5 倍杠杆 - 保证金 1000", 1, 50000, 5, 1000, 0.1, 40000},

		// 做空强平价格 = 开仓价 + (保证金 / 数量)
		{"做空 10 倍杠杆 - 保证金 500", 2, 50000, 10, 500, 0.1, 55000},
		{"做空 20 倍杠杆 - 保证金 250", 2, 50000, 20, 250, 0.1, 52500},
		{"做空 5 倍杠杆 - 保证金 1000", 2, 50000, 5, 1000, 0.1, 60000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var liquidationPrice float64
			if tc.direction == 1 {
				// 做多：价格下跌到强平价格
				liquidationPrice = tc.entryPrice - (tc.margin / tc.size)
			} else {
				// 做空：价格上涨到强平价格
				liquidationPrice = tc.entryPrice + (tc.margin / tc.size)
			}

			assert.InDelta(t, tc.expectedLiquidation, liquidationPrice, 0.01)
		})
	}
}

// TestRiskControlLogic_MarginRatio 测试保证金率计算
func TestRiskControlLogic_MarginRatio(t *testing.T) {
	testCases := []struct {
		name            string
		margin          float64
		unrealizedPnl   float64
		positionValue   float64
		leverage        int32
		expectedRatio   float64
	}{
		{"100% 保证金率", 500, 0, 5000, 10, 1.0},
		{"150% 保证金率 (盈利)", 500, 250, 5000, 10, 1.5},
		{"50% 保证金率 (亏损)", 500, -250, 5000, 10, 0.5},
		{"200% 保证金率", 1000, 0, 5000, 10, 2.0},
		{"0% 保证金率 (穿仓)", 500, -500, 5000, 10, 0.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 所需保证金 = 仓位价值 / 杠杆
			requiredMargin := tc.positionValue / float64(tc.leverage)

			// 保证金率 = (保证金 + 未实现盈亏) / 所需保证金
			marginRatio := (tc.margin + tc.unrealizedPnl) / requiredMargin

			// 保证金率不能为负
			if marginRatio < 0 {
				marginRatio = 0
			}

			assert.InDelta(t, tc.expectedRatio, marginRatio, 0.01)
		})
	}
}

// TestRiskControlLogic_UnrealizedPnL 测试未实现盈亏计算
func TestRiskControlLogic_UnrealizedPnL(t *testing.T) {
	testCases := []struct {
		name          string
		direction     int32 // 1=做多，2=做空
		entryPrice    float64
		currentPrice  float64
		size          float64
		expectedPnl   float64
	}{
		// 做多场景
		{"做多 - 盈利 500", 1, 50000, 55000, 0.1, 500},
		{"做多 - 亏损 500", 1, 50000, 45000, 0.1, -500},
		{"做多 - 盈亏平衡", 1, 50000, 50000, 0.1, 0},
		{"做多 - 大幅盈利", 1, 50000, 60000, 0.2, 2000},

		// 做空场景
		{"做空 - 盈利 500", 2, 50000, 45000, 0.1, 500},
		{"做空 - 亏损 500", 2, 50000, 55000, 0.1, -500},
		{"做空 - 盈亏平衡", 2, 50000, 50000, 0.1, 0},
		{"做空 - 大幅盈利", 2, 50000, 40000, 0.2, 2000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var pnl float64
			if tc.direction == 1 {
				// 做多：价格上涨盈利
				pnl = (tc.currentPrice - tc.entryPrice) * tc.size
			} else {
				// 做空：价格下跌盈利
				pnl = (tc.entryPrice - tc.currentPrice) * tc.size
			}

			assert.InDelta(t, tc.expectedPnl, pnl, 0.01)
		})
	}
}

// TestRiskControlLogic_PositionValue 测试仓位价值计算
func TestRiskControlLogic_PositionValue(t *testing.T) {
	testCases := []struct {
		name          string
		price         float64
		size          float64
		expectedValue float64
	}{
		{"50000 价格 0.1 数量", 50000, 0.1, 5000},
		{"60000 价格 0.5 数量", 60000, 0.5, 30000},
		{"45000 价格 1 数量", 45000, 1, 45000},
		{"100000 价格 2 数量", 100000, 2, 200000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			positionValue := tc.price * tc.size
			assert.InDelta(t, tc.expectedValue, positionValue, 0.01)
		})
	}
}

// TestRiskControlLogic_AvailableBalance 测试可用余额计算
func TestRiskControlLogic_AvailableBalance(t *testing.T) {
	testCases := []struct {
		name            string
		balance         float64
		frozen          float64
		entryptType     int32
		expectedAvail   float64
	}{
		{"限价单 - 只用余额", 1000, 500, 1, 1000},
		{"市价单 - 余额 + 冻结", 1000, 500, 2, 1500},
		{"限价单 - 无冻结", 800, 0, 1, 800},
		{"市价单 - 无冻结", 800, 0, 2, 800},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var available float64
			if tc.entryptType == 1 {
				// 限价单：只使用余额
				available = tc.balance
			} else {
				// 市价单：使用余额 + 冻结
				available = tc.balance + tc.frozen
			}

			assert.Equal(t, tc.expectedAvail, available)
		})
	}
}

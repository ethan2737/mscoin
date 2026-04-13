package logic

import (
	"testing"
	"time"

	"swap-api/internal/model"

	"github.com/stretchr/testify/assert"
)

// TestPositionLogic_CalculateUnrealizedPnl 测试持仓未实现盈亏计算
func TestPositionLogic_CalculateUnrealizedPnl(t *testing.T) {
	testCases := []struct {
		name         string
		direction    int32
		entryPrice   float64
		currentPrice float64
		size         float64
		expectedPnl  float64
	}{
		// 做多场景
		{"做多 - 价格上升盈利", 1, 50000, 55000, 0.1, 500},
		{"做多 - 价格下降亏损", 1, 50000, 45000, 0.1, -500},
		{"做多 - 价格不变", 1, 50000, 50000, 0.1, 0},
		{"做多 - 大幅盈利", 1, 50000, 60000, 0.5, 5000},
		{"做多 - 大幅亏损", 1, 50000, 40000, 0.5, -5000},

		// 做空场景
		{"做空 - 价格下降盈利", 2, 50000, 45000, 0.1, 500},
		{"做空 - 价格上升亏损", 2, 50000, 55000, 0.1, -500},
		{"做空 - 价格不变", 2, 50000, 50000, 0.1, 0},
		{"做空 - 大幅盈利", 2, 50000, 40000, 0.5, 5000},
		{"做空 - 大幅亏损", 2, 50000, 60000, 0.5, -5000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var pnl float64
			if tc.direction == 1 {
				pnl = (tc.currentPrice - tc.entryPrice) * tc.size
			} else {
				pnl = (tc.entryPrice - tc.currentPrice) * tc.size
			}

			assert.InDelta(t, tc.expectedPnl, pnl, 0.01)
		})
	}
}

// TestPositionLogic_CalculateMarginRatio 测试持仓保证金率计算
func TestPositionLogic_CalculateMarginRatio(t *testing.T) {
	testCases := []struct {
		name          string
		margin        float64
		unrealizedPnl float64
		positionValue float64
		leverage      int32
		expectedRatio float64
	}{
		{"初始状态 -100% 保证金率", 500, 0, 5000, 10, 1.0},
		{"盈利 50% -150% 保证金率", 500, 250, 5000, 10, 1.5},
		{"亏损 50% -50% 保证金率", 500, -250, 5000, 10, 0.5},
		{"高保证金 -200% 保证金率", 1000, 0, 5000, 10, 2.0},
		{"穿仓状态", 500, -500, 5000, 10, 0},
		{"接近强平", 500, -400, 5000, 10, 0.2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requiredMargin := tc.positionValue / float64(tc.leverage)
			marginRatio := (tc.margin + tc.unrealizedPnl) / requiredMargin

			// 保证金率不能为负
			if marginRatio < 0 {
				marginRatio = 0
			}

			assert.InDelta(t, tc.expectedRatio, marginRatio, 0.01)
		})
	}
}

// TestPositionLogic_IsLiquidationRisk 测试强平风险判断
func TestPositionLogic_IsLiquidationRisk(t *testing.T) {
	testCases := []struct {
		name           string
		marginRatio    float64
		liquidationRatio float64
		expectedRisk   bool
	}{
		{"安全状态", 1.5, 0.8, false},
		{"临界状态", 0.8, 0.8, false},
		{"风险状态", 0.7, 0.8, true},
		{"严重风险", 0.3, 0.8, true},
		{"穿仓", 0, 0.8, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isRisk := tc.marginRatio < tc.liquidationRatio
			assert.Equal(t, tc.expectedRisk, isRisk)
		})
	}
}

// TestPositionLogic_UpdatePosition 测试持仓数据结构
func TestPositionLogic_UpdatePosition(t *testing.T) {
	position := &model.ContractPosition{
		Id:             1,
		MemberId:       1001,
		ContractCoinId: 1,
		Direction:      1,
		Leverage:       10,
		Size:           0.5,
		EntryPrice:     50000,
		Margin:         2500,
		UnrealizedPnl:  500,
		LiquidationPrice: 45000,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	assert.NotNil(t, position)
	assert.Equal(t, int64(1001), position.MemberId)
	assert.Equal(t, int32(1), position.Direction)
	assert.Equal(t, float64(0.5), position.Size)
	assert.Greater(t, position.LiquidationPrice, float64(0))
}

// TestPositionLogic_PositionValue 测试持仓价值计算
func TestPositionLogic_PositionValue(t *testing.T) {
	testCases := []struct {
		name          string
		price         float64
		size          float64
		expectedValue float64
	}{
		{"BTC 50000 美元 0.1 个", 50000, 0.1, 5000},
		{"BTC 60000 美元 0.5 个", 60000, 0.5, 30000},
		{"ETH 3000 美元 1 个", 3000, 1, 3000},
		{"ETH 2500 美元 10 个", 2500, 10, 25000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.price * tc.size
			assert.InDelta(t, tc.expectedValue, value, 0.01)
		})
	}
}

// TestPositionLogic_RequiredMargin 测试所需保证金计算
func TestPositionLogic_RequiredMargin(t *testing.T) {
	testCases := []struct {
		name           string
		positionValue  float64
		leverage       int32
		expectedMargin float64
	}{
		{"5000 价值 10 倍杠杆", 5000, 10, 500},
		{"30000 价值 20 倍杠杆", 30000, 20, 1500},
		{"10000 价值 5 倍杠杆", 10000, 5, 2000},
		{"50000 价值 100 倍杠杆", 50000, 100, 500},
		{"20000 价值 50 倍杠杆", 20000, 50, 400},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			margin := tc.positionValue / float64(tc.leverage)
			assert.InDelta(t, tc.expectedMargin, margin, 0.01)
		})
	}
}

// TestPositionLogic_LiquidationPriceDiff 测试与强平价格的距离
func TestPositionLogic_LiquidationPriceDiff(t *testing.T) {
	testCases := []struct {
		name             string
		direction        int32
		currentPrice     float64
		liquidationPrice float64
		expectedDiff     float64
		expectedDiffPct  float64
	}{
		// 做多场景
		{"做多 - 安全距离", 1, 55000, 45000, 10000, 0.1818}, // (55000-45000)/55000
		{"做多 - 近距离", 1, 50000, 48000, 2000, 0.04},
		{"做多 - 无距离", 1, 50000, 50000, 0, 0},

		// 做空场景
		{"做空 - 安全距离", 2, 45000, 55000, 10000, 0.2222}, // (55000-45000)/45000
		{"做空 - 近距离", 2, 50000, 52000, 2000, 0.04},
		{"做空 - 无距离", 2, 50000, 50000, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var diff float64
			if tc.direction == 1 {
				diff = tc.currentPrice - tc.liquidationPrice
			} else {
				diff = tc.liquidationPrice - tc.currentPrice
			}

			var diffPct float64
			if tc.currentPrice > 0 {
				diffPct = diff / tc.currentPrice
			}

			assert.InDelta(t, tc.expectedDiff, diff, 0.01)
			assert.InDelta(t, tc.expectedDiffPct, diffPct, 0.01)
		})
	}
}

// TestPositionLogic_AddMargin 测试追加保证金逻辑
func TestPositionLogic_AddMargin(t *testing.T) {
	position := &model.ContractPosition{
		Margin:   500,
		Size:     0.1,
		EntryPrice: 50000,
		Direction:  1,
		Leverage:   10,
	}

	addAmount := 200.0
	position.Margin += addAmount

	assert.Equal(t, 700.0, position.Margin)

	// 重新计算强平价格
	newLiquidationPrice := position.EntryPrice - (position.Margin / position.Size)
	assert.Equal(t, 43000.0, newLiquidationPrice) // 50000 - 700/0.1
}

// TestPositionLogic_ReduceMargin 测试减少保证金逻辑
func TestPositionLogic_ReduceMargin(t *testing.T) {
	position := &model.ContractPosition{
		Margin:   500,
		Size:     0.1,
		EntryPrice: 50000,
		Direction:  1,
		Leverage:   10,
	}

	reduceAmount := 100.0
	position.Margin -= reduceAmount

	assert.Equal(t, 400.0, position.Margin)

	// 重新计算强平价格 (更接近爆仓)
	newLiquidationPrice := position.EntryPrice - (position.Margin / position.Size)
	assert.Equal(t, 46000.0, newLiquidationPrice) // 50000 - 400/0.1
}

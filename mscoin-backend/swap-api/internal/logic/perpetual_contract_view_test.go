package logic

import (
	"testing"
	"time"

	"swap-api/internal/model"
	"swap-api/internal/types"

	"github.com/stretchr/testify/assert"
)

func TestResolveTransferTypeSupportsLegacyDirection(t *testing.T) {
	req := &types.ContractTransferReq{
		Direction: 2,
	}

	assert.Equal(t, int32(2), resolveTransferType(req))
}

func TestBuildUcOrderPageMapsVue3Fields(t *testing.T) {
	page := buildUcOrderPage([]*model.ContractOrder{
		{
			Id:             88,
			MemberId:       1001,
			ContractCoinId: 1,
			EntrustType:    1,
			Type:           1,
			Direction:      1,
			Leverage:       20,
			Price:          64000.5,
			Amount:         0.3,
			Status:         1,
			CreatedAt:      time.Date(2026, 4, 13, 10, 0, 0, 0, time.Local),
		},
	}, map[int64]*model.ContractCoin{
		1: {
			Id:           1,
			Symbol:       "BTC/USDT",
			ContractType: 1,
			ShareNumber:  0.001,
		},
	}, 0, 10)

	assert.Equal(t, int64(1), page.TotalElements)
	assert.Equal(t, int64(0), page.Number)
	if assert.Len(t, page.Content, 1) {
		item, ok := page.Content[0].(*types.UcOrderItem)
		if assert.True(t, ok) {
			assert.Equal(t, int64(88), item.OrderId)
			assert.Equal(t, "BTC/USDT", item.ContractCoinName)
			assert.Equal(t, "BUY", item.Direction)
			assert.Equal(t, "OPEN", item.EntrustType)
			assert.Equal(t, "LIMIT_PRICE", item.Type)
			assert.Equal(t, float64(64000.5), item.EntrustPrice)
			assert.Equal(t, float64(0.3), item.Share)
			assert.Equal(t, "ENTRUST_ING", item.Status)
		}
	}
}

func TestBuildSwapPositionItemsMapsSwapIndexFields(t *testing.T) {
	items := buildSwapPositionItems([]*model.ContractPosition{
		{
			Id:               9,
			MemberId:         1001,
			ContractCoinId:   1,
			Direction:        1,
			Leverage:         25,
			Size:             0.5,
			EntryPrice:       62000,
			Margin:           1240,
			UnrealizedPnl:    20,
			LiquidationPrice: 58000,
		},
	}, map[int64]*model.ContractCoin{
		1: {
			Id:           1,
			Symbol:       "BTC/USDT",
			ContractType: 1,
		},
	})

	if assert.Len(t, items, 1) {
		item := items[0]
		assert.Equal(t, int64(9), item.Id)
		assert.Equal(t, int64(1), item.ContractCoinId)
		assert.Equal(t, "BTC/USDT", item.ContractCoinName)
		assert.Equal(t, "ALWAYS", item.ContractCoinType)
		assert.Equal(t, "BUY", item.Direction)
		assert.Equal(t, int32(25), item.Multiple)
		assert.Equal(t, float64(62000), item.OpenPrice)
		assert.Equal(t, float64(1240), item.PrincipalAmount)
		assert.Equal(t, float64(0.5), item.Position)
		assert.Equal(t, float64(0.5), item.AvailablePosition)
		assert.Equal(t, float64(58000), item.LiquidationPrice)
	}
}

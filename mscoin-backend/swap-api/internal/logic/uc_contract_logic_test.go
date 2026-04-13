package logic

import (
	"context"
	"testing"
	"time"

	"swap-api/internal/model"
	"swap-api/internal/svc"
	"swap-api/internal/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcContractLogicCurrentFiltersContractAndType(t *testing.T) {
	coinModel := &stubContractCoinModel{
		coins: map[int64]*model.ContractCoin{
			1: {Id: 1, Symbol: "BTC/USDT", ContractType: 1, Status: 1},
			2: {Id: 2, Symbol: "ETH/USDT", ContractType: 1, Status: 1},
		},
	}
	orderModel := &stubContractOrderModel{
		orders: map[int64]*model.ContractOrder{
			1: {
				Id:             1,
				MemberId:       1001,
				ContractCoinId: 1,
				EntrustType:    1,
				Type:           1,
				Direction:      1,
				Leverage:       10,
				Price:          100,
				Amount:         1,
				Status:         1,
				CreatedAt:      time.Date(2026, 4, 13, 12, 0, 0, 0, time.Local),
			},
			2: {
				Id:             2,
				MemberId:       1001,
				ContractCoinId: 1,
				EntrustType:    2,
				Type:           1,
				Direction:      1,
				Leverage:       10,
				Price:          200,
				Amount:         1,
				Status:         1,
				CreatedAt:      time.Date(2026, 4, 13, 12, 10, 0, 0, time.Local),
			},
			3: {
				Id:             3,
				MemberId:       1001,
				ContractCoinId: 2,
				EntrustType:    1,
				Type:           1,
				Direction:      1,
				Leverage:       10,
				Price:          300,
				Amount:         1,
				Status:         1,
				CreatedAt:      time.Date(2026, 4, 13, 12, 20, 0, 0, time.Local),
			},
		},
	}

	logic := NewUcContractLogic(context.Background(), &svc.ServiceContext{
		ContractOrderModel:    orderModel,
		ContractPositionModel: &stubContractPositionModel{},
		ContractCoinModel:     coinModel,
	})

	result, err := logic.Current(&types.CurrentOrderReq{
		MemberId:       1001,
		ContractCoinId: 1,
		Type:           2,
		PageNo:         0,
		PageSize:       10,
	})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, int64(1), result.TotalElements)
	if assert.Len(t, result.Content, 1) {
		item := result.Content[0].(*types.UcOrderItem)
		assert.Equal(t, int64(1), item.OrderId)
	}
}

func TestUcContractLogicHistoryFiltersContractAndType(t *testing.T) {
	coinModel := &stubContractCoinModel{
		coins: map[int64]*model.ContractCoin{
			1: {Id: 1, Symbol: "BTC/USDT", ContractType: 1, Status: 1},
		},
	}
	orderModel := &stubContractOrderModel{
		orders: map[int64]*model.ContractOrder{
			1: {
				Id:             1,
				MemberId:       1001,
				ContractCoinId: 1,
				EntrustType:    1,
				Type:           1,
				Direction:      1,
				Leverage:       10,
				Price:          100,
				Amount:         1,
				Status:         2,
				CreatedAt:      time.Date(2026, 4, 13, 12, 0, 0, 0, time.Local),
			},
			2: {
				Id:             2,
				MemberId:       1001,
				ContractCoinId: 1,
				EntrustType:    2,
				Type:           1,
				Direction:      1,
				Leverage:       10,
				Price:          200,
				Amount:         1,
				Status:         3,
				CreatedAt:      time.Date(2026, 4, 13, 12, 10, 0, 0, time.Local),
			},
		},
	}

	logic := NewUcContractLogic(context.Background(), &svc.ServiceContext{
		ContractOrderModel:    orderModel,
		ContractPositionModel: &stubContractPositionModel{},
		ContractCoinModel:     coinModel,
	})

	result, err := logic.History(&types.HistoryOrderReq{
		MemberId:       1001,
		ContractCoinId: 1,
		Type:           2,
		PageNo:         0,
		PageSize:       10,
	})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, int64(1), result.TotalElements)
	if assert.Len(t, result.Content, 1) {
		item := result.Content[0].(*types.UcOrderItem)
		assert.Equal(t, int64(1), item.OrderId)
		assert.Equal(t, "ENTRUST_SUCCESS", item.Status)
	}
}

func TestUcContractLogicThumbReturnsCoinList(t *testing.T) {
	logic := NewUcContractLogic(context.Background(), &svc.ServiceContext{
		ContractOrderModel:    &stubContractOrderModel{},
		ContractPositionModel: &stubContractPositionModel{},
		ContractCoinModel: &stubContractCoinModel{
			coins: map[int64]*model.ContractCoin{
				1: {Id: 1, Symbol: "BTC/USDT", Status: 1},
				2: {Id: 2, Symbol: "ETH/USDT", Status: 1},
			},
		},
	})

	resp, err := logic.Thumb(&types.GetContractListReq{})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Len(t, *resp, 2)
	assert.Equal(t, "BTC/USDT", (*resp)[0].Name)
}

func TestUcContractLogicCancelUpdatesCurrentOrderStatus(t *testing.T) {
	orderModel := &stubContractOrderModel{
		orders: map[int64]*model.ContractOrder{
			9: {
				Id:             9,
				MemberId:       1001,
				ContractCoinId: 1,
				EntrustType:    1,
				Type:           1,
				Direction:      1,
				Leverage:       10,
				Price:          100,
				Amount:         1,
				Status:         1,
			},
		},
	}
	logic := NewUcContractLogic(context.Background(), &svc.ServiceContext{
		ContractOrderModel:    orderModel,
		ContractPositionModel: &stubContractPositionModel{},
		ContractCoinModel:     &stubContractCoinModel{},
	})

	resp, err := logic.Cancel(&types.CancelOrderReq{
		OrderId:  9,
		MemberId: 1001,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.Success)
	assert.Equal(t, int32(3), orderModel.orders[9].Status)
}

package logic

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"swap-api/internal/model"
	"swap-api/internal/svc"
	"swap-api/internal/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type trackingContractTransactionModel struct {
	stubContractTransactionModel
	lastMemberID int64
	lastStart    time.Time
	lastEnd      time.Time
	lastType     int32
	lastLimit    int64
	lastOffset   int64
}

func (m *trackingContractTransactionModel) FindByMemberId(ctx context.Context, memberId int64, startTime, endTime time.Time, txType int32, limit, offset int64) ([]*model.ContractTransaction, error) {
	m.lastMemberID = memberId
	m.lastStart = startTime
	m.lastEnd = endTime
	m.lastType = txType
	m.lastLimit = limit
	m.lastOffset = offset
	return m.stubContractTransactionModel.FindByMemberId(ctx, memberId, startTime, endTime, txType, limit, offset)
}

func TestContractWalletLogicTransactionPassesFiltersToDao(t *testing.T) {
	transactionModel := &trackingContractTransactionModel{
		stubContractTransactionModel: stubContractTransactionModel{
			transactions: map[int64]*model.ContractTransaction{
				1: {
					Id:        1,
					MemberId:  1001,
					Unit:      "USDT",
					Type:      2,
					Amount:    88,
					CreatedAt: time.Date(2026, 4, 12, 10, 0, 0, 0, time.Local),
				},
			},
		},
	}

	svcCtx := &svc.ServiceContext{
		ContractWalletModel:      &stubContractWalletModel{},
		ContractTransactionModel: transactionModel,
	}
	logic := NewContractWalletLogic(context.Background(), svcCtx)

	result, err := logic.Transaction(&types.ContractTransactionReq{
		MemberId:  1001,
		PageNo:    2,
		PageSize:  5,
		StartTime: "1712822400000",
		EndTime:   "1712995200000",
		Type:      json.RawMessage("2"),
	})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, int32(2), transactionModel.lastType)
	assert.Equal(t, int64(5), transactionModel.lastLimit)
	assert.Equal(t, int64(5), transactionModel.lastOffset)
	assert.False(t, transactionModel.lastStart.IsZero())
	assert.False(t, transactionModel.lastEnd.IsZero())
}

func TestContractWalletLogicInfoBuildsWalletItems(t *testing.T) {
	svcCtx := &svc.ServiceContext{
		ContractWalletModel: &stubContractWalletModel{
			wallets: map[int64]*model.ContractWallet{
				1: {Id: 1, MemberId: 1001, Unit: "USDT", Balance: 1200, Frozen: 80, MainBalance: 1280},
			},
		},
		ContractTransactionModel: &stubContractTransactionModel{},
	}
	logic := NewContractWalletLogic(context.Background(), svcCtx)

	items, err := logic.Info(&types.ContractWalletReq{MemberId: 1001})

	require.NoError(t, err)
	if assert.Len(t, items, 1) {
		assert.Equal(t, "USDT", items[0].Coin.Unit)
		assert.Equal(t, 1200.0, items[0].Balance)
		assert.Equal(t, 80.0, items[0].FrozenBalance)
	}
}

func TestContractWalletLogicTransferSupportsLegacyDirectionField(t *testing.T) {
	walletModel := &stubContractWalletModel{
		wallets: map[int64]*model.ContractWallet{
			1: {Id: 1, MemberId: 1001, Unit: "USDT", Balance: 100, Frozen: 0, MainBalance: 100},
		},
	}
	svcCtx := &svc.ServiceContext{
		ContractWalletModel:      walletModel,
		ContractTransactionModel: &stubContractTransactionModel{},
	}
	logic := NewContractWalletLogic(context.Background(), svcCtx)

	resp, err := logic.Transfer(&types.ContractTransferReq{
		MemberId:  1001,
		Unit:      "USDT",
		Amount:    30,
		Direction: 1,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.Success)
	assert.Equal(t, 130.0, walletModel.wallets[1].Balance)
}

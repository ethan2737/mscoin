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

type stubContractCoinModel struct {
	coins map[int64]*model.ContractCoin
}

func (m *stubContractCoinModel) Insert(ctx context.Context, coin *model.ContractCoin) (int64, error) {
	if m.coins == nil {
		m.coins = make(map[int64]*model.ContractCoin)
	}
	if coin.Id == 0 {
		coin.Id = int64(len(m.coins) + 1)
	}
	m.coins[coin.Id] = coin
	return coin.Id, nil
}

func (m *stubContractCoinModel) FindOne(ctx context.Context, id int64) (*model.ContractCoin, error) {
	return m.coins[id], nil
}

func (m *stubContractCoinModel) FindBySymbol(ctx context.Context, symbol string) (*model.ContractCoin, error) {
	for _, coin := range m.coins {
		if coin.Symbol == symbol {
			return coin, nil
		}
	}
	return nil, nil
}

func (m *stubContractCoinModel) FindAll(ctx context.Context) ([]*model.ContractCoin, error) {
	result := make([]*model.ContractCoin, 0, len(m.coins))
	for _, coin := range m.coins {
		result = append(result, coin)
	}
	return result, nil
}

func (m *stubContractCoinModel) Update(ctx context.Context, coin *model.ContractCoin) error {
	m.coins[coin.Id] = coin
	return nil
}

func (m *stubContractCoinModel) Delete(ctx context.Context, id int64) error {
	delete(m.coins, id)
	return nil
}

type stubContractOrderModel struct {
	orders map[int64]*model.ContractOrder
	nextID int64
}

func (m *stubContractOrderModel) Insert(ctx context.Context, order *model.ContractOrder) (int64, error) {
	if m.orders == nil {
		m.orders = make(map[int64]*model.ContractOrder)
	}
	m.nextID++
	order.Id = m.nextID
	if order.CreatedAt.IsZero() {
		order.CreatedAt = time.Date(2026, 4, 13, 12, 0, 0, 0, time.Local)
	}
	m.orders[order.Id] = order
	return order.Id, nil
}

func (m *stubContractOrderModel) FindOne(ctx context.Context, id int64) (*model.ContractOrder, error) {
	return m.orders[id], nil
}

func (m *stubContractOrderModel) FindByMemberId(ctx context.Context, memberId int64, status int32) ([]*model.ContractOrder, error) {
	result := make([]*model.ContractOrder, 0)
	for _, order := range m.orders {
		if order.MemberId != memberId {
			continue
		}
		if status != 0 && order.Status != status {
			continue
		}
		result = append(result, order)
	}
	return result, nil
}

func (m *stubContractOrderModel) FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*model.ContractOrder, error) {
	result := make([]*model.ContractOrder, 0)
	for _, order := range m.orders {
		if order.ContractCoinId == contractCoinId {
			result = append(result, order)
		}
	}
	return result, nil
}

func (m *stubContractOrderModel) UpdateStatus(ctx context.Context, id int64, status int32, dealAmount float64, dealAmountUsdt float64, fee float64, profit float64) error {
	order := m.orders[id]
	if order == nil {
		return nil
	}
	order.Status = status
	order.DealAmount = dealAmount
	order.DealAmountUsdt = dealAmountUsdt
	order.Fee = fee
	order.Profit = profit
	return nil
}

func (m *stubContractOrderModel) Delete(ctx context.Context, id int64) error {
	delete(m.orders, id)
	return nil
}

type stubContractWalletModel struct {
	wallets map[int64]*model.ContractWallet
}

func (m *stubContractWalletModel) Insert(ctx context.Context, wallet *model.ContractWallet) (int64, error) {
	if m.wallets == nil {
		m.wallets = make(map[int64]*model.ContractWallet)
	}
	if wallet.Id == 0 {
		wallet.Id = int64(len(m.wallets) + 1)
	}
	m.wallets[wallet.Id] = wallet
	return wallet.Id, nil
}

func (m *stubContractWalletModel) FindOne(ctx context.Context, id int64) (*model.ContractWallet, error) {
	return m.wallets[id], nil
}

func (m *stubContractWalletModel) FindByMemberId(ctx context.Context, memberId int64) ([]*model.ContractWallet, error) {
	result := make([]*model.ContractWallet, 0)
	for _, wallet := range m.wallets {
		if wallet.MemberId == memberId {
			result = append(result, wallet)
		}
	}
	return result, nil
}

func (m *stubContractWalletModel) FindByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*model.ContractWallet, error) {
	for _, wallet := range m.wallets {
		if wallet.MemberId == memberId && wallet.Unit == unit {
			return wallet, nil
		}
	}
	return nil, nil
}

func (m *stubContractWalletModel) UpdateBalance(ctx context.Context, id int64, balance float64, frozen float64, mainBalance float64) error {
	wallet := m.wallets[id]
	if wallet == nil {
		return nil
	}
	wallet.Balance = balance
	wallet.Frozen = frozen
	wallet.MainBalance = mainBalance
	return nil
}

func (m *stubContractWalletModel) Upsert(ctx context.Context, wallet *model.ContractWallet) error {
	if m.wallets == nil {
		m.wallets = make(map[int64]*model.ContractWallet)
	}
	if wallet.Id == 0 {
		wallet.Id = int64(len(m.wallets) + 1)
	}
	m.wallets[wallet.Id] = wallet
	return nil
}

func (m *stubContractWalletModel) Delete(ctx context.Context, id int64) error {
	delete(m.wallets, id)
	return nil
}

type stubContractPositionModel struct {
	positions map[int64]*model.ContractPosition
}

func (m *stubContractPositionModel) Insert(ctx context.Context, position *model.ContractPosition) (int64, error) {
	if m.positions == nil {
		m.positions = make(map[int64]*model.ContractPosition)
	}
	if position.Id == 0 {
		position.Id = int64(len(m.positions) + 1)
	}
	m.positions[position.Id] = position
	return position.Id, nil
}

func (m *stubContractPositionModel) FindOne(ctx context.Context, id int64) (*model.ContractPosition, error) {
	return m.positions[id], nil
}

func (m *stubContractPositionModel) FindByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*model.ContractPosition, error) {
	for _, position := range m.positions {
		if position.MemberId == memberId && position.ContractCoinId == contractCoinId && position.Direction == direction {
			return position, nil
		}
	}
	return nil, nil
}

func (m *stubContractPositionModel) FindByMemberId(ctx context.Context, memberId int64) ([]*model.ContractPosition, error) {
	result := make([]*model.ContractPosition, 0)
	for _, position := range m.positions {
		if position.MemberId == memberId {
			result = append(result, position)
		}
	}
	return result, nil
}

func (m *stubContractPositionModel) Update(ctx context.Context, position *model.ContractPosition) error {
	m.positions[position.Id] = position
	return nil
}

func (m *stubContractPositionModel) Upsert(ctx context.Context, position *model.ContractPosition) error {
	if m.positions == nil {
		m.positions = make(map[int64]*model.ContractPosition)
	}
	m.positions[position.Id] = position
	return nil
}

func (m *stubContractPositionModel) Delete(ctx context.Context, id int64) error {
	delete(m.positions, id)
	return nil
}

type stubContractTransactionModel struct {
	transactions map[int64]*model.ContractTransaction
	nextID       int64
}

func (m *stubContractTransactionModel) Insert(ctx context.Context, tx *model.ContractTransaction) (int64, error) {
	if m.transactions == nil {
		m.transactions = make(map[int64]*model.ContractTransaction)
	}
	m.nextID++
	tx.Id = m.nextID
	if tx.CreatedAt.IsZero() {
		tx.CreatedAt = time.Date(2026, 4, 13, 12, 0, 0, 0, time.Local)
	}
	m.transactions[tx.Id] = tx
	return tx.Id, nil
}

func (m *stubContractTransactionModel) FindOne(ctx context.Context, id int64) (*model.ContractTransaction, error) {
	return m.transactions[id], nil
}

func (m *stubContractTransactionModel) FindByMemberId(ctx context.Context, memberId int64, startTime, endTime time.Time, txType int32, limit, offset int64) ([]*model.ContractTransaction, error) {
	result := make([]*model.ContractTransaction, 0)
	for _, tx := range m.transactions {
		if tx.MemberId != memberId {
			continue
		}
		result = append(result, tx)
	}
	return result, nil
}

func (m *stubContractTransactionModel) Delete(ctx context.Context, id int64) error {
	delete(m.transactions, id)
	return nil
}

func newTestOrderEntrustLogic() (*OrderEntrustLogic, *stubContractOrderModel, *stubContractWalletModel, *stubContractPositionModel, *stubContractTransactionModel) {
	coinModel := &stubContractCoinModel{
		coins: map[int64]*model.ContractCoin{
			1: {
				Id:           1,
				Symbol:       "BTC/USDT",
				ContractType: 1,
				MinLeverage:  10,
				Status:       1,
			},
		},
	}
	orderModel := &stubContractOrderModel{orders: make(map[int64]*model.ContractOrder)}
	walletModel := &stubContractWalletModel{
		wallets: map[int64]*model.ContractWallet{
			1: {Id: 1, MemberId: 1001, Unit: "USDT", Balance: 1000, Frozen: 0, MainBalance: 1000},
		},
	}
	positionModel := &stubContractPositionModel{positions: make(map[int64]*model.ContractPosition)}
	transactionModel := &stubContractTransactionModel{transactions: make(map[int64]*model.ContractTransaction)}

	svcCtx := &svc.ServiceContext{
		ContractCoinModel:        coinModel,
		ContractOrderModel:       orderModel,
		ContractWalletModel:      walletModel,
		ContractPositionModel:    positionModel,
		ContractTransactionModel: transactionModel,
	}

	return NewOrderEntrustLogic(context.Background(), svcCtx), orderModel, walletModel, positionModel, transactionModel
}

func TestOrderEntrustLogicAddFreezesWalletAndCreatesOrder(t *testing.T) {
	logic, orderModel, walletModel, _, _ := newTestOrderEntrustLogic()

	resp, err := logic.Add(&types.AddOrderReq{
		MemberId:       1001,
		ContractCoinId: 1,
		EntrustType:    1,
		Type:           1,
		Direction:      1,
		Leverage:       10,
		Price:          50000,
		Amount:         0.1,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.OrderId)

	wallet := walletModel.wallets[1]
	require.NotNil(t, wallet)
	assert.Equal(t, 500.0, wallet.Balance)
	assert.Equal(t, 500.0, wallet.Frozen)

	order := orderModel.orders[resp.OrderId]
	require.NotNil(t, order)
	assert.Equal(t, int64(1001), order.MemberId)
	assert.Equal(t, int32(1), order.Status)
	assert.Equal(t, float64(50000), order.Price)
}

func TestOrderEntrustLogicCurrentFiltersContractAndMapsVue3Fields(t *testing.T) {
	logic, orderModel, _, _, _ := newTestOrderEntrustLogic()
	orderModel.orders[1] = &model.ContractOrder{
		Id:             1,
		MemberId:       1001,
		ContractCoinId: 1,
		EntrustType:    1,
		Type:           1,
		Direction:      1,
		Leverage:       10,
		Price:          12345,
		Amount:         0.05,
		Status:         1,
		CreatedAt:      time.Date(2026, 4, 13, 14, 0, 0, 0, time.Local),
	}
	orderModel.orders[2] = &model.ContractOrder{
		Id:             2,
		MemberId:       1001,
		ContractCoinId: 1,
		EntrustType:    2,
		Type:           1,
		Direction:      1,
		Leverage:       10,
		Price:          10000,
		Amount:         0.03,
		Status:         1,
		CreatedAt:      time.Date(2026, 4, 13, 15, 0, 0, 0, time.Local),
	}

	resp, err := logic.Current(&types.CurrentOrderReq{
		MemberId:       1001,
		ContractCoinId: 1,
		Type:           2,
		PageNo:         0,
		PageSize:       10,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	if assert.Len(t, resp.Content, 1) {
		item := resp.Content[0]
		assert.Equal(t, int64(1), item.OrderId)
		assert.Equal(t, "BTC/USDT", item.ContractCoinName)
		assert.Equal(t, "OPEN", item.EntrustType)
		assert.Equal(t, "LIMIT_PRICE", item.Type)
		assert.Equal(t, float64(12345), item.EntrustPrice)
	}
	assert.Equal(t, int64(1), resp.TotalElements)
}

func TestOrderEntrustLogicQuickCloseUpdatesPositionWalletAndTransaction(t *testing.T) {
	logic, orderModel, walletModel, positionModel, transactionModel := newTestOrderEntrustLogic()
	positionModel.positions[1] = &model.ContractPosition{
		Id:               1,
		MemberId:         1001,
		ContractCoinId:   1,
		Direction:        1,
		Leverage:         10,
		Size:             0.2,
		EntryPrice:       50000,
		Margin:           1000,
		UnrealizedPnl:    0,
		LiquidationPrice: 45000,
	}

	resp, err := logic.QuickClose(&types.QuickCloseReq{
		MemberId:       1001,
		ContractCoinId: 1,
		Price:          51000,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.Success)

	position := positionModel.positions[1]
	require.NotNil(t, position)
	assert.Equal(t, 0.0, position.Size)
	assert.Equal(t, 0.0, position.UnrealizedPnl)

	wallet := walletModel.wallets[1]
	require.NotNil(t, wallet)
	assert.Equal(t, 1200.0, wallet.Balance)

	require.Len(t, transactionModel.transactions, 1)
	for _, tx := range transactionModel.transactions {
		assert.Equal(t, int64(1001), tx.MemberId)
		assert.Equal(t, int32(3), tx.Type)
		assert.Equal(t, 200.0, tx.Amount)
	}

	require.Len(t, orderModel.orders, 1)
	for _, order := range orderModel.orders {
		assert.Equal(t, int32(2), order.Status)
		assert.Equal(t, float64(0.2), order.Amount)
	}
}

func TestOrderEntrustLogicCancelRejectsForeignOrder(t *testing.T) {
	logic, orderModel, _, _, _ := newTestOrderEntrustLogic()
	orderModel.orders[1] = &model.ContractOrder{
		Id:             1,
		MemberId:       2002,
		ContractCoinId: 1,
		EntrustType:    1,
		Type:           1,
		Direction:      1,
		Leverage:       10,
		Price:          50000,
		Amount:         0.1,
		Status:         1,
	}

	resp, err := logic.Cancel(&types.CancelOrderReq{
		OrderId:  1,
		MemberId: 1001,
	})

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, int32(1), orderModel.orders[1].Status)
}

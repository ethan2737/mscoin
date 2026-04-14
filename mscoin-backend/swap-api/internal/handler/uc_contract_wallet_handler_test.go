package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"grpc-common/market/mclient"
	marketpb "grpc-common/market/types/market"
	"swap-api/internal/marketdata"
	"swap-api/internal/midd"
	"swap-api/internal/model"
	"swap-api/internal/svc"
	"swap-api/internal/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type handlerStubContractCoinModel struct {
	coins map[int64]*model.ContractCoin
}

func (m *handlerStubContractCoinModel) Insert(ctx context.Context, coin *model.ContractCoin) (int64, error) {
	if m.coins == nil {
		m.coins = make(map[int64]*model.ContractCoin)
	}
	if coin.Id == 0 {
		coin.Id = int64(len(m.coins) + 1)
	}
	m.coins[coin.Id] = coin
	return coin.Id, nil
}

func (m *handlerStubContractCoinModel) FindOne(ctx context.Context, id int64) (*model.ContractCoin, error) {
	return m.coins[id], nil
}

func (m *handlerStubContractCoinModel) FindBySymbol(ctx context.Context, symbol string) (*model.ContractCoin, error) {
	for _, coin := range m.coins {
		if coin.Symbol == symbol {
			return coin, nil
		}
	}
	return nil, nil
}

func (m *handlerStubContractCoinModel) FindAll(ctx context.Context) ([]*model.ContractCoin, error) {
	result := make([]*model.ContractCoin, 0, len(m.coins))
	for _, coin := range m.coins {
		result = append(result, coin)
	}
	return result, nil
}

func (m *handlerStubContractCoinModel) Update(ctx context.Context, coin *model.ContractCoin) error {
	m.coins[coin.Id] = coin
	return nil
}

func (m *handlerStubContractCoinModel) Delete(ctx context.Context, id int64) error {
	delete(m.coins, id)
	return nil
}

type handlerStubContractOrderModel struct {
	orders map[int64]*model.ContractOrder
}

func (m *handlerStubContractOrderModel) Insert(ctx context.Context, order *model.ContractOrder) (int64, error) {
	if m.orders == nil {
		m.orders = make(map[int64]*model.ContractOrder)
	}
	if order.Id == 0 {
		order.Id = int64(len(m.orders) + 1)
	}
	if order.CreatedAt.IsZero() {
		order.CreatedAt = time.Date(2026, 4, 13, 12, 0, 0, 0, time.Local)
	}
	m.orders[order.Id] = order
	return order.Id, nil
}

func (m *handlerStubContractOrderModel) FindOne(ctx context.Context, id int64) (*model.ContractOrder, error) {
	return m.orders[id], nil
}

func (m *handlerStubContractOrderModel) FindByMemberId(ctx context.Context, memberId int64, status int32) ([]*model.ContractOrder, error) {
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

func (m *handlerStubContractOrderModel) FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*model.ContractOrder, error) {
	result := make([]*model.ContractOrder, 0)
	for _, order := range m.orders {
		if order.ContractCoinId == contractCoinId {
			result = append(result, order)
		}
	}
	return result, nil
}

func (m *handlerStubContractOrderModel) UpdateStatus(ctx context.Context, id int64, status int32, dealAmount float64, dealAmountUsdt float64, fee float64, profit float64) error {
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

func (m *handlerStubContractOrderModel) Delete(ctx context.Context, id int64) error {
	delete(m.orders, id)
	return nil
}

type handlerStubContractPositionModel struct{}

func (m *handlerStubContractPositionModel) Insert(ctx context.Context, position *model.ContractPosition) (int64, error) {
	return 0, nil
}

func (m *handlerStubContractPositionModel) FindOne(ctx context.Context, id int64) (*model.ContractPosition, error) {
	return nil, nil
}

func (m *handlerStubContractPositionModel) FindByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*model.ContractPosition, error) {
	return nil, nil
}

func (m *handlerStubContractPositionModel) FindByMemberId(ctx context.Context, memberId int64) ([]*model.ContractPosition, error) {
	return nil, nil
}

func (m *handlerStubContractPositionModel) Update(ctx context.Context, position *model.ContractPosition) error {
	return nil
}

func (m *handlerStubContractPositionModel) Upsert(ctx context.Context, position *model.ContractPosition) error {
	return nil
}

func (m *handlerStubContractPositionModel) Delete(ctx context.Context, id int64) error {
	return nil
}

type handlerStubContractWalletModel struct {
	wallets map[int64]*model.ContractWallet
}

func (m *handlerStubContractWalletModel) Insert(ctx context.Context, wallet *model.ContractWallet) (int64, error) {
	if m.wallets == nil {
		m.wallets = make(map[int64]*model.ContractWallet)
	}
	if wallet.Id == 0 {
		wallet.Id = int64(len(m.wallets) + 1)
	}
	m.wallets[wallet.Id] = wallet
	return wallet.Id, nil
}

func (m *handlerStubContractWalletModel) FindOne(ctx context.Context, id int64) (*model.ContractWallet, error) {
	return m.wallets[id], nil
}

func (m *handlerStubContractWalletModel) FindByMemberId(ctx context.Context, memberId int64) ([]*model.ContractWallet, error) {
	result := make([]*model.ContractWallet, 0)
	for _, wallet := range m.wallets {
		if wallet.MemberId == memberId {
			result = append(result, wallet)
		}
	}
	return result, nil
}

func (m *handlerStubContractWalletModel) FindByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*model.ContractWallet, error) {
	for _, wallet := range m.wallets {
		if wallet.MemberId == memberId && wallet.Unit == unit {
			return wallet, nil
		}
	}
	return nil, nil
}

func (m *handlerStubContractWalletModel) UpdateBalance(ctx context.Context, id int64, balance float64, frozen float64, mainBalance float64) error {
	wallet := m.wallets[id]
	if wallet == nil {
		return nil
	}
	wallet.Balance = balance
	wallet.Frozen = frozen
	wallet.MainBalance = mainBalance
	return nil
}

func (m *handlerStubContractWalletModel) Upsert(ctx context.Context, wallet *model.ContractWallet) error {
	if m.wallets == nil {
		m.wallets = make(map[int64]*model.ContractWallet)
	}
	m.wallets[wallet.Id] = wallet
	return nil
}

func (m *handlerStubContractWalletModel) Delete(ctx context.Context, id int64) error {
	delete(m.wallets, id)
	return nil
}

type handlerStubContractTransactionModel struct {
	transactions map[int64]*model.ContractTransaction
}

type handlerStubMarketRPC struct {
	historyFn func(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.HistoryRes, error)
}

type handlerStubMarketDataClient struct {
	thumbs     []*marketdata.SymbolThumb
	plate      *marketdata.TradePlate
	trades     []*marketdata.LatestTrade
	thumbErr   error
	plateErr   error
	tradeErr   error
	lastSymbol string
	lastFull   bool
	lastSize   int
}

func (m *handlerStubMarketDataClient) LoadSymbolThumb(ctx context.Context) ([]*marketdata.SymbolThumb, error) {
	if m.thumbErr != nil {
		return nil, m.thumbErr
	}
	return m.thumbs, nil
}

func (m *handlerStubMarketDataClient) LoadTradePlate(ctx context.Context, symbol string, full bool) (*marketdata.TradePlate, error) {
	m.lastSymbol = symbol
	m.lastFull = full
	if m.plateErr != nil {
		return nil, m.plateErr
	}
	return m.plate, nil
}

func (m *handlerStubMarketDataClient) LoadLatestTrade(ctx context.Context, symbol string, size int) ([]*marketdata.LatestTrade, error) {
	m.lastSymbol = symbol
	m.lastSize = size
	if m.tradeErr != nil {
		return nil, m.tradeErr
	}
	return m.trades, nil
}

func (m *handlerStubMarketRPC) FindSymbolThumbTrend(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.SymbolThumbRes, error) {
	return nil, errors.New("not implemented")
}

func (m *handlerStubMarketRPC) FindSymbolInfo(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.ExchangeCoin, error) {
	return nil, errors.New("not implemented")
}

func (m *handlerStubMarketRPC) FindCoinInfo(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.Coin, error) {
	return nil, errors.New("not implemented")
}

func (m *handlerStubMarketRPC) HistoryKline(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.HistoryRes, error) {
	if m.historyFn == nil {
		return nil, errors.New("not implemented")
	}
	return m.historyFn(ctx, req, opts...)
}

func (m *handlerStubMarketRPC) FindExchangeCoinVisible(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.ExchangeCoinRes, error) {
	return nil, errors.New("not implemented")
}

func (m *handlerStubMarketRPC) FindAllCoin(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.CoinList, error) {
	return nil, errors.New("not implemented")
}

func (m *handlerStubMarketRPC) FindCoinById(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.Coin, error) {
	return nil, errors.New("not implemented")
}

func (m *handlerStubContractTransactionModel) Insert(ctx context.Context, tx *model.ContractTransaction) (int64, error) {
	if m.transactions == nil {
		m.transactions = make(map[int64]*model.ContractTransaction)
	}
	if tx.Id == 0 {
		tx.Id = int64(len(m.transactions) + 1)
	}
	if tx.CreatedAt.IsZero() {
		tx.CreatedAt = time.Date(2026, 4, 13, 12, 0, 0, 0, time.Local)
	}
	m.transactions[tx.Id] = tx
	return tx.Id, nil
}

func (m *handlerStubContractTransactionModel) FindOne(ctx context.Context, id int64) (*model.ContractTransaction, error) {
	return m.transactions[id], nil
}

func (m *handlerStubContractTransactionModel) FindByMemberId(ctx context.Context, memberId int64, startTime, endTime time.Time, txType int32, limit, offset int64) ([]*model.ContractTransaction, error) {
	filtered := make([]*model.ContractTransaction, 0)
	for _, tx := range m.transactions {
		if tx.MemberId != memberId {
			continue
		}
		if txType != -1 && tx.Type != txType {
			continue
		}
		if !startTime.IsZero() && tx.CreatedAt.Before(startTime) {
			continue
		}
		if !endTime.IsZero() && tx.CreatedAt.After(endTime) {
			continue
		}
		filtered = append(filtered, tx)
	}
	if offset >= int64(len(filtered)) {
		return []*model.ContractTransaction{}, nil
	}
	end := offset + limit
	if end > int64(len(filtered)) {
		end = int64(len(filtered))
	}
	return filtered[offset:end], nil
}

func (m *handlerStubContractTransactionModel) Delete(ctx context.Context, id int64) error {
	delete(m.transactions, id)
	return nil
}

type handlerPageResult struct {
	Content       []types.UcOrderItem `json:"content"`
	TotalElements int64               `json:"totalElements"`
}

type handlerThumbResult []types.ContractCoinThumbItem

type handlerTransactionPage struct {
	Content       []types.ContractTransactionItem `json:"content"`
	TotalElements int64                           `json:"totalElements"`
}

type handlerResult[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func newHandlerServiceContext() (*svc.ServiceContext, *handlerStubContractOrderModel, *handlerStubContractTransactionModel) {
	orderModel := &handlerStubContractOrderModel{
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
				Status:         3,
				CreatedAt:      time.Date(2026, 4, 13, 12, 10, 0, 0, time.Local),
			},
			3: {
				Id:             3,
				MemberId:       2002,
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

	transactionModel := &handlerStubContractTransactionModel{
		transactions: map[int64]*model.ContractTransaction{
			1: {
				Id:        1,
				MemberId:  1001,
				Unit:      "BTC/USDT",
				Type:      2,
				Amount:    88,
				CreatedAt: time.UnixMilli(1712908800000),
			},
			2: {
				Id:        2,
				MemberId:  2002,
				Unit:      "ETH/USDT",
				Type:      2,
				Amount:    99,
				CreatedAt: time.UnixMilli(1712908800000),
			},
			3: {
				Id:        3,
				MemberId:  1001,
				Unit:      "BTC/USDT",
				Type:      1,
				Amount:    66,
				CreatedAt: time.UnixMilli(1712908800000),
			},
		},
	}

	return &svc.ServiceContext{
		ContractCoinModel: &handlerStubContractCoinModel{
			coins: map[int64]*model.ContractCoin{
				1: {Id: 1, Symbol: "BTC/USDT", ContractType: 1, Status: 1},
				2: {Id: 2, Symbol: "ETH/USDT", ContractType: 1, Status: 1},
			},
		},
		ContractOrderModel:       orderModel,
		ContractPositionModel:    &handlerStubContractPositionModel{},
		ContractWalletModel:      &handlerStubContractWalletModel{},
		ContractTransactionModel: transactionModel,
	}, orderModel, transactionModel
}

func performJSONRequest(t *testing.T, method string, path string, body any, authUserID int64, handlerFunc http.HandlerFunc) *httptest.ResponseRecorder {
	t.Helper()

	var payload []byte
	var err error
	if body == nil {
		payload = []byte("{}")
	} else {
		payload, err = json.Marshal(body)
		require.NoError(t, err)
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	if authUserID > 0 {
		req = req.WithContext(context.WithValue(req.Context(), midd.AuthUserIDKey, authUserID))
	}

	recorder := httptest.NewRecorder()
	handlerFunc(recorder, req)
	return recorder
}

func TestUcContractCurrentHandlerUsesAuthUserAndBuildsPageResponse(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	recorder := performJSONRequest(t, http.MethodPost, "/uc/contract/current", map[string]any{
		"memberId":       9999,
		"contractCoinId": 1,
		"type":           2,
		"pageNo":         0,
		"pageSize":       10,
	}, 1001, NewUcContractHandler(svcCtx).Current)

	require.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())

	var result handlerResult[handlerPageResult]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, int64(1), result.Data.TotalElements)
	if assert.Len(t, result.Data.Content, 1) {
		assert.Equal(t, int64(1), result.Data.Content[0].OrderId)
		assert.Equal(t, "BTC/USDT", result.Data.Content[0].ContractCoinName)
	}
}

func TestUcContractHistoryHandlerBuildsFilteredPageResponse(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	recorder := performJSONRequest(t, http.MethodPost, "/uc/contract/history", map[string]any{
		"memberId":       9999,
		"contractCoinId": 1,
		"type":           3,
		"pageNo":         0,
		"pageSize":       10,
	}, 1001, NewUcContractHandler(svcCtx).History)

	require.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())

	var result handlerResult[handlerPageResult]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, int64(1), result.Data.TotalElements)
	if assert.Len(t, result.Data.Content, 1) {
		assert.Equal(t, int64(2), result.Data.Content[0].OrderId)
		assert.Equal(t, "ENTRUST_CANCEL", result.Data.Content[0].Status)
	}
}

func TestUcContractThumbHandlerReturnsThumbItems(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	recorder := performJSONRequest(t, http.MethodPost, "/uc/contract/coin/thumb", map[string]any{}, 1001, NewUcContractHandler(svcCtx).Thumb)

	require.Equalf(t, http.StatusOK, recorder.Code, "response body: %s", recorder.Body.String())

	var result handlerResult[handlerThumbResult]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	assert.Equal(t, 0, result.Code)
	if assert.Len(t, result.Data, 2) {
		assert.Equal(t, int64(1), result.Data[0].Id)
		assert.Equal(t, "BTC/USDT", result.Data[0].Name)
	}
}

func TestUcContractCancelHandlerUsesPathOrderIDAndUpdatesStatus(t *testing.T) {
	svcCtx, orderModel, _ := newHandlerServiceContext()
	recorder := performJSONRequest(t, http.MethodPost, "/uc/contract/entrust/cancel/1", map[string]any{
		"memberId": 9999,
	}, 1001, NewUcContractHandler(svcCtx).Cancel)

	require.Equal(t, http.StatusOK, recorder.Code)

	var result handlerResult[types.CancelOrderResp]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	assert.Equal(t, 0, result.Code)
	assert.True(t, result.Data.Success)
	assert.Equal(t, int32(3), orderModel.orders[1].Status)
}

func TestContractTransactionHandlerUsesAuthUserAndReturnsPagedContent(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	recorder := performJSONRequest(t, http.MethodPost, "/uc/asset/contract-transaction/all", map[string]any{
		"memberId":  9999,
		"type":      2,
		"startTime": "1712900000000",
		"endTime":   "1712910000000",
		"pageNo":    1,
		"pageSize":  10,
	}, 1001, NewContractWalletHandler(svcCtx).Transaction)

	require.Equal(t, http.StatusOK, recorder.Code)

	var result handlerResult[handlerTransactionPage]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, int64(1), result.Data.TotalElements)
	if assert.Len(t, result.Data.Content, 1) {
		assert.Equal(t, "BTC/USDT", result.Data.Content[0].Symbol)
		assert.Equal(t, int32(2), result.Data.Content[0].Type)
		assert.Equal(t, float64(88), result.Data.Content[0].Amount)
	}
}

func TestKlineHistoryHandlerUsesMarketRPCData(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	svcCtx.MarketRpc = &handlerStubMarketRPC{
		historyFn: func(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.HistoryRes, error) {
			assert.Equal(t, "BTC/USDT", req.Symbol)
			assert.Equal(t, int64(1710000000000), req.From)
			assert.Equal(t, int64(1710003600000), req.To)
			assert.Equal(t, "60", req.Resolution)
			return &mclient.HistoryRes{
				List: []*marketpb.History{
					{
						Time:   1710000000000,
						Open:   100,
						High:   120,
						Low:    90,
						Close:  110,
						Volume: 10,
					},
				},
			}, nil
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/swap/history?symbol=BTC%2FUSDT&from=1710000000000&to=1710003600000&resolution=60", nil)
	recorder := httptest.NewRecorder()
	klineHistoryHandler(svcCtx)(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

	var result handlerResult[[][]any]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	require.Len(t, result.Data, 1)
	assert.EqualValues(t, 1710000000000, result.Data[0][0])
	assert.EqualValues(t, 110, result.Data[0][4])
}

func TestKlineHistoryHandlerFallsBackWhenMarketRPCUnavailable(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	svcCtx.MarketRpc = &handlerStubMarketRPC{
		historyFn: func(ctx context.Context, req *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.HistoryRes, error) {
			return nil, errors.New("rpc unavailable")
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/swap/history?symbol=BTC%2FUSDT&from=1710000000000&to=1710003600000&resolution=60", nil)
	recorder := httptest.NewRecorder()
	klineHistoryHandler(svcCtx)(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

	var result handlerResult[[][]any]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	require.NotEmpty(t, result.Data)
	assert.EqualValues(t, 1710000000000, result.Data[0][0])
}

func TestSwapSymbolThumbUsesRealMarketThumbData(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	svcCtx.MarketDataClient = &handlerStubMarketDataClient{
		thumbs: []*marketdata.SymbolThumb{
			{
				Symbol: "BTC/USDT",
				Close:  70748.8,
				High:   71500.5,
				Low:    69888.1,
				Volume: 9876.54,
				Chg:    0.03456,
			},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/swap/symbol-thumb", nil)
	recorder := httptest.NewRecorder()
	NewContractInfoHandler(svcCtx).List(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

	var result handlerResult[[]map[string]any]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	require.NotEmpty(t, result.Data)
	var btc map[string]any
	for _, item := range result.Data {
		if item["symbol"] == "BTC/USDT" {
			btc = item
			break
		}
	}
	require.NotNil(t, btc)
	assert.EqualValues(t, 70748.8, btc["close"])
	assert.EqualValues(t, 71500.5, btc["high"])
	assert.EqualValues(t, 69888.1, btc["low"])
	assert.EqualValues(t, 9876.54, btc["volume"])
	assert.Equal(t, "+3.46%", btc["rose"])
}

func TestExchangePlateMiniHandlerUsesRealMarketPlate(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	marketClient := &handlerStubMarketDataClient{
		plate: &marketdata.TradePlate{
			Ask: &marketdata.TradePlateSide{
				Items: []*marketdata.TradePlateItem{
					{Price: 70760, Amount: 0.3},
					{Price: 70761, Amount: 0.7},
				},
			},
			Bid: &marketdata.TradePlateSide{
				Items: []*marketdata.TradePlateItem{
					{Price: 70750, Amount: 0.2},
					{Price: 70749, Amount: 0.4},
				},
			},
		},
	}
	svcCtx.MarketDataClient = marketClient

	recorder := performJSONRequest(t, http.MethodPost, "/swap/exchange-plate-mini", map[string]any{
		"symbol": "BTC/USDT",
	}, 0, exchangePlateMiniHandler(svcCtx))

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "BTC/USDT", marketClient.lastSymbol)
	assert.False(t, marketClient.lastFull)

	var result handlerResult[map[string]any]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	asks, ok := result.Data["asks"].([]any)
	require.True(t, ok)
	require.Len(t, asks, 2)
	firstAsk, ok := asks[0].(map[string]any)
	require.True(t, ok)
	secondAsk, ok := asks[1].(map[string]any)
	require.True(t, ok)
	assert.EqualValues(t, 0.3, firstAsk["totalAmount"])
	assert.EqualValues(t, 1.0, secondAsk["totalAmount"])
}

func TestLatestTradeHandlerUsesRealMarketTrades(t *testing.T) {
	svcCtx, _, _ := newHandlerServiceContext()
	marketClient := &handlerStubMarketDataClient{
		trades: []*marketdata.LatestTrade{
			{Symbol: "BTC/USDT", Price: 70748.8, Amount: 0.12, Time: 1713000000000, Direction: "BUY"},
			{Symbol: "BTC/USDT", Price: 70749.1, Amount: 0.33, Time: 1713000001000, Direction: "SELL"},
		},
	}
	svcCtx.MarketDataClient = marketClient

	recorder := performJSONRequest(t, http.MethodPost, "/swap/latest-trade", map[string]any{
		"symbol": "BTC/USDT",
		"size":   2,
	}, 0, latestTradeHandler(svcCtx))

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "BTC/USDT", marketClient.lastSymbol)
	assert.Equal(t, 2, marketClient.lastSize)

	var result handlerResult[[]marketdata.LatestTrade]
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &result))
	require.Len(t, result.Data, 2)
	assert.EqualValues(t, 70748.8, result.Data[0].Price)
	assert.Equal(t, "BUY", result.Data[0].Direction)
}

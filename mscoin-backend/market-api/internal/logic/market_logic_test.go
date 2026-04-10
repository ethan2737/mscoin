package logic

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"grpc-common/market/mclient"
	marketpb "grpc-common/market/types/market"
	"market-api/internal/model"
	"market-api/internal/processor"
	"market-api/internal/repo"
	"market-api/internal/svc"
	"market-api/internal/types"
)

type fakeProcessor struct {
	thumb  any
	plate  *model.TradePlateGroup
	trades []*model.LatestTrade
}

func (f *fakeProcessor) GetThumb() any { return f.thumb }
func (f *fakeProcessor) GetTradePlate(symbol string, size int) *model.TradePlateGroup {
	return f.plate
}
func (f *fakeProcessor) GetLatestTrade(symbol string, size int) []*model.LatestTrade {
	if size <= 0 || size > len(f.trades) {
		size = len(f.trades)
	}
	return f.trades[:size]
}
func (f *fakeProcessor) Process(data processor.ProcessData)   {}
func (f *fakeProcessor) AddHandler(h processor.MarketHandler) {}

type fakeMarketClient struct {
	thumbs []*marketpb.CoinThumb
	called bool
}

type fakeSnapshotStore struct {
	plate  *model.TradePlateGroup
	trades []*model.LatestTrade
}

func (f *fakeSnapshotStore) LoadTradePlate(ctx context.Context, symbol string, size int) (*model.TradePlateGroup, error) {
	return f.plate, nil
}

func (f *fakeSnapshotStore) LoadLatestTrade(ctx context.Context, symbol string, size int) ([]*model.LatestTrade, error) {
	if size <= 0 || size > len(f.trades) {
		size = len(f.trades)
	}
	return f.trades[:size], nil
}

func (f *fakeMarketClient) FindSymbolThumbTrend(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.SymbolThumbRes, error) {
	f.called = true
	return &mclient.SymbolThumbRes{List: f.thumbs}, nil
}

func (f *fakeMarketClient) FindSymbolInfo(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.ExchangeCoin, error) {
	return nil, nil
}

func (f *fakeMarketClient) FindCoinInfo(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.Coin, error) {
	return nil, nil
}

func (f *fakeMarketClient) HistoryKline(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.HistoryRes, error) {
	return nil, nil
}

func (f *fakeMarketClient) FindExchangeCoinVisible(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.ExchangeCoinRes, error) {
	return nil, nil
}

func (f *fakeMarketClient) FindAllCoin(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.CoinList, error) {
	return nil, nil
}

func (f *fakeMarketClient) FindCoinById(ctx context.Context, in *mclient.MarketReq, opts ...grpc.CallOption) (*mclient.Coin, error) {
	return nil, nil
}

func TestSymbolThumbTrendBypassesProcessorCache(t *testing.T) {
	marketClient := &fakeMarketClient{
		thumbs: []*marketpb.CoinThumb{
			{
				Symbol:       "BTC/USDT",
				Close:        105,
				LastDayClose: 100,
				Change:       5,
				Chg:          0.05,
			},
		},
	}
	processorCache := &fakeProcessor{
		thumb: []*marketpb.CoinThumb{
			{
				Symbol:       "BTC/USDT",
				Close:        105,
				LastDayClose: 100,
				Change:       0,
				Chg:          0,
			},
		},
	}
	logic := &MarketLogic{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			MarketRpc: marketClient,
			Processor: processorCache,
		},
	}

	resp, err := logic.SymbolThumbTrend(&types.MarketReq{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !marketClient.called {
		t.Fatalf("expected SymbolThumbTrend to call market rpc instead of processor cache")
	}
	if len(resp) != 1 {
		t.Fatalf("expected one thumb, got %d", len(resp))
	}
	if resp[0].Change != 5 || resp[0].Chg != 0.05 {
		t.Fatalf("expected fresh rpc thumb data, got change=%v chg=%v", resp[0].Change, resp[0].Chg)
	}
}

func TestExchangePlateMiniReadsProcessorCache(t *testing.T) {
	logic := &MarketLogic{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Processor: &fakeProcessor{
				plate: &model.TradePlateGroup{
					Ask: &model.TradePlateResult{
						Direction: "SELL",
						Symbol:    "BTC/USDT",
						Items: []*model.TradePlateItem{
							{Price: 100, Amount: 2},
						},
					},
					Bid: &model.TradePlateResult{
						Direction: "BUY",
						Symbol:    "BTC/USDT",
						Items: []*model.TradePlateItem{
							{Price: 99, Amount: 3},
						},
					},
				},
			},
		},
	}

	resp, err := logic.ExchangePlateMini(&types.MarketReq{Symbol: "BTC/USDT"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Ask == nil || len(resp.Ask.Items) != 1 {
		t.Fatalf("expected ask plate to be returned from processor cache")
	}
	if resp.Bid == nil || len(resp.Bid.Items) != 1 {
		t.Fatalf("expected bid plate to be returned from processor cache")
	}
	if resp.Ask.Items[0].Price != 100 || resp.Bid.Items[0].Price != 99 {
		t.Fatalf("unexpected plate prices: ask=%v bid=%v", resp.Ask.Items[0].Price, resp.Bid.Items[0].Price)
	}
}

func TestLatestTradeReadsProcessorCache(t *testing.T) {
	logic := &MarketLogic{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Processor: &fakeProcessor{
				trades: []*model.LatestTrade{
					{Symbol: "BTC/USDT", Price: 100, Amount: 1.5, Time: 111, Direction: "BUY"},
					{Symbol: "BTC/USDT", Price: 99, Amount: 2.5, Time: 112, Direction: "SELL"},
				},
			},
		},
	}

	resp, err := logic.LatestTrade(&types.MarketReq{Symbol: "BTC/USDT", Size: 1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(resp) != 1 {
		t.Fatalf("expected one trade, got %d", len(resp))
	}
	if resp[0].Price != 100 || resp[0].Direction != "BUY" {
		t.Fatalf("unexpected latest trade payload: %+v", resp[0])
	}
}

func TestExchangePlateMiniFallsBackToSnapshotStore(t *testing.T) {
	logic := &MarketLogic{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Processor: &fakeProcessor{},
			SnapshotStore: &fakeSnapshotStore{
				plate: &model.TradePlateGroup{
					Ask: &model.TradePlateResult{
						Direction: "SELL",
						Symbol:    "SOL/USDT",
						Items: []*model.TradePlateItem{
							{Price: 90, Amount: 2},
						},
					},
					Bid: &model.TradePlateResult{
						Direction: "BUY",
						Symbol:    "SOL/USDT",
						Items: []*model.TradePlateItem{
							{Price: 89, Amount: 3},
						},
					},
				},
			},
		},
	}

	resp, err := logic.ExchangePlateMini(&types.MarketReq{Symbol: "SOL/USDT"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Ask == nil || len(resp.Ask.Items) != 1 || resp.Ask.Items[0].Price != 90 {
		t.Fatalf("expected ask plate from snapshot store, got %+v", resp.Ask)
	}
	if resp.Bid == nil || len(resp.Bid.Items) != 1 || resp.Bid.Items[0].Price != 89 {
		t.Fatalf("expected bid plate from snapshot store, got %+v", resp.Bid)
	}
}

func TestLatestTradeFallsBackToSnapshotStore(t *testing.T) {
	logic := &MarketLogic{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Processor: &fakeProcessor{},
			SnapshotStore: &fakeSnapshotStore{
				trades: []*model.LatestTrade{
					{Symbol: "SOL/USDT", Price: 82.5, Amount: 8, Time: 123, Direction: "BUY"},
				},
			},
		},
	}

	resp, err := logic.LatestTrade(&types.MarketReq{Symbol: "SOL/USDT", Size: 20})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(resp) != 1 || resp[0].Price != 82.5 || resp[0].Direction != "BUY" {
		t.Fatalf("expected latest trade from snapshot store, got %+v", resp)
	}
}

var _ repo.TradeSnapshotStore = (*fakeSnapshotStore)(nil)

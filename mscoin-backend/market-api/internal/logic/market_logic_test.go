package logic

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"grpc-common/market/mclient"
	marketpb "grpc-common/market/types/market"
	"market-api/internal/processor"
	"market-api/internal/svc"
	"market-api/internal/types"
)

type fakeProcessor struct {
	thumb any
}

func (f *fakeProcessor) GetThumb() any                        { return f.thumb }
func (f *fakeProcessor) Process(data processor.ProcessData)   {}
func (f *fakeProcessor) AddHandler(h processor.MarketHandler) {}

type fakeMarketClient struct {
	thumbs []*marketpb.CoinThumb
	called bool
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

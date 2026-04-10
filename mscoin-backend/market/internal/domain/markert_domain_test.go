package domain

import (
	"context"
	"testing"
	"time"

	"market/internal/model"
)

type fakeKlineRepo struct {
	from   int64
	end    int64
	symbol string
	period string
	klines []*model.Kline
}

func (f *fakeKlineRepo) FindBySymbol(ctx context.Context, symbol, period string, count int64) ([]*model.Kline, error) {
	return nil, nil
}

func (f *fakeKlineRepo) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, sort string) ([]*model.Kline, error) {
	f.from = from
	f.end = end
	f.symbol = symbol
	f.period = period
	return f.klines, nil
}

func TestSymbolThumbTrendUsesRolling24HourWindow(t *testing.T) {
	now := time.Now()
	repo := &fakeKlineRepo{
		klines: []*model.Kline{
			{ClosePrice: 105, HighestPrice: 108, LowestPrice: 101, Volume: 10, Turnover: 1000, Time: now.UnixMilli()},
			{ClosePrice: 100, HighestPrice: 102, LowestPrice: 99, Volume: 9, Turnover: 900, Time: now.Add(-23 * time.Hour).UnixMilli()},
		},
	}
	domain := &MarketDomain{klineRepo: repo}

	thumbs := domain.SymbolThumbTrend([]*model.ExchangeCoin{{Symbol: "BTC/USDT"}})

	if len(thumbs) != 1 {
		t.Fatalf("expected one thumb, got %d", len(thumbs))
	}
	if repo.symbol != "BTC/USDT" || repo.period != "1H" {
		t.Fatalf("unexpected query target: symbol=%s period=%s", repo.symbol, repo.period)
	}

	window := repo.end - repo.from
	expectedWindow := int64((24 * time.Hour).Milliseconds())
	diff := window - expectedWindow
	if diff < 0 {
		diff = -diff
	}
	if diff > int64(time.Minute) {
		t.Fatalf("expected rolling 24h window, got %dms", window)
	}
}

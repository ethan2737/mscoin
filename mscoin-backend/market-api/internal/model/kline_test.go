package model

import (
	"testing"

	"grpc-common/market/types/market"
)

func TestToCoinThumbUsesExistingLastDayClose(t *testing.T) {
	kline := &Kline{
		OpenPrice:    102,
		ClosePrice:   105,
		HighestPrice: 108,
		LowestPrice:  101,
		Time:         1712707200000,
		Volume:       10,
		Turnover:     1000,
	}
	thumb := &market.CoinThumb{
		Symbol:       "BTC/USDT",
		Close:        100,
		LastDayClose: 100,
		High:         106,
		Low:          99,
		Volume:       20,
		Turnover:     2000,
		DateTime:     1712703600000,
	}

	next := kline.ToCoinThumb("BTC/USDT", thumb)

	if next.LastDayClose != 100 {
		t.Fatalf("expected lastDayClose to remain 100, got %v", next.LastDayClose)
	}
	if next.Change != 5 {
		t.Fatalf("expected 24h change amount to be 5, got %v", next.Change)
	}
	if next.Chg != 0.05 {
		t.Fatalf("expected 24h change ratio to be 0.05, got %v", next.Chg)
	}
}

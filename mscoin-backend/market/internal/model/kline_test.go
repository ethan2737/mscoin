package model

import "testing"

func TestToCoinThumbUsesLastDayCloseFor24HourChange(t *testing.T) {
	latest := &Kline{
		OpenPrice:    102,
		ClosePrice:   105,
		HighestPrice: 108,
		LowestPrice:  101,
		Time:         1712707200000,
	}
	lastDay := &Kline{
		OpenPrice:    98,
		ClosePrice:   100,
		HighestPrice: 101,
		LowestPrice:  96,
		Time:         1712620800000,
	}

	thumb := latest.ToCoinThumb("BTC/USDT", lastDay)

	if thumb.Symbol != "BTC/USDT" {
		t.Fatalf("unexpected symbol: %s", thumb.Symbol)
	}
	if thumb.LastDayClose != 100 {
		t.Fatalf("expected lastDayClose to be 100, got %v", thumb.LastDayClose)
	}
	if thumb.Change != 5 {
		t.Fatalf("expected 24h change amount to be 5, got %v", thumb.Change)
	}
	if thumb.Chg != 0.05 {
		t.Fatalf("expected 24h change ratio to be 0.05, got %v", thumb.Chg)
	}
}

func TestToCoinThumbHandlesMissingLastDayClose(t *testing.T) {
	latest := &Kline{
		OpenPrice:  80,
		ClosePrice: 82,
		Time:       1712707200000,
	}
	lastDay := &Kline{}

	thumb := latest.ToCoinThumb("SOL/USDT", lastDay)

	if thumb.LastDayClose != 0 {
		t.Fatalf("expected lastDayClose to stay 0, got %v", thumb.LastDayClose)
	}
	if thumb.Change != 0 {
		t.Fatalf("expected change to stay 0 when reference close is missing, got %v", thumb.Change)
	}
	if thumb.Chg != 0 {
		t.Fatalf("expected chg to stay 0 when reference close is missing, got %v", thumb.Chg)
	}
}

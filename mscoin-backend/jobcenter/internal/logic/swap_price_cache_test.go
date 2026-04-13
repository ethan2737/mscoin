package logic

import "testing"

func TestBuildSwapRateCacheKey(t *testing.T) {
	if got := buildSwapRateCacheKey("BTC/USDT"); got != "SWAP::BTC::USDT::SWAP::RATE" {
		t.Fatalf("unexpected cache key: %s", got)
	}
}

func TestParseSwapPriceString(t *testing.T) {
	if got := parseSwapPriceString("64123.45"); got != 64123.45 {
		t.Fatalf("unexpected parsed price: %v", got)
	}
	if got := parseSwapPriceString("bad-value"); got != 0 {
		t.Fatalf("invalid price should fall back to 0, got: %v", got)
	}
}

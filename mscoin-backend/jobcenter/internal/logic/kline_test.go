package logic

import (
	"errors"
	"testing"

	"grpc-common/market/types/market"
)

func TestNormalizeSymbolTargetsUsesVisibleCoinsAndDeduplicates(t *testing.T) {
	targets := normalizeSymbolTargets([]*market.ExchangeCoin{
		{Symbol: "BTC/USDT"},
		{Symbol: "ETH/USDT"},
		{Symbol: "BTC/USDT"},
		{Symbol: ""},
		{Symbol: "SOL/ETH"},
	})

	if len(targets) != 3 {
		t.Fatalf("expected 3 targets, got %d", len(targets))
	}

	if targets[0].Symbol != "BTC/USDT" || targets[0].InstID != "BTC-USDT" {
		t.Fatalf("unexpected first target: %+v", targets[0])
	}

	if targets[1].Symbol != "ETH/USDT" || targets[1].InstID != "ETH-USDT" {
		t.Fatalf("unexpected second target: %+v", targets[1])
	}

	if targets[2].Symbol != "SOL/ETH" || targets[2].InstID != "SOL-ETH" {
		t.Fatalf("unexpected third target: %+v", targets[2])
	}
}

func TestGetWithProxyFallbackRetriesDirectWhenProxyFails(t *testing.T) {
	callOrder := make([]string, 0, 2)
	request := func(path string, headers map[string]string, proxy string) ([]byte, error) {
		callOrder = append(callOrder, proxy)
		if proxy != "" {
			return nil, errors.New("proxy refused")
		}
		return []byte(`{"code":"0"}`), nil
	}

	body, err := getWithProxyFallback(request, "https://www.okx.com/api/v5/market/candles", map[string]string{
		"X-Test": "1",
	}, "http://localhost:10809")
	if err != nil {
		t.Fatalf("expected fallback request to succeed, got %v", err)
	}

	if string(body) != `{"code":"0"}` {
		t.Fatalf("unexpected body: %s", string(body))
	}

	if len(callOrder) != 2 || callOrder[0] != "http://localhost:10809" || callOrder[1] != "" {
		t.Fatalf("unexpected call order: %#v", callOrder)
	}
}

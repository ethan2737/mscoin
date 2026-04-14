package marketdata

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	LoadSymbolThumb(ctx context.Context) ([]*SymbolThumb, error)
	LoadTradePlate(ctx context.Context, symbol string, full bool) (*TradePlate, error)
	LoadLatestTrade(ctx context.Context, symbol string, size int) ([]*LatestTrade, error)
}

type SymbolThumb struct {
	Symbol      string    `json:"symbol"`
	Open        float64   `json:"open"`
	High        float64   `json:"high"`
	Low         float64   `json:"low"`
	Close       float64   `json:"close"`
	Chg         float64   `json:"chg"`
	Change      float64   `json:"change"`
	Volume      float64   `json:"volume"`
	Turnover    float64   `json:"turnover"`
	UsdRate     float64   `json:"usdRate"`
	BaseUsdRate float64   `json:"baseUsdRate"`
	Zone        int       `json:"zone"`
	Trend       []float64 `json:"trend"`
}

type TradePlateItem struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type TradePlateSide struct {
	Direction    string            `json:"direction"`
	MaxAmount    float64           `json:"maxAmount"`
	MinAmount    float64           `json:"minAmount"`
	HighestPrice float64           `json:"highestPrice"`
	LowestPrice  float64           `json:"lowestPrice"`
	Symbol       string            `json:"symbol"`
	Items        []*TradePlateItem `json:"items"`
}

type TradePlate struct {
	Ask *TradePlateSide `json:"ask"`
	Bid *TradePlateSide `json:"bid"`
}

type LatestTrade struct {
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Time      int64   `json:"time"`
	Direction string  `json:"direction"`
}

type httpClient struct {
	baseURL string
	client  *http.Client
}

type apiEnvelope struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func New(baseURL string) Client {
	baseURL = strings.TrimRight(baseURL, "/")
	if baseURL == "" {
		return nil
	}
	return &httpClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *httpClient) LoadSymbolThumb(ctx context.Context) ([]*SymbolThumb, error) {
	var thumbs []*SymbolThumb
	if err := c.postJSON(ctx, "/symbol-thumb", map[string]any{}, &thumbs); err != nil {
		return nil, err
	}
	return thumbs, nil
}

func (c *httpClient) LoadTradePlate(ctx context.Context, symbol string, full bool) (*TradePlate, error) {
	path := "/exchange-plate-mini"
	if full {
		path = "/exchange-plate-full"
	}
	var plate TradePlate
	if err := c.postJSON(ctx, path, map[string]any{"symbol": symbol}, &plate); err != nil {
		return nil, err
	}
	return &plate, nil
}

func (c *httpClient) LoadLatestTrade(ctx context.Context, symbol string, size int) ([]*LatestTrade, error) {
	var trades []*LatestTrade
	if err := c.postJSON(ctx, "/latest-trade", map[string]any{
		"symbol": symbol,
		"size":   size,
	}, &trades); err != nil {
		return nil, err
	}
	return trades, nil
}

func (c *httpClient) postJSON(ctx context.Context, path string, payload any, target any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("market api status=%d body=%s", resp.StatusCode, string(data))
	}
	var envelope apiEnvelope
	if err := json.Unmarshal(data, &envelope); err != nil {
		return err
	}
	if envelope.Code != 0 {
		if envelope.Message == "" {
			envelope.Message = "market api returned non-zero code"
		}
		return errors.New(envelope.Message)
	}
	if target == nil || len(envelope.Data) == 0 || string(envelope.Data) == "null" {
		return nil
	}
	return json.Unmarshal(envelope.Data, target)
}

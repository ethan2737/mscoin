package logic

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"grpc-common/market/mclient"
	markettypes "grpc-common/market/types/market"
	"jobcenter/internal/database"
	"jobcenter/internal/domain"
	"mscoin-common/tools"
)

type SwapKlineConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type SwapKlineResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type swapSymbolTarget struct {
	Symbol string
	InstID string
}

var swapFallbackSymbolTargets = []swapSymbolTarget{
	{Symbol: "BTC/USDT", InstID: "BTC-USDT-SWAP"},
	{Symbol: "ETH/USDT", InstID: "ETH-USDT-SWAP"},
	{Symbol: "SOL/USDT", InstID: "SOL-USDT-SWAP"},
}

type SwapKline struct {
	wg          sync.WaitGroup
	c           SwapKlineConfig
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
	ch          cache.Cache
	marketRpc   mclient.Market
}

func normalizeSwapSymbolTargets(coins []*markettypes.ExchangeCoin) []swapSymbolTarget {
	targets := make([]swapSymbolTarget, 0, len(coins))
	seen := make(map[string]struct{}, len(coins))
	for _, coin := range coins {
		if coin == nil {
			continue
		}

		symbol := strings.TrimSpace(coin.Symbol)
		if symbol == "" {
			continue
		}

		if _, ok := seen[symbol]; ok {
			continue
		}

		seen[symbol] = struct{}{}
		targets = append(targets, swapSymbolTarget{
			Symbol: symbol,
			InstID: strings.ReplaceAll(symbol, "/", "-") + "-SWAP",
		})
	}

	return targets
}

func swapGetWithProxyFallback(request func(string, map[string]string, string) ([]byte, error), path string, headers map[string]string, proxy string) ([]byte, error) {
	if proxy == "" {
		return request(path, headers, "")
	}

	body, err := request(path, headers, proxy)
	if err != nil {
		log.Printf("okx request via proxy failed, retrying direct connection: %v", err)
		return request(path, headers, "")
	}

	return body, nil
}

func (k *SwapKline) loadTargets() []swapSymbolTarget {
	if k.marketRpc == nil {
		return append([]swapSymbolTarget{}, swapFallbackSymbolTargets...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 使用 FindExchangeCoinVisible 获取可见的合约
	resp, err := k.marketRpc.FindExchangeCoinVisible(ctx, &mclient.MarketReq{})
	if err != nil {
		log.Printf("load visible swap coins failed, fallback to defaults: %v", err)
		return append([]swapSymbolTarget{}, swapFallbackSymbolTargets...)
	}

	targets := normalizeSwapSymbolTargets(resp.List)
	if len(targets) == 0 {
		return append([]swapSymbolTarget{}, swapFallbackSymbolTargets...)
	}

	return targets
}

func (k *SwapKline) Do(period string) {
	targets := k.loadTargets()
	if len(targets) == 0 {
		return
	}

	k.wg.Add(len(targets))
	for _, target := range targets {
		current := target
		go k.getKlineData(current, period)
	}
	k.wg.Wait()
}

func (k *SwapKline) getKlineData(target swapSymbolTarget, period string) {
	api := k.c.Host + "/api/v5/market/candles?instId=" + target.InstID + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+target.InstID+"&bar="+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass

	resp, err := swapGetWithProxyFallback(tools.GetWithHeader, api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}

	var result SwapKlineResult
	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}

	log.Println("==================执行存储 swap kline====================")
	if result.Code == "0" {
		// 存储到 MongoDB 的 swap_klines 集合
		k.klineDomain.SaveSwapKline(result.Data, target.Symbol, period)
		if period == "1m" && len(result.Data) > 0 {
			data := result.Data[0]
			k.queueDomain.SendSwapKline(data, target.Symbol)
			k.ch.Set(buildSwapRateCacheKey(target.Symbol), data[4])
		}
	}

	k.wg.Done()
	log.Println("==================End====================")
}

func NewSwapKline(c SwapKlineConfig, marketRpc mclient.Market, mongoClient *database.MongoClient, kafkaCli *database.KafkaClient, cache2 cache.Cache) *SwapKline {
	return &SwapKline{
		c:           c,
		klineDomain: domain.NewKlineDomain(mongoClient),
		queueDomain: domain.NewQueueDomain(kafkaCli),
		ch:          cache2,
		marketRpc:   marketRpc,
	}
}

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

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type symbolTarget struct {
	Symbol string
	InstID string
}

type okxGetFunc func(path string, headers map[string]string, proxy string) ([]byte, error)

var fallbackSymbolTargets = []symbolTarget{
	{Symbol: "BTC/USDT", InstID: "BTC-USDT"},
	{Symbol: "ETH/USDT", InstID: "ETH-USDT"},
	{Symbol: "SOL/USDT", InstID: "SOL-USDT"},
}

type Kline struct {
	wg          sync.WaitGroup
	c           OkxConfig
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
	ch          cache.Cache
	marketRpc   mclient.Market
}

func normalizeSymbolTargets(coins []*markettypes.ExchangeCoin) []symbolTarget {
	targets := make([]symbolTarget, 0, len(coins))
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
		targets = append(targets, symbolTarget{
			Symbol: symbol,
			InstID: strings.ReplaceAll(symbol, "/", "-"),
		})
	}

	return targets
}

func getWithProxyFallback(request okxGetFunc, path string, headers map[string]string, proxy string) ([]byte, error) {
	if proxy == "" {
		return request(path, headers, "")
	}

	body, err := request(path, headers, proxy)
	if err == nil {
		return body, nil
	}

	log.Printf("okx request via proxy failed, retrying direct connection: %v", err)
	return request(path, headers, "")
}

func (k *Kline) loadTargets() []symbolTarget {
	if k.marketRpc == nil {
		return append([]symbolTarget{}, fallbackSymbolTargets...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := k.marketRpc.FindExchangeCoinVisible(ctx, &mclient.MarketReq{})
	if err != nil {
		log.Printf("load visible exchange coins failed, fallback to defaults: %v", err)
		return append([]symbolTarget{}, fallbackSymbolTargets...)
	}

	targets := normalizeSymbolTargets(resp.List)
	if len(targets) == 0 {
		return append([]symbolTarget{}, fallbackSymbolTargets...)
	}

	return targets
}

func (k *Kline) Do(period string) {
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

func (k *Kline) getKlineData(target symbolTarget, period string) {
	api := k.c.Host + "/api/v5/market/candles?instId=" + target.InstID + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+target.InstID+"&bar="+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass

	resp, err := getWithProxyFallback(tools.GetWithHeader, api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}

	var result OkxResult
	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}

	log.Println("==================执行存储mongo====================")
	if result.Code == "0" {
		k.klineDomain.SaveBatch(result.Data, target.Symbol, period)
		if period == "1m" && len(result.Data) > 0 {
			data := result.Data[0]
			k.queueDomain.Send1mKline(data, target.Symbol)
			key := strings.ReplaceAll(target.InstID, "-", "::")
			k.ch.Set(key+"::RATE", data[4])
		}
	}

	k.wg.Done()
	log.Println("==================End====================")
}

func NewKline(c OkxConfig, marketRpc mclient.Market, mongoClient *database.MongoClient, kafkaCli *database.KafkaClient, cache2 cache.Cache) *Kline {
	return &Kline{
		c:           c,
		klineDomain: domain.NewKlineDomain(mongoClient),
		queueDomain: domain.NewQueueDomain(kafkaCli),
		ch:          cache2,
		marketRpc:   marketRpc,
	}
}

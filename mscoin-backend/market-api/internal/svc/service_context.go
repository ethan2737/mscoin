package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"market-api/internal/config"
	"market-api/internal/database"
	"market-api/internal/processor"
	"market-api/internal/repo"
	"market-api/internal/ws"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
	MarketRpc       mclient.Market
	Processor       processor.Processor
	SnapshotStore   repo.TradeSnapshotStore
}

func NewServiceContext(c config.Config, server *ws.WebsocketServer) *ServiceContext {
	//初始化processor
	kafaCli := database.NewKafkaClient(c.Kafka)
	market := mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc))
	defaultProcessor := processor.NewDefaultProcessor(kafaCli)
	defaultProcessor.Init(market)
	defaultProcessor.AddHandler(processor.NewWebsocketHandler(server))
	var snapshotStore repo.TradeSnapshotStore
	if c.ExchangeMysql.DataSource != "" {
		snapshotStore = repo.NewExchangeSnapshotStore(database.ConnMysql(c.ExchangeMysql))
	}
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
		MarketRpc:       market,
		Processor:       defaultProcessor,
		SnapshotStore:   snapshotStore,
	}
}

package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/ucclient"
	"jobcenter/internal/config"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
)

type ServiceContext struct {
	Config            config.Config
	MongoClient       *database.MongoClient
	KafkaClient       *database.KafkaClient
	Cache             cache.Cache
	AssetRpc          ucclient.Asset
	MarketRpc         mclient.Market
	SwapPositionModel model.ContractPositionModel
	SwapWalletModel   model.ContractWalletModel
	SwapCoinModel     model.ContractCoinModel
	BitCoinAddress    string
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite()
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {})

	// 初始化 swap model (MySQL)
	mysqlConn := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)
	swapPositionModel := model.NewContractPositionModel(mysqlConn)
	swapWalletModel := model.NewContractWalletModel(mysqlConn)
	swapCoinModel := model.NewContractCoinModel(mysqlConn)

	return &ServiceContext{
		Config:            c,
		MongoClient:       database.ConnectMongo(c.Mongo),
		KafkaClient:       client,
		Cache:             redisCache,
		AssetRpc:          ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		MarketRpc:         mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
		SwapPositionModel: swapPositionModel,
		SwapWalletModel:   swapWalletModel,
		SwapCoinModel:     swapCoinModel,
		BitCoinAddress:    c.Bitcoin.Address,
	}
}

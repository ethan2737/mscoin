package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"jobcenter/internal/database"
	"jobcenter/internal/logic"
)

type Config struct {
	Okx             logic.OkxConfig
	Mongo           database.MongoConfig
	Kafka           database.KafkaConfig
	CacheRedis      cache.CacheConf
	UCenterRpc      zrpc.RpcClientConf
	MarketRpc       zrpc.RpcClientConf
	Bitcoin         logic.BitCoinConfig
	SwapKline       logic.SwapKlineConfig
	SwapLiquidation logic.SwapLiquidationConfig
	Mysql           MysqlConfig
}

type MysqlConfig struct {
	DataSource string
}

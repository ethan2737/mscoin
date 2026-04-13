package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql       MysqlConfig
	CacheRedis  redis.RedisConf
	Mongo       MongoConfig
	MarketRpc   zrpc.RpcClientConf
	UcenterRpc  zrpc.RpcClientConf
}

type MysqlConfig struct {
	DataSource string
}

type MongoConfig struct {
	Url      string
	Username string
	Password string
	DataBase string
}

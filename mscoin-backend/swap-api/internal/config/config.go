package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql       MysqlConfig
	CacheRedis  redis.RedisConf
	Mongo       MongoConfig
	SwapRpc     zrpc.RpcClientConf
	MarketRpc   zrpc.RpcClientConf
	UcenterRpc  zrpc.RpcClientConf
	JWT         AuthConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
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

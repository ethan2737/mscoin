package svc

import (
	"context"
	"grpc-common/market/mclient"
	"log"
	"swap-api/internal/config"
	"swap-api/internal/dao"
	"swap-api/internal/marketdata"
	"swap-api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config config.Config

	// MySQL
	ContractCoinModel        model.ContractCoinModel
	ContractOrderModel       model.ContractOrderModel
	ContractPositionModel    model.ContractPositionModel
	ContractWalletModel      model.ContractWalletModel
	ContractTransactionModel model.ContractTransactionModel

	// DAO
	ContractCoinDao        dao.ContractCoinDao
	ContractOrderDao       dao.ContractOrderDao
	ContractPositionDao    dao.ContractPositionDao
	ContractWalletDao      dao.ContractWalletDao
	ContractTransactionDao dao.ContractTransactionDao

	// Redis
	Redis *redis.Redis

	// MongoDB
	MongoClient *mongo.Client
	MongoDb     *mongo.Database

	// RPC Clients
	SwapClient       zrpc.Client
	MarketRpc        mclient.Market
	MarketDataClient marketdata.Client
	UcenterClient    zrpc.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 MySQL
	mysqlConn := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)

	// 初始化模型
	contractCoinModel := model.NewContractCoinModel(mysqlConn)
	contractOrderModel := model.NewContractOrderModel(mysqlConn)
	contractPositionModel := model.NewContractPositionModel(mysqlConn)
	contractWalletModel := model.NewContractWalletModel(mysqlConn)
	contractTransactionModel := model.NewContractTransactionModel(mysqlConn)

	// 初始化 DAO
	contractCoinDao := dao.NewContractCoinDao(contractCoinModel)
	contractOrderDao := dao.NewContractOrderDao(contractOrderModel)
	contractPositionDao := dao.NewContractPositionDao(contractPositionModel)
	contractWalletDao := dao.NewContractWalletDao(contractWalletModel)
	contractTransactionDao := dao.NewContractTransactionDao(contractTransactionModel)

	// 初始化 Redis
	r := redis.MustNewRedis(c.CacheRedis)

	// 初始化 MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.Mongo.Url))
	if err != nil {
		log.Fatalf("failed to connect mongo: %v", err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("failed to ping mongo: %v", err)
	}

	// 初始化 RPC 客户端
	var swapClient, ucenterClient zrpc.Client
	var marketRpc mclient.Market
	if !c.SwapRpc.NonBlock {
		swapClient = zrpc.MustNewClient(c.SwapRpc)
	}
	marketRpc = mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc))
	if !c.UcenterRpc.NonBlock {
		ucenterClient = zrpc.MustNewClient(c.UcenterRpc)
	}

	return &ServiceContext{
		Config:                   c,
		ContractCoinModel:        contractCoinModel,
		ContractOrderModel:       contractOrderModel,
		ContractPositionModel:    contractPositionModel,
		ContractWalletModel:      contractWalletModel,
		ContractTransactionModel: contractTransactionModel,
		ContractCoinDao:          contractCoinDao,
		ContractOrderDao:         contractOrderDao,
		ContractPositionDao:      contractPositionDao,
		ContractWalletDao:        contractWalletDao,
		ContractTransactionDao:   contractTransactionDao,
		Redis:                    r,
		MongoClient:              client,
		MongoDb:                  client.Database(c.Mongo.DataBase),
		SwapClient:               swapClient,
		MarketRpc:                marketRpc,
		MarketDataClient:         marketdata.New(c.MarketApiBaseURL),
		UcenterClient:            ucenterClient,
	}
}

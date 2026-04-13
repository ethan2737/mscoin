package svc

import (
	"context"
	"log"
	"swap-api/internal/config"
	"swap-api/internal/dao"
	"swap-api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	SwapConn    *grpc.ClientConn
	MarketConn  *grpc.ClientConn
	UcenterConn *grpc.ClientConn
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

	// 初始化 RPC 连接 (延迟连接)
	var swapConn, marketConn, ucenterConn *grpc.ClientConn
	if !c.SwapRpc.NonBlock {
		swapConn = createGrpcConn(c.SwapRpc)
	}
	if !c.MarketRpc.NonBlock {
		marketConn = createGrpcConn(c.MarketRpc)
	}
	if !c.UcenterRpc.NonBlock {
		ucenterConn = createGrpcConn(c.UcenterRpc)
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
		SwapConn:                 swapConn,
		MarketConn:               marketConn,
		UcenterConn:              ucenterConn,
	}
}

func createGrpcConn(cfg zrpc.RpcClientConf) *grpc.ClientConn {
	conn, err := grpc.Dial(
		cfg.Endpoints[0],
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
	if err != nil {
		log.Printf("failed to dial grpc: %v", err)
		return nil
	}
	return conn
}

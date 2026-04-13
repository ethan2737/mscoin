package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"swap/internal/config"
	"swap/internal/dao"
	"swap/internal/model"
)

type ServiceContext struct {
	Config config.Config

	// Models
	ContractCoinModel       model.ContractCoinModel
	ContractOrderModel      model.ContractOrderModel
	ContractPositionModel   model.ContractPositionModel
	ContractWalletModel     model.ContractWalletModel
	ContractTransactionModel model.ContractTransactionModel

	// DAOs
	ContractCoinDao       dao.ContractCoinDao
	ContractOrderDao      dao.ContractOrderDao
	ContractPositionDao   dao.ContractPositionDao
	ContractWalletDao     dao.ContractWalletDao
	ContractTransactionDao dao.ContractTransactionDao

	// Connections
	MongoClient *mongo.Client
	Redis       *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewSqlConn(c.Mysql.DataSource, "")
	mongoClient, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Mongo.Url))
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.CacheRedis.Host,
	})

	svc := &ServiceContext{
		Config: c,
	}

	// Initialize Models
	svc.ContractCoinModel = model.NewContractCoinModel(sqlConn)
	svc.ContractOrderModel = model.NewContractOrderModel(sqlConn)
	svc.ContractPositionModel = model.NewContractPositionModel(sqlConn)
	svc.ContractWalletModel = model.NewContractWalletModel(sqlConn)
	svc.ContractTransactionModel = model.NewContractTransactionModel(sqlConn)

	// Initialize DAOs
	svc.ContractCoinDao = dao.NewContractCoinDao(svc.ContractCoinModel)
	svc.ContractOrderDao = dao.NewContractOrderDao(svc.ContractOrderModel)
	svc.ContractPositionDao = dao.NewContractPositionDao(svc.ContractPositionModel)
	svc.ContractWalletDao = dao.NewContractWalletDao(svc.ContractWalletModel)
	svc.ContractTransactionDao = dao.NewContractTransactionDao(svc.ContractTransactionModel)

	svc.Redis = redisClient
	svc.MongoClient = mongoClient

	return svc
}

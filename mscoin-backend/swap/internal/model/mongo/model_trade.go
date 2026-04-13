package model

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trade struct {
	Id        string    `bson:"_id"`
	Symbol    string    `bson:"symbol"`
	Price     float64   `bson:"price"`
	Amount    float64   `bson:"amount"`
	Direction string    `bson:"direction"`
	CreatedAt time.Time `bson:"created_at"`
}

type TradeModel interface {
	Insert(ctx context.Context, trade *Trade) error
	FindLatest(ctx context.Context, symbol string, limit int64) ([]*Trade, error)
}

type defaultTradeModel struct {
	collection *mongo.Collection
}

func NewTradeModel(db *mongo.Database) TradeModel {
	return &defaultTradeModel{
		collection: db.Collection("swap_trades"),
	}
}

func (m *defaultTradeModel) Insert(ctx context.Context, trade *Trade) error {
	_, err := m.collection.InsertOne(ctx, trade)
	return err
}

func (m *defaultTradeModel) FindLatest(ctx context.Context, symbol string, limit int64) ([]*Trade, error) {
	filter := bson.M{"symbol": symbol}
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(limit)
	cursor, err := m.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trades []*Trade
	if err := cursor.All(ctx, &trades); err != nil {
		return nil, err
	}
	return trades, nil
}

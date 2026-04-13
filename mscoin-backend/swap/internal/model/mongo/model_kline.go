package model

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Kline struct {
	Id        string    `bson:"_id"`
	Symbol    string    `bson:"symbol"`
	Period    string    `bson:"period"`
	OpenTime  time.Time `bson:"open_time"`
	CloseTime time.Time `bson:"close_time"`
	Open      float64   `bson:"open"`
	High      float64   `bson:"high"`
	Low       float64   `bson:"low"`
	Close     float64   `bson:"close"`
	Volume    float64   `bson:"volume"`
	Amount    float64   `bson:"amount"`
	CreatedAt time.Time `bson:"created_at"`
}

type KlineModel interface {
	Insert(ctx context.Context, kline *Kline) error
	FindLatest(ctx context.Context, symbol string, period string) (*Kline, error)
	FindByRange(ctx context.Context, symbol string, period string, start, end time.Time) ([]*Kline, error)
}

type defaultKlineModel struct {
	collection *mongo.Collection
}

func NewKlineModel(db *mongo.Database) KlineModel {
	return &defaultKlineModel{
		collection: db.Collection("swap_klines"),
	}
}

func (m *defaultKlineModel) Insert(ctx context.Context, kline *Kline) error {
	_, err := m.collection.InsertOne(ctx, kline)
	return err
}

func (m *defaultKlineModel) FindLatest(ctx context.Context, symbol string, period string) (*Kline, error) {
	filter := bson.M{"symbol": symbol, "period": period}
	opts := options.FindOne().SetSort(bson.D{{"open_time", -1}})
	var kline Kline
	err := m.collection.FindOne(ctx, filter, opts).Decode(&kline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &kline, nil
}

func (m *defaultKlineModel) FindByRange(ctx context.Context, symbol string, period string, start, end time.Time) ([]*Kline, error) {
	filter := bson.M{
		"symbol":     symbol,
		"period":     period,
		"open_time":  bson.M{"$gte": start, "$lte": end},
	}
	opts := options.Find().SetSort(bson.D{{"open_time", 1}})
	cursor, err := m.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var klines []*Kline
	if err := cursor.All(ctx, &klines); err != nil {
		return nil, err
	}
	return klines, nil
}

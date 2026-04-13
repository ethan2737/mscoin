package model

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Symbol struct {
	Id         string    `bson:"_id"`
	Symbol     string    `bson:"symbol"`
	Price      float64   `bson:"price"`
	Change24h  float64   `bson:"change_24h"`
	High24h    float64   `bson:"high_24h"`
	Low24h     float64   `bson:"low_24h"`
	Volume24h  float64   `bson:"volume_24h"`
	Amount24h  float64   `bson:"amount_24h"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

type SymbolModel interface {
	Upsert(ctx context.Context, symbol *Symbol) error
	FindOne(ctx context.Context, symbol string) (*Symbol, error)
	FindAll(ctx context.Context) ([]*Symbol, error)
}

type defaultSymbolModel struct {
	collection *mongo.Collection
}

func NewSymbolModel(db *mongo.Database) SymbolModel {
	return &defaultSymbolModel{
		collection: db.Collection("swap_symbols"),
	}
}

func (m *defaultSymbolModel) Upsert(ctx context.Context, symbol *Symbol) error {
	filter := bson.M{"symbol": symbol.Symbol}
	update := bson.M{
		"$set": bson.M{
			"price":       symbol.Price,
			"change_24h":  symbol.Change24h,
			"high_24h":    symbol.High24h,
			"low_24h":     symbol.Low24h,
			"volume_24h":  symbol.Volume24h,
			"amount_24h":  symbol.Amount24h,
			"updated_at":  symbol.UpdatedAt,
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := m.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (m *defaultSymbolModel) FindOne(ctx context.Context, symbol string) (*Symbol, error) {
	filter := bson.M{"symbol": symbol}
	var s Symbol
	err := m.collection.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (m *defaultSymbolModel) FindAll(ctx context.Context) ([]*Symbol, error) {
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var symbols []*Symbol
	if err := cursor.All(ctx, &symbols); err != nil {
		return nil, err
	}
	return symbols, nil
}

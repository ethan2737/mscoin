package model

import (
	"grpc-common/market/types/market"
	"mscoin-common/op"
)

type Kline struct {
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    //成交笔数
	Volume       float64 `bson:"volume,omitempty"`   //成交量
	Turnover     float64 `bson:"turnover,omitempty"` //成交额
}

// SwapKline 永续合约 K 线
type SwapKline struct {
	Symbol       string  `bson:"symbol,omitempty"`
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    // 成交笔数
	Volume       float64 `bson:"volume,omitempty"`   // 成交量
	Turnover     float64 `bson:"turnover,omitempty"` // 成交额
	TimeStr      string  `bson:"timeStr,omitempty"`
}

func (*Kline) Table(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

func (k *Kline) ToCoinThumb(symbol string, end *Kline) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice
	ct.Zone = 0
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	ct.DateTime = k.Time
	if end == nil || end.ClosePrice <= 0 {
		return ct
	}
	ct.LastDayClose = end.ClosePrice
	ct.Change = k.ClosePrice - end.ClosePrice
	ct.Chg = op.DivN(ct.Change, end.ClosePrice, 8)
	return ct
}

func DefaultCoinThumb(symbol string) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Trend = []float64{}
	return ct
}

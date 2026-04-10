package model

type TradePlateResult struct {
	Direction    string            `json:"direction"`
	MaxAmount    float64           `json:"maxAmount"`
	MinAmount    float64           `json:"minAmount"`
	HighestPrice float64           `json:"highestPrice"`
	LowestPrice  float64           `json:"lowestPrice"`
	Symbol       string            `json:"symbol"`
	Items        []*TradePlateItem `json:"items"`
}
type TradePlateItem struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type TradePlateGroup struct {
	Ask *TradePlateResult `json:"ask"`
	Bid *TradePlateResult `json:"bid"`
}

type LatestTrade struct {
	Symbol    string  `json:"symbol,omitempty"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Time      int64   `json:"time"`
	Direction string  `json:"direction"`
}

type CompleteOrder struct {
	Symbol        string  `json:"symbol"`
	Amount        float64 `json:"amount"`
	TradedAmount  float64 `json:"tradedAmount"`
	Price         float64 `json:"price"`
	CompletedTime int64   `json:"completedTime"`
	Time          int64   `json:"time"`
	Direction     int     `json:"direction"`
}

func NewTradePlateResult(direction string, symbol string) *TradePlateResult {
	return &TradePlateResult{
		Direction: direction,
		Symbol:    symbol,
		Items:     make([]*TradePlateItem, 0),
	}
}

func LatestTradeFromCompleteOrder(order *CompleteOrder) *LatestTrade {
	if order == nil {
		return nil
	}
	amount := order.TradedAmount
	if amount <= 0 {
		amount = order.Amount
	}
	direction := "BUY"
	if order.Direction == 1 {
		direction = "SELL"
	}
	time := order.CompletedTime
	if time <= 0 {
		time = order.Time
	}
	return &LatestTrade{
		Symbol:    order.Symbol,
		Amount:    amount,
		Price:     order.Price,
		Time:      time,
		Direction: direction,
	}
}

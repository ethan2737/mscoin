package repo

import (
	"context"
	"math"
	"sort"
	"strconv"

	"gorm.io/gorm"
	"market-api/internal/model"
	"mscoin-common/msdb"
)

const (
	exchangeOrderTrading   = 0
	exchangeOrderCompleted = 1
	exchangeDirectionBuy   = 0
	exchangeDirectionSell  = 1
	exchangeTypeLimitPrice = 1
)

type TradeSnapshotStore interface {
	LoadTradePlate(ctx context.Context, symbol string, size int) (*model.TradePlateGroup, error)
	LoadLatestTrade(ctx context.Context, symbol string, size int) ([]*model.LatestTrade, error)
}

type ExchangeSnapshotStore struct {
	conn *gorm.DB
}

type exchangeOrder struct {
	Id            int64   `gorm:"column:id"`
	OrderId       string  `gorm:"column:order_id"`
	Amount        float64 `gorm:"column:amount"`
	TradedAmount  float64 `gorm:"column:traded_amount"`
	Price         float64 `gorm:"column:price"`
	Symbol        string  `gorm:"column:symbol"`
	Time          int64   `gorm:"column:time"`
	CompletedTime int64   `gorm:"column:completed_time"`
	Direction     int     `gorm:"column:direction"`
	Status        int     `gorm:"column:status"`
	Type          int     `gorm:"column:type"`
}

func (*exchangeOrder) TableName() string {
	return "exchange_order"
}

func NewExchangeSnapshotStore(db *msdb.MsDB) *ExchangeSnapshotStore {
	if db == nil || db.Conn == nil {
		return nil
	}
	return &ExchangeSnapshotStore{conn: db.Conn}
}

func (s *ExchangeSnapshotStore) LoadTradePlate(ctx context.Context, symbol string, size int) (*model.TradePlateGroup, error) {
	askOrders, err := s.loadActiveOrders(ctx, symbol, exchangeDirectionSell)
	if err != nil {
		return nil, err
	}
	bidOrders, err := s.loadActiveOrders(ctx, symbol, exchangeDirectionBuy)
	if err != nil {
		return nil, err
	}
	return &model.TradePlateGroup{
		Ask: buildTradePlateResult("SELL", symbol, askOrders, size),
		Bid: buildTradePlateResult("BUY", symbol, bidOrders, size),
	}, nil
}

func (s *ExchangeSnapshotStore) LoadLatestTrade(ctx context.Context, symbol string, size int) ([]*model.LatestTrade, error) {
	if size <= 0 {
		size = 20
	}
	var orders []*exchangeOrder
	err := s.conn.WithContext(ctx).
		Model(&exchangeOrder{}).
		Where("symbol = ? AND status = ?", symbol, exchangeOrderCompleted).
		Order("completed_time DESC, id DESC").
		Limit(size).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	list := make([]*model.LatestTrade, 0, len(orders))
	for _, order := range orders {
		list = append(list, model.LatestTradeFromCompleteOrder(&model.CompleteOrder{
			Symbol:        order.Symbol,
			Amount:        order.Amount,
			TradedAmount:  order.TradedAmount,
			Price:         order.Price,
			CompletedTime: order.CompletedTime,
			Time:          order.Time,
			Direction:     order.Direction,
		}))
	}
	return list, nil
}

func (s *ExchangeSnapshotStore) loadActiveOrders(ctx context.Context, symbol string, direction int) ([]*exchangeOrder, error) {
	var orders []*exchangeOrder
	orderClause := "price ASC, time ASC, id ASC"
	if direction == exchangeDirectionBuy {
		orderClause = "price DESC, time ASC, id ASC"
	}
	err := s.conn.WithContext(ctx).
		Model(&exchangeOrder{}).
		Where("symbol = ? AND status = ? AND direction = ? AND type = ?", symbol, exchangeOrderTrading, direction, exchangeTypeLimitPrice).
		Order(orderClause).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func buildTradePlateResult(direction string, symbol string, orders []*exchangeOrder, size int) *model.TradePlateResult {
	result := model.NewTradePlateResult(direction, symbol)
	if len(orders) == 0 {
		return result
	}

	type aggregatedItem struct {
		price  float64
		amount float64
	}

	itemsByPrice := make(map[string]*aggregatedItem, len(orders))
	priceKeys := make([]float64, 0, len(orders))

	for _, order := range orders {
		remaining := math.Max(order.Amount-order.TradedAmount, 0)
		if remaining <= 0 {
			continue
		}
		key := formatTradePriceKey(order.Price)
		item, ok := itemsByPrice[key]
		if !ok {
			item = &aggregatedItem{price: order.Price}
			itemsByPrice[key] = item
			priceKeys = append(priceKeys, order.Price)
		}
		item.amount += remaining
	}

	if len(priceKeys) == 0 {
		return result
	}

	sort.Float64s(priceKeys)
	if direction == "BUY" {
		sort.Slice(priceKeys, func(i, j int) bool {
			return priceKeys[i] > priceKeys[j]
		})
	}

	if size > 0 && len(priceKeys) > size {
		priceKeys = priceKeys[:size]
	}

	result.Items = make([]*model.TradePlateItem, 0, len(priceKeys))
	result.MinAmount = math.MaxFloat64
	for _, price := range priceKeys {
		item := itemsByPrice[formatTradePriceKey(price)]
		if item == nil || item.amount <= 0 {
			continue
		}
		result.Items = append(result.Items, &model.TradePlateItem{
			Price:  item.price,
			Amount: item.amount,
		})
		if item.amount > result.MaxAmount {
			result.MaxAmount = item.amount
		}
		if item.amount < result.MinAmount {
			result.MinAmount = item.amount
		}
		if result.HighestPrice == 0 || item.price > result.HighestPrice {
			result.HighestPrice = item.price
		}
		if result.LowestPrice == 0 || item.price < result.LowestPrice {
			result.LowestPrice = item.price
		}
	}

	if len(result.Items) == 0 {
		result.MinAmount = 0
		return result
	}
	if result.MinAmount == math.MaxFloat64 {
		result.MinAmount = 0
	}
	return result
}

func formatTradePriceKey(price float64) string {
	return strconv.FormatFloat(price, 'f', 8, 64)
}

var _ TradeSnapshotStore = (*ExchangeSnapshotStore)(nil)

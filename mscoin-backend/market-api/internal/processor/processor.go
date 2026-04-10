package processor

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"market-api/internal/database"
	"market-api/internal/model"
	"sync"
)

const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"
const TradePlateTopic = "exchange_order_trade_plate"
const CompleteTradeTopic = "exchange_order_complete"
const TradePlate = "tradePlate"

type MarketHandler interface {
	HandleTrade(symbol string, trades []*model.LatestTrade)
	HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb)
	HandleTradePlate(symbol string, tp *model.TradePlateResult)
}

type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}

type Processor interface {
	GetThumb() any
	GetTradePlate(symbol string, size int) *model.TradePlateGroup
	GetLatestTrade(symbol string, size int) []*model.LatestTrade
	Process(data ProcessData)
	AddHandler(h MarketHandler)
}

type DefaultProcessor struct {
	kafkaCli *database.KafkaClient
	handlers []MarketHandler
	thumbMap map[string]*market.CoinThumb
	plateMap map[string]*model.TradePlateGroup
	tradeMap map[string][]*model.LatestTrade
	mux      sync.RWMutex
}

func NewDefaultProcessor(kafkaCli *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		kafkaCli: kafkaCli,
		handlers: make([]MarketHandler, 0),
		thumbMap: make(map[string]*market.CoinThumb),
		plateMap: make(map[string]*model.TradePlateGroup),
		tradeMap: make(map[string][]*model.LatestTrade),
	}
}

func (d *DefaultProcessor) Process(data ProcessData) {
	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		if err := json.Unmarshal(data.Data, kline); err != nil {
			logx.Error(err)
			return
		}
		d.mux.Lock()
		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumbMap)
		}
		d.mux.Unlock()
	} else if data.Type == TradePlate {
		symbol := string(data.Key)
		tp := &model.TradePlateResult{}
		if err := json.Unmarshal(data.Data, tp); err != nil {
			logx.Error(err)
			return
		}
		d.mux.Lock()
		d.updateTradePlateCache(symbol, tp)
		d.mux.Unlock()
		for _, v := range d.handlers {
			v.HandleTradePlate(symbol, tp)
		}
	} else if data.Type == TRADE {
		symbol := string(data.Key)
		order := &model.CompleteOrder{}
		if err := json.Unmarshal(data.Data, order); err != nil {
			logx.Error(err)
			return
		}
		trade := model.LatestTradeFromCompleteOrder(order)
		if trade == nil {
			return
		}
		d.mux.Lock()
		d.appendLatestTrade(symbol, trade)
		d.mux.Unlock()
		for _, v := range d.handlers {
			v.HandleTrade(symbol, []*model.LatestTrade{trade})
		}
	}
}

func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	//发送到websocket的服务
	d.handlers = append(d.handlers, h)
}

func (p *DefaultProcessor) Init(marketRpc mclient.Market) {
	p.startReadFromKafka(KLINE1M, KLINE)
	p.startReadTradePlate(TradePlateTopic)
	p.startReadTrade(CompleteTradeTopic)
	p.initThumbMap(marketRpc)
}
func (d *DefaultProcessor) GetThumb() any {
	d.mux.RLock()
	defer d.mux.RUnlock()
	cs := make([]*market.CoinThumb, len(d.thumbMap))
	i := 0
	for _, v := range d.thumbMap {
		cs[i] = v
		i++
	}
	return cs
}

func (d *DefaultProcessor) GetTradePlate(symbol string, size int) *model.TradePlateGroup {
	d.mux.RLock()
	defer d.mux.RUnlock()
	group, ok := d.plateMap[symbol]
	if !ok {
		return &model.TradePlateGroup{
			Ask: model.NewTradePlateResult("SELL", symbol),
			Bid: model.NewTradePlateResult("BUY", symbol),
		}
	}
	return &model.TradePlateGroup{
		Ask: cloneTradePlateResult(group.Ask, size),
		Bid: cloneTradePlateResult(group.Bid, size),
	}
}

func (d *DefaultProcessor) GetLatestTrade(symbol string, size int) []*model.LatestTrade {
	d.mux.RLock()
	defer d.mux.RUnlock()
	list := d.tradeMap[symbol]
	if len(list) == 0 {
		return make([]*model.LatestTrade, 0)
	}
	if size <= 0 || size > len(list) {
		size = len(list)
	}
	result := make([]*model.LatestTrade, 0, size)
	for i := 0; i < size; i++ {
		item := *list[i]
		result = append(result, &item)
	}
	return result
}

func (p *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	//一定要先start 后read
	p.kafkaCli.StartRead(topic)
	go p.dealQueueData(p.kafkaCli, tp)
}

func (p *DefaultProcessor) dealQueueData(cli *database.KafkaClient, tp string) {
	//这就是队列的数据
	for {
		msg := cli.Read()
		data := ProcessData{
			Type: tp,
			Key:  msg.Key,
			Data: msg.Data,
		}
		p.Process(data)
	}

}

func (d *DefaultProcessor) initThumbMap(marketRpc mclient.Market) {
	symbolThumbRes, err := marketRpc.FindSymbolThumbTrend(context.Background(),
		&market.MarketReq{})
	if err != nil {
		logx.Info(err)
	} else {
		coinThumbs := symbolThumbRes.List
		for _, v := range coinThumbs {
			d.thumbMap[v.Symbol] = v
		}
	}
}

func (p *DefaultProcessor) startReadTradePlate(topic string) {
	cli := p.kafkaCli.StartReadNew(topic)
	go p.dealQueueData(cli, TradePlate)
}

func (p *DefaultProcessor) startReadTrade(topic string) {
	cli := p.kafkaCli.StartReadNew(topic)
	go p.dealQueueData(cli, TRADE)
}

func (d *DefaultProcessor) updateTradePlateCache(symbol string, tp *model.TradePlateResult) {
	group, ok := d.plateMap[symbol]
	if !ok {
		group = &model.TradePlateGroup{
			Ask: model.NewTradePlateResult("SELL", symbol),
			Bid: model.NewTradePlateResult("BUY", symbol),
		}
		d.plateMap[symbol] = group
	}
	if tp.Direction == "SELL" {
		group.Ask = cloneTradePlateResult(tp, 0)
		return
	}
	group.Bid = cloneTradePlateResult(tp, 0)
}

func (d *DefaultProcessor) appendLatestTrade(symbol string, trade *model.LatestTrade) {
	if trade == nil {
		return
	}
	list := append([]*model.LatestTrade{trade}, d.tradeMap[symbol]...)
	if len(list) > 100 {
		list = list[:100]
	}
	d.tradeMap[symbol] = list
}

func cloneTradePlateResult(src *model.TradePlateResult, size int) *model.TradePlateResult {
	if src == nil {
		return model.NewTradePlateResult("", "")
	}
	if size <= 0 || size > len(src.Items) {
		size = len(src.Items)
	}
	items := make([]*model.TradePlateItem, 0, size)
	for i := 0; i < size; i++ {
		item := *src.Items[i]
		items = append(items, &item)
	}
	return &model.TradePlateResult{
		Direction:    src.Direction,
		MaxAmount:    src.MaxAmount,
		MinAmount:    src.MinAmount,
		HighestPrice: src.HighestPrice,
		LowestPrice:  src.LowestPrice,
		Symbol:       src.Symbol,
		Items:        items,
	}
}

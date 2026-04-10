package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	marketmodel "market-api/internal/model"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *MarketLogic) loadThumbs(req *types.MarketReq) ([]*market.CoinThumb, error) {
	ctx, cancelFunc := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancelFunc()
	symbolThumbRes, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(ctx, &market.MarketReq{
		Ip: req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return symbolThumbRes.List, nil
}

func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	thumbs, err := l.loadThumbs(req)
	if err != nil {
		return nil, err
	}
	if err := copier.Copy(&list, thumbs); err != nil {
		return nil, err
	}
	return
}

func (l *MarketLogic) SymbolThumb(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	thumbs, err := l.loadThumbs(req)
	if err != nil {
		return nil, err
	}
	if err := copier.Copy(&list, thumbs); err != nil {
		return nil, err
	}
	return
}

func (l *MarketLogic) SymbolInfo(req types.MarketReq) (resp *types.ExchangeCoinResp, err error) {
	ctx, cancelFunc := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancelFunc()
	esRes, err := l.svcCtx.MarketRpc.FindSymbolInfo(ctx,
		&market.MarketReq{
			Ip:     req.Ip,
			Symbol: req.Symbol,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ExchangeCoinResp{}
	if err := copier.Copy(resp, esRes); err != nil {
		return nil, err
	}
	return
}

func (l *MarketLogic) CoinInfo(req *types.MarketReq) (*types.Coin, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	coin, err := l.svcCtx.MarketRpc.FindCoinInfo(ctx, &market.MarketReq{
		Unit: req.Unit,
	})
	if err != nil {
		return nil, err
	}
	ec := &types.Coin{}
	if err := copier.Copy(&ec, coin); err != nil {
		return nil, errors.New("数据格式有误")
	}
	return ec, nil
}

func (l *MarketLogic) History(req *types.MarketReq) (*types.HistoryKline, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	historyKline, err := l.svcCtx.MarketRpc.HistoryKline(ctx, &market.MarketReq{
		Symbol:     req.Symbol,
		From:       req.From,
		To:         req.To,
		Resolution: req.Resolution,
	})
	if err != nil {
		return nil, err
	}
	histories := historyKline.List
	var list = make([][]any, len(histories))
	for i, v := range histories {
		content := make([]any, 6)
		content[0] = v.Time
		content[1] = v.Open
		content[2] = v.High
		content[3] = v.Low
		content[4] = v.Close
		content[5] = v.Volume
		list[i] = content
	}
	return &types.HistoryKline{
		List: list,
	}, nil
}

func (l *MarketLogic) ExchangePlateMini(req *types.MarketReq) (*types.TradePlateResp, error) {
	return l.loadTradePlate(req.Symbol, 10)
}

func (l *MarketLogic) ExchangePlateFull(req *types.MarketReq) (*types.TradePlateResp, error) {
	return l.loadTradePlate(req.Symbol, 0)
}

func (l *MarketLogic) LatestTrade(req *types.MarketReq) ([]*types.LatestTradeResp, error) {
	size := req.Size
	if size <= 0 {
		size = 20
	}
	trades := l.svcCtx.Processor.GetLatestTrade(req.Symbol, size)
	if len(trades) == 0 && l.svcCtx.SnapshotStore != nil {
		loadedTrades, err := l.svcCtx.SnapshotStore.LoadLatestTrade(l.ctx, req.Symbol, size)
		if err != nil {
			return nil, err
		}
		trades = loadedTrades
	}
	list := make([]*types.LatestTradeResp, 0, len(trades))
	if err := copier.Copy(&list, trades); err != nil {
		return nil, err
	}
	return list, nil
}

func (l *MarketLogic) loadTradePlate(symbol string, size int) (*types.TradePlateResp, error) {
	plate := l.svcCtx.Processor.GetTradePlate(symbol, size)
	if !hasTradePlateData(plate) && l.svcCtx.SnapshotStore != nil {
		loadedPlate, err := l.svcCtx.SnapshotStore.LoadTradePlate(l.ctx, symbol, size)
		if err != nil {
			return nil, err
		}
		plate = loadedPlate
	}
	resp := &types.TradePlateResp{}
	if err := copier.Copy(resp, plate); err != nil {
		return nil, err
	}
	return resp, nil
}

func hasTradePlateData(plate *marketmodel.TradePlateGroup) bool {
	if plate == nil {
		return false
	}
	if plate.Ask != nil && len(plate.Ask.Items) > 0 {
		return true
	}
	return plate.Bid != nil && len(plate.Bid.Items) > 0
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

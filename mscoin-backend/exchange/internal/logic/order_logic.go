package logic

import (
	"context"
	"errors"
	"exchange/internal/domain"
	"exchange/internal/model"
	"exchange/internal/svc"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"grpc-common/ucenter/types/member"
	"mscoin-common/msdb"
	"mscoin-common/msdb/tran"
)

type ExchangeOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeOrderDomain *domain.ExchangeOrderDomain
	transaction         tran.Transaction
	kafkaDomain         *domain.KafkaDomain
}

func (l *ExchangeOrderLogic) FindOrderHistory(req *order.OrderReq) (*order.OrderRes, error) {
	exchangeOrders, total, err := l.exchangeOrderDomain.FindOrderHistory(
		l.ctx,
		req.Symbol,
		req.Page,
		req.PageSize,
		req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	copier.Copy(&list, exchangeOrders)
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

func (l *ExchangeOrderLogic) FindOrderCurrent(req *order.OrderReq) (*order.OrderRes, error) {
	exchangeOrders, total, err := l.exchangeOrderDomain.FindOrderCurrent(
		l.ctx,
		req.Symbol,
		req.Page,
		req.PageSize,
		req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	copier.Copy(&list, exchangeOrders)
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

func (l *ExchangeOrderLogic) Add(req *order.OrderReq) (*order.AddOrderRes, error) {
	memberRes, err := l.svcCtx.MemberRpc.FindMemberById(l.ctx, &member.MemberReq{
		MemberId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if memberRes.TransactionStatus == 0 {
		return nil, errors.New("姝ょ敤鎴峰凡缁忚绂佹浜ゆ槗")
	}
	if req.Type == model.TypeMap[model.LimitPrice] && req.Price <= 0 {
		return nil, errors.New("闄愪环妯″紡涓嬩环鏍间笉鑳藉皬浜庣瓑浜?")
	}
	if req.Amount <= 0 {
		return nil, errors.New("鏁伴噺涓嶈兘灏忎簬绛変簬0")
	}

	exchangeCoin, err := l.svcCtx.MarketRpc.FindSymbolInfo(l.ctx, &market.MarketReq{
		Symbol: req.Symbol,
	})
	if err != nil {
		return nil, errors.New("nonsupport coin")
	}
	if exchangeCoin.Exchangeable != 1 && exchangeCoin.Enable != 1 {
		return nil, errors.New("coin forbidden")
	}

	baseSymbol := exchangeCoin.GetBaseSymbol()
	coinSymbol := exchangeCoin.GetCoinSymbol()
	cc := baseSymbol
	if req.Direction == model.DirectionMap[model.SELL] {
		cc = coinSymbol
	}

	coin, err := l.svcCtx.MarketRpc.FindCoinInfo(l.ctx, &market.MarketReq{
		Unit: cc,
	})
	if err != nil || coin == nil {
		return nil, errors.New("nonsupport coin")
	}

	if req.Type == model.TypeMap[model.MarketPrice] && req.Direction == model.DirectionMap[model.BUY] {
		if exchangeCoin.GetMinTurnover() > 0 && req.Amount < float64(exchangeCoin.GetMinTurnover()) {
			return nil, errors.New("鎴愪氦棰濊嚦灏戞槸" + fmt.Sprintf("%d", exchangeCoin.GetMinTurnover()))
		}
	} else {
		if exchangeCoin.GetMaxVolume() > 0 && exchangeCoin.GetMaxVolume() < req.Amount {
			return nil, errors.New("鏁伴噺瓒呭嚭" + fmt.Sprintf("%f", exchangeCoin.GetMaxVolume()))
		}
		if exchangeCoin.GetMinVolume() > 0 && exchangeCoin.GetMinVolume() > req.Amount {
			return nil, errors.New("鏁伴噺涓嶈兘浣庝簬" + fmt.Sprintf("%f", exchangeCoin.GetMinVolume()))
		}
	}

	baseWallet, err := l.svcCtx.AssetRpc.FindWalletBySymbol(l.ctx, &asset.AssetReq{
		UserId:   req.UserId,
		CoinName: baseSymbol,
	})
	if err != nil {
		return nil, errors.New("no wallet")
	}
	exCoinWallet, err := l.svcCtx.AssetRpc.FindWalletBySymbol(l.ctx, &asset.AssetReq{
		UserId:   req.UserId,
		CoinName: coinSymbol,
	})
	if err != nil {
		return nil, errors.New("no wallet")
	}
	if baseWallet.IsLock == 1 || exCoinWallet.IsLock == 1 {
		return nil, errors.New("wallet locked")
	}

	if req.Direction == model.DirectionMap[model.SELL] && exchangeCoin.GetMinSellPrice() > 0 {
		if req.Price < exchangeCoin.GetMinSellPrice() || req.Type == model.TypeMap[model.MarketPrice] {
			return nil, errors.New("涓嶈兘浣庝簬鏈€浣庨檺浠?" + fmt.Sprintf("%f", exchangeCoin.GetMinSellPrice()))
		}
	}
	if req.Direction == model.DirectionMap[model.BUY] && exchangeCoin.GetMaxBuyPrice() > 0 {
		if req.Price > exchangeCoin.GetMaxBuyPrice() || req.Type == model.TypeMap[model.MarketPrice] {
			return nil, errors.New("涓嶈兘浣庝簬鏈€楂橀檺浠?" + fmt.Sprintf("%f", exchangeCoin.GetMaxBuyPrice()))
		}
	}
	if req.Type == model.TypeMap[model.MarketPrice] {
		if req.Direction == model.DirectionMap[model.BUY] && exchangeCoin.EnableMarketBuy == 0 {
			return nil, errors.New("涓嶆敮鎸佸競浠疯喘涔?")
		} else if req.Direction == model.DirectionMap[model.SELL] && exchangeCoin.EnableMarketSell == 0 {
			return nil, errors.New("涓嶆敮鎸佸競浠峰嚭鍞?")
		}
	}

	count, err := l.exchangeOrderDomain.FindCurrentTradingCount(l.ctx, req.UserId, req.Symbol, req.Direction)
	if err != nil {
		return nil, err
	}
	if exchangeCoin.GetMaxTradingOrder() > 0 && count >= exchangeCoin.GetMaxTradingOrder() {
		return nil, errors.New("瓒呰繃鏈€澶ф寕鍗曟暟閲?" + fmt.Sprintf("%d", exchangeCoin.GetMaxTradingOrder()))
	}

	exchangeOrder := model.NewOrder()
	exchangeOrder.MemberId = req.UserId
	exchangeOrder.Symbol = req.Symbol
	exchangeOrder.BaseSymbol = baseSymbol
	exchangeOrder.CoinSymbol = coinSymbol
	exchangeOrder.Type = model.TypeMap.Code(req.Type)
	exchangeOrder.Direction = model.DirectionMap.Code(req.Direction)
	if exchangeOrder.Type == model.MarketPrice {
		exchangeOrder.Price = 0
	} else {
		exchangeOrder.Price = req.Price
	}
	exchangeOrder.UseDiscount = "0"
	exchangeOrder.Amount = req.Amount

	var freezeMoney float64
	err = l.transaction.Action(func(conn msdb.DbConn) error {
		money, addErr := l.exchangeOrderDomain.AddOrder(l.ctx, conn, exchangeOrder, exchangeCoin, baseWallet, exCoinWallet)
		if addErr != nil {
			return errors.New("璁㈠崟鎻愪氦澶辫触")
		}
		freezeMoney = money
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = l.kafkaDomain.SendOrderAdd(
		"add-exchange-order",
		req.UserId,
		exchangeOrder.OrderId,
		freezeMoney,
		req.Symbol,
		exchangeOrder.Direction,
		baseSymbol,
		coinSymbol)
	if err != nil {
		return nil, errors.New("鍙戞秷鎭け璐?")
	}
	return &order.AddOrderRes{
		OrderId: exchangeOrder.OrderId,
	}, nil
}

func (l *ExchangeOrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	exchangeOrder, err := l.exchangeOrderDomain.FindOrderByOrderId(l.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	oo := &order.ExchangeOrderOrigin{}
	copier.Copy(oo, exchangeOrder)
	return oo, nil
}

func (l *ExchangeOrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	exchangeOrder, err := l.exchangeOrderDomain.FindOrderByOrderId(l.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if exchangeOrder.Status != model.Trading && exchangeOrder.Status != model.Init {
		return nil, errors.New("当前订单不可撤销")
	}
	if err = l.exchangeOrderDomain.UpdateStatusCancel(l.ctx, req.OrderId); err != nil {
		return nil, err
	}
	if exchangeOrder.Status == model.Trading {
		if err = l.kafkaDomain.SendOrderCancel("exchange_order_cancel", req.OrderId); err != nil {
			return nil, errors.New("撤单消息发送失败")
		}
	}
	return &order.CancelOrderRes{}, nil
}
func NewExchangeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeOrderLogic {
	orderDomain := domain.NewExchangeOrderDomain(svcCtx.Db)
	return &ExchangeOrderLogic{
		ctx:                 ctx,
		svcCtx:              svcCtx,
		Logger:              logx.WithContext(ctx),
		exchangeOrderDomain: orderDomain,
		transaction:         tran.NewTransaction(svcCtx.Db.Conn),
		kafkaDomain:         svcCtx.KafkaDomain,
	}
}

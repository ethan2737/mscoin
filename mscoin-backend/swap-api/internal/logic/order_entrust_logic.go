package logic

import (
	"context"
	"errors"

	"swap-api/internal/dao"
	"swap-api/internal/model"
	"swap-api/internal/svc"
	"swap-api/internal/types"
)

type OrderEntrustLogic struct {
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	orderDao        dao.ContractOrderDao
	positionDao     dao.ContractPositionDao
	walletDao       dao.ContractWalletDao
	contractCoinDao dao.ContractCoinDao
	transactionDao  dao.ContractTransactionDao
}

func NewOrderEntrustLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderEntrustLogic {
	return &OrderEntrustLogic{
		ctx:             ctx,
		svcCtx:          svcCtx,
		orderDao:        dao.NewContractOrderDao(svcCtx.ContractOrderModel),
		positionDao:     dao.NewContractPositionDao(svcCtx.ContractPositionModel),
		walletDao:       dao.NewContractWalletDao(svcCtx.ContractWalletModel),
		contractCoinDao: dao.NewContractCoinDao(svcCtx.ContractCoinModel),
		transactionDao:  dao.NewContractTransactionDao(svcCtx.ContractTransactionModel),
	}
}

func (l *OrderEntrustLogic) Add(req *types.AddOrderReq) (*types.AddOrderResp, error) {
	coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.ContractCoinId))
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("合约不存在")
	}

	if err := l.checkMargin(req); err != nil {
		return nil, err
	}

	order := &model.ContractOrder{
		MemberId:       req.MemberId,
		ContractCoinId: int64(req.ContractCoinId),
		EntrustType:    req.EntrustType,
		Type:           req.Type,
		Direction:      req.Direction,
		Leverage:       req.Leverage,
		Price:          req.Price,
		Amount:         req.Amount,
		Status:         1,
	}

	if err = l.orderDao.Create(l.ctx, order); err != nil {
		return nil, err
	}

	return &types.AddOrderResp{OrderId: order.Id}, nil
}

func (l *OrderEntrustLogic) Cancel(req *types.CancelOrderReq) (*types.CancelOrderResp, error) {
	order, err := l.orderDao.GetByID(l.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}
	if order.Status != 1 {
		return nil, errors.New("订单状态不允许撤销")
	}
	if order.MemberId != req.MemberId {
		return nil, errors.New("无权操作该订单")
	}

	if err = l.orderDao.UpdateStatus(l.ctx, req.OrderId, 3, 0, 0, 0, 0); err != nil {
		return nil, err
	}

	if order.EntrustType == 1 {
		requiredMargin := (order.Price * order.Amount) / float64(order.Leverage)
		wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, order.MemberId, "USDT")
		if err != nil {
			return nil, err
		}
		if wallet != nil {
			newFrozen := wallet.Frozen - requiredMargin
			if newFrozen < 0 {
				newFrozen = 0
			}
			if err = l.walletDao.UpdateBalance(l.ctx, wallet.Id, wallet.Balance, newFrozen, wallet.MainBalance); err != nil {
				return nil, err
			}
		}
	}

	return &types.CancelOrderResp{Success: true}, nil
}

func (l *OrderEntrustLogic) Current(req *types.CurrentOrderReq) (*types.CurrentOrderResp, error) {
	orders, err := l.orderDao.GetByMemberId(l.ctx, req.MemberId, 1)
	if err != nil {
		return nil, err
	}

	coinMap, err := l.loadContractCoinMap()
	if err != nil {
		return nil, err
	}

	filtered := filterOrdersForSwap(orders, req.ContractCoinId, req.Type, false)
	pageNo := int64(req.PageNo)
	pageSize := int64(req.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}
	page := buildUcOrderPage(filtered, coinMap, pageNo, pageSize)
	items := toUcOrderItems(page.Content)
	return &types.CurrentOrderResp{
		List:          items,
		Content:       items,
		TotalElements: page.TotalElements,
		Number:        page.Number,
		TotalPages:    page.TotalPages,
		HasNext:       page.HasNext,
		IsLast:        page.IsLast,
	}, nil
}

func (l *OrderEntrustLogic) History(req *types.HistoryOrderReq) (*types.HistoryOrderResp, error) {
	orders, err := l.orderDao.GetByMemberId(l.ctx, req.MemberId, 0)
	if err != nil {
		return nil, err
	}

	coinMap, err := l.loadContractCoinMap()
	if err != nil {
		return nil, err
	}

	filtered := filterOrdersForSwap(orders, req.ContractCoinId, req.Type, true)
	pageNo := int64(req.PageNo)
	pageSize := int64(req.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}
	page := buildUcOrderPage(filtered, coinMap, pageNo, pageSize)
	items := toUcOrderItems(page.Content)
	return &types.HistoryOrderResp{
		List:          items,
		Content:       items,
		TotalElements: page.TotalElements,
		Number:        page.Number,
		TotalPages:    page.TotalPages,
		HasNext:       page.HasNext,
		IsLast:        page.IsLast,
	}, nil
}

func (l *OrderEntrustLogic) Position(req *types.PositionReq) (*types.PositionResp, error) {
	positions, err := l.positionDao.GetByMemberId(l.ctx, req.MemberId)
	if err != nil {
		return nil, err
	}

	coinMap, err := l.loadContractCoinMap()
	if err != nil {
		return nil, err
	}

	filtered := filterPositionsForSwap(positions, req.ContractCoinId)
	items := buildSwapPositionItems(filtered, coinMap)
	return &types.PositionResp{List: items}, nil
}

func (l *OrderEntrustLogic) QuickClose(req *types.QuickCloseReq) (*types.QuickCloseResp, error) {
	position, err := l.positionDao.GetByMember(l.ctx, req.MemberId, int64(req.ContractCoinId), 0)
	if err != nil {
		return nil, err
	}
	if position == nil || position.Size == 0 {
		position, err = l.positionDao.GetByMember(l.ctx, req.MemberId, int64(req.ContractCoinId), 1)
		if err != nil {
			return nil, err
		}
	}
	if position == nil || position.Size == 0 {
		return nil, errors.New("无持仓")
	}

	coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.ContractCoinId))
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("合约不存在")
	}

	currentPrice := position.EntryPrice
	if req.Price > 0 {
		currentPrice = req.Price
	}

	profit := 0.0
	if position.Direction == 1 {
		profit = (currentPrice - position.EntryPrice) * position.Size
	} else {
		profit = (position.EntryPrice - currentPrice) * position.Size
	}

	order := &model.ContractOrder{
		MemberId:       req.MemberId,
		ContractCoinId: int64(req.ContractCoinId),
		EntrustType:    2,
		Type:           2,
		Direction:      position.Direction,
		Leverage:       position.Leverage,
		Price:          currentPrice,
		Amount:         position.Size,
		Status:         2,
	}
	if err = l.orderDao.Create(l.ctx, order); err != nil {
		return nil, err
	}

	position.Size = 0
	position.UnrealizedPnl = 0
	if err = l.positionDao.Update(l.ctx, position); err != nil {
		return nil, err
	}

	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, "USDT")
	if err != nil {
		return nil, err
	}
	if wallet != nil {
		newBalance := wallet.Balance + profit
		if err = l.walletDao.UpdateBalance(l.ctx, wallet.Id, newBalance, wallet.Frozen, wallet.MainBalance); err != nil {
			return nil, err
		}

		transaction := &model.ContractTransaction{
			MemberId:       req.MemberId,
			Unit:           "USDT",
			Type:           3,
			Amount:         profit,
			Balance:        newBalance,
			RelatedOrderId: order.Id,
			Remark:         "闪电平仓",
		}
		if err = l.transactionDao.Create(l.ctx, transaction); err != nil {
			return nil, err
		}
	}

	return &types.QuickCloseResp{Success: true}, nil
}

func (l *OrderEntrustLogic) Leverage(req *types.LeverageReq) (*types.LeverageResp, error) {
	position, err := l.positionDao.GetByMember(l.ctx, req.MemberId, int64(req.ContractCoinId), 0)
	if err != nil {
		return nil, err
	}
	if position == nil {
		coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.ContractCoinId))
		if err != nil {
			return nil, err
		}
		if coin == nil {
			return nil, errors.New("合约不存在")
		}
		return &types.LeverageResp{Leverage: coin.MinLeverage}, nil
	}
	return &types.LeverageResp{Leverage: position.Leverage}, nil
}

func (l *OrderEntrustLogic) checkMargin(req *types.AddOrderReq) error {
	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, "USDT")
	if err != nil {
		return err
	}

	requiredMargin := (req.Price * req.Amount) / float64(req.Leverage)
	if req.EntrustType == 1 {
		availableBalance := 0.0
		if wallet != nil {
			availableBalance = wallet.Balance
		}
		if availableBalance < requiredMargin {
			return errors.New("保证金不足")
		}
		if wallet != nil {
			newBalance := availableBalance - requiredMargin
			newFrozen := wallet.Frozen + requiredMargin
			if err = l.walletDao.UpdateBalance(l.ctx, wallet.Id, newBalance, newFrozen, wallet.MainBalance); err != nil {
				return err
			}
		}
		return nil
	}

	totalBalance := 0.0
	if wallet != nil {
		totalBalance = wallet.Balance + wallet.Frozen
	}
	if totalBalance < requiredMargin {
		return errors.New("保证金不足")
	}
	return nil
}

func (l *OrderEntrustLogic) loadContractCoinMap() (map[int64]*model.ContractCoin, error) {
	coins, err := l.contractCoinDao.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}

	coinMap := make(map[int64]*model.ContractCoin, len(coins))
	for _, coin := range coins {
		coinMap[coin.Id] = coin
	}
	return coinMap, nil
}

func filterOrdersForSwap(orders []*model.ContractOrder, contractCoinID int32, selectedType int32, historyOnly bool) []*model.ContractOrder {
	filtered := make([]*model.ContractOrder, 0, len(orders))
	for _, order := range orders {
		if historyOnly && order.Status != 2 && order.Status != 3 {
			continue
		}
		if contractCoinID > 0 && order.ContractCoinId != int64(contractCoinID) {
			continue
		}
		if selectedType > 0 && !matchesSwapSelectedType(order, selectedType) {
			continue
		}
		filtered = append(filtered, order)
	}
	return filtered
}

func filterPositionsForSwap(positions []*model.ContractPosition, contractCoinID int32) []*model.ContractPosition {
	filtered := make([]*model.ContractPosition, 0, len(positions))
	for _, position := range positions {
		if contractCoinID > 0 && position.ContractCoinId != int64(contractCoinID) {
			continue
		}
		filtered = append(filtered, position)
	}
	return filtered
}

func matchesSwapSelectedType(order *model.ContractOrder, selectedType int32) bool {
	switch selectedType {
	case 2:
		return order.EntrustType == 1
	case 3:
		return order.EntrustType != 1
	default:
		return true
	}
}

func toUcOrderItems(content []any) []*types.UcOrderItem {
	items := make([]*types.UcOrderItem, 0, len(content))
	for _, item := range content {
		orderItem, ok := item.(*types.UcOrderItem)
		if ok && orderItem != nil {
			items = append(items, orderItem)
		}
	}
	return items
}

package logic

import (
	"context"
	"errors"
	"swap/internal/dao"
	"swap/internal/model"
	"swap/internal/svc"
	"swap/types/swap"
)

type OrderEntrustLogic struct {
	ctx             context.Context
	svc             *svc.ServiceContext
	orderDao        dao.ContractOrderDao
	positionDao     dao.ContractPositionDao
	walletDao       dao.ContractWalletDao
	contractCoinDao dao.ContractCoinDao
	transactionDao  dao.ContractTransactionDao
}

func NewOrderEntrustLogic(ctx context.Context, svc *svc.ServiceContext) *OrderEntrustLogic {
	return &OrderEntrustLogic{
		ctx:             ctx,
		svc:             svc,
		orderDao:        dao.NewContractOrderDao(svc.ContractOrderModel),
		positionDao:     dao.NewContractPositionDao(svc.ContractPositionModel),
		walletDao:       dao.NewContractWalletDao(svc.ContractWalletModel),
		contractCoinDao: dao.NewContractCoinDao(svc.ContractCoinModel),
		transactionDao:  dao.NewContractTransactionDao(svc.ContractTransactionModel),
	}
}

func (l *OrderEntrustLogic) AddOrder(req *swap.AddOrderRequest) (*swap.AddOrderResponse, error) {
	// 1. 获取合约信息
	coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.ContractCoinId))
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("合约不存在")
	}

	// 2. 风控校验：保证金是否足够
	if err := l.checkMargin(req); err != nil {
		return nil, err
	}

	// 3. 创建订单
	order := &model.ContractOrder{
		MemberId:       req.MemberId,
		ContractCoinId: int64(req.ContractCoinId),
		EntrustType:    req.EntrustType,
		Type:           req.Type,
		Direction:      req.Direction,
		Leverage:       req.Leverage,
		Price:          req.Price,
		Amount:         req.Amount,
		Status:         1, // 待成交
	}

	// 4. 插入订单到数据库
	err = l.orderDao.Create(l.ctx, order)
	if err != nil {
		return nil, err
	}

	return &swap.AddOrderResponse{
		OrderId: order.Id,
	}, nil
}

func (l *OrderEntrustLogic) CancelOrder(req *swap.CancelOrderRequest) (*swap.CancelOrderResponse, error) {
	// 1. 获取订单信息
	order, err := l.orderDao.GetByID(l.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 校验订单状态：只有待成交的订单可以撤销
	if order.Status != 1 {
		return nil, errors.New("订单状态不允许撤销")
	}

	// 3. 校验订单归属
	if order.MemberId != req.MemberId {
		return nil, errors.New("无权操作该订单")
	}

	// 4. 更新订单状态为已取消
	err = l.orderDao.UpdateStatus(l.ctx, req.OrderId, 3, 0, 0, 0, 0) // 3 = 已取消
	if err != nil {
		return nil, err
	}

	// 5. 如果是限价单且有冻结资金，需要解冻
	if order.EntrustType == 1 { // 1 = 限价单
		// 计算需要解冻的保证金
		requiredMargin := (order.Price * order.Amount) / float64(order.Leverage)

		// 获取用户钱包
		wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, order.MemberId, "USDT")
		if err != nil {
			return nil, err
		}
		if wallet != nil {
			// 解冻资金
			newFrozen := wallet.Frozen - requiredMargin
			if newFrozen < 0 {
				newFrozen = 0
			}
			err = l.walletDao.UpdateBalance(l.ctx, wallet.Id, wallet.Balance, newFrozen, wallet.MainBalance)
			if err != nil {
				return nil, err
			}
		}
	}

	return &swap.CancelOrderResponse{Success: true}, nil
}

func (l *OrderEntrustLogic) QuickClose(req *swap.QuickCloseRequest) (*swap.QuickCloseResponse, error) {
	// 1. 获取用户持仓
	position, err := l.positionDao.GetByMember(l.ctx, req.MemberId, int64(req.ContractCoinId), 0)
	if err != nil {
		return nil, err
	}
	// 如果没有同方向持仓，尝试反方向
	if position == nil || position.Size == 0 {
		position, err = l.positionDao.GetByMember(l.ctx, req.MemberId, int64(req.ContractCoinId), 1)
		if err != nil {
			return nil, err
		}
	}
	if position == nil || position.Size == 0 {
		return nil, errors.New("无持仓")
	}

	// 2. 获取合约信息
	coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.ContractCoinId))
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("合约不存在")
	}

	// 3. 获取当前市场价格（这里需要从市场服务获取，暂时用持仓均价）
	currentPrice := position.EntryPrice
	if req.Price > 0 {
		currentPrice = req.Price
	}

	// 4. 计算盈亏
	var profit float64
	if position.Direction == 1 { // 做多
		profit = (currentPrice - position.EntryPrice) * position.Size
	} else { // 做空
		profit = (position.EntryPrice - currentPrice) * position.Size
	}

	// 5. 创建平仓订单
	order := &model.ContractOrder{
		MemberId:       req.MemberId,
		ContractCoinId: int64(req.ContractCoinId),
		EntrustType:    2, // 2 = 市价单
		Type:           2, // 2 = 平仓
		Direction:      position.Direction,
		Leverage:       position.Leverage,
		Price:          currentPrice,
		Amount:         position.Size,
		Status:         2, // 2 = 已成交
	}

	err = l.orderDao.Create(l.ctx, order)
	if err != nil {
		return nil, err
	}

	// 6. 更新持仓为 0
	position.Size = 0
	position.UnrealizedPnl = 0
	err = l.positionDao.Update(l.ctx, position)
	if err != nil {
		return nil, err
	}

	// 7. 更新用户钱包余额
	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, "USDT")
	if err != nil {
		return nil, err
	}
	if wallet != nil {
		newBalance := wallet.Balance + profit
		err = l.walletDao.UpdateBalance(l.ctx, wallet.Id, newBalance, wallet.Frozen, wallet.MainBalance)
		if err != nil {
			return nil, err
		}

		// 8. 记录交易流水
		transaction := &model.ContractTransaction{
			MemberId:       req.MemberId,
			Unit:           "USDT",
			Type:           3, // 3 = 平仓盈亏
			Amount:         profit,
			Balance:        newBalance,
			RelatedOrderId: order.Id,
			Remark:         "闪电平仓",
		}
		err = l.transactionDao.Create(l.ctx, transaction)
		if err != nil {
			return nil, err
		}
	}

	return &swap.QuickCloseResponse{
		Success: true,
	}, nil
}

func (l *OrderEntrustLogic) checkMargin(req *swap.AddOrderRequest) error {
	// 1. 获取用户钱包余额
	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, "USDT")
	if err != nil {
		return err
	}

	// 2. 计算需要的保证金
	// 保证金 = (价格 * 数量) / 杠杆
	requiredMargin := (req.Price * req.Amount) / float64(req.Leverage)

	// 3. 如果是限价单，需要冻结保证金
	if req.EntrustType == 1 { // 1 = 限价单
		// 检查可用余额是否足够
		availableBalance := 0.0
		if wallet != nil {
			availableBalance = wallet.Balance
		}

		if availableBalance < requiredMargin {
			return errors.New("保证金不足")
		}

		// 冻结保证金
		if wallet != nil {
			newBalance := availableBalance - requiredMargin
			newFrozen := wallet.Frozen + requiredMargin
			err = l.walletDao.UpdateBalance(l.ctx, wallet.Id, newBalance, newFrozen, wallet.MainBalance)
			if err != nil {
				return err
			}
		}
	} else { // 市价单
		// 检查总余额是否足够
		totalBalance := 0.0
		if wallet != nil {
			totalBalance = wallet.Balance + wallet.Frozen
		}

		if totalBalance < requiredMargin {
			return errors.New("保证金不足")
		}
	}

	return nil
}

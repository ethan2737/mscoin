package logic

import (
	"context"
	"errors"
	"swap/internal/dao"
	"swap/internal/svc"
	"swap/types/swap"
)

type PositionLogic struct {
	ctx              context.Context
	svc              *svc.ServiceContext
	positionDao      dao.ContractPositionDao
	contractCoinDao  dao.ContractCoinDao
	orderDao         dao.ContractOrderDao
}

func NewPositionLogic(ctx context.Context, svc *svc.ServiceContext) *PositionLogic {
	return &PositionLogic{
		ctx:              ctx,
		svc:              svc,
		positionDao:      dao.NewContractPositionDao(svc.ContractPositionModel),
		contractCoinDao:  dao.NewContractCoinDao(svc.ContractCoinModel),
		orderDao:         dao.NewContractOrderDao(svc.ContractOrderModel),
	}
}

func (l *PositionLogic) GetPosition(req *swap.GetPositionRequest) (*swap.PositionResponse, error) {
	positions, err := l.positionDao.GetByMemberId(l.ctx, req.MemberId)
	if err != nil {
		return nil, err
	}
	list := make([]*swap.PositionInfo, 0)
	for _, pos := range positions {
		// 计算未实现盈亏
		unrealizedPnl := l.calculateUnrealizedPnl(pos)
		list = append(list, &swap.PositionInfo{
			Id:                pos.Id,
			MemberId:          pos.MemberId,
			ContractCoinId:    int32(pos.ContractCoinId),
			Direction:         pos.Direction,
			Leverage:          pos.Leverage,
			Size:              pos.Size,
			EntryPrice:        pos.EntryPrice,
			Margin:            pos.Margin,
			UnrealizedPnl:     unrealizedPnl,
			LiquidationPrice:  pos.LiquidationPrice,
		})
	}
	return &swap.PositionResponse{List: list}, nil
}

func (l *PositionLogic) calculateUnrealizedPnl(pos interface{}) float64 {
	// TODO: 实现未实现盈亏计算
	// 需要获取当前市场价格
	return 0.0
}

func (l *PositionLogic) GetCurrentOrder(req *swap.GetCurrentOrderRequest) (*swap.CurrentOrderResponse, error) {
	// 获取用户的当前委托订单（待成交状态）
	orders, err := l.orderDao.GetByMemberId(l.ctx, req.MemberId, 1) // 1 = 待成交
	if err != nil {
		return nil, err
	}
	list := make([]*swap.OrderInfo, 0)
	for _, order := range orders {
		list = append(list, &swap.OrderInfo{
			Id:             order.Id,
			MemberId:       order.MemberId,
			ContractCoinId: int32(order.ContractCoinId),
			EntrustType:    order.EntrustType,
			Type:           order.Type,
			Direction:      order.Direction,
			Leverage:       order.Leverage,
			Price:          order.Price,
			Amount:         order.Amount,
			Status:         order.Status,
		})
	}
	return &swap.CurrentOrderResponse{List: list}, nil
}

func (l *PositionLogic) GetHistoryOrder(req *swap.GetHistoryOrderRequest) (*swap.HistoryOrderResponse, error) {
	// 获取用户的历史订单（已成交和已取消状态）
	orders, err := l.orderDao.GetByMemberId(l.ctx, req.MemberId, 0) // 0 = 全部
	if err != nil {
		return nil, err
	}
	list := make([]*swap.OrderInfo, 0)
	for _, order := range orders {
		// 过滤出已成交和已取消的订单
		if order.Status == 2 || order.Status == 3 {
			list = append(list, &swap.OrderInfo{
				Id:             order.Id,
				MemberId:       order.MemberId,
				ContractCoinId: int32(order.ContractCoinId),
				EntrustType:    order.EntrustType,
				Type:           order.Type,
				Direction:      order.Direction,
				Leverage:       order.Leverage,
				Price:          order.Price,
				Amount:         order.Amount,
				Status:         order.Status,
			})
		}
	}
	return &swap.HistoryOrderResponse{List: list}, nil
}

func (l *PositionLogic) GetLeverage(req *swap.GetLeverageRequest) (*swap.LeverageResponse, error) {
	// 获取用户的杠杆倍数设置
	// 这里可以从持仓中获取，或者使用默认值
	position, err := l.positionDao.GetByMember(l.ctx, req.MemberId, int64(req.ContractCoinId), 0)
	if err != nil {
		return nil, err
	}
	if position == nil {
		// 没有持仓时，返回合约的默认杠杆
		coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.ContractCoinId))
		if err != nil {
			return nil, err
		}
		if coin == nil {
			return nil, errors.New("合约不存在")
		}
		return &swap.LeverageResponse{
			Leverage: coin.MinLeverage, // 默认使用最小杠杆
		}, nil
	}
	return &swap.LeverageResponse{
		Leverage: position.Leverage,
	}, nil
}

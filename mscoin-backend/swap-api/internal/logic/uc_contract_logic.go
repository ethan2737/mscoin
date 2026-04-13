package logic

import (
	"context"
	"errors"

	"mscoin-common/pages"
	"swap-api/internal/dao"
	"swap-api/internal/model"
	"swap-api/internal/svc"
	"swap-api/internal/types"
)

type UcContractLogic struct {
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	orderDao        dao.ContractOrderDao
	positionDao     dao.ContractPositionDao
	contractCoinDao dao.ContractCoinDao
}

func NewUcContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UcContractLogic {
	return &UcContractLogic{
		ctx:             ctx,
		svcCtx:          svcCtx,
		orderDao:        dao.NewContractOrderDao(svcCtx.ContractOrderModel),
		positionDao:     dao.NewContractPositionDao(svcCtx.ContractPositionModel),
		contractCoinDao: dao.NewContractCoinDao(svcCtx.ContractCoinModel),
	}
}

func (l *UcContractLogic) Current(req *types.CurrentOrderReq) (*pages.PageResult, error) {
	orders, err := l.orderDao.GetByMemberId(l.ctx, req.MemberId, 1)
	if err != nil {
		return nil, err
	}

	coinMap, err := l.loadContractCoinMap()
	if err != nil {
		return nil, err
	}

	pageNo := int64(req.PageNo)
	pageSize := int64(req.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}
	filtered := filterOrdersForSwap(orders, req.ContractCoinId, req.Type, false)
	return buildUcOrderPage(filtered, coinMap, pageNo, pageSize), nil
}

func (l *UcContractLogic) History(req *types.HistoryOrderReq) (*pages.PageResult, error) {
	orders, err := l.orderDao.GetByMemberId(l.ctx, req.MemberId, 0)
	if err != nil {
		return nil, err
	}

	coinMap, err := l.loadContractCoinMap()
	if err != nil {
		return nil, err
	}

	pageNo := int64(req.PageNo)
	pageSize := int64(req.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}
	filtered := filterOrdersForSwap(orders, req.ContractCoinId, req.Type, true)
	return buildUcOrderPage(filtered, coinMap, pageNo, pageSize), nil
}

func (l *UcContractLogic) Thumb(req *types.GetContractListReq) (*[]*types.ContractCoinThumbItem, error) {
	coins, err := l.contractCoinDao.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}

	list := buildThumbItems(coins)
	return &list, nil
}

func (l *UcContractLogic) Cancel(req *types.CancelOrderReq) (*types.CancelOrderResp, error) {
	order, err := l.orderDao.GetByID(l.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}
	if order.MemberId != req.MemberId {
		return nil, errors.New("无权操作该订单")
	}
	if order.Status != 1 {
		return nil, errors.New("订单状态不允许撤销")
	}
	if err := l.orderDao.UpdateStatus(l.ctx, req.OrderId, 3, 0, 0, 0, 0); err != nil {
		return nil, err
	}

	return &types.CancelOrderResp{Success: true}, nil
}

func (l *UcContractLogic) loadContractCoinMap() (map[int64]*model.ContractCoin, error) {
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

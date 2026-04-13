package logic

import (
	"context"
	"swap-api/internal/dao"
	"swap-api/internal/svc"
	"swap-api/internal/types"
)

type ContractInfoLogic struct {
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	contractCoinDao dao.ContractCoinDao
}

func NewContractInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContractInfoLogic {
	return &ContractInfoLogic{
		ctx:             ctx,
		svcCtx:          svcCtx,
		contractCoinDao: dao.NewContractCoinDao(svcCtx.ContractCoinModel),
	}
}

func (l *ContractInfoLogic) Info(req *types.GetContractInfoReq) (*types.ContractInfoResp, error) {
	coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, nil
	}
	return &types.ContractInfoResp{
		Id:                int32(coin.Id),
		Symbol:            coin.Symbol,
		ContractType:      coin.ContractType,
		PricePrecision:    coin.PricePrecision,
		QuantityPrecision: coin.QuantityPrecision,
		ShareNumber:       coin.ShareNumber,
		MakerFee:          coin.MakerFee,
		TakerFee:          coin.TakerFee,
		MinLeverage:       coin.MinLeverage,
		MaxLeverage:       coin.MaxLeverage,
		Status:            coin.Status,
	}, nil
}

func (l *ContractInfoLogic) List(req *types.GetContractListReq) (*[]*types.ContractInfoResp, error) {
	coins, err := l.contractCoinDao.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*types.ContractInfoResp, 0, len(coins))
	for _, coin := range coins {
		list = append(list, &types.ContractInfoResp{
			Id:                int32(coin.Id),
			Symbol:            coin.Symbol,
			ContractType:      coin.ContractType,
			PricePrecision:    coin.PricePrecision,
			QuantityPrecision: coin.QuantityPrecision,
			ShareNumber:       coin.ShareNumber,
			MakerFee:          coin.MakerFee,
			TakerFee:          coin.TakerFee,
			MinLeverage:       coin.MinLeverage,
			MaxLeverage:       coin.MaxLeverage,
			Status:            coin.Status,
		})
	}
	return &list, nil
}

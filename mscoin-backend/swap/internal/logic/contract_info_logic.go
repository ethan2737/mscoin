package logic

import (
	"context"
	"swap/internal/dao"
	"swap/internal/svc"
	"swap/types/swap"
)

type ContractInfoLogic struct {
	ctx             context.Context
	svc             *svc.ServiceContext
	contractCoinDao dao.ContractCoinDao
}

func NewContractInfoLogic(ctx context.Context, svc *svc.ServiceContext) *ContractInfoLogic {
	return &ContractInfoLogic{
		ctx:             ctx,
		svc:             svc,
		contractCoinDao: dao.NewContractCoinDao(svc.ContractCoinModel),
	}
}

func (l *ContractInfoLogic) GetContractInfo(req *swap.GetContractInfoRequest) (*swap.ContractInfoResponse, error) {
	coin, err := l.contractCoinDao.GetByID(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, nil
	}
	return &swap.ContractInfoResponse{
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

func (l *ContractInfoLogic) GetContractList(req *swap.GetContractListRequest) (*swap.ContractListResponse, error) {
	coins, err := l.contractCoinDao.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*swap.ContractInfoResponse, 0, len(coins))
	for _, coin := range coins {
		list = append(list, &swap.ContractInfoResponse{
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
	return &swap.ContractListResponse{List: list}, nil
}

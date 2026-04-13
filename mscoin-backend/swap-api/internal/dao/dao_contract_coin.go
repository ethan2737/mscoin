package dao

import (
	"context"
	"swap-api/internal/model"
)

type ContractCoinDao struct {
	model model.ContractCoinModel
}

func NewContractCoinDao(model model.ContractCoinModel) ContractCoinDao {
	return ContractCoinDao{model: model}
}

func (d *ContractCoinDao) Create(ctx context.Context, coin *model.ContractCoin) error {
	_, err := d.model.Insert(ctx, coin)
	return err
}

func (d *ContractCoinDao) GetByID(ctx context.Context, id int64) (*model.ContractCoin, error) {
	return d.model.FindOne(ctx, id)
}

func (d *ContractCoinDao) GetBySymbol(ctx context.Context, symbol string) (*model.ContractCoin, error) {
	return d.model.FindBySymbol(ctx, symbol)
}

func (d *ContractCoinDao) GetAll(ctx context.Context) ([]*model.ContractCoin, error) {
	return d.model.FindAll(ctx)
}

func (d *ContractCoinDao) Update(ctx context.Context, coin *model.ContractCoin) error {
	return d.model.Update(ctx, coin)
}

func (d *ContractCoinDao) Delete(ctx context.Context, id int64) error {
	return d.model.Delete(ctx, id)
}

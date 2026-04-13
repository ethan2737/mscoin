package dao

import (
	"context"
	"swap/internal/model"
)

type ContractOrderDao struct {
	model model.ContractOrderModel
}

func NewContractOrderDao(model model.ContractOrderModel) ContractOrderDao {
	return ContractOrderDao{model: model}
}

func (d *ContractOrderDao) Create(ctx context.Context, order *model.ContractOrder) error {
	_, err := d.model.Insert(ctx, order)
	return err
}

func (d *ContractOrderDao) GetByID(ctx context.Context, id int64) (*model.ContractOrder, error) {
	return d.model.FindOne(ctx, id)
}

func (d *ContractOrderDao) GetByMemberId(ctx context.Context, memberId int64, status int32) ([]*model.ContractOrder, error) {
	return d.model.FindByMemberId(ctx, memberId, status)
}

func (d *ContractOrderDao) UpdateStatus(ctx context.Context, id int64, status int32, dealAmount float64, dealAmountUsdt float64, fee float64, profit float64) error {
	return d.model.UpdateStatus(ctx, id, status, dealAmount, dealAmountUsdt, fee, profit)
}

func (d *ContractOrderDao) Delete(ctx context.Context, id int64) error {
	return d.model.Delete(ctx, id)
}

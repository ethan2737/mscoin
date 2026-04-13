package dao

import (
	"context"
	"swap-api/internal/model"
)

type ContractWalletDao struct {
	model model.ContractWalletModel
}

func NewContractWalletDao(model model.ContractWalletModel) ContractWalletDao {
	return ContractWalletDao{model: model}
}

func (d *ContractWalletDao) Create(ctx context.Context, wallet *model.ContractWallet) error {
	_, err := d.model.Insert(ctx, wallet)
	return err
}

func (d *ContractWalletDao) GetByID(ctx context.Context, id int64) (*model.ContractWallet, error) {
	return d.model.FindOne(ctx, id)
}

func (d *ContractWalletDao) GetByMemberId(ctx context.Context, memberId int64) ([]*model.ContractWallet, error) {
	return d.model.FindByMemberId(ctx, memberId)
}

func (d *ContractWalletDao) GetByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*model.ContractWallet, error) {
	return d.model.FindByMemberAndUnit(ctx, memberId, unit)
}

func (d *ContractWalletDao) UpdateBalance(ctx context.Context, id int64, balance float64, frozen float64, mainBalance float64) error {
	return d.model.UpdateBalance(ctx, id, balance, frozen, mainBalance)
}

func (d *ContractWalletDao) Upsert(ctx context.Context, wallet *model.ContractWallet) error {
	return d.model.Upsert(ctx, wallet)
}

func (d *ContractWalletDao) Delete(ctx context.Context, id int64) error {
	return d.model.Delete(ctx, id)
}

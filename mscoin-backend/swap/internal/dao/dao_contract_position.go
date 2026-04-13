package dao

import (
	"context"
	"swap/internal/model"
)

type ContractPositionDao struct {
	model model.ContractPositionModel
}

func NewContractPositionDao(model model.ContractPositionModel) ContractPositionDao {
	return ContractPositionDao{model: model}
}

func (d *ContractPositionDao) Create(ctx context.Context, position *model.ContractPosition) error {
	_, err := d.model.Insert(ctx, position)
	return err
}

func (d *ContractPositionDao) GetByID(ctx context.Context, id int64) (*model.ContractPosition, error) {
	return d.model.FindOne(ctx, id)
}

func (d *ContractPositionDao) GetByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*model.ContractPosition, error) {
	return d.model.FindByMember(ctx, memberId, contractCoinId, direction)
}

func (d *ContractPositionDao) GetByMemberId(ctx context.Context, memberId int64) ([]*model.ContractPosition, error) {
	return d.model.FindByMemberId(ctx, memberId)
}

func (d *ContractPositionDao) Update(ctx context.Context, position *model.ContractPosition) error {
	return d.model.Update(ctx, position)
}

func (d *ContractPositionDao) Upsert(ctx context.Context, position *model.ContractPosition) error {
	return d.model.Upsert(ctx, position)
}

func (d *ContractPositionDao) Delete(ctx context.Context, id int64) error {
	return d.model.Delete(ctx, id)
}

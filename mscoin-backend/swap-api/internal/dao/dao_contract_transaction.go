package dao

import (
	"context"
	"time"
	"swap-api/internal/model"
)

type ContractTransactionDao struct {
	model model.ContractTransactionModel
}

func NewContractTransactionDao(model model.ContractTransactionModel) ContractTransactionDao {
	return ContractTransactionDao{model: model}
}

func (d *ContractTransactionDao) Create(ctx context.Context, tx *model.ContractTransaction) error {
	_, err := d.model.Insert(ctx, tx)
	return err
}

func (d *ContractTransactionDao) GetByID(ctx context.Context, id int64) (*model.ContractTransaction, error) {
	return d.model.FindOne(ctx, id)
}

func (d *ContractTransactionDao) GetByMemberId(ctx context.Context, memberId int64, startTime, endTime time.Time, txType int32, limit, offset int64) ([]*model.ContractTransaction, error) {
	return d.model.FindByMemberId(ctx, memberId, startTime, endTime, txType, limit, offset)
}

func (d *ContractTransactionDao) Delete(ctx context.Context, id int64) error {
	return d.model.Delete(ctx, id)
}

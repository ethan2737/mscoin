package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ContractTransaction struct {
	Id             int64     `db:"id"`
	MemberId       int64     `db:"member_id"`
	Unit           string    `db:"unit"`
	Type           int32     `db:"type"`
	Amount         float64   `db:"amount"`
	Balance        float64   `db:"balance"`
	RelatedOrderId int64     `db:"related_order_id"`
	Remark         string    `db:"remark"`
	CreatedAt      time.Time `db:"created_at"`
}

type ContractTransactionModel interface {
	Insert(ctx context.Context, tx *ContractTransaction) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractTransaction, error)
	FindByMemberId(ctx context.Context, memberId int64, startTime, endTime time.Time, txType int32, limit, offset int64) ([]*ContractTransaction, error)
	Delete(ctx context.Context, id int64) error
}

type defaultContractTransactionModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewContractTransactionModel(conn sqlx.SqlConn) ContractTransactionModel {
	return &defaultContractTransactionModel{
		conn:  conn,
		table: "contract_transactions",
	}
}

func (m *defaultContractTransactionModel) Insert(ctx context.Context, tx *ContractTransaction) (int64, error) {
	query := `INSERT INTO contract_transactions (member_id, unit, type, amount, balance, related_order_id, remark) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := m.conn.ExecCtx(ctx, query, tx.MemberId, tx.Unit, tx.Type, tx.Amount, tx.Balance, tx.RelatedOrderId, tx.Remark)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	tx.Id = id
	return id, nil
}

func (m *defaultContractTransactionModel) FindOne(ctx context.Context, id int64) (*ContractTransaction, error) {
	query := `SELECT id, member_id, unit, type, amount, balance, IFNULL(related_order_id, 0) AS related_order_id, remark, created_at FROM contract_transactions WHERE id = ?`
	var resp ContractTransaction
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractTransactionModel) FindByMemberId(ctx context.Context, memberId int64, startTime, endTime time.Time, txType int32, limit, offset int64) ([]*ContractTransaction, error) {
	query := `SELECT id, member_id, unit, type, amount, balance, IFNULL(related_order_id, 0) AS related_order_id, remark, created_at FROM contract_transactions WHERE member_id = ?`
	args := []interface{}{memberId}

	if !startTime.IsZero() {
		query += ` AND created_at >= ?`
		args = append(args, startTime)
	}
	if !endTime.IsZero() {
		query += ` AND created_at <= ?`
		args = append(args, endTime)
	}
	if txType != -1 {
		query += ` AND type = ?`
		args = append(args, txType)
	}

	query += ` ORDER BY created_at DESC LIMIT ? OFFSET ?`
	args = append(args, limit, offset)

	var resp []*ContractTransaction
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractTransaction{}, nil
		}
		return nil, err
	}
	return resp, nil
}

func (m *defaultContractTransactionModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contract_transactions WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

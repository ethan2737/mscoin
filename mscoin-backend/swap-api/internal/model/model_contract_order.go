package model

import (
	"context"
	"time"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ContractOrder struct {
	Id             int64     `db:"id"`
	MemberId       int64     `db:"member_id"`
	ContractCoinId int64     `db:"contract_coin_id"`
	EntrustType    int32     `db:"entrust_type"`
	Type           int32     `db:"type"`
	Direction      int32     `db:"direction"`
	Leverage       int32     `db:"leverage"`
	Price          float64   `db:"price"`
	Amount         float64   `db:"amount"`
	DealAmount     float64   `db:"deal_amount"`
	Status         int32     `db:"status"`
	DealAmountUsdt float64   `db:"deal_amount_usdt"`
	Fee            float64   `db:"fee"`
	FeeUnit        string    `db:"fee_unit"`
	Profit         float64   `db:"profit"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type ContractOrderModel interface {
	Insert(ctx context.Context, order *ContractOrder) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractOrder, error)
	FindByMemberId(ctx context.Context, memberId int64, status int32) ([]*ContractOrder, error)
	FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*ContractOrder, error)
	UpdateStatus(ctx context.Context, id int64, status int32, dealAmount float64, dealAmountUsdt float64, fee float64, profit float64) error
	Delete(ctx context.Context, id int64) error
}

type defaultContractOrderModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewContractOrderModel(conn sqlx.SqlConn) ContractOrderModel {
	return &defaultContractOrderModel{
		conn:  conn,
		table: "contract_orders",
	}
}

func (m *defaultContractOrderModel) Insert(ctx context.Context, order *ContractOrder) (int64, error) {
	query := `INSERT INTO contract_orders (member_id, contract_coin_id, entrust_type, type, direction, leverage, price, amount, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := m.conn.ExecCtx(ctx, query, order.MemberId, order.ContractCoinId, order.EntrustType, order.Type, order.Direction, order.Leverage, order.Price, order.Amount, order.Status)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	order.Id = id
	return id, nil
}

func (m *defaultContractOrderModel) FindOne(ctx context.Context, id int64) (*ContractOrder, error) {
	query := `SELECT id, member_id, contract_coin_id, entrust_type, type, direction, leverage, price, amount, deal_amount, status, deal_amount_usdt, fee, fee_unit, profit, created_at, updated_at FROM contract_orders WHERE id = ?`
	var resp ContractOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractOrderModel) FindByMemberId(ctx context.Context, memberId int64, status int32) ([]*ContractOrder, error) {
	query := `SELECT id, member_id, contract_coin_id, entrust_type, type, direction, leverage, price, amount, deal_amount, status, deal_amount_usdt, fee, fee_unit, profit, created_at, updated_at FROM contract_orders WHERE member_id = ? AND status = ?`
	var resp []*ContractOrder
	err := m.conn.QueryRowsCtx(ctx, &resp, query, memberId, status)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractOrder{}, nil
		}
		return nil, err
	}
	return resp, nil
}

func (m *defaultContractOrderModel) FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*ContractOrder, error) {
	query := `SELECT id, member_id, contract_coin_id, entrust_type, type, direction, leverage, price, amount, deal_amount, status, deal_amount_usdt, fee, fee_unit, profit, created_at, updated_at FROM contract_orders WHERE contract_coin_id = ?`
	var resp []*ContractOrder
	err := m.conn.QueryRowsCtx(ctx, &resp, query, contractCoinId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractOrder{}, nil
		}
		return nil, err
	}
	return resp, nil
}

func (m *defaultContractOrderModel) UpdateStatus(ctx context.Context, id int64, status int32, dealAmount float64, dealAmountUsdt float64, fee float64, profit float64) error {
	query := `UPDATE contract_orders SET status=?, deal_amount=?, deal_amount_usdt=?, fee=?, profit=? WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, status, dealAmount, dealAmountUsdt, fee, profit, id)
	return err
}

func (m *defaultContractOrderModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contract_orders WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
